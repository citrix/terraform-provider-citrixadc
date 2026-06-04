package appfwprofile_denylist_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AppfwprofileDenylistBindingResourceModel describes the resource data model.
type AppfwprofileDenylistBindingResourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Alertonly           types.String `tfsdk:"alertonly"`
	AsDenyList          types.String `tfsdk:"as_deny_list"`
	AsDenyListAction    types.List   `tfsdk:"as_deny_list_action"`
	AsDenyListLocation  types.String `tfsdk:"as_deny_list_location"`
	AsDenyListValueType types.String `tfsdk:"as_deny_list_value_type"`
	Comment             types.String `tfsdk:"comment"`
	Isautodeployed      types.String `tfsdk:"isautodeployed"`
	Name                types.String `tfsdk:"name"`
	Resourceid          types.String `tfsdk:"resourceid"`
	State               types.String `tfsdk:"state"`
}

func (r *AppfwprofileDenylistBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwprofile_denylist_binding resource.",
			},
			"alertonly": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Send SNMP alert?",
			},
			"as_deny_list": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Deny List Value",
			},
			"as_deny_list_action": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				Description: "Deny List Action. Default value = REDIRECT",
			},
			"as_deny_list_location": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Deny List scan location",
			},
			"as_deny_list_value_type": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Deny List value type",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments about the purpose of profile, or other useful information about the profile.",
			},
			"isautodeployed": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Is the rule auto deployed by dynamic profile ?",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the profile to which to bind an exemption or rule.",
			},
			"resourceid": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "A \"id\" that identifies the rule.",
			},
			"state": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Enabled.",
			},
		},
	}
}

func appfwprofile_denylist_bindingGetThePayloadFromthePlan(ctx context.Context, data *AppfwprofileDenylistBindingResourceModel) appfw.Appfwprofiledenylistbinding {
	tflog.Debug(ctx, "In appfwprofile_denylist_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	// NOTE (Pattern 15 sanctioned exclusion): alertonly and resourceid are
	// read-only / server-assigned GET-response attributes for the denyList
	// bind branch. The CLI does not accept them for
	// `bind appfw profile ... -denylist ...`, so they are kept Computed in the
	// schema (populated from GET) but NOT sent in the write payload. (ruletype is a
	// cross-branch attribute and is not part of this resource's model at all.)
	appfwprofile_denylist_binding := appfw.Appfwprofiledenylistbinding{}
	if !data.AsDenyList.IsNull() && !data.AsDenyList.IsUnknown() {
		appfwprofile_denylist_binding.Asdenylist = data.AsDenyList.ValueString()
	}
	if !data.AsDenyListAction.IsNull() && !data.AsDenyListAction.IsUnknown() {
		var as_deny_list_actionList []string
		data.AsDenyListAction.ElementsAs(ctx, &as_deny_list_actionList, false)
		appfwprofile_denylist_binding.Asdenylistaction = as_deny_list_actionList
	}
	if !data.AsDenyListLocation.IsNull() && !data.AsDenyListLocation.IsUnknown() {
		appfwprofile_denylist_binding.Asdenylistlocation = data.AsDenyListLocation.ValueString()
	}
	if !data.AsDenyListValueType.IsNull() && !data.AsDenyListValueType.IsUnknown() {
		appfwprofile_denylist_binding.Asdenylistvaluetype = data.AsDenyListValueType.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appfwprofile_denylist_binding.Comment = data.Comment.ValueString()
	}
	if !data.Isautodeployed.IsNull() && !data.Isautodeployed.IsUnknown() {
		appfwprofile_denylist_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwprofile_denylist_binding.Name = data.Name.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		appfwprofile_denylist_binding.State = data.State.ValueString()
	}

	return appfwprofile_denylist_binding
}

// appfwprofile_denylist_bindingSetAttrFromGet is the RESOURCE setter.
// It preserves user-supplied plan/state values for the write-able / identity
// attributes (name, as_deny_list, as_deny_list_value_type,
// as_deny_list_location, as_deny_list_action, comment, state, isautodeployed) so
// that a GET response (which may normalize or default these) does not produce an
// "inconsistent result after apply" diff. It only copies the server-managed
// read-only attributes (resourceid, alertonly) from the response. The ID is
// composed once in Create and never recomputed here (Pattern 6).
func appfwprofile_denylist_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileDenylistBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileDenylistBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_denylist_bindingSetAttrFromGet Function")

	// Server-managed read-only Computed attributes - copied from the GET response
	// so the Computed values become known after apply (Pattern 7 ECHOED branch).
	if val, ok := getResponseData["alertonly"]; ok && val != nil {
		data.Alertonly = types.StringValue(val.(string))
	} else {
		data.Alertonly = types.StringNull()
	}
	if val, ok := getResponseData["isautodeployed"]; ok && val != nil {
		data.Isautodeployed = types.StringValue(val.(string))
	} else {
		data.Isautodeployed = types.StringNull()
	}
	if val, ok := getResponseData["resourceid"]; ok && val != nil {
		data.Resourceid = types.StringValue(val.(string))
	} else {
		data.Resourceid = types.StringNull()
	}

	// name, as_deny_list, as_deny_list_value_type, as_deny_list_location,
	// as_deny_list_action (string[]), comment and state are user-supplied
	// (RequiresReplace) and are intentionally NOT overwritten from the GET response
	// - the plan/state value is authoritative (Pattern 7).

	return data
}

// appfwprofile_denylist_bindingSetAttrFromGetForDatasource is the DATASOURCE
// setter. The datasource has no prior plan/state to preserve, so it faithfully
// copies every attribute from the GET response and composes the ID.
func appfwprofile_denylist_bindingSetAttrFromGetForDatasource(ctx context.Context, data *AppfwprofileDenylistBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileDenylistBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_denylist_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["alertonly"]; ok && val != nil {
		data.Alertonly = types.StringValue(val.(string))
	} else {
		data.Alertonly = types.StringNull()
	}
	if val, ok := getResponseData["as_deny_list"]; ok && val != nil {
		data.AsDenyList = types.StringValue(val.(string))
	} else {
		data.AsDenyList = types.StringNull()
	}
	if val, ok := getResponseData["as_deny_list_action"]; ok && val != nil {
		if sliceVal, ok := val.([]interface{}); ok {
			stringList := utils.ToStringList(sliceVal)
			listValue, _ := types.ListValueFrom(ctx, types.StringType, stringList)
			data.AsDenyListAction = listValue
		} else {
			data.AsDenyListAction = types.ListNull(types.StringType)
		}
	} else {
		data.AsDenyListAction = types.ListNull(types.StringType)
	}
	if val, ok := getResponseData["as_deny_list_location"]; ok && val != nil {
		data.AsDenyListLocation = types.StringValue(val.(string))
	} else {
		data.AsDenyListLocation = types.StringNull()
	}
	if val, ok := getResponseData["as_deny_list_value_type"]; ok && val != nil {
		data.AsDenyListValueType = types.StringValue(val.(string))
	} else {
		data.AsDenyListValueType = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["isautodeployed"]; ok && val != nil {
		data.Isautodeployed = types.StringValue(val.(string))
	} else {
		data.Isautodeployed = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["resourceid"]; ok && val != nil {
		data.Resourceid = types.StringValue(val.(string))
	} else {
		data.Resourceid = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}

	// Compose the ID (datasource has no Create).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("as_deny_list:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsDenyList.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("as_deny_list_location:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsDenyListLocation.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("as_deny_list_value_type:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsDenyListValueType.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
