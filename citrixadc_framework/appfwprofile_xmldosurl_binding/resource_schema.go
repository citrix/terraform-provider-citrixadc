package appfwprofile_xmldosurl_binding

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

// AppfwprofileXmldosurlBindingResourceModel describes the resource data model.
type AppfwprofileXmldosurlBindingResourceModel struct {
	Id                              types.String `tfsdk:"id"`
	Alertonly                       types.String `tfsdk:"alertonly"`
	Comment                         types.String `tfsdk:"comment"`
	Isautodeployed                  types.String `tfsdk:"isautodeployed"`
	Name                            types.String `tfsdk:"name"`
	Resourceid                      types.String `tfsdk:"resourceid"`
	State                           types.String `tfsdk:"state"`
	Xmlblockdtd                     types.String `tfsdk:"xmlblockdtd"`
	Xmlblockexternalentities        types.String `tfsdk:"xmlblockexternalentities"`
	Xmlblockpi                      types.String `tfsdk:"xmlblockpi"`
	Xmldosurl                       types.String `tfsdk:"xmldosurl"`
	Xmlmaxattributenamelength       types.Int64  `tfsdk:"xmlmaxattributenamelength"`
	Xmlmaxattributenamelengthcheck  types.String `tfsdk:"xmlmaxattributenamelengthcheck"`
	Xmlmaxattributes                types.Int64  `tfsdk:"xmlmaxattributes"`
	Xmlmaxattributescheck           types.String `tfsdk:"xmlmaxattributescheck"`
	Xmlmaxattributevaluelength      types.Int64  `tfsdk:"xmlmaxattributevaluelength"`
	Xmlmaxattributevaluelengthcheck types.String `tfsdk:"xmlmaxattributevaluelengthcheck"`
	Xmlmaxchardatalength            types.Int64  `tfsdk:"xmlmaxchardatalength"`
	Xmlmaxchardatalengthcheck       types.String `tfsdk:"xmlmaxchardatalengthcheck"`
	Xmlmaxelementchildren           types.Int64  `tfsdk:"xmlmaxelementchildren"`
	Xmlmaxelementchildrencheck      types.String `tfsdk:"xmlmaxelementchildrencheck"`
	Xmlmaxelementdepth              types.Int64  `tfsdk:"xmlmaxelementdepth"`
	Xmlmaxelementdepthcheck         types.String `tfsdk:"xmlmaxelementdepthcheck"`
	Xmlmaxelementnamelength         types.Int64  `tfsdk:"xmlmaxelementnamelength"`
	Xmlmaxelementnamelengthcheck    types.String `tfsdk:"xmlmaxelementnamelengthcheck"`
	Xmlmaxelements                  types.Int64  `tfsdk:"xmlmaxelements"`
	Xmlmaxelementscheck             types.String `tfsdk:"xmlmaxelementscheck"`
	Xmlmaxentityexpansiondepth      types.Int64  `tfsdk:"xmlmaxentityexpansiondepth"`
	Xmlmaxentityexpansiondepthcheck types.String `tfsdk:"xmlmaxentityexpansiondepthcheck"`
	Xmlmaxentityexpansions          types.Int64  `tfsdk:"xmlmaxentityexpansions"`
	Xmlmaxentityexpansionscheck     types.String `tfsdk:"xmlmaxentityexpansionscheck"`
	Xmlmaxfilesize                  types.Int64  `tfsdk:"xmlmaxfilesize"`
	Xmlmaxfilesizecheck             types.String `tfsdk:"xmlmaxfilesizecheck"`
	Xmlmaxnamespaces                types.Int64  `tfsdk:"xmlmaxnamespaces"`
	Xmlmaxnamespacescheck           types.String `tfsdk:"xmlmaxnamespacescheck"`
	Xmlmaxnamespaceurilength        types.Int64  `tfsdk:"xmlmaxnamespaceurilength"`
	Xmlmaxnamespaceurilengthcheck   types.String `tfsdk:"xmlmaxnamespaceurilengthcheck"`
	Xmlmaxnodes                     types.Int64  `tfsdk:"xmlmaxnodes"`
	Xmlmaxnodescheck                types.String `tfsdk:"xmlmaxnodescheck"`
	Xmlmaxsoaparrayrank             types.Int64  `tfsdk:"xmlmaxsoaparrayrank"`
	Xmlmaxsoaparraysize             types.Int64  `tfsdk:"xmlmaxsoaparraysize"`
	Xmlminfilesize                  types.Int64  `tfsdk:"xmlminfilesize"`
	Xmlminfilesizecheck             types.String `tfsdk:"xmlminfilesizecheck"`
	Xmlsoaparraycheck               types.String `tfsdk:"xmlsoaparraycheck"`
}

func (r *AppfwprofileXmldosurlBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwprofile_xmldosurl_binding resource.",
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
			"xmlblockdtd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML DTD is ON or OFF. Protects against recursive Document Type Declaration (DTD) entity expansion attacks. Also, SOAP messages cannot have DTDs in messages.",
			},
			"xmlblockexternalentities": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Block External Entities Check is ON or OFF. Protects against XML External Entity (XXE) attacks that force applications to parse untrusted external entities (sources) in XML documents.",
			},
			"xmlblockpi": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Block PI is ON or OFF. Protects resources from denial of service attacks as SOAP messages cannot have processing instructions (PI) in messages.",
			},
			"xmldosurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "XML DoS URL regular expression length.",
			},
			"xmlmaxattributenamelength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the longest name of any XML attribute. Protects against overflow attacks.",
			},
			"xmlmaxattributenamelengthcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Max attribute name length check is ON or OFF.",
			},
			"xmlmaxattributes": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify maximum number of attributes per XML element. Protects against overflow attacks.",
			},
			"xmlmaxattributescheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Max attributes check is ON or OFF.",
			},
			"xmlmaxattributevaluelength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the longest value of any XML attribute. Protects against overflow attacks.",
			},
			"xmlmaxattributevaluelengthcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Max atribute value length is ON or OFF.",
			},
			"xmlmaxchardatalength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the maximum size of CDATA. Protects against overflow attacks and large quantities of unparsed data within XML messages.",
			},
			"xmlmaxchardatalengthcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Max CDATA length check is ON or OFF.",
			},
			"xmlmaxelementchildren": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the maximum number of children allowed per XML element. Protects against overflow attacks.",
			},
			"xmlmaxelementchildrencheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Max element children check is ON or OFF.",
			},
			"xmlmaxelementdepth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum nesting (depth) of XML elements. This check protects against documents that have excessive hierarchy depths.",
			},
			"xmlmaxelementdepthcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Max element depth check is ON or OFF.",
			},
			"xmlmaxelementnamelength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the longest name of any element (including the expanded namespace) to protect against overflow attacks.",
			},
			"xmlmaxelementnamelengthcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Max element name length check is ON or OFF.",
			},
			"xmlmaxelements": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the maximum number of XML elements allowed. Protects against overflow attacks.",
			},
			"xmlmaxelementscheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Max elements check is ON or OFF.",
			},
			"xmlmaxentityexpansiondepth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify maximum entity expansion depth. Protects aganist Entity Expansion Attack.",
			},
			"xmlmaxentityexpansiondepthcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Max Entity Expansions Depth Check is ON or OFF.",
			},
			"xmlmaxentityexpansions": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify maximum allowed number of entity expansions. Protects aganist Entity Expansion Attack.",
			},
			"xmlmaxentityexpansionscheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Max Entity Expansions Check is ON or OFF.",
			},
			"xmlmaxfilesize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the maximum size of XML messages. Protects against overflow attacks.",
			},
			"xmlmaxfilesizecheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Max file size check is ON or OFF.",
			},
			"xmlmaxnamespaces": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify maximum number of active namespaces. Protects against overflow attacks.",
			},
			"xmlmaxnamespacescheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Max namespaces check is ON or OFF.",
			},
			"xmlmaxnamespaceurilength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the longest URI of any XML namespace. Protects against overflow attacks.",
			},
			"xmlmaxnamespaceurilengthcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Max namespace URI length check is ON or OFF.",
			},
			"xmlmaxnodes": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the maximum number of XML nodes. Protects against overflow attacks.",
			},
			"xmlmaxnodescheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Max nodes check is ON or OFF.",
			},
			"xmlmaxsoaparrayrank": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "XML Max Individual SOAP Array Rank. This is the dimension of the SOAP array.",
			},
			"xmlmaxsoaparraysize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "XML Max Total SOAP Array Size. Protects against SOAP Array Abuse attack.",
			},
			"xmlminfilesize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Enforces minimum message size.",
			},
			"xmlminfilesizecheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Min file size check is ON or OFF.",
			},
			"xmlsoaparraycheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML SOAP Array check is ON or OFF.",
			},
		},
	}
}

