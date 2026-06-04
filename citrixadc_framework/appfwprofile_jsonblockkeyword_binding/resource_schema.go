package appfwprofile_jsonblockkeyword_binding

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

// AppfwprofileJsonblockkeywordBindingResourceModel describes the resource data model.
type AppfwprofileJsonblockkeywordBindingResourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Alertonly                  types.String `tfsdk:"alertonly"`
	Comment                    types.String `tfsdk:"comment"`
	Isautodeployed             types.String `tfsdk:"isautodeployed"`
	IskeyregexJsonBlockkeyword types.String `tfsdk:"iskeyregex_json_blockkeyword"`
	Jsonblockkeyword           types.String `tfsdk:"jsonblockkeyword"`
	Jsonblockkeywordtype       types.String `tfsdk:"jsonblockkeywordtype"`
	Jsonblockkeywordurl        types.String `tfsdk:"jsonblockkeywordurl"`
	KeynameJsonBlockkeyword    types.String `tfsdk:"keyname_json_blockkeyword"`
	Name                       types.String `tfsdk:"name"`
	Resourceid                 types.String `tfsdk:"resourceid"`
	State                      types.String `tfsdk:"state"`
}

func (r *AppfwprofileJsonblockkeywordBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwprofile_jsonblockkeyword_binding resource.",
			},
			"alertonly": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Send SNMP alert?",
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
			"iskeyregex_json_blockkeyword": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Is JSON blockkeyword key a regular expression?",
			},
			"jsonblockkeyword": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Field name of json block keyword binding",
			},
			"jsonblockkeywordtype": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "JSON block keyword type",
			},
			"jsonblockkeywordurl": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The json blockkeyword URL.",
			},
			"keyname_json_blockkeyword": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "JSON block keyword keyname",
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

func appfwprofile_jsonblockkeyword_bindingGetThePayloadFromthePlan(ctx context.Context, data *AppfwprofileJsonblockkeywordBindingResourceModel) appfw.Appfwprofilejsonblockkeywordbinding {
	tflog.Debug(ctx, "In appfwprofile_jsonblockkeyword_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model.
	// NOTE (Pattern 15 sanctioned exclusion): alertonly and resourceid are
	// read-only / server-assigned GET-response attributes for the JSONBlockKeyword
	// bind branch. The CLI does not accept them for
	// `bind appfw profile ... -JSONBlockKeyword ...`, so they are kept Computed in
	// the schema (populated from GET) but NOT sent in the write payload. (ruletype is
	// a cross-branch attribute and is not part of this resource's model at all, so
	// there is nothing to exclude for it here.)
	appfwprofile_jsonblockkeyword_binding := appfw.Appfwprofilejsonblockkeywordbinding{}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appfwprofile_jsonblockkeyword_binding.Comment = data.Comment.ValueString()
	}
	if !data.Isautodeployed.IsNull() && !data.Isautodeployed.IsUnknown() {
		appfwprofile_jsonblockkeyword_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.IskeyregexJsonBlockkeyword.IsNull() && !data.IskeyregexJsonBlockkeyword.IsUnknown() {
		appfwprofile_jsonblockkeyword_binding.Iskeyregexjsonblockkeyword = data.IskeyregexJsonBlockkeyword.ValueString()
	}
	if !data.Jsonblockkeyword.IsNull() && !data.Jsonblockkeyword.IsUnknown() {
		appfwprofile_jsonblockkeyword_binding.Jsonblockkeyword = data.Jsonblockkeyword.ValueString()
	}
	if !data.Jsonblockkeywordtype.IsNull() && !data.Jsonblockkeywordtype.IsUnknown() {
		appfwprofile_jsonblockkeyword_binding.Jsonblockkeywordtype = data.Jsonblockkeywordtype.ValueString()
	}
	if !data.Jsonblockkeywordurl.IsNull() && !data.Jsonblockkeywordurl.IsUnknown() {
		appfwprofile_jsonblockkeyword_binding.Jsonblockkeywordurl = data.Jsonblockkeywordurl.ValueString()
	}
	if !data.KeynameJsonBlockkeyword.IsNull() && !data.KeynameJsonBlockkeyword.IsUnknown() {
		appfwprofile_jsonblockkeyword_binding.Keynamejsonblockkeyword = data.KeynameJsonBlockkeyword.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwprofile_jsonblockkeyword_binding.Name = data.Name.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		appfwprofile_jsonblockkeyword_binding.State = data.State.ValueString()
	}

	return appfwprofile_jsonblockkeyword_binding
}

// appfwprofile_jsonblockkeyword_bindingSetAttrFromGet is the RESOURCE setter.
// It preserves user-supplied plan/state values for the write-able / identity
// attributes (name, jsonblockkeyword, keyname_json_blockkeyword,
// jsonblockkeywordurl, iskeyregex_json_blockkeyword, jsonblockkeywordtype,
// comment, state, isautodeployed) so that a GET response (which may normalize or
// default these) does not produce an "inconsistent result after apply" diff. It
// only copies the server-managed read-only attributes (resourceid, alertonly)
// from the response. The ID is composed once in Create and never recomputed here
// (Pattern 6).
func appfwprofile_jsonblockkeyword_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileJsonblockkeywordBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileJsonblockkeywordBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_jsonblockkeyword_bindingSetAttrFromGet Function")

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

	// name, jsonblockkeyword, keyname_json_blockkeyword, jsonblockkeywordurl,
	// iskeyregex_json_blockkeyword, jsonblockkeywordtype, comment and state are
	// user-supplied (RequiresReplace) and are intentionally NOT overwritten from the
	// GET response - the plan/state value is authoritative (Pattern 7).

	return data
}

