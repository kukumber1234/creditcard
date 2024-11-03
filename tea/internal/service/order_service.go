package service

import (
	"net/http"

	dal "tea/internal/dal"
	"tea/models"
)

func PostOrder(order models.Order) {
	//logic will be here
	
	dal.OrderPost(order) //after check go to bucket dal
}

func GetOrder(data []byte, w http.ResponseWriter) {
	// logic
	
	w.Header().Set("Content-type", "application/json")
	w.Write(data)
}
