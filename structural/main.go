package main

import (
	"TMPS/structural/nginx"
	"TMPS/structural/server"
	"fmt"
	"net/http"
	"os"
)

func main() {

	nginxServer := nginx.NewNginxServer()
	appStatusURL := "/app/status"
	createuserURL := "/login/user"

	httpCode, body := nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(createuserURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createuserURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(createuserURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createuserURL, httpCode, body)

	http.Handle("/", &server.WithLog{&server.SimpleServer{"Hi world, from a simple server"}, os.Stdout})
	http.ListenAndServe(":3000", nil)
}
