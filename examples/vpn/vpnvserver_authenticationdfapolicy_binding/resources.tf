resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_vpnvserver"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_authenticationdfaaction" "tf_dfaaction" {
  name       = "tf_dfaaction"
  serverurl  = "https://example.com/"
  clientid   = "cliId"
  passphrase = "secret"
}
resource "citrixadc_authenticationdfapolicy" "td_dfapolicy" {
  name   = "td_dfapolicy"
  rule   = "ns_true"
  action = citrixadc_authenticationdfaaction.tf_dfaaction.name
}
resource "citrixadc_vpnvserver_authenticationdfapolicy_binding" "tf_bind" {
  name            = citrixadc_vpnvserver.tf_vpnvserver.name
  policy          = citrixadc_authenticationdfapolicy.td_dfapolicy.name
  priority        = 50
  groupextraction = "0"
  bindpoint       = "REQUEST"
}