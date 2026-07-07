package appfwprofile_xmldosurl_binding

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

// AppfwprofileXmldosurlBindingResourceModel describes the resource data model.
type AppfwprofileXmldosurlBindingResourceModel struct {
	Id                              types.String `tfsdk:"id"`
	Alertonly                       types.String `tfsdk:"alertonly"`
	Comment                         types.String `tfsdk:"comment"`
	Isautodeployed                  types.String `tfsdk:"isautodeployed"`
	Name                            types.String `tfsdk:"name"`
	Resourceid                      types.String `tfsdk:"resourceid"`
	Ruletype                        types.String `tfsdk:"ruletype"`
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
			"xmlblockdtd": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "State if XML DTD is ON or OFF. Protects against recursive Document Type Declaration (DTD) entity expansion attacks. Also, SOAP messages cannot have DTDs in messages.",
			},
			"xmlblockexternalentities": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "State if XML Block External Entities Check is ON or OFF. Protects against XML External Entity (XXE) attacks that force applications to parse untrusted external entities (sources) in XML documents.",
			},
			"xmlblockpi": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "State if XML Block PI is ON or OFF. Protects resources from denial of service attacks as SOAP messages cannot have processing instructions (PI) in messages.",
			},
			"xmldosurl": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "XML DoS URL regular expression length.",
			},
			"xmlmaxattributenamelength": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Specify the longest name of any XML attribute. Protects against overflow attacks.",
			},
			"xmlmaxattributenamelengthcheck": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "State if XML Max attribute name length check is ON or OFF.",
			},
			"xmlmaxattributes": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Specify maximum number of attributes per XML element. Protects against overflow attacks.",
			},
			"xmlmaxattributescheck": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "State if XML Max attributes check is ON or OFF.",
			},
			"xmlmaxattributevaluelength": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Specify the longest value of any XML attribute. Protects against overflow attacks.",
			},
			"xmlmaxattributevaluelengthcheck": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "State if XML Max atribute value length is ON or OFF.",
			},
			"xmlmaxchardatalength": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Specify the maximum size of CDATA. Protects against overflow attacks and large quantities of unparsed data within XML messages.",
			},
			"xmlmaxchardatalengthcheck": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "State if XML Max CDATA length check is ON or OFF.",
			},
			"xmlmaxelementchildren": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Specify the maximum number of children allowed per XML element. Protects against overflow attacks.",
			},
			"xmlmaxelementchildrencheck": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "State if XML Max element children check is ON or OFF.",
			},
			"xmlmaxelementdepth": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Maximum nesting (depth) of XML elements. This check protects against documents that have excessive hierarchy depths.",
			},
			"xmlmaxelementdepthcheck": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "State if XML Max element depth check is ON or OFF.",
			},
			"xmlmaxelementnamelength": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Specify the longest name of any element (including the expanded namespace) to protect against overflow attacks.",
			},
			"xmlmaxelementnamelengthcheck": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "State if XML Max element name length check is ON or OFF.",
			},
			"xmlmaxelements": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Specify the maximum number of XML elements allowed. Protects against overflow attacks.",
			},
			"xmlmaxelementscheck": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "State if XML Max elements check is ON or OFF.",
			},
			"xmlmaxentityexpansiondepth": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Specify maximum entity expansion depth. Protects aganist Entity Expansion Attack.",
			},
			"xmlmaxentityexpansiondepthcheck": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "State if XML Max Entity Expansions Depth Check is ON or OFF.",
			},
			"xmlmaxentityexpansions": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Specify maximum allowed number of entity expansions. Protects aganist Entity Expansion Attack.",
			},
			"xmlmaxentityexpansionscheck": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "State if XML Max Entity Expansions Check is ON or OFF.",
			},
			"xmlmaxfilesize": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Specify the maximum size of XML messages. Protects against overflow attacks.",
			},
			"xmlmaxfilesizecheck": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "State if XML Max file size check is ON or OFF.",
			},
			"xmlmaxnamespaces": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Specify maximum number of active namespaces. Protects against overflow attacks.",
			},
			"xmlmaxnamespacescheck": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "State if XML Max namespaces check is ON or OFF.",
			},
			"xmlmaxnamespaceurilength": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Specify the longest URI of any XML namespace. Protects against overflow attacks.",
			},
			"xmlmaxnamespaceurilengthcheck": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "State if XML Max namespace URI length check is ON or OFF.",
			},
			"xmlmaxnodes": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Specify the maximum number of XML nodes. Protects against overflow attacks.",
			},
			"xmlmaxnodescheck": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "State if XML Max nodes check is ON or OFF.",
			},
			"xmlmaxsoaparrayrank": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "XML Max Individual SOAP Array Rank. This is the dimension of the SOAP array.",
			},
			"xmlmaxsoaparraysize": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "XML Max Total SOAP Array Size. Protects against SOAP Array Abuse attack.",
			},
			"xmlminfilesize": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Enforces minimum message size.",
			},
			"xmlminfilesizecheck": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "State if XML Min file size check is ON or OFF.",
			},
			"xmlsoaparraycheck": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "State if XML SOAP Array check is ON or OFF.",
			},
		},
	}
}

