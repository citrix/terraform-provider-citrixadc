resource "citrixadc_aaauser_vpnsessionpolicy_binding" "tf_aaauser_vpnsessionpolicy_binding" {
  username = "user1"
  policy   = citrixadc_vpnsessionpolicy.tf_vpnsessionpolicy.name
  priority = 100
}

resource "citrixadc_vpnsessionaction" "tf_vpnsessionaction" {
  name                       = "newsession"
  sesstimeout                = "10"
  defaultauthorizationaction = "ALLOW"
}

resource "citrixadc_vpnsessionpolicy" "tf_vpnsessionpolicy" {
  name   = "tf_vpnsessionpolicy"
  rule   = "HTTP.REQ.HEADER(\"User-Agent\").CONTAINS(\"CitrixReceiver\").NOT"
  action = citrixadc_vpnsessionaction.tf_vpnsessionaction.name
}