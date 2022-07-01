resource "citrixadc_dnstxtrec" "dnstxtrec" {
  domain = "example1.com"
  string = [
    "block",
    "log",
    "stats"
  ]
  ttl = 3600
}


