package httputils

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"test/portal/domain"

	"github.com/go-playground/validator/v10"
)

type response struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func WriteSuccessResponse(w http.ResponseWriter, payload any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responseData := response{
		Data:    payload,
		Message: "Success",
		Status:  http.StatusOK,
	}
	err := json.NewEncoder(w).Encode(responseData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func WriteErrorResponse(w http.ResponseWriter, processErr error) {
	errorData := strings.Split(processErr.Error(), ":")
	status, err := strconv.Atoi(errorData[0])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	errResponse := response{
		Data:    nil,
		Message: processErr.Error(),
		Status:  status,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	err = json.NewEncoder(w).Encode(errResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func ValidateAndUnmarshal[T any](r *http.Request, validate *validator.Validate, data *T) error {
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		return errors.New(domain.ErrBadRequest)
	}

	err = validate.Struct(data)
	if err != nil {
		log.Println("Validate error", err)
		return errors.New(domain.ErrBadRequest + ";;" + err.Error())
	}
	return nil
}

func GetPathParamByPathPosition(r *http.Request, position int) *string {
	urlParts := strings.Split(r.URL.Path, "/")
	if len(urlParts) < position {
		return nil
	}
	paramValue := &urlParts[position]
	log.Println("Find param value", *paramValue)
	return paramValue
}
