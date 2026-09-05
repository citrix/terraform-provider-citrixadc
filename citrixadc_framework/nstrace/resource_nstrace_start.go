package nstrace

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/basic"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NstraceStartResource{}
var _ resource.ResourceWithConfigure = (*NstraceStartResource)(nil)
var _ resource.ResourceWithImportState = (*NstraceStartResource)(nil)

func NewNstraceStartResource() resource.Resource {
	return &NstraceStartResource{}
}

// NstraceStartResource defines the resource implementation.
type NstraceStartResource struct {
	client *service.NitroClient
}

// NstraceStartResourceModel describes the resource data model.
//
// nstrace is a NITRO object that supports multiple actions (start / stop) plus a
// get. Mirroring the systemscalablemgmtthreads package, each action is modelled as
// its own action-only resource. This resource wraps the `?action=start` action,
// which begins a packet trace with the configured options.
//
// There is no `set nstrace` operation (NITRO errorcode 1088), so this is a
// fire-once action: Create issues ?action=start, Read/Update are no-ops (the trace
// options are not read back into state), and Delete is a state-only removal — it
// does NOT stop the trace (use citrixadc_nstrace_stop for that). The live trace
// state is queryable via the citrixadc_nstrace data source. Attributes are
// therefore plain Optional (not Computed / not RequiresReplace): a fire-once action
// cannot be re-applied in place while a trace is running (NITRO errorcode 3984
// "One instance of nstrace is already running"), so change a running trace by
// stopping it first.
type NstraceStartResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Nf               types.Int64  `tfsdk:"nf"`
	Time             types.Int64  `tfsdk:"time"`
	Size             types.Int64  `tfsdk:"size"`
	Mode             types.List   `tfsdk:"mode"`
	Pernic           types.String `tfsdk:"pernic"`
	Filename         types.String `tfsdk:"filename"`
	Fileid           types.String `tfsdk:"fileid"`
	Filter           types.String `tfsdk:"filter"`
	Link             types.String `tfsdk:"link"`
	Nodes            types.List   `tfsdk:"nodes"`
	Filesize         types.Int64  `tfsdk:"filesize"`
	Traceformat      types.String `tfsdk:"traceformat"`
	Merge            types.String `tfsdk:"merge"`
	Doruntimecleanup types.String `tfsdk:"doruntimecleanup"`
	Tracebuffers     types.Int64  `tfsdk:"tracebuffers"`
	Skiprpc          types.String `tfsdk:"skiprpc"`
	Skiplocalssh     types.String `tfsdk:"skiplocalssh"`
	Capsslkeys       types.String `tfsdk:"capsslkeys"`
	Capdroppkt       types.String `tfsdk:"capdroppkt"`
	Inmemorytrace    types.String `tfsdk:"inmemorytrace"`
	Nodeid           types.Int64  `tfsdk:"nodeid"`
}

func (r *NstraceStartResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nstrace_start"
}

func (r *NstraceStartResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NstraceStartResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NstraceStartResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	strOpt := func(desc string) schema.StringAttribute {
		return schema.StringAttribute{Optional: true, Description: desc}
	}
	intOpt := func(desc string) schema.Int64Attribute {
		return schema.Int64Attribute{Optional: true, Description: desc}
	}
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id":   schema.StringAttribute{Computed: true, Description: "The ID of the nstrace_start resource."},
			"nf":   intOpt("Number of files to be generated in cycle."),
			"time": intOpt("Time per file (sec)."),
			"size": intOpt("Size of the captured data. Set 0 for full packet trace."),
			"mode": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "Capturing mode for trace. Any of/combination of: TX, TXB, RX, IPV6, NEW_RX, C2C, NS_FR_TX, APPFW, MPTCP, PolicyBased, HTTP_QUIC. Default: NEW_RX TXB.",
			},
			"pernic":   strOpt("Use separate trace files for each interface. Works only with cap format. Possible values = ENABLED, DISABLED"),
			"filename": strOpt("Name of the trace file."),
			"fileid":   strOpt("ID for the trace file name for uniqueness. Should be used only with -name option."),
			"filter":   strOpt("Filter expression for nstrace. Maximum length of filter is 255."),
			"link":     strOpt("Includes filtered connection's peer traffic. Possible values = ENABLED, DISABLED"),
			"nodes": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Description: "Nodes on which tracing is started.",
			},
			"filesize":         intOpt("File size, in MB, threshold for rollover. If free disk space is less than 2GB at rollover, trace stops."),
			"traceformat":      strOpt("Format in which trace will be generated. Possible values = NSCAP, PCAP"),
			"merge":            strOpt("Specify how traces across PE's are merged. Possible values = ONSTOP, ONTHEFLY, NOMERGE"),
			"doruntimecleanup": strOpt("Enable or disable runtime temp file cleanup. Possible values = ENABLED, DISABLED"),
			"tracebuffers":     intOpt("Number of 16KB trace buffers."),
			"skiprpc":          strOpt("skip RPC packets. Possible values = ENABLED, DISABLED"),
			"skiplocalssh":     strOpt("skip local SSH packets. Possible values = ENABLED, DISABLED"),
			"capsslkeys":       strOpt("Capture SSL Master keys. Not captured on FIPS machine. Possible values = ENABLED, DISABLED"),
			"capdroppkt":       strOpt("Captures Dropped Packets if set to ENABLED. Possible values = ENABLED, DISABLED"),
			"inmemorytrace":    strOpt("Logs packets in memory and dumps the trace file on stop. Possible values = ENABLED, DISABLED"),
			"nodeid":           intOpt("Unique number that identifies the cluster node."),
		},
	}
}

