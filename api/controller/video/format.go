package video

import (
	"GoAPI/common"

	"github.com/tidwall/gjson"
)

func format_douyin(data string) common.Response {
	exisit := gjson.Get(data, "status_code").Int() == 0
	if exisit {

		music := common.DouyinMusic{
			Title:   gjson.Get(data, "aweme_detail.music.title").String(),
			Author:  gjson.Get(data, "aweme_detail.music.author").String(),
			PlayUrl: gjson.Get(data, "aweme_detail.music.play_url.uri").String(),
		}

		author := common.Creator{
			UserId:    gjson.Get(data, "aweme_detail.author.total_favorited").String(),
			Nickname:  gjson.Get(data, "aweme_detail.author.nickname").String(),
			Signature: gjson.Get(data, "aweme_detail.author.signature").String(),
			AvatarUrl: gjson.Get(data, "aweme_detail.music.avatar_large.url_list.0").String(),
		}

		if gjson.Get(data, "aweme_detail.images.0").Exists() {

			img_slice := []string{}

			img_json := gjson.Get(data, "aweme_detail.images")
			img_json.ForEach(func(index, value gjson.Result) bool {
				img_slice = append(img_slice, value.Get("url_list.3").String())
				return true
			})

			result := common.DouyinInfo{
				Video:  false,
				Images: img_slice,
				Music:  music,
				Author: author,
			}

			return common.ReturnData(200, result, "请求成功")
			// return common.ReturnData(200, helper.Json.Decode(data), "请求成功")
		} else {

			play_url := "https://aweme.snssdk.com/aweme/v1/play/?video_id=" + gjson.Get(data, "aweme_detail.video.play_addr.uri").String() + "&ratio=1080p&line=0"
			redict_data := DouyinGET(play_url, nil)

			url := redict_data.Header.Get("Location")

			video := common.DouyinVideo{
				Duration:    int(gjson.Get(data, "aweme_detail.video.duration").Int()),
				Cover:       gjson.Get(data, "aweme_detail.video.dynamic_cover.url_list.0").String(),
				PlayUrl:     url,
				Description: gjson.Get(data, "aweme_detail.preview_title").String(),
			}

			result := common.DouyinInfo{
				Video:  video,
				Images: false,
				Music:  music,
				Author: author,
			}

			return common.ReturnData(200, result, "请求成功")
		}
	} else {
		return common.ReturnData(400, nil, gjson.Get(data, "status_msg").String())
	}
}
