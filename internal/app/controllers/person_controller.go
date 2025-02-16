package controllers

import (
	"health-history/internal/app/models"
	"health-history/internal/app/repositories"
	"health-history/internal/app/requests"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PersonController struct {
	PersonReporitory *repositories.PersonRepository
}

func NewPersonController() *PersonController {
	return &PersonController{PersonReporitory: &repositories.PersonRepository{}}
}

// CreatePerson create new person
// @Summary Create person
// @Description Create a new person
// @Tags persons
// @Accept json
// @Produce json
// @Param request body requests.PersonRequest true "Data of person"
// @Success 201 {object} models.Person "Successful response"
// @Failure 400 {object} map[string]string "Validation error"
// @Router /persons [post]
func (pc *PersonController) CreatePerson(c *gin.Context) {
	var person requests.PersonRequest

	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	personModel := models.Person{
		Name: person.Name,
		Age:  person.Age,
	}

	if err := pc.PersonReporitory.CreatePerson(&personModel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, person)
}

// UpdatePerson update a person
// @Summary Update  person
// @Description Update a exist person
// @Tags persons
// @Accept json
// @Produce json
// @Param request body requests.PersonRequest true "Data of person"
// @Param id path int true "Mdication ID" example(1)
// @Success 200 {object} models.Person "Successful response"
// @Failure 400 {object} map[string]string "Validation error"
// @Router /persons/:id [put]
func (pc *PersonController) UpdatePerson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var person requests.PersonRequest
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	personModel := models.Person{
		ID:   id,
		Name: person.Name,
		Age:  person.Age,
	}

	if err := pc.PersonReporitory.UpdatePerson(&personModel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, person)
}

// DeletePerson delete a person
// @Summary Delete person
// @Description Delete a person by id
// @Tags persons
// @Produce json
// @Param id path int true "Mdication ID" example(1)
// @Success 204
// @Router /persons/:id [get]
func (pc *PersonController) DeletePerson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := pc.PersonReporitory.DeletePerson(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetPersonById return a person
// @Summary Return person
// @Description Return person by id
// @Tags persons
// @Produce json
// @Param id path int true "Person ID" example(1)
// @Success 200 {object} models.Person "Successful response"
// @Failure 404 {object} map[string]string "Person not found"
// @Router /persons/:id [get]
func (pc *PersonController) GetPersonById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	person, err := pc.PersonReporitory.GetPersonById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}

	c.JSON(http.StatusOK, person)
}

// GetAllPersons return person list
// @Summary List persons
// @Description Return person list
// @Tags persons
// @Produce json
// @Success 200 {array} models.Person "Successful response"
// @Router /persons [get]
func (pc *PersonController) GetAllPersons(c *gin.Context) {
	persons, err := pc.PersonReporitory.GetAllPersons()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, persons)
}
