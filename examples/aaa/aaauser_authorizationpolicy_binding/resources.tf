resource "citrixadc_aaauser_authorizationpolicy_binding" "tf_aaauser_authorizationpolicy_binding" {
  username = "user1"
  policy   = citrixadc_authorizationpolicy.tf_authorize.name
  priority = 100
}

resource "citrixadc_authorizationpolicy" "tf_authorize" {
  name   = "tp-authorize-1"
  rule   = "true"
  action = "ALLOW"
}
