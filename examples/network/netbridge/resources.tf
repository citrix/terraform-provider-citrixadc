resource "citrixadc_vxlanvlanmap" "tf_vxlanvlanmp" {
  name = "tf_vxlanvlanmp"
}
resource "citrixadc_netbridge" "tf_netbridge" {
  name         = "tf_netbridge"
  vxlanvlanmap = citrixadc_vxlanvlanmap.tf_vxlanvlanmp.name
}