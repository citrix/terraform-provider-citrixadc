package lbaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func LbactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comment. Any type of information about this LB action.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the LB action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my lb action\" or 'my lb action').",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the LB action.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my lb action\" or my lb action').",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of an LB action. Available settings function as follows:\n* NOLBACTION - Does not consider LB action in making LB decision.\n* SELECTIONORDER - services bound to vserver with order specified in value parameter is considerd for lb/gslb decision.",
			},
			"value": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "The selection order list used during lb/gslb decision. Preference of services during lb/gslb decision is as follows - services corresponding to first order specified in the sequence is considered first, services corresponding to second order specified in the sequence is considered next and so on. For example, if -value 2 1 3 is specified here and service-1 bound to a vserver with order 1, service-2 bound to a vserver with order 2 and  service-3 bound to a vserver with order 3. Then preference of selecting services in LB decision is as follows: service-2, service-1, service-3.",
			},
		},
	}
}
