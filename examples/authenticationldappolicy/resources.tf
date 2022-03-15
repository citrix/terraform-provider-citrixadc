resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
  name          = "ldapaction"
  serverip      = "1.2.3.4"
  serverport    = 8080
  authtimeout   = 1
  ldaploginname = "username"
}
resource "citrixadc_authenticationldappolicy" "tf_authenticationldappolicy" {
  name      = "tf_authenticationldappolicy"
  rule      = "NS_TRUE"
  reqaction = citrixadc_authenticationldapaction.tf_authenticationldapaction.name
}