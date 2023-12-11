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

variable "rs_action" {
  type = map(string)
}
variable "rs_policy" {
  type = map(string)
}