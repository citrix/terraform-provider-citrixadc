package responderpolicy

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ResponderpolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the responder action to perform if the request matches this responder policy.",
			},
			"appflowaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "AppFlow action to invoke for requests that match this policy.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any type of information about this responder policy.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the messagelog action to use for requests that match this policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the responder policy.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that the policy uses to determine whether to respond to the specified request.",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform if the result of policy evaluation is undefined (UNDEF).",
			},
		},
		// The convenience-block sets are shared with the resource model; they are
		// exposed here as computed outputs so the shared model maps cleanly.
		Blocks: map[string]schema.Block{
			"globalbinding": schema.SetNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"gotopriorityexpression": schema.StringAttribute{Computed: true},
						"invoke":                 schema.BoolAttribute{Computed: true},
						"labelname":              schema.StringAttribute{Computed: true},
						"labeltype":              schema.StringAttribute{Computed: true},
						"policyname":             schema.StringAttribute{Computed: true},
						"priority":               schema.Int64Attribute{Computed: true},
						"type":                   schema.StringAttribute{Computed: true},
					},
				},
			},
			"lbvserverbinding": schema.SetNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"bindpoint":              schema.StringAttribute{Computed: true},
						"gotopriorityexpression": schema.StringAttribute{Computed: true},
						"invoke":                 schema.BoolAttribute{Computed: true},
						"labelname":              schema.StringAttribute{Computed: true},
						"labeltype":              schema.StringAttribute{Computed: true},
						"name":                   schema.StringAttribute{Computed: true},
						"priority":               schema.Int64Attribute{Computed: true},
					},
				},
			},
			"csvserverbinding": schema.SetNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"bindpoint":              schema.StringAttribute{Computed: true},
						"gotopriorityexpression": schema.StringAttribute{Computed: true},
						"invoke":                 schema.BoolAttribute{Computed: true},
						"labelname":              schema.StringAttribute{Computed: true},
						"labeltype":              schema.StringAttribute{Computed: true},
						"name":                   schema.StringAttribute{Computed: true},
						"policyname":             schema.StringAttribute{Computed: true},
						"priority":               schema.Int64Attribute{Computed: true},
						"targetlbvserver":        schema.StringAttribute{Computed: true},
					},
				},
			},
		},
	}
}
