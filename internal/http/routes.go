package http

import "avito-shop-service/internal/http/middleware"

func (s *Server) initRoutes() {
	s.router.Use(middleware.ErrorHandler())

	s.router.POST("/api/auth", s.handlers.UserHandler.CreateOrAuthUser)

	auth := s.router.Group("/api")
	auth.Use(middleware.AuthWithJWT(s.cfg))
	{
		auth.GET("/info", s.handlers.UserHandler.GetUserInfo)
		auth.GET("/buy/:id", s.handlers.PurchaseHandler.CreatePurchase)
		auth.POST("/sendCoin", s.handlers.TransactionHandler.CreateTransaction)
	}
}
