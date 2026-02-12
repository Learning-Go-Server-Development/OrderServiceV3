package security

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"

	px "github.com/GolangToolKits/go-http-proxy"
	jv "github.com/Ulbora/GoAuth2JwtValidator"
)

type OAuth2Security struct {
	ValadationHost string
	Proxy          px.Proxy
	ClientID       int64
}

func (s *OAuth2Security) ValidateToken(claim *jv.Claim, r *http.Request) bool {
	var rtn bool
	var token string
	tokenHeader := r.Header.Get("Authorization")
	tokenArray := strings.Split(tokenHeader, " ")
	if len(tokenArray) == 2 {
		token = tokenArray[1]
	}

	var vreq ValidationReq
	vreq.ClientID = s.ClientID
	vreq.Role = claim.Role
	vreq.URL = claim.URL
	vreq.AccessToken = token
	aJSON, err := json.Marshal(vreq)

	req, err := http.NewRequest(http.MethodPost, s.ValadationHost+"/rs/token/validate", bytes.NewBuffer(aJSON))
	req.Header.Set("Content-Type", "application/json")
	if err == nil {
		var resp ValidationResp
		suc, stat := s.Proxy.Do(req, &resp)
		if suc && stat == http.StatusOK {
			rtn = resp.Valid
		}
	}
	return rtn
}

func (s *OAuth2Security) New() Security {
	return s
}
