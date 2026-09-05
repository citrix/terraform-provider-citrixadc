package dnssvcbrec

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

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

// unsetOnRemoveStringModifier forces the planned value to unknown when the user
// removes a previously-set attribute from configuration while a non-empty value
// still exists in prior state. This makes Terraform detect a change (unknown !=
// prior) and call Update, which issues the NITRO ?action=unset. Without it an
// Optional+Computed attribute is "sticky": the prior value is carried forward
// and removal is a silent no-op. Marking the plan unknown also avoids a
// "provider produced inconsistent result" error after the unset reverts the
// attribute to no value.
type unsetOnRemoveStringModifier struct{}

func (m unsetOnRemoveStringModifier) Description(_ context.Context) string {
	return "Marks the value unknown when removed from config while a prior non-empty value exists, so it is unset on the appliance."
}

func (m unsetOnRemoveStringModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m unsetOnRemoveStringModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.StateValue.IsNull() {
		return
	}
	if req.ConfigValue.IsNull() && req.StateValue.ValueString() != "" {
		resp.PlanValue = types.StringUnknown()
	}
}

// unsetOnRemoveInt64Modifier is the Int64 counterpart of unsetOnRemoveStringModifier.
// defaultValue is the value the appliance reverts the attribute to after ?action=unset:
// leave it 0 for attributes that revert to no value (omitted from GET), or set it to
// the NITRO default for attributes that revert to a known non-zero value (e.g. ttl ->
// 3600). Comparing against defaultValue prevents a perpetual post-unset "known after
// apply" plan churn once state already equals the default.
type unsetOnRemoveInt64Modifier struct{ defaultValue int64 }

func (m unsetOnRemoveInt64Modifier) Description(_ context.Context) string {
	return "Marks the value unknown when removed from config while prior state differs from the default, so it is unset on the appliance."
}

func (m unsetOnRemoveInt64Modifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m unsetOnRemoveInt64Modifier) PlanModifyInt64(_ context.Context, req planmodifier.Int64Request, resp *planmodifier.Int64Response) {
	if req.StateValue.IsNull() {
		return
	}
	if req.ConfigValue.IsNull() && req.StateValue.ValueInt64() != m.defaultValue {
		resp.PlanValue = types.Int64Unknown()
	}
}

// unsetOnRemoveBoolModifier is the Bool counterpart of unsetOnRemoveStringModifier.
type unsetOnRemoveBoolModifier struct{}

func (m unsetOnRemoveBoolModifier) Description(_ context.Context) string {
	return "Marks the value unknown when removed from config while a prior true value exists, so it is unset on the appliance."
}

func (m unsetOnRemoveBoolModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m unsetOnRemoveBoolModifier) PlanModifyBool(_ context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
	if req.StateValue.IsNull() {
		return
	}
	if req.ConfigValue.IsNull() && req.StateValue.ValueBool() {
		resp.PlanValue = types.BoolUnknown()
	}
}

// DnssvcbrecResourceModel describes the resource data model.
type DnssvcbrecResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Domain               types.String `tfsdk:"domain"`
	Targetname           types.String `tfsdk:"targetname"`
	Priority             types.Int64  `tfsdk:"priority"`
	Svcbtype             types.String `tfsdk:"svcbtype"`
	Alpn                 types.String `tfsdk:"alpn"`
	Encryptedclienthello types.String `tfsdk:"encryptedclienthello"`
	Ipv4hint             types.String `tfsdk:"ipv4hint"`
	Ipv6hint             types.String `tfsdk:"ipv6hint"`
	Mandatory            types.String `tfsdk:"mandatory"`
	Nodefaultalpn        types.Bool   `tfsdk:"nodefaultalpn"`
	Nodeid               types.Int64  `tfsdk:"nodeid"`
	Port                 types.Int64  `tfsdk:"port"`
	Ttl                  types.Int64  `tfsdk:"ttl"`
}

func (r *DnssvcbrecResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dnssvcbrec resource. It is a combination of domain, targetname, priority and svcbtype.",
			},
			"domain": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Domain name for the SVCB/HTTPS record.",
			},
			"targetname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Target domain name.",
			},
			"priority": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Service priority (0 for AliasMode, >0 for ServiceMode).",
			},
			"svcbtype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("SVCB"),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Service type: SVCB or HTTPS. Possible values = SVCB, HTTPS",
			},
			"alpn": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "Comma-separated list of ALPN protocol identifiers.",
			},
			"encryptedclienthello": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "Base64-encoded ECH configuration.",
			},
			"ipv4hint": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "Comma-separated list of IPv4 hint addresses.",
			},
			"ipv6hint": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "Comma-separated list of IPv6 hint addresses.",
			},
			"mandatory": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "Comma-separated list of mandatory SvcParam keys.",
			},
			"nodefaultalpn": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					unsetOnRemoveBoolModifier{},
				},
				Description: "Indicates no default ALPN protocols.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"port": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					unsetOnRemoveInt64Modifier{},
				},
				Description: "Port number for the service.",
			},
			"ttl": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					// ttl reverts to the NITRO default 3600 on unset (a known non-zero
					// value returned by GET), so compare against 3600 to avoid a
					// perpetual "known after apply" churn after the unset.
					unsetOnRemoveInt64Modifier{defaultValue: 3600},
				},
				Description: "Time to Live (TTL) in seconds.",
			},
		},
	}
}

