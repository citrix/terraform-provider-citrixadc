package appfwprofile_excluderescontenttype_binding

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

// AppfwprofileExcluderescontenttypeBindingResourceModel describes the resource data model.
type AppfwprofileExcluderescontenttypeBindingResourceModel struct {
	Id                    types.String `tfsdk:"id"`
	Alertonly             types.String `tfsdk:"alertonly"`
	Comment               types.String `tfsdk:"comment"`
	Excluderescontenttype types.String `tfsdk:"excluderescontenttype"`
	Isautodeployed        types.String `tfsdk:"isautodeployed"`
	Name                  types.String `tfsdk:"name"`
	Resourceid            types.String `tfsdk:"resourceid"`
	State                 types.String `tfsdk:"state"`
}

func (r *AppfwprofileExcluderescontenttypeBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwprofile_excluderescontenttype_binding resource.",
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
			"excluderescontenttype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A regular expression that represents the content type of the response that are to be excluded from inspection.",
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
		},
	}
}

func appfwprofile_excluderescontenttype_bindingGetThePayloadFromtheConfig(ctx context.Context, data *AppfwprofileExcluderescontenttypeBindingResourceModel) appfw.Appfwprofileexcluderescontenttypebinding {
	tflog.Debug(ctx, "In appfwprofile_excluderescontenttype_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appfwprofile_excluderescontenttype_binding := appfw.Appfwprofileexcluderescontenttypebinding{}
	if !data.Alertonly.IsNull() {
		appfwprofile_excluderescontenttype_binding.Alertonly = data.Alertonly.ValueString()
	}
	if !data.Comment.IsNull() {
		appfwprofile_excluderescontenttype_binding.Comment = data.Comment.ValueString()
	}
	if !data.Excluderescontenttype.IsNull() {
		appfwprofile_excluderescontenttype_binding.Excluderescontenttype = data.Excluderescontenttype.ValueString()
	}
	if !data.Isautodeployed.IsNull() {
		appfwprofile_excluderescontenttype_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.Name.IsNull() {
		appfwprofile_excluderescontenttype_binding.Name = data.Name.ValueString()
	}
	if !data.Resourceid.IsNull() {
		appfwprofile_excluderescontenttype_binding.Resourceid = data.Resourceid.ValueString()
	}
	if !data.State.IsNull() {
		appfwprofile_excluderescontenttype_binding.State = data.State.ValueString()
	}

	return appfwprofile_excluderescontenttype_binding
}

func appfwprofile_excluderescontenttype_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileExcluderescontenttypeBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileExcluderescontenttypeBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_excluderescontenttype_bindingSetAttrFromGet Function")

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
	if val, ok := getResponseData["excluderescontenttype"]; ok && val != nil {
		data.Excluderescontenttype = types.StringValue(val.(string))
	} else {
		data.Excluderescontenttype = types.StringNull()
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

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("excluderescontenttype:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Excluderescontenttype.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
