package ip_api

import (
	"encoding/json"
	"log"
	"testing"
)

func getSuccessfulSingleResponse() string {
	location := Location{
		Status:        "success",
		Message:       "",
		Continent:     "North America",
		ContinentCode: "NA",
		Country:       "United States",
		CountryCode:   "US",
		Region:        "VA",
		RegionName:    "Virginia",
		City:          "Ashburn",
		District:      "",
		ZIP:           "20149",
		Lat:           39.0438,
		Lon:           -77.4874,
		Timezone:      "America/New_York",
		Currency:      "",
		ISP:           "Level 3 Communications",
		Org:           "Google Inc.",
		AS:            "AS15169 Google LLC",
		ASName:        "",
		Reverse:       "dns.google",
		Mobile:        false,
		Proxy:         false,
		Query:         "8.8.8.8",
	}
	
	result, _ := json.Marshal(location)
	
	return string(result)
}

func getSuccessfulBatchResponse() string {
	location1 := Location{
		Status:        "success",
		Message:       "",
		Continent:     "North America",
		ContinentCode: "NA",
		Country:       "United States",
		CountryCode:   "US",
		Region:        "VA",
		RegionName:    "Virginia",
		City:          "Ashburn",
		District:      "",
		ZIP:           "20149",
		Lat:           39.0438,
		Lon:           -77.4874,
		Timezone:      "America/New_York",
		Currency:      "",
		ISP:           "Level 3 Communications",
		Org:           "Google Inc.",
		AS:            "AS15169 Google LLC",
		ASName:        "",
		Reverse:       "",
		Mobile:        false,
		Proxy:         false,
		Query:         "8.8.8.8",
	}
	
	location2 := Location{
		Status:        "success",
		Message:       "",
		Continent:     "Oceania",
		ContinentCode: "OC",
		Country:       "Australia",
		CountryCode:   "AU",
		Region:        "NSW",
		RegionName:    "New South Wales",
		City:          "Sydney",
		District:      "",
		ZIP:           "1001",
		Lat:           -33.8688,
		Lon:           151.209,
		Timezone:      "Australia/Sydney",
		Currency:      "",
		ISP:           "Cloudflare, Inc.",
		Org:           "",
		AS:            "AS13335 Cloudflare, Inc.",
		ASName:        "",
		Reverse:       "",
		Mobile:        false,
		Proxy:         false,
		Query:         "1.1.1.1",
	}

	locations := []Location{location1,location2}

	result, _ := json.Marshal(locations)

	return string(result)
}

func getSuccessfulFieldListString() string {
	return "fields=status,message,continent,continentCode,country,countryCode,region,regionName,city,district,zip,lat,lon,timezone,isp,org,as,asname,reverse,mobile,proxy,query"
}

func TestSingleQuery(t *testing.T) {
	var singleQuery = Query{
		Queries: []QueryIP{
			{Query:"8.8.8.8"},
		},
		Fields:  []string{"status","message","continent","continentCode","country","countryCode","region","regionName","city","district","zip","lat","lon","timezone","isp","org","as","asname","reverse","mobile","proxy","query"},
		Lang:    "",
	}

	var location Location

	location, err := SingleQuery(singleQuery,"")

	if err != nil {
		t.Error(err)
	}

	jsonLocation, _ := json.Marshal(location)

	if string(jsonLocation) != getSuccessfulSingleResponse() {
		log.Println(string(jsonLocation))
		log.Println(getSuccessfulSingleResponse())
		t.Error("Locations did not match")
	}
}


func TestBatchQuery(t *testing.T) {
	var batchQuery = Query{
		Queries: []QueryIP{
			{Query:"8.8.8.8"},
			{Query:"1.1.1.1"},
		},
		Fields:  []string{"status","message","continent","continentCode","country","countryCode","region","regionName","city","district","zip","lat","lon","timezone","isp","org","as","asname","reverse","mobile","proxy","query"},
		Lang:    "",
	}

	var locations []Location

	locations, err := BatchQuery(batchQuery,"")

	if err != nil {
		t.Error(err)
	}

	jsonLocations, _ := json.Marshal(locations)

	if string(jsonLocations) != getSuccessfulBatchResponse() {
		log.Println(string(jsonLocations))
		log.Println(getSuccessfulBatchResponse())
		t.Error("Locations did not match")
	}
}

func TestBuildFieldList(t *testing.T) {
	fieldsList := []string{"status","message","continent","continentCode","country","countryCode","region","regionName","city","district","zip","lat","lon","timezone","isp","org","as","asname","reverse","mobile","proxy","query"}

	fieldListString := buildFieldList(fieldsList)

	if fieldListString != getSuccessfulFieldListString() {
		t.Error("Field list does not match")
	}
}