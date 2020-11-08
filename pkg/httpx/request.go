package httpx

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Request : http request wrapper
type Request struct {
	Request *http.Request
}

// GetRequest : get http request
func (req *Request) GetRequest() *http.Request {
	return req.Request
}

// GetCookieValue :
func (req *Request) GetCookieValue(cookieKey string) string {
	var val string
	cookie, err := req.GetRequest().Cookie(cookieKey)
	if err != nil {
		// do nothing here
	}
	if cookie != nil {
		val = cookie.Value
	}
	return val
}

// GetHeaderValue :
func (req *Request) GetHeaderValue(header string) string {
	return req.GetRequest().Header.Get(header)
}

// ParseParams :
func (req *Request) ParseParams(args ...interface{}) error {
	vals := mux.Vars(req.Request)
	var err error

	for i := 0; i < len(args); i += 2 {
		if len(args) < (i + 1) {
			err = errors.New("error of parsing parameters")
			break
		}

		dest := args[i+1]
		destValue := reflect.Indirect(reflect.ValueOf(dest))
		if reflect.TypeOf(dest).Kind() != reflect.Ptr {
			err = errors.New("destination value must be a pointer")
			log.Println(err)
			return err
		}

		paramName := (args[i]).(string)
		val := vals[paramName]
		// fmt.Println("vals => ", vals)
		// fmt.Println("destVal => ", destValue, "parameter name => ", parameterName, " dest => ", dest, " value => ", val)
		if val == "" {
			val = req.Request.FormValue(paramName)
			if val == "" {
				return errors.New(paramName + " is required")
			}
		}

		switch dest.(type) {
		case *string:
			destValue.SetString(val)
		case *bool:
			v, err := strconv.ParseBool(val)
			if err != nil {
				return errors.New("parse " + paramName + " as boolean value error")
			}
			destValue.SetBool(v)
		case *int:
			v, err := strconv.ParseInt(val, 10, 32)
			if err != nil {
				err = errors.New("parse " + paramName + " as int error")
				break
			}
			destValue.SetInt(v)
		case *int8:
			v, err := strconv.ParseInt(val, 10, 8)
			if err != nil {
				return errors.New("parse " + paramName + " as int8 error")
			}
			destValue.SetInt(v)
		case *int16:
			v, err := strconv.ParseInt(val, 10, 16)
			if err != nil {
				return errors.New("parse " + paramName + " as int16 error")
			}
			destValue.SetInt(v)
		case *int32:
			v, err := strconv.ParseInt(val, 10, 32)
			if err != nil {
				return errors.New("parse " + paramName + " as int32 error")
			}
			destValue.SetInt(v)
		case *int64:
			v, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return errors.New("parse " + paramName + " as int64 error")
			}
			destValue.SetInt(v)
		case *uint:
			v, err := strconv.ParseUint(val, 10, 32)
			if err != nil {
				return errors.New("parse " + paramName + " as uint32 error")
			}
			destValue.SetUint(v)
		case *uint8:
			v, err := strconv.ParseUint(val, 10, 8)
			if err != nil {
				err = errors.New("parse " + paramName + " as uint8 error")
				break
			}
			destValue.SetUint(v)
		case *uint16:
			v, err := strconv.ParseUint(val, 10, 16)
			if err != nil {
				return errors.New("parse " + paramName + " as uint16 error")
			}
			destValue.SetUint(v)
		case *uint32:
			v, err := strconv.ParseUint(val, 10, 32)
			if err != nil {
				return errors.New("parse " + paramName + " as uint32 error")
			}
			destValue.SetUint(v)
		case *uint64:
			v, err := strconv.ParseUint(val, 10, 64)
			if err != nil {
				return errors.New("parse " + paramName + " as uint64 error")
			}
			destValue.SetUint(v)
		case *float32:
			v, err := strconv.ParseFloat(val, 32)
			if err != nil {
				err = errors.New("parse " + paramName + " as float32 error")
				break
			}
			destValue.SetFloat(v)
		case *float64:
			v, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return errors.New("parse " + paramName + " as float64 error")
			}
			destValue.SetFloat(v)
		default:
			return errors.New("Unsupported type => " + reflect.TypeOf(dest).Kind().String())
		}
	}

	return err
}

// ParseBody : parse request body
func (req *Request) ParseBody(dest interface{}) error {
	decoder := json.NewDecoder(req.GetRequest().Body)
	if err := decoder.Decode(dest); err != nil && err != io.EOF {
		log.Error(err)
		return err
	}
	return nil
}
