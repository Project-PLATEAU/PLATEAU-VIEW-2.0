package integration

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=server.cfg.yml ../../../schemas/integration.yml

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

var _ StrictServerInterface = (*Server)(nil)
