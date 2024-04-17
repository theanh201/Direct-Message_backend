module controller

go 1.21.8

replace DirectBackend/entities => .././entities

replace DirectBackend/model => .././model

require (
	DirectBackend/entities v0.0.0-00010101000000-000000000000
	DirectBackend/model v0.0.0-00010101000000-000000000000
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
)
