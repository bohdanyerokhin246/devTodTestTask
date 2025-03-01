package handlers

import (
	"devTodTestTask/internal/models"
	"devTodTestTask/internal/services"
	"devTodTestTask/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CatHandler struct {
	Service *services.CatService
}

// CreateCatHandler godoc
// @Summary Create a new cat
// @Description Create a new cat and store it in the database
// @Param cat body models.Cat true "Cat data"
// @Success 201 {object} models.Cat "Successfully created cat"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /cats [post]
func (h *CatHandler) CreateCatHandler(c *gin.Context) {
	var cat models.Cat
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid input"})
		return
	}

	if err := utils.ValidateBreed(cat.Breed); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid breed"})
		return
	}

	if err := h.Service.CreateCat(&cat); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, cat)
}

// ListCatsHandler godoc
// @Summary Get list of all cats
// @Description Get a list of all cats in the database
// @Success 200 {array} models.Cat "List of cats"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /cats [get]
func (h *CatHandler) ListCatsHandler(c *gin.Context) {
	cats, err := h.Service.ListCats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, cats)
}

// CatByIDHandler godoc
// @Summary Get cat by ID
// @Description Get a specific cat by its ID
// @Param id path int true "Cat ID"
// @Success 200 {object} models.Cat "Cat data"
// @Failure 400 {object} ErrorResponse "Invalid ID format"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /cats/{id} [get]
func (h *CatHandler) CatByIDHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	cats, err := h.Service.CatByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, cats)
}

// UpdateCatHandler godoc
// @Summary Update a cat
// @Description Update an existing cat's details
// @Param cat body models.Cat true "Updated cat data"
// @Success 200 {object} ErrorResponse "Successfully updated cat"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /cats [put]
func (h *CatHandler) UpdateCatHandler(c *gin.Context) {
	var cat models.Cat
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid input"})
		return
	}

	err := h.Service.UpdateCat(&cat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, ErrorResponse{Error: "Successfully updated cat"})
}

// DeleteCatHandler godoc
// @Summary Delete a cat
// @Description Delete a cat from the database by its ID
// @Param id path int true "Cat ID"
// @Success 200 {object} ErrorResponse "Successfully deleted cat"
// @Failure 400 {object} ErrorResponse "Invalid ID format"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /cats/{id} [delete]
func (h *CatHandler) DeleteCatHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.Service.DeleteCat(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, ErrorResponse{Error: "Successfully deleted cat"})
}

type ErrorResponse struct {
	Error string `json:"error"`
}
