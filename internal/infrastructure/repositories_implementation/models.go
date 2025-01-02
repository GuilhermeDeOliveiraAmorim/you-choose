package repositories_implementation

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"gorm.io/gorm"
)

type Lists struct {
	ID            string     `gorm:"primaryKey;not null"`
	Active        bool       `gorm:"not null"`
	CreatedAt     time.Time  `gorm:"not null"`
	UpdatedAt     time.Time  `gorm:"not null"`
	DeactivatedAt *time.Time `gorm:"default:NULL"`
	Name          string     `gorm:"not null"`
	Movies        []Movies   `gorm:"many2many:list_movies;"`
}

func (m *Lists) ToEntity(movies []entities.Movie, combinations []entities.Combination) *entities.List {
	return &entities.List{
		SharedEntity: entities.SharedEntity{
			ID:            m.ID,
			Active:        m.Active,
			CreatedAt:     m.CreatedAt,
			UpdatedAt:     m.UpdatedAt,
			DeactivatedAt: m.DeactivatedAt,
		},
		Name:         m.Name,
		Movies:       movies,
		Combinations: combinations,
	}
}

type Movies struct {
	ID            string         `gorm:"primaryKey;not null"`
	Active        bool           `gorm:"not null"`
	CreatedAt     time.Time      `gorm:"not null"`
	UpdatedAt     time.Time      `gorm:"not null"`
	DeactivatedAt *time.Time     `gorm:"default:NULL"`
	Name          string         `gorm:"not null"`
	Year          int64          `gorm:"not null"`
	Poster        string         `gorm:"not null"`
	ExternalID    string         `gorm:"not null"`
	Lists         []Lists        `gorm:"many2many:list_movies;"`
	FirstVotes    []Combinations `gorm:"foreignKey:FirstMovieID"`
	SecondVotes   []Combinations `gorm:"foreignKey:SecondMovieID"`
	WinnerVotes   []Votes        `gorm:"foreignKey:WinnerID"`
}

func (m *Movies) ToEntity() *entities.Movie {
	return &entities.Movie{
		SharedEntity: entities.SharedEntity{
			ID:            m.ID,
			Active:        m.Active,
			CreatedAt:     m.CreatedAt,
			UpdatedAt:     m.UpdatedAt,
			DeactivatedAt: m.DeactivatedAt,
		},
		Name:       m.Name,
		Year:       m.Year,
		Poster:     m.Poster,
		ExternalID: m.ExternalID,
	}
}

type ListMovies struct {
	ListID        string     `gorm:"primaryKey"`
	List          Lists      `gorm:"foreignKey:ListID"`
	MovieID       string     `gorm:"primaryKey"`
	Movie         Movies     `gorm:"foreignKey:MovieID"`
	CreatedAt     time.Time  `gorm:"not null"`
	DeactivatedAt *time.Time `gorm:"default:NULL"`
}

type Votes struct {
	ID            string       `gorm:"primaryKey;not null"`
	Active        bool         `gorm:"not null"`
	CreatedAt     time.Time    `gorm:"not null"`
	DeactivatedAt *time.Time   `gorm:"default:NULL"`
	UserID        string       `gorm:"not null"`
	CombinationID string       `gorm:"not null"`
	Combination   Combinations `gorm:"foreignKey:CombinationID"`
	WinnerID      string       `gorm:"not null"`
	Winner        Movies       `gorm:"foreignKey:WinnerID"`
}

func (v *Votes) ToEntity() *entities.Vote {
	return &entities.Vote{
		ID:            v.ID,
		Active:        v.Active,
		CreatedAt:     v.CreatedAt,
		DeactivatedAt: v.DeactivatedAt,
		UserID:        v.UserID,
		CombinationID: v.CombinationID,
		WinnerID:      v.WinnerID,
	}
}

type Combinations struct {
	ID            string `gorm:"primaryKey;not null"`
	ListID        string `gorm:"not null"`
	List          Lists  `gorm:"foreignKey:ListID"`
	FirstMovieID  string `gorm:"not null"`
	FirstMovie    Movies `gorm:"foreignKey:FirstMovieID"`
	SecondMovieID string `gorm:"not null"`
	SecondMovie   Movies `gorm:"foreignKey:SecondMovieID"`
}

func (c *Combinations) ToEntity() *entities.Combination {
	return &entities.Combination{
		ID:            c.ID,
		ListID:        c.ListID,
		FirstMovieID:  c.FirstMovieID,
		SecondMovieID: c.SecondMovieID,
	}
}

func Migration(db *gorm.DB, sqlDB *sql.DB) {
	if err := db.AutoMigrate(
		Lists{},
		Movies{},
		ListMovies{},
		Votes{},
		Combinations{},
	); err != nil {
		fmt.Println("Error during migration:", err)
		return
	}
	fmt.Println("Successful migration")
}
