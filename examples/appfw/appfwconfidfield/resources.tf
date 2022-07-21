resource "citrixadc_appfwconfidfield" "tf_confidfield" {
  fieldname = "tf_confidfield"
  url       = "www.example.com/"
  isregex   = "REGEX"
  comment   = "Testing"
  state     = "DISABLED"
}