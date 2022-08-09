resource "citrixadc_nd6ravariables" "tf_nd6ravariables" {
  vlan                     = 1
  ceaserouteradv           = "NO"
  onlyunicastrtadvresponse = "YES"
  srclinklayeraddroption   = "NO"
}
