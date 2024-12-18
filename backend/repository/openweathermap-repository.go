package repository

type OpenweathermapRepository interface{}

type openweathermapRepository struct{}

func NewOpenweathermapRepository() OpenweathermapRepository {
	return &openweathermapRepository{}
}
