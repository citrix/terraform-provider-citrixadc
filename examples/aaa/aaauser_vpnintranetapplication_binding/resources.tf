resource "citrixadc_aaauser_vpnintranetapplication_binding" "tf_aaauser_vpnintranetapplication_binding" {
  username            = "user1"
  intranetapplication = citrixadc_vpnintranetapplication.tf_vpnintranetapplication.intranetapplication
}

resource "citrixadc_vpnintranetapplication" "tf_vpnintranetapplication" {
  intranetapplication = "tf_vpnintranetapplication"
  protocol            = "UDP"
  destip              = "2.3.6.5"
  interception        = "TRANSPARENT"
}
