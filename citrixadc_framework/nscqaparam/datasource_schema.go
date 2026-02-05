package nscqaparam

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NscqaparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"harqretxdelay": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "HARQ retransmission delay (in ms).",
			},
			"lr1coeflist": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "coefficients values for Label1.",
			},
			"lr1probthresh": schema.Float64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Probability threshold values for LR model to differentiate between NET1 and reset(NET2 and NET3).",
			},
			"lr2coeflist": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "coefficients values for Label 2.",
			},
			"lr2probthresh": schema.Float64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Probability threshold values for LR model to differentiate between NET2 and NET3.",
			},
			"minrttnet1": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "MIN RTT (in ms) for the first network.",
			},
			"minrttnet2": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "MIN RTT (in ms) for the second network.",
			},
			"minrttnet3": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "MIN RTT (in ms) for the third network.",
			},
			"net1cclscale": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Three congestion level scores limits corresponding to None, Low, Medium.",
			},
			"net1csqscale": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Three signal quality level scores limits corresponding to Excellent, Good, Fair.",
			},
			"net1label": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the network label.",
			},
			"net1logcoef": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Connection quality ranking Log coefficients of network 1.",
			},
			"net2cclscale": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Three congestion level scores limits corresponding to None, Low, Medium.",
			},
			"net2csqscale": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Three signal quality level scores limits corresponding to Excellent, Good, Fair.",
			},
			"net2label": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the network label 2.",
			},
			"net2logcoef": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Connnection quality ranking Log coefficients of network 2.",
			},
			"net3cclscale": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Three congestion level scores limits corresponding to None, Low, Medium.",
			},
			"net3csqscale": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Three signal quality level scores limits corresponding to Excellent, Good, Fair.",
			},
			"net3label": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the network label 3.",
			},
			"net3logcoef": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Connection quality ranking Log coefficients of network 3.",
			},
		},
	}
}
