resource "citrixadc_installer" "tf_installer" {
    url =  "file:///var/tmp/build_mana_47_24_nc_64.tgz"
    y = true
    l = false
    wait_until_reachable = true
}
