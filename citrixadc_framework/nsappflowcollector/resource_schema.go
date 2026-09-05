package nsappflowcollector

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
			// SDK v2: Required + ForceNew
			"ipaddress": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The IPv4 address of the AppFlow collector.",
			},
			// SDK v2: Required + ForceNew
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the AppFlow collector.",
			},
			// SDK v2: Optional + Computed + ForceNew (no Default in SDK v2 — value read from ADC).
			"port": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "The UDP port on which the AppFlow collector is listening.",
			},
		},
	}
}

func nsappflowcollectorGetThePayloadFromthePlan(ctx context.Context, data *NsappflowcollectorResourceModel) ns.Nsappflowcollector {
	tflog.Debug(ctx, "In nsappflowcollectorGetThePayloadFromthePlan Function")

	// Create API request body from the model
	nsappflowcollector := ns.Nsappflowcollector{}
	if !data.Ipaddress.IsNull() && !data.Ipaddress.IsUnknown() {
		nsappflowcollector.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		nsappflowcollector.Name = data.Name.ValueString()
	}
	// SDK v2 only set port when it was present in the raw config.
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		nsappflowcollector.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}

	return nsappflowcollector
}

func nsappflowcollectorSetAttrFromGet(ctx context.Context, data *NsappflowcollectorResourceModel, getResponseData map[string]interface{}) *NsappflowcollectorResourceModel {
	tflog.Debug(ctx, "In nsappflowcollectorSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	} else if data.Ipaddress.IsUnknown() {
		data.Ipaddress = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else if data.Name.IsUnknown() {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	} else if data.Port.IsUnknown() {
		// Only null when the value is unknown; never clobber a configured value.
		data.Port = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
