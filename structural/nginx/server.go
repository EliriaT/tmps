package nginx

type iServer interface {
	handleRequest(string, string) (int, string)
}
