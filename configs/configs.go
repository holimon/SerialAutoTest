package configs

type ServerConfigType struct {
	SerialCom   string
	BaudRate    int
	ServerAddr  string
	AccessToken string
	Token       string
}

var ServerConfig = ServerConfigType{}

var ClinetListened map[string]string = make(map[string]string)

var LogPathConfig = "run.log"
