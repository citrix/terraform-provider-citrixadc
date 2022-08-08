resource "citrixadc_l3param" "tf_l3param" {
  srcnat               = "DISABLED"
  icmpgenratethreshold = 150
  overridernat         = "DISABLED"
  dropdfflag           = "DISABLED"
}
