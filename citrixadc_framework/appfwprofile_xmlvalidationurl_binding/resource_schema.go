package appfwprofile_xmlvalidationurl_binding

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

// AppfwprofileXmlvalidationurlBindingResourceModel describes the resource data model.
type AppfwprofileXmlvalidationurlBindingResourceModel struct {
	Id                       types.String `tfsdk:"id"`
	Alertonly                types.String `tfsdk:"alertonly"`
	Comment                  types.String `tfsdk:"comment"`
	Isautodeployed           types.String `tfsdk:"isautodeployed"`
	Name                     types.String `tfsdk:"name"`
	Resourceid               types.String `tfsdk:"resourceid"`
	Ruletype                 types.String `tfsdk:"ruletype"`
	State                    types.String `tfsdk:"state"`
	Xmladditionalsoapheaders types.String `tfsdk:"xmladditionalsoapheaders"`
	Xmlendpointcheck         types.String `tfsdk:"xmlendpointcheck"`
	Xmlrequestschema         types.String `tfsdk:"xmlrequestschema"`
	Xmlresponseschema        types.String `tfsdk:"xmlresponseschema"`
	Xmlvalidateresponse      types.String `tfsdk:"xmlvalidateresponse"`
	Xmlvalidatesoapenvelope  types.String `tfsdk:"xmlvalidatesoapenvelope"`
	Xmlvalidationurl         types.String `tfsdk:"xmlvalidationurl"`
	Xmlwsdl                  types.String `tfsdk:"xmlwsdl"`
}

func (r *AppfwprofileXmlvalidationurlBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwprofile_xmlvalidationurl_binding resource.",
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
			"xmladditionalsoapheaders": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Allow addtional soap headers.",
			},
			"xmlendpointcheck": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Modifies the behaviour of the Request URL validation w.r.t. the Service URL.\n	If set to ABSOLUTE, the entire request URL is validated with the entire URL mentioned in Service of the associated WSDL.\n		eg: Service URL: http://example.org/ExampleService, Request URL: http//example.com/ExampleService would FAIL the validation.\n	If set to RELAIVE, only the non-hostname part of the request URL is validated against the non-hostname part of the Service URL.\n		eg: Service URL: http://example.org/ExampleService, Request URL: http//example.com/ExampleService would PASS the validation.",
			},
			"xmlrequestschema": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "XML Schema object for request validation .",
			},
			"xmlresponseschema": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "XML Schema object for response validation.",
			},
			"xmlvalidateresponse": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Validate response message.",
			},
			"xmlvalidatesoapenvelope": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Validate SOAP Evelope only.",
			},
			"xmlvalidationurl": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "XML Validation URL regular expression.",
			},
			"xmlwsdl": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "WSDL object for soap request validation.",
			},
		},
	}
}

