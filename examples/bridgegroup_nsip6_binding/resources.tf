resource "citrixadc_bridgegroup" "tf_bridgegroup" {
  bridgegroup_id     = 2
  dynamicrouting     = "DISABLED"
  ipv6dynamicrouting = "DISABLED"
}
resource "citrixadc_nsip6" "test_nsip" {
  ipv6address = "2001:db8:100::fb/64"
  type        = "VIP"
  icmp        = "DISABLED"
}
resource "citrixadc_bridgegroup_nsip6_binding" "tf_binding" {
  bridgegroup_id = citrixadc_bridgegroup.tf_bridgegroup.bridgegroup_id
  ipaddress      = citrixadc_nsip6.test_nsip.ipv6address
}