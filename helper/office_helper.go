package helper

import (
	"cleancode/models"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// parse CSV row
func ParseReservation(row []string) (models.Reservation, error) {
	if len(row) < 4 {
		return models.Reservation{}, fmt.Errorf("row has fewer than 4 columns")
	}
	capacity, _ := strconv.Atoi(strings.TrimSpace(row[0]))
	rate, _ := strconv.ParseFloat(strings.TrimSpace(row[1]), 64)
	startDate, _ := time.Parse("2006-01-02", strings.TrimSpace(row[2]))
	var endDate *time.Time
	if strings.TrimSpace(row[3]) != "" {
		parsed, err := time.Parse("2006-01-02", strings.TrimSpace(row[3]))
		if err != nil {
			return models.Reservation{}, fmt.Errorf("invalid end date")
		}
		endDate = &parsed
	}
	return models.Reservation{Capacity: capacity, MonthlyRate: rate, StartDate: startDate, EndDate: endDate}, nil
}

// Check if reservation overlaps with the target month
func Overlaps(res models.Reservation, monthStart, monthEnd time.Time) (int, bool) {
	start := res.StartDate
	if start.Before(monthStart) {
		start = monthStart
	}
	end := monthEnd
	if res.EndDate != nil && res.EndDate.Before(monthEnd) {
		end = *res.EndDate
	}
	if start.After(end) {
		return 0, false
	}
	return int(end.Sub(start).Hours()/24) + 1, true
}
func GetDaysInMonth(year int, month time.Month) int {
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

// CSV Reader
func ReadReservations() ([]models.Reservation, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("cannot open input file")
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV")
	}
	var reservations []models.Reservation
	for i, row := range rows[1:] {
		res, err := ParseReservation(row)
		if err != nil {
			log.Printf("Skipping row %d: %v | Row data: %s", i+1, err, strings.Join(row, ", "))
			continue
		}
		reservations = append(reservations, res)
	}
	return reservations, nil
}
