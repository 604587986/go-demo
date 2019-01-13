package api

type ApiRes struct {
	Code uint        `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Page struct{
	PageSize int `json:"page_size"`
	PageNum	int `json:"page_num"`
}
