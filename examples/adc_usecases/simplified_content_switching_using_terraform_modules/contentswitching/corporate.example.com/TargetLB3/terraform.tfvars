targetlb3_config = {
  name            = "webapp3.corporate.com:443-ns"
  vip             = ""
  lbmethod        = "LEASTCONNECTION"
  port            = "0"
  persistencetype = "SOURCE"
  servicetype     = "HTTP"
  timeout         = "20"
  //redirurl        = ""
  //httpprofile = "http2_profile"
}

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
