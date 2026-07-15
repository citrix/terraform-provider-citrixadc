package appfwprofile_fakeaccount_binding

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

// AppfwprofileFakeaccountBindingResourceModel describes the resource data model.
type AppfwprofileFakeaccountBindingResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Alertonly        types.String `tfsdk:"alertonly"`
	Comment          types.String `tfsdk:"comment"`
	Fakeaccount      types.String `tfsdk:"fakeaccount"`
	Formexpression   types.String `tfsdk:"formexpression"`
	FormurlFad       types.String `tfsdk:"formurl_fad"`
	Isautodeployed   types.String `tfsdk:"isautodeployed"`
	Isfieldnameregex types.String `tfsdk:"isfieldnameregex"`
	Name             types.String `tfsdk:"name"`
	Resourceid       types.String `tfsdk:"resourceid"`
	State            types.String `tfsdk:"state"`
	Tag              types.String `tfsdk:"tag"`
}

func (r *AppfwprofileFakeaccountBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwprofile_fakeaccount_binding resource.",
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
			"fakeaccount": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Field name of the fake account rule.",
			},
			"formexpression": schema.StringAttribute{
				// CLI synopsis (man bind appfw profile, fakeAccount branch):
				//   [-formExpression <expression>]  => optional.
				// formExpression and formurl_fad are mutually exclusive at runtime
				// (NITRO errorcode 390: "Only one of FormURL or FormExpression is
				// allowed"). Enforced in ValidateConfig.
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "A regular expression that defines the Fake Account.",
			},
			"formurl_fad": schema.StringAttribute{
				// CLI synopsis (man bind appfw profile, fakeAccount branch):
				//   [-formURL <expression>]  => optional.
				// Mutually exclusive with formexpression (see above).
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The fake account detection URL.",
			},
			"isautodeployed": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Is the rule auto deployed by dynamic profile ?",
			},
			"isfieldnameregex": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Is Fake Account Detection field name regex?",
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
			"tag": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "A tag expression that defines the Fake Account.",
			},
		},
	}
}

func appfwprofile_fakeaccount_bindingGetThePayloadFromthePlan(ctx context.Context, data *AppfwprofileFakeaccountBindingResourceModel) appfw.Appfwprofilefakeaccountbinding {
	tflog.Debug(ctx, "In appfwprofile_fakeaccount_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model.
	// NOTE (Pattern 15 sanctioned exclusion): alertonly and resourceid are
	// read-only / server-assigned GET-response attributes for the fakeAccount
	// bind branch. The CLI does not accept them for
	// `bind appfw profile ... -fakeAccount ...`, so they are kept Computed in the
	// schema (populated from GET) but NOT sent in the write payload. (ruletype is a
	// cross-branch attribute and is not part of this resource's model at all, so
	// there is nothing to exclude for it here.)
	appfwprofile_fakeaccount_binding := appfw.Appfwprofilefakeaccountbinding{}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appfwprofile_fakeaccount_binding.Comment = data.Comment.ValueString()
	}
	if !data.Fakeaccount.IsNull() && !data.Fakeaccount.IsUnknown() {
		appfwprofile_fakeaccount_binding.Fakeaccount = data.Fakeaccount.ValueString()
	}
	if !data.Formexpression.IsNull() && !data.Formexpression.IsUnknown() {
		appfwprofile_fakeaccount_binding.Formexpression = data.Formexpression.ValueString()
	}
	if !data.FormurlFad.IsNull() && !data.FormurlFad.IsUnknown() {
		appfwprofile_fakeaccount_binding.Formurlfad = data.FormurlFad.ValueString()
	}
	if !data.Isautodeployed.IsNull() && !data.Isautodeployed.IsUnknown() {
		appfwprofile_fakeaccount_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.Isfieldnameregex.IsNull() && !data.Isfieldnameregex.IsUnknown() {
		appfwprofile_fakeaccount_binding.Isfieldnameregex = data.Isfieldnameregex.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwprofile_fakeaccount_binding.Name = data.Name.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		appfwprofile_fakeaccount_binding.State = data.State.ValueString()
	}
	if !data.Tag.IsNull() && !data.Tag.IsUnknown() {
		appfwprofile_fakeaccount_binding.Tag = data.Tag.ValueString()
	}

	return appfwprofile_fakeaccount_binding
}

// appfwprofile_fakeaccount_bindingSetAttrFromGet is the RESOURCE setter.
// It copies the server-managed read-only attributes (resourceid, alertonly,
// isautodeployed) plus the GET-echoed config attributes (isfieldnameregex, state,
// comment) from the response so they become known after apply AND round-trip on
// import (where there is no prior plan/state). The GET row for this binding
// echoes these values verbatim (no normalization), so populating them does not
// introduce an "inconsistent result after apply" diff. The identity /
// ID-component attributes (name, fakeaccount, tag, formexpression, formurl_fad)
// are backfilled from the parsed composite ID in readXFromApi (they are always
// recoverable from the ID). The ID is composed once in Create and never
// recomputed here (Pattern 6).
func appfwprofile_fakeaccount_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileFakeaccountBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileFakeaccountBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_fakeaccount_bindingSetAttrFromGet Function")

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

	// GET-echoed config attributes - the appliance returns these verbatim, so
	// copying them makes them round-trip on import without perturbing the basic
	// apply (config value == GET value).
	if val, ok := getResponseData["isfieldnameregex"]; ok && val != nil {
		data.Isfieldnameregex = types.StringValue(val.(string))
	} else {
		data.Isfieldnameregex = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}

	// name, fakeaccount, tag, formexpression and formurl_fad are the identity /
	// ID-component attributes and are backfilled from the parsed composite ID in
	// readXFromApi (Pattern 7), so they are NOT set from the GET response here.

	return data
}

