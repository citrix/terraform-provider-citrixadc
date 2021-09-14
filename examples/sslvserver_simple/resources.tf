resource "citrixadc_sslvserver" "tf_sslvserver" {
  cipherredirect = "ENABLED"
  cipherurl = "http://www.citrix.com"
  cleartextport = "80"
  clientauth = "ENABLED"
  clientcert = "Optional"
  hsts = "ENABLED"
  includesubdomains = "YES"
  maxage = "100"
  ocspstapling = "ENABLED"
  preload = "YES"
  sendclosenotify = "YES"
  sessreuse = "ENABLED"
  sesstimeout = "180"
  snienable = "ENABLED"
  sslredirect = "ENABLED"
  strictsigdigestcheck = "ENABLED"
  tls1 = "ENABLED"
  tls11 = "ENABLED"
  tls12 = "ENABLED"
  tls13 = "ENABLED"
  tls13sessionticketsperauthcontext = "7"
  zerorttearlydata = "ENABLED"
  vservername = citrixadc_lbvserver.tf_lbvserver.name
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_vserver"
  servicetype = "SSL"
}