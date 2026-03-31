package lsnclient_nsacl6_binding

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

// LsnclientNsacl6BindingResourceModel describes the resource data model.
type LsnclientNsacl6BindingResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Acl6name   types.String `tfsdk:"acl6name"`
	Clientname types.String `tfsdk:"clientname"`
	Td         types.Int64  `tfsdk:"td"`
}

func (r *LsnclientNsacl6BindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lsnclient_nsacl6_binding resource.",
			},
			"acl6name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of any configured extended ACL6 whose action is ALLOW. The condition specified in the extended ACL6 rule is used as the condition for the LSN client.",
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

func lsnclient_nsacl6_bindingGetThePayloadFromtheConfig(ctx context.Context, data *LsnclientNsacl6BindingResourceModel) lsn.Lsnclientnsacl6binding {
	tflog.Debug(ctx, "In lsnclient_nsacl6_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lsnclient_nsacl6_binding := lsn.Lsnclientnsacl6binding{}
	if !data.Acl6name.IsNull() {
		lsnclient_nsacl6_binding.Acl6name = data.Acl6name.ValueString()
	}
	if !data.Clientname.IsNull() {
		lsnclient_nsacl6_binding.Clientname = data.Clientname.ValueString()
	}
	if !data.Td.IsNull() {
		lsnclient_nsacl6_binding.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}

	return lsnclient_nsacl6_binding
}

func lsnclient_nsacl6_bindingSetAttrFromGet(ctx context.Context, data *LsnclientNsacl6BindingResourceModel, getResponseData map[string]interface{}) *LsnclientNsacl6BindingResourceModel {
	tflog.Debug(ctx, "In lsnclient_nsacl6_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["acl6name"]; ok && val != nil {
		data.Acl6name = types.StringValue(val.(string))
	} else {
		data.Acl6name = types.StringNull()
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
	idParts = append(idParts, fmt.Sprintf("acl6name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Acl6name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("clientname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Clientname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
