targetlb2_config = {
  name            = "webapp2.corporate.com:443-ns"
  vip             = ""
  lbmethod        = "LEASTCONNECTION"
  port            = "0"
  persistencetype = "NONE"
  servicetype     = "HTTP"
  timeout = "0"
  redirurl        = ""
}

service_config = {
  service_port = "443"
  servicetype  = "SSL"
  maxclient    = "80"
  maxreq       = "10000"
  cip          = "ENABLED"
  cipheader    = "X-Forwarded-For"
  usip         = "NO"
  useproxyport = "YES"
  sp           = "ON"
  clttimeout   = "180"
  svrtimeout   = "360"
  cka          = "YES"
  tcpb         = "NO"
  cmp          = "NO"
}