func appfwprofile_xmlvalidationurl_bindingGetThePayloadFromthePlan(ctx context.Context, data *AppfwprofileXmlvalidationurlBindingResourceModel) appfw.Appfwprofilexmlvalidationurlbinding {
	tflog.Debug(ctx, "In appfwprofile_xmlvalidationurl_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	appfwprofile_xmlvalidationurl_binding := appfw.Appfwprofilexmlvalidationurlbinding{}
	if !data.Alertonly.IsNull() && !data.Alertonly.IsUnknown() {
		appfwprofile_xmlvalidationurl_binding.Alertonly = data.Alertonly.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appfwprofile_xmlvalidationurl_binding.Comment = data.Comment.ValueString()
	}
	if !data.Isautodeployed.IsNull() && !data.Isautodeployed.IsUnknown() {
		appfwprofile_xmlvalidationurl_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwprofile_xmlvalidationurl_binding.Name = data.Name.ValueString()
	}
	if !data.Resourceid.IsNull() && !data.Resourceid.IsUnknown() {
		appfwprofile_xmlvalidationurl_binding.Resourceid = data.Resourceid.ValueString()
	}
	if !data.Ruletype.IsNull() && !data.Ruletype.IsUnknown() {
		appfwprofile_xmlvalidationurl_binding.Ruletype = data.Ruletype.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		appfwprofile_xmlvalidationurl_binding.State = data.State.ValueString()
	}
	if !data.Xmladditionalsoapheaders.IsNull() && !data.Xmladditionalsoapheaders.IsUnknown() {
		appfwprofile_xmlvalidationurl_binding.Xmladditionalsoapheaders = data.Xmladditionalsoapheaders.ValueString()
	}
	if !data.Xmlendpointcheck.IsNull() && !data.Xmlendpointcheck.IsUnknown() {
		appfwprofile_xmlvalidationurl_binding.Xmlendpointcheck = data.Xmlendpointcheck.ValueString()
	}
	if !data.Xmlrequestschema.IsNull() && !data.Xmlrequestschema.IsUnknown() {
		appfwprofile_xmlvalidationurl_binding.Xmlrequestschema = data.Xmlrequestschema.ValueString()
	}
	if !data.Xmlresponseschema.IsNull() && !data.Xmlresponseschema.IsUnknown() {
		appfwprofile_xmlvalidationurl_binding.Xmlresponseschema = data.Xmlresponseschema.ValueString()
	}
	if !data.Xmlvalidateresponse.IsNull() && !data.Xmlvalidateresponse.IsUnknown() {
		appfwprofile_xmlvalidationurl_binding.Xmlvalidateresponse = data.Xmlvalidateresponse.ValueString()
	}
	if !data.Xmlvalidatesoapenvelope.IsNull() && !data.Xmlvalidatesoapenvelope.IsUnknown() {
		appfwprofile_xmlvalidationurl_binding.Xmlvalidatesoapenvelope = data.Xmlvalidatesoapenvelope.ValueString()
	}
	if !data.Xmlvalidationurl.IsNull() && !data.Xmlvalidationurl.IsUnknown() {
		appfwprofile_xmlvalidationurl_binding.Xmlvalidationurl = data.Xmlvalidationurl.ValueString()
	}
	if !data.Xmlwsdl.IsNull() && !data.Xmlwsdl.IsUnknown() {
		appfwprofile_xmlvalidationurl_binding.Xmlwsdl = data.Xmlwsdl.ValueString()
	}

	return appfwprofile_xmlvalidationurl_binding
}

// appfwprofile_xmlvalidationurl_bindingSetAttrFromGet is the RESOURCE-side setter.
// All attributes are RequiresReplace (no update endpoint) and the NITRO server may
// echo server-defaulted/normalized values for fields like alertonly and isautodeployed
// (the SDK v2 resource explicitly skipped re-setting alertonly/isautodeployed in Read
// to preserve the configured values). To avoid "inconsistent result after apply" we
// adopt the GET value only when the model field is currently null/unknown (e.g. import);
// otherwise we preserve the configured plan/state value. The ID is set once in Create
// and is preserved here.
func appfwprofile_xmlvalidationurl_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileXmlvalidationurlBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileXmlvalidationurlBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_xmlvalidationurl_bindingSetAttrFromGet Function")

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
	data.Xmladditionalsoapheaders = adopt(data.Xmladditionalsoapheaders, "xmladditionalsoapheaders")
	data.Xmlendpointcheck = adopt(data.Xmlendpointcheck, "xmlendpointcheck")
	data.Xmlrequestschema = adopt(data.Xmlrequestschema, "xmlrequestschema")
	data.Xmlresponseschema = adopt(data.Xmlresponseschema, "xmlresponseschema")
	data.Xmlvalidateresponse = adopt(data.Xmlvalidateresponse, "xmlvalidateresponse")
	data.Xmlvalidatesoapenvelope = adopt(data.Xmlvalidatesoapenvelope, "xmlvalidatesoapenvelope")
	data.Xmlvalidationurl = adopt(data.Xmlvalidationurl, "xmlvalidationurl")
	data.Xmlwsdl = adopt(data.Xmlwsdl, "xmlwsdl")

	// Re-derive the canonical id so a legacy SDK v2 id is upgraded to the new format on Read.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("xmlvalidationurl:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Xmlvalidationurl.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}

// appfwprofile_xmlvalidationurl_bindingSetAttrFromGetForDatasource is the DATASOURCE-side
// setter: it faithfully copies every field from the GET response (the datasource has no
// prior plan/state to preserve) and sets the composite ID.
func appfwprofile_xmlvalidationurl_bindingSetAttrFromGetForDatasource(ctx context.Context, data *AppfwprofileXmlvalidationurlBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileXmlvalidationurlBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_xmlvalidationurl_bindingSetAttrFromGetForDatasource Function")

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
	data.Xmladditionalsoapheaders = copyField("xmladditionalsoapheaders")
	data.Xmlendpointcheck = copyField("xmlendpointcheck")
	data.Xmlrequestschema = copyField("xmlrequestschema")
	data.Xmlresponseschema = copyField("xmlresponseschema")
	data.Xmlvalidateresponse = copyField("xmlvalidateresponse")
	data.Xmlvalidatesoapenvelope = copyField("xmlvalidatesoapenvelope")
	data.Xmlvalidationurl = copyField("xmlvalidationurl")
	data.Xmlwsdl = copyField("xmlwsdl")

	// Set ID for the datasource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("xmlvalidationurl:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Xmlvalidationurl.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
