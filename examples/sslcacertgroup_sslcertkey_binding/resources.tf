resource "citrixadc_sslcacertgroup_sslcertkey_binding" "sslcacertgroup_sslcertkey_binding_demo" {	
    cacertgroupname = citrixadc_sslcacertgroup.ns_callout_certs1.cacertgroupname
    certkeyname = citrixadc_sslcertkey.tf_cacertkey.certkey
    ocspcheck = "Mandatory"
}

resource "citrixadc_sslcertkey" "tf_sslcertkey" {
    certkey = "tf_sslcertkey"
    cert = "/var/tmp/certificate1.crt"
    key = "/var/tmp/key1.pem"
    notificationperiod = 40
    expirymonitor = "ENABLED"
}

resource "citrixadc_sslcertkey" "tf_cacertkey" {
    certkey = "tf_cacertkey"
    cert = "/var/tmp/ca.crt"
}
    
resource "citrixadc_sslcacertgroup" "ns_callout_certs1" {
    cacertgroupname = "ns_callout_certs1"
}