func appfwprofile_xmldosurl_bindingGetThePayloadFromthePlan(ctx context.Context, data *AppfwprofileXmldosurlBindingResourceModel) appfw.Appfwprofilexmldosurlbinding {
	tflog.Debug(ctx, "In appfwprofile_xmldosurl_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	appfwprofile_xmldosurl_binding := appfw.Appfwprofilexmldosurlbinding{}
	if !data.Alertonly.IsNull() && !data.Alertonly.IsUnknown() {
		appfwprofile_xmldosurl_binding.Alertonly = data.Alertonly.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appfwprofile_xmldosurl_binding.Comment = data.Comment.ValueString()
	}
	if !data.Isautodeployed.IsNull() && !data.Isautodeployed.IsUnknown() {
		appfwprofile_xmldosurl_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwprofile_xmldosurl_binding.Name = data.Name.ValueString()
	}
	if !data.Resourceid.IsNull() && !data.Resourceid.IsUnknown() {
		appfwprofile_xmldosurl_binding.Resourceid = data.Resourceid.ValueString()
	}
	if !data.Ruletype.IsNull() && !data.Ruletype.IsUnknown() {
		appfwprofile_xmldosurl_binding.Ruletype = data.Ruletype.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		appfwprofile_xmldosurl_binding.State = data.State.ValueString()
	}
	if !data.Xmlblockdtd.IsNull() && !data.Xmlblockdtd.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlblockdtd = data.Xmlblockdtd.ValueString()
	}
	if !data.Xmlblockexternalentities.IsNull() && !data.Xmlblockexternalentities.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlblockexternalentities = data.Xmlblockexternalentities.ValueString()
	}
	if !data.Xmlblockpi.IsNull() && !data.Xmlblockpi.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlblockpi = data.Xmlblockpi.ValueString()
	}
	if !data.Xmldosurl.IsNull() && !data.Xmldosurl.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmldosurl = data.Xmldosurl.ValueString()
	}
	if !data.Xmlmaxattributenamelength.IsNull() && !data.Xmlmaxattributenamelength.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlmaxattributenamelength = utils.IntPtr(int(data.Xmlmaxattributenamelength.ValueInt64()))
	}
	if !data.Xmlmaxattributenamelengthcheck.IsNull() && !data.Xmlmaxattributenamelengthcheck.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlmaxattributenamelengthcheck = data.Xmlmaxattributenamelengthcheck.ValueString()
	}
	if !data.Xmlmaxattributes.IsNull() && !data.Xmlmaxattributes.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlmaxattributes = utils.IntPtr(int(data.Xmlmaxattributes.ValueInt64()))
	}
	if !data.Xmlmaxattributescheck.IsNull() && !data.Xmlmaxattributescheck.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlmaxattributescheck = data.Xmlmaxattributescheck.ValueString()
	}
	if !data.Xmlmaxattributevaluelength.IsNull() && !data.Xmlmaxattributevaluelength.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlmaxattributevaluelength = utils.IntPtr(int(data.Xmlmaxattributevaluelength.ValueInt64()))
	}
	if !data.Xmlmaxattributevaluelengthcheck.IsNull() && !data.Xmlmaxattributevaluelengthcheck.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlmaxattributevaluelengthcheck = data.Xmlmaxattributevaluelengthcheck.ValueString()
	}
	if !data.Xmlmaxchardatalength.IsNull() && !data.Xmlmaxchardatalength.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlmaxchardatalength = utils.IntPtr(int(data.Xmlmaxchardatalength.ValueInt64()))
	}
	if !data.Xmlmaxchardatalengthcheck.IsNull() && !data.Xmlmaxchardatalengthcheck.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlmaxchardatalengthcheck = data.Xmlmaxchardatalengthcheck.ValueString()
	}
	if !data.Xmlmaxelementchildren.IsNull() && !data.Xmlmaxelementchildren.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlmaxelementchildren = utils.IntPtr(int(data.Xmlmaxelementchildren.ValueInt64()))
	}
	if !data.Xmlmaxelementchildrencheck.IsNull() && !data.Xmlmaxelementchildrencheck.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlmaxelementchildrencheck = data.Xmlmaxelementchildrencheck.ValueString()
	}
	if !data.Xmlmaxelementdepth.IsNull() && !data.Xmlmaxelementdepth.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlmaxelementdepth = utils.IntPtr(int(data.Xmlmaxelementdepth.ValueInt64()))
	}
	if !data.Xmlmaxelementdepthcheck.IsNull() && !data.Xmlmaxelementdepthcheck.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlmaxelementdepthcheck = data.Xmlmaxelementdepthcheck.ValueString()
	}
	if !data.Xmlmaxelementnamelength.IsNull() && !data.Xmlmaxelementnamelength.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlmaxelementnamelength = utils.IntPtr(int(data.Xmlmaxelementnamelength.ValueInt64()))
	}
	if !data.Xmlmaxelementnamelengthcheck.IsNull() && !data.Xmlmaxelementnamelengthcheck.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlmaxelementnamelengthcheck = data.Xmlmaxelementnamelengthcheck.ValueString()
	}
	if !data.Xmlmaxelements.IsNull() && !data.Xmlmaxelements.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlmaxelements = utils.IntPtr(int(data.Xmlmaxelements.ValueInt64()))
	}
	if !data.Xmlmaxelementscheck.IsNull() && !data.Xmlmaxelementscheck.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlmaxelementscheck = data.Xmlmaxelementscheck.ValueString()
	}
	if !data.Xmlmaxentityexpansiondepth.IsNull() && !data.Xmlmaxentityexpansiondepth.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlmaxentityexpansiondepth = utils.IntPtr(int(data.Xmlmaxentityexpansiondepth.ValueInt64()))
	}
	if !data.Xmlmaxentityexpansiondepthcheck.IsNull() && !data.Xmlmaxentityexpansiondepthcheck.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlmaxentityexpansiondepthcheck = data.Xmlmaxentityexpansiondepthcheck.ValueString()
	}
	if !data.Xmlmaxentityexpansions.IsNull() && !data.Xmlmaxentityexpansions.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlmaxentityexpansions = utils.IntPtr(int(data.Xmlmaxentityexpansions.ValueInt64()))
	}
	if !data.Xmlmaxentityexpansionscheck.IsNull() && !data.Xmlmaxentityexpansionscheck.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlmaxentityexpansionscheck = data.Xmlmaxentityexpansionscheck.ValueString()
	}
	if !data.Xmlmaxfilesize.IsNull() && !data.Xmlmaxfilesize.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlmaxfilesize = utils.IntPtr(int(data.Xmlmaxfilesize.ValueInt64()))
	}
	if !data.Xmlmaxfilesizecheck.IsNull() && !data.Xmlmaxfilesizecheck.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlmaxfilesizecheck = data.Xmlmaxfilesizecheck.ValueString()
	}
	if !data.Xmlmaxnamespaces.IsNull() && !data.Xmlmaxnamespaces.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlmaxnamespaces = utils.IntPtr(int(data.Xmlmaxnamespaces.ValueInt64()))
	}
	if !data.Xmlmaxnamespacescheck.IsNull() && !data.Xmlmaxnamespacescheck.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlmaxnamespacescheck = data.Xmlmaxnamespacescheck.ValueString()
	}
	if !data.Xmlmaxnamespaceurilength.IsNull() && !data.Xmlmaxnamespaceurilength.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlmaxnamespaceurilength = utils.IntPtr(int(data.Xmlmaxnamespaceurilength.ValueInt64()))
	}
	if !data.Xmlmaxnamespaceurilengthcheck.IsNull() && !data.Xmlmaxnamespaceurilengthcheck.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlmaxnamespaceurilengthcheck = data.Xmlmaxnamespaceurilengthcheck.ValueString()
	}
	if !data.Xmlmaxnodes.IsNull() && !data.Xmlmaxnodes.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlmaxnodes = utils.IntPtr(int(data.Xmlmaxnodes.ValueInt64()))
	}
	if !data.Xmlmaxnodescheck.IsNull() && !data.Xmlmaxnodescheck.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlmaxnodescheck = data.Xmlmaxnodescheck.ValueString()
	}
	if !data.Xmlmaxsoaparrayrank.IsNull() && !data.Xmlmaxsoaparrayrank.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlmaxsoaparrayrank = utils.IntPtr(int(data.Xmlmaxsoaparrayrank.ValueInt64()))
	}
	if !data.Xmlmaxsoaparraysize.IsNull() && !data.Xmlmaxsoaparraysize.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlmaxsoaparraysize = utils.IntPtr(int(data.Xmlmaxsoaparraysize.ValueInt64()))
	}
	if !data.Xmlminfilesize.IsNull() && !data.Xmlminfilesize.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlminfilesize = utils.IntPtr(int(data.Xmlminfilesize.ValueInt64()))
	}
	if !data.Xmlminfilesizecheck.IsNull() && !data.Xmlminfilesizecheck.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlminfilesizecheck = data.Xmlminfilesizecheck.ValueString()
	}
	if !data.Xmlsoaparraycheck.IsNull() && !data.Xmlsoaparraycheck.IsUnknown() {
		appfwprofile_xmldosurl_binding.Xmlsoaparraycheck = data.Xmlsoaparraycheck.ValueString()
	}

	return appfwprofile_xmldosurl_binding
}

