package Tikhub

import "github.com/imroc/req/v3"

type xbjson struct {
	Code   int    `json:"code"`
	Router string `json:"router"`
	Params struct {
		UserAgent    string `json:"user_agent"`
		RoomID       string `json:"room_id"`
		UserUniqueID string `json:"user_unique_id"`
	} `json:"params"`
	Data struct {
		XBogus string `json:"x_bogus"`
	} `json:"data"`
}
type Tikhub struct {
	*req.Client
	ApiKey         string
	ua             string
	room_id        string
	user_unique_id string
}

type ttwid struct {
	Code   int    `json:"code"`
	Router string `json:"router"`
	Params struct {
	} `json:"params"`
	Data struct {
		Ttwid string `json:"ttwid"`
	} `json:"data"`
}
type FetchJson struct {
	Code   int    `json:"code"`
	Router string `json:"router"`
	Params struct {
	} `json:"params"`
	Data struct {
		ID                 string `json:"id"`
		CreateTime         string `json:"create_time"`
		LastTime           string `json:"last_time"`
		UserUID            string `json:"user_uid"`
		UserUIDType        int    `json:"user_uid_type"`
		FirebaseInstanceID string `json:"firebase_instance_id"`
		UserAgent          string `json:"user_agent"`
		BrowserName        string `json:"browser_name"`
	} `json:"data"`
}
type livejson struct {
	Code   int    `json:"code"`
	Router string `json:"router"`
	Params struct {
		RoomID       string `json:"room_id"`
		UserUniqueID string `json:"user_unique_id"`
	} `json:"params"`
	Data struct {
		Data []struct {
			Common struct {
				Method     string `json:"method"`
				PlayTime   int    `json:"play_time"`
				IsShowMsg  bool   `json:"is_show_msg"`
				RoomID     int64  `json:"room_id"`
				CreateTime int64  `json:"create_time"`
				MsgID      int64  `json:"msg_id"`
			} `json:"common"`
			Content string `json:"content"`
		} `json:"data"`
		Extra struct {
			Cursor        string `json:"cursor"`
			LiveCursor    string `json:"live_cursor"`
			FetchInterval int    `json:"fetch_interval"`
			Now           int64  `json:"now"`
		} `json:"extra"`
		InternalExt string `json:"internal_ext"`
		StatusCode  int    `json:"status_code"`
	} `json:"data"`
}

func NewGithubClient(key string) *Tikhub {
	var tikhub = &Tikhub{}
	tikhub.ApiKey = key
	tikhub.Client = req.C().SetCommonBearerAuthToken(key).SetBaseURL("https://api.tikhub.io")
	return tikhub
}
func (t Tikhub) Generate_wss_xb_signature() (string, error) {
	//api/v1/douyin/web/generate_wss_xb_signature
	get, err := t.R().Get("/api/v1/douyin/web/generate_wss_xb_signature")
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
	post, err := t.R().SetBodyString(ttwid).Post("/api/v1/douyin/web/fetch_query_user")
	if err != nil {
		return "", err
	}
	f := &FetchJson{}
	err = post.UnmarshalJson(f)
	if err != nil {
		return "", err
	}
	t.ua = f.Data.UserAgent

	return f.Data.UserUID, nil
}
func (t *Tikhub) Generate_ttwid() (string, error) {
	get, err := t.R().Get("/api/v1/douyin/web/generate_ttwid")
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
	get, err := t.R().Get("api/v1/douyin/web/fetch_live_im_fetch")
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
	// 处理 l.Data
	return nil
}
