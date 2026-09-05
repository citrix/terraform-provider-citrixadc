package nsrpcnode

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func NsrpcnodeDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ipaddress": schema.StringAttribute{
				Required:    true,
				Description: "IP address of the node. This has to be in the same subnet as the NSIP address.",
			},
			"password": schema.StringAttribute{
				Sensitive:   true,
				Optional:    true,
				Computed:    true,
				Description: "Password to be used in authentication with the peer system node.",
			},
			"secure": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of the channel when talking to the node.",
			},
			"srcip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Source IP address to be used to communicate with the peer system node. The default value is 0, which means that the appliance uses the NSIP address as the source IP address.",
			},
			"validatecert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "validate the server certificate for secure SSL connections",
			},
		},
	}
}

type NsrpcnodeDataSourceModel struct {
	Id           types.String `tfsdk:"id"`
	Ipaddress    types.String `tfsdk:"ipaddress"`
	Password     types.String `tfsdk:"password"`
	Secure       types.String `tfsdk:"secure"`
	Srcip        types.String `tfsdk:"srcip"`
	Validatecert types.String `tfsdk:"validatecert"`
}

func nsrpcnodeDataSourceSetAttrFromGet(ctx context.Context, data *NsrpcnodeDataSourceModel, getResponseData map[string]interface{}) *NsrpcnodeDataSourceModel {
	tflog.Debug(ctx, "In nsrpcnodeDataSourceSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	} else {
		data.Ipaddress = types.StringNull()
	}
	// password is not returned by NITRO API (secret/ephemeral) - retain from config
	if val, ok := getResponseData["secure"]; ok && val != nil {
		data.Secure = types.StringValue(val.(string))
	} else {
		data.Secure = types.StringNull()
	}
	if val, ok := getResponseData["srcip"]; ok && val != nil {
		data.Srcip = types.StringValue(val.(string))
	} else {
		data.Srcip = types.StringNull()
	}
	if val, ok := getResponseData["validatecert"]; ok && val != nil {
		data.Validatecert = types.StringValue(val.(string))
	} else {
		data.Validatecert = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Ipaddress.ValueString()))

	return data
}
