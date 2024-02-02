package controllers

import (
	"context"
	"database/sql"
	"errors"
	"ms-customer/domain/dto"
	"ms-customer/domain/entity"
	"ms-customer/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Customer interface {
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	FindAll(c echo.Context) error
	FindById(c echo.Context) error
}

type CustomerImpl struct {
	CustomerRepo repositories.Customer
}

func NewCustomerController(cr repositories.Customer) Customer {
	return &CustomerImpl{
		CustomerRepo: cr,
	}
}

func (p *CustomerImpl) Create(c echo.Context) error {
	req := dto.CustomerReq{}
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

	data := entity.Customer{
		ID:        req.ID,
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: dateNow,
		UpdatedAt: dateNow,
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*5)
	defer cancel()

	err := p.CustomerRepo.Create(ctx, data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"msg": "internal server error",
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"msg": "created",
	})
}

func (p *CustomerImpl) Update(c echo.Context) error {
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

	d, err := p.CustomerRepo.FindById(ctx, req.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"msg": "Not Found",
		})
	}

	dateNow := time.Now().String()
	data := entity.Customer{
		ID:        d.ID,
		Name:      d.Name,
		Email:     d.Email,
		CreatedAt: dateNow,
		UpdatedAt: dateNow,
	}

	if req.Name != "" {
		data.Name = req.Name
	}
	if req.Email != "" {
		data.Email = req.Email
	}

	if err := p.CustomerRepo.Update(ctx, data); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"msg": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"msg": "ok",
	})
}

func (p *CustomerImpl) Delete(c echo.Context) error {
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

	if err := p.CustomerRepo.Delete(ctx, pid); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"msg": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"msg": "ok",
	})
}

func (p *CustomerImpl) FindAll(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*5)
	defer cancel()

	data, err := p.CustomerRepo.FindAll(ctx)
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

func (p *CustomerImpl) FindById(c echo.Context) error {
	customerId := c.Param("id")
	pid, err := strconv.Atoi(customerId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"msg": "Bad Request",
		})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*5)
	defer cancel()

	data, err := p.CustomerRepo.FindById(ctx, pid)
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
