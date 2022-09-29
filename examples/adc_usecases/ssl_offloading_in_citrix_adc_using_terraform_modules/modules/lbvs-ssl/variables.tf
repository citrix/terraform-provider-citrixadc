
variable "sslvip_config" {
  type        = map(string)
  description = "Describes the friendly name (key=lbname),VIP(key=vip), port(key=port) and service type (key=servicetype) for the LB"
}
variable "sslvsparam" {
  type        = map(string)
  description = "SSL Virtual server parameters"
}
