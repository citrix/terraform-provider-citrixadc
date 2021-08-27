resource "citrixadc_ssldtlsprofile" "tf_ssldtlsprofile" {
	name = "tf_ssldtlsprofile"
	helloverifyrequest = "ENABLED"
	maxbadmacignorecount = 128
	maxholdqlen = 64
	maxpacketsize = 125
	maxrecordsize = 250
	maxretrytime = 5
	pmtudiscovery = "DISABLED"
	terminatesession = "ENABLED"
}