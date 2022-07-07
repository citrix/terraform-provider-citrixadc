resource "citrixadc_nsacl6" "tf_nsacl6" {
  acl6name   = "tf_nsacl6"
  acl6action = "ALLOW"
  logstate   = "ENABLED"
  stateful   = "NO"
  ratelimit  = 120
  state      = "ENABLED"
  priority   = 20
  protocol   = "TCP"
}
resource "citrixadc_netprofile" "tf_netprofile" {
  name                   = "tf_netprofile"
  proxyprotocol          = "ENABLED"
  proxyprotocoltxversion = "V1"
}
resource "citrixadc_nat64" "tf_nat64" {
  name       = "tf_nat64"
  acl6name   = citrixadc_nsacl6.tf_nsacl6.acl6name
  netprofile = citrixadc_netprofile.tf_netprofile.name
}