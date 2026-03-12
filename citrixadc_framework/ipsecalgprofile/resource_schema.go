package ipsecalgprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ipsecalg"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// IpsecalgprofileResourceModel describes the resource data model.
type IpsecalgprofileResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Connfailover      types.String `tfsdk:"connfailover"`
	Espgatetimeout    types.Int64  `tfsdk:"espgatetimeout"`
	Espsessiontimeout types.Int64  `tfsdk:"espsessiontimeout"`
	Ikesessiontimeout types.Int64  `tfsdk:"ikesessiontimeout"`
	Name              types.String `tfsdk:"name"`
}

func (r *IpsecalgprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the ipsecalgprofile resource.",
			},
			"connfailover": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Mode in which the connection failover feature must operate for the IPSec Alg. After a failover, established UDP connections and ESP packet flows are kept active and resumed on the secondary appliance. Recomended setting is ENABLED.",
			},
			"espgatetimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(30),
				Description: "Timeout ESP in seconds as no ESP packets are seen after IKE negotiation",
			},
			"espsessiontimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(60),
				Description: "ESP session timeout in minutes.",
			},
			"ikesessiontimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(60),
				Description: "IKE session timeout in minutes",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the ipsec alg profile",
			},
		},
	}
}

func ipsecalgprofileGetThePayloadFromtheConfig(ctx context.Context, data *IpsecalgprofileResourceModel) ipsecalg.Ipsecalgprofile {
	tflog.Debug(ctx, "In ipsecalgprofileGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	ipsecalgprofile := ipsecalg.Ipsecalgprofile{}
	if !data.Connfailover.IsNull() {
		ipsecalgprofile.Connfailover = data.Connfailover.ValueString()
	}
	if !data.Espgatetimeout.IsNull() {
		ipsecalgprofile.Espgatetimeout = utils.IntPtr(int(data.Espgatetimeout.ValueInt64()))
	}
	if !data.Espsessiontimeout.IsNull() {
		ipsecalgprofile.Espsessiontimeout = utils.IntPtr(int(data.Espsessiontimeout.ValueInt64()))
	}
	if !data.Ikesessiontimeout.IsNull() {
		ipsecalgprofile.Ikesessiontimeout = utils.IntPtr(int(data.Ikesessiontimeout.ValueInt64()))
	}
	if !data.Name.IsNull() {
		ipsecalgprofile.Name = data.Name.ValueString()
	}

	return ipsecalgprofile
}

func ipsecalgprofileSetAttrFromGet(ctx context.Context, data *IpsecalgprofileResourceModel, getResponseData map[string]interface{}) *IpsecalgprofileResourceModel {
	tflog.Debug(ctx, "In ipsecalgprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["connfailover"]; ok && val != nil {
		data.Connfailover = types.StringValue(val.(string))
	} else {
		data.Connfailover = types.StringNull()
	}
	if val, ok := getResponseData["espgatetimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Espgatetimeout = types.Int64Value(intVal)
		}
	} else {
		data.Espgatetimeout = types.Int64Null()
	}
	if val, ok := getResponseData["espsessiontimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Espsessiontimeout = types.Int64Value(intVal)
		}
	} else {
		data.Espsessiontimeout = types.Int64Null()
	}
	if val, ok := getResponseData["ikesessiontimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ikesessiontimeout = types.Int64Value(intVal)
		}
	} else {
		data.Ikesessiontimeout = types.Int64Null()
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
