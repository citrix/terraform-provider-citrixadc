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
			"xmladditionalsoapheaders": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow addtional soap headers.",
			},
			"xmlendpointcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Modifies the behaviour of the Request URL validation w.r.t. the Service URL.\n	If set to ABSOLUTE, the entire request URL is validated with the entire URL mentioned in Service of the associated WSDL.\n		eg: Service URL: http://example.org/ExampleService, Request URL: http//example.com/ExampleService would FAIL the validation.\n	If set to RELAIVE, only the non-hostname part of the request URL is validated against the non-hostname part of the Service URL.\n		eg: Service URL: http://example.org/ExampleService, Request URL: http//example.com/ExampleService would PASS the validation.",
			},
			"xmlrequestschema": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "XML Schema object for request validation .",
			},
			"xmlresponseschema": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "XML Schema object for response validation.",
			},
			"xmlvalidateresponse": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Validate response message.",
			},
			"xmlvalidatesoapenvelope": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Validate SOAP Evelope only.",
			},
			"xmlvalidationurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "XML Validation URL regular expression.",
			},
			"xmlwsdl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "WSDL object for soap request validation.",
			},
		},
	}
}

func appfwprofile_xmlvalidationurl_bindingGetThePayloadFromtheConfig(ctx context.Context, data *AppfwprofileXmlvalidationurlBindingResourceModel) appfw.Appfwprofilexmlvalidationurlbinding {
	tflog.Debug(ctx, "In appfwprofile_xmlvalidationurl_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appfwprofile_xmlvalidationurl_binding := appfw.Appfwprofilexmlvalidationurlbinding{}
	if !data.Alertonly.IsNull() {
		appfwprofile_xmlvalidationurl_binding.Alertonly = data.Alertonly.ValueString()
	}
	if !data.Comment.IsNull() {
		appfwprofile_xmlvalidationurl_binding.Comment = data.Comment.ValueString()
	}
	if !data.Isautodeployed.IsNull() {
		appfwprofile_xmlvalidationurl_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.Name.IsNull() {
		appfwprofile_xmlvalidationurl_binding.Name = data.Name.ValueString()
	}
	if !data.Resourceid.IsNull() {
		appfwprofile_xmlvalidationurl_binding.Resourceid = data.Resourceid.ValueString()
	}
	if !data.State.IsNull() {
		appfwprofile_xmlvalidationurl_binding.State = data.State.ValueString()
	}
	if !data.Xmladditionalsoapheaders.IsNull() {
		appfwprofile_xmlvalidationurl_binding.Xmladditionalsoapheaders = data.Xmladditionalsoapheaders.ValueString()
	}
	if !data.Xmlendpointcheck.IsNull() {
		appfwprofile_xmlvalidationurl_binding.Xmlendpointcheck = data.Xmlendpointcheck.ValueString()
	}
	if !data.Xmlrequestschema.IsNull() {
		appfwprofile_xmlvalidationurl_binding.Xmlrequestschema = data.Xmlrequestschema.ValueString()
	}
	if !data.Xmlresponseschema.IsNull() {
		appfwprofile_xmlvalidationurl_binding.Xmlresponseschema = data.Xmlresponseschema.ValueString()
	}
	if !data.Xmlvalidateresponse.IsNull() {
		appfwprofile_xmlvalidationurl_binding.Xmlvalidateresponse = data.Xmlvalidateresponse.ValueString()
	}
	if !data.Xmlvalidatesoapenvelope.IsNull() {
		appfwprofile_xmlvalidationurl_binding.Xmlvalidatesoapenvelope = data.Xmlvalidatesoapenvelope.ValueString()
	}
	if !data.Xmlvalidationurl.IsNull() {
		appfwprofile_xmlvalidationurl_binding.Xmlvalidationurl = data.Xmlvalidationurl.ValueString()
	}
	if !data.Xmlwsdl.IsNull() {
		appfwprofile_xmlvalidationurl_binding.Xmlwsdl = data.Xmlwsdl.ValueString()
	}

	return appfwprofile_xmlvalidationurl_binding
}

func appfwprofile_xmlvalidationurl_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileXmlvalidationurlBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileXmlvalidationurlBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_xmlvalidationurl_bindingSetAttrFromGet Function")

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
	if val, ok := getResponseData["xmladditionalsoapheaders"]; ok && val != nil {
		data.Xmladditionalsoapheaders = types.StringValue(val.(string))
	} else {
		data.Xmladditionalsoapheaders = types.StringNull()
	}
	if val, ok := getResponseData["xmlendpointcheck"]; ok && val != nil {
		data.Xmlendpointcheck = types.StringValue(val.(string))
	} else {
		data.Xmlendpointcheck = types.StringNull()
	}
	if val, ok := getResponseData["xmlrequestschema"]; ok && val != nil {
		data.Xmlrequestschema = types.StringValue(val.(string))
	} else {
		data.Xmlrequestschema = types.StringNull()
	}
	if val, ok := getResponseData["xmlresponseschema"]; ok && val != nil {
		data.Xmlresponseschema = types.StringValue(val.(string))
	} else {
		data.Xmlresponseschema = types.StringNull()
	}
	if val, ok := getResponseData["xmlvalidateresponse"]; ok && val != nil {
		data.Xmlvalidateresponse = types.StringValue(val.(string))
	} else {
		data.Xmlvalidateresponse = types.StringNull()
	}
	if val, ok := getResponseData["xmlvalidatesoapenvelope"]; ok && val != nil {
		data.Xmlvalidatesoapenvelope = types.StringValue(val.(string))
	} else {
		data.Xmlvalidatesoapenvelope = types.StringNull()
	}
	if val, ok := getResponseData["xmlvalidationurl"]; ok && val != nil {
		data.Xmlvalidationurl = types.StringValue(val.(string))
	} else {
		data.Xmlvalidationurl = types.StringNull()
	}
	if val, ok := getResponseData["xmlwsdl"]; ok && val != nil {
		data.Xmlwsdl = types.StringValue(val.(string))
	} else {
		data.Xmlwsdl = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("xmlvalidationurl:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Xmlvalidationurl.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
