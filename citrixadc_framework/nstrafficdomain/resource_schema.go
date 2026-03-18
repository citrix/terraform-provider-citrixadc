package nstrafficdomain

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NstrafficdomainResourceModel describes the resource data model.
type NstrafficdomainResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Aliasname types.String `tfsdk:"aliasname"`
	Td        types.Int64  `tfsdk:"td"`
	Vmac      types.String `tfsdk:"vmac"`
}

func (r *NstrafficdomainResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nstrafficdomain resource.",
			},
			"aliasname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of traffic domain  being added.",
			},
			"td": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Integer value that uniquely identifies a traffic domain.",
			},
			"vmac": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Associate the traffic domain with a VMAC address instead of with VLANs. The Citrix ADC then sends the VMAC address of the traffic domain in all responses to ARP queries for network entities in that domain. As a result, the ADC can segregate subsequent incoming traffic for this traffic domain on the basis of the destination MAC address, because the destination MAC address is the VMAC address of the traffic domain. After creating entities on a traffic domain, you can easily manage and monitor them by performing traffic domain level operations.",
			},
		},
	}
}

func nstrafficdomainGetThePayloadFromtheConfig(ctx context.Context, data *NstrafficdomainResourceModel) ns.Nstrafficdomain {
	tflog.Debug(ctx, "In nstrafficdomainGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nstrafficdomain := ns.Nstrafficdomain{}
	if !data.Aliasname.IsNull() {
		nstrafficdomain.Aliasname = data.Aliasname.ValueString()
	}
	if !data.Td.IsNull() {
		nstrafficdomain.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Vmac.IsNull() {
		nstrafficdomain.Vmac = data.Vmac.ValueString()
	}

	return nstrafficdomain
}

func nstrafficdomainSetAttrFromGet(ctx context.Context, data *NstrafficdomainResourceModel, getResponseData map[string]interface{}) *NstrafficdomainResourceModel {
	tflog.Debug(ctx, "In nstrafficdomainSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["aliasname"]; ok && val != nil {
		data.Aliasname = types.StringValue(val.(string))
	} else {
		data.Aliasname = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else {
		data.Td = types.Int64Null()
	}
	if val, ok := getResponseData["vmac"]; ok && val != nil {
		data.Vmac = types.StringValue(val.(string))
	} else {
		data.Vmac = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Td.ValueInt64()))

	return data
}
