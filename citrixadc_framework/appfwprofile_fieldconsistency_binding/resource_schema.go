package appfwprofile_fieldconsistency_binding

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

// AppfwprofileFieldconsistencyBindingResourceModel describes the resource data model.
type AppfwprofileFieldconsistencyBindingResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Alertonly        types.String `tfsdk:"alertonly"`
	Comment          types.String `tfsdk:"comment"`
	Fieldconsistency types.String `tfsdk:"fieldconsistency"`
	FormactionurlFfc types.String `tfsdk:"formactionurl_ffc"`
	Isautodeployed   types.String `tfsdk:"isautodeployed"`
	IsregexFfc       types.String `tfsdk:"isregex_ffc"`
	Name             types.String `tfsdk:"name"`
	Resourceid       types.String `tfsdk:"resourceid"`
	State            types.String `tfsdk:"state"`
}

func (r *AppfwprofileFieldconsistencyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwprofile_fieldconsistency_binding resource.",
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
			"fieldconsistency": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The web form field name.",
			},
			"formactionurl_ffc": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The web form action URL.",
			},
			"isautodeployed": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the rule auto deployed by dynamic profile ?",
			},
			"isregex_ffc": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the web form field name a regular expression?",
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
		},
	}
}

func appfwprofile_fieldconsistency_bindingGetThePayloadFromtheConfig(ctx context.Context, data *AppfwprofileFieldconsistencyBindingResourceModel) appfw.Appfwprofilefieldconsistencybinding {
	tflog.Debug(ctx, "In appfwprofile_fieldconsistency_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appfwprofile_fieldconsistency_binding := appfw.Appfwprofilefieldconsistencybinding{}
	if !data.Alertonly.IsNull() {
		appfwprofile_fieldconsistency_binding.Alertonly = data.Alertonly.ValueString()
	}
	if !data.Comment.IsNull() {
		appfwprofile_fieldconsistency_binding.Comment = data.Comment.ValueString()
	}
	if !data.Fieldconsistency.IsNull() {
		appfwprofile_fieldconsistency_binding.Fieldconsistency = data.Fieldconsistency.ValueString()
	}
	if !data.FormactionurlFfc.IsNull() {
		appfwprofile_fieldconsistency_binding.Formactionurlffc = data.FormactionurlFfc.ValueString()
	}
	if !data.Isautodeployed.IsNull() {
		appfwprofile_fieldconsistency_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.IsregexFfc.IsNull() {
		appfwprofile_fieldconsistency_binding.Isregexffc = data.IsregexFfc.ValueString()
	}
	if !data.Name.IsNull() {
		appfwprofile_fieldconsistency_binding.Name = data.Name.ValueString()
	}
	if !data.Resourceid.IsNull() {
		appfwprofile_fieldconsistency_binding.Resourceid = data.Resourceid.ValueString()
	}
	if !data.State.IsNull() {
		appfwprofile_fieldconsistency_binding.State = data.State.ValueString()
	}

	return appfwprofile_fieldconsistency_binding
}

func appfwprofile_fieldconsistency_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileFieldconsistencyBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileFieldconsistencyBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_fieldconsistency_bindingSetAttrFromGet Function")

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
	if val, ok := getResponseData["fieldconsistency"]; ok && val != nil {
		data.Fieldconsistency = types.StringValue(val.(string))
	} else {
		data.Fieldconsistency = types.StringNull()
	}
	if val, ok := getResponseData["formactionurl_ffc"]; ok && val != nil {
		data.FormactionurlFfc = types.StringValue(val.(string))
	} else {
		data.FormactionurlFfc = types.StringNull()
	}
	if val, ok := getResponseData["isautodeployed"]; ok && val != nil {
		data.Isautodeployed = types.StringValue(val.(string))
	} else {
		data.Isautodeployed = types.StringNull()
	}
	if val, ok := getResponseData["isregex_ffc"]; ok && val != nil {
		data.IsregexFfc = types.StringValue(val.(string))
	} else {
		data.IsregexFfc = types.StringNull()
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

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("fieldconsistency:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Fieldconsistency.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("formactionurl_ffc:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.FormactionurlFfc.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
