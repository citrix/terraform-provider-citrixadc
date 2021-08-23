resource citrixadc_lbvserver_botpolicy_binding demo_lbvserver_botpolicy_binding {
  name                   = citrixadc_lbvserver.demo_lb.name
  policyname             = citrixadc_botpolicy.demo_botpolicy.name
  labeltype              = "reqvserver" # Possible values = reqvserver, resvserver, policylabel
  labelname              = citrixadc_lbvserver.demo_lb.name
  priority               = 100
  bindpoint              = "REQUEST" # Possible values = REQUEST, RESPONSE
  gotopriorityexpression = "END"
  invoke                 = true         # boolean
}

resource "citrixadc_lbvserver" "demo_lb" {
  name        = "demo_lb"
  servicetype = "HTTP"
}

resource "citrixadc_botpolicy" "demo_botpolicy" {
  name        = "demo_botpolicy"
  profilename = citrixadc_botprofile.tf_botprofile.name
  rule        = "true"
  comment     = "COMMENT FOR BOTPOLICY"
}

resource "citrixadc_botprofile" "tf_botprofile" {
	name = "tf_botprofile"
	errorurl = "http://www.citrix.com"
	trapurl = "/http://www.citrix.com"
	comment = "tf_botprofile comment"
	bot_enable_white_list = "ON"
	bot_enable_black_list = "ON"
	bot_enable_rate_limit = "ON"
	devicefingerprint = "ON"
	devicefingerprintaction = ["LOG", "RESET"]
	bot_enable_ip_reputation = "ON"
	trap = "ON"
	trapaction = ["LOG", "RESET"]
	bot_enable_tps = "ON"
}
