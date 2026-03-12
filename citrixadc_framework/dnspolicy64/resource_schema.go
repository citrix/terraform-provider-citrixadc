package dnspolicy64

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Dnspolicy64ResourceModel describes the resource data model.
type Dnspolicy64ResourceModel struct {
	Id     types.String `tfsdk:"id"`
	Action types.String `tfsdk:"action"`
	Name   types.String `tfsdk:"name"`
	Rule   types.String `tfsdk:"rule"`
}

func (r *Dnspolicy64Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dnspolicy64 resource.",
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DNS64 action to perform when the rule evaluates to TRUE. The built in actions function as follows:\n* A default dns64 action with prefix <default prefix> and mapped and exclude are any\nYou can create custom actions by using the add dns action command in the CLI or the DNS64 > Actions > Create DNS64 Action dialog box in the Citrix ADC configuration utility.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the DNS64 policy.",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Expression against which DNS traffic is evaluated.\nNote:\n* On the command line interface, if the expression includes blank spaces, the entire expression must be enclosed in double quotation marks.\n* If the expression itself includes double quotation marks, you must escape the quotations by using the  character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.\nExample: CLIENT.IP.SRC.IN_SUBENT(23.34.0.0/16)",
			},
		},
	}
}

func dnspolicy64GetThePayloadFromtheConfig(ctx context.Context, data *Dnspolicy64ResourceModel) dns.Dnspolicy64 {
	tflog.Debug(ctx, "In dnspolicy64GetThePayloadFromtheConfig Function")

	// Create API request body from the model
	dnspolicy64 := dns.Dnspolicy64{}
	if !data.Action.IsNull() {
		dnspolicy64.Action = data.Action.ValueString()
	}
	if !data.Name.IsNull() {
		dnspolicy64.Name = data.Name.ValueString()
	}
	if !data.Rule.IsNull() {
		dnspolicy64.Rule = data.Rule.ValueString()
	}

	return dnspolicy64
}

func dnspolicy64SetAttrFromGet(ctx context.Context, data *Dnspolicy64ResourceModel, getResponseData map[string]interface{}) *Dnspolicy64ResourceModel {
	tflog.Debug(ctx, "In dnspolicy64SetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	} else {
		data.Action = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["rule"]; ok && val != nil {
		data.Rule = types.StringValue(val.(string))
	} else {
		data.Rule = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
