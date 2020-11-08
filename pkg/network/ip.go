package network

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
)

// IP 2 geo location service
// http://ip-api.com/json/
//

// query => IP or domain
// IP: 104.237.95.14

// English
// {
// 	"as": "AS9312 xTom Hong Kong Limited",
// 	"city": "Sheung Wan",
// 	"country": "Hong Kong",
// 	"countryCode": "HK",
// 	"isp": "Univera Network",
// 	"lat": 22.2874,
// 	"lon": 114.151,
// 	"org": "HostAware HK",
// 	"query": "104.237.95.14",
// 	"region": "HCW",
// 	"regionName": "Central and Western",
// 	"status": "success",
// 	"timezone": "Asia/Hong_Kong",
// 	"zip": "00000"
// }

// Localization
// ex: http://ip-api.com/json/104.237.95.14?lang=zh-CN
// {
// 	"as": "AS9312 xTom Hong Kong Limited",
// 	"city": "Sheung Wan",
// 	"country": "香港",
// 	"countryCode": "HK",
// 	"isp": "Univera Network",
// 	"lat": 22.2874,
// 	"lon": 114.151,
// 	"org": "HostAware HK",
// 	"query": "104.237.95.14",
// 	"region": "HCW",
// 	"regionName": "中西區",
// 	"status": "success",
// 	"timezone": "Asia/Hong_Kong",
// 	"zip": "00000"
// }

// IPGeo :
type IPGeo struct {
	Query       string  `json:"query"`
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float32 `json:"lat"`
	Lon         float32 `json:"lon"`
	Timezone    string  `json:"timezone"`
	ISP         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Message     string  `json:"message"`
}

// GetIPGeo :
func GetIPGeo(query string, localization string) (*IPGeo, error) {
	resp, err := http.Get("http://ip-api.com/json/" + query + "?lang=" + localization)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ipg := new(IPGeo)
	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		if err = json.Unmarshal(body, ipg); err != nil {
			return nil, err
		}
		if !strings.EqualFold(ipg.Status, "success") {
			return nil, errors.New(ipg.Message)
		}

		return ipg, nil
	} else {
		return nil, errors.New("http status: " + resp.Status)
	}
}

// GetOutboundIP : Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
