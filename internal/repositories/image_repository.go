package repositories

type ImageRepository interface {
	SaveImage(poster string) (string, error)
}
