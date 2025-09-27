package example

import (
	"context"

	"github.com/Dizzrt/ellie/transport/http"
	"github.com/gin-gonic/gin"
)

type ExampleHTTPServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
}

func RegisterExampleHTTPServer(s *http.Server, srv ExampleHTTPServer) {
	r := s.Engine()

	r.GET("/example/hello/:name", _Example_Hello_HTTP_Handler(s, srv))

	s.Handler = r
}

func _Example_Hello_HTTP_Handler(hs *http.Server, srv ExampleHTTPServer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req HelloRequest
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, hs.WrapHTTPResponse(nil, err))
			ctx.Abort()
			return
		}

		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, hs.WrapHTTPResponse(nil, err))
			ctx.Abort()
			return
		}

		resp, err := srv.Hello(ctx.Request.Context(), &req)
		if err != nil {
			ctx.JSON(http.HTTPStatusCodeFromError(err), hs.WrapHTTPResponse(resp, err))
			ctx.Abort()
			return
		}

		hs.EncodeResponse(ctx, resp, err)
	}
}
