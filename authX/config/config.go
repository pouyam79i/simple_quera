package config

// config data structure
type ServerInfo struct {
	IP   string `yaml:"ip"`
	Port string `yaml:"port"`
}

type DataBaseInfo struct {
	IP       string `yaml:"ip"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Table    string `yaml:"table"`
}

type API_CONNECTIONS struct {
	MainKey string `yaml:"mainKey"`
}

type Configs struct {
	Server         ServerInfo      `yaml:"server"`
	DBI            DataBaseInfo    `yaml:"db"`
	ApiConnections API_CONNECTIONS `yaml:"apiConnections"`
}

// API communication data structure

type API_SIGNIN_REQUEST struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type API_SIGNIN_RESULT struct {
	Result bool   `json:"result"`
	Token  string `json:"token"`
	Info   string `json:"info"`
}

type API_IDENTIFY_REQUEST struct {
	Token string `json:"token"`
}

type API_IDENTIFY_RESULT struct {
	Result bool   `json:"result"`
	Info   string `json:"info"`
}

// DB data retrieval
type USER_DATA struct {
	EMAIL string `json:"EMAIL"`
	TOKEN string `json:"TOKEN"`
	PWD   string `json:"PWD"`
}
