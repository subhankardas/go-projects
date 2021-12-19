package api

import (
	"encoding/json"
	"go-projects/microservices/students-service/model"
	"net/http"
)

const (
	HEALTHY string = "HEALTHY"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.BaseResponse{Status: HEALTHY, Code: http.StatusOK})
}
