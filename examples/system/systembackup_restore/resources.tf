resource "citrixadc_systembackup_restore" "tf_systembackup_restore" {
  filename   = "my_restore_file.tgz"
  skipbackup = "true"
}
