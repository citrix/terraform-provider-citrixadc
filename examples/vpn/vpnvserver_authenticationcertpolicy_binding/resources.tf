resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_vserver"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_authenticationcertaction" "tf_certaction" {
  name                       = "tf_certaction"
  twofactor                  = "ON"
  defaultauthenticationgroup = "new_group"
  usernamefield              = "Subject:CN"
  groupnamefield             = "subject:grp"
}
resource "citrixadc_authenticationcertpolicy" "tf_certpolicy" {
  name      = "tf_certpolicy"
  rule      = "ns_true"
  reqaction = citrixadc_authenticationcertaction.tf_certaction.name
}
resource "citrixadc_vpnvserver_authenticationcertpolicy_binding" "tf_bind" {
  name      = citrixadc_vpnvserver.tf_vpnvserver.name
  policy    = citrixadc_authenticationcertpolicy.tf_certpolicy.name
  priority  = 80
  secondary = false
  bindpoint = "REQUEST"
}