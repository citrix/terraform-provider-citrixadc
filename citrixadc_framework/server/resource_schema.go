package server

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/basic"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// ServerResourceModel describes the resource data model.
type ServerResourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Internal           types.Bool   `tfsdk:"internal"`
	Comment            types.String `tfsdk:"comment"`
	Delay              types.Int64  `tfsdk:"delay"`
	Domain             types.String `tfsdk:"domain"`
	Domainresolvenow   types.Bool   `tfsdk:"domainresolvenow"`
	Domainresolveretry types.Int64  `tfsdk:"domainresolveretry"`
	Graceful           types.String `tfsdk:"graceful"`
	Ipaddress          types.String `tfsdk:"ipaddress"`
	Ipv6address        types.String `tfsdk:"ipv6address"`
	Name               types.String `tfsdk:"name"`
	Newname            types.String `tfsdk:"newname"`
	Querytype          types.String `tfsdk:"querytype"`
	State              types.String `tfsdk:"state"`
	Td                 types.Int64  `tfsdk:"td"`
	Translationip      types.String `tfsdk:"translationip"`
	Translationmask    types.String `tfsdk:"translationmask"`
}

func (r *ServerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the server resource.",
			},
			"internal": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Display names of the servers that have been created for internal use.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any information about the server.",
			},
			"delay": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Time, in seconds, after which all the services configured on the server are disabled.",
			},
			"domain": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Domain name of the server. For a domain based configuration, you must create the server first.",
			},
			"domainresolvenow": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Immediately send a DNS query to resolve the server's domain name.",
			},
			"domainresolveretry": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(5),
				Description: "Time, in seconds, for which the NetScaler must wait, after DNS resolution fails, before sending the next DNS query to resolve the domain name.",
			},
			"graceful": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Shut down gracefully, without accepting any new connections, and disabling each service when all of its connections are closed.",
			},
			"ipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv4 or IPv6 address of the server. If you create an IP address based server, you can specify the name of the server, instead of its IP address, when creating a service. Note: If you do not create a server entry, the server IP address that you enter when you create a service becomes the name of the server.",
			},
			"ipv6address": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Support IPv6 addressing mode. If you configure a server with the IPv6 addressing mode, you cannot use the server in the IPv4 addressing mode.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the server.\nMust begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\nCan be changed after the name is created.",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the server. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},
			"querytype": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("A"),
				Description: "Specify the type of DNS resolution to be done on the configured domain to get the backend services. Valid query types are A, AAAA and SRV with A being the default querytype. The type of DNS resolution done on the domains in SRV records is inherited from ipv6 argument.",
			},
			"state": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Initial state of the server.",
			},
			"td": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"translationip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address used to transform the server's DNS-resolved IP address.",
			},
			"translationmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The netmask of the translation ip",
			},
		},
	}
}

func serverGetThePayloadFromtheConfig(ctx context.Context, data *ServerResourceModel) basic.Server {
	tflog.Debug(ctx, "In serverGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	server := basic.Server{}
	if !data.Internal.IsNull() {
		server.Internal = data.Internal.ValueBool()
	}
	if !data.Comment.IsNull() {
		server.Comment = data.Comment.ValueString()
	}
	if !data.Delay.IsNull() {
		server.Delay = utils.IntPtr(int(data.Delay.ValueInt64()))
	}
	if !data.Domain.IsNull() {
		server.Domain = data.Domain.ValueString()
	}
	if !data.Domainresolvenow.IsNull() {
		server.Domainresolvenow = data.Domainresolvenow.ValueBool()
	}
	if !data.Domainresolveretry.IsNull() {
		server.Domainresolveretry = utils.IntPtr(int(data.Domainresolveretry.ValueInt64()))
	}
	if !data.Graceful.IsNull() {
		server.Graceful = data.Graceful.ValueString()
	}
	if !data.Ipaddress.IsNull() {
		server.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Ipv6address.IsNull() {
		server.Ipv6address = data.Ipv6address.ValueString()
	}
	if !data.Name.IsNull() {
		server.Name = data.Name.ValueString()
	}
	if !data.Newname.IsNull() {
		server.Newname = data.Newname.ValueString()
	}
	if !data.Querytype.IsNull() {
		server.Querytype = data.Querytype.ValueString()
	}
	if !data.State.IsNull() {
		server.State = data.State.ValueString()
	}
	if !data.Td.IsNull() {
		server.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Translationip.IsNull() {
		server.Translationip = data.Translationip.ValueString()
	}
	if !data.Translationmask.IsNull() {
		server.Translationmask = data.Translationmask.ValueString()
	}

	return server
}

func serverSetAttrFromGet(ctx context.Context, data *ServerResourceModel, getResponseData map[string]interface{}) *ServerResourceModel {
	tflog.Debug(ctx, "In serverSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["Internal"]; ok && val != nil {
		data.Internal = types.BoolValue(val.(bool))
	} else {
		data.Internal = types.BoolNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["delay"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Delay = types.Int64Value(intVal)
		}
	} else {
		data.Delay = types.Int64Null()
	}
	if val, ok := getResponseData["domain"]; ok && val != nil {
		data.Domain = types.StringValue(val.(string))
	} else {
		data.Domain = types.StringNull()
	}
	if val, ok := getResponseData["domainresolvenow"]; ok && val != nil {
		data.Domainresolvenow = types.BoolValue(val.(bool))
	} else {
		data.Domainresolvenow = types.BoolNull()
	}
	if val, ok := getResponseData["domainresolveretry"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Domainresolveretry = types.Int64Value(intVal)
		}
	} else {
		data.Domainresolveretry = types.Int64Null()
	}
	if val, ok := getResponseData["graceful"]; ok && val != nil {
		data.Graceful = types.StringValue(val.(string))
	} else {
		data.Graceful = types.StringNull()
	}
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	} else {
		data.Ipaddress = types.StringNull()
	}
	if val, ok := getResponseData["ipv6address"]; ok && val != nil {
		data.Ipv6address = types.StringValue(val.(string))
	} else {
		data.Ipv6address = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}
	if val, ok := getResponseData["querytype"]; ok && val != nil {
		data.Querytype = types.StringValue(val.(string))
	} else {
		data.Querytype = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else {
		data.Td = types.Int64Null()
	}
	if val, ok := getResponseData["translationip"]; ok && val != nil {
		data.Translationip = types.StringValue(val.(string))
	} else {
		data.Translationip = types.StringNull()
	}
	if val, ok := getResponseData["translationmask"]; ok && val != nil {
		data.Translationmask = types.StringValue(val.(string))
	} else {
		data.Translationmask = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated
	data.Id = types.StringValue(fmt.Sprintf("%s", data.Name.ValueString()))

	return data
}
