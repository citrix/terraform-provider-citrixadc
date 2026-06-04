package appfwprofile_blockkeyword_binding

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

// AppfwprofileBlockkeywordBindingResourceModel describes the resource data model.
type AppfwprofileBlockkeywordBindingResourceModel struct {
	Id                             types.String `tfsdk:"id"`
	Alertonly                      types.String `tfsdk:"alertonly"`
	AsBlockkeywordFormurl          types.String `tfsdk:"as_blockkeyword_formurl"`
	AsFieldnameIsregexBlockkeyword types.String `tfsdk:"as_fieldname_isregex_blockkeyword"`
	Blockkeyword                   types.String `tfsdk:"blockkeyword"`
	Blockkeywordtype               types.String `tfsdk:"blockkeywordtype"`
	Comment                        types.String `tfsdk:"comment"`
	Fieldname                      types.String `tfsdk:"fieldname"`
	Isautodeployed                 types.String `tfsdk:"isautodeployed"`
	Name                           types.String `tfsdk:"name"`
	Resourceid                     types.String `tfsdk:"resourceid"`
	State                          types.String `tfsdk:"state"`
}

func (r *AppfwprofileBlockkeywordBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwprofile_blockkeyword_binding resource.",
			},
			"alertonly": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Send SNMP alert?",
			},
			"as_blockkeyword_formurl": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The blockkeyword form action URL.",
			},
			"as_fieldname_isregex_blockkeyword": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Is block keyword field name regular expression?",
			},
			"blockkeyword": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Field name of the block keyword binding.",
			},
			"blockkeywordtype": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "block keyword type(literal|PCRE)",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments about the purpose of profile, or other useful information about the profile.",
			},
			"fieldname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "A block keyword field name",
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

func appfwprofile_blockkeyword_bindingGetThePayloadFromthePlan(ctx context.Context, data *AppfwprofileBlockkeywordBindingResourceModel) appfw.Appfwprofileblockkeywordbinding {
	tflog.Debug(ctx, "In appfwprofile_blockkeyword_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	// NOTE (Pattern 15 sanctioned exclusion): alertonly and resourceid are
	// read-only / server-assigned GET-response attributes for the blockKeyword
	// bind branch. The CLI does not accept them for
	// `bind appfw profile ... -blockKeyword ...`, so they are kept Computed in the
	// schema (populated from GET) but NOT sent in the write payload. (ruletype is a
	// cross-branch attribute and is not part of this resource's model at all.)
	appfwprofile_blockkeyword_binding := appfw.Appfwprofileblockkeywordbinding{}
	if !data.AsBlockkeywordFormurl.IsNull() && !data.AsBlockkeywordFormurl.IsUnknown() {
		appfwprofile_blockkeyword_binding.Asblockkeywordformurl = data.AsBlockkeywordFormurl.ValueString()
	}
	if !data.AsFieldnameIsregexBlockkeyword.IsNull() && !data.AsFieldnameIsregexBlockkeyword.IsUnknown() {
		appfwprofile_blockkeyword_binding.Asfieldnameisregexblockkeyword = data.AsFieldnameIsregexBlockkeyword.ValueString()
	}
	if !data.Blockkeyword.IsNull() && !data.Blockkeyword.IsUnknown() {
		appfwprofile_blockkeyword_binding.Blockkeyword = data.Blockkeyword.ValueString()
	}
	if !data.Blockkeywordtype.IsNull() && !data.Blockkeywordtype.IsUnknown() {
		appfwprofile_blockkeyword_binding.Blockkeywordtype = data.Blockkeywordtype.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appfwprofile_blockkeyword_binding.Comment = data.Comment.ValueString()
	}
	if !data.Fieldname.IsNull() && !data.Fieldname.IsUnknown() {
		appfwprofile_blockkeyword_binding.Fieldname = data.Fieldname.ValueString()
	}
	if !data.Isautodeployed.IsNull() && !data.Isautodeployed.IsUnknown() {
		appfwprofile_blockkeyword_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwprofile_blockkeyword_binding.Name = data.Name.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		appfwprofile_blockkeyword_binding.State = data.State.ValueString()
	}

	return appfwprofile_blockkeyword_binding
}

// appfwprofile_blockkeyword_bindingSetAttrFromGet is the RESOURCE setter.
// It preserves user-supplied plan/state values for the write-able / identity
// attributes (name, blockkeyword, fieldname, as_blockkeyword_formurl, comment,
// state, as_fieldname_isregex_blockkeyword, blockkeywordtype, isautodeployed) so
// that a GET response (which may normalize or default these) does not produce an
// "inconsistent result after apply" diff. It only copies the server-managed
// read-only attributes (resourceid, alertonly) from the response. The ID is
// composed once in Create and never recomputed here (Pattern 6).
func appfwprofile_blockkeyword_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileBlockkeywordBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileBlockkeywordBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_blockkeyword_bindingSetAttrFromGet Function")

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

	// name, blockkeyword, fieldname, as_blockkeyword_formurl, comment, state,
	// as_fieldname_isregex_blockkeyword and blockkeywordtype are user-supplied
	// (RequiresReplace) and are intentionally NOT overwritten from the GET response
	// - the plan/state value is authoritative (Pattern 7).

	return data
}

// appfwprofile_blockkeyword_bindingSetAttrFromGetForDatasource is the DATASOURCE
// setter. The datasource has no prior plan/state to preserve, so it faithfully
// copies every attribute from the GET response and composes the ID.
func appfwprofile_blockkeyword_bindingSetAttrFromGetForDatasource(ctx context.Context, data *AppfwprofileBlockkeywordBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileBlockkeywordBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_blockkeyword_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["alertonly"]; ok && val != nil {
		data.Alertonly = types.StringValue(val.(string))
	} else {
		data.Alertonly = types.StringNull()
	}
	if val, ok := getResponseData["as_blockkeyword_formurl"]; ok && val != nil {
		data.AsBlockkeywordFormurl = types.StringValue(val.(string))
	} else {
		data.AsBlockkeywordFormurl = types.StringNull()
	}
	if val, ok := getResponseData["as_fieldname_isregex_blockkeyword"]; ok && val != nil {
		data.AsFieldnameIsregexBlockkeyword = types.StringValue(val.(string))
	} else {
		data.AsFieldnameIsregexBlockkeyword = types.StringNull()
	}
	if val, ok := getResponseData["blockkeyword"]; ok && val != nil {
		data.Blockkeyword = types.StringValue(val.(string))
	} else {
		data.Blockkeyword = types.StringNull()
	}
	if val, ok := getResponseData["blockkeywordtype"]; ok && val != nil {
		data.Blockkeywordtype = types.StringValue(val.(string))
	} else {
		data.Blockkeywordtype = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["fieldname"]; ok && val != nil {
		data.Fieldname = types.StringValue(val.(string))
	} else {
		data.Fieldname = types.StringNull()
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
	idParts = append(idParts, fmt.Sprintf("as_blockkeyword_formurl:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsBlockkeywordFormurl.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("blockkeyword:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Blockkeyword.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("fieldname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Fieldname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
