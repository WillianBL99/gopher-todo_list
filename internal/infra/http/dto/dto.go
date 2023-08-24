package dto

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"
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

func requiredFields(s interface{}) []string {
	reqFields := make([]string, 0)
	v := reflect.ValueOf(s)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).String() == "" {
			reqFields = append(reqFields, strings.ToLower(v.Type().Field(i).Name))
		}
	}
	return reqFields
}
