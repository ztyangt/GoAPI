package video

import (
	"GoAPI/common"
	"GoAPI/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func videoGET(url string, params map[string]any) (result helper.CurlResult) {

	var headers = make(map[string]any)

	headers["Referer"] = "https://www.douyin.com/"
	headers["cookie"] = "douyin.com; __ac_nonce=0640c590e0055a7762c5a; __ac_signature=_02B4Z6wo00f01xfHYmAAAIDCnIz5guZX5YcX52bAAKIH95; ttwid=1%7CgOM77GJEWySpR1H3GC3divVW-IuN4dwVl1MH5sy88-I%7C1678530831%7C14a3d223290290bebb7de89402be8c1a6c76d909c4e7f7a7b412bd385d831415; home_can_add_dy_2_desktop=%220%22; strategyABtestKey=%221678530833.708%22; passport_csrf_token=5f03dcbad9659e5814d4cb0e42b2947f; passport_csrf_token_default=5f03dcbad9659e5814d4cb0e42b2947f; s_v_web_id=verify_lf3tx4un_dIA9t3Dl_lvJQ_4PU9_9Ous_EPsfwKjvJWIM; bd_ticket_guard_client_data=eyJiZC10aWNrZXQtZ3VhcmQtdmVyc2lvbiI6MiwiYmQtdGlja2V0LWd1YXJkLWNsaWVudC1jc3IiOiItLS0tLUJFR0lOIENFUlRJRklDQVRFIFJFUVVFU1QtLS0tLVxyXG5NSUlCRHpDQnRRSUJBREFuTVFzd0NRWURWUVFHRXdKRFRqRVlNQllHQTFVRUF3d1BZbVJmZEdsamEyVjBYMmQxXHJcbllYSmtNRmt3RXdZSEtvWkl6ajBDQVFZSUtvWkl6ajBEQVFjRFFnQUV0TldaREFVaVAxZExuRnR1VUx3OEtQbkJcclxuZDdvNU53V3ozVXMwR08vOWJvOGVVNXg5R3VaZzVoSVVSelM5VnJLU3EzRG82YjhDOWZkWlcwbk9PYTFxdmFBc1xyXG5NQ29HQ1NxR1NJYjNEUUVKRGpFZE1Cc3dHUVlEVlIwUkJCSXdFSUlPZDNkM0xtUnZkWGxwYmk1amIyMHdDZ1lJXHJcbktvWkl6ajBFQXdJRFNRQXdSZ0loQUtrNXYrczhpQkVkak9pWWs4cFRkM0hSekpMVHBOV2l6cDg4c1VKdmZ1SS9cclxuQWlFQTY2S0pqc0RBZ1hlaXJlMjBBUkNKUGpvRGdYK2Rwb2UvTTV1b0QyOUhoYWM9XHJcbi0tLS0tRU5EIENFUlRJRklDQVRFIFJFUVVFU1QtLS0tLVxyXG4ifQ==; msToken=HiWhRV6SkPuMTefqLVTmkMzmsVzXZX9eLXiQzX7vA1EXIqmzW_9ZhLrVxfCePOMmp-ZCLSk1RbcF8k5X39ZGPwkhYZF6AK3yKIb-pHKTC4E=; msToken=U0Y5NBIQs1FzL4zf3A9qUXTHSATU5AT0qwRDLcVNYp-dGl6yJzUJABJlgSbMbUOMp2AGDuXh9MTZdeTyOQhv8K_mXh0q79RiNzWvr5WtAWFTVGbG2T0GuAcuzbFzOQ==; ttcid=2fd95b10a9f34221908b7d1e015fc35b15; tt_scid=4NoT5NyVRyUA7.XCYVILadREPVI48hUIn9oy1R.vU0ogkNJ5cElc5nolHmNgAqqo0c8b"
	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36"

	return helper.HTTP.GET(url, params, headers)
}

func DouyinFunc(ctx *gin.Context) {
	id, _ := ctx.GetQuery("id")
	{
		data := GetDouyin(id)
		ctx.JSON(http.StatusOK, common.ReturnData(200, helper.Json.Decode(data), "请求成功"))
	}
}
