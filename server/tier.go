package server

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"regexp"

	"github.com/dimfeld/httptreemux"
)

func callTier() {
	resp := `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><GetCollectorProfileResponse ETSServiceVersion="collector-2.12.17-20150306" xmlns="http://ws.loyalty.com/tp/ets/2008/04/01/collector/types" xmlns:ns2="http://ws.loyalty.com/tp/ets/2008/02/01/common" xmlns:ns3="http://ws.loyalty.com/tp/ets/2008/04/01/collector-common" xmlns:ns4="http://ws.loyalty.com/tp/ets/2008/04/01/collector" xmlns:ns5="http://ws.loyalty.com/tp/ets/2008/02/01/email" xmlns:ns6="http://ws.loyalty.com/tp/ets/2008/02/01/account"><ns4:MosaikTier SegmentPrefix="BMPH" SegmentSuffix="5C"/><ns4:MosaikTier SegmentPrefix="BMWJ" SegmentSuffix="5C"/><ns4:Person DateOfBirth="1969-11-30-05:00" FirstName="LMGCANSAVE" Gender="F" LastName="L3R3Z5" Prefix="MS" Suffix=" "><ns4:HomeAddress City="NORTH YORK" Country="CAN" PostalCode="M2P2B7" Province="ON" Status="0" StreetAddress1="4110 YONGE STREET" StreetAddress2="SUITE 200"/><ns4:HomePhone>4169804860</ns4:HomePhone><ns4:BusinessPhone>4169804867</ns4:BusinessPhone></ns4:Person></Collector><Balance Amount="31980708" LastMaintenanceTime="2014-09-12T00:28:56-04:00"/><CashMilesBalance Amount="3902" LastMaintenanceTime="2011-07-13T10:54:11-04:00"/></GetCollectorProfileResponse></soap:Body></soap:Envelope><Collector AMCashEligible="true" AMCashRegion="2" AccountStatus="A" AccountTier="G" AccountType="I" AddressType="H" CollectorNumber="50000131287" EnrollSourceCodeLevel1="LMGCAN" EnrollSourceCodeLevel2="SAVE" HouseholdSize="4" IncomeLevel="4" LanguageCode="fr-CA" MailProfile="0" NumberOfCards="2" StatementsPerYear=" " xsi:type="ns4:ConsumerCollectorType" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">`
	r, err := regexp.Compile(`<Collector[\s\S]+>`)
	fmt.Println(r.MatchString(resp))

}

func userTierHandler(w http.ResponseWriter, r *http.Request, ps map[string]string) {
	// TODO userTierHandler

	fmt.Fprintf(w, "hello, %s!\n", ps["num"])
}

// InitTier func sets up routing for tier path
func InitTier(router *httptreemux.TreeMux) {
	router.GET("/user/tier/:num", userTierHandler)
}
