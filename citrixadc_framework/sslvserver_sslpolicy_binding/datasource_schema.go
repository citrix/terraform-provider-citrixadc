package sslvserver_sslpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslvserverSslpolicyBindingDataSourceModel is the data-source-specific model,
// decoupled from SslvserverSslpolicyBindingResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attribute the resource deliberately omits
// (polinherit). The Framework's per-attribute model <-> schema reflection
// requires this model to have exactly the attributes the data-source schema
// declares.
type SslvserverSslpolicyBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Type                   types.String `tfsdk:"type"`
	Vservername            types.String `tfsdk:"vservername"`

	// Read-only (GET-only) attribute from the NITRO doc read-only set
	// (zion73x_readonly/sslvserver_sslpolicy_binding.json). Never settable;
	// populated from GET.
	Polinherit types.Int64 `tfsdk:"polinherit"`
}

func SslvserverSslpolicyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"invoke": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Invoke flag. This attribute is relevant only for ADVANCED policies",
			},
			"labelname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the label to invoke if the current policy rule evaluates to TRUE.",
			},
			"labeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of policy label invocation.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "The name of the SSL policy binding.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The priority of the policies bound to this SSL service",
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "Bind point to which to bind the policy. Possible Values: REQUEST, INTERCEPT_REQ and CLIENTHELLO_REQ. These bindpoints mean:\n1. REQUEST: Policy evaluation will be done at appplication above SSL. This bindpoint is default and is used for actions based on clientauth and client cert.\n2. INTERCEPT_REQ: Policy evaluation will be done during SSL handshake to decide whether to intercept or not. Actions allowed with this type are: INTERCEPT, BYPASS and RESET.\n3. CLIENTHELLO_REQ: Policy evaluation will be done during handling of Client Hello Request. Action allowed with this type is: RESET, FORWARD and PICKCACERTGRP.",
			},
			"vservername": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL virtual server.",
			},

			// Read-only (GET-only) attribute surfaced by the data source (this is
			// intentionally NOT modeled on the resource). Computed.
			"polinherit": schema.Int64Attribute{
				Computed:    true,
				Description: "Whether the bound policy is a inherited policy or not.",
			},
		},
	}
}

// sslvserver_sslpolicy_bindingDataSourceSetAttrFromGet projects a NITRO
// sslvserver_sslpolicy_binding GET response onto the data-source model. A data
// source has no plan/apply reconciliation, so attributes are simply filled from
// the GET (or left Null when the GET omits them) via the shared utils.MapGet*
// helpers, and the composite ID is set.
func sslvserver_sslpolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *SslvserverSslpolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In sslvserver_sslpolicy_bindingDataSourceSetAttrFromGet Function")

	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Invoke = utils.MapGetBool(g, "invoke")
	data.Labelname = utils.MapGetString(g, "labelname")
	data.Labeltype = utils.MapGetString(g, "labeltype")
	data.Policyname = utils.MapGetString(g, "policyname")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Type = utils.MapGetString(g, "type")
	data.Vservername = utils.MapGetString(g, "vservername")

	// Read-only (GET-only) attribute.
	data.Polinherit = utils.MapGetInt64(g, "polinherit")

	// Set composite ID for the datasource.
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("priority:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Priority.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("type:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Type.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("vservername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vservername.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
