resource "citrixadc_authenticationlocalpolicy" "tf_authenticationlocalpolicy" {
  name   = "tf_authenticationlocalpolicy"
  rule   = "ns_true"
}