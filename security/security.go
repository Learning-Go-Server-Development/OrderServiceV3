package security

import (
	"net/http"

	jv "github.com/Ulbora/GoAuth2JwtValidator"
)

type Security interface {
	ValidateToken(claim *jv.Claim, r *http.Request) bool
}
