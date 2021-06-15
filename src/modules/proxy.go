package modules

import (
	"encoding/json"
	"github.com/asmcos/requests"
	"io"
	"io/ioutil"
	"net/http"
	"proxy/config"
	"proxy/kits"
)

func Proxy(w http.ResponseWriter, req *http.Request) {
	//获取配置信息
	JsonData := config.ApiJson{Success: false, Message: "", Data: "{}"}
	r := requests.Requests()
	//r.Debug = 1
	AgentData := config.AgentJson{}
	Url := ""
	switch req.RequestURI {
	case "/web/v1/cmdb/agent":
		Url = config.CmdbUrl + req.RequestURI
	case "/web/v1/monitor/agent":
		Url = config.MonitorUrl + req.RequestURI
	}
	if Url != "" {
		ReqData, _ := ioutil.ReadAll(req.Body)
		err := json.Unmarshal(ReqData, &AgentData)
		d, _ := json.Marshal(AgentData)
		resp, err := r.PostJson(Url, string(d))
		if err == nil {
			err = resp.Json(&JsonData)
		}
	}
	Data, _ := json.Marshal(JsonData)
	_, err := io.WriteString(w, string(Data))
	if err != nil {
		kits.Log("Proxy server: "+err.Error(), "error", "Proxy")
	} else {
		kits.Log(req.RemoteAddr+" "+req.Method+" "+req.RequestURI, "info", "Proxy")
	}
}

func ProxyServer() {
	http.HandleFunc("/xxx/cmdb/agent", Proxy)
	http.HandleFunc("/xxx/monitor/agent", Proxy)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		kits.Log("Proxy server: "+err.Error(), "error", "ProxyServer")
	}
}
