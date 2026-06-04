package appfwprofile_appfwconfidfield_binding

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

// AppfwprofileAppfwconfidfieldBindingResourceModel describes the resource data model.
type AppfwprofileAppfwconfidfieldBindingResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Alertonly      types.String `tfsdk:"alertonly"`
	CffieldUrl     types.String `tfsdk:"cffield_url"`
	Comment        types.String `tfsdk:"comment"`
	Confidfield    types.String `tfsdk:"confidfield"`
	Isautodeployed types.String `tfsdk:"isautodeployed"`
	IsregexCffield types.String `tfsdk:"isregex_cffield"`
	Name           types.String `tfsdk:"name"`
	Resourceid     types.String `tfsdk:"resourceid"`
	State          types.String `tfsdk:"state"`
}

func (r *AppfwprofileAppfwconfidfieldBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwprofile_appfwconfidfield_binding resource.",
			},
			"alertonly": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Send SNMP alert?",
			},
			"cffield_url": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "URL of the web page that contains the web form.",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments about the purpose of profile, or other useful information about the profile.",
			},
			"confidfield": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the form field to designate as confidential.",
			},
			"isautodeployed": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Is the rule auto deployed by dynamic profile ?",
			},
			"isregex_cffield": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Is Fake Account Detection field name regular expression?",
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

func appfwprofile_appfwconfidfield_bindingGetThePayloadFromthePlan(ctx context.Context, data *AppfwprofileAppfwconfidfieldBindingResourceModel) appfw.Appfwprofileappfwconfidfieldbinding {
	tflog.Debug(ctx, "In appfwprofile_appfwconfidfield_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	appfwprofile_appfwconfidfield_binding := appfw.Appfwprofileappfwconfidfieldbinding{}
	// NOTE (Pattern 15 sanctioned exclusion): alertonly, isautodeployed and resourceid
	// are read-only / cross-branch GET-response attributes for the confidField bind branch.
	// The CLI does not accept them for `bind appfw profile ... -confidField ...`, so they
	// are kept Computed in the schema (populated from GET) but NOT sent in the write payload.
	if !data.CffieldUrl.IsNull() && !data.CffieldUrl.IsUnknown() {
		appfwprofile_appfwconfidfield_binding.Cffieldurl = data.CffieldUrl.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appfwprofile_appfwconfidfield_binding.Comment = data.Comment.ValueString()
	}
	if !data.Confidfield.IsNull() && !data.Confidfield.IsUnknown() {
		appfwprofile_appfwconfidfield_binding.Confidfield = data.Confidfield.ValueString()
	}
	if !data.IsregexCffield.IsNull() && !data.IsregexCffield.IsUnknown() {
		appfwprofile_appfwconfidfield_binding.Isregexcffield = data.IsregexCffield.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwprofile_appfwconfidfield_binding.Name = data.Name.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		appfwprofile_appfwconfidfield_binding.State = data.State.ValueString()
	}

	return appfwprofile_appfwconfidfield_binding
}

// appfwprofile_appfwconfidfield_bindingSetAttrFromGet is the RESOURCE setter.
// It preserves user-supplied plan/state values for the write-able / identity
// attributes (name, confidfield, cffield_url, comment, isregex_cffield, state)
// so that a GET response (which may normalize or default these) does not produce
// an "inconsistent result after apply" diff. It only copies the server-managed
// read-only attributes (resourceid, isautodeployed, alertonly) from the response.
// The ID is composed once in Create and never recomputed here (Pattern 6).
func appfwprofile_appfwconfidfield_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileAppfwconfidfieldBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileAppfwconfidfieldBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_appfwconfidfield_bindingSetAttrFromGet Function")

	// Server-managed read-only attributes - safe to copy from the GET response.
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

	// name, confidfield, cffield_url, comment, isregex_cffield and state are
	// user-supplied (RequiresReplace) and are intentionally NOT overwritten from
	// the GET response - the plan/state value is authoritative (Pattern 7).

	return data
}

// appfwprofile_appfwconfidfield_bindingSetAttrFromGetForDatasource is the
// DATASOURCE setter. The datasource has no prior plan/state to preserve, so it
// faithfully copies every attribute from the GET response and composes the ID.
func appfwprofile_appfwconfidfield_bindingSetAttrFromGetForDatasource(ctx context.Context, data *AppfwprofileAppfwconfidfieldBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileAppfwconfidfieldBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_appfwconfidfield_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["alertonly"]; ok && val != nil {
		data.Alertonly = types.StringValue(val.(string))
	} else {
		data.Alertonly = types.StringNull()
	}
	if val, ok := getResponseData["cffield_url"]; ok && val != nil {
		data.CffieldUrl = types.StringValue(val.(string))
	} else {
		data.CffieldUrl = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["confidfield"]; ok && val != nil {
		data.Confidfield = types.StringValue(val.(string))
	} else {
		data.Confidfield = types.StringNull()
	}
	if val, ok := getResponseData["isautodeployed"]; ok && val != nil {
		data.Isautodeployed = types.StringValue(val.(string))
	} else {
		data.Isautodeployed = types.StringNull()
	}
	if val, ok := getResponseData["isregex_cffield"]; ok && val != nil {
		data.IsregexCffield = types.StringValue(val.(string))
	} else {
		data.IsregexCffield = types.StringNull()
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
	idParts = append(idParts, fmt.Sprintf("cffield_url:%s", utils.UrlEncode(fmt.Sprintf("%v", data.CffieldUrl.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("confidfield:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Confidfield.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
