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
resource "citrixadc_vpnglobal_authenticationsamlpolicy_binding" "tf_bind" {
  policyname      = citrixadc_authenticationsamlpolicy.tf_samlpolicy.name
  priority        = 80
  groupextraction = false
  secondary       = false
}