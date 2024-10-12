package ip_api

import (
	"encoding/json"
	"testing"
)

const allFieldCsv string = "status,message,continent,continentCode,country,countryCode,region,regionName,city,district,zip,lat,lon,timezone,offset,currency,isp,org,as,asname,reverse,mobile,proxy,query,hosting"

var googleAllFieldsDefaultLangRequestPointers = struct {
	lat     float32
	lon     float32
	offset  int
	mobile  bool
	proxy   bool
	hosting bool
}{
	lat:     39.03,
	lon:     -77.5,
	offset:  -14400,
	mobile:  false,
	proxy:   false,
	hosting: true,
}

var googleAllFieldsDefaultLangRequestTest = struct {
	query struct {
		name  string
		query Query
	}
	out struct {
		Location
		err error
	}
}{
	query: struct {
		name  string
		query Query
	}{
		name: "googleAllFieldsDefaultLang",
		query: Query{
			Queries: []QueryIP{
				{Query: "8.8.8.8"},
			},
			Fields: allFieldCsv,
			Lang:   "",
		},
	},
	out: struct {
		Location
		err error
	}{Location: Location{
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
		Lat:           &googleAllFieldsDefaultLangRequestPointers.lat,
		Lon:           &googleAllFieldsDefaultLangRequestPointers.lon,
		Timezone:      "America/New_York",
		Offset:        &googleAllFieldsDefaultLangRequestPointers.offset,
		Currency:      "USD",
		ISP:           "Google LLC",
		Org:           "Google Public DNS",
		AS:            "AS15169 Google LLC",
		ASName:        "GOOGLE",
		Reverse:       "dns.google",
		Mobile:        &googleAllFieldsDefaultLangRequestPointers.mobile,
		Proxy:         &googleAllFieldsDefaultLangRequestPointers.proxy,
		Hosting:       &googleAllFieldsDefaultLangRequestPointers.hosting,
		Query:         "8.8.8.8",
	},
		err: nil,
	},
}

var cloudflareAllFieldsDefaultLangRequestPointers = struct {
	lat     float32
	lon     float32
	offset  int
	mobile  bool
	proxy   bool
	hosting bool
}{
	lat:     -27.4766,
	lon:     153.0166,
	offset:  36000,
	mobile:  false,
	proxy:   false,
	hosting: true,
}

var cloudflareAllFieldsDefaultLangRequestTest = struct {
	query struct {
		name  string
		query Query
	}
	out struct {
		Location
		err error
	}
}{
	query: struct {
		name  string
		query Query
	}{
		name: "cloudflareAllFieldsDefaultLang",
		query: Query{
			Queries: []QueryIP{
				{Query: "1.1.1.1"},
			},
			Fields: allFieldCsv,
			Lang:   "",
		},
	},
	out: struct {
		Location
		err error
	}{Location: Location{
		Status:        "success",
		Message:       "",
		Continent:     "Oceania",
		ContinentCode: "OC",
		Country:       "Australia",
		CountryCode:   "AU",
		Region:        "QLD",
		RegionName:    "Queensland",
		City:          "South Brisbane",
		District:      "",
		ZIP:           "4101",
		Lat:           &cloudflareAllFieldsDefaultLangRequestPointers.lat,
		Lon:           &cloudflareAllFieldsDefaultLangRequestPointers.lon,
		Timezone:      "Australia/Brisbane",
		Offset:        &cloudflareAllFieldsDefaultLangRequestPointers.offset,
		Currency:      "AUD",
		ISP:           "Cloudflare, Inc",
		Org:           "APNIC and Cloudflare DNS Resolver project",
		AS:            "AS13335 Cloudflare, Inc.",
		ASName:        "CLOUDFLARENET",
		Reverse:       "one.one.one.one",
		Mobile:        &cloudflareAllFieldsDefaultLangRequestPointers.mobile,
		Proxy:         &cloudflareAllFieldsDefaultLangRequestPointers.proxy,
		Hosting:       &cloudflareAllFieldsDefaultLangRequestPointers.hosting,
		Query:         "1.1.1.1",
	},
		err: nil,
	},
}

