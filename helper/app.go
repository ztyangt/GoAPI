package helper

func init() {
	Is.Ip = isIp
	Is.Url = isUrl
	Is.Email = isEmail
	Is.Phone = isPhone
	Is.Empty = isEmpty
	Is.True = isTrue
	Is.False = isFalse
	Is.Domain = isDomain
	Get.Type = getType
	In.Array = inArray
	Array.Filter = arrayFilter
	Json.Encode = JsonEncode
	Json.Decode = JsonDecode
	HTTP.GET = GET
	HTTP.POST = POST
	Net.Tcping = NetTcping
	Struct.Set = StructSet
	Struct.Get = StructGet
	Struct.Del = StructDel
	Struct.Has = StructHas
	Struct.Keys = StructKeys
	Struct.Values = StructValues
	Struct.Len = StructLen
	Struct.Map = StructMap
	Struct.Slice = StructSlice
}

var Get struct {
	Type func(value any) string
}

var In struct {
	Array func(value any, array []any) bool
}

var Is struct {
	Ip     func(ip any) bool
	Url    func(url any) bool
	Email  func(email any) bool
	Phone  func(phone any) bool
	Empty  func(value any) bool
	True   func(value any) bool
	False  func(value any) bool
	Domain func(domain any) bool
}

var Array struct {
	Filter func(array []string) []string
}

var Json struct {
	Encode func(value any) (result string)
	Decode func(value string) (result any)
}

var HTTP struct {
	GET  func(url string, params map[string]any, headers map[string]any) (result CurlResult)
	POST func(url string, params map[string]any, headers map[string]any) (result CurlResult)
}

var Net struct {
	Tcping func(host any, opts ...map[string]any) (ok bool, detail []map[string]any)
}

var Struct struct {
	Set    func(obj any, key string, val any)
	Get    func(obj any, key string) (result any)
	Del    func(obj any, key string)
	Has    func(obj any, key string) (ok bool)
	Keys   func(obj any) (slice []string)
	Values func(obj any) (slice []any)
	Len    func(obj any) (length int)
	Map    func(obj any) (result map[string]any)
	Slice  func(obj any) (slice []any)
}
