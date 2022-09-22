variable "lb_config" {
  type        = map(string)
  description = "Describes the friendly name (key=lbname),VIP(key=vip), port(key=port) and service type (key=servicetype) for the LB"
}

variable "backend_service_config" {
  type        = map(string)
  description = "Describes port(key=port) and service type (key=servicetype) for the backend services"
}

variable "backend_services" {
  description = "The list of backend services (ip addresses)"
  type        = list(string)
}

variable "ssl_config" {
  description = "Map : certkey = cert name, certfile = file with cert, keyfile = file with certificate key"
  type        = map(string)
}

variable "http_monitor_config" {
  type = map(string)
}

