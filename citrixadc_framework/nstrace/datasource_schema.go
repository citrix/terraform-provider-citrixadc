package nstrace

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NstraceDataSourceModel describes the data source data model.
//
// The start/stop actions are modelled as separate action-only resources
// (resource_nstrace_start.go / resource_nstrace_stop.go). This data source reads
// the live trace configuration/state via the NITRO get.
type NstraceDataSourceModel struct {
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

func NstraceDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"nf": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of files to be generated in cycle.",
			},
			"time": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time per file (sec).",
			},
			"size": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Size of the captured data. Set 0 for full packet trace.",
			},
			"mode": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Capturing mode for trace.",
			},
			"pernic": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use separate trace files for each interface. Works only with cap format.",
			},
			"filename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the trace file.",
			},
			"fileid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ID for the trace file name for uniqueness.",
			},
			"filter": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Filter expression for nstrace.",
			},
			"link": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Includes filtered connection's peer traffic.",
			},
			"nodes": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "Nodes on which tracing is started.",
			},
			"filesize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "File size, in MB, treshold for rollover.",
			},
			"traceformat": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Format in which trace will be generated.",
			},
			"merge": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify how traces across PE's are merged.",
			},
			"doruntimecleanup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable runtime temp file cleanup.",
			},
			"tracebuffers": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of 16KB trace buffers.",
			},
			"skiprpc": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "skip RPC packets.",
			},
			"skiplocalssh": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "skip local SSH packets.",
			},
			"capsslkeys": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Capture SSL Master keys.",
			},
			"capdroppkt": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Captures Dropped Packets if set to ENABLED.",
			},
			"inmemorytrace": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Logs packets in appliance's memory and dumps the trace file on stopping the nstrace operation.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
		},
	}
}

// nstraceSetAttrFromGet maps a NITRO get response onto the data source model.
func nstraceSetAttrFromGet(ctx context.Context, data *NstraceDataSourceModel, getResponseData map[string]interface{}) *NstraceDataSourceModel {
	tflog.Debug(ctx, "In nstraceSetAttrFromGet Function")

	setInt := func(key string, dst *types.Int64) {
		if val, ok := getResponseData[key]; ok && val != nil {
			if intVal, err := utils.ConvertToInt64(val); err == nil {
				*dst = types.Int64Value(intVal)
				return
			}
		}
		*dst = types.Int64Null()
	}
	setStr := func(key string, dst *types.String) {
		if val, ok := getResponseData[key]; ok && val != nil {
			if s, ok := val.(string); ok {
				*dst = types.StringValue(s)
				return
			}
		}
		*dst = types.StringNull()
	}

	setInt("nf", &data.Nf)
	setInt("time", &data.Time)
	setInt("size", &data.Size)
	if val, ok := getResponseData["mode"]; ok && val != nil {
		if sliceVal, ok := val.([]interface{}); ok {
			lv, _ := types.ListValueFrom(ctx, types.StringType, utils.ToStringList(sliceVal))
			data.Mode = lv
		} else {
			data.Mode = types.ListNull(types.StringType)
		}
	} else {
		data.Mode = types.ListNull(types.StringType)
	}
	setStr("pernic", &data.Pernic)
	setStr("filename", &data.Filename)
	setStr("fileid", &data.Fileid)
	setStr("filter", &data.Filter)
	setStr("link", &data.Link)
	if val, ok := getResponseData["nodes"]; ok && val != nil {
		if sliceVal, ok := val.([]interface{}); ok {
			int64List := make([]int64, 0, len(sliceVal))
			for _, item := range sliceVal {
				if intVal, err := utils.ConvertToInt64(item); err == nil {
					int64List = append(int64List, intVal)
				}
			}
			lv, _ := types.ListValueFrom(ctx, types.Int64Type, int64List)
			data.Nodes = lv
		} else {
			data.Nodes = types.ListNull(types.Int64Type)
		}
	} else {
		data.Nodes = types.ListNull(types.Int64Type)
	}
	setInt("filesize", &data.Filesize)
	setStr("traceformat", &data.Traceformat)
	setStr("merge", &data.Merge)
	setStr("doruntimecleanup", &data.Doruntimecleanup)
	setInt("tracebuffers", &data.Tracebuffers)
	setStr("skiprpc", &data.Skiprpc)
	setStr("skiplocalssh", &data.Skiplocalssh)
	setStr("capsslkeys", &data.Capsslkeys)
	setStr("capdroppkt", &data.Capdroppkt)
	setStr("inmemorytrace", &data.Inmemorytrace)
	setInt("nodeid", &data.Nodeid)

	// Singleton - static ID.
	data.Id = types.StringValue("nstrace-config")

	return data
}
