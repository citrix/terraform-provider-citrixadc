package aaagroup_vpnsecureprivateaccessprofile_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AaagroupVpnsecureprivateaccessprofileBindingResourceModel describes the resource data model.
type AaagroupVpnsecureprivateaccessprofileBindingResourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Gotopriorityexpression     types.String `tfsdk:"gotopriorityexpression"`
	Groupname                  types.String `tfsdk:"groupname"`
	Secureprivateaccessprofile types.String `tfsdk:"secureprivateaccessprofile"`
}

func (r *AaagroupVpnsecureprivateaccessprofileBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the aaagroup_vpnsecureprivateaccessprofile_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					// GH #1436
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE. Specify one of the following values: NEXT, END, USE_INVOCATION_RESULT, or an expression that evaluates to a number.",
			},
			"groupname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the group that you are binding.",
			},
			"secureprivateaccessprofile": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the Secure Private Access Profile bound to the group.",
			},
		},
	}
}

func aaagroup_vpnsecureprivateaccessprofile_bindingGetThePayloadFromthePlan(ctx context.Context, data *AaagroupVpnsecureprivateaccessprofileBindingResourceModel) aaa.Aaagroupvpnsecureprivateaccessprofilebinding {
	tflog.Debug(ctx, "In aaagroup_vpnsecureprivateaccessprofile_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	aaagroup_vpnsecureprivateaccessprofile_binding := aaa.Aaagroupvpnsecureprivateaccessprofilebinding{}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		aaagroup_vpnsecureprivateaccessprofile_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Groupname.IsNull() && !data.Groupname.IsUnknown() {
		aaagroup_vpnsecureprivateaccessprofile_binding.Groupname = data.Groupname.ValueString()
	}
	if !data.Secureprivateaccessprofile.IsNull() && !data.Secureprivateaccessprofile.IsUnknown() {
		aaagroup_vpnsecureprivateaccessprofile_binding.Secureprivateaccessprofile = data.Secureprivateaccessprofile.ValueString()
	}

	return aaagroup_vpnsecureprivateaccessprofile_binding
}

func aaagroup_vpnsecureprivateaccessprofile_bindingSetAttrFromGet(ctx context.Context, data *AaagroupVpnsecureprivateaccessprofileBindingResourceModel, getResponseData map[string]interface{}) *AaagroupVpnsecureprivateaccessprofileBindingResourceModel {
	tflog.Debug(ctx, "In aaagroup_vpnsecureprivateaccessprofile_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["groupname"]; ok && val != nil {
		data.Groupname = types.StringValue(val.(string))
	} else {
		data.Groupname = types.StringNull()
	}
	if val, ok := getResponseData["secureprivateaccessprofile"]; ok && val != nil {
		data.Secureprivateaccessprofile = types.StringValue(val.(string))
	} else {
		data.Secureprivateaccessprofile = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("groupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Groupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("secureprivateaccessprofile:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Secureprivateaccessprofile.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
