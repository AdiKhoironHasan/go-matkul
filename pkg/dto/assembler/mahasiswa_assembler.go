package assembler

import (
	"github.com/AdiKhoironHasan/matkul/internal/models"
	"github.com/AdiKhoironHasan/matkul/pkg/dto"
)

func ToSaveMatkul(d *dto.MatkulReqDTO) *models.MatkulModels {
	return &models.MatkulModels{
		ID:   d.ID,
		Name: d.Nama,
	}
}
