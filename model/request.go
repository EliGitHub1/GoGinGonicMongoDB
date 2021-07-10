package model

import (
	"github.com/aws/aws-sdk-go/service/dataexchange"	"golang.org/x/text/date"
)


type Request struct {
	Id                  int64  `bson:"id" json:"id" binding:"required"`
	Author              int64  `json:"author" binding:"required"`
	Status              string `json:"status"`
	Resolved_by         int64  `json:"resolved_by"`
	Request_created_at  string `json:"request_created_at"`
	Vacation_start_date string `json:"vacation_start_date"`
	Vacation_end_date   string     `json:"vacation_end_date"`
}
