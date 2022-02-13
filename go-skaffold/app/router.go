package app

func mapRoutes(s *server) {
	// Ping
	s.router.Get("/ping", s.pingHandler.Get())
}
