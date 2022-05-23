resource "citrixadc_l4param" "tf_l4param" {
  l2connmethod = "MacVlanChannel"
  l4switch     = "DISABLED"
}