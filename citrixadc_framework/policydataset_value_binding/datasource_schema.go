package policydataset_value_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func PolicydatasetValueBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this dataset or a data bound to this dataset.",
			},
			"endrange": schema.StringAttribute{
				Required:    true,
				Description: "The dataset entry is a range from <value> through <end_range>, inclusive. endRange cannot be used if value is an ipv4 or ipv6 subnet and endRange cannot itself be a subnet.",
			},
			"index": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The index of the value (ipv4, ipv6, number) associated with the set.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the dataset to which to bind the value.",
			},
			"value": schema.StringAttribute{
				Required:    true,
				Description: "Value of the specified type that is associated with the dataset. For ipv4 and ipv6, value can be a subnet using the slash notation address/n, where address is the beginning of the subnet and n is the number of left-most bits set in the subnet mask, defining the end of the subnet. The start address will be masked by the subnet mask if necessary, for example for 192.128.128.0/10, the start address will be 192.128.0.0.",
			},
		},
	}
}
