package api

import (
	"chat_room/protobuf/common"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 默认返回
func HttpResp(c *gin.Context, resp *common.CommonResp, err error) {
	// 非业务上错误
	if err != nil {
		failResp(c, err)
		return
	}
	// 业务错误
	data := map[string]interface{}{}
	json.Unmarshal([]byte(resp.ResultJson), &data)
	// 成功返回
	successResp(c, data)
}

// 成功
func successResp(c *gin.Context, data map[string]interface{}){
	time := time.Now().Unix()
	data["message"] = "success"
	data["error_code"] = "0000"
	data["unix_time"] = time
	c.JSON(http.StatusOK, data)
}


// 失败
func failResp(c *gin.Context, err error) {

}
