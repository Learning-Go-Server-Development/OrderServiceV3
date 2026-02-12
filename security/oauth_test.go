package security_test

import (
	"net/http"
	"testing"

	px "github.com/GolangToolKits/go-http-proxy"
	"github.com/Learning-Go-Server-Development/OrderServiceV3/security"
	jv "github.com/Ulbora/GoAuth2JwtValidator"
)

func TestOAuth2Security_ValidateToken(t *testing.T) {
	var ss security.OAuth2Security
	ss.ClientID = 52
	ss.ValadationHost = "http://www.goauth2.com"
	ss.Proxy = &px.GoProxy{}
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		claim      *jv.Claim
		r          *http.Request
		authHeader string
		want       bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			claim: &jv.Claim{
				Role: "user",
				URL:  "/rs/order/add",
			},
			want:       true,
			authHeader: "Bearer eNq8lMtyozwQhZ/oT2GInWLp2EaWisvPTcjaWZIvAuG4hjEgnn5KjDOVxaTsxRT77tPnfGr1QaMzA1xGEsF8gLNQwgZekjlfwQWsrgSvkPty0CjeEdUw26tgeX2DtRq4NgWijQm29mso/RUaRAGNUMmAUuySHHcOUhTgmwDqzLaBqTmLIjY1W2Z5MwFO2q/DlqWwgTV+NUODjDvButLhOp4FaSf3hWfB8qMPB95FGeyidd4f004K4P6gxTw/kHcFyw+5K8KSknCAspM7u7/SYm6lZlaZa3+FLrz2bsKEqsNmX+Cb8NyS17gydfvCa7g2wVHLAMZ5pQZYXtlBo/U9S8rsXqW50duMvvLabWgKF/CCB0rG/J+17wz0rbBxhTFC9555kAUNrO6zN17DgOvQIsF5VclIwpZfgnZPEsVSd8ZBchYg/6qJkhVchBly/W0j43u+0UOuNiOfbNNAZfznhu+M21gbFp+1cdE3zBGKejhN7z1RtrS/822Y+g4afNvVFGDt215FV/BLhlFjFlys5t9yip/jRIIp+DhP83FwR4Fn0XQKRvmzu1TRolcT7dLr87sklJhkl3YPOe0cPAjg/qQEttTG1tecv3VyfUz7tz863ngjRoZBlvf+Co0+TUZBQsUl/GamuvdUfai/YWmySVdz7ZaChBazZ4pL90xBos1toySYymMXpQ89trxO1J2h5jW+TfOmp4dvyreopVtcTuNn+dAPc1BFCRp85/3MnXiCv7jpn/+LiWIAT3S34LN360zB3zkRYr2caHmCy/8tqzsNx72Hluy/YnMMhm28ICIJ3GW2KCIVEfm6/BUAAP//8e4NqQ==",
		},
		{
			name: "test 2",
			claim: &jv.Claim{
				Role: "user",
				URL:  "/rs/order/add",
			},
			want: false,
			//bad token
			authHeader: "Bearer eNq8lMtyozwQhZ///DlqWwgTV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			r, _ := http.NewRequest("POST", "/ffllist", nil)
			r.Header.Set("Authorization", tt.authHeader)
			r.Header.Set("Content-Type", "application/json")
			s := ss.New()
			got := s.ValidateToken(tt.claim, r)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("ValidateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
