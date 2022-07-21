resource "citrixadc_crvserver" "crvserver" {
    name = "my_vserver"
    servicetype = "HTTP"
    arp = "OFF"
}