module gomrp

go 1.18

replace scheduler => ..\\scheduler

require github.com/microsoft/go-mssqldb v0.13.2

require (
	github.com/golang-sql/civil v0.0.0-20190719163853-cb61b32ac6fe // indirect
	github.com/golang-sql/sqlexp v0.0.0-20170517235910-f1bb20e5a188 // indirect
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897 // indirect
)
