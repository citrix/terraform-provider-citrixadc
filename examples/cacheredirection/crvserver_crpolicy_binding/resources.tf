resource "citrixadc_crpolicy" "crpolicy" {
    policyname = "crpolicy1"
    rule = "CLIENT.IP.SRC.IN_SUBNET(1.1.1.1/24)"
    action = "ORIGIN"
}
resource "citrixadc_crvserver" "crvserver" {
    name = "my_vserver"
    servicetype = "HTTP"
    arp = "OFF"
}
resource "citrixadc_crvserver_crpolicy_binding" "crvserver_crpolicy_binding" {
    name = citrixadc_crvserver.crvserver.name
    policyname = citrixadc_crpolicy.crpolicy.policyname
    priority = 10 
}