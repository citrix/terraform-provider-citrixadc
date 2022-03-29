resource "citrixadc_authenticationtacacsaction" "tf_tacacsaction" {
  name            = "tf_tacacsaction"
  serverip        = "1.2.3.4"
  serverport      = 8080
  authtimeout     = 5
  authorization   = "ON"
  accounting      = "ON"
  auditfailedcmds = "ON"
  groupattrname   = "group"
}
resource "citrixadc_authenticationtacacspolicy" "tf_tacacspolicy" {
  name      = "tf_tacacspolicy"
  rule      = "NS_FALSE"
  reqaction = citrixadc_authenticationtacacsaction.tf_tacacsaction.name

}
resource "citrixadc_vpnglobal_authenticationtacacspolicy_binding" "tf_bind" {
  policyname      = citrixadc_authenticationtacacspolicy.tf_tacacspolicy.name
  priority        = 80
  secondary       = false
  groupextraction = false
}