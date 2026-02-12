package security

import (
	"net/http"

	jv "github.com/Ulbora/GoAuth2JwtValidator"
)

type ValidationReq struct {
	ClientID    int64  `json:"clientId"`
	Role        string `json:"role"`
	URL         string `json:"url"`
	AccessToken string `json:"accessToken"`
}

type ValidationResp struct {
	Valid bool `json:"valid"`
}

type Security interface {
	ValidateToken(claim *jv.Claim, r *http.Request) bool
}
