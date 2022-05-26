resource "citrixadc_nstrafficdomain" "tf_trafficdomain" {
  td        = 2
  aliasname = "tf_trafficdomain"
  vmac      = "DISABLED"
}
resource "citrixadc_bridgegroup" "tf_bridgegroup" {
  bridgegroup_id     = 2
  dynamicrouting     = "DISABLED"
  ipv6dynamicrouting = "DISABLED"
}
resource "citrixadc_nstrafficdomain_bridgegroup_binding" "tf_binding" {
  td          = citrixadc_nstrafficdomain.tf_trafficdomain.td
  bridgegroup = citrixadc_bridgegroup.tf_bridgegroup.bridgegroup_id
}