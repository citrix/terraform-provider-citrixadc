resource "citrixadc_hafailover" "ensure_secondary_is_secondary" {
    provider = citrixadc.secondary

    ipaddress = var.secondary_nsip
    state = "Secondary"
    force = true
}

resource "citrixadc_installer" "tf_installer_secondary" {
    provider = citrixadc.secondary

    url =  "file:///var/tmp/build_artesa_17_42_nc_64.tgz"
    y = true
    l = false
    wait_until_reachable = true

    depends_on = [citrixadc_hafailover.ensure_secondary_is_secondary]
}

resource "time_sleep" "wait_for_secondary_nitro" {
    create_duration = "120s"

    depends_on = [citrixadc_installer.tf_installer_secondary]
}

resource "citrixadc_hafailover" "ensure_secondary_is_primary" {
    provider = citrixadc.secondary

    ipaddress = var.secondary_nsip
    state = "Primary"
    force = true

    depends_on = [time_sleep.wait_for_secondary_nitro]
}

resource "citrixadc_installer" "tf_installer_primary" {
    provider = citrixadc.primary

    url =  "file:///var/tmp/build_artesa_17_42_nc_64.tgz"
    y = true
    l = false
    wait_until_reachable = true

    depends_on = [citrixadc_hafailover.ensure_secondary_is_primary]
}

resource "time_sleep" "wait_for_primary_nitro" {
    create_duration = "120s"

    depends_on = [citrixadc_installer.tf_installer_primary]
}

resource "citrixadc_hafailover" "ensure_primary_is_primary" {
    provider = citrixadc.primary

    ipaddress = var.primary_nsip
    state = "Primary"
    force = true

    depends_on = [time_sleep.wait_for_primary_nitro]
}
