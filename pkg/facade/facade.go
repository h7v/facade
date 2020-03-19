package facade

import (
	"facade/pkg/facade/serverFunctionality"
	"strings"
)

func NewServer() *Server {
	return &Server{
		connectClients: &serverFunctionality.ConnectClients{},
		showStatistics: &serverFunctionality.ShowStatistics{},
	}
}

// Server implements server and facade.
type Server struct {
	connectClients *serverFunctionality.ConnectClients
	showStatistics *serverFunctionality.ShowStatistics
}

// Work returns that server must do.
func (s *Server) Work() string {
	result := []string{
		s.connectClients.Connect(),
		s.showStatistics.Show(),
	}
	return strings.Join(result, "\n")
}
