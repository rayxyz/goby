package httpx

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// ParseResponseBody :
func ParseResponseBody(body io.Reader) (*Response, error) {
	buf, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	response := new(Response)
	if err = json.Unmarshal(buf, response); err != nil {
		return nil, err
	}
	return response, nil
}
