package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// WebSocketServer 是一个WebSocket服务器
type WebSocketServer struct {
	upgrader websocket.Upgrader
	clients  map[*websocket.Conn]bool
}

// NewWebSocketServer 创建一个新的 WebSocketServer
func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				// 验证请求的Origin是否在允许的列表中，这里允许所有的Origin
				return true
			},
		},
		clients: make(map[*websocket.Conn]bool),
	}
}

// handleWebSocket 处理WebSocket连接
func (wss *WebSocketServer) HandleWebSocket(resp http.ResponseWriter, req *http.Request) {
	// 将HTTP连接升级为WebSocket连接
	conn, err := wss.upgrader.Upgrade(resp, req, nil)
	if err != nil {
		log.Printf("Error upgrading to WebSocket: %v\n", err)
		return
	}

	// 发送 "OK" 消息给客户端
	err = conn.WriteMessage(websocket.TextMessage, []byte("OK"))
	if err != nil {
		log.Printf("Error sending 'OK' message to client: %v\n", err)
		// 处理发送消息失败的情况
		// ...
	}

	// 将新连接加入客户端列表
	wss.clients[conn] = true

	// 读取并处理WebSocket消息
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message from WebSocket: %v\n", err)
			// 处理读取消息错误的情况
			// ...
			break
		}
	}

	// 关闭连接并从客户端列表中删除
	conn.Close()
	delete(wss.clients, conn)
}

// BroadcastMessage 广播消息给所有连接的客户端
func (wss *WebSocketServer) BroadcastMessage(message []byte) {
	for client := range wss.clients {
		err := client.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Printf("Error sending message to client: %v\n", err)
			// 处理发送消息失败的情况
			// ...
		}
	}
}
