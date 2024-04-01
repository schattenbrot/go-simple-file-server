package config

import (
	"github.com/namsral/flag"
)

var Port int
var Cors bool

func NewConfig() {
	flag.IntVar(&Port, "port", 8080, "the server port")
	flag.BoolVar(&Cors, "cors", false, "the server cors")
	flag.Parse()
}
