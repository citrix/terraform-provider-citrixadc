resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name           = "tf_authenticationvserver"
  servicetype    = "SSL"
  comment        = "new"
  authentication = "ON"
  state          = "DISABLED"
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
resource "citrixadc_authenticationvserver_authenticationcertpolicy_binding" "tf_bind" {
  name      = citrixadc_authenticationvserver.tf_authenticationvserver.name
  policy    = citrixadc_authenticationcertpolicy.tf_certpolicy.name
  priority  = 90
  bindpoint = "AAA_REQUEST"
}