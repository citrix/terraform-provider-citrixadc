resource "citrixadc_nslicense" "license" {
    license_file = "CNS_V10000_SERVER_PLT_Retail.lic"

    ssh_host_pubkey = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDaA2H70ONYk1JDPHmqKNoOYzLZeR8jNu252P63OsI+N1k4hHQUPeysV20vzeDqgtDOoOkb90By9ryRTjGDOzxers04B23+BM+gaTFp0ONNr8uCLNt5mtZXK6dp2JjYpysl3qmpDDZ4qYhoDikliL05+bO/3dEpK6kOo25DjwjHsJDK8HovAiLdHg7v6Y6PTbJseT/+pae+0P0/gBFY901cEeB/DJqzyH7Qd1lUuUroy9buROTVhkF5VdaaPQJK8YX2oH8ocoqQOHxrSfh3U0+OuboQSyle5MnFjO88yRJrRwpT1ooJGse3xWf/0Zd5/gbuZTzswqPen2x0JN3iIvpekKItcTEegy9JlVFPEtcLeO738uYJxJuSen2HECmtl9LFjtFkLRkC5/t7qZK3SCvkKaEF/ol2K53aOPd5P9K6mYtc9xJvgtX1gntuDMuxNZBoZCeX/+5dxL0SAro9bBY0ArwpnhAo7xYgdY7F7RsXvNBJuZZiZQvFJNqnFtteKbk="

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
