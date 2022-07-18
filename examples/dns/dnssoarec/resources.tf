resource "citrixadc_dnssoarec" "tf_dnssoarec" {
	domain =  "hello.com"
	originserver  = "10.2.3.5"
	contact =  "other"
	expire = 1800
	refresh = 4000
}
