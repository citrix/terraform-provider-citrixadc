resource "citrixadc_authenticationdfaaction" "tf_dfaaction" {
  name       = "tf_dfaaction"
  serverurl  = "https://example.com/"
  clientid   = "cliId"
  passphrase = "secret"
}
resource "citrixadc_authenticationdfapolicy" "td_dfapolicy" {
  name   = "td_dfapolicy"
  rule   = "NS_TRUE"
  action = citrixadc_authenticationdfaaction.tf_dfaaction.name
}