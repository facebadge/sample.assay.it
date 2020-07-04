package main

import (
	"github.com/fogfish/gurl"
	ƒ "github.com/fogfish/gurl/http/recv"
	ø "github.com/fogfish/gurl/http/send"
)

// TestOk successful
func TestOk() gurl.Arrow {
	return gurl.HTTP(
		ø.GET("https://assay.it"),
		ƒ.Code(200),
	)
}

//
func main() {}
