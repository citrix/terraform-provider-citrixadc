resource "citrixadc_appfwsettings" "tf_appfwsettings" {
  defaultprofile           = "APPFW_BYPASS"
  undefaction              = "APPFW_BLOCK"
  sessiontimeout           = 900
  learnratelimit           = 400
  sessionlifetime          = 0
  sessioncookiename        = "citrix_ns_id"
  importsizelimit          = 134217728
  signatureautoupdate      = "OFF"
  signatureurl             = "https://s3.amazonaws.com/NSAppFwSignatures/SignaturesMapping.xml"
  cookiepostencryptprefix  = "ENC"
  geolocationlogging       = "OFF"
  ceflogging               = "OFF"
  entitydecoding           = "OFF"
  useconfigurablesecretkey = "OFF"
  sessionlimit             = 100000
  malformedreqaction = [
    "block",
    "log",
    "stats"
  ]
  centralizedlearning = "OFF"
  proxyport           = 8080
}
