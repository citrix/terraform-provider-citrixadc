resource "citrixadc_aaagroup_vpnintranetapplication_binding" "tf_aaagroup_vpnintranetapplication_binding" {
  groupname           = "my_group"
  intranetapplication = citrixadc_vpnintranetapplication.tf_vpnintranetapplication.intranetapplication
}

resource "citrixadc_vpnintranetapplication" "tf_vpnintranetapplication" {
  intranetapplication = "tf_vpnintranetapplication"
  protocol            = "UDP"
  destip              = "2.3.6.5"
  interception        = "TRANSPARENT"
}
