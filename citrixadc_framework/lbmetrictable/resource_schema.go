package lbmetrictable

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/lb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LbmetrictableResourceModel describes the resource data model.
type LbmetrictableResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Snmpoid     types.String `tfsdk:"snmpoid"`
	Metric      types.String `tfsdk:"metric"`
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
			"snmpoid": schema.StringAttribute{
				Required:    true,
				Description: "New SNMP OID of the metric.",
			},
			"metric": schema.StringAttribute{
				Required:    true,
				Description: "Name of the metric for which to change the SNMP OID.",
			},
			"metrictable": schema.StringAttribute{
				Required:    true,
				Description: "Name for the metric table. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my metrictable\" or 'my metrictable').",
			},
		},
	}
}

func lbmetrictableGetThePayloadFromtheConfig(ctx context.Context, data *LbmetrictableResourceModel) lb.Lbmetrictable {
	tflog.Debug(ctx, "In lbmetrictableGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lbmetrictable := lb.Lbmetrictable{}
	if !data.Snmpoid.IsNull() {
		lbmetrictable.Snmpoid = data.Snmpoid.ValueString()
	}
	if !data.Metric.IsNull() {
		lbmetrictable.Metric = data.Metric.ValueString()
	}
	if !data.Metrictable.IsNull() {
		lbmetrictable.Metrictable = data.Metrictable.ValueString()
	}

	return lbmetrictable
}

func lbmetrictableSetAttrFromGet(ctx context.Context, data *LbmetrictableResourceModel, getResponseData map[string]interface{}) *LbmetrictableResourceModel {
	tflog.Debug(ctx, "In lbmetrictableSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["Snmpoid"]; ok && val != nil {
		data.Snmpoid = types.StringValue(val.(string))
	} else {
		data.Snmpoid = types.StringNull()
	}
	if val, ok := getResponseData["metric"]; ok && val != nil {
		data.Metric = types.StringValue(val.(string))
	} else {
		data.Metric = types.StringNull()
	}
	if val, ok := getResponseData["metrictable"]; ok && val != nil {
		data.Metrictable = types.StringValue(val.(string))
	} else {
		data.Metrictable = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Metrictable.ValueString())

	return data
}
