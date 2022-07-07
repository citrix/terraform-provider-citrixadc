resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tfvserver_example"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_authenticationlocalpolicy" "tf_localpolicy" {
  name = "tf_localpolicy"
  rule = "ns_true"
}
resource "citrixadc_vpnvserver_authenticationlocalpolicy_binding" "tf_bind" {
  name            = citrixadc_vpnvserver.tf_vpnvserver.name
  policy          = citrixadc_authenticationlocalpolicy.tf_localpolicy.name
  priority        = 90
  groupextraction = false
  secondary       = false
  bindpoint       = "REQUEST"
}