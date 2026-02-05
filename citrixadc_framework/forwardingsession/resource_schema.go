package forwardingsession

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// ForwardingsessionResourceModel describes the resource data model.
type ForwardingsessionResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Acl6name         types.String `tfsdk:"acl6name"`
	Aclname          types.String `tfsdk:"aclname"`
	Connfailover     types.String `tfsdk:"connfailover"`
	Name             types.String `tfsdk:"name"`
	Netmask          types.String `tfsdk:"netmask"`
	Network          types.String `tfsdk:"network"`
	Processlocal     types.String `tfsdk:"processlocal"`
	Sourceroutecache types.String `tfsdk:"sourceroutecache"`
	Td               types.Int64  `tfsdk:"td"`
}

func (r *ForwardingsessionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the forwardingsession resource.",
			},
			"acl6name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of any configured ACL6 whose action is ALLOW. The rule of the ACL6 is used as a forwarding session rule.",
			},
			"aclname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of any configured ACL whose action is ALLOW. The rule of the ACL is used as a forwarding session rule.",
			},
			"connfailover": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Synchronize connection information with the secondary appliance in a high availability (HA) pair. That is, synchronize all connection-related information for the forwarding session.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the forwarding session rule. Can begin with a letter, number, or the underscore character (_), and can consist of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the rule is created.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my rule\" or 'my rule').",
			},
			"netmask": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Subnet mask associated with the network.",
			},
			"network": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "An IPv4 network address or IPv6 prefix of a network from which the forwarded traffic originates or to which it is destined.",
			},
			"processlocal": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enabling this option on forwarding session will not steer the packet to flow processor. Instead, packet will be routed.",
			},
			"sourceroutecache": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Cache the source ip address and mac address of the DA servers.",
			},
			"td": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
		},
	}
}

func forwardingsessionGetThePayloadFromtheConfig(ctx context.Context, data *ForwardingsessionResourceModel) network.Forwardingsession {
	tflog.Debug(ctx, "In forwardingsessionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	forwardingsession := network.Forwardingsession{}
	if !data.Acl6name.IsNull() {
		forwardingsession.Acl6name = data.Acl6name.ValueString()
	}
	if !data.Aclname.IsNull() {
		forwardingsession.Aclname = data.Aclname.ValueString()
	}
	if !data.Connfailover.IsNull() {
		forwardingsession.Connfailover = data.Connfailover.ValueString()
	}
	if !data.Name.IsNull() {
		forwardingsession.Name = data.Name.ValueString()
	}
	if !data.Netmask.IsNull() {
		forwardingsession.Netmask = data.Netmask.ValueString()
	}
	if !data.Network.IsNull() {
		forwardingsession.Network = data.Network.ValueString()
	}
	if !data.Processlocal.IsNull() {
		forwardingsession.Processlocal = data.Processlocal.ValueString()
	}
	if !data.Sourceroutecache.IsNull() {
		forwardingsession.Sourceroutecache = data.Sourceroutecache.ValueString()
	}
	if !data.Td.IsNull() {
		forwardingsession.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}

	return forwardingsession
}

func forwardingsessionSetAttrFromGet(ctx context.Context, data *ForwardingsessionResourceModel, getResponseData map[string]interface{}) *ForwardingsessionResourceModel {
	tflog.Debug(ctx, "In forwardingsessionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["acl6name"]; ok && val != nil {
		data.Acl6name = types.StringValue(val.(string))
	} else {
		data.Acl6name = types.StringNull()
	}
	if val, ok := getResponseData["aclname"]; ok && val != nil {
		data.Aclname = types.StringValue(val.(string))
	} else {
		data.Aclname = types.StringNull()
	}
	if val, ok := getResponseData["connfailover"]; ok && val != nil {
		data.Connfailover = types.StringValue(val.(string))
	} else {
		data.Connfailover = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	} else {
		data.Netmask = types.StringNull()
	}
	if val, ok := getResponseData["network"]; ok && val != nil {
		data.Network = types.StringValue(val.(string))
	} else {
		data.Network = types.StringNull()
	}
	if val, ok := getResponseData["processlocal"]; ok && val != nil {
		data.Processlocal = types.StringValue(val.(string))
	} else {
		data.Processlocal = types.StringNull()
	}
	if val, ok := getResponseData["sourceroutecache"]; ok && val != nil {
		data.Sourceroutecache = types.StringValue(val.(string))
	} else {
		data.Sourceroutecache = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else {
		data.Td = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
