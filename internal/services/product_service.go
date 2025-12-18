package services

import (
	"errors"
	"go-rest-api/internal/dto"
	"go-rest-api/internal/models"
	"go-rest-api/internal/repositories"

	"gorm.io/gorm"
)

type ProductService interface {
	Create(req dto.CreateProductRequest, userID uint) (*models.Product, error)
	GetAll(query dto.ProductQuery) ([]models.Product, int64, error)
	GetByID(id uint) (*models.Product, error)
	Update(id uint, req dto.UpdateProductRequest, userID uint) (*models.Product, error)
	Delete(id uint, userID uint) error
}

type productService struct {
	productRepo repositories.ProductRepository
}

func NewProductService(productRepo repositories.ProductRepository) ProductService {
	return &productService{productRepo: productRepo}
}

func (s *productService) Create(req dto.CreateProductRequest, userID uint) (*models.Product, error) {
	product := &models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		CreatedBy:   userID,
	}

	if err := s.productRepo.Create(product); err != nil {
		return nil, err
	}

	return s.productRepo.FindByID(product.ID)
}

func (s *productService) GetAll(query dto.ProductQuery) ([]models.Product, int64, error) {
	return s.productRepo.FindAll(query)
}

func (s *productService) GetByID(id uint) (*models.Product, error) {
	product, err := s.productRepo.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return product, nil
}

func (s *productService) Update(id uint, req dto.UpdateProductRequest, userID uint) (*models.Product, error) {
	product, err := s.GetByID(id)
	if err != nil {
		return nil, err
	}

	// if product.CreatedBy != userID {
	// 	return nil, errors.New("unauthorized to update this product")
	// }

	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.Stock = req.Stock

	if req.Name != "" {
		product.Name = req.Name
	}
	if req.Description != "" {
		product.Description = req.Description
	}
	if req.Price > 0 {
		product.Price = req.Price
	}
	if req.Stock >= 0 {
		product.Stock = req.Stock
	}

	if err := s.productRepo.Update(product); err != nil {
		return nil, err
	}

	return s.productRepo.FindByID(id)
}

func (s *productService) Delete(id uint, userID uint) error {
	_, err := s.GetByID(id)
	if err != nil {
		return err
	}

	// if product.CreatedBy != userID {
	// 	return errors.New("unauthorized to delete this product")
	// }

	return s.productRepo.Delete(id)
}
