package app

import (
	"vacation_requests/controllers/employee_controller"
	"vacation_requests/controllers/manager_controller"
)

func mapUrls() {
	//employee_controller
	r.GET("/employee/GetRemianingDays", employee_controller.GetRemainingDays)
	r.GET("/employee/GetVactionRequest", employee_controller.GetVactionRequest)
	r.POST("/employee/CreateVactionRequest", employee_controller.CreateVactionRequest)

	//manager_controller
	r.GET("/GetRequests/:filter_type", manager_controller.GetRequests)
	r.GET("/GetRequstByEmployeeId/:emplyeeId", manager_controller.GetRequstByEmployeeId)
	r.GET("/GetOverlappedRequstId/:request_id", manager_controller.GetOverlappedIds)
	// r.POST("/ProcessRequestById:/:request_id", manager_controller.ProcessRequestById)
}
