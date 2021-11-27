package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//IAPIHandlers interface definition
type IAPIHandlers interface {
	SetupHandlers()
}

// APIHandlers public API implementation
type APIHandlers struct {
	core *InstanceCore
}

// NewAPIHandlers create APIHandlers instance
func NewAPIHandlers(core *InstanceCore) IAPIHandlers {
	return &APIHandlers{
		core: core,
	}
}

func readBody(res http.ResponseWriter, req *http.Request) (*[]byte, error) {
	res.Header().Add("Access-Control-Allow-Origin", "*")
	res.Header().Add("Access-Control-Allow-Credentials", "true")
	res.Header().Add("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
	res.Header().Add("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, "+
		"Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, "+
		"Access-Control-Request-Headers")

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprint(res, BuildResponse(ErrorReadRequestBody, ErrorReadRequestBodyMsg))
		return nil, err
	}
	if len(body) == 0 {
		return nil, fmt.Errorf("preflight request detected")
	}
	return &body, nil
}

// SetupHandlers setup API handlers instance
func (api *APIHandlers) SetupHandlers() {
	http.HandleFunc("/v1/GetLunarArrivalTime", api.lunarArrivalTimeHandler)
}

func (api *APIHandlers) lunarArrivalTimeHandler(res http.ResponseWriter, req *http.Request) {
	body, err := readBody(res, req)
	if err != nil {
		log.Error(err)
		return
	}

	lunarArrivalTime, requestError := api.core.GetArrivalTimeManager().GetLunarArrivalTime(string(*body))
	if requestError != nil {
		log.Error(requestError)
		res.Write([]byte(BuildResponse(requestError.Code, requestError.Error.Error())))
		return
	}

	respBody, _ := json.Marshal(lunarArrivalTime)
	res.Write(respBody)
}
