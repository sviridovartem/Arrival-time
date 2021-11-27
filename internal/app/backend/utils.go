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
	seconds := (float64(inputTime.Unix()) + LunarConstant) / LunarSecond
	n := ""
	if seconds < 0 {
		n = "-"
		seconds = -seconds
	}
	lunarYear := int(seconds / SecondInYear)
	seconds = seconds - float64(SecondInYear*lunarYear)
	lunarYear = lunarYear + 1

	lunarDay := int(seconds / SecondInMonth)
	seconds = seconds - float64(SecondInMonth*lunarDay)
	lunarDay = lunarDay + 1
	lunarDayResult := TwoDigitRepresentation(lunarDay)

	lunarCycle := int(seconds / SecondInDay)
	seconds = seconds - float64(SecondInDay*lunarCycle)
	lunarCycle = lunarCycle + 1
	lunarCycleResult := TwoDigitRepresentation(lunarCycle)

	lunarHours := int(seconds / SecondInHours)
	seconds = seconds - float64(SecondInHours*lunarHours)
	lunarHoursResult := TwoDigitRepresentation(lunarHours)

	lunarMinutes := int(seconds / SecondInMinutes)
	seconds = seconds - float64(SecondInMinutes*lunarMinutes)
	lunarMinutesResult := TwoDigitRepresentation(lunarMinutes)

	lunarSecondsResult := TwoDigitRepresentation(int(seconds / 1))

	if lunarYear+lunarDay+lunarCycle == 0 {
		return fmt.Sprint(n, "V ", lunarHoursResult, ":", lunarMinutesResult, ":", lunarSecondsResult)
	} else if lunarYear+lunarDay == 0 {
		return fmt.Sprint(n, lunarCycleResult, " V ", lunarHoursResult, ":", lunarMinutesResult, ":", lunarSecondsResult)
	} else if lunarYear == 0 {
		return fmt.Sprint(n, lunarDayResult, "-", lunarCycleResult, " V ", lunarHoursResult, ":", lunarMinutesResult, ":", lunarSecondsResult)
	} else {
		return fmt.Sprint(n, lunarYear, "-", lunarDayResult, "-", lunarCycleResult, " V ", lunarHoursResult, ":", lunarMinutesResult, ":", lunarSecondsResult)
	}

}

// TwoDigitRepresentation returns string value of two character
func TwoDigitRepresentation(n int) string {
	if n < 10 {
		return fmt.Sprintf("0%d", n)
	} else {
		return fmt.Sprint(n)
	}
}
