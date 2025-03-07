package handler

import "github.com/gin-gonic/gin"

type Handler struct{}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
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
				tracks.GET("/:id", h.getTrackById)
				tracks.PUT("/:id", h.updateTrack)
				tracks.DELETE("/:id", h.deleteTrack)

			}
		}
		/*tracks := api.Group("/tracks")
		{
			tracks.GET("/:id")
			tracks.PUT("/:id")
			tracks.DELETE("/:id")
		}*/
	}
	return router
}
