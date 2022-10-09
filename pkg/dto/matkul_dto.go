package dto

import (
	"errors"

	"github.com/AdiKhoironHasan/matkul/pkg/common/crypto"
	"github.com/AdiKhoironHasan/matkul/pkg/common/validator"
	"github.com/AdiKhoironHasan/matkul/pkg/env"
	util "github.com/AdiKhoironHasan/matkul/pkg/utils"
)

type MatkulReqDTO struct {
	Authorization string `json:"Authorization" valid:"required" validname:"authorization"`
	Signature     string `json:"signature" valid:"required" validname:"signature"`
	DateTime      string `json:"datetime" valid:"required" validname:"datetime"`
	ID            int64  `json:"id_dosen" valid:"required" validname:"id"`
	Nama          string `json:"nama" valid:"required" validname:"nama"`
}

func (dto *MatkulReqDTO) Validate() error {
	v := validator.NewValidate(dto)
	v.SetCustomValidation(true, func() error {
		return dto.customValidation()
	})
	return v.Validate()
}

func (dto *MatkulReqDTO) customValidation() error {

	signature := crypto.EncodeSHA256HMAC(util.GetBTBPrivKeySignature(), dto.Authorization, dto.DateTime)
	if signature != dto.Signature {
		if env.IsProduction() {
			return errors.New("invalid signature")
		}
		return errors.New("invalid signature" + " --> " + signature)
	}

	return nil
}

type GetDosenIDResDTO struct {
	// ID int64 `json:"id_dosen"`
	// datas string `json:"data"`
	Data []GetDosenDataResDTO `json:"data"`
}

type GetDosenDataResDTO struct {
	IdDosen int64 `json:"id_dosen"`
}
