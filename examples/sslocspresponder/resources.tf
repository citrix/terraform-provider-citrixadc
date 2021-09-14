resource "citrixadc_sslocspresponder" "tf_sslocspresponder" {
		name = "tf_sslocspresponder"
		url = "http://www.citrix.com"
		batchingdelay = 5
		batchingdepth = 2
		cache = "ENABLED"
		cachetimeout = 1
		httpmethod = "GET"
		insertclientcert = "YES"
		ocspurlresolvetimeout = 100
		producedattimeskew = 300
		resptimeout = 100
		trustresponder = false
		usenonce = "NO"
	}