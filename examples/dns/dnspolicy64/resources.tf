resource "citrixadc_dnspolicy64" "dnspolicy64" {
  name  = "policy_1"
  rule = "dns.req.question.type.ne(aaaa)"
  action = "default_DNS64_action"
}