package dto

import (
	"github.com/AdiKhoironHasan/matkul/pkg/common/validator"
)

type MatkulReqDTO struct {
	ID   int64  `json:"id_dosen" valid:"required" validname:"id"`
	Nama string `json:"nama" valid:"required" validname:"nama"`
}

func (dto *MatkulReqDTO) Validate() error {
	v := validator.NewValidate(dto)

	return v.Validate()
}

type GetDosenIDResDTO struct {
	// ID int64 `json:"id_dosen"`
	// datas string `json:"data"`
	Data []GetDosenDataResDTO `json:"data"`
}

type GetDosenDataResDTO struct {
	IdDosen int64 `json:"id_dosen"`
	// NoRumah string `json:"no_rumah"`
}
