resource "citrixadc_contentinspectionprofile" "tf_contentinspectionprofile" {
  name             = "my_ci_profile"
  type             = "INLINEINSPECTION"
  ingressinterface = "LA/2"
  egressinterface  = "LA/3"
}
