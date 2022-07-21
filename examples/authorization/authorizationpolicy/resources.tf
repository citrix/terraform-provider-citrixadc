resource "citrixadc_authorizationpolicy" "authorize1" {
  name   = "tp-authorize-1"
  rule   = "true"
  action = "ALLOW"
}

resource "citrixadc_authorizationpolicy" "authorize2" {
  name   = "tp-authorize-2"
  rule   = "true"
  action = "DENY"
}
