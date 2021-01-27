package main

import (
	"fmt"

	"github.com/vroomy/common"
)

const (
	defaultMethods        = "POST, GET, OPTIONS, PUT, DELETE"
	defaultAllowedHeaders = "Accept, Content-Type, Content-Length, Accept-Encoding, Origin, X-UTC-OFFSET"
)

// Init is called when Vroomy initializes the plugin
func Init(env map[string]string) (err error) {
	return
}

// CORs will enable CORs
func CORs(args ...string) (h common.Handler, err error) {
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

func newHandler(url, methods, allowedHeaders string) common.Handler {
	return func(ctx common.Context) {
		hdr := ctx.Writer().Header()
		hdr.Set("Access-Control-Allow-Origin", url)
		hdr.Set("Access-Control-Allow-Methods", methods)
		hdr.Set("Access-Control-Allow-Headers", allowedHeaders)
		return
	}
}
