package repository

import "github.com/AdiKhoironHasan/matkul/internal/models"

type Repository interface {
	SaveMatkul(dataMahasiswa *models.MatkulModels) error
}
