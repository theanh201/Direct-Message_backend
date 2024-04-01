module main

go 1.21.8

replace DirectBackend/api => ./api

require (
	DirectBackend/api v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.1
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
)
