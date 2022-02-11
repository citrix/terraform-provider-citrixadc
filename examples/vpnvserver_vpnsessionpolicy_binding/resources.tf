resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "vpn_vserver"
  servicetype = "SSL"
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

resource "citrixadc_vpnvserver_vpnsessionpolicy_binding" "tf_bind" {
  name      = citrixadc_vpnvserver.tf_vpnvserver.name
  policy    = citrixadc_vpnsessionpolicy.tf_vpnsessionpolicy.name
  priority  = 20
}
