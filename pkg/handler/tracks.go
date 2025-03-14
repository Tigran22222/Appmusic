package handler

import (
	"appmusic"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getAllTracks(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	albumId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid album id param")
		return
	}

	tracks, err := h.services.Track.GetAll(userId, albumId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, tracks)
}

func (h *Handler) getTrackByIdFromAlbum(c *gin.Context) {
	albumID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid album ID"})
		return
	}

	trackID, err := strconv.Atoi(c.Param("trackId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid track ID"})
		return
	}

	userId, err := getUserId(c) // если у тебя есть userId
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	track, err := h.services.Track.GetByIdFromAlbum(userId, albumID, trackID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Track not found"})
		return
	}

	c.JSON(http.StatusOK, track)
}

func (h *Handler) updateByIdFromAlbum(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	albumId, err := strconv.Atoi(c.Param("id")) // <- добавлено получение albumId
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid album id param")
		return
	}

	trackId, err := strconv.Atoi(c.Param("trackId")) // <- исправлено, раньше было "id"
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid track id param")
		return
	}

	var input appmusic.UpdateTrackInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Track.UpdateByIdFromAlbum(userId, albumId, trackId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteTrackByIdFromAlbum(c *gin.Context) {
	albumID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid album ID"})
		return
	}

	trackID, err := strconv.Atoi(c.Param("trackId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid track ID"})
		return
	}

	userId, err := getUserId(c) // если у тебя есть userId
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err = h.services.Track.DeleteByIdFromAlbum(userId, albumID, trackID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})

}
func (h *Handler) createTrackInAlbum(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	albumId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid album id param")
		return
	}

	var input appmusic.Track
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Track.Create(userId, albumId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
