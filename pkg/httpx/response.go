package httpx

import (
	"encoding/json"
	"net/http"
	"strings"

	"goby/pkg/dict"

	log "github.com/sirupsen/logrus"
)

// Response : http ResponseWriter wrapper
type Response struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"msg"`
	Stream  bool        `json:"-"`

	Writer     http.ResponseWriter `json:"-"`
	hasWritten bool
}

// GetWriter : get http response writer
func (resp *Response) GetWriter() http.ResponseWriter {
	return resp.Writer
}

// WriteJSON : write response data
func (resp *Response) WriteJSON(data interface{}) {
	resp.Data = data
	dataBytes, err := json.Marshal(resp)
	if err != nil {
		log.Panic(err)
	}
	resp.Writer.Header().Set("Content-Type", "application/json")
	resp.Writer.Write(dataBytes)
	resp.hasWritten = true
}

// Write : write response data
func (resp *Response) Write(data []byte) {
	// resp.Data = data
	// dataBytes, err := json.Marshal(resp)
	// if err != nil {
	// 	log.Panic(err)
	// }
	resp.Writer.Write(data)
	resp.hasWritten = true
}

// WriteStream :
func (resp *Response) WriteStream(data interface{}) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		log.Panic(err)
	}
	resp.Writer.Write(dataBytes)
	resp.hasWritten = true
}

func (resp *Response) handleStreamingError() {
	resp.WriteStream("Server error: " + resp.Message)
}

// handle http status
func (resp *Response) handleErrrorStatuses(w http.ResponseWriter) {
	if resp.Stream {
		resp.handleStreamingError()
		return
	}
	if !resp.hasWritten {
		if resp.Status == http.StatusOK {
			resp.Status = StatusRunningError
			if strings.EqualFold(resp.Message, dict.Blank) {
				resp.Message = StatusText(StatusRunningError)
			}
		} else {
			if !strings.EqualFold(http.StatusText(resp.Status), dict.Blank) {
				if strings.EqualFold(resp.Message, dict.Blank) {
					resp.Message = http.StatusText(resp.Status)
				}
			} else {
				if strings.EqualFold(StatusText(resp.Status), dict.Blank) {
					resp.Status = StatusUnknown
				}
				if strings.EqualFold(resp.Message, dict.Blank) {
					resp.Message = StatusText(resp.Status)
				}
			}
		}
		resp.WriteJSON(nil)
	} else {
		if strings.EqualFold(StatusText(resp.Status), dict.Blank) {
			resp.Status = StatusUnknown
		}
		if strings.EqualFold(resp.Message, dict.Blank) {
			resp.Message = StatusText(resp.Status)
		}
	}
}

// // ParseResponse :
// func ParseResponse(response http.Response) (*Response, error) {
// 	resp := new(Response)
// 	data, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if err := json.Unmarshal(data, resp); err != nil {
// 		return nil, err
// 	}
// 	return resp, nil
// }
