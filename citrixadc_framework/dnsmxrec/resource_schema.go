package dnsmxrec

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// DnsmxrecResourceModel describes the resource data model.
type DnsmxrecResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Domain    types.String `tfsdk:"domain"`
	Ecssubnet types.String `tfsdk:"ecssubnet"`
	Mx        types.String `tfsdk:"mx"`
	Nodeid    types.Int64  `tfsdk:"nodeid"`
	Pref      types.Int64  `tfsdk:"pref"`
	Ttl       types.Int64  `tfsdk:"ttl"`
}

func (r *DnsmxrecResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dnsmxrec resource.",
			},
			"domain": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Domain name for which to add the MX record.",
			},
			"ecssubnet": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subnet for which the cached MX record need to be removed.",
			},
			"mx": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Host name of the mail exchange server.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"pref": schema.Int64Attribute{
				Required:    true,
				Description: "Priority number to assign to the mail exchange server. A domain name can have multiple mail servers, with a priority number assigned to each server. The lower the priority number, the higher the mail server's priority. When other mail servers have to deliver mail to the specified domain, they begin with the mail server with the lowest priority number, and use other configured mail servers, in priority order, as backups.",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(3600),
				Description: "Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.",
			},
		},
	}
}

// dnsmxrecGetThePayloadFromtheConfig builds the create (add) payload.
func dnsmxrecGetThePayloadFromtheConfig(ctx context.Context, data *DnsmxrecResourceModel) dns.Dnsmxrec {
	tflog.Debug(ctx, "In dnsmxrecGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	dnsmxrec := dns.Dnsmxrec{}
	if !data.Domain.IsNull() && !data.Domain.IsUnknown() {
		dnsmxrec.Domain = data.Domain.ValueString()
	}
	if !data.Ecssubnet.IsNull() && !data.Ecssubnet.IsUnknown() {
		dnsmxrec.Ecssubnet = data.Ecssubnet.ValueString()
	}
	if !data.Mx.IsNull() && !data.Mx.IsUnknown() {
		dnsmxrec.Mx = data.Mx.ValueString()
	}
	if !data.Nodeid.IsNull() && !data.Nodeid.IsUnknown() {
		dnsmxrec.Nodeid = utils.IntPtr(int(data.Nodeid.ValueInt64()))
	}
	if !data.Pref.IsNull() && !data.Pref.IsUnknown() {
		dnsmxrec.Pref = utils.IntPtr(int(data.Pref.ValueInt64()))
	}
	if !data.Ttl.IsNull() && !data.Ttl.IsUnknown() {
		dnsmxrec.Ttl = utils.IntPtr(int(data.Ttl.ValueInt64()))
	}

	return dnsmxrec
}

// dnsmxrecGetTheUpdatablePayloadFromThePlan builds the update (put) payload.
// domain and mx are mandatory in the NITRO update payload and are RequiresReplace,
// so they are always included; only the changed updateable fields are added,
// mirroring the SDK v2 update behavior.
func dnsmxrecGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *DnsmxrecResourceModel, state *DnsmxrecResourceModel) (dns.Dnsmxrec, bool) {
	tflog.Debug(ctx, "In dnsmxrecGetTheUpdatablePayloadFromThePlan Function")

	dnsmxrec := dns.Dnsmxrec{
		Domain: data.Domain.ValueString(),
		Mx:     data.Mx.ValueString(),
	}
	hasChange := false
	if !data.Ecssubnet.Equal(state.Ecssubnet) {
		if !data.Ecssubnet.IsNull() && !data.Ecssubnet.IsUnknown() {
			dnsmxrec.Ecssubnet = data.Ecssubnet.ValueString()
		}
		hasChange = true
	}
	if !data.Nodeid.Equal(state.Nodeid) {
		if !data.Nodeid.IsNull() && !data.Nodeid.IsUnknown() {
			dnsmxrec.Nodeid = utils.IntPtr(int(data.Nodeid.ValueInt64()))
		}
		hasChange = true
	}
	if !data.Pref.Equal(state.Pref) {
		if !data.Pref.IsNull() && !data.Pref.IsUnknown() {
			dnsmxrec.Pref = utils.IntPtr(int(data.Pref.ValueInt64()))
		}
		hasChange = true
	}
	if !data.Ttl.Equal(state.Ttl) {
		if !data.Ttl.IsNull() && !data.Ttl.IsUnknown() {
			dnsmxrec.Ttl = utils.IntPtr(int(data.Ttl.ValueInt64()))
		}
		hasChange = true
	}

	return dnsmxrec, hasChange
}

func dnsmxrecSetAttrFromGet(ctx context.Context, data *DnsmxrecResourceModel, getResponseData map[string]interface{}) *DnsmxrecResourceModel {
	tflog.Debug(ctx, "In dnsmxrecSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["domain"]; ok && val != nil {
		data.Domain = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["ecssubnet"]; ok && val != nil {
		data.Ecssubnet = types.StringValue(val.(string))
	} else if data.Ecssubnet.IsUnknown() {
		// ecssubnet is Optional+Computed; NITRO omits it when unset. Preserve a
		// configured value and only resolve an unknown (unconfigured) value to null.
		data.Ecssubnet = types.StringNull()
	}
	if val, ok := getResponseData["mx"]; ok && val != nil {
		data.Mx = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		}
	} else if data.Nodeid.IsUnknown() {
		// nodeid is Optional+Computed; its NITRO default is omitted from GET.
		// Preserve a configured value and only resolve an unknown value to null.
		data.Nodeid = types.Int64Null()
	}
	if val, ok := getResponseData["pref"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Pref = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["ttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ttl = types.Int64Value(intVal)
		}
	} else if data.Ttl.IsUnknown() {
		data.Ttl = types.Int64Null()
	}

	// Set ID for the resource
	// Single unique attribute (domain) - use plain value as ID (matches SDK v2)
	data.Id = types.StringValue(data.Domain.ValueString())

	return data
}
