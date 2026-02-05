package nscapacity

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NscapacityDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"bandwidth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "System bandwidth limit.",
			},
			"edition": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Product edition.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password to use when authenticating with ADM Agent for LAS licensing.",
			},
			"platform": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "appliance platform type.",
			},
			"unit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bandwidth unit.",
			},
			"username": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Username to authenticate with ADM Agent for LAS licensing. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
			"vcpu": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "licensed using vcpu pool.",
			},
		},
	}
}
