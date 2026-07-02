package cloudparameter

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CloudparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"activationcode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Activation code for the NGS Connector instance",
			},
			"connectorresidence": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Identifies whether the connector is located Onprem, Aws or Azure",
			},
			"controllerfqdn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "FQDN of the controller to which the Citrix ADC SDProxy Connects",
			},
			"controllerport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number of the controller to which the Citrix ADC SDProxy connects",
			},
			"customerid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Customer ID of the citrix cloud customer",
			},
			"deployment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Describes if the customer is a Staging/Production or Dev Citrix Cloud customer",
			},
			"instanceid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Instance ID of the customer provided by Trust",
			},
			"resourcelocation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Resource Location of the customer provided by Trust",
			},
		},
	}
}