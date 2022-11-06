package fizzbuzz

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup, h *Handler) {
    r.GET("/fizzbuzz", h.GetFizzbuzz)
}
