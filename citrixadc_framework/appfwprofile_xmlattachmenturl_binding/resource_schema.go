package appfwprofile_xmlattachmenturl_binding

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

// AppfwprofileXmlattachmenturlBindingResourceModel describes the resource data model.
type AppfwprofileXmlattachmenturlBindingResourceModel struct {
	Id                            types.String `tfsdk:"id"`
	Alertonly                     types.String `tfsdk:"alertonly"`
	Comment                       types.String `tfsdk:"comment"`
	Isautodeployed                types.String `tfsdk:"isautodeployed"`
	Name                          types.String `tfsdk:"name"`
	Resourceid                    types.String `tfsdk:"resourceid"`
	State                         types.String `tfsdk:"state"`
	Xmlattachmentcontenttype      types.String `tfsdk:"xmlattachmentcontenttype"`
	Xmlattachmentcontenttypecheck types.String `tfsdk:"xmlattachmentcontenttypecheck"`
	Xmlattachmenturl              types.String `tfsdk:"xmlattachmenturl"`
	Xmlmaxattachmentsize          types.Int64  `tfsdk:"xmlmaxattachmentsize"`
	Xmlmaxattachmentsizecheck     types.String `tfsdk:"xmlmaxattachmentsizecheck"`
}

func (r *AppfwprofileXmlattachmenturlBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwprofile_xmlattachmenturl_binding resource.",
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
			"xmlattachmentcontenttype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify content-type regular expression.",
			},
			"xmlattachmentcontenttypecheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML attachment content-type check is ON or OFF. Protects against XML requests with illegal attachments.",
			},
			"xmlattachmenturl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "XML attachment URL regular expression length.",
			},
			"xmlmaxattachmentsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify maximum attachment size.",
			},
			"xmlmaxattachmentsizecheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Max attachment size Check is ON or OFF. Protects against XML requests with large attachment data.",
			},
		},
	}
}

func appfwprofile_xmlattachmenturl_bindingGetThePayloadFromtheConfig(ctx context.Context, data *AppfwprofileXmlattachmenturlBindingResourceModel) appfw.Appfwprofilexmlattachmenturlbinding {
	tflog.Debug(ctx, "In appfwprofile_xmlattachmenturl_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appfwprofile_xmlattachmenturl_binding := appfw.Appfwprofilexmlattachmenturlbinding{}
	if !data.Alertonly.IsNull() {
		appfwprofile_xmlattachmenturl_binding.Alertonly = data.Alertonly.ValueString()
	}
	if !data.Comment.IsNull() {
		appfwprofile_xmlattachmenturl_binding.Comment = data.Comment.ValueString()
	}
	if !data.Isautodeployed.IsNull() {
		appfwprofile_xmlattachmenturl_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.Name.IsNull() {
		appfwprofile_xmlattachmenturl_binding.Name = data.Name.ValueString()
	}
	if !data.Resourceid.IsNull() {
		appfwprofile_xmlattachmenturl_binding.Resourceid = data.Resourceid.ValueString()
	}
	if !data.State.IsNull() {
		appfwprofile_xmlattachmenturl_binding.State = data.State.ValueString()
	}
	if !data.Xmlattachmentcontenttype.IsNull() {
		appfwprofile_xmlattachmenturl_binding.Xmlattachmentcontenttype = data.Xmlattachmentcontenttype.ValueString()
	}
	if !data.Xmlattachmentcontenttypecheck.IsNull() {
		appfwprofile_xmlattachmenturl_binding.Xmlattachmentcontenttypecheck = data.Xmlattachmentcontenttypecheck.ValueString()
	}
	if !data.Xmlattachmenturl.IsNull() {
		appfwprofile_xmlattachmenturl_binding.Xmlattachmenturl = data.Xmlattachmenturl.ValueString()
	}
	if !data.Xmlmaxattachmentsize.IsNull() {
		appfwprofile_xmlattachmenturl_binding.Xmlmaxattachmentsize = utils.IntPtr(int(data.Xmlmaxattachmentsize.ValueInt64()))
	}
	if !data.Xmlmaxattachmentsizecheck.IsNull() {
		appfwprofile_xmlattachmenturl_binding.Xmlmaxattachmentsizecheck = data.Xmlmaxattachmentsizecheck.ValueString()
	}

	return appfwprofile_xmlattachmenturl_binding
}

func appfwprofile_xmlattachmenturl_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileXmlattachmenturlBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileXmlattachmenturlBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_xmlattachmenturl_bindingSetAttrFromGet Function")

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
	if val, ok := getResponseData["xmlattachmentcontenttype"]; ok && val != nil {
		data.Xmlattachmentcontenttype = types.StringValue(val.(string))
	} else {
		data.Xmlattachmentcontenttype = types.StringNull()
	}
	if val, ok := getResponseData["xmlattachmentcontenttypecheck"]; ok && val != nil {
		data.Xmlattachmentcontenttypecheck = types.StringValue(val.(string))
	} else {
		data.Xmlattachmentcontenttypecheck = types.StringNull()
	}
	if val, ok := getResponseData["xmlattachmenturl"]; ok && val != nil {
		data.Xmlattachmenturl = types.StringValue(val.(string))
	} else {
		data.Xmlattachmenturl = types.StringNull()
	}
	if val, ok := getResponseData["xmlmaxattachmentsize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxattachmentsize = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxattachmentsize = types.Int64Null()
	}
	if val, ok := getResponseData["xmlmaxattachmentsizecheck"]; ok && val != nil {
		data.Xmlmaxattachmentsizecheck = types.StringValue(val.(string))
	} else {
		data.Xmlmaxattachmentsizecheck = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("xmlattachmenturl:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Xmlattachmenturl.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
