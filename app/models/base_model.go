package models

type Response struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data,omitempty"`
	Error      *ErrorLog   `json:"error,omitempty"`
	Page       int         `json:"page,omitempty"`
	PerPage    int         `json:"per_page,omitempty"`
	Total      int64       `json:"total,omitempty"`
	StatCode   string      `json:"stat_code,omitempty"`
	StatMsg    string      `json:"stat_msg,omitempty"`
}

type ErrorLog struct {
	Line          string `json:"line,omitempty"`
	Filename      string `json:"filename,omitempty"`
	Function      string `json:"function,omitempty"`
	Message       string `json:"message,omitempty"`
	SystemMessage string `json:"system_message,omitempty"`
	Error         error  `json:"-"`
	StatusCode    int    `json:"-"`
}
