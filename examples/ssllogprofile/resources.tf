resource "citrixadc_ssllogprofile" "foo" {
    name = "foo"
    ssllogclauth = "DISABLED"
    ssllogclauthfailures = "ENABLED"
    sslloghs = "ENABLED"
    sslloghsfailures = "ENABLED"	
}
