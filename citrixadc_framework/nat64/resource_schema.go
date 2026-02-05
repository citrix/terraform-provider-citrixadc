package nat64

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Nat64ResourceModel describes the resource data model.
type Nat64ResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Acl6name   types.String `tfsdk:"acl6name"`
	Name       types.String `tfsdk:"name"`
	Netprofile types.String `tfsdk:"netprofile"`
}

func (r *Nat64Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nat64 resource.",
			},
			"acl6name": schema.StringAttribute{
				Required:    true,
				Description: "Name of any configured ACL6 whose action is ALLOW.  IPv6 Packets matching the condition of this ACL6 rule and destination IP address of these packets matching the NAT64 IPv6 prefix are considered for NAT64 translation.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the NAT64 rule. Must begin with a letter, number, or the underscore character (_), and can consist of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore characters. Cannot be changed after the rule is created. Choose a name that helps identify the NAT64 rule.",
			},
			"netprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the configured netprofile. The Citrix ADC selects one of the IP address in the netprofile as the source IP address of the translated IPv4 packet to be sent to the IPv4 server.",
			},
		},
	}
}

func nat64GetThePayloadFromtheConfig(ctx context.Context, data *Nat64ResourceModel) network.Nat64 {
	tflog.Debug(ctx, "In nat64GetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nat64 := network.Nat64{}
	if !data.Acl6name.IsNull() {
		nat64.Acl6name = data.Acl6name.ValueString()
	}
	if !data.Name.IsNull() {
		nat64.Name = data.Name.ValueString()
	}
	if !data.Netprofile.IsNull() {
		nat64.Netprofile = data.Netprofile.ValueString()
	}

	return nat64
}

func nat64SetAttrFromGet(ctx context.Context, data *Nat64ResourceModel, getResponseData map[string]interface{}) *Nat64ResourceModel {
	tflog.Debug(ctx, "In nat64SetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["acl6name"]; ok && val != nil {
		data.Acl6name = types.StringValue(val.(string))
	} else {
		data.Acl6name = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["netprofile"]; ok && val != nil {
		data.Netprofile = types.StringValue(val.(string))
	} else {
		data.Netprofile = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
