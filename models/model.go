package models

import "time"

type ErrorResponse struct {
	Error string `json:"error"`
}
type Reservation struct {
	Capacity    int
	MonthlyRate float64
	StartDate   time.Time
	EndDate     *time.Time
}
type Request struct {
	Month string `json:"month" example:"2014-05"`
}

	// type Response struct {
	// 	Month              string  `json:"month"`
	// 	Revenue            float64 `json:"revenue"`
	// 	UnreservedCapacity int     `json:"unreserved_capacity"`
	// }
