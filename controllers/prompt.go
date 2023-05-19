package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/martian/v3/log"
	"net/http"
	"personal-finance-api/models"
	"personal-finance-api/services"
)

// GeneratePlan godoc
//
//	@Summary		GeneratePlan
//	@Description	Generate a plan based on inputs
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.PlanInput	true	"query params"
//	@Success		201		{object}	models.ApiResponse
//	@Failure		400		{object}	models.ApiResponse
//	@Failure		500		{object}	models.ApiResponse
//	@Router			/generate [post]
func GeneratePlan(c *gin.Context) {
	var input models.PlanInput
	if err := c.ShouldBindJSON(&input); err != nil {
		dr := models.CreateResponse(err, "Missing info to create a plan", nil)
		c.JSON(http.StatusBadRequest, dr)
		return
	}

	dc := models.DataCollected{Email: input.Email}

	exists, err := models.CheckIfEmailAlreadyExists("users", input.Email)
	if err != nil {
		log.Errorf(err.Error())
		dr := models.CreateResponse(err, "An error occurred while retrieving data to firestore", nil)
		c.JSON(http.StatusInternalServerError, dr)
		return
	}
	if !exists {
		_, err = models.WriteToFirestore("users", dc)

		if err != nil {
			log.Errorf(err.Error())
			dr := models.CreateResponse(err, "An error occurred while writing data to firestore", nil)
			c.JSON(http.StatusInternalServerError, dr)
			return
		}
	}

	prompt := models.PromptFromPlan(input)
	content, err := services.SendPrompt(prompt)

	if err != nil || content == "" {
		log.Errorf(err.Error())
		dr := models.CreateResponse(err, "An error occurred while sending data to chatgpt", nil)
		c.JSON(http.StatusInternalServerError, dr)
		return
	}
	services.SendEmail(content, input.Email)
	dr := models.CreateResponse(err, "Success!", content)
	c.JSON(http.StatusOK, dr)
}
