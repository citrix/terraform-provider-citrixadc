resource "citrixadc_botprofile" "tf_botprofile" {
  name                     = "tf_botprofile"
  errorurl                 = "http://www.citrix.com"
  trapurl                  = "/http://www.citrix.com"
  comment                  = "tf_botprofile comment"
  bot_enable_white_list    = "ON"
  bot_enable_black_list    = "ON"
  bot_enable_rate_limit    = "ON"
  devicefingerprint        = "ON"
  devicefingerprintaction  = ["LOG", "RESET"]
  bot_enable_ip_reputation = "ON"
  trap                     = "ON"
  trapaction               = ["LOG", "RESET"]
  bot_enable_tps           = "ON"
}
resource "citrixadc_botprofile_logexpression_binding" "tf_binding" {
  name                     = citrixadc_botprofile.tf_botprofile.name
  logexpression            = "true"
  bot_log_expression_name  = "tf_logname"
  bot_log_expression_value = "HTTP.REQ.BODY.CONTAINS(\"ANDROID\")"
  bot_bind_comment         = "LogTesting"
}