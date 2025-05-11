package Tikhub

import "github.com/imroc/req/v3"

type Tikhub struct {
	r              *req.Client
	Cursor         string
	Live_cursor    string
	Internal_ext   string
	ApiKey         string
	ua             string
	room_id        string
	user_unique_id string
}

func NewGithubClient(key string, ua string) *Tikhub {
	var tikhub = &Tikhub{}
	tikhub.ua = ua
	tikhub.ApiKey = key

	tikhub.r = req.C().SetCommonBearerAuthToken(key).SetBaseURL("https://api.tikhub.io")
	return tikhub
}
func (t Tikhub) Generate_wss_xb_signature() (string, error) {
	//api/v1/douyin/web/generate_wss_xb_signature
	get, err := t.r.R().Get("/api/v1/douyin/web/generate_wss_xb_signature")
	if err != nil {
		return "", err
	}
	x := &xbjson{}
	err = get.UnmarshalJson(x)
	if err != nil {
		return "", err
	}

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
func (t Tikhub) Fetch_live_im_fetch() error {
	///api/v1/douyin/web/fetch_live_im_fetch
	get, err := t.r.R().Get("api/v1/douyin/web/fetch_live_im_fetch")
	if err != nil {
		return err
	}
	l := &livejson{}
	err = get.UnmarshalJson(l)
	if err != nil {
		return err
	}
	t.room_id = l.Params.RoomID
	t.user_unique_id = l.Params.UserUniqueID
	t.Internal_ext = l.Data.InternalExt
	t.Live_cursor = l.Data.Extra.LiveCursor
	t.Cursor = l.Data.Extra.Cursor
	// 处理 l.Data
	return nil
}
