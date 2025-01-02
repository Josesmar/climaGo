package inbouind

import (
	"clima-cep/internal/domain"
	"encoding/json"
	"log"
	"math"
	"net/http"

	"github.com/gorilla/mux"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type TemperatureResponse struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func RegisterHandler(r *mux.Router) {
	r.HandleFunc("/climate/{zipcode}", handleClimate).Methods(http.MethodGet)
}

func handleClimate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	zipcode := vars["zipcode"]

	if len(zipcode) != 8 {
		respondWithError(w, http.StatusUnprocessableEntity, "invalid zipcode")
		return
	}

	climateService := domain.NewClimateService(nil)
	climate, err := climateService.GetClimate(zipcode)
	if err != nil {
		if err == domain.ErrZipcodeNotFound {
			respondWithError(w, http.StatusNotFound, "can not find zipcode")
		} else {
			log.Printf("Error fetching climate: %v", err)
			respondWithError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	response := TemperatureResponse{
		TempC: climate.TempC,
		TempF: math.Round(climate.TempF*100) / 100,
		TempK: climate.TempK,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{Message: message})
}
