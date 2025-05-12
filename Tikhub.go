package Tikhub

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/imroc/req/v3"
)

type Tikhub struct {
	r              *req.Client
	Cursor         string
	Live_cursor    string
	Internal_ext   string
	ApiKey         string
	ua             string
	room_id        string
	user_unique_id string
	XBogus         string
	browser_name   string
}

// RandUserAgent 随机UA
func RandUserAgent() string {
	ua := gofakeit.UserAgent()
	return ua
	//fmt.Println("随机UA:", ua)
}
func NewGithubClient(key string, ua string) *Tikhub {
	var tikhub = &Tikhub{}
	tikhub.ua = ua
	tikhub.ApiKey = key

	tikhub.r = req.C().SetCommonBearerAuthToken(key).SetBaseURL("https://api.tikhub.io")
	return tikhub
}

// GenerateWsLink 生成ws链接
func GenerateWsLink(key, userAgent, webcastId string) (WsLink, error) {
	var wslink WsLink
	t := NewGithubClient(key, userAgent)
	ttwid, err := t.GenerateTtwid(userAgent)
	wslink.Ttwid = ttwid.Data.Ttwid
	if err != nil {
		return WsLink{}, err
	}
	roomId, err := t.WebcastId2RoomId(webcastId)
	if err != nil {
		return WsLink{}, err
	}
	User, err := t.FetchQueryUser(ttwid.Data.Ttwid)

	if err != nil {
		return WsLink{}, err
	}
	fetch, err := t.FetchLiveImFetch(roomId.Data.RoomID, User.Data.UserUID)
	if err != nil {
		return WsLink{}, err
	}
	signature, err := t.GenerateWssXbSignature(userAgent, roomId.Data.RoomID, User.Data.UserUID)
	if err != nil {
		return WsLink{}, err
	}
	//browserInfo := strings.Split(User.Data.UserAgent, "Mozilla")[1]
	//parsedURL := strings.Replace(browserInfo[1:], " ", "%20", -1)
	wslink.Url = fmt.Sprintf("wss://webcast5-ws-web-hl.douyin.com/webcast/im/push/v2/?aid=6383&app_name=douyin_web&browser_language=zh-CN&browser_name=%s&browser_online=true&browser_platform=Win32&browser_version=%s&compress=gzip&cookie_enabled=true&device_platform=web&did_rule=3&endpoint=live_pc&heartbeatDuration=0&host=https://live.douyin.com&identity=audience&im_path=/webcast/im/fetch/&insert_task_id=&live_id=1&live_reason=&need_persist_msg_count=15&screen_height=1080&screen_width=1920&support_wrds=1&tz_name=Asia/Shanghai&update_version_code=1.0.14-beta.0&version_code=180800&webcast_sdk_version=1.0.14-beta.0&room_id=%s&user_unique_id=%s&cursor=%s&internal_ext=%s&signature=%s",
		User.Data.BrowserName, userAgent,
		roomId.Data.RoomID, User.Data.UserUID, fetch.Data.Extra.Cursor, fetch.Data.InternalExt, signature.Data.XBogus)
	return wslink, err
}

// GenerateWssXbSignature 生成弹幕xb签名
func (t *Tikhub) GenerateWssXbSignature(userAgent, roomId, userUniqueId string) (Xbjson, error) {
	//api/v1/douyin/web/generate_wss_xb_signature
	get, err := t.r.R().SetQueryParam("user_agent", userAgent).
		SetQueryParam("room_id", roomId).
		SetQueryParam("user_unique_id", userUniqueId).
		Get("/api/v1/douyin/web/generate_wss_xb_signature")
	if err != nil {
		return Xbjson{}, fmt.Errorf("GenerateWssXbSignature请求失败: %v", err)
	}
	x := Xbjson{}
	//println(get.String())
	err = get.UnmarshalJson(&x)
	if err != nil {
		return Xbjson{}, fmt.Errorf("GenerateWssXbSignature解析失败: %v", err)
	}
	t.XBogus = x.Data.XBogus
	return x, nil

}

