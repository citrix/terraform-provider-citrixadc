resource "citrixadc_pcpprofile" "tf_pcpprofile" {
  name               = "my_pcpprofile"
  mapping            = "ENABLED"
  peer               = "ENABLED"
}
