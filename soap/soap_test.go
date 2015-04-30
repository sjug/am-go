package soap

import (
	"encoding/xml"
	"regexp"
	"testing"

	"github.com/sjug/am-go/structure"
	. "github.com/smartystreets/goconvey/convey"
)

func TestEncodeRequestMatch(t *testing.T) {
	Convey("Given some TierRequest struct", t, func() {
		expectedRequest := &TierRequest{
			NsEnv:  "http://schemas.xmlsoap.org/soap/envelope/",
			NsType: "http://ws.loyalty.com/tp/ets/2008/04/01/collector/types",
			TierBodys: TierBody{
				GetCollectorProfiles: GetCollectorProfile{
					Contexts: CollectorContext{
						Channel:  "WEB",
						Source:   "WEB",
						Language: "en-CA"},
					Number: 5,
				}}}
		So(expectedRequest, ShouldHaveSameTypeAs, &TierRequest{})

		Convey("When the struct is marshaled to XML", func() {
			xmlstring, _ := xml.MarshalIndent(expectedRequest, "", "  ")
			xmlstring = []byte(xml.Header + string(xmlstring))
			So(xmlstring, ShouldHaveSameTypeAs, []byte("test"))

			Convey("The marshalled XML byte slice should match the one returned from the encodeRequest func", func() {
				realXML, _ := encodeRequest(5)
				So(realXML, ShouldResemble, xmlstring)
			})
		})
	})

}

func TestRegexMatch(t *testing.T) {
	Convey("Given a SOAP response", t, func() {
		fullReponse := `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
    <soap:Body>
      <GetCollectorProfileResponse
        xmlns="http://ws.loyalty.com/tp/ets/2008/04/01/collector/types"
        xmlns:ns2="http://ws.loyalty.com/tp/ets/2008/02/01/common"
        xmlns:ns3="http://ws.loyalty.com/tp/ets/2008/04/01/collector-common"
        xmlns:ns4="http://ws.loyalty.com/tp/ets/2008/04/01/collector"
        xmlns:ns5="http://ws.loyalty.com/tp/ets/2008/02/01/email"
        xmlns:ns6="http://ws.loyalty.com/tp/ets/2008/02/01/account"
        ETSServiceVersion="collector-2.12.17-20150306">
        <Collector
          xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
          AMCashEligible="true"
          AMCashRegion="2"
          AccountStatus="A"
          AccountTier="B"
          AccountType="I"
          AddressType="H"
          CollectorNumber="50001366830"
          EnrollSourceCodeLevel1="PREA01"
          EnrollSourceCodeLevel2="TEST"
          LanguageCode="fr-CA"
          MailProfile="0"
          NumberOfCards="1"
          xsi:type="ns4:ConsumerCollectorType">
          <ns4:MosaikTier SegmentPrefix="BMPH" SegmentSuffix="3E" />
          <ns4:Person
            DateOfBirth="1983-02-22-05:00"
            FirstName="ANGELIQUE"
            Gender="F"
            LastName="LAURENT"
            Prefix="MS" Suffix=" ">
            <ns4:HomeAddress
              City="TORONTO"
              Country="CAN"
              PostalCode="M5G2L1"
              Province="ON"
              Status="0"
              StreetAddress1="600-438 UNIVERSITY AVE" />
            <ns4:HomePhone>4165522367</ns4:HomePhone>
            <ns4:BusinessPhone>4165522367</ns4:BusinessPhone>
          </ns4:Person>
        </Collector>
        <Balance Amount="97695" LastMaintenanceTime="2014-03-11T19:50:32-04:00" />
        <CashMilesBalance Amount="0" LastMaintenanceTime="1969-12-31T19:00:00-05:00" />
        <ContactDetails
          ChangedTime="2013-03-06T12:43:20.600-05:00"
          Channel="MOBAPP"
          CollectorNumber="50001366830"
          ContactType="EMAIL"
          EffectiveStartTime="2013-03-06T12:43:20.572-05:00"
          EnableOptions="true"
          Format="T"
          FromSecureSource="true"
          Source="MOBLOYAPP"
          Status="V"
          Value="hting@loyalty.com"
          Verified="false" />
      </GetCollectorProfileResponse>
    </soap:Body>
  </soap:Envelope>`
		So(fullReponse, ShouldHaveSameTypeAs, "text")

		Convey("Regular expression parsing output", func() {
			r, _ := regexp.Compile(`<Collector[\s\S]*?">`)
			mockResponse := r.FindString(fullReponse)
			So(fullReponse, ShouldContainSubstring, mockResponse)

			Convey("Function regexResponse should match parsed string", func() {
				actualResponse := regexResponse(fullReponse)
				So(actualResponse, ShouldEqual, mockResponse)
			})
		})
	})

}

func TestParseXML(t *testing.T) {
	Convey("Given a collector response", t, func() {
		resp := `<Collector
      xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
      AMCashEligible="true"
      AMCashRegion="2"
      AccountStatus="A"
      AccountTier="B"
      AccountType="I"
      AddressType="H"
      CollectorNumber="50001366830"
      EnrollSourceCodeLevel1="PREA01"
      EnrollSourceCodeLevel2="TEST"
      LanguageCode="fr-CA"
      MailProfile="0"
      NumberOfCards="1"
      xsi:type="ns4:ConsumerCollectorType">
      <ns4:MosaikTier SegmentPrefix="BMPH" SegmentSuffix="3E" />
      <ns4:Person
        DateOfBirth="1983-02-22-05:00"
        FirstName="ANGELIQUE"
        Gender="F"
        LastName="LAURENT"
        Prefix="MS" Suffix=" ">
        <ns4:HomeAddress
          City="TORONTO"
          Country="CAN"
          PostalCode="M5G2L1"
          Province="ON"
          Status="0"
          StreetAddress1="600-438 UNIVERSITY AVE" />
        <ns4:HomePhone>4165522367</ns4:HomePhone>
        <ns4:BusinessPhone>4165522367</ns4:BusinessPhone>
      </ns4:Person>
    </Collector>`
		So(resp, ShouldHaveSameTypeAs, "text")

		Convey("Response object should contain unmarshaled XML data", func() {
			var responseObject CollectorResponse
			xml.Unmarshal([]byte(resp), &responseObject)
			So(responseObject.Tier, ShouldNotBeBlank)

			Convey("Collector Tier object should match that returned by parseXML function", func() {
				mockCollectorTier := structure.CollectorTier{CollectorTier: responseObject.Tier}
				realCollectorTier, _ := parseXML(resp)
				So(realCollectorTier, ShouldResemble, &mockCollectorTier)
			})

		})
	})

}
