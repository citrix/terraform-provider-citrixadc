resource "citrixadc_aaagroup_authorizationpolicy_binding" "tf_aaagroup_authorizationpolicy_binding" {
  groupname = "my_group"
  policy   = citrixadc_authorizationpolicy.tf_authorize.name
  priority = 100
}

resource "citrixadc_authorizationpolicy" "tf_authorize" {
  name   = "tp-authorize-1"
  rule   = "true"
  action = "ALLOW"
}