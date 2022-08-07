resource "citrixadc_clusternodegroup_vpnvserver_binding" "tf_clusternodegroup_vpnvserver_binding" {
  name    = "my_vpn_group"
  vserver = "my_vpnvserver"
}
