resource "citrixadc_bridgegroup" "tf_bridgegroup" {
  bridgegroup_id     = 2
  dynamicrouting     = "DISABLED"
  ipv6dynamicrouting = "DISABLED"
}