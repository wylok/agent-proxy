package config

const (
	FilePath   = "/xxx"
	PidPath    = FilePath + "/xxx"
	PidFile    = PidPath + "agent-proxy.pid"
	LogFile    = PidPath + "agent-proxy.log"
	MonitorUrl = "http://xxx:8002"
	CmdbUrl    = "http://xxx:8003"
)

type ApiJson struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type AgentJson struct {
	AgentId string `json:"agent_id"`
	Version string `json:"version"`
	Data    string `json:"data"`
}
