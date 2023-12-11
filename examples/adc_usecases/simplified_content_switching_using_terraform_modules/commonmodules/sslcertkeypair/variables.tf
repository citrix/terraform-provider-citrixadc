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

variable "ssl_certkey_name" {
  type = string
}
variable "ssl_certificate_path" {
  type = string
}
variable "ssl_key_path" {
  type = string
}
variable "filename" {
  type = map(string)
}




