package ip_api

//URI for the free IP-API
const FREE_API_URI = "http://ip-api.com/json/"

//URI for the pro IP-API
const PRO_API_URI = "http://pro.ip-api.com/json/"

//All supported fields for IP-API
var API_FIELDS = []string {"status","message","continent","continentCode","country","countryCode","region","regionName","city","district","zip","lat","lon","timezone","isp","org","as","asname","reverse","mobile","proxy"}

//The default fields for IP-API
var API_FIELD_DEFAULTS = []string {"status","message","country","countryCode","region","regionName","city","zip","lat","lon","timezone","isp","org","as"}

var API_KEY string