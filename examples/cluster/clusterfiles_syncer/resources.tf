resource "citrixadc_clusterfiles_syncer" "syncer" {
    timestamp = timestamp()
    mode = ["all", "misc"]
}
