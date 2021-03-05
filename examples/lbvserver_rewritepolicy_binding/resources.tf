resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
  name   = "tf_test_rewrite_policy"
  action = "DROP"
  rule   = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"
}

resource "citrixadc_lbvserver_rewritepolicy_binding" "tf_bind" {
    name = citrixadc_lbvserver.tf_lbvserver.name
    policyname = citrixadc_rewritepolicy.tf_rewrite_policy.name
    priority = 110
    bindpoint = "REQUEST"
}
