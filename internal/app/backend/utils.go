package backend

import (
	"encoding/json"
	"fmt"
	"time"

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
			Name:   TimeInputLayoutName,
			Usage:  TimeInputLayoutUsage,
			Value:  TimeInputLayoutDefault,
			EnvVar: TimeInputLayoutEnvVar,
		},
		cli.IntFlag{
			Name:   DeliverTimeName,
			Usage:  DeliverTimeUsage,
			Value:  DeliverTimeDefault,
			EnvVar: DeliverTimeEnvVar,
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

// Lunar Time Calculation
func GetLunarTime(inputTime time.Time) string {
	LST_f := 0.9843529666671
	LST_c := 14159025
	m := 60
	h := 3600
	p := 86400
	d := 2592000
	y := 31104000

	//todo fix here everyting
	s := (float64(inputTime.Unix()) + float64(LST_c)) / LST_f

	n := ""
	if s < 0 {
		n = "-"
	}
	if s < 0 {
		s = -s
	}
	Y := int(s / float64(y))
	k := float64(y * Y)
	Y = Y + 1
	s = s - k
	D := int(s / float64(d))
	s = s - float64(d*(D))
	D = D + 1
	dd := DD(D)
	C := int(s / float64(p))
	s = s - float64(p*C)
	C = C + 1
	cc := DD(C)
	H := int(s / float64(h))
	s = s - float64(h*H)
	hh := DD(H)
	M := int(s / float64(m))
	s = s - float64(m*M)
	mm := DD(M)
	ss := DD(int(s / 1))

	if Y+D+C == 0 {
		return (n + "V " + hh + ":" + mm + ":" + ss)
	} else if Y+D == 0 {
		return (n + cc + " V " + hh + ":" + mm + ":" + ss)
	} else if Y == 0 {
		return (n + dd + "-" + cc + " V " + hh + ":" + mm + ":" + ss)
	} else {
		return (n + fmt.Sprint(Y) + "-" + dd + "-" + cc + " V " + hh + ":" + mm + ":" + ss)
	}

}

// Lunar Time Calculation (DD)
func DD(n int) string {
	if n < 10 {
		return fmt.Sprintf("0%d", n)
	} else {
		return fmt.Sprint(n)
	}
}
