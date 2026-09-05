package cloudtrafficroutes

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/cloud"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// unsetOnRemoveStringModifier forces the planned value to unknown when the user
// removes a previously-set attribute from configuration while a non-empty value
// still exists in prior state. This makes Terraform detect a change (unknown !=
// prior) and call Update, which issues the NITRO ?action=unset. Without it an
// Optional+Computed attribute is "sticky": the prior value is carried forward
// and removal is a silent no-op. Because these attributes revert to no value
// (absent from GET) after unset, marking the plan unknown also avoids a
// "provider produced inconsistent result" error, which a static Default would
// trigger.
type unsetOnRemoveStringModifier struct{}

func (m unsetOnRemoveStringModifier) Description(_ context.Context) string {
	return "Marks the value unknown when removed from config while a prior non-empty value exists, so it is unset on the appliance."
}

func (m unsetOnRemoveStringModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m unsetOnRemoveStringModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.StateValue.IsNull() {
		return
	}
	if req.ConfigValue.IsNull() && req.StateValue.ValueString() != "" {
		resp.PlanValue = types.StringUnknown()
	}
}

// CloudtrafficroutesResourceModel describes the resource data model.
type CloudtrafficroutesResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Name             types.String `tfsdk:"name"`
	Targetvpcnetwork types.String `tfsdk:"targetvpcnetwork"`
	Destrange        types.String `tfsdk:"destrange"`
	Nexthopip        types.String `tfsdk:"nexthopip"`
	Ownernode        types.Int64  `tfsdk:"ownernode"`
}

func (r *CloudtrafficroutesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cloudtrafficroutes resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the traffic cloud route.",
			},
			"targetvpcnetwork": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "Target VPC network name.",
			},
			"destrange": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "Destination IP range in CIDR format.",
			},
			"nexthopip": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "Next hop IP address.",
			},
			"ownernode": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "cluster owner node id for the nexthopipaddress.",
			},
		},
	}
}

func cloudtrafficroutesGetThePayloadFromthePlan(ctx context.Context, data *CloudtrafficroutesResourceModel) cloud.Cloudtrafficroutes {
	tflog.Debug(ctx, "In cloudtrafficroutesGetThePayloadFromthePlan Function")

	// Create API request body from the model
	cloudtrafficroutes := cloud.Cloudtrafficroutes{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		cloudtrafficroutes.Name = data.Name.ValueString()
	}
	if !data.Targetvpcnetwork.IsNull() && !data.Targetvpcnetwork.IsUnknown() {
		cloudtrafficroutes.Targetvpcnetwork = data.Targetvpcnetwork.ValueString()
	}
	if !data.Destrange.IsNull() && !data.Destrange.IsUnknown() {
		cloudtrafficroutes.Destrange = data.Destrange.ValueString()
	}
	if !data.Nexthopip.IsNull() && !data.Nexthopip.IsUnknown() {
		cloudtrafficroutes.Nexthopip = data.Nexthopip.ValueString()
	}
	if !data.Ownernode.IsNull() && !data.Ownernode.IsUnknown() {
		cloudtrafficroutes.Ownernode = utils.IntPtr(int(data.Ownernode.ValueInt64()))
	}

	return cloudtrafficroutes
}

func cloudtrafficroutesSetAttrFromGet(ctx context.Context, data *CloudtrafficroutesResourceModel, getResponseData map[string]interface{}) *CloudtrafficroutesResourceModel {
	tflog.Debug(ctx, "In cloudtrafficroutesSetAttrFromGet Function")

	// Convert API response to model.
	// For Optional+Computed attributes whose NITRO value is omitted from GET, only
	// null the attribute when it was unknown (unset in config); otherwise preserve
	// the configured/state value to avoid "inconsistent result after apply".
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else if data.Name.IsUnknown() {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["targetvpcnetwork"]; ok && val != nil {
		data.Targetvpcnetwork = types.StringValue(val.(string))
	} else if data.Targetvpcnetwork.IsUnknown() {
		data.Targetvpcnetwork = types.StringNull()
	}
	if val, ok := getResponseData["destrange"]; ok && val != nil {
		data.Destrange = types.StringValue(val.(string))
	} else if data.Destrange.IsUnknown() {
		data.Destrange = types.StringNull()
	}
	if val, ok := getResponseData["nexthopip"]; ok && val != nil {
		data.Nexthopip = types.StringValue(val.(string))
	} else if data.Nexthopip.IsUnknown() {
		data.Nexthopip = types.StringNull()
	}
	if val, ok := getResponseData["ownernode"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ownernode = types.Int64Value(intVal)
		}
	} else if data.Ownernode.IsUnknown() {
		data.Ownernode = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute (name) - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
