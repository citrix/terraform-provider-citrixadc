package nslicenseproxyserver

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NslicenseproxyserverResourceModel describes the resource data model.
type NslicenseproxyserverResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Port       types.Int64  `tfsdk:"port"`
	Serverip   types.String `tfsdk:"serverip"`
	Servername types.String `tfsdk:"servername"`
}

func (r *NslicenseproxyserverResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nslicenseproxyserver resource.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			// port is the only updateable attribute (SDK v2: Required, no ForceNew).
			"port": schema.Int64Attribute{
				Required:    true,
				Description: "License proxy server port.",
			},
			// serverip was Optional + ForceNew in SDK v2. Optional+Computed+ForceNew ->
			// RequiresReplaceIfConfigured (only replace when the configured value changes)
			// with UseStateForUnknown so a purely-computed value does not churn the plan.
			"serverip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the License proxy server.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
			},
			// servername was Optional + ForceNew in SDK v2 (backward-compat wins over the
			// metadata is_updateable=true). Same modifier treatment as serverip.
			"servername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Fully qualified domain name of the License proxy server.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
			},
		},
	}
}

func nslicenseproxyserverGetThePayloadFromtheConfig(ctx context.Context, data *NslicenseproxyserverResourceModel) ns.Nslicenseproxyserver {
	tflog.Debug(ctx, "In nslicenseproxyserverGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nslicenseproxyserver := ns.Nslicenseproxyserver{}
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		nslicenseproxyserver.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Serverip.IsNull() && !data.Serverip.IsUnknown() {
		nslicenseproxyserver.Serverip = data.Serverip.ValueString()
	}
	if !data.Servername.IsNull() && !data.Servername.IsUnknown() {
		nslicenseproxyserver.Servername = data.Servername.ValueString()
	}

	return nslicenseproxyserver
}

// nslicenseproxyserverSetAttrFromGet is the resource-side state setter. It maps the
// NITRO GET response onto the model WITHOUT overwriting the resource ID (the ID is
// assigned once in Create from the configured serverip/servername value and carried
// unchanged in state thereafter, matching the SDK v2 contract). Else-branches only
// null a value when it is Unknown so a known/configured value NITRO omits from GET is
// never clobbered (omit-on-default guard).
func nslicenseproxyserverSetAttrFromGet(ctx context.Context, data *NslicenseproxyserverResourceModel, getResponseData map[string]interface{}) *NslicenseproxyserverResourceModel {
	tflog.Debug(ctx, "In nslicenseproxyserverSetAttrFromGet Function")

	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	} else if data.Port.IsUnknown() {
		data.Port = types.Int64Null()
	}
	if val, ok := getResponseData["serverip"]; ok && val != nil && val.(string) != "" {
		data.Serverip = types.StringValue(val.(string))
	} else if data.Serverip.IsUnknown() {
		data.Serverip = types.StringNull()
	}
	if val, ok := getResponseData["servername"]; ok && val != nil && val.(string) != "" {
		data.Servername = types.StringValue(val.(string))
	} else if data.Servername.IsUnknown() {
		data.Servername = types.StringNull()
	}

	// NOTE: data.Id is intentionally NOT modified here. The ID is the plain
	// serverip-or-servername value assigned in Create (SDK v2 backward-compat).

	return data
}

// nslicenseproxyserverSetAttrFromGetForDatasource is the datasource-side setter. It
// copies all attributes from the GET response and assigns the ID from the
// serverip-or-servername value (serverip takes precedence, matching SDK v2).
func nslicenseproxyserverSetAttrFromGetForDatasource(ctx context.Context, data *NslicenseproxyserverResourceModel, getResponseData map[string]interface{}) *NslicenseproxyserverResourceModel {
	tflog.Debug(ctx, "In nslicenseproxyserverSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["serverip"]; ok && val != nil {
		data.Serverip = types.StringValue(val.(string))
	} else {
		data.Serverip = types.StringNull()
	}
	if val, ok := getResponseData["servername"]; ok && val != nil {
		data.Servername = types.StringValue(val.(string))
	} else {
		data.Servername = types.StringNull()
	}

	// ID = serverip-or-servername value (serverip precedence), matching SDK v2.
	var name string
	if !data.Serverip.IsNull() && data.Serverip.ValueString() != "" {
		name = data.Serverip.ValueString()
	} else if !data.Servername.IsNull() && data.Servername.ValueString() != "" {
		name = data.Servername.ValueString()
	}
	data.Id = types.StringValue(name)

	return data
}
