resource "citrixadc_crvserver" "crvserver" {
  name        = "my_vserver"
  servicetype = "HTTP"
  arp         = "OFF"
}
resource "citrixadc_lbvserver" "foo_lbvserver" {
  name        = "test_policy_lbv"
  servicetype = "HTTP"
  ipv46       = "192.122.3.30"
  port        = 8000
  comment     = "hello"
}
resource "citrixadc_csaction" "tf_csaction" {
  name            = "test_csaction"
  targetlbvserver = citrixadc_lbvserver.foo_lbvserver.name
}
resource "citrixadc_cspolicy" "foo_cspolicy" {
  policyname = "test_cspolicy"
  rule       = "TRUE"
  action     = citrixadc_csaction.tf_csaction.name
}
resource "citrixadc_service" "tf_service" {
  lbvserver = citrixadc_lbvserver.foo_lbvserver.name
  name = "tf_service1"
  port = 8080
  ip = "10.202.22.111"
  servicetype = "HTTP"
  cachetype = "TRANSPARENT"
}
resource "citrixadc_crvserver_cspolicy_binding" "crvserver_cspolicy_binding" {
  name       = citrixadc_crvserver.crvserver.name
  policyname = citrixadc_cspolicy.foo_cspolicy.policyname
  priority   = 90
}