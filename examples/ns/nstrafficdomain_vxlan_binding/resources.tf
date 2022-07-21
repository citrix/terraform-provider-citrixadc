resource "citrixadc_nstrafficdomain" "tf_trafficdomain" {
  td        = 2
  aliasname = "tf_trafficdomain"
}
resource "citrixadc_vxlan" "tf_vxlan" {
  vxlanid            = 123
  port               = 33
  dynamicrouting     = "DISABLED"
  ipv6dynamicrouting = "DISABLED"
  innervlantagging   = "ENABLED"
}
resource "citrixadc_nstrafficdomain_vxlan_binding" "tf_binding" {
  td    = citrixadc_nstrafficdomain.tf_trafficdomain.td
  vxlan = citrixadc_vxlan.tf_vxlan.vxlanid
}