package config

import(
	"os"
	"fmt"
)

const ProductPageUrl = "https://catalog.app.iherb.com/product/"


//default headers values.  Could be overwritten in .env file
var (
	HttpDefaultAcceptLanguageHeader = "en-US,en;q=0.8"
	HttpDefaultPlatformHeader       = "Linux"
	HttpDefaultRegionTypeHeader     = "GLOBAL"
	HttpDefaultIhPrefHeader         = "lc=en-US;cc=USD;ctc=AM;wp=kilograms"
	HttpDefaultPrefHeader           = "{\"ctc\":\"AM\",\"crc\":\"USD\",\"crs\":\"2\",\"lac\":\"en-US\",\"storeid\":0,\"som\":\"kilograms\"}"
	HttpDefaultUserAgentHeader      = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36"
	HttpDefaultContentTypeHeader    = "application/json; charset=UTF-8"
)


func LoadConfig(){
	if platform := os.Getenv("PLATFORM"); platform!= ""{
		HttpDefaultPlatformHeader = platform
	}
	if userAgent := os.Getenv("USER_AGENT"); userAgent!= ""{
		HttpDefaultUserAgentHeader = userAgent
	}
	setPreferences()
}

func setPreferences(){
	lang := os.Getenv("CATALOG_LANGUAGE")
	currency := os.Getenv("CURRENCY")
	country := os.Getenv("COUNTRY")
	weightUnits := os.Getenv("COUNTRY")
	if lang!="" && currency!="" && country!="" && weightUnits!=""{
		HttpDefaultIhPrefHeader = fmt.Sprintf("lc=%s;cc=%s;ctc=%s;wp=%s",lang,currency,country,weightUnits)
		HttpDefaultPrefHeader = fmt.Sprintf("{\"ctc\":\"%s\",\"crc\":\"%s\",\"crs\":\"2\",\"lac\":\"%s\",\"storeid\":0,\"som\":\"%s\"}",country,currency,lang,weightUnits)
	}
}