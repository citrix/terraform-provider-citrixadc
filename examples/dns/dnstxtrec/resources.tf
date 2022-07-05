resource "citrixadc_dnstxtrec" "dnstxtrec" {
  domain = "asoighewgoadfa.net"
  string = [
                "v=spf1 a mxrec include:websitewelcome.com ~all"
            ]
  ttl = 3600
}


