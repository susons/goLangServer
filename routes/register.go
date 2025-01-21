package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/project_api/models"
	"github.com/gin-gonic/gin"
)

func registerEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf(`Couldnt register the event with id: %v`, id)})
		return
	}

	event, err := models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt not fetch event"})
		return
	}

	err = event.Register(event.UserId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt not register event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered event"})

}

func cancelRegister(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	userId := context.GetInt64("userId")

	var event models.Event
	event.Id = id

	event.CancelRegister(userId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf(`Couldnt cancel registration for the event with id: %v`, id)})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "Cancelled"})

}
