package repositories

import (
	"errors"
	"go-rest-api/internal/dto"
	"go-rest-api/internal/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *models.Product) error
	FindAll(query dto.ProductQuery) ([]models.Product, int64, error)
	FindByID(id uint) (*models.Product, error)
	Update(product *models.Product) error
	Delete(id uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) FindAll(query dto.ProductQuery) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64

	db := r.db.Model(&models.Product{}).Preload("User")

	if query.Search != "" {
		db = db.Where("name ILIKE ?", "%"+query.Search+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, errors.New("product not found")
		}
		return nil, 0, err
	}

	if query.Page < 1 {
		query.Page = 1
	}
	if query.PerPage < 1 {
		query.PerPage = 10
	}

	offset := (query.Page - 1) * query.PerPage
	if err := db.Offset(offset).Limit(query.PerPage).Find(&products).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, errors.New("product not found")
		}
		return nil, 0, err
	}

	return products, total, nil
}

func (r *productRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := r.db.Preload("User").First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}
