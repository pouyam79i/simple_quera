package config

const (
	API_CODEX_URL = "https://api.codex.jaagrav.in"
)

type ApiCall struct {
	Email string   `json:"email"`
	Code  CodexAPI `json:"code"`
}

type CodexAPI struct {
	Code     string `json:"code"`
	Language string `json:"language"`
	Input    string `json:"input"`
}

type ResCodeX struct {
	TimeStamp int    `json:"timeStamp"`
	Status    int    `json:"status"`
	Output    string `json:"output"`
	Error     string `json:"error"`
	Language  string `json:"language"`
	Info      string `json:"info"`
}

type Response struct {
	Result bool   `json:"result"`
	Info   string `json:"info"`
}
