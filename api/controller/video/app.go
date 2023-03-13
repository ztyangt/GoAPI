package video

import (
	"GoAPI/common"
	"regexp"
)

func GetDouyin(id string) common.Response {

	redict_data := DouyinGET("https://v.douyin.com/"+id, nil)
	re := regexp.MustCompile(`[\d]+`)
	mid := re.FindString(redict_data.Header.Get("Location"))

	url := "https://www.douyin.com/aweme/v1/web/aweme/detail?aweme_id=" + mid + "&aid=6383&version_code=19990127"
	Xbogus := GetBogus("aweme_id=" + mid + "&aid=6383&version_code=19990127")

	new_url := url + "&X-Bogus=" + Xbogus

	r := DouyinGET(new_url, nil)
	result := format_douyin(r.Body)

	return result
}
