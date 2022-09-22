---
subcategory: "LSN"
---

# Resource: lsntransportprofile

The lsntransportprofile resource is used to create lsntransportprofile.


## Example usage

```hcl
resource "citrixadc_lsntransportprofile" "tf_lsntransportprofile" {
  transportprofilename = "my_lsn_transportprofile"
  transportprotocol    = "TCP"
  portquota            = 10
  sessionquota         = 10
  groupsessionlimit    = 1000
}
```


## Argument Reference

* `transportprofilename` - (Required) Name for the LSN transport profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN transport profile is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "lsn transport profile1" or 'lsn transport profile1').
* `transportprotocol` - (Required) Protocol for which to set the LSN transport profile parameters.
* `finrsttimeout` - (Optional) Timeout, in seconds, for a TCP LSN session after a FIN or RST message is received from one of the endpoints.  If a TCP LSN session is idle (after the Citrix ADC receives a FIN or RST message) for a time that exceeds this value, the Citrix ADC ADC removes the session.  Since the LSN feature of the Citrix ADC does not maintain state information of any TCP LSN sessions, this timeout accommodates the transmission of the FIN or RST, and ACK messages from the other endpoint so that both endpoints can properly close the connection.
* `groupsessionlimit` - (Optional) Maximum number of concurrent LSN sessions(for the specified protocol) allowed for all subscriber of a group to which this profile has bound. This limit will get split across the Citrix ADCs packet engines and rounded down. When the number of LSN sessions reaches the limit for a group in packet engine, the Citrix ADC does not allow the subscriber of that group to open additional sessions through that packet engine.
* `portpreserveparity` - (Optional) Enable port parity between a subscriber port and its mapped LSN NAT port. For example, if a subscriber initiates a connection from an odd numbered port, the Citrix ADC allocates an odd numbered LSN NAT port for this connection.  You must set this parameter for proper functioning of protocols that require the source port to be even or odd numbered, for example, in peer-to-peer applications that use RTP or RTCP protocol.
* `portpreserverange` - (Optional) If a subscriber initiates a connection from a well-known port (0-1023), allocate a NAT port from the well-known port range (0-1023) for this connection. For example, if a subscriber initiates a connection from port 80, the Citrix ADC can allocate port 100 as the NAT port for this connection.  This parameter applies to dynamic NAT without port block allocation. It also applies to Deterministic NAT if the range of ports allocated includes well-known ports.  When all the well-known ports of all the available NAT IP addresses are used in different subscriber's connections (LSN sessions), and a subscriber initiates a connection from a well-known port, the Citrix ADC drops this connection.
* `portquota` - (Optional) Maximum number of LSN NAT ports to be used at a time by each subscriber for the specified protocol. For example, each subscriber can be limited to a maximum of 500 TCP NAT ports. When the LSN NAT mappings for a subscriber reach the limit, the Citrix ADC does not allocate additional NAT ports for that subscriber.
* `sessionquota` - (Optional) Maximum number of concurrent LSN sessions allowed for each subscriber for the specified protocol.  When the number of LSN sessions reaches the limit for a subscriber, the Citrix ADC does not allow the subscriber to open additional sessions.
* `sessiontimeout` - (Optional) Timeout, in seconds, for an idle LSN session. If an LSN session is idle for a time that exceeds this value, the Citrix ADC removes the session.  This timeout does not apply for a TCP LSN session when a FIN or RST message is received from either of the endpoints.
* `stuntimeout` - (Optional) STUN protocol timeout
* `syncheck` - (Optional) Silently drop any non-SYN packets for connections for which there is no LSN-NAT session present on the Citrix ADC.   If you disable this parameter, the Citrix ADC accepts any non-SYN packets and creates a new LSN session entry for this connection.   Following are some reasons for the Citrix ADC to receive such packets:  * LSN session for a connection existed but the Citrix ADC removed this session because the LSN session was idle for a time that exceeded the configured session timeout. * Such packets can be a part of a DoS attack.
* `synidletimeout` - (Optional) SYN Idle timeout


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lsntransportprofile. It has the same value as the `transportprofilename` attribute.


## Import

A lsntransportprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_lsntransportprofile.tf_lsntransportprofile my_lsn_transportprofile
```
