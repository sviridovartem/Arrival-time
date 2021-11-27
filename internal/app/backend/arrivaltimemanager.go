package backend

import (
	"encoding/json"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

// IArrivalTimeManager interface of ArrivalTimeManager
type IArrivalTimeManager interface {
	GetLunarArrivalTime(string) (*LunarArrivalTimeResponse, *RequestError)
}

// ArrivalTimeManager arrival time manager instance
type ArrivalTimeManager struct {
	core *InstanceCore
}

// NewArrivalTimeManager create a new arrival time manager  instance
func NewArrivalTimeManager(core *InstanceCore) IArrivalTimeManager {
	return &ArrivalTimeManager{
		core: core,
	}
}

// ShipmentsInfo contain info about local time and delivery date
type ShipmentsInfo struct {
	LocalTime   string
	DeliverDays int //optional param
}

// LunarArrivalTimeResponse response for GetLunarArrivalTime
type LunarArrivalTimeResponse struct {
	LunarArrivalTime string
}

// RequestError structure indicate errors while request
type RequestError struct {
	Code  int
	Error error
}

// GetLunarArrivalTime returns lunar arrival time
func (atm *ArrivalTimeManager) GetLunarArrivalTime(shipmentsInfo string) (*LunarArrivalTimeResponse, *RequestError) {
	log.Info("[+] GetLunarArrivalTime")
	defer log.Info("[-] GetLunarArrivalTime")

	shipmentsInfoData := ShipmentsInfo{}
	err := json.Unmarshal([]byte(shipmentsInfo), &shipmentsInfoData)
	if err != nil {
		log.Error(err)
		return nil, &RequestError{
			Code:  ErrorInvalidRequestFormat,
			Error: fmt.Errorf(ErrorInvalidRequestFormatMsg),
		}
	}

	parsed_time, err := time.Parse(atm.core.context.GlobalString(TimeInputLayoutName), shipmentsInfoData.LocalTime)
	if err != nil {
		log.Error("invalid input time, ", err)
		return nil, &RequestError{
			Code:  ErrorInvalidFormatTimeBody,
			Error: fmt.Errorf(ErrorInvalidFormatTimeMsg, err),
		}
	}
	log.Info("parsed time: ", parsed_time)

	utc_time := parsed_time.UTC()
	log.Info("utc_time time: ", utc_time)

	deliver_time := atm.core.context.GlobalInt(DeliverTimeName)
	if shipmentsInfoData.DeliverDays != 0 {
		deliver_time = shipmentsInfoData.DeliverDays
	}

	utc_deliver_time_earth_time := utc_time.AddDate(0, 0, deliver_time)
	log.Info("utc_time time: ", utc_deliver_time_earth_time)
	log.Info("utc_time time nanoseconds: ", utc_deliver_time_earth_time.Unix())

	response := LunarArrivalTimeResponse{}
	response.LunarArrivalTime = GetLunarTime(utc_deliver_time_earth_time)

	return &response, nil
}
