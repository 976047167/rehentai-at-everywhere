package base

type HTTPServer struct {
	client *HentaiAtHomeClient
}

func NewHTTPServer(c *HentaiAtHomeClient) *HTTPServer {
	return &HTTPServer{c}
}

func (h HTTPServer) StartCSnnectionListener(port int) {
	//
}
