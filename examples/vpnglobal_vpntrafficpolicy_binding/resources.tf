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
resource "citrixadc_vpngobal_vpntrafficpolicy_binding" "tf_bond" {
  policyname = citrixadc_vpntrafficpolicy.tf_vpntrafficpolicy.name
  priority   = 20
}