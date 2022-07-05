resource "citrixadc_dnscnamerec" "dnscnamerec" {
	aliasname = "citrixadc.cloud.com"
    canonicalname = "ctxwsp-citrixadc-fdproxy-global.trafficmanager.net"
    ttl = 3600
}