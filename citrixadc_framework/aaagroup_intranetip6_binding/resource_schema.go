package aaagroup_intranetip6_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AaagroupIntranetip6BindingResourceModel describes the resource data model.
type AaagroupIntranetip6BindingResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Groupname   types.String `tfsdk:"groupname"`
	Intranetip6 types.String `tfsdk:"intranetip6"`
	Numaddr     types.Int64  `tfsdk:"numaddr"`
}

func (r *AaagroupIntranetip6BindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the aaagroup_intranetip6_binding resource.",
			},
			"groupname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the group that you are binding.",
			},
			"intranetip6": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The Intranet IP6(s) bound to the group",
			},
			"numaddr": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Numbers of ipv6 address bound starting with intranetip6",
			},
		},
	}
}

func aaagroup_intranetip6_bindingGetThePayloadFromthePlan(ctx context.Context, data *AaagroupIntranetip6BindingResourceModel) aaa.Aaagroupintranetip6binding {
	tflog.Debug(ctx, "In aaagroup_intranetip6_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	aaagroup_intranetip6_binding := aaa.Aaagroupintranetip6binding{}
	if !data.Groupname.IsNull() && !data.Groupname.IsUnknown() {
		aaagroup_intranetip6_binding.Groupname = data.Groupname.ValueString()
	}
	if !data.Intranetip6.IsNull() && !data.Intranetip6.IsUnknown() {
		aaagroup_intranetip6_binding.Intranetip6 = data.Intranetip6.ValueString()
	}
	if !data.Numaddr.IsNull() && !data.Numaddr.IsUnknown() {
		aaagroup_intranetip6_binding.Numaddr = utils.IntPtr(int(data.Numaddr.ValueInt64()))
	}

	return aaagroup_intranetip6_binding
}

// aaagroup_intranetip6_bindingSetAttrFromGet is the resource-side setter.
// It refreshes the identity attributes from the GET response and does NOT
// (re)compute the ID - the resource sets the ID exactly once in Create.
func aaagroup_intranetip6_bindingSetAttrFromGet(ctx context.Context, data *AaagroupIntranetip6BindingResourceModel, getResponseData map[string]interface{}) *AaagroupIntranetip6BindingResourceModel {
	tflog.Debug(ctx, "In aaagroup_intranetip6_bindingSetAttrFromGet Function")

	if val, ok := getResponseData["groupname"]; ok && val != nil {
		data.Groupname = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["intranetip6"]; ok && val != nil {
		data.Intranetip6 = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["numaddr"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Numaddr = types.Int64Value(intVal)
		}
	}

	return data
}

// aaagroup_intranetip6_bindingSetAttrFromGetForDatasource faithfully copies
// every field from the GET response and composes the ID, because the datasource
// never calls Create and has no prior state to preserve.
func aaagroup_intranetip6_bindingSetAttrFromGetForDatasource(ctx context.Context, data *AaagroupIntranetip6BindingResourceModel, getResponseData map[string]interface{}) *AaagroupIntranetip6BindingResourceModel {
	tflog.Debug(ctx, "In aaagroup_intranetip6_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["groupname"]; ok && val != nil {
		data.Groupname = types.StringValue(val.(string))
	} else {
		data.Groupname = types.StringNull()
	}
	if val, ok := getResponseData["intranetip6"]; ok && val != nil {
		data.Intranetip6 = types.StringValue(val.(string))
	} else {
		data.Intranetip6 = types.StringNull()
	}
	if val, ok := getResponseData["numaddr"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Numaddr = types.Int64Value(intVal)
		}
	} else {
		data.Numaddr = types.Int64Null()
	}

	// Compose the composite ID: groupname,intranetip6,numaddr.
	// IPv6 colons are percent-encoded by UrlEncode so they never collide with
	// the key:value / comma delimiters used by ParseIdString.
	data.Id = types.StringValue(aaagroup_intranetip6_bindingComposeId(
		data.Groupname.ValueString(),
		data.Intranetip6.ValueString(),
		data.Numaddr.ValueInt64(),
	))

	return data
}

// aaagroup_intranetip6_bindingComposeId builds the colon-safe composite ID.
// Each value is UrlEncoded so the IPv6 address's literal ':' characters become
// '%3A' and cannot be mistaken for the key:value separator that ParseIdString
// keys on (it finds the FIRST ':' in each comma-separated segment).
func aaagroup_intranetip6_bindingComposeId(groupname string, intranetip6 string, numaddr int64) string {
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("groupname:%s", utils.UrlEncode(groupname)))
	idParts = append(idParts, fmt.Sprintf("intranetip6:%s", utils.UrlEncode(intranetip6)))
	idParts = append(idParts, fmt.Sprintf("numaddr:%s", utils.UrlEncode(fmt.Sprintf("%d", numaddr))))
	return strings.Join(idParts, ",")
}
