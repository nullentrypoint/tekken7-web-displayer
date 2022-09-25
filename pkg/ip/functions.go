package ip

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
)

func GetInfo(ip string) (Response, error) {

	netIP := net.ParseIP(ip)
	if netIP.IsPrivate() {
		return Response{}, fmt.Errorf("private ip: %v", ip)
	}

	requestURL := baseUrl + ip
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return Response{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Response{}, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Response{}, err
	}

	var response Response
	if err := unmarshal(resBody, &response); err != nil {
		return Response{}, err
	}

	return response, nil
}

func Itoa(ipInt int64) string {
	// need to do two bit shifting and “0xff” masking
	b0 := strconv.FormatInt((ipInt>>24)&0xff, 10)
	b1 := strconv.FormatInt((ipInt>>16)&0xff, 10)
	b2 := strconv.FormatInt((ipInt>>8)&0xff, 10)
	b3 := strconv.FormatInt((ipInt & 0xff), 10)
	return b0 + "." + b1 + "." + b2 + "." + b3
}

func unmarshal(data []byte, ptr *Response) error {
	if err := xml.Unmarshal(data, ptr); err != nil {
		return err
	}

	return nil
}

func (r Response) String() string {
	out := r.Country
	if r.RegionName != "" {
		out += ", " + r.RegionName
	}

	if r.City != "" && r.City != r.RegionName {
		out += ", " + r.City
	}

	return out
} 

