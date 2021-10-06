resource "citrixadc_dnsparameter" "tf_dnsparameter" {
  cacheecszeroprefix         = "ENABLED"
  cachehitbypass             = "DISABLED"
  cachenoexpire              = "DISABLED"
  cacherecords               = "YES"
  dns64timeout               = 1000
  dnsrootreferral            = "DISABLED"
  dnssec                     = "ENABLED"
  ecsmaxsubnets              = 0
  maxcachesize               = 0
  maxnegativecachesize       = 0
  maxnegcachettl             = 604800
  maxpipeline                = 255
  maxttl                     = 604800
  maxudppacketsize           = 1280
  minttl                     = 0
  namelookuppriority         = "WINS"
  nxdomainratelimitthreshold = 0
  recursion                  = "DISABLED"
  resolutionorder            = "OnlyAQuery"
  retries                    = 5
  splitpktqueryprocessing    = "ALLOW"
}
