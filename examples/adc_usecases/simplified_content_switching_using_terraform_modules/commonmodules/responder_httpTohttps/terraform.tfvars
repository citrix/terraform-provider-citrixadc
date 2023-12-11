#Declare Responder action parameters
rs_action = {
  name   = "httptohttps_redirect"
  type   = "redirect"
  target = "\"https://\" + HTTP.REQ.HOSTNAME.HTTP_URL_SAFE + HTTP.REQ.URL.PATH_AND_QUERY.HTTP_URL_SAFE"
}
rs_policy = {
  name = "httptohttps_pol"
  rule = "HTTP.REQ.IS_VALID"
}
