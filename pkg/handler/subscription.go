package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yaroshevichM/software-school-case/pkg/models"
)

func (h *Handler) createSubscription(c *gin.Context) {
	var input models.CreateSubscriptionInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusConflict, err.Error())
		return
	}

	subscription, err := h.services.Subscription.Create(input)

	if err != nil {
		newErrorResponse(c, http.StatusConflict, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": subscription,
	})
}
