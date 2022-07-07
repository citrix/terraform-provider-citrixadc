resource "citrixadc_nssimpleacl" "tf_nssimpleacl" {
  aclname   = "tf_nssimpleacl"
  aclaction = "DENY"
  srcip     = "1.2.3.1"
  destport  = 123
  protocol  = "UDP"
  ttl       = 600
}