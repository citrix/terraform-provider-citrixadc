resource "citrixadc_botpolicy" "tf_botpolicy" {
  name        = "tf_botpolicy"
  profilename = "BOT_BYPASS"
  rule        = "true"
  comment     = "COMMENT FOR BOTPOLICY"
}
resource "citrixadc_botglobal_botpolicy_binding" "tf_binding" {
  policyname     = citrixadc_botpolicy.tf_botpolicy.name
  priority       = 90
  type           = "REQ_DEFAULT"
}