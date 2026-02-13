package security_test

import (
	"net/http"
	"testing"

	px "github.com/GolangToolKits/go-http-proxy"
	"github.com/Learning-Go-Server-Development/OrderServiceV3/security"
	jv "github.com/Ulbora/GoAuth2JwtValidator"
)

func TestOAuth2Security_ValidateToken(t *testing.T) {
	var token = "eNq8lM1y6jgQhd8o5Z8A8ZI4WEhjmcI/stAOSSTIlgkzHmzLTz8lhtzK4qZgcctLV3WfPudzqw8GHTkQaqMQLEboJgq28JTORAjnsD5TEqLg6WDQdkd1y72ohtV5ARs9CmMLZLelxNm/QRWHaJQltEIVB1rzU/q+85FmgFwk0Ee+xrbmKMutrVlzJ3Il+DBxk3Q8gy1syLMdinPh43w5JKPocdirfRk5sPq03wN+Wzk4X/bvWa8kCP5h5aw40FcNq0+1K5OK0WSEqlc7bzizcuZkdlZVmDhEJ9FEF2lDNUm7L8lFRkElGlLbun0ZtcLY4KjjgJCi1iOszvxg0NstS8a9QWeF1Vu5OOtV0QQty+AcnsjI6DX/V+0rB0MnPVITgtCtZ4Zz3ML6NnsVtRwEPitTUtS12ijYiRPu9jTVPAtcAdKjBMV3TZSGcJ7kKIjXrdre8l09FHp15ZOvWqit/8LydYVHjGXxVbsth5b7UrOIZNmtZ5MvvZ98W6axj8bYCwwDxMReVLMQfstw1XDxyWn/LKftY5wonoKP/zAfn/QMRA7LpmBUPLpLNSsHPdEuPT++S1LLSXZpd5fTziejBMG/jMKOecT5nvN/ncK8Z8Pil050vRFXhjgvhjhEV582o6SJFgr+MFPfeuohMT+wtNlUYIQJKkkTh3uuFio4MpAae9sYxVN57DfZXY+daFJ9Y2hEQy7T/NOPu/9UrFHH1qSaxs/yrh/uo5pRNMb+61H42wne4mp4/C2mmgMy0d2Cj96tIwO/50Sp89SPs9Z/cZ6Z/MzDhvR/064RL86idJNxc/pQi81hns1i8leI/wsAAP///yQN5w=="

	var ss security.MockOAuth2Security

	//var ss security.OAuth2Security
	ss.ClientID = 52
	ss.ValidationHost = "http://www.goauth2.com"
	ss.Proxy = &px.GoProxy{}
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		claim      *jv.Claim
		r          *http.Request
		authHeader string
		want       bool
		mockResp   bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			claim: &jv.Claim{
				Role: "user",
				URL:  "/rs/order/add",
			},
			want:       true,
			authHeader: "Bearer " + token,
			mockResp:   true,
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
			mockResp:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			ss.MockValid = tt.mockResp
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
