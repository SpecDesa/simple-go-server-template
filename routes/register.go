package routes

import (
	"net/http"
	"strconv"

	"example.com/rest/models"
	"example.com/rest/utils"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {

	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse event id."})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Event not found."})
		return
	}

	err = event.Register(userId)
	if utils.CheckIfErrorContextJson("Event not found", err, http.StatusInternalServerError, context) {
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered!"})

}

func cancelRegistration(context *gin.Context) {

	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if utils.CheckIfErrorContextJson("Couldn't parse event id", err, http.StatusBadRequest, context) {
		return
	}

	event, err := models.GetEventById(eventId)
	if utils.CheckIfErrorContextJson("Event not found", err, http.StatusBadRequest, context) {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Event not found."})
		return
	}

	err = event.DeleteRegistration(userId)
	if utils.CheckIfErrorContextJson("Event not found or user not signed up", err, http.StatusInternalServerError, context) {
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registration removed."})
}

func getRegistrations(context *gin.Context) {
	events, err := models.GetAllRegistrations()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get all registrations"})
		return
	}

	context.JSON(http.StatusOK, events)
}
