resource "citrixadc_nslicense" "license" {
    license_file = "CNS_V10000_SERVER_PLT_Retail.lic"

    # All options below are shown with the default values
    # You can omit any parameters for which the default value is acceptable

    # Reboot after uploading the license file
    # Alter this to false if you intend to do a manual reboot later
    reboot = true

    # All times below are relevant only when reboot == true

    # Initial delay after reboot before starting polling target ADC
    poll_delay = "30s"

    # Interval between polls
    poll_interval = "30s"

    # How long to wait for a poll connection timeout
    poll_timeout = "10s"

    # Total time to wait for each operation
    timeouts {
        create = "20m"
        update = "20m"
        delete = "20m"
    }
}
