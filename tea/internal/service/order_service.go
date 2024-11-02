package service

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"tea/models"
)

func PostOrder(order models.Order) {
	writeJson, _ := json.MarshalIndent(order, "", "\t")
	err := os.WriteFile("../data/files.json", writeJson, 0o644)
	if err != nil {
		log.Fatalf("Error writing json file: %v", err)
	}
}

func GetOrder(data []byte, w http.ResponseWriter) {
	w.Header().Set("Content-type", "application/json")
	w.Write(data)
}
