package ip_api

import (
	"encoding/json"
	"log"
	"testing"
)

func getSuccessfulSingleResponse() string {
	var lat float32 = 40.7357
	var lon float32 = -74.1724
	var mobile = false
	var proxy = false
	var hosting = true

	location := Location{
		Status:        "success",
		Message:       "",
		Continent:     "North America",
		ContinentCode: "NA",
		Country:       "United States",
		CountryCode:   "US",
		Region:        "NJ",
		RegionName:    "New Jersey",
		City:          "Newark",
		District:      "",
		ZIP:           "07175",
		Lat:           &lat,
		Lon:           &lon,
		Timezone:      "America/New_York",
		Currency:      "",
		ISP:           "Google LLC",
		Org:           "Level 3",
		AS:            "AS15169 Google LLC",
		ASName:        "GOOGLE",
		Reverse:       "dns.google",
		Mobile:        &mobile,
		Proxy:         &proxy,
		Hosting:       &hosting,
		Query:         "8.8.8.8",
	}
	
	result, _ := json.Marshal(location)
	
	return string(result)
}

func getSuccessfulBatchResponse() string {
	var lat1 float32 = 40.7357
	var lon1 float32 = -74.1724
	var mobile1 = false
	var proxy1 = false
	var hosting1 = true

	location1 := Location{
		Status:        "success",
		Message:       "",
		Continent:     "North America",
		ContinentCode: "NA",
		Country:       "United States",
		CountryCode:   "US",
		Region:        "NJ",
		RegionName:    "New Jersey",
		City:          "Newark",
		District:      "",
		ZIP:           "07175",
		Lat:           &lat1,
		Lon:           &lon1,
		Timezone:      "America/New_York",
		Currency:      "",
		ISP:           "Google LLC",
		Org:           "Level 3",
		AS:            "AS15169 Google LLC",
		ASName:        "GOOGLE",
		Reverse:       "",
		Mobile:        &mobile1,
		Proxy:         &proxy1,
		Hosting:       &hosting1,
		Query:         "8.8.8.8",
	}

	var lat2 float32 = -33.8688
	var lon2 float32 = 151.209
	var mobile2 = false
	var proxy2 = false
	var hosting2 = true
	
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
		Lat:           &lat2,
		Lon:           &lon2,
		Timezone:      "Australia/Sydney",
		Currency:      "",
		ISP:           "Cloudflare, Inc.",
		Org:           "",
		AS:            "AS13335 Cloudflare, Inc.",
		ASName:        "CLOUDFLARENET",
		Reverse:       "",
		Mobile:        &mobile2,
		Proxy:         &proxy2,
		Hosting:       &hosting2,
		Query:         "1.1.1.1",
	}

	var lat3 float32 = 39.0438
	
	location3 := Location{
		Status:        "success",
		Message:       "",
		Continent:     "Северная Америка",
		ContinentCode: "",
		Country:       "США",
		CountryCode:   "",
		Region:        "VA",
		RegionName:    "",
		City:          "Ашберн",
		District:      "",
		ZIP:           "20149",
		Lat:           &lat3,
		Lon:           nil,
		Timezone:      "",
		Currency:      "",
		ISP:           "",
		Org:           "",
		AS:            "",
		ASName:        "",
		Reverse:       "",
		Mobile:        nil,
		Proxy:         nil,
		Hosting:       nil,
		Query:         "",
	}

	locations := []Location{location1,location2,location3}

	result, _ := json.Marshal(locations)

	return string(result)
}

func getSuccessfulFieldListString() string {
	return "fields=status,message,continent,continentCode,country,countryCode,region,regionName,city,district,zip,lat,lon,timezone,isp,org,as,asname,reverse,mobile,proxy,query,hosting"
}

func TestSingleQuery(t *testing.T) {
	var singleQuery = Query{
		Queries: []QueryIP{
			{Query:"8.8.8.8"},
		},
		Fields:  "status,message,continent,continentCode,country,countryCode,region,regionName,city,district,zip,lat,lon,timezone,isp,org,as,asname,reverse,mobile,proxy,query,hosting",
		Lang:    "",
	}

	var location *Location

	location, err := SingleQuery(singleQuery,"","",true)

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
			{Query:"8.8.4.4",Fields:"status,message,continent,country,region,city,zip,lat",Lang:"ru"},
		},
		Fields:  "status,message,continent,continentCode,country,countryCode,region,regionName,city,district,zip,lat,lon,timezone,isp,org,as,asname,reverse,mobile,proxy,query,hosting",
		Lang:    "",
	}

	var locations []Location

	locations, err := BatchQuery(batchQuery,"","",true)

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
	fieldsList := "status,message,continent,continentCode,country,countryCode,region,regionName,city,district,zip,lat,lon,timezone,isp,org,as,asname,reverse,mobile,proxy,query,hosting"

	fieldListString := buildFieldList(fieldsList)

	if fieldListString != getSuccessfulFieldListString() {
		t.Error("Field list does not match")
	}
}