package backend

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// test SetupAppFlags function
func TestSetupAppFlags(t *testing.T) {
	setupAppFlags := SetupAppFlags()
	assert.NotNil(t, setupAppFlags)
}

// test BuildResponse function response structure
func TestBuildResponse(t *testing.T) {
	buildResponse := BuildResponse(1, "success")
	assert.NotEmpty(t, buildResponse)
	assert.Equal(t, buildResponse, "{\"ErrorCode\":1,\"ErrorMessage\":\"success\"}")

	buildResponse = BuildResponse(ErrorReadRequestBody, ErrorReadRequestBodyMsg)
	assert.NotEmpty(t, buildResponse)
	assert.Equal(t, buildResponse, "{\"ErrorCode\":1000,\"ErrorMessage\":\"Can't read request body\"}")
}

// test GetLunarTime function for getting correct date result
func TestGetLunarTime(t *testing.T) {
	parsed_time, _ := time.Parse("2006-01-02 15:04:05 -0700", "2020-01-01 10:10:30 +0300")
	getLunarTime := GetLunarTime(parsed_time)
	assert.NotEmpty(t, getLunarTime)
	assert.Equal(t, getLunarTime, "52-12-30 V 01:47:20")

	parsed_time, _ = time.Parse("2006-01-02 15:04:05 -0700", "1600-04-10 10:10:30 +0400")
	getLunarTime = GetLunarTime(parsed_time)
	assert.NotEmpty(t, getLunarTime)
	assert.Equal(t, getLunarTime, "-381-08-10 V 19:39:02")

	parsed_time, _ = time.Parse("2006-01-02 15:04:05 -0700", "1969-07-21 02:56:15 +0000")
	getLunarTime = GetLunarTime(parsed_time)
	assert.NotEmpty(t, getLunarTime)
	assert.Equal(t, getLunarTime, "1-01-01 V 00:00:00")
}

// test GenPassword function
func TestTwoDigitRepresentation(t *testing.T) {
	twoDigitRepresentation := TwoDigitRepresentation(-5)
	assert.NotEmpty(t, twoDigitRepresentation)
	assert.Equal(t, twoDigitRepresentation, "-")

	twoDigitRepresentation = TwoDigitRepresentation(0)
	assert.NotEmpty(t, twoDigitRepresentation)
	assert.Equal(t, twoDigitRepresentation, "00")

	twoDigitRepresentation = TwoDigitRepresentation(8)
	assert.NotEmpty(t, twoDigitRepresentation)
	assert.Equal(t, twoDigitRepresentation, "08")

	twoDigitRepresentation = TwoDigitRepresentation(10)
	assert.NotEmpty(t, twoDigitRepresentation)
	assert.Equal(t, twoDigitRepresentation, "10")

	twoDigitRepresentation = TwoDigitRepresentation(15)
	assert.NotEmpty(t, twoDigitRepresentation)
	assert.Equal(t, twoDigitRepresentation, "15")
}
