package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-logr/logr"
	"github.com/imperialhound/friend-foe-api/internal/validator"
)

type ValidatorHandler struct {
	ctx       context.Context
	logger    logr.Logger
	validator *validator.Validator
}

func NewValidatorHandler(ctx context.Context, logger logr.Logger, authServer string) *ValidatorHandler {
	validator := validator.New(ctx, logger, authServer)
	return &ValidatorHandler{
		ctx:       ctx,
		logger:    logger,
		validator: validator,
	}
}

// ValidateToken checks if request has a valid JWT else will return 401
func (v *ValidatorHandler) ValidateToken(w http.ResponseWriter, r *http.Request) {
	logger := v.logger

	// Create empty response map to populate with status
	resp := make(map[string]string)

	logger.Info("validating JWT token")
	// Extract token from header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		msg := "no bearer token found"
		logger.Info(msg)
		http.Error(w, msg, http.StatusUnauthorized)
		return
	}

	token := strings.Split(authHeader, " ")[1]

	// Validate JWT
	valid, err := v.validator.Validate(token)
	if err != nil {
		msg := "failed to validate token"
		logger.Error(err, msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// TODO: Invalid signatures return as error
	// Add step to check if error due to invalid signature
	if valid {
		resp["status"] = "Valid JWT. Welcome in and have a cup of tea"
	} else {
		resp["status"] = "Invalid JWT! Leave before I set the dogs loose!"
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		msg := "failed to marshal responce"
		logger.Error(err, msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.Write(jsonResp)
}
