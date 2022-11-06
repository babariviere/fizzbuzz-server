package fizzbuzz

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
    Service Service
}

// GetFizzbuzz godoc
//
// @Summary     return a fizzbuzz like string, separated by comma, of custom strings and integers provided by user.
// @Description by default, the parameters will use default fizzbuzz parameters if none are provided.
// @Produce     plain
// @Param       int1  query    int    false "integer for first modulo"  minimum(1)
// @Param       int2  query    int    false "integer for second modulo" minimum(1)
// @Param       str1  query    string false "string for first modulo"   example(fizz)
// @Param       str2  query    string false "string for second modulo"  example(buzz)
// @Param       limit query    int    false "number of fizzbuzz to generate"
// @Success     200   {string} string
// @Failure     400   {string} string
// @Router      /fizzbuzz [get]
func (h *Handler) GetFizzbuzz(c *gin.Context) {
    var req GetFizzbuzzReq

    if err := c.BindQuery(&req); err != nil {
        c.String(http.StatusBadRequest, err.Error())
        return
    }

    req.OrDefault()
    if req.Limit == 0 {
        req.Limit = 100
    }

    resp, err := h.Service.GetFizzbuzz(req)
    if err != nil {
        c.String(http.StatusBadRequest, err.Error())
        return
    }

    c.String(http.StatusOK, resp)
}
