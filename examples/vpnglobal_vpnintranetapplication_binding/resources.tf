resource "citrixadc_vpnintranetapplication" "tf_vpnintranetapplication" {
  intranetapplication = "tf_vpnintranetapplication"
  protocol            = "UDP"
  destip              = "2.3.6.5"
  interception        = "TRANSPARENT"
}
resource "citrixadc_vpnglobal_vpnintranetapplication_binding" "tf_bind" {
  intranetapplication =  citrixadc_vpnintranetapplication.tf_vpnintranetapplication.intranetapplication
}
