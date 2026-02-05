package nslicenseproxyserver

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
			},
			"port": schema.Int64Attribute{
				Required:    true,
				Description: "License proxy server port.",
			},
			"serverip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the License proxy server.",
			},
			"servername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Fully qualified domain name of the License proxy server.",
			},
		},
	}
}

func nslicenseproxyserverGetThePayloadFromtheConfig(ctx context.Context, data *NslicenseproxyserverResourceModel) ns.Nslicenseproxyserver {
	tflog.Debug(ctx, "In nslicenseproxyserverGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nslicenseproxyserver := ns.Nslicenseproxyserver{}
	if !data.Port.IsNull() {
		nslicenseproxyserver.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Serverip.IsNull() {
		nslicenseproxyserver.Serverip = data.Serverip.ValueString()
	}
	if !data.Servername.IsNull() {
		nslicenseproxyserver.Servername = data.Servername.ValueString()
	}

	return nslicenseproxyserver
}

func nslicenseproxyserverSetAttrFromGet(ctx context.Context, data *NslicenseproxyserverResourceModel, getResponseData map[string]interface{}) *NslicenseproxyserverResourceModel {
	tflog.Debug(ctx, "In nslicenseproxyserverSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	} else {
		data.Port = types.Int64Null()
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

	// Set ID for the resource based on which identifiers are present
	var idParts []string
	if !data.Serverip.IsNull() {
		idParts = append(idParts, fmt.Sprintf("serverip:%s", data.Serverip.ValueString()))
	}
	if !data.Servername.IsNull() {
		idParts = append(idParts, fmt.Sprintf("servername:%s", data.Servername.ValueString()))
	}
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
