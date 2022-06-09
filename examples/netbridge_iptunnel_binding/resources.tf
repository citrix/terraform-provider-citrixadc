resource "citrixadc_vxlanvlanmap" "tf_vxlanvlanmp" {
  name = "tf_vxlanvlanmp"
}
resource "citrixadc_nsip" "nsip" {
  ipaddress = "2.2.2.1"
  type      = "VIP"
  netmask   = "255.255.255.0"
}
resource "citrixadc_netbridge" "tf_netbridge" {
  name         = "tf_netbridge"
  vxlanvlanmap = citrixadc_vxlanvlanmap.tf_vxlanvlanmp.name
}
resource "citrixadc_iptunnel" "tf_iptunnel" {
  name             = "tf_iptunnel"
  remote           = "66.0.0.11"
  remotesubnetmask = "255.255.255.255"
  local            = citrixadc_nsip.nsip.ipaddress
  protocol         = "GRE"
}
resource "citrixadc_netbridge_iptunnel_binding" "tf_binding" {
  name   = citrixadc_netbridge.tf_netbridge.name
  tunnel = citrixadc_iptunnel.tf_iptunnel.name
}