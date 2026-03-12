package transformaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func TransformactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this URL Transformation action.",
			},
			"cookiedomainfrom": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Pattern that matches the domain to be transformed in Set-Cookie headers.",
			},
			"cookiedomaininto": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "PCRE-format regular expression that describes the transformation to be performed on cookie domains that match the cookieDomainFrom pattern. \nNOTE: The cookie domain to be transformed is extracted from the request.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the URL transformation action.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the URL Transformation action is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, my transform action or my transform action).",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Positive integer specifying the priority of the action within the profile. A lower number specifies a higher priority. Must be unique within the list of actions bound to the profile. Policies are evaluated in the order of their priority numbers, and the first policy that matches is applied.",
			},
			"profilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the URL Transformation profile with which to associate this action.",
			},
			"requrlfrom": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "PCRE-format regular expression that describes the request URL pattern to be transformed.",
			},
			"requrlinto": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "PCRE-format regular expression that describes the transformation to be performed on URLs that match the reqUrlFrom pattern.",
			},
			"resurlfrom": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "PCRE-format regular expression that describes the response URL pattern to be transformed.",
			},
			"resurlinto": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "PCRE-format regular expression that describes the transformation to be performed on URLs that match the resUrlFrom pattern.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable this action.",
			},
		},
	}
}
