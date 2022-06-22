resource "citrixadc_interface" "test_interface" {
    interface_id = "1/1"
    hamonitor = "ON"
    autoneg = "ENABLED"
    mtu = 1500
}
