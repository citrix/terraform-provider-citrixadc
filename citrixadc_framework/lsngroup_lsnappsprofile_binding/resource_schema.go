package lsngroup_lsnappsprofile_binding

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

// LsngroupLsnappsprofileBindingResourceModel describes the resource data model.
type LsngroupLsnappsprofileBindingResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Appsprofilename types.String `tfsdk:"appsprofilename"`
	Groupname       types.String `tfsdk:"groupname"`
}

func (r *LsngroupLsnappsprofileBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lsngroup_lsnappsprofile_binding resource.",
			},
			"appsprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the LSN application profile to bind to the specified LSN group. For each set of destination ports, bind a profile for each protocol for which you want to specify settings.\n\nBy default, one LSN application profile with default settings for TCP, UDP, and ICMP protocols for all destination ports is bound to an LSN group during its creation.  This profile is called a default application profile.\n\nWhen you bind an LSN application profile, with a specified set of destination ports, to an LSN group, the bound profile overrides the default LSN application profile for that protocol at that set of destination ports.",
			},
			"groupname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the LSN group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN group is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"lsn group1\" or 'lsn group1').",
			},
		},
	}
}

func lsngroup_lsnappsprofile_bindingGetThePayloadFromtheConfig(ctx context.Context, data *LsngroupLsnappsprofileBindingResourceModel) lsn.Lsngrouplsnappsprofilebinding {
	tflog.Debug(ctx, "In lsngroup_lsnappsprofile_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lsngroup_lsnappsprofile_binding := lsn.Lsngrouplsnappsprofilebinding{}
	if !data.Appsprofilename.IsNull() {
		lsngroup_lsnappsprofile_binding.Appsprofilename = data.Appsprofilename.ValueString()
	}
	if !data.Groupname.IsNull() {
		lsngroup_lsnappsprofile_binding.Groupname = data.Groupname.ValueString()
	}

	return lsngroup_lsnappsprofile_binding
}

func lsngroup_lsnappsprofile_bindingSetAttrFromGet(ctx context.Context, data *LsngroupLsnappsprofileBindingResourceModel, getResponseData map[string]interface{}) *LsngroupLsnappsprofileBindingResourceModel {
	tflog.Debug(ctx, "In lsngroup_lsnappsprofile_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["appsprofilename"]; ok && val != nil {
		data.Appsprofilename = types.StringValue(val.(string))
	} else {
		data.Appsprofilename = types.StringNull()
	}
	if val, ok := getResponseData["groupname"]; ok && val != nil {
		data.Groupname = types.StringValue(val.(string))
	} else {
		data.Groupname = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("appsprofilename:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Appsprofilename.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("groupname:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Groupname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
