resource "citrixadc_botpolicylabel" "tf_botpolicylabel" {
  labelname = "tf_botpolicylabel"
}
resource "citrixadc_botpolicy" "tf_botpolicy" {
  name        = "tf_botpolicy"
  profilename = "BOT_BYPASS"
  rule        = "true"
  comment     = "COMMENT FOR BOTPOLICY"
}
resource "citrixadc_botpolicylabel_botpolicy_binding" "tf_binding" {
  labelname  = citrixadc_botpolicylabel.tf_botpolicylabel.labelname
  policyname = citrixadc_botpolicy.tf_botpolicy.name
  priority   = 50
}