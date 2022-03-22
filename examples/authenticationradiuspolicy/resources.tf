resource "citrixadc_authenticationradiusaction" "tf_radiusaction" {
  name         = "tf_radiusaction"
  radkey       = "secret"
  serverip     = "1.2.3.4"
  serverport   = 8080
  authtimeout  = 2
  radnasip     = "DISABLED"
  passencoding = "chap"
}
resource "citrixadc_authenticationradiuspolicy" "tf_radiuspolicy" {
  name      = "tf_radiuspolicy"
  rule      = "NS_TRUE"
  reqaction = citrixadc_authenticationradiusaction.tf_radiusaction.name
}