---
subcategory: "Basic"
---

# Data Source: nstrace

The nstrace data source allows you to retrieve information about the current
nstrace (packet trace) operation on the appliance.


## Example usage

```terraform
data "citrixadc_nstrace" "tf_nstrace" {
}

output "filename" {
  value = data.citrixadc_nstrace.tf_nstrace.filename
}

output "traceformat" {
  value = data.citrixadc_nstrace.tf_nstrace.traceformat
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `nf` - Number of files to be generated in cycle.
* `time` - Time per file (sec).
* `size` - Size of the captured data. 0 means full packet trace.
* `mode` - Capturing mode for trace.
* `pernic` - Whether separate trace files are used for each interface.
* `filename` - Name of the trace file.
* `fileid` - ID for the trace file name for uniqueness.
* `filter` - Filter expression for nstrace.
* `link` - Whether filtered connection's peer traffic is included.
* `nodes` - Nodes on which tracing is started.
* `filesize` - File size, in MB, threshold for rollover.
* `traceformat` - Format in which trace is generated.
* `merge` - How traces across PE's are merged.
* `doruntimecleanup` - Whether runtime temp file cleanup is enabled.
* `tracebuffers` - Number of 16KB trace buffers.
* `skiprpc` - Whether RPC packets are skipped.
* `skiplocalssh` - Whether local SSH packets are skipped.
* `capsslkeys` - Whether SSL Master keys are captured.
* `capdroppkt` - Whether dropped packets are captured.
* `inmemorytrace` - Whether packets are logged in the appliance's memory.
* `nodeid` - Unique number that identifies the cluster node.
* `id` - The id of the nstrace. It is a system-generated identifier.
