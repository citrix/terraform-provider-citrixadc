resource "citrixadc_vlan_nsip6_binding" "tf_vlan_nsip6_binding" {
  vlanid    = 2
  ipaddress = "2001::a/96"
}
