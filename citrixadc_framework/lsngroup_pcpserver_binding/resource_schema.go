package lsngroup_pcpserver_binding

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

// LsngroupPcpserverBindingResourceModel describes the resource data model.
type LsngroupPcpserverBindingResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Groupname types.String `tfsdk:"groupname"`
	Pcpserver types.String `tfsdk:"pcpserver"`
}

func (r *LsngroupPcpserverBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lsngroup_pcpserver_binding resource.",
			},
			"groupname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the LSN group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN group is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"lsn group1\" or 'lsn group1').",
			},
			"pcpserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the PCP server to be associated with lsn group.",
			},
		},
	}
}

func lsngroup_pcpserver_bindingGetThePayloadFromtheConfig(ctx context.Context, data *LsngroupPcpserverBindingResourceModel) lsn.Lsngrouppcpserverbinding {
	tflog.Debug(ctx, "In lsngroup_pcpserver_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lsngroup_pcpserver_binding := lsn.Lsngrouppcpserverbinding{}
	if !data.Groupname.IsNull() {
		lsngroup_pcpserver_binding.Groupname = data.Groupname.ValueString()
	}
	if !data.Pcpserver.IsNull() {
		lsngroup_pcpserver_binding.Pcpserver = data.Pcpserver.ValueString()
	}

	return lsngroup_pcpserver_binding
}

func lsngroup_pcpserver_bindingSetAttrFromGet(ctx context.Context, data *LsngroupPcpserverBindingResourceModel, getResponseData map[string]interface{}) *LsngroupPcpserverBindingResourceModel {
	tflog.Debug(ctx, "In lsngroup_pcpserver_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["groupname"]; ok && val != nil {
		data.Groupname = types.StringValue(val.(string))
	} else {
		data.Groupname = types.StringNull()
	}
	if val, ok := getResponseData["pcpserver"]; ok && val != nil {
		data.Pcpserver = types.StringValue(val.(string))
	} else {
		data.Pcpserver = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("groupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Groupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("pcpserver:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Pcpserver.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
