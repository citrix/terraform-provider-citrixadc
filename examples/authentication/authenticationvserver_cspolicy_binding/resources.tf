resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name           = "tf_authenticationvserver"
  servicetype    = "SSL"
  comment        = "new"
  authentication = "ON"
  state          = "DISABLED"
}
resource "citrixadc_lbvserver" "foo_lbvserver" {
  name        = "test_policy_lb"
  servicetype = "HTTP"
  ipv46       = "192.122.3.3"
  port        = 8000
  comment     = "hello"
}
resource "citrixadc_csaction" "tf_csaction" {
  name            = "tf_csaction"
  targetlbvserver = citrixadc_lbvserver.foo_lbvserver.name
}
resource "citrixadc_cspolicy" "foo_cspolicy" {
  policyname = "test_policy"
  rule       = "TRUE"
  action     = citrixadc_csaction.tf_csaction.name
}
resource "citrixadc_authenticationvserver_cspolicy_binding" "tf_bind" {
  name      = citrixadc_authenticationvserver.tf_authenticationvserver.name
  policy    = citrixadc_cspolicy.foo_cspolicy.policyname
  priority  = 90
  bindpoint = "REQUEST" #doesnot unbind for RESPONSE
}