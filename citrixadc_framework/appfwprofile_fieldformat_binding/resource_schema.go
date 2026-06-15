package appfwprofile_fieldformat_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AppfwprofileFieldformatBindingResourceModel describes the resource data model.
type AppfwprofileFieldformatBindingResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Alertonly            types.String `tfsdk:"alertonly"`
	Comment              types.String `tfsdk:"comment"`
	Fieldformat          types.String `tfsdk:"fieldformat"`
	Fieldformatmaxlength types.Int64  `tfsdk:"fieldformatmaxlength"`
	Fieldformatminlength types.Int64  `tfsdk:"fieldformatminlength"`
	Fieldtype            types.String `tfsdk:"fieldtype"`
	FormactionurlFf      types.String `tfsdk:"formactionurl_ff"`
	Isautodeployed       types.String `tfsdk:"isautodeployed"`
	IsregexFf            types.String `tfsdk:"isregexff"`
	Name                 types.String `tfsdk:"name"`
	Resourceid           types.String `tfsdk:"resourceid"`
	Ruletype             types.String `tfsdk:"ruletype"`
	State                types.String `tfsdk:"state"`
}

func (r *AppfwprofileFieldformatBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwprofile_fieldformat_binding resource.",
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
			"fieldformat": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the form field to which a field format will be assigned.",
			},
			"fieldformatmaxlength": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "The maximum allowed length for data in this form field.",
			},
			"fieldformatminlength": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "The minimum allowed length for data in this form field.",
			},
			"fieldtype": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The field type you are assigning to this form field.",
			},
			"formactionurl_ff": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Action URL of the form field to which a field format will be assigned.",
			},
			"isautodeployed": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Is the rule auto deployed by dynamic profile ?",
			},
			"isregexff": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Is the form field name a regular expression?",
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

func appfwprofile_fieldformat_bindingGetThePayloadFromthePlan(ctx context.Context, data *AppfwprofileFieldformatBindingResourceModel) appfw.Appfwprofilefieldformatbinding {
	tflog.Debug(ctx, "In appfwprofile_fieldformat_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	appfwprofile_fieldformat_binding := appfw.Appfwprofilefieldformatbinding{}
	if !data.Alertonly.IsNull() && !data.Alertonly.IsUnknown() {
		appfwprofile_fieldformat_binding.Alertonly = data.Alertonly.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appfwprofile_fieldformat_binding.Comment = data.Comment.ValueString()
	}
	if !data.Fieldformat.IsNull() && !data.Fieldformat.IsUnknown() {
		appfwprofile_fieldformat_binding.Fieldformat = data.Fieldformat.ValueString()
	}
	if !data.Fieldformatmaxlength.IsNull() && !data.Fieldformatmaxlength.IsUnknown() {
		appfwprofile_fieldformat_binding.Fieldformatmaxlength = utils.IntPtr(int(data.Fieldformatmaxlength.ValueInt64()))
	}
	if !data.Fieldformatminlength.IsNull() && !data.Fieldformatminlength.IsUnknown() {
		appfwprofile_fieldformat_binding.Fieldformatminlength = utils.IntPtr(int(data.Fieldformatminlength.ValueInt64()))
	}
	if !data.Fieldtype.IsNull() && !data.Fieldtype.IsUnknown() {
		appfwprofile_fieldformat_binding.Fieldtype = data.Fieldtype.ValueString()
	}
	if !data.FormactionurlFf.IsNull() && !data.FormactionurlFf.IsUnknown() {
		appfwprofile_fieldformat_binding.Formactionurlff = data.FormactionurlFf.ValueString()
	}
	if !data.Isautodeployed.IsNull() && !data.Isautodeployed.IsUnknown() {
		appfwprofile_fieldformat_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.IsregexFf.IsNull() && !data.IsregexFf.IsUnknown() {
		appfwprofile_fieldformat_binding.Isregexff = data.IsregexFf.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwprofile_fieldformat_binding.Name = data.Name.ValueString()
	}
	if !data.Resourceid.IsNull() && !data.Resourceid.IsUnknown() {
		appfwprofile_fieldformat_binding.Resourceid = data.Resourceid.ValueString()
	}
	if !data.Ruletype.IsNull() && !data.Ruletype.IsUnknown() {
		appfwprofile_fieldformat_binding.Ruletype = data.Ruletype.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		appfwprofile_fieldformat_binding.State = data.State.ValueString()
	}

	return appfwprofile_fieldformat_binding
}

func appfwprofile_fieldformat_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileFieldformatBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileFieldformatBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_fieldformat_bindingSetAttrFromGet Function")

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
	if val, ok := getResponseData["fieldformat"]; ok && val != nil {
		data.Fieldformat = types.StringValue(val.(string))
	} else {
		data.Fieldformat = types.StringNull()
	}
	if val, ok := getResponseData["fieldformatmaxlength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Fieldformatmaxlength = types.Int64Value(intVal)
		}
	} else {
		data.Fieldformatmaxlength = types.Int64Null()
	}
	if val, ok := getResponseData["fieldformatminlength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Fieldformatminlength = types.Int64Value(intVal)
		}
	} else {
		data.Fieldformatminlength = types.Int64Null()
	}
	if val, ok := getResponseData["fieldtype"]; ok && val != nil {
		data.Fieldtype = types.StringValue(val.(string))
	} else {
		data.Fieldtype = types.StringNull()
	}
	if val, ok := getResponseData["formactionurl_ff"]; ok && val != nil {
		data.FormactionurlFf = types.StringValue(val.(string))
	} else {
		data.FormactionurlFf = types.StringNull()
	}
	if val, ok := getResponseData["isautodeployed"]; ok && val != nil {
		data.Isautodeployed = types.StringValue(val.(string))
	} else {
		data.Isautodeployed = types.StringNull()
	}
	if val, ok := getResponseData["isregex_ff"]; ok && val != nil {
		data.IsregexFf = types.StringValue(val.(string))
	} else {
		data.IsregexFf = types.StringNull()
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

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("fieldformat:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Fieldformat.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("formactionurl_ff:%s", utils.UrlEncode(fmt.Sprintf("%v", data.FormactionurlFf.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
