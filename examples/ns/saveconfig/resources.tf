resource "citrixadc_nsconfig_save" "tf_save" {
    timestamp  = "2020-03-24T12:37:06Z"

    # Will not error when save is already in progress
    concurrent_save_ok = true

    # Set to non zero value to retry the save config operation
    # Will throw error if limit is surpassed
    concurrent_save_retries = 1

    # Time interval between save retries
    concurrent_save_interval = "10s"

    # Total timeout for all retries
    concurrent_save_timeout = "5m"
}
