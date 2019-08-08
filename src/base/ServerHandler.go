package base

import (
	"fmt"
	"net/url"
)

const (
	ACT_SERVER_STAT           = "server_stat"
	ACT_GET_BLACKLIST         = "get_blacklist"
	ACT_CLIENT_LOGIN          = "client_login"
	ACT_CLIENT_SETTINGS       = "client_settings"
	ACT_CLIENT_START          = "client_start"
	ACT_CLIENT_SUSPEND        = "client_suspend"
	ACT_CLIENT_RESUME         = "client_resume"
	ACT_CLIENT_STOP           = "client_stop"
	ACT_STILL_ALIVE           = "still_alive"
	ACT_STATIC_RANGE_FETCH    = "srfetch"
	ACT_DOWNLOADER_FETCH      = "dlfetch"
	ACT_DOWNLOADER_FAILREPORT = "dlfails"
	ACT_OVERLOAD              = "overload"
)

func GetServerConnectionURL(act, add string) (serverConnectionURL *url.URL, err error) {
	if act == ACT_SERVER_STAT {
		rawurl := fmt.Sprintf("%s%s/%sclientbuild=%d&act=%s", CLIENT_RPC_PROTOCOL, GetRPCServerHost(), CLIENT_RPC_FILE, CLIENT_BUILD, act)

		serverConnectionURL, err = url.Parse(rawurl)
		if err != nil {
			return
		}

	} else {
		rawurl := fmt.Sprintf("%s%s/%s%s", CLIENT_RPC_PROTOCOL, GetRPCServerHost(), CLIENT_RPC_FILE, GetURLQueryString(act, add))

		serverConnectionURL, err = url.Parse(rawurl)
		if err != nil {
			return
		}
	}

	return
}

func GetURLQueryString(act, add string) string {
	correctedTime := GetServerTime()
	actKey := GetSHA1String(fmt.Sprintf("hentai@home-%s-%s-%d-%d-%s", act, add, clientID, correctedTime, clientKey))
	return fmt.Sprintf("clientbuild=%d&act=%s&add=%s&cid=%d&acttime=%d&actkey=%s", CLIENT_BUILD, act, add, clientID, correctedTime, actKey)
}

type ServerResponse struct {
}

// handle notifications and other communications with the hentai@home server
type ServerHandler struct {
}
