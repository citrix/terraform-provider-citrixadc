resource "citrixadc_netprofile" "tf_netprofile" {
  name                   = "tf_netprofile"
  proxyprotocol          = "ENABLED"
  proxyprotocoltxversion = "V1"
}
resource "citrixadc_netprofile_srcportset_binding" "tf_binding" {
  name         = citrixadc_netprofile.tf_netprofile.name
  srcportrange = "2000"
}