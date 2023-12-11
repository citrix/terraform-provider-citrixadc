#Declare require content switching vServer parameters
sslcsvip_config = {
  name        = "corpservices.example.com:443"
  vip         = "192.168.100.10"
  port        = "443"
  servicetype = "SSL"
  ersa        = "DISABLED"
  ssl3        = "DISABLED"
  tls1        = "DISABLED"
  tls11       = "DISABLED"
  httpprofile = "nshttp_default_strict_validation"
}
httpcsvip_config = {
  name        = "corpservices.example.com:80"
  vip         = "192.168.100.10"
  port        = "80"
  servicetype = "HTTP"
  redirecturl = ""
  httpprofile = "nshttp_default_strict_validation"
}

#Default Target LB vServer parameters
httpvip_config = {
  name            = "corpwebapp.corporate.com:443"
  vip             = ""
  lbmethod        = "LEASTCONNECTION"
  port            = "0"
  persistencetype = "SOURCEIP"
  servicetype     = "HTTP"
  timeout         = "20"
  httpprofile     = "nshttp_default_internal_apps"
}

#Declare require SSL service parameters to be bound to default target LB vserver"
service_config = {
  service_port = "443"
  servicetype  = "SSL"
  maxclient    = "100"
  maxreq       = "0"
  cip          = "ENABLED"
  cipheader    = "X-Forwarded-For"
  usip         = "NO"
  useproxyport = "YES"
  sp           = "OFF"
  clttimeout   = "180"
  svrtimeout   = "360"
  cka          = "YES"
  tcpb         = "YES"
  cmp          = "NO"
  //httpprofile = "http2_profile"
}
#Declare content switching action name and target LB here
cs_action1 = {
  name = "csact.webapp1"
}
#Declare content switching action name and target LB here
cs_pol1 = {
  name     = "cspol.webapp1"
  rule     = "HTTP.REQ.URL.PATH.GET(1).EQ(\"webapp1\")"
  priority = 50
}
#Declare content switching action name and target LB here
cs_action2 = {
  name = "csact.webapp2"
}
#Declare content switching action name and target LB here
cs_pol2 = {
  name     = "cspol.webapp2"
  rule     = "HTTP.REQ.URL.PATH.GET(1).EQ(\"webapp2\")"
  priority = 60
}
cs_action3 = {
  name = "csact.webapp3"
}
#Declare content switching action name and target LB here
cs_pol3 = {
  name     = "cspol.webapp3"
  rule     = "HTTP.REQ.URL.PATH.GET(1).EQ(\"webapp3\")"
  priority = 70
}
