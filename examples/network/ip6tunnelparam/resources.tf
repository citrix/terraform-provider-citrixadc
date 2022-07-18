resource "citrixadc_nsip6" "name" {
  ipv6address = "2001:db8:100::fa/64"
  type        = "VIP"
  icmp        = "DISABLED"
}
resource "citrixadc_ip6tunnelparam" "tf_ip6tunnelparam" {
  srcip                = split("/",citrixadc_nsip6.name.ipv6address)[0]
  dropfrag             = "NO"
  dropfragcputhreshold = 1
  srciproundrobin      = "NO"
  useclientsourceipv6  = "NO"
}