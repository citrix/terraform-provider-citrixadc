resource "citrixadc_servicegroup" "backend" {
  servicegroupname    = "test_service_group"
  servicetype         = "HTTP"
  servicegroupmembers = ["192.168.1.1:80:10"]
}

