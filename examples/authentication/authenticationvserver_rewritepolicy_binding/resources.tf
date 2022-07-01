resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name           = "tf_authenticationvserver"
  servicetype    = "SSL"
  comment        = "new"
  authentication = "ON"
  state          = "DISABLED"
}
resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
  name   = "tf_rewrite_policy"
  action = "DROP"
  rule   = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"
}
resource "citrixadc_authenticationvserver_rewritepolicy_binding" "tf_bind" {
  name                   = citrixadc_authenticationvserver.tf_authenticationvserver.name
  policy                 = citrixadc_rewritepolicy.tf_rewrite_policy.name
  priority               = 90
  bindpoint              = "RESPONSE"
  gotopriorityexpression = "END"
  groupextraction        = "false"
}