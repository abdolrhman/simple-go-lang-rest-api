package controller

import (
	"github.com/abdolrhman/simple-go-lang-rest-api/model"
	"github.com/abdolrhman/simple-go-lang-rest-api/service"
	"github.com/gin-gonic/gin"
)

var err error

// GetCustomers godoc
// @Summary List customers
// @Description get customers
// @Tags customers
// @Accept  json
// @Produce  json
// @Success 200 "Success"
// @Router /customers/ [get]
func (base *Controller) GetCustomers(c *gin.Context) {
	var args model.Args

	// Define and get sorting field
	args.Sort = c.DefaultQuery("Sort", "ID")

	// Define and get sorting order field
	args.Order = c.DefaultQuery("Order", "DESC")

	// Define and get offset for pagination
	args.Offset = c.DefaultQuery("Offset", "0")

	// Define and get limit for pagination
	args.Limit = c.DefaultQuery("Limit", "25")

	// Get search keyword for Search Scope
	args.Search = c.DefaultQuery("Search", "")

	// Fetch results from database
	customers, filteredData, totalData, err := service.GetCustomers(base.DB, args)
	if err != nil {
		c.AbortWithStatus(404)
	}

	// Fill return data struct
	data := model.Data{
		TotalData:    totalData,
		FilteredData: filteredData,
		Data:         customers,
	}

	c.JSON(200, data)
}
