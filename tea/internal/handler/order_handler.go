package handler

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"

	service "tea/internal/service"
	models "tea/models"
)

func PostOrders(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	file, _ := io.ReadAll(r.Body)

	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)
	logger.Error("Error reding file in: ")

	json.Unmarshal(file, &order)
	service.PostOrder(order)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("../data/files.json")
	defer file.Close()
	data, _ := io.ReadAll(file)
	service.GetOrder(data, w)
}

func GetOrdersID(w http.ResponseWriter, r *http.Request) {
}

func PutOrdersID(w http.ResponseWriter, r *http.Request) {
}

func DeleteOrdersID(w http.ResponseWriter, r *http.Request) {
}

func PostOrdersIDClose(w http.ResponseWriter, r *http.Request) {
}
