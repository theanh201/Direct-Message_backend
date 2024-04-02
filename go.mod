module main

go 1.21.8

replace DirectBackend/api => ./api

replace DirectBackend/db => ./db

require (

)

require (
	DirectBackend/db v0.0.0-00010101000000-000000000000
	DirectBackend/api v0.0.0-00010101000000-000000000000
	filippo.io/edwards25519 v1.1.0
	github.com/go-sql-driver/mysql v1.8.1
	github.com/gorilla/mux v1.8.1
)
