sslvip_config = {
  name            = "example.com:443"
  vip             = "192.168.50.250"
  port            = "443"
  servicetype     = "SSL"
  lbmethod        = "LEASTCONNECTION"
  persistencetype = "NONE"
  timeout         = "0"
  httpprofile = ""
  redirectfromport = "0"
  httpsredirecturl = ""
}
sslvsparam = {
  ersa            = "DISABLED"
  ssl3            = "DISABLED"
  tls1            = "DISABLED"
  tls11           = "DISABLED"
}

httpvip_config = {
  name            = "example.com:80"
  vip             = "192.168.50.250"
  lbmethod        = "LEASTCONNECTION"
  port            = "80"
  persistencetype = "NONE"
  servicetype     = "HTTP"
  redirurl        = "https://example.com"
}
service_config = {
  service_port = "18082"
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
}

ssl_certkey_name     = "terracert"
ssl_certificate_path = "../sslcerts/terracert.cert"
ssl_key_path = "../sslcerts/terracert.key"
filename = {
  certname = "terracert.cert"
  keyname = "terracert.key"
}
