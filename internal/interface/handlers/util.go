package handlers

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
	"github.com/gin-gonic/gin"
)

func handleErrors(c *gin.Context, errs []util.ProblemDetails) {
	if len(errs) > 0 {
		for _, err := range errs {
			if err.Status == 500 {
				c.JSON(err.Status, gin.H{"error": err})
				return
			} else {
				c.JSON(err.Status, gin.H{"error": err})
				return
			}
		}
	}
}
