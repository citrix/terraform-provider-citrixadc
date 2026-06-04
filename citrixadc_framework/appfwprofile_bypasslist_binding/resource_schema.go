package appfwprofile_bypasslist_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AppfwprofileBypasslistBindingResourceModel describes the resource data model.
type AppfwprofileBypasslistBindingResourceModel struct {
	Id                    types.String `tfsdk:"id"`
	Alertonly             types.String `tfsdk:"alertonly"`
	AsBypassList          types.String `tfsdk:"as_bypass_list"`
	AsBypassListAction    types.String `tfsdk:"as_bypass_list_action"`
	AsBypassListLocation  types.String `tfsdk:"as_bypass_list_location"`
	AsBypassListValueType types.String `tfsdk:"as_bypass_list_value_type"`
	Comment               types.String `tfsdk:"comment"`
	Isautodeployed        types.String `tfsdk:"isautodeployed"`
	Name                  types.String `tfsdk:"name"`
	Resourceid            types.String `tfsdk:"resourceid"`
	State                 types.String `tfsdk:"state"`
}

func (r *AppfwprofileBypasslistBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwprofile_bypasslist_binding resource.",
			},
			"alertonly": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Send SNMP alert?",
			},
			"as_bypass_list": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Bypass List Value",
			},
			"as_bypass_list_action": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Bypass List Action",
			},
			"as_bypass_list_location": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Bypass List scan location",
			},
			"as_bypass_list_value_type": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Bypass List value type",
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

func appfwprofile_bypasslist_bindingGetThePayloadFromthePlan(ctx context.Context, data *AppfwprofileBypasslistBindingResourceModel) appfw.Appfwprofilebypasslistbinding {
	tflog.Debug(ctx, "In appfwprofile_bypasslist_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	// NOTE (Pattern 15 sanctioned exclusion): alertonly and resourceid are
	// read-only / server-assigned GET-response attributes for the bypassList
	// bind branch. The CLI does not accept them for
	// `bind appfw profile ... -bypasslist ...`, so they are kept Computed in the
	// schema (populated from GET) but NOT sent in the write payload. (ruletype is a
	// cross-branch attribute and is not part of this resource's model at all.)
	appfwprofile_bypasslist_binding := appfw.Appfwprofilebypasslistbinding{}
	if !data.AsBypassList.IsNull() && !data.AsBypassList.IsUnknown() {
		appfwprofile_bypasslist_binding.Asbypasslist = data.AsBypassList.ValueString()
	}
	if !data.AsBypassListAction.IsNull() && !data.AsBypassListAction.IsUnknown() {
		appfwprofile_bypasslist_binding.Asbypasslistaction = data.AsBypassListAction.ValueString()
	}
	if !data.AsBypassListLocation.IsNull() && !data.AsBypassListLocation.IsUnknown() {
		appfwprofile_bypasslist_binding.Asbypasslistlocation = data.AsBypassListLocation.ValueString()
	}
	if !data.AsBypassListValueType.IsNull() && !data.AsBypassListValueType.IsUnknown() {
		appfwprofile_bypasslist_binding.Asbypasslistvaluetype = data.AsBypassListValueType.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appfwprofile_bypasslist_binding.Comment = data.Comment.ValueString()
	}
	if !data.Isautodeployed.IsNull() && !data.Isautodeployed.IsUnknown() {
		appfwprofile_bypasslist_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwprofile_bypasslist_binding.Name = data.Name.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		appfwprofile_bypasslist_binding.State = data.State.ValueString()
	}

	return appfwprofile_bypasslist_binding
}

// appfwprofile_bypasslist_bindingSetAttrFromGet is the RESOURCE setter.
// It preserves user-supplied plan/state values for the write-able / identity
// attributes (name, as_bypass_list, as_bypass_list_value_type,
// as_bypass_list_location, as_bypass_list_action, comment, state,
// isautodeployed) so that a GET response (which may normalize or default these)
// does not produce an "inconsistent result after apply" diff. It only copies the
// server-managed read-only attributes (resourceid, alertonly) from the response.
// The ID is composed once in Create and never recomputed here (Pattern 6).
func appfwprofile_bypasslist_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileBypasslistBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileBypasslistBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_bypasslist_bindingSetAttrFromGet Function")

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

	// name, as_bypass_list, as_bypass_list_value_type, as_bypass_list_location,
	// as_bypass_list_action, comment and state are user-supplied (RequiresReplace)
	// and are intentionally NOT overwritten from the GET response - the plan/state
	// value is authoritative (Pattern 7).

	return data
}

// appfwprofile_bypasslist_bindingSetAttrFromGetForDatasource is the DATASOURCE
// setter. The datasource has no prior plan/state to preserve, so it faithfully
// copies every attribute from the GET response and composes the ID.
func appfwprofile_bypasslist_bindingSetAttrFromGetForDatasource(ctx context.Context, data *AppfwprofileBypasslistBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileBypasslistBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_bypasslist_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["alertonly"]; ok && val != nil {
		data.Alertonly = types.StringValue(val.(string))
	} else {
		data.Alertonly = types.StringNull()
	}
	if val, ok := getResponseData["as_bypass_list"]; ok && val != nil {
		data.AsBypassList = types.StringValue(val.(string))
	} else {
		data.AsBypassList = types.StringNull()
	}
	if val, ok := getResponseData["as_bypass_list_action"]; ok && val != nil {
		data.AsBypassListAction = types.StringValue(val.(string))
	} else {
		data.AsBypassListAction = types.StringNull()
	}
	if val, ok := getResponseData["as_bypass_list_location"]; ok && val != nil {
		data.AsBypassListLocation = types.StringValue(val.(string))
	} else {
		data.AsBypassListLocation = types.StringNull()
	}
	if val, ok := getResponseData["as_bypass_list_value_type"]; ok && val != nil {
		data.AsBypassListValueType = types.StringValue(val.(string))
	} else {
		data.AsBypassListValueType = types.StringNull()
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
	idParts = append(idParts, fmt.Sprintf("as_bypass_list:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsBypassList.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("as_bypass_list_location:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsBypassListLocation.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("as_bypass_list_value_type:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsBypassListValueType.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
