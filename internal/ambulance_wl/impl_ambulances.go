package ambulance_wl

import (
	"net/http"

	"github.com/Marek-FIIT/ambulance-webapi/internal/db_service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Kópia zakomentovanej časti z api_ambulances.go
// CreateAmbulance - Saves new ambulance definition
func (this *implAmbulancesAPI) CreateAmbulance(ctx *gin.Context) {
	value, exists := ctx.Get("db_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db not found",
				"error":   "db not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[Ambulance])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db context is not of required type",
				"error":   "cannot cast db context to db_service.DbService",
			})
		return
	}

	ambulance := Ambulance{}
	err := ctx.BindJSON(&ambulance)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "Invalid request body",
				"error":   err.Error(),
			})
		return
	}

	if ambulance.Id == "" {
		ambulance.Id = uuid.New().String()
	}

	err = db.CreateDocument(ctx, ambulance.Id, &ambulance)

	switch err {
	case nil:
		ctx.JSON(
			http.StatusCreated,
			ambulance,
		)
	case db_service.ErrConflict:
		ctx.JSON(
			http.StatusConflict,
			gin.H{
				"status":  "Conflict",
				"message": "Ambulance already exists",
				"error":   err.Error(),
			},
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to create ambulance in database",
				"error":   err.Error(),
			},
		)
	}
}

// DeleteAmbulance - Deletes specific ambulance
func (this *implAmbulancesAPI) DeleteAmbulance(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNotImplemented)
}
