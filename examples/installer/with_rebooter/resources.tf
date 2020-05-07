resource "citrixadc_installer" "tf_installer" {
    url =  "file:///var/tmp/build_mana_47_24_nc_64.tgz"
    y = false
    l = false
    wait_until_reachable = false
}

resource "citrixadc_rebooter" "tf_rebooter" {
  timestamp            = timestamp()
  warm                 = false
  wait_until_reachable = true
  depends_on = [ citrixadc_installer.tf_installer ]
}
