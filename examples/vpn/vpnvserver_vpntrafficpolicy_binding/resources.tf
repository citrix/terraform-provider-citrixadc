resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_examplevserver"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_vpntrafficaction" "foo" {
  fta  = "ON"
  hdx  = "ON"
  name = "Testingaction"
  qual = "tcp"
  sso  = "ON"
}
resource "citrixadc_vpntrafficpolicy" "tf_vpntrafficpolicy" {
  name   = "tf_vpntrafficpolicy"
  rule   = "HTTP.REQ.HEADER(\"User-Agent\").CONTAINS(\"CitrixReceiver\").NOT"
  action = citrixadc_vpntrafficaction.foo.name
}
resource "citrixadc_vpnvserver_vpntrafficpolicy_binding" "tf_bind" {
  name                   = citrixadc_vpnvserver.tf_vpnvserver.name
  policy                 = citrixadc_vpntrafficpolicy.tf_vpntrafficpolicy.name
  priority               = 200
  bindpoint              = "REQUEST"
  gotopriorityexpression = "NEXT"
}