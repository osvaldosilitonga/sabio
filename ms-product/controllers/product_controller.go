package controllers

import (
	"context"
	"database/sql"
	"errors"
	"ms-product/domain/dto"
	"ms-product/domain/entity"
	"ms-product/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Product interface {
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	FindAll(c echo.Context) error
	FindById(c echo.Context) error
}

type ProductImpl struct {
	ProducRepo repositories.Product
}

func NewProductController(pr repositories.Product) Product {
	return &ProductImpl{
		ProducRepo: pr,
	}
}

func (p *ProductImpl) Create(c echo.Context) error {
	req := dto.ProductReq{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"msg":    "Bad Request",
			"detail": err,
		})
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"msg":    "Bad Request",
			"detail": err,
		})
	}

	dateNow := time.Now().String()

	data := entity.Product{
		ID:        req.ID,
		Name:      req.Name,
		Price:     req.Price,
		Stock:     req.Stock,
		CreatedAt: dateNow,
		UpdatedAt: dateNow,
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*5)
	defer cancel()

	err := p.ProducRepo.Create(ctx, data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"msg": "internal server error",
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"msg": "created",
	})
}

func (p *ProductImpl) Update(c echo.Context) error {
	req := dto.UpdateReq{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"msg":    "Bad Request",
			"detail": err,
		})
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"msg":    "Bad Request",
			"detail": err,
		})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*5)
	defer cancel()

	d, err := p.ProducRepo.FindById(ctx, req.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"msg": "Not Found",
		})
	}

	dateNow := time.Now().String()
	data := entity.Product{
		ID:        d.ID,
		Name:      d.Name,
		Price:     d.Price,
		Stock:     d.Stock,
		CreatedAt: dateNow,
		UpdatedAt: dateNow,
	}

	if req.Name != "" {
		data.Name = req.Name
	}
	if req.Price > 0 {
		data.Price = req.Price
	}
	if req.Stock > 0 {
		data.Stock = req.Stock
	}

	if err := p.ProducRepo.Update(ctx, data); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"msg": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"msg": "ok",
	})
}

func (p *ProductImpl) Delete(c echo.Context) error {
	productId := c.Param("id")
	pid, err := strconv.Atoi(productId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"msg":    "Bad Request",
			"detail": err,
		})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*5)
	defer cancel()

	if err := p.ProducRepo.Delete(ctx, pid); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"msg": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"msg": "ok",
	})
}

func (p *ProductImpl) FindAll(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*5)
	defer cancel()

	data, err := p.ProducRepo.FindAll(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"msg": "Internal Server Error",
		})
	}

	if len(data) < 1 {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, echo.Map{
				"msg": "Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"msg":  "ok",
		"data": data,
	})
}

func (p *ProductImpl) FindById(c echo.Context) error {
	productId := c.Param("id")
	pid, err := strconv.Atoi(productId)

	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*5)
	defer cancel()

	data, err := p.ProducRepo.FindById(ctx, pid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"msg": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"msg":  "ok",
		"data": data,
	})
}
