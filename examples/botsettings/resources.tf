resource "citrixadc_botsettings" "default" {
  sessiontimeout      = "950"
  proxyport           = "80"
  sessioncookiename   = "citrixbotid"
  dfprequestlimit     = "3"
  signatureautoupdate = "ON"
  trapurlautogenerate = "ON"
  trapurlinterval     = "3800"
  trapurllength       = "33"
}
