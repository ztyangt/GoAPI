package app

import (
	"GoAPI/api/controller/music/netease"
	"GoAPI/api/controller/music/tencent"
	"GoAPI/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MusicAPI(route *gin.RouterGroup) {
	musicRoute := route.Group("music")
	{
		musicRoute.GET("/:type", func(ctx *gin.Context) {
			server, err := ctx.GetQuery("server")
			if err != true {
				ctx.JSON(http.StatusBadRequest, common.ReturnData(400, nil, "厂商标识为空"))
			}
			switch server {
			case "netease":
				netease.MusicFunc(ctx)
			case "tencent":
				tencent.MusicFunc(ctx)
			}
		})
	}
}
