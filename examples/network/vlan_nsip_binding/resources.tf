resource "citrixadc_vlan" "tf_vlan" {
    vlanid = 40
    aliasname = "Management VLAN"
}

resource "citrixadc_nsip" "tf_snip" {
    ipaddress = "10.222.74.146"
    type = "SNIP"
    netmask = "255.255.255.0"
    icmp = "ENABLED"
    state = "ENABLED"
}

resource "citrixadc_vlan_nsip_binding" "tf_bind" {
    vlanid = citrixadc_vlan.tf_vlan.vlanid
    ipaddress = citrixadc_nsip.tf_snip.ipaddress
    netmask = citrixadc_nsip.tf_snip.netmask
}
