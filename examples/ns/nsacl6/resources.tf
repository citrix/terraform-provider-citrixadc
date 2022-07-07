resource "citrixadc_nstrafficdomain" "tf_trafficdomain" {
  td        = 2
  aliasname = "tf_trafficdomain"
  vmac      = "ENABLED"
}
resource "citrixadc_nsacl6" "tf_nsacl6" {
  acl6name   = "tf_nsacl6"
  acl6action = "ALLOW"
  td         = citrixadc_nstrafficdomain.tf_trafficdomain.td
  logstate   = "ENABLED"
  stateful   = "NO"
  ratelimit  = 120
  state      = "ENABLED"
  priority   = 20
  protocol   = "TCP"
}