package base

import (
	"math/rand"
	"sync"
)

// the client build is among other things used by the server to determine the client's capabilities.
// any forks should use the build number as an indication of compatibility with mainline, rather than an internal build number.
const (
	CLIENT_BUILD        = 134
	CLIENT_KEY_LENGTH   = 20
	MAX_KEY_TIME_DRIFT  = 300
	MAX_CONNECTION_BASE = 20
	TCP_PACKET_SIZE     = 1460
)

const (
	CLIENT_VERSION        = "1.4.2"
	CLIENT_RPC_PROTOCOL   = "http://"
	CLIENT_RPC_HOST       = "rpc.hentaiathome.net"
	CLIENT_RPC_FILE       = "clientapi13.php?"
	CLIENT_LOGIN_FILENAME = "client_login"
)

const (
	CONTENT_TYPE_DEFAULT = "text/html; charset=iso-8859-1"
	CONTENT_TYPE_OCTET   = "application/octet-stream"
	CONTENT_TYPE_JPG     = "image/jpeg"
	CONTENT_TYPE_PNG     = "image/png"
	CONTENT_TYPE_GIF     = "image/gif"
	CONTENT_TYPE_WEBM    = "video/webm"
)

var rpcServerLock sync.Mutex
var rpcServers []string
var rpcServerCurrent, rpcServerLastFailed string

var clientKey string
var clientID, serverTimeDelta int

func GetRPCServerHost() string {
	rpcServerLock.Lock()
	defer rpcServerLock.Unlock()

	if rpcServerCurrent == "" {
		nServers := len(rpcServers)

		switch nServers {
		case 0:
			return CLIENT_RPC_HOST
		case 1:
			rpcServerCurrent = rpcServers[0]
		default:
			var rpcServerSelector int = rand.Intn(nServers)
			var scanDirection int

			if scanDirection = -1; rand.Float32() < 0.5 {
				scanDirection = 1
			}

			for {
				candidate := rpcServers[(rpcServerSelector+nServers)%nServers]

				if candidate == rpcServerLastFailed {
					Debug(rpcServerLastFailed, "was marked as last failed")
					rpcServerSelector += scanDirection
					continue
				}

				rpcServerCurrent = candidate

				Debug("Selected rpcServerCurrent=" + rpcServerCurrent)

				break
			}
		}

	}

	return rpcServerCurrent
}
