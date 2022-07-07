resource "citrixadc_lbgroup" "tf_lbgroup" {
	name = "tf_lbgroup"
	persistencetype = "COOKIEINSERT"
	persistencebackup = "SOURCEIP"
	backuppersistencetimeout = 15.0
	persistmask = "255.255.254.0"
	cookiename = "tf_cookie_1"
	v6persistmasklen = 96
	timeout = 15.0
}