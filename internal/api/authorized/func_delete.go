package authorized

import (
	"net/http"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/gin-boilerplate/internal/api"
	"github.com/ChangSZ/gin-boilerplate/internal/code"
	"github.com/ChangSZ/gin-boilerplate/pkg/validator"
)

type deleteRequest struct {
	Id string `uri:"id"` // HashID
}

type deleteResponse struct {
	Id int64 `json:"id"` // 主键ID
}

// Delete 删除调用方
// @Summary 删除调用方
// @Description 删除调用方
// @Tags API.authorized
// @Accept json
// @Produce json
// @Param id path string true "hashId"
// @Success 200 {object} deleteResponse
// @Failure 400 {object} code.Failure
// @Router /api/authorized/{id} [delete]
// @Security LoginToken
func (h *handler) Delete(ctx *gin.Context) {
	req := new(deleteRequest)
	res := new(deleteResponse)
	if err := ctx.ShouldBindUri(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, validator.GetValidationError(err).Error())
		return
	}

	ids, err := h.hashids.HashidsDecode(req.Id)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.HashIdsDecodeError, err)
		return
	}

	id := int64(ids[0])

	err = h.service.Delete(ctx, id)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.AuthorizedDeleteError, err)
		return
	}

	res.Id = id
	api.ResponseOK(ctx, res)
}
