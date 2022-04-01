resource "citrixadc_vpnglobal_intranetip_binding" "tf_bind" {
  intranetip = "2.3.4.5"
  netmask    = "255.255.255.0"
}