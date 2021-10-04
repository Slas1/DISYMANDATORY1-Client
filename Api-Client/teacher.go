package client

type teacher struct {
	ID             int32      `json:"teacherId"`
	Name           string     `json:"name"`
	NRatings       int32      `json:"nRatings"`
	AvgRatings     float32    `json:"avgRatings"`
	ActiveStudents []*student `json:"activeStudents"`
}
