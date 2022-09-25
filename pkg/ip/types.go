package ip

import "encoding/xml"

const baseUrl = "http://ip-api.com/xml/"

type Response struct {
	XMLName     xml.Name `xml:"query"`
	Text        string   `xml:",chardata"`
	Status      string   `xml:"status"`
	Country     string   `xml:"country"`
	CountryCode string   `xml:"countryCode"`
	Region      string   `xml:"region"`
	RegionName  string   `xml:"regionName"`
	City        string   `xml:"city"`
	Zip         string   `xml:"zip"`
	Lat         string   `xml:"lat"`
	Lon         string   `xml:"lon"`
	Timezone    string   `xml:"timezone"`
	Isp         string   `xml:"isp"`
	Org         string   `xml:"org"`
	As          string   `xml:"as"`
	Query       string   `xml:"query"`
}