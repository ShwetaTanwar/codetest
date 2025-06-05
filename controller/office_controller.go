package controller

import (
	"cleancode/helper"
	"cleancode/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Month string `json:"month" example:"2014-05"`
}

// CalculateHandler  		Calculate monthly revenue and unreserved capacity
// CalculateHandler			godoc
// @Tags					OfficeReservation API
// @Summary					Calculate revenue and capacity for a month
// @Description				Accepts JSON with month (YYYY-MM) and returns total revenue and unreserved capacity for that month based on CSV reservation data.
// @Accept					json
// @Produce					json
// @Param					request body models.Request true "Request body should have a 'month' field in YYYY-MM format. Example: { \"month\": \"2014-05\" }"
// @Success					200 {object} string
// @Failure					400 {object} models.ErrorResponse
// @Router					/calculate [post]
func CalculateHandler(c *gin.Context) {
	var req models.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	yearMonth, err := time.Parse("2006-01", req.Month)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	monthStart := time.Date(yearMonth.Year(), yearMonth.Month(), 1, 0, 0, 0, 0, time.UTC)
	monthEnd := time.Date(yearMonth.Year(), yearMonth.Month(), helper.GetDaysInMonth(yearMonth.Year(), yearMonth.Month()), 0, 0, 0, 0, time.UTC)
	daysInMonth := helper.GetDaysInMonth(yearMonth.Year(), yearMonth.Month())
	reservations, err := helper.ReadReservations()
	if err != nil {
		log.Fatal("Failed to read reservations: ", err)
	}
	totalRevenue := 0.0
	unreservedCapacity := 0
	for _, res := range reservations {
		overlapDays, isReserved := helper.Overlaps(res, monthStart, monthEnd) //revenue or unreserved
		if isReserved {
			proRated := (res.MonthlyRate / float64(daysInMonth)) * float64(overlapDays)
			totalRevenue += proRated
		} else {
			unreservedCapacity += res.Capacity
		}
	}
	c.Header("Content-Type", "text/plain")
	c.String(http.StatusOK,
		"%s: expected revenue: $%.2f, expected total capacity of the unreserved offices: %d",
		req.Month,
		totalRevenue,
		unreservedCapacity,
	)
	// c.JSON(http.StatusOK, Response{
	// 	Month:              req.Month,
	// 	Revenue:            totalRevenue,
	// 	UnreservedCapacity: unreservedCapacity,
	// })
}

// ManualHandler displays the CSV data of reservations text file.
// @Summary      Show reservations table
// @Description  Displays the CSV data.
// @Tags         OfficeReservation API
// @Produce      text/html
// @Param 		Authorization header string true "Insert your access token" default(Bearer )
// @Success      200 {string} string "HTML page with CSV data"
// @Failure      500 {object} models.ErrorResponse
// @Router       /manual [get]
func ManualHandler(c *gin.Context) {
	data, err := os.ReadFile("input.txt") // Make sure path is correct
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to read input.txt: %v", err)
		return
	}
	c.Data(http.StatusOK, "text/plain", data)
}
