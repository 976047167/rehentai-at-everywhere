package base

var shutdown bool = false

type HentaiAtHomeClient struct {
	clientApi     *ClientAPI
	serverHandler *ServerHandler
	httpServer    *HTTPServer
}

func NewClient() *HentaiAtHomeClient {
	client := HentaiAtHomeClient{}
	return &client
}
func (c HentaiAtHomeClient) Run() {
	SetActiveClient(&c)
	c.clientApi = NewClientAPI(&c)
	c.serverHandler = NewServerHandler(&c)
	c.httpServer = NewHTTPServer(&c)
	clientPort := GetClientPort()
	c.httpServer.StartCSnnectionListener(clientPort)
	for {
		if shutdown == true {
			break
		}
		//TODO:
	}
}
