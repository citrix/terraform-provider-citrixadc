resource "citrixadc_icaaccessprofile" "tf_icaaccessprofile" {
  name                   = "my_ica_accessprofile"
  connectclientlptports  = "DEFAULT"
  localremotedatasharing = "DEFAULT"
}
