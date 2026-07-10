package appfwprofile_xmlattachmenturl_binding

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

// AppfwprofileXmlattachmenturlBindingResourceModel describes the resource data model.
type AppfwprofileXmlattachmenturlBindingResourceModel struct {
	Id                            types.String `tfsdk:"id"`
	Alertonly                     types.String `tfsdk:"alertonly"`
	Comment                       types.String `tfsdk:"comment"`
	Isautodeployed                types.String `tfsdk:"isautodeployed"`
	Name                          types.String `tfsdk:"name"`
	Resourceid                    types.String `tfsdk:"resourceid"`
	Ruletype                      types.String `tfsdk:"ruletype"`
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
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments about the purpose of profile, or other useful information about the profile.",
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
			"xmlattachmentcontenttype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Specify content-type regular expression.",
			},
			"xmlattachmentcontenttypecheck": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "State if XML attachment content-type check is ON or OFF. Protects against XML requests with illegal attachments.",
			},
			"xmlattachmenturl": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "XML attachment URL regular expression length.",
			},
			"xmlmaxattachmentsize": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Specify maximum attachment size.",
			},
			"xmlmaxattachmentsizecheck": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "State if XML Max attachment size Check is ON or OFF. Protects against XML requests with large attachment data.",
			},
		},
	}
}

func appfwprofile_xmlattachmenturl_bindingGetThePayloadFromthePlan(ctx context.Context, data *AppfwprofileXmlattachmenturlBindingResourceModel) appfw.Appfwprofilexmlattachmenturlbinding {
	tflog.Debug(ctx, "In appfwprofile_xmlattachmenturl_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	appfwprofile_xmlattachmenturl_binding := appfw.Appfwprofilexmlattachmenturlbinding{}
	if !data.Alertonly.IsNull() && !data.Alertonly.IsUnknown() {
		appfwprofile_xmlattachmenturl_binding.Alertonly = data.Alertonly.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appfwprofile_xmlattachmenturl_binding.Comment = data.Comment.ValueString()
	}
	if !data.Isautodeployed.IsNull() && !data.Isautodeployed.IsUnknown() {
		appfwprofile_xmlattachmenturl_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwprofile_xmlattachmenturl_binding.Name = data.Name.ValueString()
	}
	if !data.Resourceid.IsNull() && !data.Resourceid.IsUnknown() {
		appfwprofile_xmlattachmenturl_binding.Resourceid = data.Resourceid.ValueString()
	}
	if !data.Ruletype.IsNull() && !data.Ruletype.IsUnknown() {
		appfwprofile_xmlattachmenturl_binding.Ruletype = data.Ruletype.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		appfwprofile_xmlattachmenturl_binding.State = data.State.ValueString()
	}
	if !data.Xmlattachmentcontenttype.IsNull() && !data.Xmlattachmentcontenttype.IsUnknown() {
		appfwprofile_xmlattachmenturl_binding.Xmlattachmentcontenttype = data.Xmlattachmentcontenttype.ValueString()
	}
	if !data.Xmlattachmentcontenttypecheck.IsNull() && !data.Xmlattachmentcontenttypecheck.IsUnknown() {
		appfwprofile_xmlattachmenturl_binding.Xmlattachmentcontenttypecheck = data.Xmlattachmentcontenttypecheck.ValueString()
	}
	if !data.Xmlattachmenturl.IsNull() && !data.Xmlattachmenturl.IsUnknown() {
		appfwprofile_xmlattachmenturl_binding.Xmlattachmenturl = data.Xmlattachmenturl.ValueString()
	}
	if !data.Xmlmaxattachmentsize.IsNull() && !data.Xmlmaxattachmentsize.IsUnknown() {
		appfwprofile_xmlattachmenturl_binding.Xmlmaxattachmentsize = utils.IntPtr(int(data.Xmlmaxattachmentsize.ValueInt64()))
	}
	if !data.Xmlmaxattachmentsizecheck.IsNull() && !data.Xmlmaxattachmentsizecheck.IsUnknown() {
		appfwprofile_xmlattachmenturl_binding.Xmlmaxattachmentsizecheck = data.Xmlmaxattachmentsizecheck.ValueString()
	}

	return appfwprofile_xmlattachmenturl_binding
}

// appfwprofile_xmlattachmenturl_bindingSetAttrFromGet is the RESOURCE-side setter.
// All attributes are RequiresReplace (no update endpoint) and the NITRO server may
// echo server-defaulted/normalized values for fields like alertonly, isautodeployed,
// resourceid, ruletype. To avoid "inconsistent result after apply" we adopt the GET
// value only when the model field is currently null/unknown (e.g. import); otherwise
// we preserve the configured plan/state value. The ID is set once in Create and is
// preserved here.
func appfwprofile_xmlattachmenturl_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileXmlattachmenturlBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileXmlattachmenturlBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_xmlattachmenturl_bindingSetAttrFromGet Function")

	adopt := func(cur types.String, key string) types.String {
		if !cur.IsNull() && !cur.IsUnknown() {
			return cur
		}
		if val, ok := getResponseData[key]; ok && val != nil {
			return types.StringValue(val.(string))
		}
		return types.StringNull()
	}

	data.Alertonly = adopt(data.Alertonly, "alertonly")
	data.Comment = adopt(data.Comment, "comment")
	data.Isautodeployed = adopt(data.Isautodeployed, "isautodeployed")
	data.Name = adopt(data.Name, "name")
	data.Resourceid = adopt(data.Resourceid, "resourceid")
	data.Ruletype = adopt(data.Ruletype, "ruletype")
	data.State = adopt(data.State, "state")
	data.Xmlattachmentcontenttype = adopt(data.Xmlattachmentcontenttype, "xmlattachmentcontenttype")
	data.Xmlattachmentcontenttypecheck = adopt(data.Xmlattachmentcontenttypecheck, "xmlattachmentcontenttypecheck")
	data.Xmlattachmenturl = adopt(data.Xmlattachmenturl, "xmlattachmenturl")
	if !data.Xmlmaxattachmentsize.IsNull() && !data.Xmlmaxattachmentsize.IsUnknown() {
		// preserve configured value
	} else if val, ok := getResponseData["xmlmaxattachmentsize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxattachmentsize = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxattachmentsize = types.Int64Null()
	}
	data.Xmlmaxattachmentsizecheck = adopt(data.Xmlmaxattachmentsizecheck, "xmlmaxattachmentsizecheck")

	return data
}

