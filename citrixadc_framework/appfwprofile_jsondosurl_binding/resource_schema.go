package appfwprofile_jsondosurl_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AppfwprofileJsondosurlBindingResourceModel describes the resource data model.
type AppfwprofileJsondosurlBindingResourceModel struct {
	Id                          types.String `tfsdk:"id"`
	Alertonly                   types.String `tfsdk:"alertonly"`
	Comment                     types.String `tfsdk:"comment"`
	Isautodeployed              types.String `tfsdk:"isautodeployed"`
	Jsondosurl                  types.String `tfsdk:"jsondosurl"`
	Jsonmaxarraylength          types.Int64  `tfsdk:"jsonmaxarraylength"`
	Jsonmaxarraylengthcheck     types.String `tfsdk:"jsonmaxarraylengthcheck"`
	Jsonmaxcontainerdepth       types.Int64  `tfsdk:"jsonmaxcontainerdepth"`
	Jsonmaxcontainerdepthcheck  types.String `tfsdk:"jsonmaxcontainerdepthcheck"`
	Jsonmaxdocumentlength       types.Int64  `tfsdk:"jsonmaxdocumentlength"`
	Jsonmaxdocumentlengthcheck  types.String `tfsdk:"jsonmaxdocumentlengthcheck"`
	Jsonmaxobjectkeycount       types.Int64  `tfsdk:"jsonmaxobjectkeycount"`
	Jsonmaxobjectkeycountcheck  types.String `tfsdk:"jsonmaxobjectkeycountcheck"`
	Jsonmaxobjectkeylength      types.Int64  `tfsdk:"jsonmaxobjectkeylength"`
	Jsonmaxobjectkeylengthcheck types.String `tfsdk:"jsonmaxobjectkeylengthcheck"`
	Jsonmaxstringlength         types.Int64  `tfsdk:"jsonmaxstringlength"`
	Jsonmaxstringlengthcheck    types.String `tfsdk:"jsonmaxstringlengthcheck"`
	Name                        types.String `tfsdk:"name"`
	Resourceid                  types.String `tfsdk:"resourceid"`
	State                       types.String `tfsdk:"state"`
}

func (r *AppfwprofileJsondosurlBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwprofile_jsondosurl_binding resource.",
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
			"jsondosurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The URL on which we need to enforce the specified JSON denial-of-service (JSONDoS) attack protections.\nAn JSON DoS configuration consists of the following items:\n* URL. PCRE-format regular expression for the URL.\n* Maximum-document-length-check toggle.  ON to enable this check, OFF to disable it.\n* Maximum document length. Positive integer representing the maximum length of the JSON document.\n* Maximum-container-depth-check toggle. ON to enable, OFF to disable.\n * Maximum container depth. Positive integer representing the maximum container depth of the JSON document.\n* Maximum-object-key-count-check toggle. ON to enable, OFF to disable.\n* Maximum object key count. Positive integer representing the maximum allowed number of keys in any of the  JSON object.\n* Maximum-object-key-length-check toggle. ON to enable, OFF to disable.\n* Maximum object key length. Positive integer representing the maximum allowed length of key in any of the  JSON object.\n* Maximum-array-value-count-check toggle. ON to enable, OFF to disable.\n* Maximum array value count. Positive integer representing the maximum allowed number of values in any of the JSON array.\n* Maximum-string-length-check toggle. ON to enable, OFF to disable.\n* Maximum string length. Positive integer representing the maximum length of string in JSON.",
			},
			"jsonmaxarraylength": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(10000),
				Description: "Maximum array length in the any of JSON object. This check protects against arrays having large lengths.",
			},
			"jsonmaxarraylengthcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if JSON Max array value count check is ON or OFF.",
			},
			"jsonmaxcontainerdepth": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(5),
				Description: "Maximum allowed nesting depth  of JSON document. JSON allows one to nest the containers (object and array) in any order to any depth. This check protects against documents that have excessive depth of hierarchy.",
			},
			"jsonmaxcontainerdepthcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if JSON Max depth check is ON or OFF.",
			},
			"jsonmaxdocumentlength": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(20000000),
				Description: "Maximum document length of JSON document, in bytes.",
			},
			"jsonmaxdocumentlengthcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if JSON Max document length check is ON or OFF.",
			},
			"jsonmaxobjectkeycount": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(10000),
				Description: "Maximum key count in the any of JSON object. This check protects against objects that have large number of keys.",
			},
			"jsonmaxobjectkeycountcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if JSON Max object key count check is ON or OFF.",
			},
			"jsonmaxobjectkeylength": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(128),
				Description: "Maximum key length in the any of JSON object. This check protects against objects that have large keys.",
			},
			"jsonmaxobjectkeylengthcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if JSON Max object key length check is ON or OFF.",
			},
			"jsonmaxstringlength": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1000000),
				Description: "Maximum string length in the JSON. This check protects against strings that have large length.",
			},
			"jsonmaxstringlengthcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if JSON Max string value count check is ON or OFF.",
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

