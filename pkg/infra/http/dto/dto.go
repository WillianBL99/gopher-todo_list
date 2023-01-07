package dto

import (
	"encoding/json"
	"net/http"
)

type DTO interface {
	Validate() error
}

func ToDTO(r *http.Request, dto DTO) error {
	var err error
	err = json.NewDecoder(r.Body).Decode(dto)
	err = dto.Validate()

	return err
}