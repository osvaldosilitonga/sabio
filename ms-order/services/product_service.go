package service

import (
	"context"
	"encoding/json"
	"fmt"
	"ms-order/domain/dto"
	"net/http"
	"os"
)

type Product interface {
	FindById(ctx context.Context, id int) (dto.Product, int, error)
}

type ProductImpl struct{}

func NewProductService() Product {
	return &ProductImpl{}
}

func (c *ProductImpl) FindById(ctx context.Context, id int) (dto.Product, int, error) {
	baseUrl := os.Getenv("PRODUCT_SERVICE_BASE_URL")

	data := dto.Product{}

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

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return data, http.StatusInternalServerError, err
	}

	return data, resp.StatusCode, nil

	// switch resp.StatusCode {
	// case http.StatusOK:
	// 	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
	// 		return data, http.StatusInternalServerError, err
	// 	}

	// 	log.Println(data, "<--------- data")

	// 	return data, http.StatusOK, nil

	// case http.StatusNotFound:
	// 	return data, http.StatusNotFound, errors.New("Not Found")

	// default:
	// 	return data, http.StatusInternalServerError, errors.New("Internal Server Error")
	// }
}
