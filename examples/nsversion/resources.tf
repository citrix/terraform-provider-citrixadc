data "citrixadc_nsversion" "nsversion" {
#installedversion = true
}

resource "citrixadc_installer" "tf_installer" {
    url =  "file:///var/tmp/build_mana_83_27_nc_64.tgz"
    y = true
    l = false
    wait_until_reachable = true

    count = local.do_install ? 1 : 0
}

locals {
    do_install = ( data.citrixadc_nsversion.nsversion.version != "Netscaler NS13.0: Build 83.27.nc" )
}
