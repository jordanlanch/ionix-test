package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jordanlanch/ionix-test/domain"
)

type VaccinationController struct {
	VaccinationUsecase domain.VaccinationUsecase
}

func NewVaccinationController(vu domain.VaccinationUsecase) *VaccinationController {
	return &VaccinationController{
		VaccinationUsecase: vu,
	}
}

func (vc *VaccinationController) Create(c *gin.Context) {
	var vac domain.Vaccination
	if err := c.ShouldBindJSON(&vac); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := vc.VaccinationUsecase.CreateVaccination(c.Request.Context(), &vac)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, vac)
}

func (vc *VaccinationController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var vac domain.Vaccination
	if err := c.ShouldBindJSON(&vac); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = vc.VaccinationUsecase.UpdateVaccination(c.Request.Context(), id, &vac)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vaccination updated successfully"})
}

func (vc *VaccinationController) FetchAll(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	offsetStr := c.DefaultQuery("offset", "0")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset parameter"})
		return
	}

	pagination := domain.Pagination{Limit: &limit, Offset: &offset}

	vaccinations, err := vc.VaccinationUsecase.GetAllVaccinations(c.Request.Context(), &pagination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, vaccinations)
}

func (vc *VaccinationController) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	vac, err := vc.VaccinationUsecase.GetVaccinationById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, vac)
}

func (vc *VaccinationController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = vc.VaccinationUsecase.DeleteVaccination(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vaccination deleted successfully"})
}
