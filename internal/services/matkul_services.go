package services

import (
	integ "github.com/AdiKhoironHasan/matkul/internal/integration"
	"github.com/AdiKhoironHasan/matkul/internal/repository"
	"github.com/AdiKhoironHasan/matkul/pkg/dto"
	"github.com/AdiKhoironHasan/matkul/pkg/dto/assembler"
)

type service struct {
	repo      repository.Repository
	IntegServ integ.IntegServices
}

func NewService(repo repository.Repository, IntegServ integ.IntegServices) Services {
	return &service{repo, IntegServ}
}

func (s *service) SaveMatkul(req *dto.MatkulReqDTO) error {
	// var resp := *dto.GetDosenIDResDTO

	err := s.IntegServ.GetDosenID(req)
	if err != nil {
		return err
	}

	dtMatkul := assembler.ToSaveMatkul(req)

	err = s.repo.SaveMatkul(dtMatkul)
	if err != nil {
		return err
	}

	return nil
}
