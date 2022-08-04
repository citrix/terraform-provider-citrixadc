resource "citrixadc_snmpuser" "tf_snmpuser" {
  name       = "test_user"
  group      = "test_group"
  authtype   = "MD5"
  authpasswd = "this_is_my_password"
  privtype   = "DES"
  privpasswd = "this_is_my_password2"
}


resource "citrixadc_snmpgroup" "tf_snmpgroup" {
  name    = "test_group"
  securitylevel = "authNoPriv"
  readviewname = "test_readviewname"
}
