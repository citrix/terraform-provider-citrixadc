resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
  name          = "tf_ldapaction"
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
resource "citrixadc_systemglobal_authenticationldappolicy_binding" "tf_bind" {
  policyname     = citrixadc_authenticationldappolicy.tf_authenticationldappolicy.name
  globalbindtype = "RNAT_GLOBAL"
  priority       = 88
  feature        = "SYSTEM"
}