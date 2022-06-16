resource "citrixadc_dnsaction" "dnsaction" {
  actionname       = "tf_action1"
  actiontype       = "ViewName"
  //ipaddress        = ["192.0.2.20","192.0.2.56","198.51.100.10"]
  //ttl              = 3600
  viewname         = "view1"
  //preferredloclist = ["NA.tx.ns1.*.*.*","NA.tx.ns2.*.*.*","NA.tx.ns3.*.*.*"]
  dnsprofilename   = "tf_profile1"

}
//add dns action dns_act_response_rewrite Rewrite_Response -IPAddress 192.0.2.20 192.0.2.56 198.51.100.10