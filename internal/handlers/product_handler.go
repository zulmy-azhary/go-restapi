package handlers

import (
	"go-rest-api/internal/dto"
	"go-rest-api/internal/services"
	"go-rest-api/internal/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	productService services.ProductService
}

func NewProductHandler(productService services.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) Create(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	var req dto.CreateProductRequest
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse("Invalid request body"))
	}

	product, err := h.productService.Create(req, userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(utils.SuccessResponse("Product created successfully", product))
}

func (h *ProductHandler) GetAll(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("perPage", "10"))
	search := c.Query("search", "")

	query := dto.ProductQuery{
		Page:    page,
		PerPage: perPage,
		Search:  search,
	}

	products, total, err := h.productService.GetAll(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(utils.PaginatedResponse("Products retrieved successfully", products, page, perPage, int(total)))
}

func (h *ProductHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse("Invalid product ID"))
	}

	product, err := h.productService.GetByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.ErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(utils.SuccessResponse("Product retrieved successfully", product))
}

func (h *ProductHandler) Update(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse("Invalid product ID"))
	}

	var req dto.UpdateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse("Invalid request body"))
	}

	product, err := h.productService.Update(uint(id), req, userID)
	if err != nil {
		if err.Error() == "unauthorized to update this product" {
			return c.Status(fiber.StatusForbidden).JSON(utils.ErrorResponse(err.Error()))
		}
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(utils.SuccessResponse("Product updated successfully", product))
}

func (h *ProductHandler) Delete(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse("Invalid product ID"))
	}

	if err := h.productService.Delete(uint(id), userID); err != nil {
		if err.Error() == "unauthorized to delete this product" {
			return c.Status(fiber.StatusForbidden).JSON(utils.ErrorResponse(err.Error()))
		}
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(utils.SuccessResponse("Product deleted successfully"))
}
