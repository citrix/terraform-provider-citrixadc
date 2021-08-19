resource "citrixadc_botpolicy" "demo_botpolicy1" {
  name        = "demo_botpolicy1"
  profilename = "BOT_BYPASS"
  rule        = "true"
  comment     = "COMMENT FOR BOTPOLICY"
}