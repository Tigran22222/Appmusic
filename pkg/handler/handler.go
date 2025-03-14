package handler

import (
	"appmusic/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		albums := api.Group("/albums")
		{
			albums.POST("/", h.createAlbum)
			albums.GET("/", h.getAllAlbums)
			albums.GET("/:id", h.getAlbumById)
			albums.PUT("/:id", h.updateAlbum)
			albums.DELETE("/:id", h.deleteAlbum)

			tracks := albums.Group("/:id/tracks")
			{
				tracks.POST("/", h.createTrackInAlbum)
				tracks.GET("/", h.getAllTracks)
				tracks.GET("/:trackId", h.getTrackByIdFromAlbum)
				tracks.DELETE("/:trackId", h.deleteTrackByIdFromAlbum)
				tracks.PUT("/:trackId", h.updateByIdFromAlbum)

			}
		}
		tracks := api.Group("/tracks")
		{
			tracks.GET("/:trackId", h.getTrackByIdFromAlbum)
			tracks.PUT("/:id", h.updateByIdFromAlbum)
			tracks.DELETE("/:trackId", h.deleteTrackByIdFromAlbum)
		}
	}
	return router
}
