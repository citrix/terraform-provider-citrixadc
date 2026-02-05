package nsdhcpparams

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NsdhcpparamsResourceModel describes the resource data model.
type NsdhcpparamsResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Dhcpclient types.String `tfsdk:"dhcpclient"`
	Saveroute  types.String `tfsdk:"saveroute"`
}

func (r *NsdhcpparamsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsdhcpparams resource.",
			},
			"dhcpclient": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enables DHCP client to acquire IP address from the DHCP server in the next boot. When set to OFF, disables the DHCP client in the next boot.",
			},
			"saveroute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "DHCP acquired routes are saved on the Citrix ADC.",
			},
		},
	}
}

func nsdhcpparamsGetThePayloadFromtheConfig(ctx context.Context, data *NsdhcpparamsResourceModel) ns.Nsdhcpparams {
	tflog.Debug(ctx, "In nsdhcpparamsGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nsdhcpparams := ns.Nsdhcpparams{}
	if !data.Dhcpclient.IsNull() {
		nsdhcpparams.Dhcpclient = data.Dhcpclient.ValueString()
	}
	if !data.Saveroute.IsNull() {
		nsdhcpparams.Saveroute = data.Saveroute.ValueString()
	}

	return nsdhcpparams
}

func nsdhcpparamsSetAttrFromGet(ctx context.Context, data *NsdhcpparamsResourceModel, getResponseData map[string]interface{}) *NsdhcpparamsResourceModel {
	tflog.Debug(ctx, "In nsdhcpparamsSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["dhcpclient"]; ok && val != nil {
		data.Dhcpclient = types.StringValue(val.(string))
	} else {
		data.Dhcpclient = types.StringNull()
	}
	if val, ok := getResponseData["saveroute"]; ok && val != nil {
		data.Saveroute = types.StringValue(val.(string))
	} else {
		data.Saveroute = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("nsdhcpparams-config")

	return data
}
