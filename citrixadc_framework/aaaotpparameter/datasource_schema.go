package aaaotpparameter

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AaaotpparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"encryption": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "To encrypt otp secret in AD or not. Default value is OFF",
			},
			"maxotpdevices": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of otp devices user can register. Default value is 4. Max value is 255",
			},
			"otptype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Input flag to generate OTP for the given type. Possible values = gwtest",
			},
		},
	}
}
