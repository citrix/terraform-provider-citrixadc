resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name           = "tf_vserver_examples"
  servicetype    = "SSL"
  ipv46          = "3.3.3.3"
  port           = 443
}
resource "citrixadc_authenticationsamlaction" "tf_samlaction" {
  name                    = "tf_samlaction"
  metadataurl             = "http://www.example.com"
  samltwofactor           = "OFF"
  requestedauthncontext   = "minimum"
  digestmethod            = "SHA1"
  signaturealg            = "RSA-SHA256"
  metadatarefreshinterval = 1
}
resource "citrixadc_authenticationsamlpolicy" "tf_samlpolicy" {
  name      = "tf_samlpolicy"
  rule      = "NS_TRUE"
  reqaction = citrixadc_authenticationsamlaction.tf_samlaction.name
}
resource "citrixadc_vpnvserver_authenticationsamlpolicy_binding" "tf_bind" {
  name = citrixadc_vpnvserver.tf_vpnvserver.name
  policy = citrixadc_authenticationsamlpolicy.tf_samlpolicy.name
  priority = 80
  bindpoint = "ICA_REQUEST"
}