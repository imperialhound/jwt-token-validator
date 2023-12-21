package validator

import (
	"context"
	"strings"
	"time"

	"github.com/MicahParks/keyfunc/v2"
	"github.com/go-logr/logr"
	"github.com/golang-jwt/jwt/v5"
)

type Validator struct {
	logger     logr.Logger
	authServer string
	ctx        context.Context
	options    keyfunc.Options
}

func New(ctx context.Context, logger logr.Logger, authServer string) *Validator {
	options := keyfunc.Options{
		RefreshInterval:  time.Hour,
		RefreshRateLimit: time.Minute * 5,
		RefreshTimeout:   time.Second * 10,
	}

	return &Validator{
		logger:     logger,
		authServer: authServer,
		ctx:        ctx,
		options:    options,
	}
}

// Validate verifies the signature of a JWT utilizing a JWKS from
// a predefined URL and will return true if valid
func (v *Validator) Validate(token string) (bool, error) {
	logger := v.logger

	wellKnowns := "/.well-known/jwks.json"

	JWKSURL := strings.Join([]string{v.authServer, wellKnowns}, "")

	logger.Info("Create JWKS from source", "url", JWKSURL)
	jwks, err := keyfunc.Get(JWKSURL, v.options)
	if err != nil {
		return false, err
	}

	logger.Info("Parsing the JWT")
	parsed, err := jwt.Parse(token, jwks.Keyfunc)
	if err != nil {
		return false, err
	}

	logger.Info("Validating JWT")
	if !parsed.Valid {
		return false, nil
	}
	logger.Info("JWT is valid")
	return true, nil
}
