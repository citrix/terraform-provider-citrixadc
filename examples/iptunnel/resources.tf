resource "citrixadc_iptunnel" "tf_iptunnel" {
    name = "tf_iptunnel"
    remote = "66.0.0.11"
    remotesubnetmask = "255.255.255.255"
    local = "*"
}
