resource "citrixadc_lsnpool" "tf_lsnpool" {
  poolname            = "my_lsn_pool"
  nattype             = "DYNAMIC"
  portblockallocation = "DISABLED"
  maxportrealloctmq   = 50
  portrealloctimeout  = 50
}
