package nspartition

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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
			// SDK v2: Optional + Computed, NOT ForceNew.
			"force": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Switches to new admin partition without prompt for saving configuration. Configuration will not be saved",
			},
			// SDK v2: Optional + Computed, no Default (value read from ADC).
			"maxbandwidth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum bandwidth, in Kbps, that the partition can consume. A zero value indicates the bandwidth is unrestricted on the partition and it can consume up to the system limits.",
			},
			"maxconn": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of concurrent connections that can be open in the partition. A zero value indicates no limit on number of open connections.",
			},
			"maxmemlimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum memory, in megabytes, allocated to the partition.  A zero value indicates the memory is unlimited on the partition and it can consume up to the system limits.",
			},
			"minbandwidth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum bandwidth, in Kbps, that the partition can consume. A zero value indicates the bandwidth is unrestricted on the partition and it can consume up to the system limits",
			},
			"partitionmac": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Special MAC address for the partition which is used for communication over shared vlans in this partition. If not specified, the MAC address is auto-generated.",
			},
			// SDK v2: Required + ForceNew -> RequiresReplace.
			"partitionname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the Partition. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},
			// SDK v2: Optional + Computed, NOT ForceNew.
			"save": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Switches to new admin partition without prompt for saving configuration. Configuration will be saved",
			},
		},
	}
}

func nspartitionGetThePayloadFromthePlan(ctx context.Context, data *NspartitionResourceModel) ns.Nspartition {
	tflog.Debug(ctx, "In nspartitionGetThePayloadFromthePlan Function")

	// Create API request body from the model
	nspartition := ns.Nspartition{}
	if !data.Force.IsNull() && !data.Force.IsUnknown() {
		nspartition.Force = data.Force.ValueBool()
	}
	if !data.Maxbandwidth.IsNull() && !data.Maxbandwidth.IsUnknown() {
		nspartition.Maxbandwidth = utils.IntPtr(int(data.Maxbandwidth.ValueInt64()))
	}
	if !data.Maxconn.IsNull() && !data.Maxconn.IsUnknown() {
		nspartition.Maxconn = utils.IntPtr(int(data.Maxconn.ValueInt64()))
	}
	if !data.Maxmemlimit.IsNull() && !data.Maxmemlimit.IsUnknown() {
		nspartition.Maxmemlimit = utils.IntPtr(int(data.Maxmemlimit.ValueInt64()))
	}
	if !data.Minbandwidth.IsNull() && !data.Minbandwidth.IsUnknown() {
		nspartition.Minbandwidth = utils.IntPtr(int(data.Minbandwidth.ValueInt64()))
	}
	if !data.Partitionmac.IsNull() && !data.Partitionmac.IsUnknown() {
		nspartition.Partitionmac = data.Partitionmac.ValueString()
	}
	if !data.Partitionname.IsNull() && !data.Partitionname.IsUnknown() {
		nspartition.Partitionname = data.Partitionname.ValueString()
	}
	if !data.Save.IsNull() && !data.Save.IsUnknown() {
		nspartition.Save = data.Save.ValueBool()
	}

	return nspartition
}

func nspartitionSetAttrFromGet(ctx context.Context, data *NspartitionResourceModel, getResponseData map[string]interface{}) *NspartitionResourceModel {
	tflog.Debug(ctx, "In nspartitionSetAttrFromGet Function")

	// Convert API response to model.
	// Note: NITRO GET does not echo the action-only flags "force" and "save",
	// so guard the else-branch to only null when the value is unknown; never
	// clobber a known configured value (omit-on-default trap).
	if val, ok := getResponseData["force"]; ok && val != nil {
		data.Force = types.BoolValue(val.(bool))
	} else if data.Force.IsUnknown() {
		data.Force = types.BoolNull()
	}
	if val, ok := getResponseData["maxbandwidth"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxbandwidth = types.Int64Value(intVal)
		}
	} else if data.Maxbandwidth.IsUnknown() {
		data.Maxbandwidth = types.Int64Null()
	}
	if val, ok := getResponseData["maxconn"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxconn = types.Int64Value(intVal)
		}
	} else if data.Maxconn.IsUnknown() {
		data.Maxconn = types.Int64Null()
	}
	if val, ok := getResponseData["maxmemlimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxmemlimit = types.Int64Value(intVal)
		}
	} else if data.Maxmemlimit.IsUnknown() {
		data.Maxmemlimit = types.Int64Null()
	}
	if val, ok := getResponseData["minbandwidth"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Minbandwidth = types.Int64Value(intVal)
		}
	} else if data.Minbandwidth.IsUnknown() {
		data.Minbandwidth = types.Int64Null()
	}
	if val, ok := getResponseData["partitionmac"]; ok && val != nil {
		data.Partitionmac = types.StringValue(val.(string))
	} else if data.Partitionmac.IsUnknown() {
		data.Partitionmac = types.StringNull()
	}
	if val, ok := getResponseData["partitionname"]; ok && val != nil {
		data.Partitionname = types.StringValue(val.(string))
	} else if data.Partitionname.IsUnknown() {
		data.Partitionname = types.StringNull()
	}
	if val, ok := getResponseData["save"]; ok && val != nil {
		data.Save = types.BoolValue(val.(bool))
	} else if data.Save.IsUnknown() {
		data.Save = types.BoolNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain partitionname value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Partitionname.ValueString()))

	return data
}
