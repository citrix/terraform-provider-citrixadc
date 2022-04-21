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
resource "citrixadc_botprofile_ipreputation_binding" "tf_binding" {
  name              = citrixadc_botprofile.tf_botprofile.name
  bot_ipreputation  = "true"
  category          = "BOTNETS"
  bot_iprep_action  = ["LOG", "REDIRECT"]
  bot_bind_comment  = "TestingIpreputation"
  bot_iprep_enabled = "ON"
  logmessage        = "MessageTesting"
}