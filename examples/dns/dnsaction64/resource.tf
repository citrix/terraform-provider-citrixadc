resource "citrixadc_dnsaction64" "dnsaction64" {
	actionname = "default_DNS64_action1"
    prefix = "64:ff9c::/96"
    mappedrule = "DNS.RR.TYPE.EQ(A) && !(DNS.RR.RDATA.IP.IN_SUBNET(0.0.0.0/8)"
    excluderule = "DNS.RR.TYPE.EQ(AAAA) && DNS.RR.RDATA.IPV6.IN_SUBNET(::ffff:0:0/96)"
}
