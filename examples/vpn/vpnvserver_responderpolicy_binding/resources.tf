resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_examplevserver"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_responderpolicy" "tf_responder_policy" {
  name   = "tf_responder_policy"
  action = "NOOP"
  rule   = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"nosuchthing\")"
}
resource "citrixadc_vpnvserver_responderpolicy_binding" "tf_bind" {
  name      = citrixadc_vpnvserver.tf_vpnvserver.name
  policy    = citrixadc_responderpolicy.tf_responder_policy.name
  priority  = 200
}