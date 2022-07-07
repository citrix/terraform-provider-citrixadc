resource "citrixadc_iptunnel" "tf_iptunnel" {
  name             = "tf_iptunnel"
  remote           = "66.0.0.11"
  remotesubnetmask = "255.255.255.255"
  local            = "*"
}
resource "citrixadc_nspbr6" "tf_nspbr6" {
  name     = "tf_nspbr6"
  action   = "ALLOW"
  protocol = "ICMPV6"
  priority = 20
  state    = "ENABLED"
  iptunnel = citrixadc_iptunnel.tf_iptunnel.name
}