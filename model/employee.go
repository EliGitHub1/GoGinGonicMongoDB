package model

type Employee struct {
	Id             int64 `bson:"id" json:"id" binding:"required"`
	Remaining_days int64 `json:"remaining_days" binding:"required"`
}
