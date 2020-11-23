package leego

import (
	"net/http"
	"testing"
)

func TestEngine(t *testing.T) {
	router := NewEngine()

	router.GET("/get", func(resp http.ResponseWriter, req *http.Request) {
		resp.WriteHeader(200)
		resp.Write([]byte("GET"))
	})
	router.POST("/post", func(resp http.ResponseWriter, req *http.Request) {
		resp.WriteHeader(200)
		resp.Write([]byte("post"))
	})
	router.PUT("/put", func(resp http.ResponseWriter, req *http.Request) {
		resp.WriteHeader(200)
		resp.Write([]byte("put"))
	})
	router.PATCH("/patch", func(resp http.ResponseWriter, req *http.Request) {
		resp.WriteHeader(200)
		resp.Write([]byte("patch"))
	})
	router.DELETE("/delete", func(resp http.ResponseWriter, req *http.Request) {
		resp.WriteHeader(200)
		resp.Write([]byte("delete"))
	})
	router.HEAD("/head", func(resp http.ResponseWriter, req *http.Request) {
		resp.WriteHeader(211)
	})
	router.OPTIONS("/options", func(resp http.ResponseWriter, req *http.Request) {
		resp.WriteHeader(200)
		resp.Write([]byte("options"))
	})

	router.ListenAndServe(":8080")
}
