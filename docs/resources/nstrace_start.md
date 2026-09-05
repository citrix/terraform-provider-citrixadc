---
subcategory: "Basic"
---

# Resource: nstrace_start

This resource starts an nstrace packet capture on the Citrix ADC with the configured options.

Starting and stopping a trace are separate action resources: `citrixadc_nstrace_start` and `citrixadc_nstrace_stop`. Removing this resource from your configuration does **not** stop a running trace — use `citrixadc_nstrace_stop` to stop it. Use the `citrixadc_nstrace` data source to read the live trace state.

~> **NOTE:** Only one trace can run at a time; a trace cannot be started while another is already running. Stop the running trace first.


## Example usage

```hcl
resource "citrixadc_nstrace_start" "capture" {
  nf          = 24
  time        = 3600
  size        = 0
  mode        = ["NEW_RX", "TXB"]
  filename    = "mytrace"
  traceformat = "NSCAP"
  filesize    = 1024
}
```


## Argument Reference

All arguments are optional (the appliance applies its own defaults for those omitted):

* `nf` - Number of files to be generated in cycle.
* `time` - Time per file (sec).
* `size` - Size of the captured data. Set 0 for full packet trace.
* `mode` - (List of String) Capturing mode(s): TX, TXB, RX, IPV6, NEW_RX, C2C, NS_FR_TX, APPFW, MPTCP, PolicyBased, HTTP_QUIC.
* `pernic` - Use separate trace files for each interface (cap format only). ENABLED/DISABLED.
* `filename` - Name of the trace file.
* `fileid` - ID for the trace file name for uniqueness (use only with `filename`).
* `filter` - Filter expression for nstrace (max length 255).
* `link` - Include filtered connection's peer traffic. ENABLED/DISABLED.
* `nodes` - (List of Number) Nodes on which tracing is started.
* `filesize` - File size (MB) threshold for rollover.
* `traceformat` - NSCAP or PCAP.
* `merge` - How traces across PEs are merged: ONSTOP, ONTHEFLY, NOMERGE.
* `doruntimecleanup` - Runtime temp file cleanup. ENABLED/DISABLED.
* `tracebuffers` - Number of 16KB trace buffers.
* `skiprpc` - Skip RPC packets. ENABLED/DISABLED.
* `skiplocalssh` - Skip local SSH packets. ENABLED/DISABLED.
* `capsslkeys` - Capture SSL master keys. ENABLED/DISABLED.
* `capdroppkt` - Capture dropped packets. ENABLED/DISABLED.
* `inmemorytrace` - Log packets in memory, dump on stop. ENABLED/DISABLED.
* `nodeid` - Cluster node id.


## Attribute Reference

* `id` - The id of the nstrace_start resource. It is set to `nstrace_start`.
