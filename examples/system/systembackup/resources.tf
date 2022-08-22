resource "citrixadc_systembackup" "tf_systembackup" {
  filename         = "my_restore_file"
  level            = "basic"
  uselocaltimezone = "true"
}
