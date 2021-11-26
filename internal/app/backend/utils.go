package backend

import (
	"encoding/json"

	"github.com/urfave/cli"
)

type errorResponse struct {
	ErrorCode    int
	ErrorMessage string
}

// SetupAppFlags sets array of supported command line arguments
func SetupAppFlags() []cli.Flag {

	return []cli.Flag{
		cli.StringFlag{
			Name:   PortName,
			Usage:  PortUsage,
			Value:  PortDefault,
			EnvVar: PortEnvVar,
		},
		cli.StringFlag{
			Name:   TLSModeName,
			Usage:  TLSModeUsage,
			Value:  TLSModeDefault,
			EnvVar: TLSModeEnvVar,
		},
		cli.StringFlag{
			Name:   CerFileName,
			Usage:  CerFileUsage,
			Value:  CerFileDefault,
			EnvVar: CerFileEnvVar,
		},
		cli.StringFlag{
			Name:   KeyFileName,
			Usage:  KeyFileUsage,
			Value:  KeyFileDefault,
			EnvVar: KeyFileEnvVar,
		},
		cli.StringFlag{
			Name:   InsecureSkipVerifyName,
			Usage:  InsecureSkipVerifyUsage,
			Value:  InsecureSkipVerifyDefault,
			EnvVar: InsecureSkipVerifyEnvVar,
		},
	}
}

// BuildResponse create JSON based response
func BuildResponse(code int, message string) string {
	resp := &errorResponse{
		ErrorCode:    code,
		ErrorMessage: message,
	}
	b, _ := json.Marshal(resp)
	return string(b)
}
