package ip_api

//URI for the free IP-API
const FreeAPIURI = "http://ip-api.com/json/"

//URI for the pro IP-API
const ProAPIURI = "http://pro.ip-api.com/json/"

type Location struct {
	Status 			string	`json:"status,omitempty"`
	Message			string	`json:"message,omitempty"`
	Continent		string	`json:"continent,omitempty"`
	ContinentCode	string	`json:"continentCode,omitempty"`
	Country			string	`json:"country,omitempty"`
	CountryCode		string	`json:"countryCode,omitempty"`
	Region			string	`json:"region,omitempty"`
	RegionName		string	`json:"regionName,omitempty"`
	City			string	`json:"city,omitempty"`
	District		string	`json:"district,omitempty"`
	ZIP				string	`json:"zip,omitempty"`
	Lat				float32	`json:"lat,omitempty"`
	Lon				float32	`json:"lon,omitempty"`
	Timezone		string	`json:"timezone,omitempty"`
	Currency		string	`json:"currency,omitempty"`
	ISP				string	`json:"isp,omitempty"`
	Org				string	`json:"org,omitempty"`
	AS				string	`json:"as,omitempty"`
	ASName			string	`json:"asame,omitempty"`
	Reverse			string	`json:"reverse,omitempty"`
	Mobile			bool	`json:"mobile,omitempty"`
	Proxy			bool	`json:"proxy,omitempty"`
	Query			string	`json:"query,omitempty"`
}