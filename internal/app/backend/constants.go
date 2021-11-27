package backend

// ApplicationName is literal name of application
const ApplicationName = "Arrival Time Backend Service"

// ApplicationUsage is description of application
const ApplicationUsage = "Arrival Time Backend Service is used for getting time information between two points"

// PortName - param name
const PortName = "port"

// PortUsage - description of port param
const PortUsage = "Port number of HTTP(S) listener"

// PortDefault default value of port param
const PortDefault = "8080"

// PortEnvVar - name of environment variable for port param
const PortEnvVar = "BACKEND_HTTP_LISTENER_PORT"

// TLSModeName name of environment variable of tls mode
const TLSModeName = "TLSMode"

// TLSModeUsage possible input values
const TLSModeUsage = "0 - http, 1 - https"

// TLSModeDefault default value for TLS mode
const TLSModeDefault = "0"

// TLSModeEnvVar name of environment variable of TLS mode
const TLSModeEnvVar = "TLS_MODE"

// CerFileName name of environment variable of cer file
const CerFileName = "cer_file"

// CerFileUsage input value
const CerFileUsage = "Certificate filename with extension"

// CerFileDefault default value for certificate file
const CerFileDefault = "DeviceCert.crt"

// CerFileEnvVar name of environment certificate file variable
const CerFileEnvVar = "CER_FILE"

// KeyFileName name of environment variable of key file
const KeyFileName = "key_file"

// KeyFileUsage input value
const KeyFileUsage = "Key filename with extension"

// KeyFileDefault default value for key file
const KeyFileDefault = "DeviceKey.key"

// KeyFileEnvVar name of environment key file variable
const KeyFileEnvVar = "KEY_FILE"

// TimeInputLayoutName name of environment variable of Time Input
const TimeInputLayoutName = "time_input_layout_name"

// TimeInputLayoutUsage input value
const TimeInputLayoutUsage = "The default format of date input"

// TimeInputLayoutDefault default value for Time Input
const TimeInputLayoutDefault = "2006-01-02 15:04:05 -0700"

// TimeInputLayoutEnvVar name of environment Time Input variable
const TimeInputLayoutEnvVar = "TIME_INPUT_LAYOUT"

// DeliverTimeName name of environment variable of Time Deliver from earth to moon
const DeliverTimeName = "deliver_time_name"

// DeliverTimeUsage input value
const DeliverTimeUsage = "The default value of moon deliver time"

// DeliverTimeDefault default value for DeliverTime
const DeliverTimeDefault = 4

// DeliverTimeEnvVar name of environment Deliver Time variable
const DeliverTimeEnvVar = "DELIVER_TIME_LAYOUT"

// Error codes section - begin

// NoError indicate when no error occurs
const NoError = 0

// NoErrorMsg message of NoError
const NoErrorMsg = "successful"

// ErrorReadRequestBody is error code that returns to a client when can't read request body
const ErrorReadRequestBody = 1000

// ErrorReadRequestBodyMsg error message for ErrorReadRequestBody error code
const ErrorReadRequestBodyMsg = "Can't read request body"

// InvalidFormatTime is error code that returns to a client when can't parse input time
const ErrorInvalidFormatTimeBody = 1001

// ErrorInvalidFormatTimeMsg error message for ErrorInvalidFormatTimeBody error code
const ErrorInvalidFormatTimeMsg = "invalid input time %v"

// ErrorInvalidRequestFormat indicate validation errors
const ErrorInvalidRequestFormat = 1002

// ErrorInvalidRequestFormatMsg message for ErrorInvalidRequestFormat
const ErrorInvalidRequestFormatMsg = "invalid request format"

// ErrorUnknown unknown error code
const ErrorUnknown = 1003

// ErrorUnknownMsg unknown error message
const ErrorUnknownMsg = "unknown error occurs [%v]"

// Error codes section - end