func appfwprofile_jsondosurl_bindingGetThePayloadFromtheConfig(ctx context.Context, data *AppfwprofileJsondosurlBindingResourceModel) appfw.Appfwprofilejsondosurlbinding {
	tflog.Debug(ctx, "In appfwprofile_jsondosurl_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appfwprofile_jsondosurl_binding := appfw.Appfwprofilejsondosurlbinding{}
	if !data.Alertonly.IsNull() {
		appfwprofile_jsondosurl_binding.Alertonly = data.Alertonly.ValueString()
	}
	if !data.Comment.IsNull() {
		appfwprofile_jsondosurl_binding.Comment = data.Comment.ValueString()
	}
	if !data.Isautodeployed.IsNull() {
		appfwprofile_jsondosurl_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.Jsondosurl.IsNull() {
		appfwprofile_jsondosurl_binding.Jsondosurl = data.Jsondosurl.ValueString()
	}
	if !data.Jsonmaxarraylength.IsNull() {
		appfwprofile_jsondosurl_binding.Jsonmaxarraylength = utils.IntPtr(int(data.Jsonmaxarraylength.ValueInt64()))
	}
	if !data.Jsonmaxarraylengthcheck.IsNull() {
		appfwprofile_jsondosurl_binding.Jsonmaxarraylengthcheck = data.Jsonmaxarraylengthcheck.ValueString()
	}
	if !data.Jsonmaxcontainerdepth.IsNull() {
		appfwprofile_jsondosurl_binding.Jsonmaxcontainerdepth = utils.IntPtr(int(data.Jsonmaxcontainerdepth.ValueInt64()))
	}
	if !data.Jsonmaxcontainerdepthcheck.IsNull() {
		appfwprofile_jsondosurl_binding.Jsonmaxcontainerdepthcheck = data.Jsonmaxcontainerdepthcheck.ValueString()
	}
	if !data.Jsonmaxdocumentlength.IsNull() {
		appfwprofile_jsondosurl_binding.Jsonmaxdocumentlength = utils.IntPtr(int(data.Jsonmaxdocumentlength.ValueInt64()))
	}
	if !data.Jsonmaxdocumentlengthcheck.IsNull() {
		appfwprofile_jsondosurl_binding.Jsonmaxdocumentlengthcheck = data.Jsonmaxdocumentlengthcheck.ValueString()
	}
	if !data.Jsonmaxobjectkeycount.IsNull() {
		appfwprofile_jsondosurl_binding.Jsonmaxobjectkeycount = utils.IntPtr(int(data.Jsonmaxobjectkeycount.ValueInt64()))
	}
	if !data.Jsonmaxobjectkeycountcheck.IsNull() {
		appfwprofile_jsondosurl_binding.Jsonmaxobjectkeycountcheck = data.Jsonmaxobjectkeycountcheck.ValueString()
	}
	if !data.Jsonmaxobjectkeylength.IsNull() {
		appfwprofile_jsondosurl_binding.Jsonmaxobjectkeylength = utils.IntPtr(int(data.Jsonmaxobjectkeylength.ValueInt64()))
	}
	if !data.Jsonmaxobjectkeylengthcheck.IsNull() {
		appfwprofile_jsondosurl_binding.Jsonmaxobjectkeylengthcheck = data.Jsonmaxobjectkeylengthcheck.ValueString()
	}
	if !data.Jsonmaxstringlength.IsNull() {
		appfwprofile_jsondosurl_binding.Jsonmaxstringlength = utils.IntPtr(int(data.Jsonmaxstringlength.ValueInt64()))
	}
	if !data.Jsonmaxstringlengthcheck.IsNull() {
		appfwprofile_jsondosurl_binding.Jsonmaxstringlengthcheck = data.Jsonmaxstringlengthcheck.ValueString()
	}
	if !data.Name.IsNull() {
		appfwprofile_jsondosurl_binding.Name = data.Name.ValueString()
	}
	if !data.Resourceid.IsNull() {
		appfwprofile_jsondosurl_binding.Resourceid = data.Resourceid.ValueString()
	}
	if !data.State.IsNull() {
		appfwprofile_jsondosurl_binding.State = data.State.ValueString()
	}

	return appfwprofile_jsondosurl_binding
}

func appfwprofile_jsondosurl_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileJsondosurlBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileJsondosurlBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_jsondosurl_bindingSetAttrFromGet Function")

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
	if val, ok := getResponseData["jsondosurl"]; ok && val != nil {
		data.Jsondosurl = types.StringValue(val.(string))
	} else {
		data.Jsondosurl = types.StringNull()
	}
	if val, ok := getResponseData["jsonmaxarraylength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Jsonmaxarraylength = types.Int64Value(intVal)
		}
	} else {
		data.Jsonmaxarraylength = types.Int64Null()
	}
	if val, ok := getResponseData["jsonmaxarraylengthcheck"]; ok && val != nil {
		data.Jsonmaxarraylengthcheck = types.StringValue(val.(string))
	} else {
		data.Jsonmaxarraylengthcheck = types.StringNull()
	}
	if val, ok := getResponseData["jsonmaxcontainerdepth"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Jsonmaxcontainerdepth = types.Int64Value(intVal)
		}
	} else {
		data.Jsonmaxcontainerdepth = types.Int64Null()
	}
	if val, ok := getResponseData["jsonmaxcontainerdepthcheck"]; ok && val != nil {
		data.Jsonmaxcontainerdepthcheck = types.StringValue(val.(string))
	} else {
		data.Jsonmaxcontainerdepthcheck = types.StringNull()
	}
	if val, ok := getResponseData["jsonmaxdocumentlength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Jsonmaxdocumentlength = types.Int64Value(intVal)
		}
	} else {
		data.Jsonmaxdocumentlength = types.Int64Null()
	}
	if val, ok := getResponseData["jsonmaxdocumentlengthcheck"]; ok && val != nil {
		data.Jsonmaxdocumentlengthcheck = types.StringValue(val.(string))
	} else {
		data.Jsonmaxdocumentlengthcheck = types.StringNull()
	}
	if val, ok := getResponseData["jsonmaxobjectkeycount"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Jsonmaxobjectkeycount = types.Int64Value(intVal)
		}
	} else {
		data.Jsonmaxobjectkeycount = types.Int64Null()
	}
	if val, ok := getResponseData["jsonmaxobjectkeycountcheck"]; ok && val != nil {
		data.Jsonmaxobjectkeycountcheck = types.StringValue(val.(string))
	} else {
		data.Jsonmaxobjectkeycountcheck = types.StringNull()
	}
	if val, ok := getResponseData["jsonmaxobjectkeylength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Jsonmaxobjectkeylength = types.Int64Value(intVal)
		}
	} else {
		data.Jsonmaxobjectkeylength = types.Int64Null()
	}
	if val, ok := getResponseData["jsonmaxobjectkeylengthcheck"]; ok && val != nil {
		data.Jsonmaxobjectkeylengthcheck = types.StringValue(val.(string))
	} else {
		data.Jsonmaxobjectkeylengthcheck = types.StringNull()
	}
	if val, ok := getResponseData["jsonmaxstringlength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Jsonmaxstringlength = types.Int64Value(intVal)
		}
	} else {
		data.Jsonmaxstringlength = types.Int64Null()
	}
	if val, ok := getResponseData["jsonmaxstringlengthcheck"]; ok && val != nil {
		data.Jsonmaxstringlengthcheck = types.StringValue(val.(string))
	} else {
		data.Jsonmaxstringlengthcheck = types.StringNull()
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
	idParts = append(idParts, fmt.Sprintf("jsondosurl:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Jsondosurl.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
