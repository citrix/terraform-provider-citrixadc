variable "vip_config" {
    type = "map"
    "description" = "Describes the friendly name (key=lbname),VIP(key=vip), port(key=port) and service type (key=servicetype) for the LB"
}

variable "backend_service_config" {
    type = "map"
    "description" = "Describes port(key=port) and service type (key=servicetype) for the backend services"
}


variable "backend_services" {
    description = "The list of backend services (ip address:port list)"
    type = "list"
}

