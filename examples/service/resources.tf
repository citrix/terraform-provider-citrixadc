resource "citrixadc_service" "tf_service" {
    servicetype = "HTTP"
    name = "tf_service"
    ipaddress = "10.77.33.22"
    ip = "10.77.33.22"
    port = "80"
}
