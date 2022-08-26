resource "citrixadc_autoscaleprofile" "profile1" {
    name = "profile1"
    type = "CLOUDSTACK"
    apikey = "abc123"
    url = "https://1.1.1.1"
    sharedsecret = "abc123"
}