package response

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Unauthorized(w http.ResponseWriter, schema any, logger *logrus.Logger) {
	message, err := json.Marshal(schema)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"schema": schema,
		}).Error("Internal Marshal Error")
		return
	}
	sendResponse(w, http.StatusUnauthorized, message)
}

func InternalServerResponse(w http.ResponseWriter, schema any, logger *logrus.Logger) {
	message, err := json.Marshal(schema)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"schema": schema,
		}).Error("InternalServerResponse Marshal Error")
		return
	}
	sendResponse(w, http.StatusInternalServerError, message)
}

func UnauthorizedResponse(w http.ResponseWriter, schema any, logger *logrus.Logger) {
	message, err := json.Marshal(schema)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"schema": schema,
		}).Error("UnauthorizedResponse Marshal Error")
		return
	}
	sendResponse(w, http.StatusUnauthorized, message)
}

func OKResponse(w http.ResponseWriter, schema any, logger *logrus.Logger) {
	message, err := json.Marshal(schema)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"schema": schema,
		}).Error("OKResponse Marshal Error")
		return
	}
	sendResponse(w, http.StatusOK, message)
}

func CreatedResponse(w http.ResponseWriter, schema any, logger *logrus.Logger) {
	message, err := json.Marshal(schema)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"schema": schema,
		}).Error("CreatedResponse Marshal Error")
		return
	}
	sendResponse(w, http.StatusCreated, message)
}

func sendResponse(w http.ResponseWriter, statusCode int, message []byte) {
	w.WriteHeader(statusCode)
	w.Write(message)
}
