package plugin

import "github.com/vroomy/common"

func newHandler(url, methods, allowedHeaders string) common.Handler {
	return func(ctx common.Context) {
		hdr := ctx.Writer().Header()
		hdr.Set("Access-Control-Allow-Origin", url)
		hdr.Set("Access-Control-Allow-Methods", methods)
		hdr.Set("Access-Control-Allow-Headers", allowedHeaders)
		return
	}
}
