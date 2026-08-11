package dnsnameserver

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DnsnameserverResourceModel describes the resource data model.
type DnsnameserverResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Dnsprofilename types.String `tfsdk:"dnsprofilename"`
	Dnsvservername types.String `tfsdk:"dnsvservername"`
	Ip             types.String `tfsdk:"ip"`
	Local          types.Bool   `tfsdk:"local"`
	State          types.String `tfsdk:"state"`
	Type           types.String `tfsdk:"type"`
}

func (r *DnsnameserverResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dnsnameserver resource.",
			},
			"dnsprofilename": schema.StringAttribute{
				// SDK v2: Optional + Computed, updateable (no ForceNew).
				// Default "" (no documented server default; unset removes the
				// profile association). The static default makes the attribute
				// non-sticky on config removal so Update fires and the unset
				// operation can revert it.
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(""),
				Description: "Name of the DNS profile to be associated with the name server",
			},
			"dnsvservername": schema.StringAttribute{
				// SDK v2: Optional + ForceNew (not Computed)
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of a DNS virtual server. Overrides any IP address-based name servers configured on the Citrix ADC.",
			},
			"ip": schema.StringAttribute{
				// SDK v2: Optional + ForceNew (not Computed)
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IP address of an external name server or, if the Local parameter is set, IP address of a local DNS server (LDNS).",
			},
			"local": schema.BoolAttribute{
				// SDK v2: Optional + Computed + ForceNew
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Mark the IP address as one that belongs to a local recursive DNS server on the Citrix ADC. The appliance recursively resolves queries received on an IP address that is marked as being local. For recursive resolution to work, the global DNS parameter, Recursion, must also be set.\n\nIf no name server is marked as being local, the appliance functions as a stub resolver and load balances the name servers.",
			},
			"state": schema.StringAttribute{
				// SDK v2: Optional + Computed + ForceNew (no explicit Default; ADC default ENABLED is returned by GET)
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Administrative state of the name server.",
			},
			"type": schema.StringAttribute{
				// SDK v2: Optional + Computed + ForceNew (no explicit Default; ADC default UDP is returned by GET)
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Protocol used by the name server. UDP_TCP is not valid if the name server is a DNS virtual server configured on the appliance.",
			},
		},
	}
}

func dnsnameserverGetThePayloadFromthePlan(ctx context.Context, data *DnsnameserverResourceModel) dns.Dnsnameserver {
	tflog.Debug(ctx, "In dnsnameserverGetThePayloadFromthePlan Function")

	// Create API request body from the model
	dnsnameserver := dns.Dnsnameserver{}
	if !data.Ip.IsNull() && !data.Ip.IsUnknown() {
		dnsnameserver.Ip = data.Ip.ValueString()
	}
	if !data.Dnsvservername.IsNull() && !data.Dnsvservername.IsUnknown() {
		dnsnameserver.Dnsvservername = data.Dnsvservername.ValueString()
	}
	if !data.Dnsprofilename.IsNull() && !data.Dnsprofilename.IsUnknown() && data.Dnsprofilename.ValueString() != "" {
		dnsnameserver.Dnsprofilename = data.Dnsprofilename.ValueString()
	}
	if !data.Local.IsNull() && !data.Local.IsUnknown() {
		dnsnameserver.Local = data.Local.ValueBool()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		dnsnameserver.State = data.State.ValueString()
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		dnsnameserver.Type = data.Type.ValueString()
	}

	return dnsnameserver
}

// interfaceToBool robustly converts a NITRO JSON value into a Go bool.
func interfaceToBool(val interface{}) bool {
	switch v := val.(type) {
	case bool:
		return v
	case string:
		return v == "true" || v == "True" || v == "TRUE"
	default:
		return false
	}
}

