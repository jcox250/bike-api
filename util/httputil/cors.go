package httputil

import "github.com/rs/cors"

var AllowedOrigins []string
var AllowedHeaders []string
var AllowedMethods []string

func NewCors() *cors.Cors {
	c := cors.New(cors.Options{
		AllowedOrigins: AllowedOrigins,
		AllowedHeaders: AllowedHeaders,
		AllowedMethods: AllowedMethods,
	})
	return c
}
