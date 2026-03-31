package lsnappsprofile_lsnappsattributes_binding

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

// LsnappsprofileLsnappsattributesBindingResourceModel describes the resource data model.
type LsnappsprofileLsnappsattributesBindingResourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Appsattributesname types.String `tfsdk:"appsattributesname"`
	Appsprofilename    types.String `tfsdk:"appsprofilename"`
}

func (r *LsnappsprofileLsnappsattributesBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lsnappsprofile_lsnappsattributes_binding resource.",
			},
			"appsattributesname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the LSN application port ATTRIBUTES command to bind to the specified LSN Appsprofile. Properties of the Appsprofile will be applicable to this APPSATTRIBUTES",
			},
			"appsprofilename": schema.StringAttribute{
				Required:    true,
				Description: "Name for the LSN application profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN application profile is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"lsn application profile1\" or 'lsn application profile1').",
			},
		},
	}
}

func lsnappsprofile_lsnappsattributes_bindingGetThePayloadFromtheConfig(ctx context.Context, data *LsnappsprofileLsnappsattributesBindingResourceModel) lsn.Lsnappsprofilelsnappsattributesbinding {
	tflog.Debug(ctx, "In lsnappsprofile_lsnappsattributes_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lsnappsprofile_lsnappsattributes_binding := lsn.Lsnappsprofilelsnappsattributesbinding{}
	if !data.Appsattributesname.IsNull() {
		lsnappsprofile_lsnappsattributes_binding.Appsattributesname = data.Appsattributesname.ValueString()
	}
	if !data.Appsprofilename.IsNull() {
		lsnappsprofile_lsnappsattributes_binding.Appsprofilename = data.Appsprofilename.ValueString()
	}

	return lsnappsprofile_lsnappsattributes_binding
}

func lsnappsprofile_lsnappsattributes_bindingSetAttrFromGet(ctx context.Context, data *LsnappsprofileLsnappsattributesBindingResourceModel, getResponseData map[string]interface{}) *LsnappsprofileLsnappsattributesBindingResourceModel {
	tflog.Debug(ctx, "In lsnappsprofile_lsnappsattributes_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["appsattributesname"]; ok && val != nil {
		data.Appsattributesname = types.StringValue(val.(string))
	} else {
		data.Appsattributesname = types.StringNull()
	}
	if val, ok := getResponseData["appsprofilename"]; ok && val != nil {
		data.Appsprofilename = types.StringValue(val.(string))
	} else {
		data.Appsprofilename = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("appsattributesname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Appsattributesname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("appsprofilename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Appsprofilename.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
