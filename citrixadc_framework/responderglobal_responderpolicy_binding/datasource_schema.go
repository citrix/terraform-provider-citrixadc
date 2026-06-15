package responderglobal_responderpolicy_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ResponderglobalResponderpolicyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"globalbindtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"invoke": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.",
			},
			"labelname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the policy label to invoke. If the current policy evaluates to TRUE, the invoke parameter is set, and Label Type is policylabel.",
			},
			"labeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of invocation, Available settings function as follows:\n* vserver - Forward the request to the specified virtual server.\n* policylabel - Invoke the specified policy label.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the responder policy.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the priority of the policy.",
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "Specifies the bind point whose policies you want to display. Available settings function as follows:\n* REQ_OVERRIDE - Request override. Binds the policy to the priority request queue.\n* REQ_DEFAULT - Binds the policy to the default request queue.\n* OTHERTCP_REQ_OVERRIDE - Binds the policy to the non-HTTP TCP priority request queue.\n* OTHERTCP_REQ_DEFAULT - Binds the policy to the non-HTTP TCP default request queue..\n* SIPUDP_REQ_OVERRIDE - Binds the policy to the SIP UDP priority response queue..\n* SIPUDP_REQ_DEFAULT - Binds the policy to the SIP UDP default response queue.\n* RADIUS_REQ_OVERRIDE - Binds the policy to the RADIUS priority response queue..\n* RADIUS_REQ_DEFAULT - Binds the policy to the RADIUS default response queue.\n* MSSQL_REQ_OVERRIDE - Binds the policy to the Microsoft SQL priority response queue..\n* MSSQL_REQ_DEFAULT - Binds the policy to the Microsoft SQL default response queue.\n* MYSQL_REQ_OVERRIDE - Binds the policy to the MySQL priority response queue.\n* MYSQL_REQ_DEFAULT - Binds the policy to the MySQL default response queue.\n* HTTPQUIC_REQ_OVERRIDE - Binds the policy to the HTTP_QUIC override response queue.\n* HTTPQUIC_REQ_DEFAULT - Binds the policy to the HTTP_QUIC default response queue.",
			},
		},
	}
}
