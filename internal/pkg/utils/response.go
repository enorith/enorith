package utils

type RedirectResp struct {
	To   string `json:"to"`
	Code int    `json:"code"`
}

func RedirectTo(to string) RedirectResp {
	return RedirectResp{To: to, Code: 30201}
}
