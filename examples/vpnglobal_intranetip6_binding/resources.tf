resource "citrixadc_vpnglobal_intranetip6_binding" "tf_bind" {
  intranetip6 = "2.3.4.5"
  numaddr     = "45"
}