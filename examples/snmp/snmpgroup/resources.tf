resource "citrixadc_snmpgroup" "tf_snmpgroup" {
  name    = "test_group"
  securitylevel = "authNoPriv"
  readviewname = "test_name"
}