package appfwprofile_trustedlearningclients_binding

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

// AppfwprofileTrustedlearningclientsBindingResourceModel describes the resource data model.
type AppfwprofileTrustedlearningclientsBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Alertonly              types.String `tfsdk:"alertonly"`
	Comment                types.String `tfsdk:"comment"`
	Isautodeployed         types.String `tfsdk:"isautodeployed"`
	Name                   types.String `tfsdk:"name"`
	Resourceid             types.String `tfsdk:"resourceid"`
	Ruletype               types.String `tfsdk:"ruletype"`
	State                  types.String `tfsdk:"state"`
	Trustedlearningclients types.String `tfsdk:"trustedlearningclients"`
}

func (r *AppfwprofileTrustedlearningclientsBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwprofile_trustedlearningclients_binding resource.",
			},
			"alertonly": schema.StringAttribute{
				// GET overrides the configured value (e.g. ON -> OFF); preserve the
				// configured value in state, so Computed is dropped (Pattern 7).
				Optional: true,
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
			"isautodeployed": schema.StringAttribute{
				// GET overrides the configured value (e.g. AUTODEPLOYED ->
				// NOTAUTODEPLOYED); preserve the configured value (Pattern 7).
				Optional: true,
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
				// GET never echoes resourceid; preserve the configured value and drop
				// Computed to avoid known-after-apply churn (Pattern 7).
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "A \"id\" that identifies the rule.",
			},
			"ruletype": schema.StringAttribute{
				// Present in SDK v2 + NITRO struct; GET never echoes it, so Optional
				// only (re-added per migration family pattern a).
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Specifies rule type of binding",
			},
			"state": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Enabled.",
			},
			"trustedlearningclients": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Specify trusted host/network IP",
			},
		},
	}
}

func appfwprofile_trustedlearningclients_bindingGetThePayloadFromthePlan(ctx context.Context, data *AppfwprofileTrustedlearningclientsBindingResourceModel) appfw.Appfwprofiletrustedlearningclientsbinding {
	tflog.Debug(ctx, "In appfwprofile_trustedlearningclients_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	appfwprofile_trustedlearningclients_binding := appfw.Appfwprofiletrustedlearningclientsbinding{}
	if !data.Alertonly.IsNull() && !data.Alertonly.IsUnknown() {
		appfwprofile_trustedlearningclients_binding.Alertonly = data.Alertonly.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appfwprofile_trustedlearningclients_binding.Comment = data.Comment.ValueString()
	}
	if !data.Isautodeployed.IsNull() && !data.Isautodeployed.IsUnknown() {
		appfwprofile_trustedlearningclients_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwprofile_trustedlearningclients_binding.Name = data.Name.ValueString()
	}
	if !data.Resourceid.IsNull() && !data.Resourceid.IsUnknown() {
		appfwprofile_trustedlearningclients_binding.Resourceid = data.Resourceid.ValueString()
	}
	if !data.Ruletype.IsNull() && !data.Ruletype.IsUnknown() {
		appfwprofile_trustedlearningclients_binding.Ruletype = data.Ruletype.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		appfwprofile_trustedlearningclients_binding.State = data.State.ValueString()
	}
	if !data.Trustedlearningclients.IsNull() && !data.Trustedlearningclients.IsUnknown() {
		appfwprofile_trustedlearningclients_binding.Trustedlearningclients = data.Trustedlearningclients.ValueString()
	}

	return appfwprofile_trustedlearningclients_binding
}

// appfwprofile_trustedlearningclients_bindingSetAttrFromGet is the RESOURCE setter.
// It only reads back fields the NITRO GET echoes faithfully (name, trustedlearningclients,
// state, comment). It deliberately PRESERVES the plan/state value for:
//   - alertonly, isautodeployed: server overrides the configured value (Pattern 7).
//   - resourceid, ruletype: never echoed by GET (Pattern 7).
//
// This mirrors the SDK v2 resource which commented out d.Set for alertonly/isautodeployed.
// ID is set once in Create, so it is not recomputed here (Pattern 6).
func appfwprofile_trustedlearningclients_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileTrustedlearningclientsBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileTrustedlearningclientsBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_trustedlearningclients_bindingSetAttrFromGet Function")

	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["trustedlearningclients"]; ok && val != nil {
		data.Trustedlearningclients = types.StringValue(val.(string))
	}

	// Re-derive the canonical id so a legacy SDK v2 id is upgraded to the new format on Read.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("trustedlearningclients:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Trustedlearningclients.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}

// appfwprofile_trustedlearningclients_bindingSetAttrFromGetForDatasource is the
// DATASOURCE setter (Pattern 7 split). It faithfully copies every field the GET
// response returns and sets the composite ID, since the datasource has no Create.
func appfwprofile_trustedlearningclients_bindingSetAttrFromGetForDatasource(ctx context.Context, data *AppfwprofileTrustedlearningclientsBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileTrustedlearningclientsBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_trustedlearningclients_bindingSetAttrFromGetForDatasource Function")

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
	if val, ok := getResponseData["trustedlearningclients"]; ok && val != nil {
		data.Trustedlearningclients = types.StringValue(val.(string))
	} else {
		data.Trustedlearningclients = types.StringNull()
	}

	// Set composite ID for the datasource
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("trustedlearningclients:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Trustedlearningclients.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
