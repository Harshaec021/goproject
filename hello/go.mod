module github.com/Harshaec021/goproject/hello

go 1.17

replace github.com/Harshaec021/goproject/morestrings => ../morestrings

require (
	github.com/Harshaec021/goproject/morestrings v0.0.0-00010101000000-000000000000 // indirect
	github.com/Harshaec021/goproject/utils v0.0.0-20221008105338-4c857463dea6 // indirect
	golang.org/x/net v0.0.0-20221004154528-8021a29435af // indirect
)

replace github.com/Harshaec021/goproject/utils => ../utils
