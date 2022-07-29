resource "citrixadc_clusternode" "tf_clusternode" {
  nodeid             = 1
  ipaddress          = "10.222.74.150"
  state              = "ACTIVE"
}
