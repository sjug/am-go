package soap

import (
	"bytes"
	"encoding/xml"
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
	GetCollectorProfiles GetCollectorProfile `Collectorxml:"typ:GetCollectorProfileRequest"`
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
	r, _ := regexp.Compile(`<Collector[\s\S]+>`)
	newResp := r.FindString(resp)
	var c CollectorResponse
	xml.Unmarshal([]byte(newResp), &c)
	tempCollector := structure.CollectorTier{CollectorTier: c.Tier}
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
	client := &http.Client{}
	body := nopCloser{bytes.NewBufferString(string(xmlstring))}
	req, err := http.NewRequest("PUT", "http://example.com/someresource", body)
	if err != nil {
		return "error", err
	}
	req.Header.Add("Accept", "application/xml")
	req.Header.Add("Content-Type", "application/xml")
	//req.ContentLength = int64(len(string(msgbody)))

	resp, resperr := client.Do(req)
	if err != nil {
		return "error", resperr
	}

	strResp, _ := ioutil.ReadAll(resp.Body)
	//resp := `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><GetCollectorProfileResponse ETSServiceVersion="collector-2.12.17-20150306" xmlns="http://ws.loyalty.com/tp/ets/2008/04/01/collector/types" xmlns:ns2="http://ws.loyalty.com/tp/ets/2008/02/01/common" xmlns:ns3="http://ws.loyalty.com/tp/ets/2008/04/01/collector-common" xmlns:ns4="http://ws.loyalty.com/tp/ets/2008/04/01/collector" xmlns:ns5="http://ws.loyalty.com/tp/ets/2008/02/01/email" xmlns:ns6="http://ws.loyalty.com/tp/ets/2008/02/01/account"><ns4:MosaikTier SegmentPrefix="BMPH" SegmentSuffix="5C"/><ns4:MosaikTier SegmentPrefix="BMWJ" SegmentSuffix="5C"/><ns4:Person DateOfBirth="1969-11-30-05:00" FirstName="LMGCANSAVE" Gender="F" LastName="L3R3Z5" Prefix="MS" Suffix=" "><ns4:HomeAddress City="NORTH YORK" Country="CAN" PostalCode="M2P2B7" Province="ON" Status="0" StreetAddress1="4110 YONGE STREET" StreetAddress2="SUITE 200"/><ns4:HomePhone>4169804860</ns4:HomePhone><ns4:BusinessPhone>4169804867</ns4:BusinessPhone></ns4:Person></Collector><Balance Amount="31980708" LastMaintenanceTime="2014-09-12T00:28:56-04:00"/><CashMilesBalance Amount="3902" LastMaintenanceTime="2011-07-13T10:54:11-04:00"/></GetCollectorProfileResponse></soap:Body></soap:Envelope><Collector AMCashEligible="true" AMCashRegion="2" AccountStatus="A" AccountTier="G" AccountType="I" AddressType="H" CollectorNumber="50000131287" EnrollSourceCodeLevel1="LMGCAN" EnrollSourceCodeLevel2="SAVE" HouseholdSize="4" IncomeLevel="4" LanguageCode="fr-CA" MailProfile="0" NumberOfCards="2" StatementsPerYear=" " xsi:type="ns4:ConsumerCollectorType" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">`
	return string(strResp), nil
}
