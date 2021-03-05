resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_responderpolicy" "tf_responder_policy" {
  name   = "tf_responder_policy"
  action = "NOOP"
  rule   = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"nosuchthing\")"
}

resource "citrixadc_lbvserver_responderpolicy_binding" "tf_bind" {
    name = citrixadc_lbvserver.tf_lbvserver.name
    policyname = citrixadc_responderpolicy.tf_responder_policy.name
    priority = 120
    bindpoint = "REQUEST"
}
