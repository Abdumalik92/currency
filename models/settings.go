package models

// Settings app settings
type Settings struct {
	AppParams      Params           `json:"app"`
	PostgresParams PostgresSettings `json:"postgresParams"`
	SecretKey      SecretKey        `json:"secretKey"`
}

// Params contains params of server meta data
type Params struct {
	ServerName string `json:"serverName"`
	PortRun    string `json:"portRun"`
	LogFile    string `json:"logFile"`
	ServerURL  string `json:"serverURL"`
}

// MySQLParams conteins params of mySql db server
type PostgresSettings struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Server   string `json:"server"`
	Port     int    `json:"port"`
	DataBase string `json:"database"`
}
type SecretKey struct {
	Key         string `json:"key"`
	PasswordSMS string `json:"passwordSMS"`
}
