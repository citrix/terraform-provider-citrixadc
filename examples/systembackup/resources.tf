resource "citrixadc_systembackup" "tf_systembackup" {
  filename         = "my_restore_file.tgz"
}

resource "citrixadc_systembackup_create" "tf_systembackup_create" {
  filename         = "my_restore_file"
  level            = "basic"
  uselocaltimezone = "true"
}

resource "citrixadc_systembackup_restore" "tf_systembackup_restore" {
  filename   = "my_restore_file.tgz"
  skipbackup = "false"
}
