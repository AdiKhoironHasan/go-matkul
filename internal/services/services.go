package services

import (
	"github.com/AdiKhoironHasan/matkul/pkg/dto"
)

type Services interface {
	SaveMatkul(req *dto.MatkulReqDTO) error
}
