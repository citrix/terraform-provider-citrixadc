package appfwprofile_xmlvalidationurl_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwprofileXmlvalidationurlBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"alertonly": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
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
				Required:    true,
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
