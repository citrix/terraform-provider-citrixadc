package server

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/basic"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
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
			// SDK v2: Optional+Computed, NOT ForceNew.
			"internal": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Display names of the servers that have been created for internal use.",
			},
			// SDK v2: Optional+Computed, NOT ForceNew.
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any information about the server.",
			},
			// SDK v2: Optional+Computed, NOT ForceNew. Used only by the disable action.
			"delay": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, after which all the services configured on the server are disabled.",
			},
			// SDK v2: Optional+Computed+ForceNew -> RequiresReplaceIfConfigured.
			"domain": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Domain name of the server. For a domain based configuration, you must create the server first.",
			},
			// SDK v2: Optional+Computed, NOT ForceNew.
			"domainresolvenow": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Immediately send a DNS query to resolve the server's domain name.",
			},
			// SDK v2: Optional+Computed (no default), NOT ForceNew.
			"domainresolveretry": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, for which the NetScaler must wait, after DNS resolution fails, before sending the next DNS query to resolve the domain name.",
			},
			// SDK v2: Optional+Computed, NOT ForceNew. Used only by the disable action.
			"graceful": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Shut down gracefully, without accepting any new connections, and disabling each service when all of its connections are closed.",
			},
			// SDK v2: Optional+Computed, NOT ForceNew.
			"ipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv4 or IPv6 address of the server. If you create an IP address based server, you can specify the name of the server, instead of its IP address, when creating a service. Note: If you do not create a server entry, the server IP address that you enter when you create a service becomes the name of the server.",
			},
			// SDK v2: Optional+Computed+ForceNew -> RequiresReplaceIfConfigured.
			"ipv6address": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Support IPv6 addressing mode. If you configure a server with the IPv6 addressing mode, you cannot use the server in the IPv4 addressing mode.",
			},
			// SDK v2: Optional+Computed+ForceNew (auto-generated when omitted) ->
			// RequiresReplaceIfConfigured so an auto-generated name never forces replace.
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Name for the server.\nMust begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\nCan be changed after the name is created.",
			},
			// SDK v2 has no newname / rename support (name is ForceNew). Kept as a
			// harmless Optional-only input; excluded from the add payload and never
			// read back so it does not churn.
			"newname": schema.StringAttribute{
				Optional:    true,
				Description: "New name for the server. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},
			// SDK v2: Optional+Computed (no default), NOT ForceNew, updateable.
			"querytype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the type of DNS resolution to be done on the configured domain to get the backend services. Valid query types are A, AAAA and SRV with A being the default querytype. The type of DNS resolution done on the domains in SRV records is inherited from ipv6 argument.",
			},
			// SDK v2: Optional+Computed (no default), NOT ForceNew. Drives the
			// enable/disable action in Update.
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Initial state of the server.",
			},
			// SDK v2: Optional+Computed+ForceNew -> RequiresReplaceIfConfigured.
			"td": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			// SDK v2: Optional+Computed, NOT ForceNew.
			"translationip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address used to transform the server's DNS-resolved IP address.",
			},
			// SDK v2: Optional+Computed, NOT ForceNew.
			"translationmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The netmask of the translation ip",
			},
		},
	}
}

