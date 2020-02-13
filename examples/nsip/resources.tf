resource "citrixadc_nsip" "nsip" {
    ipaddress = "192.168.2.55"
    type = "VIP"
    netmask = "255.255.255.0"
    icmp = "ENABLED"
    state = "ENABLED"
}
