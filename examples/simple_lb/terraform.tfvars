vip_config = {
  vip = "10.22.22.22"
}

backend_service_config = {
   clttimeout = 40
   backend_port = 80
}

backend_services = [
  "172.33.33.33",
  "172.33.44.33",
  "172.33.44.35",
  "172.33.44.34",
  "172.33.44.39",
]
