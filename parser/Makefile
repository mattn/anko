all : parser.go

parser.go : parser.go.y
	go tool yacc -o $@ parser.go.y
