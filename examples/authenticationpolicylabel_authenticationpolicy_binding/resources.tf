resource "citrixadc_authenticationpolicylabel" "tf_authenticationpolicylabel" {
  labelname = "tf_authenticationpolicylabel"
  type      = "AAATM_REQ"
  comment   = "Testingresource"
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
resource "citrixadc_authenticationpolicylabel_authenticationpolicy_binding" "tf_bind" {
  labelname  = citrixadc_authenticationpolicylabel.tf_authenticationpolicylabel.labelname
  policyname = citrixadc_authenticationpolicy.tf_authenticationpolicy.name
  priority   = 20
}