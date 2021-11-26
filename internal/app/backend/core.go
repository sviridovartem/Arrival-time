package backend

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/jasonlvhit/gocron"
	"github.com/urfave/cli"
)

// InstanceCore instance of core module structure
type InstanceCore struct {
	context            *cli.Context
	arrivalTimeManager IArrivalTimeManager
}

func newInstance(ctx *cli.Context) *InstanceCore {
	return &InstanceCore{
		context: ctx,
	}
}
func checkCertificate(cerFileCreateTime time.Time, keyFileCreateTime time.Time, cerFile string, keyFile string) {
	cerFileInfo, err := os.Stat(cerFile)
	if err != nil {
		panic(fmt.Errorf("can't get CerFileInfo in checkCertificate func, %v", err))
	}
	keyFileInfo, err := os.Stat(keyFile)
	if err != nil {
		panic(fmt.Errorf("can't get KeyFileInfo in checkCertificate func, %v", err))
	}

	if cerFileInfo.ModTime() != cerFileCreateTime ||
		keyFileInfo.ModTime() != keyFileCreateTime {
		panic(fmt.Errorf("certificate or key updated"))
	}
}

func scheduleCertificateCheck(cerFileCreateTime time.Time, keyFileCreateTime time.Time, cerFile string, keyFile string) {
	gocron.Every(5).Seconds().Do(checkCertificate, cerFileCreateTime, keyFileCreateTime, cerFile, keyFile)
	<-gocron.Start()
}

// Main - Entry point of backend core
func Main(context *cli.Context) error {
	instance := newInstance(context)
	api := NewAPIHandlers(instance)
	api.SetupHandlers()

	instance.SetArrivalTimeManager(NewArrivalTimeManager(instance))

	if context.GlobalString(TLSModeName) == "0" {
		srv := &http.Server{
			Addr:    fmt.Sprintf(":%s", context.GlobalString(PortName)),
			Handler: nil,
		}
		return srv.ListenAndServe()
	} else if context.GlobalString(TLSModeName) == "1" {
		cert, err := tls.LoadX509KeyPair(context.GlobalString(CerFileName), context.GlobalString(KeyFileName))
		if err != nil {
			panic(err)
		}
		CerFileInfo, err := os.Stat(context.GlobalString(CerFileName))
		if err != nil {
			panic(fmt.Errorf("can't get CerFileInfo, %v", err))
		}
		KeyFileInfo, err := os.Stat(context.GlobalString(KeyFileName))
		if err != nil {
			panic(fmt.Errorf("can't get KeyFileInfo, %v", err))
		}
		go scheduleCertificateCheck(
			CerFileInfo.ModTime(),
			KeyFileInfo.ModTime(),
			context.GlobalString(CerFileName),
			context.GlobalString(KeyFileName))

		srv := &http.Server{
			Addr:      fmt.Sprintf(":%s", context.GlobalString(PortName)),
			Handler:   nil,
			TLSConfig: &tls.Config{Certificates: []tls.Certificate{cert}},
		}
		return srv.ListenAndServeTLS("", "")
	} else {
		panic(fmt.Errorf("you have wrong TLSMode param, it can be only '0' or '1' but you have %s", context.GlobalString(TLSModeName)))
	}
}

// SetArrivalTimeManager setup ArrivalTimeManager instance
func (instance *InstanceCore) SetArrivalTimeManager(atm IArrivalTimeManager) {
	instance.arrivalTimeManager = atm
}

// GetArrivalTimeManager returns arrival time manager instance
func (instance *InstanceCore) GetArrivalTimeManager() IArrivalTimeManager {
	return instance.arrivalTimeManager
}
