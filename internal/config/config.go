package config

import (
	"log"
	"strings"

	"github.com/namsral/flag"
	"github.com/thanhpk/randstr"
)

var Env string
var Domain string
var Port int
var Cors struct {
	IsEnabled      bool
	AllowedOrigins []string
}
var Tokens struct {
	ReadTokens      []string
	ReadWriteTokens []string
}

func NewConfig() {
	flag.StringVar(&Env, "env", "dev", "the server environment")
	flag.StringVar(&Domain, "domain", "http://localhost", "the server domain")
	flag.IntVar(&Port, "port", 8080, "the server port")
	flag.BoolVar(&Cors.IsEnabled, "cors", false, "the server cors")
	allowedOrigins := flag.String("allowed_origins", "*", "the server allowed cors origins (split the origins with ,)")
	readTokens := flag.String("write_tokens", "", "the server's required read tokens (split tokens with ,)")
	readWriteTokens := flag.String("read_write_tokens", "", "the server's required read-write tokens (split tokens with ,)")
	flag.Parse()

	// Set cors options
	Cors.AllowedOrigins = strings.Split(*allowedOrigins, ",")

	// Set Read and write tokens
	Tokens.ReadTokens = strings.Split(*readTokens, ",")
	Tokens.ReadWriteTokens = strings.Split(*readWriteTokens, ",")

	if *readTokens == "" || *readWriteTokens == "" {
		log.Println("Please save the tokens securely!")
	}

	if *readTokens == "" {
		Tokens.ReadTokens = []string{randstr.Hex(128)}
		log.Println("No Read Token provided, a token got automatically generated:", Tokens.ReadTokens[0])
	}

	if *readWriteTokens == "" {
		Tokens.ReadWriteTokens = []string{randstr.Hex(128)}
		log.Println("No Read-Write Token provided, a token got automatically generated:", Tokens.ReadWriteTokens[0])
	}
}