// appfwprofile_fakeaccount_bindingSetAttrFromGetForDatasource is the DATASOURCE
// setter. The datasource has no prior plan/state to preserve, so it faithfully
// copies every attribute from the GET response and composes the ID.
func appfwprofile_fakeaccount_bindingSetAttrFromGetForDatasource(ctx context.Context, data *AppfwprofileFakeaccountBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileFakeaccountBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_fakeaccount_bindingSetAttrFromGetForDatasource Function")

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
	if val, ok := getResponseData["fakeaccount"]; ok && val != nil {
		data.Fakeaccount = types.StringValue(val.(string))
	} else {
		data.Fakeaccount = types.StringNull()
	}
	if val, ok := getResponseData["formexpression"]; ok && val != nil {
		data.Formexpression = types.StringValue(val.(string))
	} else {
		data.Formexpression = types.StringNull()
	}
	if val, ok := getResponseData["formurl_fad"]; ok && val != nil {
		data.FormurlFad = types.StringValue(val.(string))
	} else {
		data.FormurlFad = types.StringNull()
	}
	if val, ok := getResponseData["isautodeployed"]; ok && val != nil {
		data.Isautodeployed = types.StringValue(val.(string))
	} else {
		data.Isautodeployed = types.StringNull()
	}
	if val, ok := getResponseData["isfieldnameregex"]; ok && val != nil {
		data.Isfieldnameregex = types.StringValue(val.(string))
	} else {
		data.Isfieldnameregex = types.StringNull()
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
	if val, ok := getResponseData["tag"]; ok && val != nil {
		data.Tag = types.StringValue(val.(string))
	} else {
		data.Tag = types.StringNull()
	}

	// Compose the ID (datasource has no Create). formexpression and formurl_fad
	// are an at-most-one pair; only the populated arm is included so the absent
	// arm is never emitted as an empty "key:" segment.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("fakeaccount:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Fakeaccount.ValueString()))))
	if !data.Formexpression.IsNull() && data.Formexpression.ValueString() != "" {
		idParts = append(idParts, fmt.Sprintf("formexpression:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Formexpression.ValueString()))))
	}
	if !data.FormurlFad.IsNull() && data.FormurlFad.ValueString() != "" {
		idParts = append(idParts, fmt.Sprintf("formurl_fad:%s", utils.UrlEncode(fmt.Sprintf("%v", data.FormurlFad.ValueString()))))
	}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("tag:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Tag.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