// dnsnameserverSetAttrFromGet maps the NITRO GET response onto the resource model.
// name is the primary identifier (ip or dnsvservername) parsed from the resource ID,
// and dnsType is the configured protocol parsed from the ID (preserves UDP_TCP).
func dnsnameserverSetAttrFromGet(ctx context.Context, data *DnsnameserverResourceModel, getResponseData map[string]interface{}, name string, dnsType string) *DnsnameserverResourceModel {
	tflog.Debug(ctx, "In dnsnameserverSetAttrFromGet Function")

	// dnsprofilename: reflect the empty default when NITRO omits it (e.g. after
	// an unset) so state matches the schema Default "" and no inconsistent-result
	// error is raised.
	if val, ok := getResponseData["dnsprofilename"]; ok && val != nil {
		data.Dnsprofilename = types.StringValue(val.(string))
	} else {
		data.Dnsprofilename = types.StringValue("")
	}

	// Identity attributes: the matched entry has either ip == name or
	// dnsvservername == name. Set exactly the identity attribute that matched
	// and null the other, so the state mirrors the (mutually exclusive) config.
	entryIp, _ := getResponseData["ip"].(string)
	entryVs, _ := getResponseData["dnsvservername"].(string)
	if name != "" && entryIp == name {
		data.Ip = types.StringValue(name)
		data.Dnsvservername = types.StringNull()
	} else if name != "" && entryVs == name {
		data.Dnsvservername = types.StringValue(name)
		data.Ip = types.StringNull()
	} else {
		// Fallback (should not happen for a matched entry): preserve config,
		// only resolving unknowns to null.
		if data.Ip.IsUnknown() {
			data.Ip = types.StringNull()
		}
		if data.Dnsvservername.IsUnknown() {
			data.Dnsvservername = types.StringNull()
		}
	}

	// local: NITRO frequently omits this from GET. Preserve the configured/prior
	// value so a user-set true/false is not clobbered; only resolve unknown -> null.
	if val, ok := getResponseData["local"]; ok && val != nil {
		data.Local = types.BoolValue(interfaceToBool(val))
	} else if data.Local.IsUnknown() {
		data.Local = types.BoolNull()
	}

	// state
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else if data.State.IsUnknown() {
		data.State = types.StringNull()
	}

	// type: use the configured protocol from the ID (preserves UDP_TCP, which the
	// ADC splits into individual UDP/TCP entries) rather than the matched entry type.
	if dnsType != "" {
		data.Type = types.StringValue(dnsType)
	} else if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else if data.Type.IsUnknown() {
		data.Type = types.StringNull()
	}

	return data
}

// dnsnameserverSetAttrFromGetForDatasource maps the NITRO GET response onto the
// model for the datasource: it copies all readable attributes verbatim from the
// server and derives the composite id (name,type) from the response.
func dnsnameserverSetAttrFromGetForDatasource(ctx context.Context, data *DnsnameserverResourceModel, getResponseData map[string]interface{}) *DnsnameserverResourceModel {
	tflog.Debug(ctx, "In dnsnameserverSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["dnsprofilename"]; ok && val != nil {
		data.Dnsprofilename = types.StringValue(val.(string))
	} else {
		data.Dnsprofilename = types.StringNull()
	}
	if val, ok := getResponseData["ip"]; ok && val != nil {
		data.Ip = types.StringValue(val.(string))
	} else {
		data.Ip = types.StringNull()
	}
	if val, ok := getResponseData["dnsvservername"]; ok && val != nil {
		data.Dnsvservername = types.StringValue(val.(string))
	} else {
		data.Dnsvservername = types.StringNull()
	}
	if val, ok := getResponseData["local"]; ok && val != nil {
		data.Local = types.BoolValue(interfaceToBool(val))
	} else {
		data.Local = types.BoolNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}
	typeVal := ""
	if val, ok := getResponseData["type"]; ok && val != nil {
		typeVal = val.(string)
		data.Type = types.StringValue(typeVal)
	} else {
		data.Type = types.StringNull()
	}

	// Derive the composite id in the SDK v2 format: "<name>,<type>".
	name := ""
	if !data.Ip.IsNull() && data.Ip.ValueString() != "" {
		name = data.Ip.ValueString()
	} else if !data.Dnsvservername.IsNull() && data.Dnsvservername.ValueString() != "" {
		name = data.Dnsvservername.ValueString()
	}
	data.Id = types.StringValue(name + "," + typeVal)

	return data
}
