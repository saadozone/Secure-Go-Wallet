package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saadozone/gin-gorm-rest/internal/dto"
	"github.com/saadozone/gin-gorm-rest/internal/model"
	"github.com/saadozone/gin-gorm-rest/pkg/utils"
)

func (h *Handler) Profile(c *gin.Context) {
	user := c.MustGet("user").(*model.User)

	input := &dto.WalletRequestBody{}
	input.UserID = int(user.ID)
	wallet, err := h.walletService.GetWalletByUserId(input)
	if err != nil {
		response := utils.ErrorResponse("show profile failed", http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formattedUser := dto.FormatUserDetail(user, wallet)
	response := utils.SuccessResponse("show profile success", http.StatusOK, formattedUser)
	c.JSON(http.StatusOK, response)
}
