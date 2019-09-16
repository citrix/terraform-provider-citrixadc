variable "lb_config" {
  type        = map(string)
  description = "Describes the friendly name (key=lbname),VIP(key=vip), port(key=port) and service type (key=servicetype) for the LB"
}

variable "backend_service_config_cart" {
  type        = map(string)
  description = "Describes port(key=port) and service type (key=servicetype) for the cart backend services"
}

variable "backend_service_config_catalog" {
  type        = map(string)
  description = "Describes port(key=port) and service type (key=servicetype) for the catalog backend services"
}

variable "backend_services_cart" {
  description = "The list of backend services (ip addresses) for the cart service"
  type        = list(string)
}

variable "backend_services_catalog" {
  description = "The list of backend services (ip addresses) for the catalog service"
  type        = list(string)
}

variable "ssl_config" {
  description = "Map : certkey = cert name, certfile = file with cert, keyfile = file with certificate key"
  type        = map(string)
}

variable "http_monitor_config_cart" {
  type = map(string)
}

variable "http_monitor_config_catalog" {
  type = map(string)
}

