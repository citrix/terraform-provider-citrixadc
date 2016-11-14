lb_config = {
  lbname = "ProductionLB"
  vip = "10.22.22.22"
  servicetype = "HTTP"
  port = 80
}

backend_service_config = {
  servicetype = "HTTP"
  port = 8080
}

backend_services = [
  "172.33.33.33",
  "172.33.44.33",
  "172.33.44.34",
  "172.33.44.35",
]
