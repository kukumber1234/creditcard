package main

import (
	"log"
	"net/http"

	handler "tea/internal/handler"
)

type Person struct {
	Name      string   `json:"Name"`
	Age       int      `json:"Age"`
	IsStudent bool     `json:"IsStudent"`
	Languages []string `json:"Languages"`
	Address   Address  `json:"Address"`
}

type Address struct {
	City       string `json:"City"`
	PostalCode string `json:"PostalCode"`
}

func main() {
	// fmt.Println("NewTextHandler")
	// // Создание текстового обработчика и логгера
	// handler := slog.NewTextHandler(os.Stdout, nil)
	// logger := slog.New(handler)

	// // Запись простого сообщения с уровнем Info
	// logger.Info("Application started", "version", "1.0.0")

	// // Запись сообщения уровня Debug с дополнительными данными
	// logger.Debug("Debugging details", "module", "auth", "userID", 12345)

	// // Запись сообщения уровня Error
	// logger.Error("Failed to connect to database", "dbHost", "localhost", "retries", 3)

	// fmt.Println("\nNewJSONHandler")
	// handler2 := slog.NewJSONHandler(os.Stdout, nil)
	// logger2 := slog.New(handler2)

	// logger2.Info("Server started", "port", 8080)

	// fmt.Println("\nNewJSONHandler by levels")
	// logger3 := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelWarn}))
	// logger3.Info("This message will not appear") // не будет записано
	// logger3.Warn("This is a warning message")    // будет записано
	// logger3.Error("This is an error message")    // будет записано

	mux := http.NewServeMux()
	mux.HandleFunc("POST /orders", handler.PostOrders)
	mux.HandleFunc("GET /orders", handler.GetOrders)
	mux.HandleFunc("GET /orders/{id}", handler.GetOrdersID)
	mux.HandleFunc("PUT /orders/{id}", handler.PutOrdersID)
	mux.HandleFunc("DELETE /orders/{id}", handler.DeleteOrdersID)
	mux.HandleFunc("POST /orders/{id}/close", handler.PostOrdersIDClose)

	mux.HandleFunc("POST /menu", handler.PostMenu)
	mux.HandleFunc("GET /menu", handler.GetMenu)
	mux.HandleFunc("GET /menu/{id}", handler.GetMenuID)
	mux.HandleFunc("PUT /menu/{id}", handler.PutMenuID)
	mux.HandleFunc("DELETE /menu/{id}", handler.DeleteMenuID)

	mux.HandleFunc("POST /inventory", handler.PostInventory)
	mux.HandleFunc("GET /inventory", handler.GetInventory)
	mux.HandleFunc("GET /inventory/{id}", handler.GetInventoryID)
	mux.HandleFunc("PUT /inventory/{id}", handler.PutInventoryID)
	mux.HandleFunc("DELETE /inventory/{id}", handler.DeleteInventoryID)

	mux.HandleFunc("GET /reports/total-sales", handler.GetTotalSales)
	mux.HandleFunc("GET /reports/popular-items", handler.GetPopularItems)

	log.Fatal(http.ListenAndServe(":7070", mux))
}
