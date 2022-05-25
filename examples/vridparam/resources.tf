resource "citrixadc_vridparam" "tf_vridparam" {
  sendtomaster  = "DISABLED"
  hellointerval = 1000
  deadinterval  = 3
}