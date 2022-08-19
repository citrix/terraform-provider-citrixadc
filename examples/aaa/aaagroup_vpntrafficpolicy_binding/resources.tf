resource "citrixadc_aaagroup_vpntrafficpolicy_binding" "tf_aaagroup_vpntrafficpolicy_binding" {
  groupname = "my_group"
  policy    = citrixadc_vpntrafficpolicy.tf_vpntrafficpolicy.name
  priority  = 100
}

resource "citrixadc_vpntrafficaction" "foo" {
  fta        = "ON"
  hdx        = "ON"
  name       = "Testingaction"
  qual       = "tcp"
  sso        = "ON"
}
resource "citrixadc_vpntrafficpolicy" "tf_vpntrafficpolicy" {
  name   = "tf_vpntrafficpolicy"
  rule   = "HTTP.REQ.HEADER(\"User-Agent\").CONTAINS(\"CitrixReceiver\").NOT"
  action = citrixadc_vpntrafficaction.foo.name
}