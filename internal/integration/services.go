package integration

import "github.com/AdiKhoironHasan/matkul/pkg/dto"

type IntegServices interface {
	GetDosenID(req *dto.MatkulReqDTO) error
}
