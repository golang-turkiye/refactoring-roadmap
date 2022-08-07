package response

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

func InternalServerResponse(w http.ResponseWriter, schema any, logger *logrus.Logger) {
	message, err := json.Marshal(schema)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"schema": schema,
		}).Error("InternalServerResponse Marshal Error")
		return
	}
	if err := sendResponse(w, http.StatusInternalServerError, message); err != nil {
		logger.Error("Response can not be wrote")
	}
}

func UnauthorizedResponse(w http.ResponseWriter, schema any, logger *logrus.Logger) {
	message, err := json.Marshal(schema)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"schema": schema,
		}).Error("UnauthorizedResponse Marshal Error")
		return
	}
	if err := sendResponse(w, http.StatusUnauthorized, message); err != nil {
		logger.Error("Response can not be wrote")
	}
}

func OKResponse(w http.ResponseWriter, schema any, logger *logrus.Logger) {
	message, err := json.Marshal(schema)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"schema": schema,
		}).Error("OKResponse Marshal Error")
		return
	}
	if err := sendResponse(w, http.StatusOK, message); err != nil {
		logger.Error("Response can not be wrote")
	}
}

func CreatedResponse(w http.ResponseWriter, schema any, logger *logrus.Logger) {
	message, err := json.Marshal(schema)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"schema": schema,
		}).Error("CreatedResponse Marshal Error")
		return
	}
	if err := sendResponse(w, http.StatusCreated, message); err != nil {
		logger.Error("Response can not be wrote")
	}
}

func sendResponse(w http.ResponseWriter, statusCode int, message []byte) error {
	w.WriteHeader(statusCode)
	if _, err := w.Write(message); err != nil {
		return err
	}
	return nil
}
