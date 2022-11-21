package nginx

import "TMPS/structural/token"

type Nginx struct {
	application       *Application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func NewNginxServer() *Nginx {
	pasetoMaker, _ := token.NewPasetoMaker("12345678901234567890123456789012")
	return &Nginx{
		application:       &Application{pasetoMaker},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

func (n *Nginx) HandleRequest(url, method string) (int, string) {
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}
	// aici
	return n.application.handleRequest(url, method)
}

func (n *Nginx) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}
	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}