// appfwprofile_jsonblockkeyword_bindingSetAttrFromGetForDatasource is the
// DATASOURCE setter. The datasource has no prior plan/state to preserve, so it
// faithfully copies every attribute from the GET response and composes the ID.
func appfwprofile_jsonblockkeyword_bindingSetAttrFromGetForDatasource(ctx context.Context, data *AppfwprofileJsonblockkeywordBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileJsonblockkeywordBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_jsonblockkeyword_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["alertonly"]; ok && val != nil {
		data.Alertonly = types.StringValue(val.(string))
	} else {
		data.Alertonly = types.StringNull()
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
	if val, ok := getResponseData["iskeyregex_json_blockkeyword"]; ok && val != nil {
		data.IskeyregexJsonBlockkeyword = types.StringValue(val.(string))
	} else {
		data.IskeyregexJsonBlockkeyword = types.StringNull()
	}
	if val, ok := getResponseData["jsonblockkeyword"]; ok && val != nil {
		data.Jsonblockkeyword = types.StringValue(val.(string))
	} else {
		data.Jsonblockkeyword = types.StringNull()
	}
	if val, ok := getResponseData["jsonblockkeywordtype"]; ok && val != nil {
		data.Jsonblockkeywordtype = types.StringValue(val.(string))
	} else {
		data.Jsonblockkeywordtype = types.StringNull()
	}
	if val, ok := getResponseData["jsonblockkeywordurl"]; ok && val != nil {
		data.Jsonblockkeywordurl = types.StringValue(val.(string))
	} else {
		data.Jsonblockkeywordurl = types.StringNull()
	}
	if val, ok := getResponseData["keyname_json_blockkeyword"]; ok && val != nil {
		data.KeynameJsonBlockkeyword = types.StringValue(val.(string))
	} else {
		data.KeynameJsonBlockkeyword = types.StringNull()
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
	idParts = append(idParts, fmt.Sprintf("jsonblockkeyword:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Jsonblockkeyword.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("jsonblockkeywordurl:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Jsonblockkeywordurl.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("keyname_json_blockkeyword:%s", utils.UrlEncode(fmt.Sprintf("%v", data.KeynameJsonBlockkeyword.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
