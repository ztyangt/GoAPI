package service

import (
	"GinAPI/helper"
	"fmt"

	"github.com/tidwall/gjson"
)

type Creator struct {
	Id string `json:"id"`
	Nickname string `json:"nickname"`
	Signing string `json:"signing"`
	AvatarUrl   string `json:"avatarUrl"`

}

type List struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Artist  any    `json:"artist"`
	Description   string `json:"description"`
	PicUrl   string `json:"pic_url"`
	UrlId   string `json:"url_id"`
	LyricId string `json:"lyric_id"`
}

type Song struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Artist  any    `json:"artist"`
	Album   string `json:"album"`
	PicUrl   string `json:"pic_url"`
	UrlId   string `json:"url_id"`
	LyricId string `json:"lyric_id"`
}


type Url struct {
	Id    string `json:"id"`
	Size  int64    `json:"size"`
	Url  	string    `json:"url"`
	Duration   int64 `json:"duration"`
}

func df_song() string {
	return helper.Json.Encode(Song{
		Id:      "",
		Name:    "",
		Artist:  nil,
		Album:   "",
		PicUrl:   "",
		UrlId:   "",
		LyricId: "",
	})
}

func FormatList(data string) string {

	list := gjson.Get(data, "playlist.tracks")
	exist := list.Exists()
	
	if exist {

		fmt.Println(list)

		return "{}"

		// var artist = []string{}
		// id := gjson.Get(data, "songs.0.id").String()
		// name := gjson.Get(data, "songs.0.name").String()

		// artist_str := gjson.Get(data, "songs.0.ar")
		// artist_str.ForEach(func(index, value gjson.Result) bool {
		// 	artist = append(artist,gjson.Get(value.String(), "name").String())
		// 	return true 
		// })

		// album := gjson.Get(data, "songs.0.al.name").String()
		// pic_str := gjson.Get(data, "songs.0.al.picUrl").String()

		// return helper.Json.Encode(Song{
		// 	Id:      id,
    //   Name:    name,
    //   Artist:  artist,
    //   Album:   album,
    //   PicUrl:   pic_str,
    //   UrlId:   id,
    //   LyricId: id,
		// })
	}else{
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
			artist = append(artist,gjson.Get(value.String(), "name").String())
			return true 
		})

		album := gjson.Get(data, "songs.0.al.name").String()
		pic_str := gjson.Get(data, "songs.0.al.picUrl").String()

		return helper.Json.Encode(Song{
			Id:      id,
      Name:    name,
      Artist:  artist,
      Album:   album,
      PicUrl:   pic_str,
      UrlId:   id,
      LyricId: id,
		})
	}else{
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

		return helper.Json.Encode(Url{
			Id:      id,
      Size:    size,
      Url:  		url,
      Duration:   duration,
		})

	}else{
		return "{}"
	}
}