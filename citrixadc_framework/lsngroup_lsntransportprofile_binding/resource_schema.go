package lsngroup_lsntransportprofile_binding

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

// LsngroupLsntransportprofileBindingResourceModel describes the resource data model.
type LsngroupLsntransportprofileBindingResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Groupname            types.String `tfsdk:"groupname"`
	Transportprofilename types.String `tfsdk:"transportprofilename"`
}

func (r *LsngroupLsntransportprofileBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lsngroup_lsntransportprofile_binding resource.",
			},
			"groupname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the LSN group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN group is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"lsn group1\" or 'lsn group1').",
			},
			"transportprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the LSN transport profile to bind to the specified LSN group. Bind a profile for each protocol for which you want to specify settings.\n\nBy default, one LSN transport profile with default settings for TCP, UDP, and ICMP protocols is bound to an LSN group during its creation. This profile is called a default transport.\n\nAn LSN transport profile that you bind to an LSN group overrides the default LSN transport profile for that protocol.",
			},
		},
	}
}

func lsngroup_lsntransportprofile_bindingGetThePayloadFromtheConfig(ctx context.Context, data *LsngroupLsntransportprofileBindingResourceModel) lsn.Lsngrouplsntransportprofilebinding {
	tflog.Debug(ctx, "In lsngroup_lsntransportprofile_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lsngroup_lsntransportprofile_binding := lsn.Lsngrouplsntransportprofilebinding{}
	if !data.Groupname.IsNull() {
		lsngroup_lsntransportprofile_binding.Groupname = data.Groupname.ValueString()
	}
	if !data.Transportprofilename.IsNull() {
		lsngroup_lsntransportprofile_binding.Transportprofilename = data.Transportprofilename.ValueString()
	}

	return lsngroup_lsntransportprofile_binding
}

func lsngroup_lsntransportprofile_bindingSetAttrFromGet(ctx context.Context, data *LsngroupLsntransportprofileBindingResourceModel, getResponseData map[string]interface{}) *LsngroupLsntransportprofileBindingResourceModel {
	tflog.Debug(ctx, "In lsngroup_lsntransportprofile_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["groupname"]; ok && val != nil {
		data.Groupname = types.StringValue(val.(string))
	} else {
		data.Groupname = types.StringNull()
	}
	if val, ok := getResponseData["transportprofilename"]; ok && val != nil {
		data.Transportprofilename = types.StringValue(val.(string))
	} else {
		data.Transportprofilename = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("groupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Groupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("transportprofilename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Transportprofilename.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
