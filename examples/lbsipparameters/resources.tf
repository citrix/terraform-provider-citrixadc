resource "citrixadc_lbsipparameters" "tf_lbsipparameters" {
	addrportvip = "ENABLED"
	retrydur = 100
	rnatdstport = 80
	rnatsecuredstport = 81
	rnatsecuresrcport = 82
	rnatsrcport = 83
	sip503ratethreshold = 15
}
