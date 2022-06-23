resource "citrixadc_dnstxtrec" "tf_dnstxtrec" {
  domain = "example.com"
  string = [
    "block",
    "log",
    "stats"
  ]
  ttl = "3600"
}


