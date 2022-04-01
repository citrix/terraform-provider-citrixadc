resource "citrixadc_vpnclientlessaccesspolicy" "tf_vpnclientlessaccesspolicy" {
  name        = "tf_vpnclientlessaccesspolicy"
  profilename = "ns_cvpn_default_profile"
  rule        = "true"
}
resource "citrixadc_vpnglobal_vpnclientlessaccesspolicy_binding" "tf_bind" {
  policyname     = citrixadc_vpnclientlessaccesspolicy.tf_vpnclientlessaccesspolicy.name
  priority       = 90
  globalbindtype = "RNAT_GLOBAL"
  secondary      = "false"
  type           = "RES_OVERRIDE"
}