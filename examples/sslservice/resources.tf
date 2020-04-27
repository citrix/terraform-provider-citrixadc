resource "citrixadc_lbvserver" "test_lbvserver" {
    name = "test_lbvserver"
    ipv46 = "10.33.55.33"
    port = 80

}

resource "citrixadc_service" "test_service" {
    servicetype = "SSL"
    name = "test_service"
    ipaddress = "10.77.33.22"
    ip = "10.77.33.22"
    port = "443"
    lbvserver = citrixadc_lbvserver.test_lbvserver.name
    snienable = "ENABLED"
}
