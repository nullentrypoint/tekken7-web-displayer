package ip

import (
	"log"
	"testing"
)

const example1 = `<query>
<status>success</status>
<country>Germany</country>
<countryCode>DE</countryCode>
<region>HE</region>
<regionName>Hesse</regionName>
<city>Frankfurt am Main</city>
<zip>60313</zip>
<lat>50.1109</lat>
<lon>8.68213</lon>
<timezone>Europe/Berlin</timezone>
<isp>Google LLC</isp>
<org>Google LLC</org>
<as>AS15169 Google LLC</as>
<query>142.250.184.206</query>
</query>`

func TestUnmarshal(t *testing.T) {
	var response Response
	if err := unmarshal([]byte(example1), &response); err != nil {
		t.Fatal(err)
	}
	log.Println(response)
}