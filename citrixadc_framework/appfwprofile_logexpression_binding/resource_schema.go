package appfwprofile_logexpression_binding

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

// AppfwprofileLogexpressionBindingResourceModel describes the resource data model.
type AppfwprofileLogexpressionBindingResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Alertonly       types.String `tfsdk:"alertonly"`
	AsLogexpression types.String `tfsdk:"as_logexpression"`
	Comment         types.String `tfsdk:"comment"`
	Isautodeployed  types.String `tfsdk:"isautodeployed"`
	Logexpression   types.String `tfsdk:"logexpression"`
	Name            types.String `tfsdk:"name"`
	Resourceid      types.String `tfsdk:"resourceid"`
	State           types.String `tfsdk:"state"`
}

func (r *AppfwprofileLogexpressionBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwprofile_logexpression_binding resource.",
			},
			"alertonly": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Send SNMP alert?",
			},
			"as_logexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "LogExpression to log when violation happened on appfw profile",
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
			"logexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of LogExpression object.",
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

func appfwprofile_logexpression_bindingGetThePayloadFromtheConfig(ctx context.Context, data *AppfwprofileLogexpressionBindingResourceModel) appfw.Appfwprofilelogexpressionbinding {
	tflog.Debug(ctx, "In appfwprofile_logexpression_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appfwprofile_logexpression_binding := appfw.Appfwprofilelogexpressionbinding{}
	if !data.Alertonly.IsNull() {
		appfwprofile_logexpression_binding.Alertonly = data.Alertonly.ValueString()
	}
	if !data.AsLogexpression.IsNull() {
		appfwprofile_logexpression_binding.Aslogexpression = data.AsLogexpression.ValueString()
	}
	if !data.Comment.IsNull() {
		appfwprofile_logexpression_binding.Comment = data.Comment.ValueString()
	}
	if !data.Isautodeployed.IsNull() {
		appfwprofile_logexpression_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.Logexpression.IsNull() {
		appfwprofile_logexpression_binding.Logexpression = data.Logexpression.ValueString()
	}
	if !data.Name.IsNull() {
		appfwprofile_logexpression_binding.Name = data.Name.ValueString()
	}
	if !data.Resourceid.IsNull() {
		appfwprofile_logexpression_binding.Resourceid = data.Resourceid.ValueString()
	}
	if !data.State.IsNull() {
		appfwprofile_logexpression_binding.State = data.State.ValueString()
	}

	return appfwprofile_logexpression_binding
}

func appfwprofile_logexpression_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileLogexpressionBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileLogexpressionBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_logexpression_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["alertonly"]; ok && val != nil {
		data.Alertonly = types.StringValue(val.(string))
	} else {
		data.Alertonly = types.StringNull()
	}
	if val, ok := getResponseData["as_logexpression"]; ok && val != nil {
		data.AsLogexpression = types.StringValue(val.(string))
	} else {
		data.AsLogexpression = types.StringNull()
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
	if val, ok := getResponseData["logexpression"]; ok && val != nil {
		data.Logexpression = types.StringValue(val.(string))
	} else {
		data.Logexpression = types.StringNull()
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
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("logexpression:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Logexpression.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
