---
subcategory: "SSL"
---

# Resource: sslparameter

The sslparameter resource is used to update the ADC SSL parameters.


## Example usage

```hcl
resource "citrixadc_sslparameter" "default" {
  pushflag       = "2"
  denysslreneg   = "NONSECURE"
  defaultprofile = "ENABLED"
}
```


## Argument Reference

* `quantumsize` - (Optional) Amount of data to collect before the data is pushed to the crypto hardware for encryption. For large downloads, a larger quantum size better utilizes the crypto resources. Possible values: [ 4096, 8192, 16384 ]
* `crlmemorysizemb` - (Optional) 
* `strictcachecks` - (Optional) Enable strict CA certificate checks on the appliance. Possible values: [ YES, NO ]
* `ssltriggertimeout` - (Optional) Time, in milliseconds, after which encryption is triggered for transactions that are not tracked on the Citrix ADC because their length is not known. There can be a delay of up to 10ms from the specified timeout value before the packet is pushed into the queue.
* `sendclosenotify` - (Optional) Send an SSL Close-Notify message to the client at the end of a transaction. Possible values: [ YES, NO ]
* `encrypttriggerpktcount` - (Optional) 
* `denysslreneg` - (Optional) Deny renegotiation in specified circumstances. Available settings function as follows: * NO - Allow SSL renegotiation. * FRONTEND_CLIENT - Deny secure and nonsecure SSL renegotiation initiated by the client. * FRONTEND_CLIENTSERVER - Deny secure and nonsecure SSL renegotiation initiated by the client or the Citrix ADC during policy-based client authentication. * ALL - Deny all secure and nonsecure SSL renegotiation. * NONSECURE - Deny nonsecure SSL renegotiation. Allows only clients that support RFC 5746. Possible values: [ NO, FRONTEND_CLIENT, FRONTEND_CLIENTSERVER, ALL, NONSECURE ]
* `insertionencoding` - (Optional) Encoding method used to insert the subject or issuer's name in HTTP requests to servers. Possible values: [ Unicode, UTF-8 ]
* `ocspcachesize` - (Optional) Size, per packet engine, in megabytes, of the OCSP cache. A maximum of 10% of the packet engine memory can be assigned. Because the maximum allowed packet engine memory is 4GB, the maximum value that can be assigned to the OCSP cache is approximately 410 MB.
* `pushflag` - (Optional) Insert PUSH flag into decrypted, encrypted, or all records. If the PUSH flag is set to a value other than 0, the buffered records are forwarded on the basis of the value of the PUSH flag. Available settings function as follows: 0 - Auto (PUSH flag is not set.) 1 - Insert PUSH flag into every decrypted record. 2 -Insert PUSH flag into every encrypted record. 3 - Insert PUSH flag into every decrypted and encrypted record.
* `dropreqwithnohostheader` - (Optional) Host header check for SNI enabled sessions. If this check is enabled and the HTTP request does not contain the host header for SNI enabled sessions(i.e vserver or profile bound to vserver has SNI enabled and 'Client Hello' arrived with SNI extension), the request is dropped. Possible values: [ YES, NO ]
* `snihttphostmatch` - (Optional) Controls how the HTTP 'Host' header value is validated. These checks are performed only if the session is SNI enabled (i.e when vserver or profile bound to vserver has SNI enabled and 'Client Hello' arrived with SNI extension) and HTTP request contains 'Host' header. Available settings function as follows: CERT   - Request is forwarded if the 'Host' value is covered by the certificate used to establish this SSL session. Note: 'CERT' matching mode cannot be applied in TLS 1.3 connections established by resuming from a previous TLS 1.3 session. On these connections, 'STRICT' matching mode will be used instead. STRICT - Request is forwarded only if value of 'Host' header in HTTP is identical to the 'Server name' value passed in 'Client Hello' of the SSL connection. NO     - No validation is performed on the HTTP 'Host' header value. Possible values: [ NO, CERT, STRICT ]
* `pushenctriggertimeout` - (Optional) PUSH encryption trigger timeout value. The timeout value is applied only if you set the Push Encryption Trigger parameter to Timer in the SSL virtual server settings.
* `cryptodevdisablelimit` - (Optional) Limit to the number of disabled SSL chips after which the ADC restarts. A value of zero implies that the ADC does not automatically restart.
* `undefactioncontrol` - (Optional) Name of the undefined built-in control action: CLIENTAUTH, NOCLIENTAUTH, NOOP, RESET, or DROP.
* `undefactiondata` - (Optional) Name of the undefined built-in data action: NOOP, RESET or DROP.
* `defaultprofile` - (Optional) Global parameter used to enable default profile feature. Possible values: [ ENABLED, DISABLED ]
* `softwarecryptothreshold` - (Optional) Citrix ADC CPU utilization threshold (in percentage) beyond which crypto operations are not done in software. A value of zero implies that CPU is not utilized for doing crypto in software.
* `hybridfipsmode` - (Optional) When this mode is enabled, system will use additional crypto hardware to accelerate symmetric crypto operations. Possible values: [ ENABLED, DISABLED ]
* `sslierrorcache` - (Optional) Enable or disable dynamically learning and caching the learned information to make the subsequent interception or bypass decision. When enabled, NS does the lookup of this cached data to do early bypass. Possible values: [ ENABLED, DISABLED ]
* `sslimaxerrorcachemem` - (Optional) Specify the maximum memory that can be used for caching the learned data. This memory is used as a LRU cache so that the old entries gets replaced with new entry once the set memory limit is fully utilised. A value of 0 decides the limit automatically.
* `insertcertspace` - (Optional) To insert space between lines in the certificate header of request. Possible values: [ YES, NO ]
* `ndcppcompliancecertcheck` - (Optional) Applies when the Citrix ADC appliance acts as a client (back-end connection). Settings apply as follows: YES - During certificate verification, ignore the common name if SAN is present in the certificate. NO - Do not ignore common name. Possible values: [ YES, NO ]
* `heterogeneoussslhw` - (Optional) To support both cavium and coleto based platforms in cluster environment, this mode has to be enabled. Possible values: [ ENABLED, DISABLED ]
* `operationqueuelimit` - (Optional) Limit in percentage of capacity of the crypto operations queue beyond which new SSL connections are not accepted until the queue is reduced.
* `sigdigesttype` - (Optional) Signature Digest Algorithms that are supported by appliance. Default value is "ALL" and it will enable the following algorithms depending on the platform. On VPX: ECDSA-SHA1 ECDSA-SHA224 ECDSA-SHA256 ECDSA-SHA384 ECDSA-SHA512 RSA-SHA1 RSA-SHA224 RSA-SHA256 RSA-SHA384 RSA-SHA512 DSA-SHA1 DSA-SHA224 DSA-SHA256 DSA-SHA384 DSA-SHA512 On MPX with Nitrox-III and coleto cards: RSA-SHA1 RSA-SHA224 RSA-SHA256 RSA-SHA384 RSA-SHA512 ECDSA-SHA1 ECDSA-SHA224 ECDSA-SHA256 ECDSA-SHA384 ECDSA-SHA512 Others: RSA-SHA1 RSA-SHA224 RSA-SHA256 RSA-SHA384 RSA-SHA512. Note:ALL doesnot include RSA-MD5 for any platform.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslparameter. It is a unique string prefixed with "tf-sslparameter-"
