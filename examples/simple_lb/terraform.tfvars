vip_config = {
  vip = "10.22.22.22"
}

backend_service_config = {
   clttimeout = 40
}

backend_services = [
  "172.33.33.33:8080",
  "172.33.44.33:80",
  "172.33.44.35:80",
  "172.33.44.34:80",
]