// serverGetThePayloadFromthePlan builds the CREATE payload. It mirrors the SDK v2
// createServerFunc: delay, graceful and newname are excluded (delay/graceful are
// disable-action-only; newname is rename-only and unsupported here).
func serverGetThePayloadFromthePlan(ctx context.Context, data *ServerResourceModel) basic.Server {
	tflog.Debug(ctx, "In serverGetThePayloadFromthePlan Function")

	server := basic.Server{}
	if !data.Internal.IsNull() && !data.Internal.IsUnknown() {
		server.Internal = data.Internal.ValueBool()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		server.Comment = data.Comment.ValueString()
	}
	// delay is disable-action-only; excluded from the add payload (matches SDK v2).
	if !data.Domain.IsNull() && !data.Domain.IsUnknown() {
		server.Domain = data.Domain.ValueString()
	}
	if !data.Domainresolvenow.IsNull() && !data.Domainresolvenow.IsUnknown() {
		server.Domainresolvenow = data.Domainresolvenow.ValueBool()
	}
	if !data.Domainresolveretry.IsNull() && !data.Domainresolveretry.IsUnknown() {
		server.Domainresolveretry = utils.IntPtr(int(data.Domainresolveretry.ValueInt64()))
	}
	// graceful is disable-action-only; excluded from the add payload (matches SDK v2).
	if !data.Ipaddress.IsNull() && !data.Ipaddress.IsUnknown() {
		server.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Ipv6address.IsNull() && !data.Ipv6address.IsUnknown() {
		server.Ipv6address = data.Ipv6address.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		server.Name = data.Name.ValueString()
	}
	// newname is rename-only; excluded from the add payload.
	if !data.Querytype.IsNull() && !data.Querytype.IsUnknown() {
		server.Querytype = data.Querytype.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		server.State = data.State.ValueString()
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		server.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Translationip.IsNull() && !data.Translationip.IsUnknown() {
		server.Translationip = data.Translationip.ValueString()
	}
	if !data.Translationmask.IsNull() && !data.Translationmask.IsUnknown() {
		server.Translationmask = data.Translationmask.ValueString()
	}

	return server
}

// serverSetAttrFromGet is the RESOURCE state setter. It preserves configured/known
// values that NITRO omits from GET (the omit-on-default trap): an absent attribute
// is nulled only when the current model value is Unknown, never when it is a known
// (configured or prior-state) value. delay/graceful are not returned by GET (SDK v2
// never read them) so they are only resolved from Unknown -> Null and otherwise
// preserved. newname is never read back.
func serverSetAttrFromGet(ctx context.Context, data *ServerResourceModel, getResponseData map[string]interface{}) *ServerResourceModel {
	tflog.Debug(ctx, "In serverSetAttrFromGet Function")

	if val, ok := getResponseData["Internal"]; ok && val != nil {
		data.Internal = types.BoolValue(val.(bool))
	} else if data.Internal.IsUnknown() {
		data.Internal = types.BoolNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else if data.Comment.IsUnknown() {
		data.Comment = types.StringNull()
	}
	// delay is not returned by GET; preserve known value, resolve unknown -> null.
	if data.Delay.IsUnknown() {
		data.Delay = types.Int64Null()
	}
	if val, ok := getResponseData["domain"]; ok && val != nil {
		data.Domain = types.StringValue(val.(string))
	} else if data.Domain.IsUnknown() {
		data.Domain = types.StringNull()
	}
	if val, ok := getResponseData["domainresolvenow"]; ok && val != nil {
		data.Domainresolvenow = types.BoolValue(val.(bool))
	} else if data.Domainresolvenow.IsUnknown() {
		data.Domainresolvenow = types.BoolNull()
	}
	if val, ok := getResponseData["domainresolveretry"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Domainresolveretry = types.Int64Value(intVal)
		}
	} else if data.Domainresolveretry.IsUnknown() {
		data.Domainresolveretry = types.Int64Null()
	}
	// graceful is not returned by GET; preserve known value, resolve unknown -> null.
	if data.Graceful.IsUnknown() {
		data.Graceful = types.StringNull()
	}
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	} else if data.Ipaddress.IsUnknown() {
		data.Ipaddress = types.StringNull()
	}
	if val, ok := getResponseData["ipv6address"]; ok && val != nil {
		data.Ipv6address = types.StringValue(val.(string))
	} else if data.Ipv6address.IsUnknown() {
		data.Ipv6address = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else if data.Name.IsUnknown() {
		data.Name = types.StringNull()
	}
	// newname is rename-only and never returned by GET; leave the model value as-is.
	if val, ok := getResponseData["querytype"]; ok && val != nil {
		data.Querytype = types.StringValue(val.(string))
	} else if data.Querytype.IsUnknown() {
		data.Querytype = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else if data.State.IsUnknown() {
		data.State = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else if data.Td.IsUnknown() {
		data.Td = types.Int64Null()
	}
	if val, ok := getResponseData["translationip"]; ok && val != nil {
		data.Translationip = types.StringValue(val.(string))
	} else if data.Translationip.IsUnknown() {
		data.Translationip = types.StringNull()
	}
	if val, ok := getResponseData["translationmask"]; ok && val != nil {
		data.Translationmask = types.StringValue(val.(string))
	} else if data.Translationmask.IsUnknown() {
		data.Translationmask = types.StringNull()
	}

	// ID is the plain server name (SDK v2: d.SetId(serverName)).
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}

// serverSetAttrFromGetForDatasource is the DATASOURCE state setter. It populates
// every attribute directly from the GET response (config carries only the lookup
// key) and always sets the ID.
func serverSetAttrFromGetForDatasource(ctx context.Context, data *ServerResourceModel, getResponseData map[string]interface{}) *ServerResourceModel {
	tflog.Debug(ctx, "In serverSetAttrFromGetForDatasource Function")

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
		} else {
			data.Delay = types.Int64Null()
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
		} else {
			data.Domainresolveretry = types.Int64Null()
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
	data.Newname = types.StringNull()
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
		} else {
			data.Td = types.Int64Null()
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

	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
