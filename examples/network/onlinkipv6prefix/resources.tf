resource "citrixadc_onlinkipv6prefix" "tf_onlinkipv6prefix" {
  ipv6prefix      = "8000::/64"
  onlinkprefix    = "YES"
  autonomusprefix = "NO"
}
