module db

go 1.21.8

replace DirectBackend/api => .././api

require github.com/go-sql-driver/mysql v1.8.1

require filippo.io/edwards25519 v1.1.0 // indirect
