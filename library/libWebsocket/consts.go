package libWebsocket

type WebSocketEvent string

const (
	EventConnectionEstablished WebSocketEvent = "connection.established" // 连接建立
	EventConnectionClosed      WebSocketEvent = "connection.closed"      // 连接关闭

	EventDeviceOnline  WebSocketEvent = "device.online"  // 设备上线
	EventDeviceOffline WebSocketEvent = "device.offline" // 设备下线
	EventDeviceAlarm   WebSocketEvent = "device.alarm"   // 设备告警
)

// EventHandler 消息处理器
type EventHandler func(client *Client, req *WRequest)

type EventHandlers map[WebSocketEvent]EventHandler

const (
	Token = "urbThHk9T//V2yPgUiqRmWUvaOVpMuDr9AafqDSrugwUq+NiFSyWOhay6fG1QcXH15ggHV0vp0Tewik2RkEUR7KAf7KjIKUlr8V2MHOwFAOCab68bTulpoculmAangQwzFo81kWTZnDlu5pf9bhvi6c3J8oKKPIFcYVdQ5FxmA3owniKsJKSfL1uFwkN+9oyB4qK732GlUkmG5QKJgEUO0RpK+xvLoxuXWerGuPfreXErNm7sP64+0u8a88ewU6M7LiA0qSwNUd+tdgEq/D3PARnL/wB6uoyOEJ9yoBmBKv5woCAFAc5TqHzdZFPQ7ejyhIgjvuwWHabgI/fukYmHd7k+d9ev17Ohe7I0ImkV65uLBz5XSjwEsDcyd0PW3cS5jTALiMKutiNAC8IpbdP46/bjiP64D0KpSkNjjvtMAhuNpCfTiHDLX0yyenNp3rtcfe8yPJhCqQbdi8WMj1lBFwWVz/JgTEhdd7pe6TJ7jgHUfq8XpJu9nECAngH32J9C+DxDVmlJruJ94hW5ZPAf9ZsK5/+It4bK45T+81FZkf78ylPB10jed36xlQ01MOsd+FketIl7+2vfdYCXppv0I2+0Kv/7vglWygwE+NXR1W9KfFpOM5+tzUiP98szmmZVp1m0Hho76wMHzOrta18c4jLsGTrG8wgvkY34hrBlt1pevXdC8BywDQqenZh29JYSSItESgypOUkIZ/zk8ciaYLZAC+dMywvkfZesL9OXMbP89G7vB3Gbzl5HlU7G3z27qVbzB8Pfmdcc3CrQDoxr91kNxMm0Kc+pS1Mg8cdYeMI0n3dRFKRY2rDoLBv4eP+F/U1HSjVBV4JkDaZrBB0ROM8D6yF3IH8RUHcUpTF4tv6qm33F335VBmFbXzch9Dzjv0djqScs8cker1N6uJWuQ=="
)

type DataConnectionEstablished struct {
	ClientID         string `json:"client_id"`
	UserID           string `json:"user_id"`
	UserName         string `json:"user_name"`
	FirstConnectTime int64  `json:"first_connect_time"`
}