// FetchQueryUser 查询抖音用户基本信息
func (t *Tikhub) FetchQueryUser(ttwid string) (FetchJson, error) {
	///api/v1/douyin/web/fetch_query_user
	tw := fmt.Sprintf(`ttwid=%s;`, ttwid)

	post, err := t.r.R().SetBody(tw).Post("/api/v1/douyin/web/fetch_query_user")
	if err != nil {
		return FetchJson{}, fmt.Errorf("FetchQueryUser请求失败: %v", err)
	}
	f := FetchJson{}
	//println(tw, string(post.Request.Body))
	err = post.UnmarshalJson(&f)
	if err != nil {
		return FetchJson{}, fmt.Errorf("FetchQueryUser解析失败: %v", err)
	}
	if f.Code != 200 {
		return FetchJson{}, fmt.Errorf("FetchQueryUser/返回失败: %v", post.String())
	}
	t.ua = f.Data.UserAgent
	t.browser_name = f.Data.BrowserName
	t.user_unique_id = f.Data.UserUID
	return f, nil
}

// GenerateTtwid 生成ttwid
func (t *Tikhub) GenerateTtwid(userAgent string) (Ttwid, error) {
	get, err := t.r.R().SetQueryParam("user_agent", userAgent).Get("/api/v1/douyin/web/generate_ttwid")
	if err != nil {
		return Ttwid{}, fmt.Errorf("GenerateTtwid请求失败: %v", err)
	}
	// 生成一个随机的 ttwid
	//println(get.String())
	ttwid := Ttwid{}
	err = get.UnmarshalJson(&ttwid)
	if err != nil {
		return Ttwid{}, fmt.Errorf("GenerateTtwid解析失败: %v", err)
	}
	return ttwid, nil
}

// FetchLiveImFetch 抖音直播间弹幕参数获取
func (t *Tikhub) FetchLiveImFetch(roomId, userUniqueId string) (Livejson, error) {
	///api/v1/douyin/web/fetch_live_im_fetch
	get, err := t.r.R().SetQueryParam("room_id", roomId).
		SetQueryParam("user_unique_id", userUniqueId).
		Get("api/v1/douyin/web/fetch_live_im_fetch")
	if err != nil {
		return Livejson{}, fmt.Errorf("FetchLiveImFetch请求失败: %v", err)
	}
	l := Livejson{}
	//println(roomId, userUniqueId, get.String())

	err = get.UnmarshalJson(&l)
	if err != nil {
		return Livejson{}, fmt.Errorf("FetchLiveImFetch解析失败: %v", err)
	}
	return l, nil
}

// WebcastId2RoomId 直播间号转房间号
func (t *Tikhub) WebcastId2RoomId(room string) (Roomidjson, error) {
	get, err := t.r.R().SetQueryParam("webcast_id", room).Get("/api/v1/douyin/web/webcast_id_2_room_id")
	if err != nil {
		return Roomidjson{}, fmt.Errorf("WebcastId2RoomId请求失败: %v", err)
	}
	json := Roomidjson{}
	err = get.UnmarshalJson(&json)
	if err != nil {
		return Roomidjson{}, fmt.Errorf("WebcastId2RoomId解析失败: %v", err)
	}
	t.room_id = json.Data.RoomID
	return json, nil
}
func (t *Tikhub) SprintUrl() string {
	roomid, err := t.WebcastId2RoomId(t.room_id)
	if err != nil {
		return ""
	}
	sprint := fmt.Sprintf("wss://webcast5-ws-web-lf.douyin.com/webcast/im/push/v2/?aid=6383&app_name="+
		"douyin_web&browser_language=zh-CN&browser_name=%s&browser_online=true&browser_platform=Win32"+
		"&browser_version=%s&compress=gzip&cookie_enabled=true&device_platform=web&did_rule=3"+
		"&endpoint=live_pc&heartbeatDuration=0&host=https://live.douyin.com&identity=audience"+
		"&im_path=/webcast/im/fetch/&insert_task_id=&live_id=1&live_reason="+
		"&need_persist_msg_count=15&screen_height=1080&screen_width=1920"+
		"&support_wrds=1&tz_name=Asia/Shanghai&update_version_code=1.0.14-beta.0"+
		"&version_code=180800&webcast_sdk_version=1.0.14-beta.0&room_id=%s&user_unique_id=%s"+
		"&cursor=%s&internal_ext=%s&signature=%s", t.browser_name, t.ua, roomid.Data.RoomID, t.user_unique_id, t.Cursor, t.Internal_ext, t.XBogus)
	return sprint
}
