variable "lb_config" {
    type = "map"
    "description" = "Describes the friendly name (key=lbname),VIP(key=vip), port(key=port) and service type (key=servicetype) for the LB"
}

variable "backend_service_config_cart" {
    type = "map"
    "description" = "Describes port(key=port) and service type (key=servicetype) for the cart backend services"
}

variable "backend_service_config_catalog" {
    type = "map"
    "description" = "Describes port(key=port) and service type (key=servicetype) for the catalog backend services"
}


variable "backend_services_cart" {
    description = "The list of backend services (ip addresses) for the cart service"
    type = "list"
}

variable "backend_services_catalog" {
    description = "The list of backend services (ip addresses) for the catalog service"
    type = "list"
}

variable "ssl_config" {
    description = "Map : certkey = cert name, certfile = file with cert, keyfile = file with certificate key"
    type = "map"
}

variable "http_monitor_config_cart" {
    type = "map"
}

variable "http_monitor_config_catalog" {
    type = "map"
}

