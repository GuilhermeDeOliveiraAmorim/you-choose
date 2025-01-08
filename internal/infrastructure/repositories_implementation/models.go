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
	TypeList      string     `gorm:"not null"`
	Movies        []Movies   `gorm:"many2many:list_movies;"`
	Brands        []Brands   `gorm:"many2many:list_brands;"`
}

func (m *Lists) ToEntity(items []interface{}, combinations []entities.Combination, complete bool) *entities.List {
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
			TypeList:     m.TypeList,
			Items:        items,
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
		Name:     m.Name,
		Cover:    m.Cover,
		TypeList: m.TypeList,
	}
}

type Movies struct {
	ID            string     `gorm:"primaryKey;not null"`
	Active        bool       `gorm:"not null"`
	CreatedAt     time.Time  `gorm:"not null"`
	UpdatedAt     *time.Time `gorm:"default:NULL"`
	DeactivatedAt *time.Time `gorm:"default:NULL"`
	Name          string     `gorm:"not null"`
	Year          int64      `gorm:"not null"`
	Poster        string     `gorm:"not null"`
	ExternalID    string     `gorm:"not null"`
	VotesCount    int        `gorm:"not null"`
	Lists         []Lists    `gorm:"many2many:list_movies;"`
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
	ID           string `gorm:"primaryKey;not null"`
	ListID       string `gorm:"not null"`
	List         Lists  `gorm:"foreignKey:ListID"`
	FirstItemID  string `gorm:"not null"`
	SecondItemID string `gorm:"not null"`
}

func (c *Combinations) ToEntity() *entities.Combination {
	return &entities.Combination{
		ID:           c.ID,
		ListID:       c.ListID,
		FirstItemID:  c.FirstItemID,
		SecondItemID: c.SecondItemID,
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

type Brands struct {
	ID            string     `gorm:"primaryKey;not null"`
	Name          string     `gorm:"not null"`
	Logo          string     `gorm:"not null"`
	VotesCount    int        `gorm:"not null"`
	Active        bool       `gorm:"not null"`
	CreatedAt     time.Time  `gorm:"not null"`
	UpdatedAt     *time.Time `gorm:"default:NULL"`
	DeactivatedAt *time.Time `gorm:"default:NULL"`
	Lists         []Lists    `gorm:"many2many:list_brands;"`
}

type ListBrands struct {
	ListID        string     `gorm:"primaryKey"`
	List          Lists      `gorm:"foreignKey:ListID"`
	BrandID       string     `gorm:"primaryKey"`
	Brand         Brands     `gorm:"foreignKey:BrandID"`
	CreatedAt     time.Time  `gorm:"not null"`
	DeactivatedAt *time.Time `gorm:"default:NULL"`
}

type ListMovies struct {
	ListID        string     `gorm:"primaryKey"`
	List          Lists      `gorm:"foreignKey:ListID"`
	MovieID       string     `gorm:"primaryKey"`
	Movie         Movies     `gorm:"foreignKey:MovieID"`
	CreatedAt     time.Time  `gorm:"not null"`
	DeactivatedAt *time.Time `gorm:"default:NULL"`
}

func (b *Brands) ToEntity() *entities.Brand {
	return &entities.Brand{
		SharedEntity: entities.SharedEntity{
			ID:            b.ID,
			Active:        b.Active,
			CreatedAt:     b.CreatedAt,
			UpdatedAt:     b.UpdatedAt,
			DeactivatedAt: b.DeactivatedAt,
		},
		Name:       b.Name,
		Logo:       b.Logo,
		VotesCount: b.VotesCount,
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
		Brands{},
		ListBrands{},
	); err != nil {
		fmt.Println("Error during migration:", err)
		return
	}
	fmt.Println("Successful migration")
}
