package mapdmr

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// MapdmrResourceModel describes the resource data model.
type MapdmrResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Bripv6prefix types.String `tfsdk:"bripv6prefix"`
	Name         types.String `tfsdk:"name"`
}

func (r *MapdmrResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the mapdmr resource.",
			},
			"bripv6prefix": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IPv6 prefix of Border Relay (Citrix ADC) device.MAP-T CE will send ipv6 packets to this ipv6 prefix.The DMR IPv6 prefix length SHOULD be 64 bits long by default and in any case MUST NOT exceed 96 bits",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the Default Mapping Rule. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the MAP Default Mapping Rule is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"add network MapDmr map1 -BRIpv6Prefix 2003::/96\").\n			Default Mapping Rule (DMR) is defined in terms of the IPv6 prefix advertised by one or more BRs, which provide external connectivity.  A typical MAP-T CE will install an IPv4 default route using this rule.  A BR will use this rule when translating all outside IPv4 source addresses to the IPv6 MAP domain.",
			},
		},
	}
}

func mapdmrGetThePayloadFromtheConfig(ctx context.Context, data *MapdmrResourceModel) network.Mapdmr {
	tflog.Debug(ctx, "In mapdmrGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	mapdmr := network.Mapdmr{}
	if !data.Bripv6prefix.IsNull() {
		mapdmr.Bripv6prefix = data.Bripv6prefix.ValueString()
	}
	if !data.Name.IsNull() {
		mapdmr.Name = data.Name.ValueString()
	}

	return mapdmr
}

func mapdmrSetAttrFromGet(ctx context.Context, data *MapdmrResourceModel, getResponseData map[string]interface{}) *MapdmrResourceModel {
	tflog.Debug(ctx, "In mapdmrSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["bripv6prefix"]; ok && val != nil {
		data.Bripv6prefix = types.StringValue(val.(string))
	} else {
		data.Bripv6prefix = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
