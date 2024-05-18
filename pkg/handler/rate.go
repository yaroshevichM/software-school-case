package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getRate(c *gin.Context) {
	rateAmount, err := h.services.Rate.GetUSDtoUAHRate()

	if err != nil {
		newErrorResponse(c, http.StatusConflict, "Error when recieve rate")
	}

	c.JSON(http.StatusOK, rateAmount)
}