var googleSpecificFieldRuLangRequestPointers = struct {
	lat float32
}{
	lat: 39.03,
}

var googleSpecificFieldRuLangRequestTest = struct {
	query struct {
		name  string
		query Query
	}
	out struct {
		Location
		err error
	}
}{
	query: struct {
		name  string
		query Query
	}{
		name: "googleSpecificFieldRuLang",
		query: Query{
			Queries: []QueryIP{
				{Query: "8.8.4.4"},
			},
			Fields: "status,message,continent,country,region,city,zip,lat",
			Lang:   "ru",
		},
	},
	out: struct {
		Location
		err error
	}{Location: Location{
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
		Lat:           &googleSpecificFieldRuLangRequestPointers.lat,
		Lon:           nil,
		Timezone:      "",
		Offset:        nil,
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
	},
		err: nil,
	},
}

var singleQuery = googleAllFieldsDefaultLangRequestTest

func TestSingleQuery(t *testing.T) {
	t.Run(singleQuery.query.name, func(t *testing.T) {
		location, err := SingleQuery(singleQuery.query.query, "", "", true)

		if err != nil && (err.Error() != singleQuery.out.err.Error()) {
			t.Errorf("got: %#v, want: %#v", err, singleQuery.out.err)
		}

		jsonLocation, _ := json.Marshal(location)
		jsonTestLocation, _ := json.Marshal(singleQuery.out.Location)

		if string(jsonLocation) != string(jsonTestLocation) {
			t.Errorf("got: %#v, want: %#v", string(jsonLocation), string(jsonTestLocation))
		}
	})
}

var batchQuery = struct {
	query Query
	out   struct {
		locations []Location
		err       error
	}
}{
	query: Query{
		Queries: []QueryIP{
			{Query: googleAllFieldsDefaultLangRequestTest.query.query.Queries[0].Query},
			{Query: cloudflareAllFieldsDefaultLangRequestTest.query.query.Queries[0].Query},
			{Query: googleSpecificFieldRuLangRequestTest.query.query.Queries[0].Query, Fields: googleSpecificFieldRuLangRequestTest.query.query.Fields, Lang: googleSpecificFieldRuLangRequestTest.query.query.Lang},
		},
		Fields: allFieldCsv,
		Lang:   "",
	},
	out: struct {
		locations []Location
		err       error
	}{locations: []Location{
		googleAllFieldsDefaultLangRequestTest.out.Location,
		cloudflareAllFieldsDefaultLangRequestTest.out.Location,
		googleSpecificFieldRuLangRequestTest.out.Location,
	}, err: nil},
}

func TestBatchQuery(t *testing.T) {
	t.Run("batchQuery", func(t *testing.T) {
		locations, err := BatchQuery(batchQuery.query, "", "", true)

		if err != nil && (err.Error() != batchQuery.out.err.Error()) {
			t.Errorf("got: %#v, want: %#v", err, batchQuery.out.err)
		}

		jsonLocations, _ := json.Marshal(locations)

		// This is somewhat of a hack
		//   batch doesn't support reverse field, and to reduce duplicate test data
		//   we just make the test data for this field an empty string for testing
		for i := range batchQuery.out.locations {
			batchQuery.out.locations[i].Reverse = ""
		}

		jsonTestLocations, _ := json.Marshal(batchQuery.out.locations)

		if string(jsonLocations) != string(jsonTestLocations) {
			t.Errorf("got: %#v, want: %#v", string(jsonLocations), string(jsonTestLocations))
		}
	})
}

var buildFieldListTests = []struct {
	in  string
	out string
}{
	{
		in:  allFieldCsv,
		out: "fields=" + allFieldCsv,
	},
}

func TestBuildFieldList(t *testing.T) {
	for _, tt := range buildFieldListTests {
		t.Run(tt.in, func(t *testing.T) {
			s := buildFieldList(tt.in)

			if s != tt.out {
				t.Errorf("got %q, want %q", s, tt.out)
			}
		})
	}
}
