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
	UpdatedAt     *time.Time `gorm:"default:NULL"`
	DeactivatedAt *time.Time `gorm:"default:NULL"`
	Name          string     `gorm:"not null"`
	Cover         string     `gorm:"not null"`
	Movies        []Movies   `gorm:"many2many:list_movies;"`
}

func (m *Lists) ToEntity(movies []entities.Movie, combinations []entities.Combination, complete bool) *entities.List {
	if complete {
		return &entities.List{
			SharedEntity: entities.SharedEntity{
				ID:            m.ID,
				Active:        m.Active,
				CreatedAt:     m.CreatedAt,
				UpdatedAt:     m.UpdatedAt,
				DeactivatedAt: m.DeactivatedAt,
			},
			Name:         m.Name,
			Cover:        m.Cover,
			Movies:       movies,
			Combinations: combinations,
		}
	}

	return &entities.List{
		SharedEntity: entities.SharedEntity{
			ID:            m.ID,
			Active:        m.Active,
			CreatedAt:     m.CreatedAt,
			UpdatedAt:     m.UpdatedAt,
			DeactivatedAt: m.DeactivatedAt,
		},
		Name:  m.Name,
		Cover: m.Cover,
	}
}

type Movies struct {
	ID            string         `gorm:"primaryKey;not null"`
	Active        bool           `gorm:"not null"`
	CreatedAt     time.Time      `gorm:"not null"`
	UpdatedAt     *time.Time     `gorm:"default:NULL"`
	DeactivatedAt *time.Time     `gorm:"default:NULL"`
	Name          string         `gorm:"not null"`
	Year          int64          `gorm:"not null"`
	Poster        string         `gorm:"not null"`
	ExternalID    string         `gorm:"not null"`
	Lists         []Lists        `gorm:"many2many:list_movies;"`
	FirstVotes    []Combinations `gorm:"foreignKey:FirstMovieID"`
	SecondVotes   []Combinations `gorm:"foreignKey:SecondMovieID"`
	WinnerVotes   []Votes        `gorm:"foreignKey:WinnerID"`
	VotesCount    int            `gorm:"not null"`
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
		VotesCount: m.VotesCount,
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
	User          Users        `gorm:"foreignKey:UserID"`
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

type Users struct {
	ID            string     `gorm:"primaryKey;not null"`
	Name          string     `gorm:"not null"`
	Email         string     `gorm:"unique;not null"`
	Password      string     `gorm:"not null"`
	IsAdmin       bool       `gorm:"not null"`
	Active        bool       `gorm:"not null"`
	CreatedAt     time.Time  `gorm:"not null"`
	UpdatedAt     *time.Time `gorm:"default:NULL"`
	DeactivatedAt *time.Time `gorm:"default:NULL"`
}

func (u *Users) ToEntity() *entities.User {
	login := entities.Login{
		Email:    u.Email,
		Password: u.Password,
	}

	return &entities.User{
		SharedEntity: entities.SharedEntity{
			ID:            u.ID,
			Active:        u.Active,
			CreatedAt:     u.CreatedAt,
			UpdatedAt:     u.UpdatedAt,
			DeactivatedAt: u.DeactivatedAt,
		},
		Name:  u.Name,
		Login: login,
	}
}

func Migration(db *gorm.DB, sqlDB *sql.DB) {
	if err := db.AutoMigrate(
		Lists{},
		Movies{},
		ListMovies{},
		Votes{},
		Combinations{},
		Users{},
	); err != nil {
		fmt.Println("Error during migration:", err)
		return
	}
	fmt.Println("Successful migration")
}
