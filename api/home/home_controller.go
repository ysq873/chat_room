package home

import (
	"chat_room/api"
	"chat_room/protobuf/common"
	"errors"
	"github.com/gin-gonic/gin"
)

// todo:1 首页
func Index(c *gin.Context) {
	resp := new(common.CommonResp)
	api.HttpResp(c, resp, errors.New(""))
}
