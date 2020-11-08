package httpx

import (
	"errors"
	"strings"

	"goby/pkg/dict"
)

// Some http custom statuses
const (
	StatusParseParamsError     = 2000
	StatusUnauthorized         = 2001
	StatusQueryDataError       = 2002
	StatusSaveDataError        = 2003
	StatusUpdateDataError      = 2004
	StatusDeleteDataError      = 2005
	StatusRunningError         = 2006
	StatusParseDateTimeError   = 2007
	StatusDataTooLarge         = 2008
	StatusDataNotFound         = 2009
	StatusOperatoinFailed      = 2010
	StatusResourceAlreayExists = 2011
	StatusInvalid              = 2012
	StatusActionDenied         = 2013
	StatusGetTokenError        = 2014
	StatusNoTokenFound         = 2015
	StatusServiceLockedup      = 2016
	StatusUnknown              = 2100
)

var statusText = map[int]string{
	StatusParseParamsError:     "Parse parameters error",
	StatusUnauthorized:         "Unauthorized",
	StatusQueryDataError:       "Query data error",
	StatusSaveDataError:        "Save data error",
	StatusUpdateDataError:      "Update data error",
	StatusDeleteDataError:      "Delete data error",
	StatusRunningError:         "Running error",
	StatusParseDateTimeError:   "Parse date & time error",
	StatusDataTooLarge:         "Data too large",
	StatusDataNotFound:         "Data not found",
	StatusOperatoinFailed:      "Operation failed",
	StatusResourceAlreayExists: "Resource already exists",
	StatusInvalid:              "Invalid",
	StatusActionDenied:         "Action has been denied",
	StatusGetTokenError:        "Get user token error",
	StatusNoTokenFound:         "No token found",
	StatusServiceLockedup:      "Service has been locked up to specific client due to limit policy",
	StatusUnknown:              "Unknown status",
}

// StatusText returns a text for the error Status.
// Returns empty string if the Status is unknow.
func StatusText(code int) string {
	return statusText[code]
}

// StatusTextWithExtra :
func StatusTextWithExtra(code int, extra string) string {
	return strings.Join([]string{statusText[code], extra}, dict.RightArrow)
}

// NewError :
func NewError(errMsg string) error {
	return errors.New(errMsg)
}
