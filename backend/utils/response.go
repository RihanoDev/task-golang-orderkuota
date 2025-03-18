package utils

import (
	"encoding/json"
	"net/http"
	"orderkuota/dto"

	"github.com/rs/zerolog/log"
)

func ResponseJSON(w http.ResponseWriter, data interface{}, code int, status string, message string) {
	resp := dto.SuccessResponseDTO[interface{}]{
		Status:  status,
		Code:    code,
		Message: message,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error encoding response" + err.Error()))

		log.Error().Err(err).Str("error", err.Error()).Msg("Failed to encode response")
		return
	}

	log.Info().
		Str("status", status).
		Int("code", code).
		Str("message", message).
		Msg("Response Success sent successfully")
}

func ErrorResponse(w http.ResponseWriter, code int, status string, message string) {
	resp := dto.ErrorResponseDTO{
		Status:  status,
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		w.WriteHeader(code)
		w.Write([]byte("Error encoding response" + err.Error()))

		log.Error().
			Err(err).
			Str("error", err.Error()).
			Msg("Failed to encode response")
		return
	}

	log.Info().
		Str("status", status).
		Int("code", code).
		Str("message", message).
		Msg("Response Error sent successfully")
}
