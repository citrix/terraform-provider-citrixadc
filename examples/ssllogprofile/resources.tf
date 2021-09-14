resource "citrixadc_ssllogprofile" "tf_ssllgoprofile" {
    name = "foo"
    ssllogclauth = "DISABLED"
    ssllogclauthfailures = "ENABLED"
    sslloghs = "ENABLED"
    sslloghsfailures = "ENABLED"	
}
