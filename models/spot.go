package models

type Spot struct {
	ID          int     `json:"id,omitempty" bson:"_id,omitempty"`
	Address     string  `json:"address"`
	City        string  `json:"city"`
	XCoordinate float64 `json:"x_coordinate"`
	YCoordinate float64 `json:"y_coordinate"`
	Status      string  `json:"status"`
	LessorID    int     `json:"lessor_id"`
}
