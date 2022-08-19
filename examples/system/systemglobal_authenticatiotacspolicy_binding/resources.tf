resource "citrixadc_systemglobal_authenticationtatacspolicy_binding" "tf_systemglobal_authenticationtatacspolicy_binding" {
  policyname = citrixadc_authenticationtacacspolicy.tf_tacacspolicy.name
  priority   = 50
}

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
