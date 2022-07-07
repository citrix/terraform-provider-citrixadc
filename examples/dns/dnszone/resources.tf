resource "citrixadc_dnszone" "dnszone" {
  zonename      = "tf_zone1"
  proxymode     = "YES"
  dnssecoffload = "DISABLED"
  nsec          = "DISABLED"
}
