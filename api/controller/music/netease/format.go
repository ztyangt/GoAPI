package netease

import (
	"GoAPI/common"
	"GoAPI/helper"

	"github.com/tidwall/gjson"
)

func df_song() string {
	return helper.Json.Encode(common.Song{
		Id:      "",
		Name:    "",
		Artist:  nil,
		Album:   "",
		PicUrl:  "",
		UrlId:   "",
		LyricId: "",
		MvId:    "0",
		Source:  "netease",
	})
}

func FormatList(data string) string {

	list := gjson.Get(data, "playlist.tracks")
	exist := list.Exists()

	if exist {

		var play_list = []common.Song{}

		user_id := gjson.Get(data, "playlist.creator.userId").String()
		nickname := gjson.Get(data, "playlist.creator.nickname").String()
		signature := gjson.Get(data, "playlist.creator.signature").String()
		avatar_url := gjson.Get(data, "playlist.creator.avatarUrl").String()

		creator := common.Creator{
			UserId:    user_id,
			Nickname:  nickname,
			Signature: signature,
			AvatarUrl: avatar_url,
		}

		id := gjson.Get(data, "playlist.id").String()
		name := gjson.Get(data, "playlist.name").String()
		description := gjson.Get(data, "playlist.description").String()
		cover_url := gjson.Get(data, "playlist.coverImgUrl").String()
		play_list_str := gjson.Get(data, "playlist.tracks")

		play_list_str.ForEach(func(index, value gjson.Result) bool {

			var artist = []string{}

			artist_str := value.Get("ar")
			artist_str.ForEach(func(index, item gjson.Result) bool {
				artist = append(artist, item.Get("name").String())
				return true
			})

			id := value.Get("id").String()

			song := common.Song{
				Id:      id,
				Name:    value.Get("name").String(),
				Artist:  artist,
				Album:   value.Get("al.name").String(),
				PicUrl:  value.Get("al.picUrl").String(),
				UrlId:   id,
				LyricId: id,
				MvId:    value.Get("mv").String(),
				Source:  "netease",
			}

			play_list = append(play_list, song)
			return true
		})

		return helper.Json.Encode(common.List{
			Creator:     creator,
			Id:          id,
			Name:        name,
			Description: description,
			CoverUrl:    cover_url,
			PlayList:    play_list,
		})
	} else {
		return "{}"
	}
}

func FormatSong(data string) string {

	exist := gjson.Get(data, "songs").Exists()

	if exist {
		var artist = []string{}
		id := gjson.Get(data, "songs.0.id").String()
		name := gjson.Get(data, "songs.0.name").String()

		artist_str := gjson.Get(data, "songs.0.ar")
		artist_str.ForEach(func(index, value gjson.Result) bool {
			artist = append(artist, gjson.Get(value.String(), "name").String())
			return true
		})

		album := gjson.Get(data, "songs.0.al.name").String()
		pic_str := gjson.Get(data, "songs.0.al.picUrl").String()
		mv_id := gjson.Get(data, "songs.0.mv").String()

		return helper.Json.Encode(common.Song{
			Id:      id,
			Name:    name,
			Artist:  artist,
			Album:   album,
			PicUrl:  pic_str,
			UrlId:   id,
			LyricId: id,
			MvId:    mv_id,
			Source:  "netease",
		})
	} else {
		return "{}"
	}
}

func FormatUrl(data string) string {

	exist := gjson.Get(data, "data.0").Exists()

	if exist {
		id := gjson.Get(data, "data.0.id").String()
		size := gjson.Get(data, "data.0.size").Int()
		url := gjson.Get(data, "data.0.url").String()
		duration := gjson.Get(data, "data.0.time").Int()

		return helper.Json.Encode(common.Url{
			Id:       id,
			Size:     size,
			Url:      url,
			Duration: duration,
		})

	} else {
		return "{}"
	}
}

func FormatMV(data string, id string) string {
	exist := gjson.Get(data, "data.id").Exists()
	if exist {
		var url = []string{}
		var result = common.MvUrl{
			Id:  id,
			Url: "",
		}

		url_list := gjson.Get(data, "data.brs")
		url_list.ForEach(func(index, value gjson.Result) bool {
			url = append(url, value.String())
			return true
		})

		result.Url = url

		return helper.Json.Encode(result)
	} else {
		return "{}"
	}
}

func FormatComments(data string, ctype string) string {

	var exist bool
	var comments_str gjson.Result

	if ctype == "1" {
		comments_str = gjson.Get(data, "hotComments")
		exist = gjson.Get(data, "hotComments.0").Exists()
	} else {
		comments_str = gjson.Get(data, "comments")
		exist = gjson.Get(data, "comments.0").Exists()
	}

	if exist {
		var comments = []common.Comment{}

		comments_str.ForEach(func(index, value gjson.Result) bool {
			comments = append(comments, common.Comment{
				CommentId: value.Get("commentId").String(),
				UserId:    value.Get("user.userId").String(),
				Nickname:  value.Get("user.nickname").String(),
				AvatarUrl: value.Get("user.avatarUrl").String(),
				Content:   value.Get("content").String(),
				Time:      value.Get("time").String(),
				TimeStr:   value.Get("timeStr").String(),
				Ip:        value.Get("ipLocation.ip").String(),
				Location:  value.Get("ipLocation.location").String(),
			})
			return true
		})

		return helper.Json.Encode(comments)
	} else {
		return "[]"
	}
}

func FormatLyric(data string) string {
	exist := gjson.Get(data, "lrc").Exists()
	if exist {
		return data
	} else {
		return "{}"
	}
}

func FormatUser(data string) string {
	return data
}

func FormatSearch(data string, type_ string) string {

	exist := gjson.Get(data, "result").Exists()
	if exist {

		switch type_ {
		case "10":
			return gjson.Get(data, "result.albums").String()
		case "100":
			return gjson.Get(data, "result.artists").String()
		case "1000":
			return gjson.Get(data, "result.playlists").String()
		case "1004":
			return gjson.Get(data, "result.mvs").String()
		case "1006":
			return gjson.Get(data, "result.songs").String()
		case "1009":
			return gjson.Get(data, "result.djRadios").String()
		default:
			var search_result = []common.Song{}
			search_list_str := gjson.Get(data, "result.songs")
			search_list_str.ForEach(func(index, value gjson.Result) bool {

				var artist = []string{}

				artist_str := value.Get("ar")
				artist_str.ForEach(func(index, item gjson.Result) bool {
					artist = append(artist, item.Get("name").String())
					return true
				})

				id := value.Get("id").String()

				song := common.Song{
					Id:      id,
					Name:    value.Get("name").String(),
					Artist:  artist,
					Album:   value.Get("al.name").String(),
					PicUrl:  value.Get("al.PicUrl").String(),
					UrlId:   id,
					LyricId: id,
					MvId:    value.Get("mv").String(),
					Source:  "netease",
				}
				search_result = append(search_result, song)
				return true
			})
			return helper.Json.Encode(search_result)
		}

	} else {
		return "[]"
	}
}
