package server

// import (
// 	"context"
// 	"net/http"
// 	"time"

// 	"github.com/labstack/echo/v4"
// )

// type Server struct {
// 	echo.Echo
// 	httpServer	*http.Server
// }

// func New() *Server {
// 	return &Server{}
// }

// func (s *Server) Run(port string, handler http.Handler) error {
// 	s.httpServer = &http.Server{
// 		Addr: ":" + port,
// 		Handler: handler,
// 		MaxHeaderBytes: 1 << 20,
// 		ReadTimeout: 10 * time.Second,
// 		WriteTimeout: 10 * time.Second,
// 	}

// 	return s.httpServer.ListenAndServe()
// }

// func (s *Server) Shutdown(ctx context.Context) error {
// 	return s.httpServer.Shutdown(ctx)
// }