package nginx

import (
	"TMPS/structural/token"
	"time"
)

type Application struct {
	tokenMaker token.TokenMaker
}

func (a *Application) handleRequest(url, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return 200, "Ok"
	}

	if url == "/login/user" && method == "POST" {
		accessToken, _ := a.tokenMaker.CreateToken("email", time.Duration(120))
		return 200, accessToken
	}
	return 404, "Not Ok"
}
