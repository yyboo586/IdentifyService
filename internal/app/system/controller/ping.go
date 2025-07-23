package controller

import (
	"IdentifyService/library/libWebsocket"
)

var Ping = new(pingController)

type pingController struct{}

func (c *pingController) Ping(client *libWebsocket.Client, req *libWebsocket.WRequest) {
	libWebsocket.SendSuccess(client, req.Event)
}
