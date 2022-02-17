resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
  name          = "ldapaction"
  serverip      = "1.2.3.4"
  serverport    = 8080
  authtimeout   = 1
  ldaploginname = "username"
}
resource "citrixadc_authenticationpolicy" "tf_authenticationpolicy" {
  name   = "tf_authenticationpolicy"
  rule   = "true"
  action = citrixadc_authenticationldapaction.tf_authenticationldapaction.name
}