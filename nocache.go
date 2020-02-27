// Package nocache provider a baa middleware for http no-cache control.
package nocache

import (
	"time"

	"github.com/go-baa/baa"
)

// Unix epoch time
var epoch = time.Unix(0, 0).Format(time.RFC1123)

// Taken from https://github.com/mytrile/nocache
var noCacheHeaders = map[string]string{
	"Expires":         epoch,
	"Cache-Control":   "no-cache, private, max-age=0",
	"Pragma":          "no-cache",
	"X-Accel-Expires": "0",
}

var etagHeaders = []string{
	"ETag",
	"If-Modified-Since",
	"If-Match",
	"If-None-Match",
	"If-Range",
	"If-Unmodified-Since",
}

// New returns a baa middleware for http no-cache control
func New() baa.HandlerFunc {
	return func(c *baa.Context) {
		// Set our NoCache headers
		for k, v := range noCacheHeaders {
			c.Resp.Header().Set(k, v)
		}

		c.Next()
	}
}

// NewFunc returns a baa HandlerFunc for http no-cache control
func NewFunc() baa.HandlerFunc {
	return func(c *baa.Context) {
		// Set our NoCache headers
		for k, v := range noCacheHeaders {
			c.Resp.Header().Set(k, v)
		}
	}
}
