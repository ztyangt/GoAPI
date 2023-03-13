package common

type Creator struct {
	UserId    string `json:"user_id"`
	Nickname  string `json:"nickname"`
	Signature string `json:"signature"`
	AvatarUrl string `json:"avatar_url"`
}

type List struct {
	Creator     Creator `json:"creator"`
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	CoverUrl    string  `json:"cover_url"`
	SongList    any     `json:"song_list"`
}

type Song struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Artist  any    `json:"artist"`
	Album   string `json:"album"`
	PicUrl  string `json:"pic_url"`
	UrlId   string `json:"url_id"`
	LyricId string `json:"lyric_id"`
	MvId    string `json:"mv_id"`
	Source  string `json:"source"`
}

type Url struct {
	Id       string `json:"id"`
	Size     int64  `json:"size"`
	Url      string `json:"url"`
	Duration int64  `json:"duration"`
}

type MvUrl struct {
	Id  string `json:"id"`
	Url any    `json:"url"`
}

type Comment struct {
	CommentId string `json:"comment_id"`
	UserId    string `json:"user_id"`
	Nickname  string `json:"nickname"`
	AvatarUrl string `json:"avatar_url"`
	Content   string `json:"content"`
	Time      string `json:"time"`
	TimeStr   string `json:"time_str"`
	Ip        string `json:"ip"`
	Location  string `json:"location"`
}

type DouyinVideo struct {
	Duration    int    `json:"duration"`
	Cover       string `json:"cover"`
	PlayUrl     string `json:"play_url"`
	Description string `json:"description"`
}

type DouyinMusic struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	PlayUrl string `json:"play_url"`
}

type DouyinInfo struct {
	Video  any         `json:"video"`
	Images any         `json:"images"`
	Music  DouyinMusic `json:"music"`
	Author Creator     `json:"author"`
}
