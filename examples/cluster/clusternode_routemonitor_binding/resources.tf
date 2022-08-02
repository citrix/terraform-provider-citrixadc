resource "citrixadc_clusternode_routemonitor_binding" "tf_clusternode_routemonitor_binding" {
  nodeid       = 1
  routemonitor = "10.222.74.128"
  netmask      = "255.255.255.192"
}
