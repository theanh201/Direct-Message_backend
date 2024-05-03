module main

go 1.21.8

replace DirectBackend/model => ./model

replace DirectBackend/controller => ./controller

replace DirectBackend/entities => ./entities

require (
	DirectBackend/controller v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.1
)

require (
	DirectBackend/entities v0.0.0-00010101000000-000000000000 // indirect
	DirectBackend/model v0.0.0-00010101000000-000000000000 // indirect
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/gorilla/websocket v1.5.1 // indirect
	golang.org/x/net v0.17.0 // indirect
)
