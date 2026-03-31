package lsnclient_nsacl_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/lsn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LsnclientNsaclBindingResourceModel describes the resource data model.
type LsnclientNsaclBindingResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Aclname    types.String `tfsdk:"aclname"`
	Clientname types.String `tfsdk:"clientname"`
	Td         types.Int64  `tfsdk:"td"`
}

func (r *LsnclientNsaclBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lsnclient_nsacl_binding resource.",
			},
			"aclname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name(s) of any configured extended ACL(s) whose action is ALLOW.\nThe condition specified in the extended ACL rule identifies the traffic from an LSN subscriber for which the Citrix ADC is to perform large scale NAT.",
			},
			"clientname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the LSN client entity. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN client is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"lsn client1\" or 'lsn client1').",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the traffic domain on which this subscriber or the subscriber network (as specified by the network parameter) belongs. \nIf you do not specify an ID, the subscriber or the subscriber network becomes part of the default traffic domain.",
			},
		},
	}
}

func lsnclient_nsacl_bindingGetThePayloadFromtheConfig(ctx context.Context, data *LsnclientNsaclBindingResourceModel) lsn.Lsnclientnsaclbinding {
	tflog.Debug(ctx, "In lsnclient_nsacl_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lsnclient_nsacl_binding := lsn.Lsnclientnsaclbinding{}
	if !data.Aclname.IsNull() {
		lsnclient_nsacl_binding.Aclname = data.Aclname.ValueString()
	}
	if !data.Clientname.IsNull() {
		lsnclient_nsacl_binding.Clientname = data.Clientname.ValueString()
	}
	if !data.Td.IsNull() {
		lsnclient_nsacl_binding.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}

	return lsnclient_nsacl_binding
}

func lsnclient_nsacl_bindingSetAttrFromGet(ctx context.Context, data *LsnclientNsaclBindingResourceModel, getResponseData map[string]interface{}) *LsnclientNsaclBindingResourceModel {
	tflog.Debug(ctx, "In lsnclient_nsacl_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["aclname"]; ok && val != nil {
		data.Aclname = types.StringValue(val.(string))
	} else {
		data.Aclname = types.StringNull()
	}
	if val, ok := getResponseData["clientname"]; ok && val != nil {
		data.Clientname = types.StringValue(val.(string))
	} else {
		data.Clientname = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else {
		data.Td = types.Int64Null()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("aclname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Aclname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("clientname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Clientname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
