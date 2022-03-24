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