package config

const (
	API_CODEX_URL = "https://api.codex.jaagrav.in"

	API_JB_RMQ_URL = "localhost:5672"

	API_AUTHX_SIGNIN   = "http://localhost:8090/signin?key="
	API_AUTHX_VALIDATE = "http://localhost:8090/validate?key="
	AUTHX_KEY          = "202249e7150520e8acc59f818754c4d3ac4f166b1494f043dea0ece6125285b4" // change if needed try to generate it on online websites
)

type ServerInfo struct {
	Name  string `json:"name"`
	IP    string `json:"ip"`
	Port  string `json:"port"`
	Debug bool   `json:"debug"`
}

type SingInInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ClientMSG struct {
	Result bool   `json:"result"`
	Token  string `json:"token"`
	Info   string `json:"info"`
}

type JustToken struct {
	Token string `json:"token"`
}

type ValidatorResult struct {
	Result bool   `json:"result"`
	Info   string `json:"info"`
}

type CodexAPI struct {
	Code     string `json:"code"`
	Language string `json:"language"`
	Input    string `json:"input"`
}

type ClientCode struct {
	Token string   `json:"token"`
	CodeX CodexAPI `json:"codex"`
}

type ResCodeX struct {
	TimeStamp int    `json:"timeStamp"`
	Status    int    `json:"status"`
	Output    string `json:"output"`
	Error     string `json:"error"`
	Language  string `json:"language"`
	Info      string `json:"info"`
}

// send message to job builder via rabbitMQ
// use below pattern
type JB_MSG struct {
	Data  string `json:"data"`
	Token string `json:"token"`
}
