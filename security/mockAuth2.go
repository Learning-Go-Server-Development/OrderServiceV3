package security

import (
	"net/http"

	px "github.com/GolangToolKits/go-http-proxy"
	jv "github.com/Ulbora/GoAuth2JwtValidator"
)

type MockOAuth2Security struct {
	ValidationHost string
	Proxy          px.Proxy
	ClientID       int64
	MockValid      bool
}

func (s *MockOAuth2Security) ValidateToken(claim *jv.Claim, r *http.Request) bool {
	return s.MockValid
}

func (s *MockOAuth2Security) New() Security {
	return s
}
