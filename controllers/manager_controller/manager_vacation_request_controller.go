package manager_controller

import "github.com/gin-gonic/gin"

// - See an overview of all requests
// - Filter by pending and approved
func GetRequests(c *gin.Context) {}

// - See an overview for each individual employee
func GetRequstByEmployeeId(c *gin.Context) {}

// - See an overview of overlapping requests
func GetOverlappedIds(c *gin.Context) {}

// Process an individual request and either approve or reject it
func ProcessRequestById(c *gin.Context) {}