// appfwprofile_xmlattachmenturl_bindingSetAttrFromGetForDatasource is the
// DATASOURCE-side setter: it faithfully copies every field from the GET response
// (the datasource has no prior plan/state to preserve) and sets the composite ID.
func appfwprofile_xmlattachmenturl_bindingSetAttrFromGetForDatasource(ctx context.Context, data *AppfwprofileXmlattachmenturlBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileXmlattachmenturlBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_xmlattachmenturl_bindingSetAttrFromGetForDatasource Function")

	copyField := func(key string) types.String {
		if val, ok := getResponseData[key]; ok && val != nil {
			return types.StringValue(val.(string))
		}
		return types.StringNull()
	}

	data.Alertonly = copyField("alertonly")
	data.Comment = copyField("comment")
	data.Isautodeployed = copyField("isautodeployed")
	data.Name = copyField("name")
	data.Resourceid = copyField("resourceid")
	data.Ruletype = copyField("ruletype")
	data.State = copyField("state")
	data.Xmlattachmentcontenttype = copyField("xmlattachmentcontenttype")
	data.Xmlattachmentcontenttypecheck = copyField("xmlattachmentcontenttypecheck")
	data.Xmlattachmenturl = copyField("xmlattachmenturl")
	if val, ok := getResponseData["xmlmaxattachmentsize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxattachmentsize = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxattachmentsize = types.Int64Null()
	}
	data.Xmlmaxattachmentsizecheck = copyField("xmlmaxattachmentsizecheck")

	// Set ID for the datasource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("xmlattachmenturl:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Xmlattachmenturl.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
