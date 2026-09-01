package sslservice_sslpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SslserviceSslpolicyBindingDataSourceModel is the data-source-specific model. A
// data source is a pure read surface, so in addition to the read/write attributes
// (surfaced as Computed outputs) it exposes the read-only (GET-only) NITRO
// attributes that the resource intentionally omits.
type SslserviceSslpolicyBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Servicename            types.String `tfsdk:"servicename"`
	Type                   types.String `tfsdk:"type"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/sslservice_sslpolicy_binding.json). Never settable;
	// populated from GET, null when the appliance omits them.
	Polinherit types.Int64 `tfsdk:"polinherit"`
}

func SslserviceSslpolicyBindingDataSourceSchema() schema.Schema {
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
				Description: "The SSL policy binding.",
			},
			"priority": schema.Int64Attribute{
				Required:    true,
				Description: "The priority of the policies bound to this SSL service",
			},
			"servicename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL service for which to set advanced configuration.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The phase of the SSL connection in which the policy rule is evaluated.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"polinherit": schema.Int64Attribute{
				Computed:    true,
				Description: "Whether the bound policy is a inherited policy or not.",
			},
		},
	}
}

// sslservice_sslpolicy_bindingDataSourceSetAttrFromGet projects a NITRO
// sslservice_sslpolicy_binding GET response onto the data-source model and sets the
// composite ID. Attributes are filled from the GET (or left Null when the GET omits
// them) via the shared utils.MapGet* helpers.
func sslservice_sslpolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *SslserviceSslpolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In sslservice_sslpolicy_bindingDataSourceSetAttrFromGet Function")

	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Invoke = utils.MapGetBool(g, "invoke")
	data.Labelname = utils.MapGetString(g, "labelname")
	data.Labeltype = utils.MapGetString(g, "labeltype")
	data.Policyname = utils.MapGetString(g, "policyname")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Servicename = utils.MapGetString(g, "servicename")
	data.Type = utils.MapGetString(g, "type")

	// Read-only (GET-only) attributes.
	data.Polinherit = utils.MapGetInt64(g, "polinherit")

	// Datasource has no Create — set the composite ID here.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("priority:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Priority.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("servicename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicename.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
