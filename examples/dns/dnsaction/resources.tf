resource "citrixadc_dnsaction" "dnsaction" {
  actionname       = "tf_action1"
  actiontype       = "Rewrite_Response"
  ipaddress        = ["192.0.2.20","192.0.2.56","198.51.130.10"]
  dnsprofilename   = "tf_profile1"

}