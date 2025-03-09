package service

import (
	"errors"
	"teste/internal/models"
)

func ValidateBody(data *models.Produto) error {

	if data.Valor < 0 {
		return errors.New("O valor deve ser maior que 0")
	}

	return nil

}
