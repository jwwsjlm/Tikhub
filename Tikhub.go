package Tikhub

import (
	"fmt"
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
}

func NewGithubClient(key string, ua string) *Tikhub {
	var tikhub = &Tikhub{}
	tikhub.ua = ua
	tikhub.ApiKey = key

	tikhub.r = req.C().SetCommonBearerAuthToken(key).SetBaseURL("https://api.tikhub.io")
	return tikhub
}
func (t *Tikhub) Generate_wss_xb_signature() (string, error) {
	//api/v1/douyin/web/generate_wss_xb_signature
	get, err := t.r.R().SetQueryParam("user_agent", t.ua).SetQueryParam("room_id", t.room_id).SetQueryParam("user_unique_id", t.user_unique_id).Get("/api/v1/douyin/web/generate_wss_xb_signature")
	if err != nil {
		return "", err
	}
	x := &xbjson{}
	//println(get.String())
	err = get.UnmarshalJson(x)
	if err != nil {
		return "", err
	}
	t.XBogus = x.Data.XBogus
	return x.Data.XBogus, nil

}
func (t *Tikhub) Fetch_query_user(ttwid string) (string, error) {
	///api/v1/douyin/web/fetch_query_user
	post, err := t.r.R().SetBodyString(ttwid).Post("/api/v1/douyin/web/fetch_query_user")
	if err != nil {
		return "", err
	}
	f := &fetchJson{}
	err = post.UnmarshalJson(f)
	if err != nil {
		return "", err
	}
	t.ua = f.Data.UserAgent
	t.user_unique_id = f.Data.UserUID
	return f.Data.UserUID, nil
}
func (t *Tikhub) Generate_ttwid() (string, error) {
	get, err := t.r.R().SetQueryParam("user_agent", t.ua).Get("/api/v1/douyin/web/generate_ttwid")
	if err != nil {
		return "", err
	}
	// 生成一个随机的 ttwid
	var ttwid ttwid
	err = get.UnmarshalJson(&ttwid)
	if err != nil {
		return "", err
	}
	return "ttwid=" + ttwid.Data.Ttwid, nil
}
func (t *Tikhub) Fetch_live_im_fetch(room_id string) error {
	///api/v1/douyin/web/fetch_live_im_fetch
	get, err := t.r.R().SetQueryParam("room_id", room_id).SetQueryParam("user_unique_id", t.user_unique_id).Get("api/v1/douyin/web/fetch_live_im_fetch")
	if err != nil {
		return err
	}
	l := &livejson{}
	//println(get.String())
	err = get.UnmarshalJson(l)
	if err != nil {
		return err
	}
	//println(l.Data.InternalExt)
	t.room_id = l.Params.RoomID
	t.user_unique_id = l.Params.UserUniqueID
	t.Internal_ext = l.Data.InternalExt
	t.Live_cursor = l.Data.Extra.LiveCursor
	t.Cursor = l.Data.Extra.Cursor
	// 处理 l.Data
	return nil
}
func (t *Tikhub) SprintUrl() string {
	sprint := fmt.Sprintf("wss://webcast5-ws-web-lf.douyin.com/webcast/im/push/v2/?aid=6383&app_name="+
		"douyin_web&browser_language=zh-CN&browser_name=Mozilla&browser_online=true&browser_platform=Win32"+
		"&browser_version=%s&compress=gzip&cookie_enabled=true&device_platform=web&did_rule=3"+
		"&endpoint=live_pc&heartbeatDuration=0&host=https://live.douyin.com&identity=audience"+
		"&im_path=/webcast/im/fetch/&insert_task_id=&live_id=1&live_reason="+
		"&need_persist_msg_count=15&screen_height=1080&screen_width=1920"+
		"&support_wrds=1&tz_name=Asia/Shanghai&update_version_code=1.0.14-beta.0"+
		"&version_code=180800&webcast_sdk_version=1.0.14-beta.0&room_id=%s&user_unique_id=%s"+
		"&cursor=%s&internal_ext=%s&signature=%s", t.ua, t.room_id, t.user_unique_id, t.Cursor, t.Internal_ext, t.XBogus)
	return sprint
}
