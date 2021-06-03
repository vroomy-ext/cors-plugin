package plugin

import (
	"fmt"
	"log"

	"github.com/vroomy/common"
	"github.com/vroomy/plugins"
)

var p Plugin

const (
	defaultMethods        = "POST, GET, OPTIONS, PUT, DELETE"
	defaultAllowedHeaders = "Accept, Content-Type, Content-Length, Accept-Encoding, Origin, X-UTC-OFFSET"
)

func init() {
	if err := plugins.Register("cors", &p); err != nil {
		log.Fatal(err)
	}
}

type Plugin struct {
	plugins.BasePlugin
}

// CORs will enable CORs
func (p *Plugin) CORs(args ...string) (h common.Handler, err error) {
	var (
		url            string
		methods        string
		allowedHeaders string
	)

	switch len(args) {
	case 1:
		url = args[0]
		methods = defaultMethods
		allowedHeaders = defaultAllowedHeaders
	case 2:
		url = args[0]
		methods = args[1]
		allowedHeaders = defaultAllowedHeaders
	case 3:
		url = args[0]
		methods = args[1]
		allowedHeaders = args[2]

	default:
		err = fmt.Errorf("invalid number of arguments, expected 1, 2, or 3 arguments and recieved %d", len(args))
		return
	}

	h = newHandler(url, methods, allowedHeaders)
	return
}
