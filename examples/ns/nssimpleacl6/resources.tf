resource "citrixadc_nssimpleacl6" "tf_nssimpleacl6" {
  aclname   = "tf_nssimpleacl6"
  aclaction = "DENY"
  srcipv6   = "3ffe:192:168:215::82"
  destport  = 123
  protocol  = "TCP"
  ttl       = 600
}