func (r *NstraceStartResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NstraceStartResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Starting nstrace (action-only resource)")
	payload := nstraceStartGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes start as POST ?action=start. Verb casing matches the URL.
	err := r.client.ActOnResource(service.Nstrace.Type(), &payload, "start")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to start nstrace, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Triggered nstrace start")

	// Synthetic ID for the action-only resource.
	data.Id = types.StringValue("nstrace_start")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstraceStartResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// start is a fire-once action; the options are not read back. Read is a
	// preserve-state no-op. Query live trace state via the data source instead.
	var data NstraceStartResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for nstrace_start; trace options are not read back")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstraceStartResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// nstrace has no set/update operation and a running trace cannot be restarted
	// in place (errorcode 3984). Update is a no-op that preserves the ID; to change
	// a running trace, stop it (citrixadc_nstrace_stop) and re-apply.
	var data, state NstraceStartResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for nstrace_start; NITRO has no set operation for nstrace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstraceStartResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// start is a one-shot action. Its inverse is the separate
	// citrixadc_nstrace_stop resource, so Delete simply removes the resource from
	// Terraform state (it does NOT stop a running trace).
	tflog.Debug(ctx, "Delete is a no-op for nstrace_start; use citrixadc_nstrace_stop to stop the trace")
}

func nstraceStartGetThePayloadFromthePlan(ctx context.Context, data *NstraceStartResourceModel) basic.Nstrace {
	tflog.Debug(ctx, "In nstraceStartGetThePayloadFromthePlan Function")

	nstrace := basic.Nstrace{}
	if !data.Nf.IsNull() && !data.Nf.IsUnknown() {
		nstrace.Nf = utils.IntPtr(int(data.Nf.ValueInt64()))
	}
	if !data.Time.IsNull() && !data.Time.IsUnknown() {
		nstrace.Time = utils.IntPtr(int(data.Time.ValueInt64()))
	}
	if !data.Size.IsNull() && !data.Size.IsUnknown() {
		nstrace.Size = utils.IntPtr(int(data.Size.ValueInt64()))
	}
	if !data.Mode.IsNull() && !data.Mode.IsUnknown() {
		var modeList []string
		data.Mode.ElementsAs(ctx, &modeList, false)
		nstrace.Mode = modeList
	}
	if !data.Pernic.IsNull() && !data.Pernic.IsUnknown() {
		nstrace.Pernic = data.Pernic.ValueString()
	}
	if !data.Filename.IsNull() && !data.Filename.IsUnknown() {
		nstrace.Filename = data.Filename.ValueString()
	}
	if !data.Fileid.IsNull() && !data.Fileid.IsUnknown() {
		nstrace.Fileid = data.Fileid.ValueString()
	}
	if !data.Filter.IsNull() && !data.Filter.IsUnknown() {
		nstrace.Filter = data.Filter.ValueString()
	}
	if !data.Link.IsNull() && !data.Link.IsUnknown() {
		nstrace.Link = data.Link.ValueString()
	}
	if !data.Nodes.IsNull() && !data.Nodes.IsUnknown() {
		var nodesList []int64
		data.Nodes.ElementsAs(ctx, &nodesList, false)
		intList := make([]int, len(nodesList))
		for i, v := range nodesList {
			intList[i] = int(v)
		}
		nstrace.Nodes = intList
	}
	if !data.Filesize.IsNull() && !data.Filesize.IsUnknown() {
		nstrace.Filesize = utils.IntPtr(int(data.Filesize.ValueInt64()))
	}
	if !data.Traceformat.IsNull() && !data.Traceformat.IsUnknown() {
		nstrace.Traceformat = data.Traceformat.ValueString()
	}
	if !data.Merge.IsNull() && !data.Merge.IsUnknown() {
		nstrace.Merge = data.Merge.ValueString()
	}
	if !data.Doruntimecleanup.IsNull() && !data.Doruntimecleanup.IsUnknown() {
		nstrace.Doruntimecleanup = data.Doruntimecleanup.ValueString()
	}
	if !data.Tracebuffers.IsNull() && !data.Tracebuffers.IsUnknown() {
		nstrace.Tracebuffers = utils.IntPtr(int(data.Tracebuffers.ValueInt64()))
	}
	if !data.Skiprpc.IsNull() && !data.Skiprpc.IsUnknown() {
		nstrace.Skiprpc = data.Skiprpc.ValueString()
	}
	if !data.Skiplocalssh.IsNull() && !data.Skiplocalssh.IsUnknown() {
		nstrace.Skiplocalssh = data.Skiplocalssh.ValueString()
	}
	if !data.Capsslkeys.IsNull() && !data.Capsslkeys.IsUnknown() {
		nstrace.Capsslkeys = data.Capsslkeys.ValueString()
	}
	if !data.Capdroppkt.IsNull() && !data.Capdroppkt.IsUnknown() {
		nstrace.Capdroppkt = data.Capdroppkt.ValueString()
	}
	if !data.Inmemorytrace.IsNull() && !data.Inmemorytrace.IsUnknown() {
		nstrace.Inmemorytrace = data.Inmemorytrace.ValueString()
	}
	if !data.Nodeid.IsNull() && !data.Nodeid.IsUnknown() {
		nstrace.Nodeid = utils.IntPtr(int(data.Nodeid.ValueInt64()))
	}

	return nstrace
}
