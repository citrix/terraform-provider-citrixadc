package transformaction

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// TransformactionDataSourceModel is the data-source-specific model, decoupled
// from TransformactionResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits
// (continuematching, ...). Every non-key attribute is Computed.
type TransformactionDataSourceModel struct {
	Id               types.String `tfsdk:"id"`
	Comment          types.String `tfsdk:"comment"`
	Cookiedomainfrom types.String `tfsdk:"cookiedomainfrom"`
	Cookiedomaininto types.String `tfsdk:"cookiedomaininto"`
	Name             types.String `tfsdk:"name"` // Required lookup key
	Priority         types.Int64  `tfsdk:"priority"`
	Profilename      types.String `tfsdk:"profilename"`
	Requrlfrom       types.String `tfsdk:"requrlfrom"`
	Requrlinto       types.String `tfsdk:"requrlinto"`
	Resurlfrom       types.String `tfsdk:"resurlfrom"`
	Resurlinto       types.String `tfsdk:"resurlinto"`
	State            types.String `tfsdk:"state"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/transformaction.json). Never settable; populated from GET.
	Continuematching types.String `tfsdk:"continuematching"`
}

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

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"continuematching": schema.StringAttribute{
				Computed:    true,
				Description: "Continue transforming using the next rule in the list. Possible values = ON, OFF",
			},
		},
	}
}

// transformactionDataSourceSetAttrFromGet projects a NITRO transformaction GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that projection.
func transformactionDataSourceSetAttrFromGet(ctx context.Context, data *TransformactionDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In transformactionDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Comment = utils.MapGetString(g, "comment")
	data.Cookiedomainfrom = utils.MapGetString(g, "cookiedomainfrom")
	data.Cookiedomaininto = utils.MapGetString(g, "cookiedomaininto")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Profilename = utils.MapGetString(g, "profilename")
	data.Requrlfrom = utils.MapGetString(g, "requrlfrom")
	data.Requrlinto = utils.MapGetString(g, "requrlinto")
	data.Resurlfrom = utils.MapGetString(g, "resurlfrom")
	data.Resurlinto = utils.MapGetString(g, "resurlinto")
	data.State = utils.MapGetString(g, "state")

	// Read-only attributes.
	data.Continuematching = utils.MapGetString(g, "continuematching")
}
