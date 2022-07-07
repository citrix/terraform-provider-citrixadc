resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_vserver_example"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_authenticationloginschema" "tf_loginschema" {
  name                    = "tf_loginschema"
  authenticationschema    = "LoginSchema/SingleAuth.xml"
  ssocredentials          = "YES"
  authenticationstrength  = "30"
  passwordcredentialindex = "10"
}
resource "citrixadc_authenticationloginschemapolicy" "tf_loginschemapolicy" {
  name    = "tf_loginschemapolicy"
  rule    = "true"
  action  = citrixadc_authenticationloginschema.tf_loginschema.name
  comment = "samplenew_testing"
}
resource "citrixadc_vpnvserver_authenticationloginschemapolicy_binding" "tf_bind" {
  name            = citrixadc_vpnvserver.tf_vpnvserver.name
  policy          = citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy.name
  priority        = 80
  secondary       = "false"
  groupextraction = "false"
  bindpoint       = "REQUEST" // doesnot unbind for other values
}