package nsappflowcollector

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NsappflowcollectorResourceModel describes the resource data model.
type NsappflowcollectorResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Ipaddress types.String `tfsdk:"ipaddress"`
	Name      types.String `tfsdk:"name"`
	Port      types.Int64  `tfsdk:"port"`
}

func (r *NsappflowcollectorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsappflowcollector resource.",
			},
			"ipaddress": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The IPv4 address of the AppFlow collector.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the AppFlow collector.",
			},
			"port": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Default:     int64default.StaticInt64(4739),
				Description: "The UDP port on which the AppFlow collector is listening.",
			},
		},
	}
}

func nsappflowcollectorGetThePayloadFromtheConfig(ctx context.Context, data *NsappflowcollectorResourceModel) ns.Nsappflowcollector {
	tflog.Debug(ctx, "In nsappflowcollectorGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nsappflowcollector := ns.Nsappflowcollector{}
	if !data.Ipaddress.IsNull() {
		nsappflowcollector.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Name.IsNull() {
		nsappflowcollector.Name = data.Name.ValueString()
	}
	if !data.Port.IsNull() {
		nsappflowcollector.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}

	return nsappflowcollector
}

func nsappflowcollectorSetAttrFromGet(ctx context.Context, data *NsappflowcollectorResourceModel, getResponseData map[string]interface{}) *NsappflowcollectorResourceModel {
	tflog.Debug(ctx, "In nsappflowcollectorSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	} else {
		data.Ipaddress = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	} else {
		data.Port = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
