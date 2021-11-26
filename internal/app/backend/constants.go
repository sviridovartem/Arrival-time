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

// The default value is only needed when SSL/TLS certificate is not valid on server.
// In production InsecureSkipVerify should be set to false.

// InsecureSkipVerifyName name of environment variable of Insecure Skip Verify
const InsecureSkipVerifyName = "insecure_skip_verify_name"

// InsecureSkipVerifyUsage input value
const InsecureSkipVerifyUsage = "The default value is only needed when SSL/TLS certificate is not valid on server"

// InsecureSkipVerifyDefault default value for Insecure Skip Verify
const InsecureSkipVerifyDefault = "true"

// InsecureSkipVerifyEnvVar name of environment Insecure Skip Verify variable
const InsecureSkipVerifyEnvVar = "INSECURE_SKIP_VERIFY"

// Error codes section - begin

// NoError indicate when no error occurs
const NoError = 0

// NoErrorMsg message of NoError
const NoErrorMsg = "successful"

// ErrorReadRequestBody is error code that returns to a client when can't read request body
const ErrorReadRequestBody = 1000

// ErrorReadRequestBodyMsg error message for ErrorReadRequestBody error code
const ErrorReadRequestBodyMsg = "Can't read request body"

// ErrorInvalidRequestFormat indicate validation errors
const ErrorInvalidRequestFormat = 1003

// ErrorInvalidRequestFormatMsg message for ErrorInvalidRequestFormat
const ErrorInvalidRequestFormatMsg = "invalid request format"

// ErrorUnknown unknown error code
const ErrorUnknown = 1004

// ErrorUnknownMsg unknown error message
const ErrorUnknownMsg = "unknown error occurs [%v]"

// Error codes section - end
