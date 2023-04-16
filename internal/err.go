package internal

/*
#include "miniaudio.h"
#cgo LDFLAGS: -lm
*/
import "C"

type MaResult = C.ma_result

const (
	MaResultSuccess                    = MaResult(C.MA_SUCCESS)
	MaResultFailed                     = MaResult(C.MA_ERROR)
	MaResultInvalidArgs                = MaResult(C.MA_INVALID_ARGS)
	MaResultInvalidOperation           = MaResult(C.MA_INVALID_OPERATION)
	MaResultOutOfMemory                = MaResult(C.MA_OUT_OF_MEMORY)
	MaResultOutOfRange                 = MaResult(C.MA_OUT_OF_RANGE)
	MaResultAccessDenied               = MaResult(C.MA_ACCESS_DENIED)
	MaResultDoesNotExists              = MaResult(C.MA_DOES_NOT_EXIST)
	MaResultAlreadyExists              = MaResult(C.MA_ALREADY_EXISTS)
	MaResultTooManyOpenFiles           = MaResult(C.MA_TOO_MANY_OPEN_FILES)
	MaResultInvalidFile                = MaResult(C.MA_INVALID_FILE)
	MaResultTooBig                     = MaResult(C.MA_TOO_BIG)
	MaResultPathTooLong                = MaResult(C.MA_PATH_TOO_LONG)
	MaResultNameTooLong                = MaResult(C.MA_NAME_TOO_LONG)
	MaResultNotDirectory               = MaResult(C.MA_NOT_DIRECTORY)
	MaResultIsDirectory                = MaResult(C.MA_IS_DIRECTORY)
	MaResultDirectoryNotEmpty          = MaResult(C.MA_DIRECTORY_NOT_EMPTY)
	MaResultAtEnd                      = MaResult(C.MA_AT_END)
	MaResultNoSpace                    = MaResult(C.MA_NO_SPACE)
	MaResultBusy                       = MaResult(C.MA_BUSY)
	MaResultIoResultor                 = MaResult(C.MA_IO_ERROR)
	MaResultInterrupt                  = MaResult(C.MA_INTERRUPT)
	MaResultUnavailable                = MaResult(C.MA_UNAVAILABLE)
	MaResultAlreadyInUse               = MaResult(C.MA_ALREADY_IN_USE)
	MaResultBadAddress                 = MaResult(C.MA_BAD_ADDRESS)
	MaResultBadSeek                    = MaResult(C.MA_BAD_SEEK)
	MaResultBadPipe                    = MaResult(C.MA_BAD_PIPE)
	MaResultDeadlock                   = MaResult(C.MA_DEADLOCK)
	MaResultTooManyLinks               = MaResult(C.MA_TOO_MANY_LINKS)
	MaResultNotImplemented             = MaResult(C.MA_NOT_IMPLEMENTED)
	MaResultNoMessage                  = MaResult(C.MA_NO_MESSAGE)
	MaResultBadMessage                 = MaResult(C.MA_BAD_MESSAGE)
	MaResultNoDataAvailable            = MaResult(C.MA_NO_DATA_AVAILABLE)
	MaResultInvalidData                = MaResult(C.MA_INVALID_DATA)
	MaResultTimeout                    = MaResult(C.MA_TIMEOUT)
	MaResultNoNetwork                  = MaResult(C.MA_NO_NETWORK)
	MaResultNotUnique                  = MaResult(C.MA_NOT_UNIQUE)
	MaResultNotSocket                  = MaResult(C.MA_NOT_SOCKET)
	MaResultNoAddress                  = MaResult(C.MA_NO_ADDRESS)
	MaResultBadProtocol                = MaResult(C.MA_BAD_PROTOCOL)
	MaResultProtocolUnavailable        = MaResult(C.MA_PROTOCOL_UNAVAILABLE)
	MaResultProtocolNotSupported       = MaResult(C.MA_PROTOCOL_NOT_SUPPORTED)
	MaResultProtocolFamilyNotSupported = MaResult(C.MA_PROTOCOL_FAMILY_NOT_SUPPORTED)
	MaResultAddressFamilyNotSupported  = MaResult(C.MA_ADDRESS_FAMILY_NOT_SUPPORTED)
	MaResultSocketNotSupported         = MaResult(C.MA_SOCKET_NOT_SUPPORTED)
	MaResultConnectionReset            = MaResult(C.MA_CONNECTION_RESET)
	MaResultAlreadyConnected           = MaResult(C.MA_ALREADY_CONNECTED)
	MaResultNotConnected               = MaResult(C.MA_NOT_CONNECTED)
	MaResultConnectionRefused          = MaResult(C.MA_CONNECTION_REFUSED)
	MaResultNoHost                     = MaResult(C.MA_NO_HOST)
	MaResultInProgress                 = MaResult(C.MA_IN_PROGRESS)
	MaResultCancelled                  = MaResult(C.MA_CANCELLED)
	MaResultMemoryAlreadyMapped        = MaResult(C.MA_MEMORY_ALREADY_MAPPED)

	/* General miniaudio-specific errors. */

	MaResultFormatNotSupported     = MaResult(C.MA_FORMAT_NOT_SUPPORTED)
	MaResultDeviceTypeNotSupported = MaResult(C.MA_DEVICE_TYPE_NOT_SUPPORTED)
	MaResultShareModeNotSupported  = MaResult(C.MA_SHARE_MODE_NOT_SUPPORTED)
	MaResultNoBackend              = MaResult(C.MA_NO_BACKEND)
	MaResultNoDevice               = MaResult(C.MA_NO_DEVICE)
	MaResultApiNotFound            = MaResult(C.MA_API_NOT_FOUND)
	MaResultInvalidDeviceConfig    = MaResult(C.MA_INVALID_DEVICE_CONFIG)
	MaResultLoop                   = MaResult(C.MA_LOOP)

	/* State errors. */
	MaResultDeviceNotInitialized     = MaResult(C.MA_DEVICE_NOT_INITIALIZED)
	MaResultDeviceAlreadyInitialized = MaResult(C.MA_DEVICE_ALREADY_INITIALIZED)
	MaResultDeviceNotStarted         = MaResult(C.MA_DEVICE_NOT_STARTED)
	MaResultDeviceNotStopped         = MaResult(C.MA_DEVICE_NOT_STOPPED)

	/* Operation errors. */
	MaResultFailedToInitBackend        = MaResult(C.MA_FAILED_TO_INIT_BACKEND)
	MaResultFailedToOpenBackendDevice  = MaResult(C.MA_FAILED_TO_OPEN_BACKEND_DEVICE)
	MaResultFailedToStartBackendDevice = MaResult(C.MA_FAILED_TO_START_BACKEND_DEVICE)
	MaResultFailedToStopBackendDevice  = MaResult(C.MA_FAILED_TO_STOP_BACKEND_DEVICE)
)

