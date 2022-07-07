resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name           = "tf_authenticationvserver"
  servicetype    = "SSL"
  comment        = "new"
  authentication = "ON"
  state          = "DISABLED"
}
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
resource "citrixadc_authenticationvserver_authenticationpolicy_binding" "tf_bind" {
  name     = citrixadc_authenticationvserver.tf_authenticationvserver.name
  policy   = citrixadc_authenticationpolicy.tf_authenticationpolicy.name
  priority = 30
}