// appfwprofile_xmldosurl_bindingSetAttrFromGet is the RESOURCE-side setter.
// All attributes are RequiresReplace (no update endpoint) and the NITRO server may
// echo server-defaulted/normalized values for fields like alertonly, isautodeployed,
// resourceid, ruletype. To avoid "inconsistent result after apply" we adopt the GET
// value only when the model field is currently null/unknown (e.g. import); otherwise
// we preserve the configured plan/state value. The ID is set once in Create and is
// preserved here.
func appfwprofile_xmldosurl_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileXmldosurlBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileXmldosurlBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_xmldosurl_bindingSetAttrFromGet Function")

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
	data.Xmlblockdtd = adopt(data.Xmlblockdtd, "xmlblockdtd")
	data.Xmlblockexternalentities = adopt(data.Xmlblockexternalentities, "xmlblockexternalentities")
	data.Xmlblockpi = adopt(data.Xmlblockpi, "xmlblockpi")
	data.Xmldosurl = adopt(data.Xmldosurl, "xmldosurl")
	if !data.Xmlmaxattributenamelength.IsNull() && !data.Xmlmaxattributenamelength.IsUnknown() {
		// preserve configured value
	} else if val, ok := getResponseData["xmlmaxattributenamelength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxattributenamelength = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxattributenamelength = types.Int64Null()
	}
	data.Xmlmaxattributenamelengthcheck = adopt(data.Xmlmaxattributenamelengthcheck, "xmlmaxattributenamelengthcheck")
	if !data.Xmlmaxattributes.IsNull() && !data.Xmlmaxattributes.IsUnknown() {
		// preserve configured value
	} else if val, ok := getResponseData["xmlmaxattributes"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxattributes = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxattributes = types.Int64Null()
	}
	data.Xmlmaxattributescheck = adopt(data.Xmlmaxattributescheck, "xmlmaxattributescheck")
	if !data.Xmlmaxattributevaluelength.IsNull() && !data.Xmlmaxattributevaluelength.IsUnknown() {
		// preserve configured value
	} else if val, ok := getResponseData["xmlmaxattributevaluelength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxattributevaluelength = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxattributevaluelength = types.Int64Null()
	}
	data.Xmlmaxattributevaluelengthcheck = adopt(data.Xmlmaxattributevaluelengthcheck, "xmlmaxattributevaluelengthcheck")
	if !data.Xmlmaxchardatalength.IsNull() && !data.Xmlmaxchardatalength.IsUnknown() {
		// preserve configured value
	} else if val, ok := getResponseData["xmlmaxchardatalength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxchardatalength = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxchardatalength = types.Int64Null()
	}
	data.Xmlmaxchardatalengthcheck = adopt(data.Xmlmaxchardatalengthcheck, "xmlmaxchardatalengthcheck")
	if !data.Xmlmaxelementchildren.IsNull() && !data.Xmlmaxelementchildren.IsUnknown() {
		// preserve configured value
	} else if val, ok := getResponseData["xmlmaxelementchildren"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxelementchildren = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxelementchildren = types.Int64Null()
	}
	data.Xmlmaxelementchildrencheck = adopt(data.Xmlmaxelementchildrencheck, "xmlmaxelementchildrencheck")
	if !data.Xmlmaxelementdepth.IsNull() && !data.Xmlmaxelementdepth.IsUnknown() {
		// preserve configured value
	} else if val, ok := getResponseData["xmlmaxelementdepth"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxelementdepth = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxelementdepth = types.Int64Null()
	}
	data.Xmlmaxelementdepthcheck = adopt(data.Xmlmaxelementdepthcheck, "xmlmaxelementdepthcheck")
	if !data.Xmlmaxelementnamelength.IsNull() && !data.Xmlmaxelementnamelength.IsUnknown() {
		// preserve configured value
	} else if val, ok := getResponseData["xmlmaxelementnamelength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxelementnamelength = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxelementnamelength = types.Int64Null()
	}
	data.Xmlmaxelementnamelengthcheck = adopt(data.Xmlmaxelementnamelengthcheck, "xmlmaxelementnamelengthcheck")
	if !data.Xmlmaxelements.IsNull() && !data.Xmlmaxelements.IsUnknown() {
		// preserve configured value
	} else if val, ok := getResponseData["xmlmaxelements"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxelements = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxelements = types.Int64Null()
	}
	data.Xmlmaxelementscheck = adopt(data.Xmlmaxelementscheck, "xmlmaxelementscheck")
	if !data.Xmlmaxentityexpansiondepth.IsNull() && !data.Xmlmaxentityexpansiondepth.IsUnknown() {
		// preserve configured value
	} else if val, ok := getResponseData["xmlmaxentityexpansiondepth"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxentityexpansiondepth = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxentityexpansiondepth = types.Int64Null()
	}
	data.Xmlmaxentityexpansiondepthcheck = adopt(data.Xmlmaxentityexpansiondepthcheck, "xmlmaxentityexpansiondepthcheck")
	if !data.Xmlmaxentityexpansions.IsNull() && !data.Xmlmaxentityexpansions.IsUnknown() {
		// preserve configured value
	} else if val, ok := getResponseData["xmlmaxentityexpansions"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxentityexpansions = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxentityexpansions = types.Int64Null()
	}
	data.Xmlmaxentityexpansionscheck = adopt(data.Xmlmaxentityexpansionscheck, "xmlmaxentityexpansionscheck")
	if !data.Xmlmaxfilesize.IsNull() && !data.Xmlmaxfilesize.IsUnknown() {
		// preserve configured value
	} else if val, ok := getResponseData["xmlmaxfilesize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxfilesize = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxfilesize = types.Int64Null()
	}
	data.Xmlmaxfilesizecheck = adopt(data.Xmlmaxfilesizecheck, "xmlmaxfilesizecheck")
	if !data.Xmlmaxnamespaces.IsNull() && !data.Xmlmaxnamespaces.IsUnknown() {
		// preserve configured value
	} else if val, ok := getResponseData["xmlmaxnamespaces"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxnamespaces = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxnamespaces = types.Int64Null()
	}
	data.Xmlmaxnamespacescheck = adopt(data.Xmlmaxnamespacescheck, "xmlmaxnamespacescheck")
	if !data.Xmlmaxnamespaceurilength.IsNull() && !data.Xmlmaxnamespaceurilength.IsUnknown() {
		// preserve configured value
	} else if val, ok := getResponseData["xmlmaxnamespaceurilength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxnamespaceurilength = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxnamespaceurilength = types.Int64Null()
	}
	data.Xmlmaxnamespaceurilengthcheck = adopt(data.Xmlmaxnamespaceurilengthcheck, "xmlmaxnamespaceurilengthcheck")
	if !data.Xmlmaxnodes.IsNull() && !data.Xmlmaxnodes.IsUnknown() {
		// preserve configured value
	} else if val, ok := getResponseData["xmlmaxnodes"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxnodes = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxnodes = types.Int64Null()
	}
	data.Xmlmaxnodescheck = adopt(data.Xmlmaxnodescheck, "xmlmaxnodescheck")
	if !data.Xmlmaxsoaparrayrank.IsNull() && !data.Xmlmaxsoaparrayrank.IsUnknown() {
		// preserve configured value
	} else if val, ok := getResponseData["xmlmaxsoaparrayrank"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxsoaparrayrank = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxsoaparrayrank = types.Int64Null()
	}
	if !data.Xmlmaxsoaparraysize.IsNull() && !data.Xmlmaxsoaparraysize.IsUnknown() {
		// preserve configured value
	} else if val, ok := getResponseData["xmlmaxsoaparraysize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxsoaparraysize = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxsoaparraysize = types.Int64Null()
	}
	if !data.Xmlminfilesize.IsNull() && !data.Xmlminfilesize.IsUnknown() {
		// preserve configured value
	} else if val, ok := getResponseData["xmlminfilesize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlminfilesize = types.Int64Value(intVal)
		}
	} else {
		data.Xmlminfilesize = types.Int64Null()
	}
	data.Xmlminfilesizecheck = adopt(data.Xmlminfilesizecheck, "xmlminfilesizecheck")
	data.Xmlsoaparraycheck = adopt(data.Xmlsoaparraycheck, "xmlsoaparraycheck")

	// Re-derive the canonical id so a legacy SDK v2 id is upgraded to the new format on Read.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("xmldosurl:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Xmldosurl.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}

