variable "lb_config" {
    type = "map"
    "description" = "Describes the friendly name (key=lbname),VIP(key=vip), port(key=port) and service type (key=servicetype) for the LB"
}

variable "backend_service_config" {
    type = "map"
    "description" = "Describes port(key=port) and service type (key=servicetype) for the backend services"
}


variable "backend_services" {
    description = "The list of backend services (ip addresses)"
    type = "list"
}

variable "ssl_config" {
    description = "Map : certkey = cert name, certfile = file with cert, keyfile = file with certificate key"
    type = "map"
}

variable "http_monitor_config" {
    type = "map"
}

