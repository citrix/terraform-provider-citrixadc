resource "citrixadc_gslbvserver_domain_binding" "tf_gslbvserver_domain_binding"{
  name = citrixadc_gslbvserver.tf_gslbvserver.name
  domainname = "www.example.com"
  backupipflag = false
}
resource "citrixadc_gslbvserver" "tf_gslbvserver" {
  dnsrecordtype = "A"
  name          = "GSLB-East-Coast-Vserver"
  servicetype   = "HTTP"
  domain {
    domainname = "www.fooco.co"
    ttl        = "60"
  }
  domain {
    domainname = "www.barco.com"
    ttl        = "65"
  }
}

