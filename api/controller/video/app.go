package video

import (
	"GoAPI/common"
	"regexp"
)

func GetDouyin(id string) common.Response {

	redict_data := DouyinGET("https://v.douyin.com/"+id, nil)
	re := regexp.MustCompile(`[\d]+`)
	mid := re.FindString(redict_data.Header.Get("Location"))

	params := make(map[string]any)
	params["aweme_id"] = mid

	r := DouyinGET("https://www.douyin.com/aweme/v1/web/aweme/detail?aid=6383&version_code=19990127", params)
	result := format_douyin(r.Body)

	return result
}
