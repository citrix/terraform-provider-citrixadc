resource "citrixadc_dnspolicy" "dnspolicy" {
  name = "policy_A"
  rule = "CLIENT.IP.SRC.IN_SUBNET(1.1.1.1/24)"
  drop = "YES"
}
resource "citrixadc_dnsglobal_dnspolicy_binding" "dnsglobal_dnspolicy_binding" {
  policyname = citrixadc_dnspolicy.dnspolicy.name
  priority   = 30
  type       = "REQ_DEFAULT"
}