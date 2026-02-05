package nspartition

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NspartitionResourceModel describes the resource data model.
type NspartitionResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Force         types.Bool   `tfsdk:"force"`
	Maxbandwidth  types.Int64  `tfsdk:"maxbandwidth"`
	Maxconn       types.Int64  `tfsdk:"maxconn"`
	Maxmemlimit   types.Int64  `tfsdk:"maxmemlimit"`
	Minbandwidth  types.Int64  `tfsdk:"minbandwidth"`
	Partitionmac  types.String `tfsdk:"partitionmac"`
	Partitionname types.String `tfsdk:"partitionname"`
	Save          types.Bool   `tfsdk:"save"`
}

func (r *NspartitionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nspartition resource.",
			},
			"force": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Switches to new admin partition without prompt for saving configuration. Configuration will not be saved",
			},
			"maxbandwidth": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(10240),
				Description: "Maximum bandwidth, in Kbps, that the partition can consume. A zero value indicates the bandwidth is unrestricted on the partition and it can consume up to the system limits.",
			},
			"maxconn": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1024),
				Description: "Maximum number of concurrent connections that can be open in the partition. A zero value indicates no limit on number of open connections.",
			},
			"maxmemlimit": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(10),
				Description: "Maximum memory, in megabytes, allocated to the partition.  A zero value indicates the memory is unlimited on the partition and it can consume up to the system limits.",
			},
			"minbandwidth": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(10240),
				Description: "Minimum bandwidth, in Kbps, that the partition can consume. A zero value indicates the bandwidth is unrestricted on the partition and it can consume up to the system limits",
			},
			"partitionmac": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Special MAC address for the partition which is used for communication over shared vlans in this partition. If not specified, the MAC address is auto-generated.",
			},
			"partitionname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the Partition. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},
			"save": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Switches to new admin partition without prompt for saving configuration. Configuration will be saved",
			},
		},
	}
}

func nspartitionGetThePayloadFromtheConfig(ctx context.Context, data *NspartitionResourceModel) ns.Nspartition {
	tflog.Debug(ctx, "In nspartitionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nspartition := ns.Nspartition{}
	if !data.Force.IsNull() {
		nspartition.Force = data.Force.ValueBool()
	}
	if !data.Maxbandwidth.IsNull() {
		nspartition.Maxbandwidth = utils.IntPtr(int(data.Maxbandwidth.ValueInt64()))
	}
	if !data.Maxconn.IsNull() {
		nspartition.Maxconn = utils.IntPtr(int(data.Maxconn.ValueInt64()))
	}
	if !data.Maxmemlimit.IsNull() {
		nspartition.Maxmemlimit = utils.IntPtr(int(data.Maxmemlimit.ValueInt64()))
	}
	if !data.Minbandwidth.IsNull() {
		nspartition.Minbandwidth = utils.IntPtr(int(data.Minbandwidth.ValueInt64()))
	}
	if !data.Partitionmac.IsNull() {
		nspartition.Partitionmac = data.Partitionmac.ValueString()
	}
	if !data.Partitionname.IsNull() {
		nspartition.Partitionname = data.Partitionname.ValueString()
	}
	if !data.Save.IsNull() {
		nspartition.Save = data.Save.ValueBool()
	}

	return nspartition
}

func nspartitionSetAttrFromGet(ctx context.Context, data *NspartitionResourceModel, getResponseData map[string]interface{}) *NspartitionResourceModel {
	tflog.Debug(ctx, "In nspartitionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["force"]; ok && val != nil {
		data.Force = types.BoolValue(val.(bool))
	} else {
		data.Force = types.BoolNull()
	}
	if val, ok := getResponseData["maxbandwidth"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxbandwidth = types.Int64Value(intVal)
		}
	} else {
		data.Maxbandwidth = types.Int64Null()
	}
	if val, ok := getResponseData["maxconn"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxconn = types.Int64Value(intVal)
		}
	} else {
		data.Maxconn = types.Int64Null()
	}
	if val, ok := getResponseData["maxmemlimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxmemlimit = types.Int64Value(intVal)
		}
	} else {
		data.Maxmemlimit = types.Int64Null()
	}
	if val, ok := getResponseData["minbandwidth"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Minbandwidth = types.Int64Value(intVal)
		}
	} else {
		data.Minbandwidth = types.Int64Null()
	}
	if val, ok := getResponseData["partitionmac"]; ok && val != nil {
		data.Partitionmac = types.StringValue(val.(string))
	} else {
		data.Partitionmac = types.StringNull()
	}
	if val, ok := getResponseData["partitionname"]; ok && val != nil {
		data.Partitionname = types.StringValue(val.(string))
	} else {
		data.Partitionname = types.StringNull()
	}
	if val, ok := getResponseData["save"]; ok && val != nil {
		data.Save = types.BoolValue(val.(bool))
	} else {
		data.Save = types.BoolNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Partitionname.ValueString())

	return data
}
