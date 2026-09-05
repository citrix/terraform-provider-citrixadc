package nscentralmanagementserver

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NscentralmanagementserverDataSourceModel is the data-source-specific model,
// decoupled from NscentralmanagementserverResourceModel. A data source is a pure
// read surface, so it exposes the read/write attributes (as Computed outputs)
// AND the read-only ADM-service metadata the resource deliberately omits
// (instanceid, customerid, admserviceenvironment, admserviceconnectionstatus).
type NscentralmanagementserverDataSourceModel struct {
	Id                types.String `tfsdk:"id"`
	Activationcode    types.String `tfsdk:"activationcode"`
	Adcpassword       types.String `tfsdk:"adcpassword"`
	Adcusername       types.String `tfsdk:"adcusername"`
	Deviceprofilename types.String `tfsdk:"deviceprofilename"`
	Ipaddress         types.String `tfsdk:"ipaddress"`
	Password          types.String `tfsdk:"password"`
	Servername        types.String `tfsdk:"servername"`
	Type              types.String `tfsdk:"type"`
	Username          types.String `tfsdk:"username"`
	Validatecert      types.String `tfsdk:"validatecert"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/nscentralmanagementserver.json). Never settable; populated from GET.
	Instanceid                 types.String `tfsdk:"instanceid"`
	Customerid                 types.String `tfsdk:"customerid"`
	Admserviceenvironment      types.String `tfsdk:"admserviceenvironment"`
	Admserviceconnectionstatus types.String `tfsdk:"admserviceconnectionstatus"`
}

func NscentralmanagementserverDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"activationcode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Activation code is used to register to ADM service",
			},
			"adcpassword": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "ADC password used to create device profile on ADM",
			},
			"adcusername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ADC username used to create device profile on ADM",
			},
			"deviceprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Device profile is created on ADM and contains the user name and password of the instance(s).",
			},
			"ipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Ip Address of central management server.",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "Password for access to central management server. Required for any user account.",
			},
			"servername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Fully qualified domain name of the central management server or service-url to locate ADM service.",
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "Type of the central management server. Must be either CLOUD or ONPREM depending on whether the server is on the cloud or on premise.",
			},
			"username": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Username for access to central management server. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or\nsingle quotation marks (for example, \"my ns centralmgmtserver\" or \"my ns centralmgmtserver\").",
			},
			"validatecert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "validate the server certificate for secure SSL connections.",
			},

			// Read-only (GET-only) ADM-service metadata surfaced by the data source.
			"instanceid": schema.StringAttribute{
				Computed:    true,
				Description: "Instance ID of the customer provided by Trust.",
			},
			"customerid": schema.StringAttribute{
				Computed:    true,
				Description: "Customer ID of the citrix cloud customer.",
			},
			"admserviceenvironment": schema.StringAttribute{
				Computed:    true,
				Description: "ADM service environment (PRODUCTION/STAGING/DEV).",
			},
			"admserviceconnectionstatus": schema.StringAttribute{
				Computed:    true,
				Description: "Built-in agent's (mastools) connection status to ADM service.",
			},
		},
	}
}

// nscentralmanagementserverDataSourceSetAttrFromGet projects a NITRO
// nscentralmanagementserver GET response onto the data-source model. Attributes
// are filled from the GET (or left Null when the GET omits them) via the shared
// utils.MapGet* helpers; secret/write-only inputs the GET never returns are Null.
func nscentralmanagementserverDataSourceSetAttrFromGet(ctx context.Context, data *NscentralmanagementserverDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In nscentralmanagementserverDataSourceSetAttrFromGet Function")

	if v, ok := g["type"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Type = types.StringValue(utils.AnyToString(v))
	}

	data.Activationcode = utils.MapGetString(g, "activationcode")
	data.Adcusername = utils.MapGetString(g, "adcusername")
	data.Deviceprofilename = utils.MapGetString(g, "deviceprofilename")
	data.Ipaddress = utils.MapGetString(g, "ipaddress")
	data.Servername = utils.MapGetString(g, "servername")
	data.Username = utils.MapGetString(g, "username")
	data.Validatecert = utils.MapGetString(g, "validatecert")

	// Secret / write-only inputs the GET never returns -> Null.
	data.Adcpassword = types.StringNull()
	data.Password = types.StringNull()

	// Read-only ADM-service metadata.
	data.Instanceid = utils.MapGetString(g, "instanceid")
	data.Customerid = utils.MapGetString(g, "customerid")
	data.Admserviceenvironment = utils.MapGetString(g, "admserviceenvironment")
	data.Admserviceconnectionstatus = utils.MapGetString(g, "admserviceconnectionstatus")
}
