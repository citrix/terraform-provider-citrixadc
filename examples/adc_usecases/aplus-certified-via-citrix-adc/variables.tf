# SSL Service 1
variable "service1_name" {
  default = "http_service1"
}

variable "service1_ip" {
}

variable "service1_port" {
}

variable "service1_servicetype" {
}

# SSL Service 2
variable "service2_name" {
  default = "http_service2"
}

variable "service2_ip" {
}

variable "service2_port" {
}

variable "service2_servicetype" {
}

# Production SSL LB Vservice
variable "production_lb_name" {
  default = "ssl_vserver"
}

variable "production_lb_ip" {
}

# SSL CertKey
variable "ssl_certkey_name" {
  default = "ssl-certkey1"
}

variable "ssl_certificate_path" {
}

variable "ssl_key_path" {
}

variable "ssl_cacert_path" {
}

variable "ssl_cacert_name" {
  default = "ssl-cacert"
}
