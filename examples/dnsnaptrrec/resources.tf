resource "citrixadc_dnsnaptrrec" "dnsnaptrrec" {
    domain = "example.com"
    order = 10
    preference = 2
    flags = 
    services = 
    regexp = 
    replacement = 
    ttl = 3600
}