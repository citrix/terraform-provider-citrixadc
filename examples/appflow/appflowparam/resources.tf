resource "citrixadc_appflowparam" "tf_appflowparam" {
  templaterefresh     = 200
  flowrecordinterval  = 100
  httpcookie          = "ENABLED"
  httplocation        = "ENABLED"
}
