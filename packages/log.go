package packages

import (
	"log"
)

func init() {
	Packages["log"] = map[string]interface{}{
		"Fatal":     log.Fatal,
		"Fatalf":    log.Fatalf,
		"Fatalln":   log.Fatalln,
		"Flags":     log.Flags,
		"New":       log.New,
		"Output":    log.Output,
		"Panic":     log.Panic,
		"Panicf":    log.Panicf,
		"Panicln":   log.Panicln,
		"Prefix":    log.Prefix,
		"Print":     log.Print,
		"Printf":    log.Printf,
		"Println":   log.Println,
		"SetFlags":  log.SetFlags,
		"SetOutput": log.SetOutput,
		"SetPrefix": log.SetPrefix,
	}
}
