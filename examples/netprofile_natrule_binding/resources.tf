resource "citrixadc_netprofile" "tf_netprofile" {
  name                   = "tf_netprofile"
  proxyprotocol          = "ENABLED"
  proxyprotocoltxversion = "V1"
}
resource "citrixadc_netprofile_natrule_binding" "tf_binding" {
  name      = citrixadc_netprofile.tf_netprofile.name
  natrule   = "10.10.10.10"
  netmask   = "255.255.255.255"
  rewriteip = "3.3.3.3"
}