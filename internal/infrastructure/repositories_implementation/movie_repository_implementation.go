package repositories_implementation

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"

	"cloud.google.com/go/storage"
)

type MovieRepository struct {
	gorm       *gorm.DB
	BucketName string
}

func NewMovieRepository(gorm *gorm.DB, bucketName string) *MovieRepository {
	return &MovieRepository{
		gorm:       gorm,
		BucketName: bucketName,
	}
}

func (c *MovieRepository) CreateMovie(movie entities.Movie) error {
	tx := c.gorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := tx.Create(&Movies{
		ID:            movie.ID,
		Active:        movie.Active,
		CreatedAt:     movie.CreatedAt,
		UpdatedAt:     movie.UpdatedAt,
		DeactivatedAt: movie.DeactivatedAt,
		Name:          movie.Name,
		Year:          movie.Year,
		Poster:        movie.Poster,
		ExternalID:    movie.ExternalID,
		VotesCount:    movie.VotesCount,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (c *MovieRepository) GetMovieByID(movieID string) (entities.Movie, error) {
	var movieModel Movies

	result := c.gorm.Model(&Movies{}).Where("id =? AND active =?", movieID, true).First(&movieModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entities.Movie{}, errors.New("movie not found")
		}
		return entities.Movie{}, errors.New(result.Error.Error())
	}

	return *movieModel.ToEntity(), nil
}

func (c *MovieRepository) ThisMovieExist(movieExternalID string) (bool, error) {
	var movieModel Movies

	result := c.gorm.Model(&Movies{}).Where("external_id =?", movieExternalID).First(&movieModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, errors.New(result.Error.Error())
	}

	return true, nil
}

func (c *MovieRepository) GetMoviesByIDs(moviesIDs []string) ([]entities.Movie, error) {
	var moviesModel []Movies

	result := c.gorm.Model(&Movies{}).Where("id IN?", moviesIDs).Find(&moviesModel)
	if result.Error != nil {
		return nil, errors.New(result.Error.Error())
	}

	var movies []entities.Movie
	for _, movieModel := range moviesModel {
		movies = append(movies, *movieModel.ToEntity())
	}

	return movies, nil
}

func (c *MovieRepository) SavePoster(poster string) (string, error) {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to create storage client: %v", err)
	}
	defer client.Close()

	resp, err := http.Get(poster)
	if err != nil {
		return "", fmt.Errorf("failed to download poster: %v", err)
	}
	defer resp.Body.Close()

	imageData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read poster data: %v", err)
	}

	objectName := ulid.Make().String()

	bucket := client.Bucket(c.BucketName)

	writer := bucket.Object(objectName).NewWriter(ctx)

	writer.ContentType = http.DetectContentType(imageData)

	if _, err := writer.Write(imageData); err != nil {
		writer.Close()
		return "", fmt.Errorf("failed to upload poster to bucket: %v", err)
	}

	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("failed to finalize upload to bucket: %v", err)
	}

	return objectName, nil
}

func (c *MovieRepository) UpdadeMovie(movie entities.Movie) error {
	tx := c.gorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := tx.Model(&Movies{}).Where("id =?", movie.ID).Updates(Movies{
		Active:        movie.Active,
		Name:          movie.Name,
		Year:          movie.Year,
		Poster:        movie.Poster,
		VotesCount:    movie.VotesCount,
		DeactivatedAt: movie.DeactivatedAt,
		UpdatedAt:     movie.UpdatedAt,
		ExternalID:    movie.ExternalID,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
