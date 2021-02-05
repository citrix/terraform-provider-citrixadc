resource "citrixadc_nsip6" "test_nsip" {
    ipv6address = "2001:db8:100::fb/64"
    type = "VIP"
    icmp = "DISABLED"
    #mptcpadvertise = "YES" # "YES" | "NO"
}
