resource "citrixadc_dnspolicy" "dnspolicy" {
  name = "policy1"
  rule = "dns.req.question.type.ne(aaaa)"
  drop = "NO"
}

