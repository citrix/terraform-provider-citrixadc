resource "citrixadc_policyhttpcallout" "tf_policyhttpcallout" {
	name = "tf_policyhttpcallout"
	bodyexpr = "client.ip.src"
	cacheforsecs = 5
	comment = "Demo comment"
	headers = ["cip(client.ip.src)", "hdr(http.req.header(\"HDR\"))"]
	hostexpr = "http.req.header(\"Host\")"
	httpmethod = "GET"
	parameters = ["param1(\"name1\")", "param2(http.req.header(\"hdr\"))"]
	resultexpr = "http.res.body(10000).length"
	returntype = "TEXT"
	scheme = "http"
	vserver = citrixadc_lbvserver.tf_lbvserver.name
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
	name = "tf_lbvserver"
	ipv46 = "10.202.11.11"
	port = 80
	servicetype = "HTTP"
}