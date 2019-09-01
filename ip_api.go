package ip_api

//URI for the free IP-API
const FreeAPIURI = "http://ip-api.com/json/"

//URI for the pro IP-API
const ProAPIURI = "http://pro.ip-api.com/json/"

//All supported fields for IP-API
var APIFields = []string {"status","message","continent","continentCode","country","countryCode","region","regionName","city","district","zip","lat","lon","timezone","isp","org","as","asname","reverse","mobile","proxy"}

//The default fields for IP-API
var APIDefaultFields = []string {"status","message","country","countryCode","region","regionName","city","zip","lat","lon","timezone","isp","org","as"}

var APIKey string