func (e MaResult) Error() string {
	switch e {
	case MaResultSuccess:
		return "Success"
	case MaResultFailed:
		return "Failed"
	case MaResultInvalidArgs:
		return "InvalidArgs"
	case MaResultInvalidOperation:
		return "InvalidOperation"
	case MaResultOutOfMemory:
		return "OutOfMemory"
	case MaResultOutOfRange:
		return "OutOfRange"
	case MaResultAccessDenied:
		return "AccessDenied"
	case MaResultDoesNotExists:
		return "DoesNotExists"
	case MaResultAlreadyExists:
		return "AlreadyExists"
	case MaResultTooManyOpenFiles:
		return "TooManyOpenFiles"
	case MaResultInvalidFile:
		return "InvalidFile"
	case MaResultTooBig:
		return "TooBig"
	case MaResultPathTooLong:
		return "PathTooLong"
	case MaResultNameTooLong:
		return "NameTooLong"
	case MaResultNotDirectory:
		return "NotDirectory"
	case MaResultIsDirectory:
		return "IsDirectory"
	case MaResultDirectoryNotEmpty:
		return "DirectoryNotEmpty"
	case MaResultAtEnd:
		return "AtEnd"
	case MaResultNoSpace:
		return "NoSpace"
	case MaResultBusy:
		return "Busy"
	case MaResultIoResultor:
		return "IoMaResultor"
	case MaResultInterrupt:
		return "Interrupt"
	case MaResultUnavailable:
		return "Unavailable"
	case MaResultAlreadyInUse:
		return "AlreadyInUse"
	case MaResultBadAddress:
		return "BadAddress"
	case MaResultBadSeek:
		return "BadSeek"
	case MaResultBadPipe:
		return "BadPipe"
	case MaResultDeadlock:
		return "Deadlock"
	case MaResultTooManyLinks:
		return "TooManyLinks"
	case MaResultNotImplemented:
		return "NotImplemented"
	case MaResultNoMessage:
		return "NoMessage"
	case MaResultBadMessage:
		return "BadMessage"
	case MaResultNoDataAvailable:
		return "NoDataAvailable"
	case MaResultInvalidData:
		return "InvalidData"
	case MaResultTimeout:
		return "Timeout"
	case MaResultNoNetwork:
		return "NoNetwork"
	case MaResultNotUnique:
		return "NotUnique"
	case MaResultNotSocket:
		return "NotSocket"
	case MaResultNoAddress:
		return "NoAddress"
	case MaResultBadProtocol:
		return "BadProtocol"
	case MaResultProtocolUnavailable:
		return "ProtocolUnavailable"
	case MaResultProtocolNotSupported:
		return "ProtocolNotSupported"
	case MaResultProtocolFamilyNotSupported:
		return "ProtocolFamilyNotSupported"
	case MaResultAddressFamilyNotSupported:
		return "AddressFamilyNotSupported"
	case MaResultSocketNotSupported:
		return "SocketNotSupported"
	case MaResultConnectionReset:
		return "ConnectionReset"
	case MaResultAlreadyConnected:
		return "AlreadyConnected"
	case MaResultNotConnected:
		return "NotConnected"
	case MaResultConnectionRefused:
		return "ConnectionRefused"
	case MaResultNoHost:
		return "NoHost"
	case MaResultInProgress:
		return "InProgress"
	case MaResultCancelled:
		return "Cancelled"
	case MaResultMemoryAlreadyMapped:
		return "MemoryAlreadyMapped"
	case MaResultFormatNotSupported:
		return "FormatNotSupported"
	case MaResultDeviceTypeNotSupported:
		return "DeviceTypeNotSupported"
	case MaResultShareModeNotSupported:
		return "ShareModeNotSupported"
	case MaResultNoBackend:
		return "NoBackend"
	case MaResultNoDevice:
		return "NoDevice"
	case MaResultApiNotFound:
		return "ApiNotFound"
	case MaResultInvalidDeviceConfig:
		return "InvalidDeviceConfig"
	case MaResultLoop:
		return "Loop"
	case MaResultDeviceNotInitialized:
		return "DeviceNotInitialized"
	case MaResultDeviceAlreadyInitialized:
		return "DeviceAlreadyInitialized"
	case MaResultDeviceNotStarted:
		return "DeviceNotStarted"
	case MaResultDeviceNotStopped:
		return "DeviceNotStopped"
	case MaResultFailedToInitBackend:
		return "FailedToInitBackend"
	case MaResultFailedToOpenBackendDevice:
		return "FailedToOpenBackendDevice"
	case MaResultFailedToStartBackendDevice:
		return "FailedToStartBackendDevice"
	case MaResultFailedToStopBackendDevice:
		return "FailedToStopBackendDevice"
	default:
		return "Failed"
	}
}
