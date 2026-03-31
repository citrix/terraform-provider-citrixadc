package lsnappsprofile_port_binding

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

// LsnappsprofilePortBindingResourceModel describes the resource data model.
type LsnappsprofilePortBindingResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Appsprofilename types.String `tfsdk:"appsprofilename"`
	Lsnport         types.String `tfsdk:"lsnport"`
}

func (r *LsnappsprofilePortBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lsnappsprofile_port_binding resource.",
			},
			"appsprofilename": schema.StringAttribute{
				Required:    true,
				Description: "Name for the LSN application profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN application profile is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"lsn application profile1\" or 'lsn application profile1').",
			},
			"lsnport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Port numbers or range of port numbers to match against the destination port of the incoming packet from a subscriber. When the destination port is matched, the LSN application profile is applied for the LSN session. Separate a range of ports with a hyphen. For example, 40-90.",
			},
		},
	}
}

func lsnappsprofile_port_bindingGetThePayloadFromtheConfig(ctx context.Context, data *LsnappsprofilePortBindingResourceModel) lsn.Lsnappsprofileportbinding {
	tflog.Debug(ctx, "In lsnappsprofile_port_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lsnappsprofile_port_binding := lsn.Lsnappsprofileportbinding{}
	if !data.Appsprofilename.IsNull() {
		lsnappsprofile_port_binding.Appsprofilename = data.Appsprofilename.ValueString()
	}
	if !data.Lsnport.IsNull() {
		lsnappsprofile_port_binding.Lsnport = data.Lsnport.ValueString()
	}

	return lsnappsprofile_port_binding
}

func lsnappsprofile_port_bindingSetAttrFromGet(ctx context.Context, data *LsnappsprofilePortBindingResourceModel, getResponseData map[string]interface{}) *LsnappsprofilePortBindingResourceModel {
	tflog.Debug(ctx, "In lsnappsprofile_port_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["appsprofilename"]; ok && val != nil {
		data.Appsprofilename = types.StringValue(val.(string))
	} else {
		data.Appsprofilename = types.StringNull()
	}
	if val, ok := getResponseData["lsnport"]; ok && val != nil {
		data.Lsnport = types.StringValue(val.(string))
	} else {
		data.Lsnport = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("appsprofilename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Appsprofilename.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("lsnport:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Lsnport.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