func appfwprofile_xmldosurl_bindingGetThePayloadFromtheConfig(ctx context.Context, data *AppfwprofileXmldosurlBindingResourceModel) appfw.Appfwprofilexmldosurlbinding {
	tflog.Debug(ctx, "In appfwprofile_xmldosurl_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appfwprofile_xmldosurl_binding := appfw.Appfwprofilexmldosurlbinding{}
	if !data.Alertonly.IsNull() {
		appfwprofile_xmldosurl_binding.Alertonly = data.Alertonly.ValueString()
	}
	if !data.Comment.IsNull() {
		appfwprofile_xmldosurl_binding.Comment = data.Comment.ValueString()
	}
	if !data.Isautodeployed.IsNull() {
		appfwprofile_xmldosurl_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.Name.IsNull() {
		appfwprofile_xmldosurl_binding.Name = data.Name.ValueString()
	}
	if !data.Resourceid.IsNull() {
		appfwprofile_xmldosurl_binding.Resourceid = data.Resourceid.ValueString()
	}
	if !data.State.IsNull() {
		appfwprofile_xmldosurl_binding.State = data.State.ValueString()
	}
	if !data.Xmlblockdtd.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlblockdtd = data.Xmlblockdtd.ValueString()
	}
	if !data.Xmlblockexternalentities.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlblockexternalentities = data.Xmlblockexternalentities.ValueString()
	}
	if !data.Xmlblockpi.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlblockpi = data.Xmlblockpi.ValueString()
	}
	if !data.Xmldosurl.IsNull() {
		appfwprofile_xmldosurl_binding.Xmldosurl = data.Xmldosurl.ValueString()
	}
	if !data.Xmlmaxattributenamelength.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlmaxattributenamelength = utils.IntPtr(int(data.Xmlmaxattributenamelength.ValueInt64()))
	}
	if !data.Xmlmaxattributenamelengthcheck.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlmaxattributenamelengthcheck = data.Xmlmaxattributenamelengthcheck.ValueString()
	}
	if !data.Xmlmaxattributes.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlmaxattributes = utils.IntPtr(int(data.Xmlmaxattributes.ValueInt64()))
	}
	if !data.Xmlmaxattributescheck.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlmaxattributescheck = data.Xmlmaxattributescheck.ValueString()
	}
	if !data.Xmlmaxattributevaluelength.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlmaxattributevaluelength = utils.IntPtr(int(data.Xmlmaxattributevaluelength.ValueInt64()))
	}
	if !data.Xmlmaxattributevaluelengthcheck.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlmaxattributevaluelengthcheck = data.Xmlmaxattributevaluelengthcheck.ValueString()
	}
	if !data.Xmlmaxchardatalength.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlmaxchardatalength = utils.IntPtr(int(data.Xmlmaxchardatalength.ValueInt64()))
	}
	if !data.Xmlmaxchardatalengthcheck.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlmaxchardatalengthcheck = data.Xmlmaxchardatalengthcheck.ValueString()
	}
	if !data.Xmlmaxelementchildren.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlmaxelementchildren = utils.IntPtr(int(data.Xmlmaxelementchildren.ValueInt64()))
	}
	if !data.Xmlmaxelementchildrencheck.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlmaxelementchildrencheck = data.Xmlmaxelementchildrencheck.ValueString()
	}
	if !data.Xmlmaxelementdepth.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlmaxelementdepth = utils.IntPtr(int(data.Xmlmaxelementdepth.ValueInt64()))
	}
	if !data.Xmlmaxelementdepthcheck.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlmaxelementdepthcheck = data.Xmlmaxelementdepthcheck.ValueString()
	}
	if !data.Xmlmaxelementnamelength.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlmaxelementnamelength = utils.IntPtr(int(data.Xmlmaxelementnamelength.ValueInt64()))
	}
	if !data.Xmlmaxelementnamelengthcheck.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlmaxelementnamelengthcheck = data.Xmlmaxelementnamelengthcheck.ValueString()
	}
	if !data.Xmlmaxelements.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlmaxelements = utils.IntPtr(int(data.Xmlmaxelements.ValueInt64()))
	}
	if !data.Xmlmaxelementscheck.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlmaxelementscheck = data.Xmlmaxelementscheck.ValueString()
	}
	if !data.Xmlmaxentityexpansiondepth.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlmaxentityexpansiondepth = utils.IntPtr(int(data.Xmlmaxentityexpansiondepth.ValueInt64()))
	}
	if !data.Xmlmaxentityexpansiondepthcheck.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlmaxentityexpansiondepthcheck = data.Xmlmaxentityexpansiondepthcheck.ValueString()
	}
	if !data.Xmlmaxentityexpansions.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlmaxentityexpansions = utils.IntPtr(int(data.Xmlmaxentityexpansions.ValueInt64()))
	}
	if !data.Xmlmaxentityexpansionscheck.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlmaxentityexpansionscheck = data.Xmlmaxentityexpansionscheck.ValueString()
	}
	if !data.Xmlmaxfilesize.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlmaxfilesize = utils.IntPtr(int(data.Xmlmaxfilesize.ValueInt64()))
	}
	if !data.Xmlmaxfilesizecheck.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlmaxfilesizecheck = data.Xmlmaxfilesizecheck.ValueString()
	}
	if !data.Xmlmaxnamespaces.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlmaxnamespaces = utils.IntPtr(int(data.Xmlmaxnamespaces.ValueInt64()))
	}
	if !data.Xmlmaxnamespacescheck.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlmaxnamespacescheck = data.Xmlmaxnamespacescheck.ValueString()
	}
	if !data.Xmlmaxnamespaceurilength.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlmaxnamespaceurilength = utils.IntPtr(int(data.Xmlmaxnamespaceurilength.ValueInt64()))
	}
	if !data.Xmlmaxnamespaceurilengthcheck.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlmaxnamespaceurilengthcheck = data.Xmlmaxnamespaceurilengthcheck.ValueString()
	}
	if !data.Xmlmaxnodes.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlmaxnodes = utils.IntPtr(int(data.Xmlmaxnodes.ValueInt64()))
	}
	if !data.Xmlmaxnodescheck.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlmaxnodescheck = data.Xmlmaxnodescheck.ValueString()
	}
	if !data.Xmlmaxsoaparrayrank.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlmaxsoaparrayrank = utils.IntPtr(int(data.Xmlmaxsoaparrayrank.ValueInt64()))
	}
	if !data.Xmlmaxsoaparraysize.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlmaxsoaparraysize = utils.IntPtr(int(data.Xmlmaxsoaparraysize.ValueInt64()))
	}
	if !data.Xmlminfilesize.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlminfilesize = utils.IntPtr(int(data.Xmlminfilesize.ValueInt64()))
	}
	if !data.Xmlminfilesizecheck.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlminfilesizecheck = data.Xmlminfilesizecheck.ValueString()
	}
	if !data.Xmlsoaparraycheck.IsNull() {
		appfwprofile_xmldosurl_binding.Xmlsoaparraycheck = data.Xmlsoaparraycheck.ValueString()
	}

	return appfwprofile_xmldosurl_binding
}

