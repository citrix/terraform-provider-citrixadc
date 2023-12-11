variable "username" {
  description = "ADC administrator username"
  type        = string
  sensitive   = true
}
variable "password" {
  description = "ADC administrator password"
  type        = string
  sensitive   = true
}
variable "targetlb1_config" {
  type        = map(string)
  description = "Describes the friendly name (key=lbname),VIP(key=vip), port(key=port) and service type (key=servicetype) for the LB"
}
variable "servers" {
  description = "Name and IP address of the server"
  type = map(object({
    ipaddress = string
    state     = string
    name      = string
    comment   = string
  }))
  default = {
    "testserver10" = {
      ipaddress = "192.168.50.10"
      state     = "ENABLED"
      name      = "testserver10"
      comment   = "webapp1"
    }
    "testserver11" = {
      ipaddress = "192.168.50.11"
      state     = "ENABLED"
      name      = "testserver11"
      comment   = "webapp1"
    }
  }
}
variable "service_config" {
  description = "service configuration"
  type        = map(string)
}

