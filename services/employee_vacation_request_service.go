package services

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"vacation_requests/model"
	"vacation_requests/utils/db"
	"vacation_requests/utils/rest_errors"

	"gopkg.in/mgo.v2/bson"
)

func CreateRequest(req model.Request) rest_errors.RestErr {

	client, err := db.GetMongoClient()
	if err != nil {
		return rest_errors.NewRestError("Error:function GetMongoClient failed", http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	if isRequstAlreadyExist(req) {
		return rest_errors.NewRestError("Error:request already exists", http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	days := findNumberOfDaysRequested(req)
	if days == 0 {
		return rest_errors.NewRestError("Error:bad dates input", http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	} else if days > 30 {
		//assumimg all request are for the same year
		return rest_errors.NewRestError("Error:Asked for more than 30", http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	collection := client.Database(db.DB_REQUESTS).Collection(db.COLLECTION_REQUESTS)

	_, errInsertOne := collection.InsertOne(context.TODO(), req)
	if errInsertOne != nil {
		return rest_errors.NewRestError("Error:function InsertOne failed", http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
	return nil
}

func GetRemainingDays(req model.Request) (int, rest_errors.RestErr) {
	client, rest_err := db.GetMongoClient()
	if rest_err != nil {
		return 0, rest_errors.NewRestError("Error:function GetMongoClient failed", http.StatusInternalServerError, http.StatusText(500))
	}
	collection := client.Database(db.DB_REQUESTS).Collection(db.COLLECTION_EMPLOYEES)
	var res bson.M
	cursor, err := collection.FindOne(context.TODO(), bson.M{"id": req.Author}).Decode(&res); err != nil {
		return 0, rest_errors.NewRestError("Error:failed to fetch data", http.StatusInternalServerError, http.StatusText(500))
	}

	return res, nil
}

func isRequstAlreadyExist(req model.Request) bool {
	client, _ := db.GetMongoClient()
	collection := client.Database(db.DB_REQUESTS).Collection(db.COLLECTION_REQUESTS)
	var res bson.M
	if err := collection.FindOne(context.TODO(), bson.M{"id": req.Id}).Decode(&res); err != nil {
		log.Fatal(err)
	}
	if res != nil {
		return true
	}
	return false
}

func findNumberOfDaysRequested(req model.Request) int {
	layout := "2006-01-02T15:04:05.000Z"
	start := req.Vacation_start_date
	end := req.Vacation_end_date

	t1, err := time.Parse(layout, start)
	if err != nil {
		return 0
	}

	t2, err := time.Parse(layout, end)
	if err != nil {
		return 0
	}

	if t2.After(t1) == false {
		return 0
	}
	return t2.Sub(t1).Hours() / 24
}

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}
