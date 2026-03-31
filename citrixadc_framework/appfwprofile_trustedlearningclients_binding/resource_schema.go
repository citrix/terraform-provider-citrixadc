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
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Send SNMP alert?",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments about the purpose of profile, or other useful information about the profile.",
			},
			"isautodeployed": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the rule auto deployed by dynamic profile ?",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the profile to which to bind an exemption or rule.",
			},
			"resourceid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A \"id\" that identifies the rule.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enabled.",
			},
			"trustedlearningclients": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify trusted host/network IP",
			},
		},
	}
}

func appfwprofile_trustedlearningclients_bindingGetThePayloadFromtheConfig(ctx context.Context, data *AppfwprofileTrustedlearningclientsBindingResourceModel) appfw.Appfwprofiletrustedlearningclientsbinding {
	tflog.Debug(ctx, "In appfwprofile_trustedlearningclients_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appfwprofile_trustedlearningclients_binding := appfw.Appfwprofiletrustedlearningclientsbinding{}
	if !data.Alertonly.IsNull() {
		appfwprofile_trustedlearningclients_binding.Alertonly = data.Alertonly.ValueString()
	}
	if !data.Comment.IsNull() {
		appfwprofile_trustedlearningclients_binding.Comment = data.Comment.ValueString()
	}
	if !data.Isautodeployed.IsNull() {
		appfwprofile_trustedlearningclients_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.Name.IsNull() {
		appfwprofile_trustedlearningclients_binding.Name = data.Name.ValueString()
	}
	if !data.Resourceid.IsNull() {
		appfwprofile_trustedlearningclients_binding.Resourceid = data.Resourceid.ValueString()
	}
	if !data.State.IsNull() {
		appfwprofile_trustedlearningclients_binding.State = data.State.ValueString()
	}
	if !data.Trustedlearningclients.IsNull() {
		appfwprofile_trustedlearningclients_binding.Trustedlearningclients = data.Trustedlearningclients.ValueString()
	}

	return appfwprofile_trustedlearningclients_binding
}

func appfwprofile_trustedlearningclients_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileTrustedlearningclientsBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileTrustedlearningclientsBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_trustedlearningclients_bindingSetAttrFromGet Function")

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
	if val, ok := getResponseData["trustedlearningclients"]; ok && val != nil {
		data.Trustedlearningclients = types.StringValue(val.(string))
	} else {
		data.Trustedlearningclients = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("trustedlearningclients:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Trustedlearningclients.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
