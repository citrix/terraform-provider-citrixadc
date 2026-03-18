package nsservicefunction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NsservicefunctionResourceModel describes the resource data model.
type NsservicefunctionResourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Ingressvlan         types.Int64  `tfsdk:"ingressvlan"`
	Servicefunctionname types.String `tfsdk:"servicefunctionname"`
}

func (r *NsservicefunctionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsservicefunction resource.",
			},
			"ingressvlan": schema.Int64Attribute{
				Required:    true,
				Description: "VLAN ID on which the traffic from service function reaches Citrix ADC.",
			},
			"servicefunctionname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the service function to be created. Leading character must be a number or letter. Other characters allowed, after the first character, are @ _ - . (period) : (colon) # and space ( ).",
			},
		},
	}
}

func nsservicefunctionGetThePayloadFromtheConfig(ctx context.Context, data *NsservicefunctionResourceModel) ns.Nsservicefunction {
	tflog.Debug(ctx, "In nsservicefunctionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nsservicefunction := ns.Nsservicefunction{}
	if !data.Ingressvlan.IsNull() {
		nsservicefunction.Ingressvlan = utils.IntPtr(int(data.Ingressvlan.ValueInt64()))
	}
	if !data.Servicefunctionname.IsNull() {
		nsservicefunction.Servicefunctionname = data.Servicefunctionname.ValueString()
	}

	return nsservicefunction
}

func nsservicefunctionSetAttrFromGet(ctx context.Context, data *NsservicefunctionResourceModel, getResponseData map[string]interface{}) *NsservicefunctionResourceModel {
	tflog.Debug(ctx, "In nsservicefunctionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ingressvlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ingressvlan = types.Int64Value(intVal)
		}
	} else {
		data.Ingressvlan = types.Int64Null()
	}
	if val, ok := getResponseData["servicefunctionname"]; ok && val != nil {
		data.Servicefunctionname = types.StringValue(val.(string))
	} else {
		data.Servicefunctionname = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Servicefunctionname.ValueString())

	return data
}
