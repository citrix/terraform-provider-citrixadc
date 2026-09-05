package nsaigwprofile

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

func NsaigwprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the AIGW Profile.",
			},
			"endpointtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The type of AI GW endpoint type. Possible values = azureopenai",
			},
			"profiletype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The binding entity for the aigw profile. Possible values = frontend, backend",
			},
			"tokenquota": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Token capacity of the backend server.",
			},
			"quotarefreshfrequency": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Quota refresh rate, in minutes.",
			},
			"authtoken": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "Authentication token/API Key for the AI GW Endpoint.",
			},
		},
	}
}

type NsaigwprofileDataSourceModel struct {
	Id                    types.String `tfsdk:"id"`
	Name                  types.String `tfsdk:"name"`
	Endpointtype          types.String `tfsdk:"endpointtype"`
	Profiletype           types.String `tfsdk:"profiletype"`
	Tokenquota            types.Int64  `tfsdk:"tokenquota"`
	Quotarefreshfrequency types.Int64  `tfsdk:"quotarefreshfrequency"`
	Authtoken             types.String `tfsdk:"authtoken"`
}

func nsaigwprofileDataSourceSetAttrFromGet(ctx context.Context, data *NsaigwprofileDataSourceModel, getResponseData map[string]interface{}) *NsaigwprofileDataSourceModel {
	tflog.Debug(ctx, "In nsaigwprofileDataSourceSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["endpointtype"]; ok && val != nil {
		data.Endpointtype = types.StringValue(val.(string))
	} else {
		data.Endpointtype = types.StringNull()
	}
	if val, ok := getResponseData["profiletype"]; ok && val != nil {
		data.Profiletype = types.StringValue(val.(string))
	} else {
		data.Profiletype = types.StringNull()
	}
	if val, ok := getResponseData["tokenquota"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Tokenquota = types.Int64Value(intVal)
		}
	} else {
		data.Tokenquota = types.Int64Null()
	}
	if val, ok := getResponseData["quotarefreshfrequency"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Quotarefreshfrequency = types.Int64Value(intVal)
		}
	} else {
		data.Quotarefreshfrequency = types.Int64Null()
	}
	// authtoken is a secret and is not read back into state - retain from config

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
