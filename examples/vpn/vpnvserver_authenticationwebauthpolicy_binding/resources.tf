resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_examplevserver"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_authenticationwebauthaction" "tf_webauthaction" {
  name                       = "tf_webauthaction"
  serverip                   = "1.2.3.4"
  serverport                 = 8080
  fullreqexpr                = "TRUE"
  scheme                     = "http"
  successrule                = "http.RES.STATUS.EQ(200)"
  defaultauthenticationgroup = "new_group"
}
resource "citrixadc_authenticationwebauthpolicy" "tf_webauthpolicy" {
  name   = "tf_webauthpolicy"
  rule   = "NS_TRUE"
  action = citrixadc_authenticationwebauthaction.tf_webauthaction.name
}
resource "citrixadc_vpnvserver_authenticationwebauthpolicy_binding" "tf_bind" {
  name      = citrixadc_vpnvserver.tf_vpnvserver.name
  policy    = citrixadc_authenticationwebauthpolicy.tf_webauthpolicy.name
  priority  = 80
  bindpoint = "OTHERTCP_REQUEST"
}