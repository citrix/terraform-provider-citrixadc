package lsngroup_lsnlogprofile_binding

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

// LsngroupLsnlogprofileBindingResourceModel describes the resource data model.
type LsngroupLsnlogprofileBindingResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Groupname      types.String `tfsdk:"groupname"`
	Logprofilename types.String `tfsdk:"logprofilename"`
}

func (r *LsngroupLsnlogprofileBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lsngroup_lsnlogprofile_binding resource.",
			},
			"groupname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the LSN group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN group is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"lsn group1\" or 'lsn group1').",
			},
			"logprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the LSN logging Profile.",
			},
		},
	}
}

func lsngroup_lsnlogprofile_bindingGetThePayloadFromtheConfig(ctx context.Context, data *LsngroupLsnlogprofileBindingResourceModel) lsn.Lsngrouplsnlogprofilebinding {
	tflog.Debug(ctx, "In lsngroup_lsnlogprofile_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lsngroup_lsnlogprofile_binding := lsn.Lsngrouplsnlogprofilebinding{}
	if !data.Groupname.IsNull() {
		lsngroup_lsnlogprofile_binding.Groupname = data.Groupname.ValueString()
	}
	if !data.Logprofilename.IsNull() {
		lsngroup_lsnlogprofile_binding.Logprofilename = data.Logprofilename.ValueString()
	}

	return lsngroup_lsnlogprofile_binding
}

func lsngroup_lsnlogprofile_bindingSetAttrFromGet(ctx context.Context, data *LsngroupLsnlogprofileBindingResourceModel, getResponseData map[string]interface{}) *LsngroupLsnlogprofileBindingResourceModel {
	tflog.Debug(ctx, "In lsngroup_lsnlogprofile_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["groupname"]; ok && val != nil {
		data.Groupname = types.StringValue(val.(string))
	} else {
		data.Groupname = types.StringNull()
	}
	if val, ok := getResponseData["logprofilename"]; ok && val != nil {
		data.Logprofilename = types.StringValue(val.(string))
	} else {
		data.Logprofilename = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("groupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Groupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("logprofilename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Logprofilename.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
