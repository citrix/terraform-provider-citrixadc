resource "citrixadc_dnsaddrec" "dnsaddrec" {
  hostname  = "ab.root-servers.net"
  ipaddress = "65.200.211.129"
  ttl       = 3600
}