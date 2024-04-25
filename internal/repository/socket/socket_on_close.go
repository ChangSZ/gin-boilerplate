package socket

import "github.com/ChangSZ/gin-boilerplate/pkg/log"

func (s *server) OnClose() {
	err := s.socket.Close()
	if err != nil {
		log.Error("socket on closed error: ", err)
	}
}
