package lbmetrictable

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/lb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LbmetrictableResourceModel describes the resource data model.
// Backward-compat: the SDK v2 resource (citrixadc/resource_citrixadc_lbmetrictable.go)
// exposed only the "metrictable" attribute (Required, ForceNew). The metric/SNMPOID
// pairs belong to the separate lbmetrictable_metric_binding resource, so they are not
// part of this resource's contract.
type LbmetrictableResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Metrictable types.String `tfsdk:"metrictable"`
}

func (r *LbmetrictableResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lbmetrictable resource.",
			},
			"metrictable": schema.StringAttribute{
				Required: true,
				// SDK v2 ForceNew: true -> RequiresReplace()
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the metric table. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my metrictable\" or 'my metrictable').",
			},
		},
	}
}

func lbmetrictableGetThePayloadFromthePlan(ctx context.Context, data *LbmetrictableResourceModel) lb.Lbmetrictable {
	tflog.Debug(ctx, "In lbmetrictableGetThePayloadFromthePlan Function")

	// Create API request body from the model
	lbmetrictable := lb.Lbmetrictable{}
	if !data.Metrictable.IsNull() && !data.Metrictable.IsUnknown() {
		lbmetrictable.Metrictable = data.Metrictable.ValueString()
	}

	return lbmetrictable
}

func lbmetrictableSetAttrFromGet(ctx context.Context, data *LbmetrictableResourceModel, getResponseData map[string]interface{}) *LbmetrictableResourceModel {
	tflog.Debug(ctx, "In lbmetrictableSetAttrFromGet Function")

	// Convert API response to model.
	// metrictable is the unique name/key; only adopt the GET value when present so a
	// configured/state value is never clobbered.
	if val, ok := getResponseData["metrictable"]; ok && val != nil {
		data.Metrictable = types.StringValue(val.(string))
	}

	// Set ID for the resource.
	// Case 2: Single unique attribute - use plain value as ID.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Metrictable.ValueString()))

	return data
}