// appfwprofile_xmldosurl_bindingSetAttrFromGetForDatasource is the DATASOURCE-side
// setter: it faithfully copies every field from the GET response (the datasource has
// no prior plan/state to preserve) and sets the composite ID.
func appfwprofile_xmldosurl_bindingSetAttrFromGetForDatasource(ctx context.Context, data *AppfwprofileXmldosurlBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileXmldosurlBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_xmldosurl_bindingSetAttrFromGetForDatasource Function")

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
	data.Xmlblockdtd = copyField("xmlblockdtd")
	data.Xmlblockexternalentities = copyField("xmlblockexternalentities")
	data.Xmlblockpi = copyField("xmlblockpi")
	data.Xmldosurl = copyField("xmldosurl")
	if val, ok := getResponseData["xmlmaxattributenamelength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxattributenamelength = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxattributenamelength = types.Int64Null()
	}
	data.Xmlmaxattributenamelengthcheck = copyField("xmlmaxattributenamelengthcheck")
	if val, ok := getResponseData["xmlmaxattributes"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxattributes = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxattributes = types.Int64Null()
	}
	data.Xmlmaxattributescheck = copyField("xmlmaxattributescheck")
	if val, ok := getResponseData["xmlmaxattributevaluelength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxattributevaluelength = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxattributevaluelength = types.Int64Null()
	}
	data.Xmlmaxattributevaluelengthcheck = copyField("xmlmaxattributevaluelengthcheck")
	if val, ok := getResponseData["xmlmaxchardatalength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxchardatalength = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxchardatalength = types.Int64Null()
	}
	data.Xmlmaxchardatalengthcheck = copyField("xmlmaxchardatalengthcheck")
	if val, ok := getResponseData["xmlmaxelementchildren"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxelementchildren = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxelementchildren = types.Int64Null()
	}
	data.Xmlmaxelementchildrencheck = copyField("xmlmaxelementchildrencheck")
	if val, ok := getResponseData["xmlmaxelementdepth"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxelementdepth = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxelementdepth = types.Int64Null()
	}
	data.Xmlmaxelementdepthcheck = copyField("xmlmaxelementdepthcheck")
	if val, ok := getResponseData["xmlmaxelementnamelength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxelementnamelength = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxelementnamelength = types.Int64Null()
	}
	data.Xmlmaxelementnamelengthcheck = copyField("xmlmaxelementnamelengthcheck")
	if val, ok := getResponseData["xmlmaxelements"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxelements = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxelements = types.Int64Null()
	}
	data.Xmlmaxelementscheck = copyField("xmlmaxelementscheck")
	if val, ok := getResponseData["xmlmaxentityexpansiondepth"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxentityexpansiondepth = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxentityexpansiondepth = types.Int64Null()
	}
	data.Xmlmaxentityexpansiondepthcheck = copyField("xmlmaxentityexpansiondepthcheck")
	if val, ok := getResponseData["xmlmaxentityexpansions"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxentityexpansions = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxentityexpansions = types.Int64Null()
	}
	data.Xmlmaxentityexpansionscheck = copyField("xmlmaxentityexpansionscheck")
	if val, ok := getResponseData["xmlmaxfilesize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxfilesize = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxfilesize = types.Int64Null()
	}
	data.Xmlmaxfilesizecheck = copyField("xmlmaxfilesizecheck")
	if val, ok := getResponseData["xmlmaxnamespaces"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxnamespaces = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxnamespaces = types.Int64Null()
	}
	data.Xmlmaxnamespacescheck = copyField("xmlmaxnamespacescheck")
	if val, ok := getResponseData["xmlmaxnamespaceurilength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxnamespaceurilength = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxnamespaceurilength = types.Int64Null()
	}
	data.Xmlmaxnamespaceurilengthcheck = copyField("xmlmaxnamespaceurilengthcheck")
	if val, ok := getResponseData["xmlmaxnodes"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlmaxnodes = types.Int64Value(intVal)
		}
	} else {
		data.Xmlmaxnodes = types.Int64Null()
	}
	data.Xmlmaxnodescheck = copyField("xmlmaxnodescheck")
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
	data.Xmlminfilesizecheck = copyField("xmlminfilesizecheck")
	data.Xmlsoaparraycheck = copyField("xmlsoaparraycheck")

	// Set ID for the datasource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("xmldosurl:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Xmldosurl.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
