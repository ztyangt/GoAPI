package tencent

import (
	"GoAPI/common"
	"strconv"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

func FormatSong(data string) common.Response {

	exist := gjson.Get(data, "data.0").Exists()

	if exist {
		var artist = []string{}
		id := gjson.Get(data, "data.0.mid").String()
		name := gjson.Get(data, "data.0.name").String()

		artist_str := gjson.Get(data, "data.0.singer")
		artist_str.ForEach(func(index, value gjson.Result) bool {
			artist = append(artist, value.Get("name").String())
			return true
		})

		album := gjson.Get(data, "data.0.album.name").String()
		mv_id := gjson.Get(data, "data.0.mv.vid").String()

		pic_str := "https://y.gtimg.cn/music/photo_new/T002R800x800M000" + gjson.Get(data, "data.0.album.mid").String() + ".jpg"

		return common.ReturnData(200, common.Song{
			Id:      id,
			Name:    name,
			Artist:  artist,
			Album:   album,
			PicUrl:  pic_str,
			UrlId:   id,
			LyricId: id,
			MvId:    mv_id,
			Source:  "tencent",
		}, "请求成功")
	} else {
		return common.ReturnData(400, nil, "请求失败")
	}
}

func FormatUrl(data string, id string) common.Response {

	exist := gjson.Get(data, "data.0").Exists()

	if exist {

		uid := time.Now().Unix()
		filename := ""
		media_mid := gjson.Get(data, "data.0.file.media_mid").String()

		result := common.Url{
			Id:       id,
			Size:     0,
			Url:      "",
			Duration: 0,
		}

		var ty [7][4]string = [7][4]string{
			{"size_flac", "999", "F000", `flac",`},
			{"size_320mp3", "320", "M800", `mp3",`},
			{"size_192aac", "192", "C600", `m4a",`},
			{"size_128mp3", "128", "M500", `mp3",`},
			{"size_96aac", "96", "C400", `m4a",`},
			{"size_48aac", "48", "C200", `m4a",`},
			{"size_24aac", "24", "C100", `m4a"`},
		}
		for i := range ty {
			filename = filename + `"` + ty[i][2] + media_mid + "." + ty[i][3]
		}
		param_str := `{"req_0":{"module":"vkey.GetVkeyServer","method":"CgiGetVkey","param":{"guid":"` + strconv.FormatInt(uid, 10) + `","loginflag":1,"platform":"20","songmid":["` + id + `","` + id + `","` + id + `","` + id + `","` + id + `","` + id + `","` + id + `"],"filename":[` + filename + `],"songtype":[0,0,0,0,0,0,0]}}}`

		params := make(map[string]any)
		params["needNewCode"] = "0"
		params["platform"] = "yqq.json"
		params["format"] = "json"
		params["data"] = param_str
		resp := musicGET("https://u.y.qq.com/cgi-bin/musicu.fcg", params)

		vkeys := gjson.Get(resp.Body, "req_0.data.midurlinfo")
		vkeys.ForEach(func(index, value gjson.Result) bool {
			if gjson.Get(data, "data.0.file."+ty[index.Int()][0]).Exists() && index.Int() > 0 && strings.Contains(value.Get("purl").String(), ".mp3?") {
				result.Url = gjson.Get(resp.Body, "req_0.data.sip.0").String() + value.Get("purl").String()
				result.Size = gjson.Get(data, "data.0.file."+ty[index.Int()][0]).Int()
			}
			return true
		})

		return common.ReturnData(200, result, "请求成功")
	} else {
		return common.ReturnData(400, nil, "请求失败")
	}
}

func FormatComments(data string, ctype string) common.Response {

	var exist bool
	var comments_str gjson.Result

	if ctype == "1" {
		comments_str = gjson.Get(data, "hot_comment.commentlist")
		exist = gjson.Get(data, "hot_comment.commentlist.0").Exists()
	} else {
		comments_str = gjson.Get(data, "comment.commentlist")
		exist = gjson.Get(data, "comment.commentlist.0").Exists()
	}

	if exist {
		var comments = []common.Comment{}

		comments_str.ForEach(func(index, value gjson.Result) bool {
			comments = append(comments, common.Comment{
				CommentId: value.Get("commentid").String(),
				UserId:    value.Get("encrypt_uin").String(),
				Nickname:  value.Get("nick").String(),
				AvatarUrl: value.Get("avatarurl").String(),
				Content:   value.Get("rootcommentcontent").String(),
				Time:      value.Get("time").String(),
				TimeStr:   "",
				Ip:        "",
				Location:  "",
			})
			return true
		})

		return common.ReturnData(200, comments, "请求成功")
	} else {
		return common.ReturnData(400, nil, "请求失败")
	}
}

func FormatList(data string) common.Response {

	list := gjson.Get(data, "data.cdlist.0")
	exist := list.Exists()

	if exist {

		var play_list = []common.Song{}

		user_id := list.Get("uin").String()
		nickname := list.Get("nickname").String()
		signature := ""
		avatar_url := list.Get("headurl").String()

		creator := common.Creator{
			UserId:    user_id,
			Nickname:  nickname,
			Signature: signature,
			AvatarUrl: avatar_url,
		}

		id := list.Get("disstid").String()
		name := list.Get("dissname").String()
		description := list.Get("desc").String()
		cover_url := strings.Replace(list.Get("logo").String(), "300?n=1", "", 1)
		play_list_str := list.Get("songlist")

		play_list_str.ForEach(func(index, value gjson.Result) bool {

			var artist = []string{}

			artist_str := value.Get("singer")
			artist_str.ForEach(func(index, item gjson.Result) bool {
				artist = append(artist, item.Get("name").String())
				return true
			})

			id := value.Get("mid").String()

			song := common.Song{
				Id:      id,
				Name:    value.Get("name").String(),
				Artist:  artist,
				Album:   value.Get("album.name").String(),
				PicUrl:  "https://y.gtimg.cn/music/photo_new/T002R800x800M000" + value.Get("album.mid").String() + ".jpg",
				UrlId:   id,
				LyricId: id,
				MvId:    value.Get("mv.vid").String(),
				Source:  "tencent",
			}

			play_list = append(play_list, song)
			return true
		})

		return common.ReturnData(200, common.List{
			Creator:     creator,
			Id:          id,
			Name:        name,
			Description: description,
			CoverUrl:    cover_url,
			SongList:    play_list,
		}, "请求成功")
	} else {
		return common.ReturnData(400, nil, "请求失败")
	}
}

func FormatMV(data string, id string) common.Response {
	list := gjson.Get(data, "getMvUrl.data."+id+".mp4")
	exist := list.Get("0").Exists()
	if exist {
		var url = []string{}
		var result = common.MvUrl{
			Id:  id,
			Url: "",
		}

		list.ForEach(func(index, value gjson.Result) bool {
			if value.Get("url.0").Exists() {
				url = append(url, value.Get("url.1").String()+value.Get("vkey").String()+"/"+value.Get("cn").String())
			}
			return true
		})

		result.Url = url

		return common.ReturnData(400, result, "请求成功")
	} else {
		return common.ReturnData(400, nil, "请求失败")
	}
}