func appfwprofile_xmldosurl_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileXmldosurlBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileXmldosurlBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_xmldosurl_bindingSetAttrFromGet Function")

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
	if val, ok := getResponseData["xmlblockdtd"]; ok && val != nil {
		data.Xmlblockdtd = types.StringValue(val.(string))
	} else {
		data.Xmlblockdtd = types.StringNull()
	}
	if val, ok := getResponseData["xmlblockexternalentities"]; ok && val != nil {
		data.Xmlblockexternalentities = types.StringValue(val.(string))
	} else {
		data.Xmlblockexternalentities = types.StringNull()
	}
	if val, ok := getResponseData["xmlblockpi"]; ok && val != nil {
		data.Xmlblockpi = types.StringValue(val.(string))
	} else {
		data.Xmlblockpi = types.StringNull()
	}
	if val, ok := getResponseData["xmldosurl"]; ok && val != nil {
		data.Xmldosurl = types.StringValue(val.(string))
	} else {
		data.Xmldosurl = types.StringNull()
	}
	if val, ok := getResponseData["xmlmaxattributenamelength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxattributenamelength = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxattributenamelength = types.Int64Null()
	}
	if val, ok := getResponseData["xmlmaxattributenamelengthcheck"]; ok && val != nil {
		data.Xmlmaxattributenamelengthcheck = types.StringValue(val.(string))
	} else {
		data.Xmlmaxattributenamelengthcheck = types.StringNull()
	}
	if val, ok := getResponseData["xmlmaxattributes"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxattributes = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxattributes = types.Int64Null()
	}
	if val, ok := getResponseData["xmlmaxattributescheck"]; ok && val != nil {
		data.Xmlmaxattributescheck = types.StringValue(val.(string))
	} else {
		data.Xmlmaxattributescheck = types.StringNull()
	}
	if val, ok := getResponseData["xmlmaxattributevaluelength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxattributevaluelength = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxattributevaluelength = types.Int64Null()
	}
	if val, ok := getResponseData["xmlmaxattributevaluelengthcheck"]; ok && val != nil {
		data.Xmlmaxattributevaluelengthcheck = types.StringValue(val.(string))
	} else {
		data.Xmlmaxattributevaluelengthcheck = types.StringNull()
	}
	if val, ok := getResponseData["xmlmaxchardatalength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxchardatalength = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxchardatalength = types.Int64Null()
	}
	if val, ok := getResponseData["xmlmaxchardatalengthcheck"]; ok && val != nil {
		data.Xmlmaxchardatalengthcheck = types.StringValue(val.(string))
	} else {
		data.Xmlmaxchardatalengthcheck = types.StringNull()
	}
	if val, ok := getResponseData["xmlmaxelementchildren"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxelementchildren = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxelementchildren = types.Int64Null()
	}
	if val, ok := getResponseData["xmlmaxelementchildrencheck"]; ok && val != nil {
		data.Xmlmaxelementchildrencheck = types.StringValue(val.(string))
	} else {
		data.Xmlmaxelementchildrencheck = types.StringNull()
	}
	if val, ok := getResponseData["xmlmaxelementdepth"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxelementdepth = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxelementdepth = types.Int64Null()
	}
	if val, ok := getResponseData["xmlmaxelementdepthcheck"]; ok && val != nil {
		data.Xmlmaxelementdepthcheck = types.StringValue(val.(string))
	} else {
		data.Xmlmaxelementdepthcheck = types.StringNull()
	}
	if val, ok := getResponseData["xmlmaxelementnamelength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxelementnamelength = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxelementnamelength = types.Int64Null()
	}
	if val, ok := getResponseData["xmlmaxelementnamelengthcheck"]; ok && val != nil {
		data.Xmlmaxelementnamelengthcheck = types.StringValue(val.(string))
	} else {
		data.Xmlmaxelementnamelengthcheck = types.StringNull()
	}
	if val, ok := getResponseData["xmlmaxelements"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxelements = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxelements = types.Int64Null()
	}
	if val, ok := getResponseData["xmlmaxelementscheck"]; ok && val != nil {
		data.Xmlmaxelementscheck = types.StringValue(val.(string))
	} else {
		data.Xmlmaxelementscheck = types.StringNull()
	}
	if val, ok := getResponseData["xmlmaxentityexpansiondepth"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxentityexpansiondepth = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxentityexpansiondepth = types.Int64Null()
	}
	if val, ok := getResponseData["xmlmaxentityexpansiondepthcheck"]; ok && val != nil {
		data.Xmlmaxentityexpansiondepthcheck = types.StringValue(val.(string))
	} else {
		data.Xmlmaxentityexpansiondepthcheck = types.StringNull()
	}
	if val, ok := getResponseData["xmlmaxentityexpansions"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxentityexpansions = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxentityexpansions = types.Int64Null()
	}
	if val, ok := getResponseData["xmlmaxentityexpansionscheck"]; ok && val != nil {
		data.Xmlmaxentityexpansionscheck = types.StringValue(val.(string))
	} else {
		data.Xmlmaxentityexpansionscheck = types.StringNull()
	}
	if val, ok := getResponseData["xmlmaxfilesize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxfilesize = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxfilesize = types.Int64Null()
	}
	if val, ok := getResponseData["xmlmaxfilesizecheck"]; ok && val != nil {
		data.Xmlmaxfilesizecheck = types.StringValue(val.(string))
	} else {
		data.Xmlmaxfilesizecheck = types.StringNull()
	}
	if val, ok := getResponseData["xmlmaxnamespaces"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxnamespaces = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxnamespaces = types.Int64Null()
	}
	if val, ok := getResponseData["xmlmaxnamespacescheck"]; ok && val != nil {
		data.Xmlmaxnamespacescheck = types.StringValue(val.(string))
	} else {
		data.Xmlmaxnamespacescheck = types.StringNull()
	}
	if val, ok := getResponseData["xmlmaxnamespaceurilength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxnamespaceurilength = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxnamespaceurilength = types.Int64Null()
	}
	if val, ok := getResponseData["xmlmaxnamespaceurilengthcheck"]; ok && val != nil {
		data.Xmlmaxnamespaceurilengthcheck = types.StringValue(val.(string))
	} else {
		data.Xmlmaxnamespaceurilengthcheck = types.StringNull()
	}
	if val, ok := getResponseData["xmlmaxnodes"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxnodes = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxnodes = types.Int64Null()
	}
	if val, ok := getResponseData["xmlmaxnodescheck"]; ok && val != nil {
		data.Xmlmaxnodescheck = types.StringValue(val.(string))
	} else {
		data.Xmlmaxnodescheck = types.StringNull()
	}
	if val, ok := getResponseData["xmlmaxsoaparrayrank"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxsoaparrayrank = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxsoaparrayrank = types.Int64Null()
	}
	if val, ok := getResponseData["xmlmaxsoaparraysize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxsoaparraysize = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxsoaparraysize = types.Int64Null()
	}
	if val, ok := getResponseData["xmlminfilesize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlminfilesize = types.Int64Value(intVal)
		}
	} else {
		data.Xmlminfilesize = types.Int64Null()
	}
	if val, ok := getResponseData["xmlminfilesizecheck"]; ok && val != nil {
		data.Xmlminfilesizecheck = types.StringValue(val.(string))
	} else {
		data.Xmlminfilesizecheck = types.StringNull()
	}
	if val, ok := getResponseData["xmlsoaparraycheck"]; ok && val != nil {
		data.Xmlsoaparraycheck = types.StringValue(val.(string))
	} else {
		data.Xmlsoaparraycheck = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("xmldosurl:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Xmldosurl.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
