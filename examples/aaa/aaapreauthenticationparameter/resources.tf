resource "citrixadc_aaapreauthenticationparameter" "tf_aaapreauthenticationparameter" {
  preauthenticationaction = "DENY"
  deletefiles    = "/var/tmp/*.files"
}