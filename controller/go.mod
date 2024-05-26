module controller

go 1.21.8

replace DirectBackend/entities => .././entities

replace DirectBackend/model => .././model

require (
	DirectBackend/entities v0.0.0-00010101000000-000000000000
	DirectBackend/model v0.0.0-00010101000000-000000000000
)

require github.com/ZEGOCLOUD/zego_server_assistant/token/go/src v0.0.0-20231103072415-8c895c31df9d // indirect

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/gorilla/mux v1.8.1
	github.com/gorilla/websocket v1.5.1
	golang.org/x/net v0.17.0 // indirect
)
