package Tikhub

type Xbjson struct {
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

type Ttwid struct {
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
type Livejson struct {
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
type Roomidjson struct {
	Code   int    `json:"code"`
	Router string `json:"router"`
	Params struct {
		WebcastID string `json:"webcast_id"`
	} `json:"params"`
	Data struct {
		RoomID string `json:"room_id"`
	} `json:"data"`
}
