resource "citrixadc_crvserver" "crvserver" {
  name        = "my_vserver"
  servicetype = "HTTP"
  arp         = "OFF"
}
resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
  name   = "tf_rewrite_policy"
  action = "DROP"
  rule   = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"
}
resource "citrixadc_crvserver_rewritepolicy_binding" "crvserver_rewritepolicy_binding" {
  name       = citrixadc_crvserver.crvserver.name
  policyname = citrixadc_rewritepolicy.tf_rewrite_policy.name
  priority   = 10
  bindpoint  = "RESPONSE"
}  