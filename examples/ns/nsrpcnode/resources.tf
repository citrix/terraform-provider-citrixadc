resource "citrixadc_nsrpcnode" "tf_nsrpcnode" {
    ipaddress = "10.78.60.201"
    password = "secret"
    secure = "ON"
    srcip = "10.78.60.201"
}
