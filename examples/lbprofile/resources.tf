resource "citrixadc_lbprofile" "tf_lbprofile" {
    lbprofilename = "tf_lbprofile"
    dbslb = "ENABLED"
	processlocal = "DISABLED"
	httponlycookieflag = "ENABLED"
	lbhashfingers = 258
	lbhashalgorithm = "PRAC"
	storemqttclientidandusername = "YES"
}
