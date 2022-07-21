
locals {
    do_install = ( data.citrixadc_nsversion.nsversion.version != "Netscaler NS13.1: Build 17.42.nc" )
}

data "citrixadc_nsversion" "nsversion" {
    provider = citrixadc.primary
    installedversion = true
}

resource "citrixadc_hafailover" "ensure_secondary_is_secondary" {
    provider = citrixadc.secondary

    ipaddress = var.secondary_nsip
    state = "Secondary"
    force = true

    count = local.do_install ? 1 : 0
}

resource "citrixadc_installer" "tf_installer_secondary" {
    provider = citrixadc.secondary

    url =  "file:///var/tmp/build_artesa_17_42_nc_64.tgz"
    y = true
    l = false
    wait_until_reachable = true

    depends_on = [citrixadc_hafailover.ensure_secondary_is_secondary]

    count = local.do_install ? 1 : 0
}

resource "time_sleep" "wait_for_secondary_nitro" {
    create_duration = "120s"

    depends_on = [citrixadc_installer.tf_installer_secondary]

    count = local.do_install ? 1 : 0
}

resource "citrixadc_hafailover" "ensure_secondary_is_primary" {
    provider = citrixadc.secondary

    ipaddress = var.secondary_nsip
    state = "Primary"
    force = true

    depends_on = [time_sleep.wait_for_secondary_nitro]

    count = local.do_install ? 1 : 0
}

resource "citrixadc_installer" "tf_installer_primary" {
    provider = citrixadc.primary

    url =  "file:///var/tmp/build_artesa_17_42_nc_64.tgz"
    y = true
    l = false
    wait_until_reachable = true

    depends_on = [citrixadc_hafailover.ensure_secondary_is_primary]

    count = local.do_install ? 1 : 0
}

resource "time_sleep" "wait_for_primary_nitro" {
    create_duration = "120s"

    depends_on = [citrixadc_installer.tf_installer_primary]

    count = local.do_install ? 1 : 0
}

resource "citrixadc_hafailover" "ensure_primary_is_primary" {
    provider = citrixadc.primary

    ipaddress = var.primary_nsip
    state = "Primary"
    force = true

    depends_on = [time_sleep.wait_for_primary_nitro]

    count = local.do_install ? 1 : 0
}
