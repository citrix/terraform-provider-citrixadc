resource "citrixadc_vlan_channel_binding" "tf_vlan_channel_binding" {
  vlanid = 2
  ifnum  = "LA/2"
  tagged = false
}
