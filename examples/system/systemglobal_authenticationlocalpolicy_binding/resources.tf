resource "citrixadc_systemglobal_authenticationlocalpolicy_binding" "tf_systemglobal_authenticationlocalpolicy_binding" {
  policyname = citrixadc_authenticationlocalpolicy.tf_authenticationlocalpolicy.name
  priority   = 50
}

resource "citrixadc_authenticationlocalpolicy" "tf_authenticationlocalpolicy" {
  name   = "tf_authenticationlocalpolicy"
  rule   = "ns_true"
}