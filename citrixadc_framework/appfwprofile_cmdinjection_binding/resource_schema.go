package appfwprofile_cmdinjection_binding

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

// AppfwprofileCmdinjectionBindingResourceModel describes the resource data model.
type AppfwprofileCmdinjectionBindingResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Alertonly         types.String `tfsdk:"alertonly"`
	AsScanLocationCmd types.String `tfsdk:"as_scan_location_cmd"`
	AsValueExprCmd    types.String `tfsdk:"as_value_expr_cmd"`
	AsValueTypeCmd    types.String `tfsdk:"as_value_type_cmd"`
	Cmdinjection      types.String `tfsdk:"cmdinjection"`
	Comment           types.String `tfsdk:"comment"`
	FormactionurlCmd  types.String `tfsdk:"formactionurl_cmd"`
	Isautodeployed    types.String `tfsdk:"isautodeployed"`
	IsregexCmd        types.String `tfsdk:"isregex_cmd"`
	IsvalueregexCmd   types.String `tfsdk:"isvalueregex_cmd"`
	Name              types.String `tfsdk:"name"`
	Resourceid        types.String `tfsdk:"resourceid"`
	State             types.String `tfsdk:"state"`
}

func (r *AppfwprofileCmdinjectionBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwprofile_cmdinjection_binding resource.",
			},
			"alertonly": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Send SNMP alert?",
			},
			"as_scan_location_cmd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Location of command injection exception - form field, header or cookie.",
			},
			"as_value_expr_cmd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The web form/header/cookie value expression.",
			},
			"as_value_type_cmd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of the relaxed web form value",
			},
			"cmdinjection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the relaxed web form field/header/cookie",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments about the purpose of profile, or other useful information about the profile.",
			},
			"formactionurl_cmd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The web form action URL.",
			},
			"isautodeployed": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the rule auto deployed by dynamic profile ?",
			},
			"isregex_cmd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the relaxed web form field name/header/cookie a regular expression?",
			},
			"isvalueregex_cmd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the web form field/header/cookie value a regular expression?",
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

func appfwprofile_cmdinjection_bindingGetThePayloadFromtheConfig(ctx context.Context, data *AppfwprofileCmdinjectionBindingResourceModel) appfw.Appfwprofilecmdinjectionbinding {
	tflog.Debug(ctx, "In appfwprofile_cmdinjection_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appfwprofile_cmdinjection_binding := appfw.Appfwprofilecmdinjectionbinding{}
	if !data.Alertonly.IsNull() {
		appfwprofile_cmdinjection_binding.Alertonly = data.Alertonly.ValueString()
	}
	if !data.AsScanLocationCmd.IsNull() {
		appfwprofile_cmdinjection_binding.Asscanlocationcmd = data.AsScanLocationCmd.ValueString()
	}
	if !data.AsValueExprCmd.IsNull() {
		appfwprofile_cmdinjection_binding.Asvalueexprcmd = data.AsValueExprCmd.ValueString()
	}
	if !data.AsValueTypeCmd.IsNull() {
		appfwprofile_cmdinjection_binding.Asvaluetypecmd = data.AsValueTypeCmd.ValueString()
	}
	if !data.Cmdinjection.IsNull() {
		appfwprofile_cmdinjection_binding.Cmdinjection = data.Cmdinjection.ValueString()
	}
	if !data.Comment.IsNull() {
		appfwprofile_cmdinjection_binding.Comment = data.Comment.ValueString()
	}
	if !data.FormactionurlCmd.IsNull() {
		appfwprofile_cmdinjection_binding.Formactionurlcmd = data.FormactionurlCmd.ValueString()
	}
	if !data.Isautodeployed.IsNull() {
		appfwprofile_cmdinjection_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.IsregexCmd.IsNull() {
		appfwprofile_cmdinjection_binding.Isregexcmd = data.IsregexCmd.ValueString()
	}
	if !data.IsvalueregexCmd.IsNull() {
		appfwprofile_cmdinjection_binding.Isvalueregexcmd = data.IsvalueregexCmd.ValueString()
	}
	if !data.Name.IsNull() {
		appfwprofile_cmdinjection_binding.Name = data.Name.ValueString()
	}
	if !data.Resourceid.IsNull() {
		appfwprofile_cmdinjection_binding.Resourceid = data.Resourceid.ValueString()
	}
	if !data.State.IsNull() {
		appfwprofile_cmdinjection_binding.State = data.State.ValueString()
	}

	return appfwprofile_cmdinjection_binding
}

func appfwprofile_cmdinjection_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileCmdinjectionBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileCmdinjectionBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_cmdinjection_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["alertonly"]; ok && val != nil {
		data.Alertonly = types.StringValue(val.(string))
	} else {
		data.Alertonly = types.StringNull()
	}
	if val, ok := getResponseData["as_scan_location_cmd"]; ok && val != nil {
		data.AsScanLocationCmd = types.StringValue(val.(string))
	} else {
		data.AsScanLocationCmd = types.StringNull()
	}
	if val, ok := getResponseData["as_value_expr_cmd"]; ok && val != nil {
		data.AsValueExprCmd = types.StringValue(val.(string))
	} else {
		data.AsValueExprCmd = types.StringNull()
	}
	if val, ok := getResponseData["as_value_type_cmd"]; ok && val != nil {
		data.AsValueTypeCmd = types.StringValue(val.(string))
	} else {
		data.AsValueTypeCmd = types.StringNull()
	}
	if val, ok := getResponseData["cmdinjection"]; ok && val != nil {
		data.Cmdinjection = types.StringValue(val.(string))
	} else {
		data.Cmdinjection = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["formactionurl_cmd"]; ok && val != nil {
		data.FormactionurlCmd = types.StringValue(val.(string))
	} else {
		data.FormactionurlCmd = types.StringNull()
	}
	if val, ok := getResponseData["isautodeployed"]; ok && val != nil {
		data.Isautodeployed = types.StringValue(val.(string))
	} else {
		data.Isautodeployed = types.StringNull()
	}
	if val, ok := getResponseData["isregex_cmd"]; ok && val != nil {
		data.IsregexCmd = types.StringValue(val.(string))
	} else {
		data.IsregexCmd = types.StringNull()
	}
	if val, ok := getResponseData["isvalueregex_cmd"]; ok && val != nil {
		data.IsvalueregexCmd = types.StringValue(val.(string))
	} else {
		data.IsvalueregexCmd = types.StringNull()
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
	idParts = append(idParts, fmt.Sprintf("as_scan_location_cmd:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.AsScanLocationCmd.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("as_value_expr_cmd:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.AsValueExprCmd.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("as_value_type_cmd:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.AsValueTypeCmd.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("cmdinjection:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Cmdinjection.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("formactionurl_cmd:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.FormactionurlCmd.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
