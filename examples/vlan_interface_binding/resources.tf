resource "citrixadc_vlan" "tf_vlan" {
    vlanid = 40
    aliasname = "Management VLAN"
}

resource "citrixadc_vlan_interface_binding" "tf_bind" {
    vlanid = citrixadc_vlan.tf_vlan.vlanid
    ifnum = "1/1"
}
