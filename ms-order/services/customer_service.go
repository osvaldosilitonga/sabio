package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"ms-order/domain/dto"
	"net/http"
	"os"
)

type Customer interface {
	FindById(ctx context.Context, id int) (dto.Customer, int, error)
}

type CustomerImpl struct{}

func NewCustomerService() Customer {
	return &CustomerImpl{}
}

func (c *CustomerImpl) FindById(ctx context.Context, id int) (dto.Customer, int, error) {
	baseUrl := os.Getenv("CUSTOMER_SERVICE_BASE_URL")

	data := dto.Customer{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%v/%v", baseUrl, id), nil)
	if err != nil {
		return data, http.StatusInternalServerError, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return data, http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return data, http.StatusInternalServerError, err
		}

		return data, http.StatusOK, nil

	case http.StatusNotFound:
		return data, http.StatusNotFound, errors.New("Not Found")

	default:
		return data, http.StatusInternalServerError, errors.New("Internal Server Error")
	}
}
