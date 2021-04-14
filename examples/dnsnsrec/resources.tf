resource "citrixadc_dnsnsrec" "tf_dnsnsrec" {
    domain = "www.test.com"
    nameserver = "192.168.1.100"
}
