package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saadozone/gin-gorm-rest/pkg/utils"
)

func (h *Handler) NoRoute(c *gin.Context) {
	response := utils.ErrorResponse("page not found", http.StatusNotFound, "page not found")
	c.JSON(http.StatusNotFound, response)
}