func dnssvcbrecGetThePayloadFromthePlan(ctx context.Context, data *DnssvcbrecResourceModel) dns.Dnssvcbrec {
	tflog.Debug(ctx, "In dnssvcbrecGetThePayloadFromthePlan Function")

	// Create API request body from the model
	dnssvcbrec := dns.Dnssvcbrec{}
	if !data.Domain.IsNull() && !data.Domain.IsUnknown() {
		dnssvcbrec.Domain = data.Domain.ValueString()
	}
	if !data.Targetname.IsNull() && !data.Targetname.IsUnknown() {
		dnssvcbrec.Targetname = data.Targetname.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		dnssvcbrec.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Svcbtype.IsNull() && !data.Svcbtype.IsUnknown() {
		dnssvcbrec.Svcbtype = data.Svcbtype.ValueString()
	}
	if !data.Alpn.IsNull() && !data.Alpn.IsUnknown() {
		dnssvcbrec.Alpn = data.Alpn.ValueString()
	}
	if !data.Encryptedclienthello.IsNull() && !data.Encryptedclienthello.IsUnknown() {
		dnssvcbrec.Encryptedclienthello = data.Encryptedclienthello.ValueString()
	}
	if !data.Ipv4hint.IsNull() && !data.Ipv4hint.IsUnknown() {
		dnssvcbrec.Ipv4hint = data.Ipv4hint.ValueString()
	}
	if !data.Ipv6hint.IsNull() && !data.Ipv6hint.IsUnknown() {
		dnssvcbrec.Ipv6hint = data.Ipv6hint.ValueString()
	}
	if !data.Mandatory.IsNull() && !data.Mandatory.IsUnknown() {
		dnssvcbrec.Mandatory = data.Mandatory.ValueString()
	}
	if !data.Nodefaultalpn.IsNull() && !data.Nodefaultalpn.IsUnknown() {
		dnssvcbrec.Nodefaultalpn = data.Nodefaultalpn.ValueBool()
	}
	if !data.Nodeid.IsNull() && !data.Nodeid.IsUnknown() {
		dnssvcbrec.Nodeid = utils.IntPtr(int(data.Nodeid.ValueInt64()))
	}
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		dnssvcbrec.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Ttl.IsNull() && !data.Ttl.IsUnknown() {
		dnssvcbrec.Ttl = utils.IntPtr(int(data.Ttl.ValueInt64()))
	}

	return dnssvcbrec
}

func dnssvcbrecSetAttrFromGet(ctx context.Context, data *DnssvcbrecResourceModel, getResponseData map[string]interface{}) *DnssvcbrecResourceModel {
	tflog.Debug(ctx, "In dnssvcbrecSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["domain"]; ok && val != nil {
		data.Domain = types.StringValue(val.(string))
	} else if data.Domain.IsUnknown() {
		data.Domain = types.StringNull()
	}
	if val, ok := getResponseData["targetname"]; ok && val != nil {
		data.Targetname = types.StringValue(val.(string))
	} else if data.Targetname.IsUnknown() {
		data.Targetname = types.StringNull()
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	} else if data.Priority.IsUnknown() {
		data.Priority = types.Int64Null()
	}
	if val, ok := getResponseData["svcbtype"]; ok && val != nil {
		data.Svcbtype = types.StringValue(val.(string))
	} else if data.Svcbtype.IsUnknown() {
		data.Svcbtype = types.StringNull()
	}
	if val, ok := getResponseData["alpn"]; ok && val != nil {
		data.Alpn = types.StringValue(val.(string))
	} else if data.Alpn.IsUnknown() {
		// NITRO omits alpn from GET when unset; only null it when it was unknown
		// (never configured or just unset) so a user-supplied value is preserved.
		data.Alpn = types.StringNull()
	}
	if val, ok := getResponseData["encryptedclienthello"]; ok && val != nil {
		data.Encryptedclienthello = types.StringValue(val.(string))
	} else if data.Encryptedclienthello.IsUnknown() {
		data.Encryptedclienthello = types.StringNull()
	}
	if val, ok := getResponseData["ipv4hint"]; ok && val != nil {
		data.Ipv4hint = types.StringValue(val.(string))
	} else if data.Ipv4hint.IsUnknown() {
		data.Ipv4hint = types.StringNull()
	}
	if val, ok := getResponseData["ipv6hint"]; ok && val != nil {
		data.Ipv6hint = types.StringValue(val.(string))
	} else if data.Ipv6hint.IsUnknown() {
		data.Ipv6hint = types.StringNull()
	}
	if val, ok := getResponseData["mandatory"]; ok && val != nil {
		data.Mandatory = types.StringValue(val.(string))
	} else if data.Mandatory.IsUnknown() {
		data.Mandatory = types.StringNull()
	}
	if val, ok := getResponseData["nodefaultalpn"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.Nodefaultalpn = types.BoolValue(b)
		} else if s, ok := val.(string); ok {
			data.Nodefaultalpn = types.BoolValue(s == "true" || s == "True")
		}
	} else if data.Nodefaultalpn.IsUnknown() {
		// NITRO omits nodefaultalpn (default false) from GET; preserve a set value.
		data.Nodefaultalpn = types.BoolNull()
	}
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		}
	} else if data.Nodeid.IsUnknown() {
		data.Nodeid = types.Int64Null()
	}
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	} else if data.Port.IsUnknown() {
		data.Port = types.Int64Null()
	}
	if val, ok := getResponseData["ttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ttl = types.Int64Value(intVal)
		}
	} else if data.Ttl.IsUnknown() {
		data.Ttl = types.Int64Null()
	}

	// Set ID for the resource - composite "domain,targetname,priority,svcbtype"
	// mirrors the record identity used by the NITRO delete operation.
	data.Id = types.StringValue(fmt.Sprintf("%s,%s,%d,%s",
		data.Domain.ValueString(),
		data.Targetname.ValueString(),
		data.Priority.ValueInt64(),
		data.Svcbtype.ValueString()))

	return data
}
