resource "citrixadc_bridgegroup" "tf_bridgegroup" {
  bridgegroup_id     = 2
  dynamicrouting     = "DISABLED"
  ipv6dynamicrouting = "DISABLED"
}
resource "citrixadc_nsip" "nsip" {
  ipaddress = "2.2.2.3"
  type      = "VIP"
  netmask   = "255.255.255.0"
}
resource "citrixadc_bridgegroup_nsip_binding" "tf_binding" {
  bridgegroup_id = citrixadc_bridgegroup.tf_bridgegroup.bridgegroup_id
  ipaddress      = citrixadc_nsip.nsip.ipaddress
  netmask        = citrixadc_nsip.nsip.netmask
}