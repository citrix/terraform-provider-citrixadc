package appfwprofile_csrftag_binding

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

// AppfwprofileCsrftagBindingResourceModel describes the resource data model.
type AppfwprofileCsrftagBindingResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Alertonly         types.String `tfsdk:"alertonly"`
	Comment           types.String `tfsdk:"comment"`
	Csrfformactionurl types.String `tfsdk:"csrfformactionurl"`
	Csrftag           types.String `tfsdk:"csrftag"`
	Isautodeployed    types.String `tfsdk:"isautodeployed"`
	Name              types.String `tfsdk:"name"`
	Resourceid        types.String `tfsdk:"resourceid"`
	Ruletype          types.String `tfsdk:"ruletype"`
	State             types.String `tfsdk:"state"`
}

func (r *AppfwprofileCsrftagBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwprofile_csrftag_binding resource.",
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
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments about the purpose of profile, or other useful information about the profile.",
			},
			"csrfformactionurl": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The web form action URL.",
			},
			"csrftag": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The web form originating URL.",
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
			"ruletype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Specifies rule type of binding.",
			},
			"state": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Enabled.",
			},
		},
	}
}

func appfwprofile_csrftag_bindingGetThePayloadFromthePlan(ctx context.Context, data *AppfwprofileCsrftagBindingResourceModel) appfw.Appfwprofilecsrftagbinding {
	tflog.Debug(ctx, "In appfwprofile_csrftag_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	appfwprofile_csrftag_binding := appfw.Appfwprofilecsrftagbinding{}
	if !data.Alertonly.IsNull() && !data.Alertonly.IsUnknown() {
		appfwprofile_csrftag_binding.Alertonly = data.Alertonly.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appfwprofile_csrftag_binding.Comment = data.Comment.ValueString()
	}
	if !data.Csrfformactionurl.IsNull() && !data.Csrfformactionurl.IsUnknown() {
		appfwprofile_csrftag_binding.Csrfformactionurl = data.Csrfformactionurl.ValueString()
	}
	if !data.Csrftag.IsNull() && !data.Csrftag.IsUnknown() {
		appfwprofile_csrftag_binding.Csrftag = data.Csrftag.ValueString()
	}
	if !data.Isautodeployed.IsNull() && !data.Isautodeployed.IsUnknown() {
		appfwprofile_csrftag_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwprofile_csrftag_binding.Name = data.Name.ValueString()
	}
	if !data.Resourceid.IsNull() && !data.Resourceid.IsUnknown() {
		appfwprofile_csrftag_binding.Resourceid = data.Resourceid.ValueString()
	}
	if !data.Ruletype.IsNull() && !data.Ruletype.IsUnknown() {
		appfwprofile_csrftag_binding.Ruletype = data.Ruletype.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		appfwprofile_csrftag_binding.State = data.State.ValueString()
	}

	return appfwprofile_csrftag_binding
}

// appfwprofile_csrftag_bindingSetAttrFromGet is used by the resource Read/Create/Update.
// It preserves the user-supplied identity-key attributes (name, csrftag, csrfformactionurl)
// from the plan/state so a server-normalized echo does not trigger an
// "inconsistent result after apply" error, while adopting server values for the
// non-key attributes.
func appfwprofile_csrftag_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileCsrftagBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileCsrftagBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_csrftag_bindingSetAttrFromGet Function")

	// Convert API response to model
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
	// Identity-key attributes (csrfformactionurl, csrftag, name) are preserved from
	// the existing plan/state; never overwrite with a possibly-normalized server echo.
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
	if val, ok := getResponseData["ruletype"]; ok && val != nil {
		data.Ruletype = types.StringValue(val.(string))
	} else {
		data.Ruletype = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}

	return data
}

// appfwprofile_csrftag_bindingSetAttrFromGetForDatasource faithfully copies every field
// from the GET response and composes the ID, for use by the datasource Read (which has
// no prior plan/state to preserve).
func appfwprofile_csrftag_bindingSetAttrFromGetForDatasource(ctx context.Context, data *AppfwprofileCsrftagBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileCsrftagBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_csrftag_bindingSetAttrFromGetForDatasource Function")

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
	if val, ok := getResponseData["csrfformactionurl"]; ok && val != nil {
		data.Csrfformactionurl = types.StringValue(val.(string))
	} else {
		data.Csrfformactionurl = types.StringNull()
	}
	if val, ok := getResponseData["csrftag"]; ok && val != nil {
		data.Csrftag = types.StringValue(val.(string))
	} else {
		data.Csrftag = types.StringNull()
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
	if val, ok := getResponseData["ruletype"]; ok && val != nil {
		data.Ruletype = types.StringValue(val.(string))
	} else {
		data.Ruletype = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}

	// Set ID for the datasource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("csrfformactionurl:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Csrfformactionurl.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("csrftag:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Csrftag.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
