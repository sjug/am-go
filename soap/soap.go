package soap

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/sjug/am-go/structure"
)

// Generic Reader
type nopCloser struct {
	io.Reader
}

// CollectorResponse struct is the collector response XML structure
type CollectorResponse struct {
	XMLName      xml.Name `xml:"Collector"`
	CashEligible bool     `xml:"AMCashEligible,attr"`
	CashRegion   int      `xml:"AMCashRegion,attr"`
	Status       string   `xml:"AccountStatus,attr"`
	Tier         string   `xml:"AccountTier,attr"`
	Type         string   `xml:"AccountType,attr"`
	Number       string   `xml:"CollectorNumber,attr"`
	SourceLvl    string   `xml:"EnrollSourceCodeLevel1,attr"`
	SourceLvl2   string   `xml:"EnrollSourceCodeLevel2,attr"`
	HouseSize    int      `xml:"HouseholdSize,attr"`
	IncomeLvl    int      `xml:"IncomeLevel,attr"`
	Language     string   `xml:"LanguageCode,attr"`
	Mail         int      `xml:"MailProfile,attr"`
	Cards        int      `xml:"NumberOfCards,attr"`
	Statemnts    string   `xml:"StatementsPerYear,attr"`
	XsiType      string   `xml:"xsi:type,attr"`
	Xmlns        string   `xml:"xmlns:xsi,attr"`
}

// TierRequest is the outer most XML envelope of soap request
type TierRequest struct {
	XMLName   xml.Name `xml:"soapenv:Envelope"`
	NsEnv     string   `xml:"xmlns:soapenv,attr"`
	NsType    string   `xml:"xmlns:typ,attr"`
	Header    string   `xml:"soapenv:Header"`
	TierBodys TierBody `xml:"soapenv:Body"`
}

// TierBody is an emtpy container with the GetCollectorProfile struct
type TierBody struct {
	GetCollectorProfiles GetCollectorProfile `xml:"typ:GetCollectorProfileRequest"`
}

// GetCollectorProfile struct has the context and collector number
type GetCollectorProfile struct {
	Contexts CollectorContext `xml:"typ:Context"`
	Number   int              `xml:"typ:CollectorNumber"`
}

// CollectorContext contanins a few variables as attributes
type CollectorContext struct {
	Channel  string `xml:"Channel,attr"`
	Source   string `xml:"Source,attr"`
	Language string `xml:"LanguageCode,attr"`
}

// GetTierFromSoap function calls tier soap service and gets parsed for tier info
func GetTierFromSoap(number int) (*structure.CollectorTier, error) {
	resp, _ := callSoap(number)
	r, _ := regexp.Compile(`<Collector[\s\S]*?">`)
	newResp := r.FindString(resp)
	var c CollectorResponse
	xml.Unmarshal([]byte(newResp), &c)
	tempCollector := structure.CollectorTier{CollectorTier: c.Tier}
	if c.Tier == "" {
		return &tempCollector, errors.New("Tier missing")
	}
	return &tempCollector, nil
}

func callSoap(number int) (string, error) {
	a := &TierRequest{
		NsEnv:  "http://schemas.xmlsoap.org/soap/envelope/",
		NsType: "http://ws.loyalty.com/tp/ets/2008/04/01/collector/types",
		TierBodys: TierBody{
			GetCollectorProfiles: GetCollectorProfile{
				Contexts: CollectorContext{
					Channel:  "WEB",
					Source:   "WEB",
					Language: "en-CA"},
				Number: number,
			}}}

	xmlstring, err := xml.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error", err
	}
	xmlstring = []byte(xml.Header + string(xmlstring))
	client := http.Client{}
	body := nopCloser{bytes.NewBufferString(string(xmlstring))}
	req, err := http.NewRequest("PUT", "http://www.int-test.one.ets.app.loyalty.com/collector/services/CollectorService", body)
	if err != nil {
		return "Couldn't form http request", err
	}
	req.Header.Add("Accept", "application/xml")
	req.Header.Add("Content-Type", "text/xml;charset=UTF-8")
	req.Header.Add("SOAPAction", "\"getCollectorDetails\"")
	req.ContentLength = int64(len(string(xmlstring)))

	resp, resperr := client.Do(req)
	if err != nil {
		return "HTTP response is broken", resperr
	}

	strResp, _ := ioutil.ReadAll(resp.Body)
	return string(strResp), nil
}
