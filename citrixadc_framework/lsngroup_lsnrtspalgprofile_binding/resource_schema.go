package lsngroup_lsnrtspalgprofile_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/lsn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LsngroupLsnrtspalgprofileBindingResourceModel describes the resource data model.
type LsngroupLsnrtspalgprofileBindingResourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Groupname          types.String `tfsdk:"groupname"`
	Rtspalgprofilename types.String `tfsdk:"rtspalgprofilename"`
}

func (r *LsngroupLsnrtspalgprofileBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lsngroup_lsnrtspalgprofile_binding resource.",
			},
			"groupname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the LSN group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN group is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"lsn group1\" or 'lsn group1').",
			},
			"rtspalgprofilename": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the LSN RTSP ALG Profile.",
			},
		},
	}
}

func lsngroup_lsnrtspalgprofile_bindingGetThePayloadFromthePlan(ctx context.Context, data *LsngroupLsnrtspalgprofileBindingResourceModel) lsn.Lsngrouplsnrtspalgprofilebinding {
	tflog.Debug(ctx, "In lsngroup_lsnrtspalgprofile_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	lsngroup_lsnrtspalgprofile_binding := lsn.Lsngrouplsnrtspalgprofilebinding{}
	if !data.Groupname.IsNull() && !data.Groupname.IsUnknown() {
		lsngroup_lsnrtspalgprofile_binding.Groupname = data.Groupname.ValueString()
	}
	if !data.Rtspalgprofilename.IsNull() && !data.Rtspalgprofilename.IsUnknown() {
		lsngroup_lsnrtspalgprofile_binding.Rtspalgprofilename = data.Rtspalgprofilename.ValueString()
	}

	return lsngroup_lsnrtspalgprofile_binding
}

func lsngroup_lsnrtspalgprofile_bindingSetAttrFromGet(ctx context.Context, data *LsngroupLsnrtspalgprofileBindingResourceModel, getResponseData map[string]interface{}) *LsngroupLsnrtspalgprofileBindingResourceModel {
	tflog.Debug(ctx, "In lsngroup_lsnrtspalgprofile_bindingSetAttrFromGet Function")

	// Convert API response to model. The ID is set once in Create and preserved across reads,
	// so it is not recomputed here.
	if val, ok := getResponseData["groupname"]; ok && val != nil {
		data.Groupname = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["rtspalgprofilename"]; ok && val != nil {
		data.Rtspalgprofilename = types.StringValue(val.(string))
	}

	return data
}

// lsngroup_lsnrtspalgprofile_bindingSetAttrFromGetForDatasource faithfully copies every field
// from the GET response and composes the ID, since the datasource has no Create to seed those values.
func lsngroup_lsnrtspalgprofile_bindingSetAttrFromGetForDatasource(ctx context.Context, data *LsngroupLsnrtspalgprofileBindingResourceModel, getResponseData map[string]interface{}) *LsngroupLsnrtspalgprofileBindingResourceModel {
	tflog.Debug(ctx, "In lsngroup_lsnrtspalgprofile_bindingSetAttrFromGetForDatasource Function")

	// Convert API response to model
	if val, ok := getResponseData["groupname"]; ok && val != nil {
		data.Groupname = types.StringValue(val.(string))
	} else {
		data.Groupname = types.StringNull()
	}
	if val, ok := getResponseData["rtspalgprofilename"]; ok && val != nil {
		data.Rtspalgprofilename = types.StringValue(val.(string))
	} else {
		data.Rtspalgprofilename = types.StringNull()
	}

	// Set ID for the datasource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("groupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Groupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("rtspalgprofilename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Rtspalgprofilename.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
