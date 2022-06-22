resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_vpnvserver"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_authenticationoauthidpprofile" "tf_idpprofile" {
  name         = "tf_idpprofile"
  clientid     = "cliId"
  clientsecret = "secret"
  redirecturl  = "http://www.example.com/1/"
}
resource "citrixadc_authenticationoauthidppolicy" "tf_idppolicy" {
  name    = "tf_idppolicy"
  rule    = "true"
  action  = citrixadc_authenticationoauthidpprofile.tf_idpprofile.name
  comment = "aboutpolicy"
}
resource "citrixadc_vpnvserver_authenticationoauthidppolicy_binding" "tf_bind" {
  name      = citrixadc_vpnvserver.tf_vpnvserver.name
  policy    = citrixadc_authenticationoauthidppolicy.tf_idppolicy.name
  priority  = 70
  bindpoint = "REQUEST" // Doesn't unbind for other values
  secondary = "false"
}