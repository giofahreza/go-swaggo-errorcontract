package models

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type SuccessResponse struct {
	Data interface{} `json:"data"`
}

type ErrorContract struct {
	Code    int    `json:"code"`
	UserMsg string `json:"user_msg"`
	SysMsg  string `json:"sys_msg"`
	Time    string `json:"time"`
	DocsURL string `json:"docs_url"`
}
