package dnssoarec

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// DnssoarecResourceModel describes the resource data model.
type DnssoarecResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Contact      types.String `tfsdk:"contact"`
	Domain       types.String `tfsdk:"domain"`
	Ecssubnet    types.String `tfsdk:"ecssubnet"`
	Expire       types.Int64  `tfsdk:"expire"`
	Minimum      types.Int64  `tfsdk:"minimum"`
	Nodeid       types.Int64  `tfsdk:"nodeid"`
	Originserver types.String `tfsdk:"originserver"`
	Refresh      types.Int64  `tfsdk:"refresh"`
	Retry        types.Int64  `tfsdk:"retry"`
	Serial       types.Int64  `tfsdk:"serial"`
	Ttl          types.Int64  `tfsdk:"ttl"`
}

func (r *DnssoarecResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dnssoarec resource.",
			},
			"contact": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Email address of the contact to whom domain issues can be addressed. In the email address, replace the @ sign with a period (.). For example, enter domainadmin.example.com instead of domainadmin@example.com.",
			},
			"domain": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Domain name for which to add the SOA record.",
			},
			"ecssubnet": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subnet for which the cached SOA record need to be removed.",
			},
			"expire": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, after which the zone data on a secondary name server can no longer be considered authoritative because all refresh and retry attempts made during the period have failed. After the expiry period, the secondary server stops serving the zone. Typically one week. Not used by the primary server.",
			},
			"minimum": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Default time to live (TTL) for all records in the zone. Can be overridden for individual records.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"originserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Domain name of the name server that responds authoritatively for the domain.",
			},
			"refresh": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, for which a secondary server must wait between successive checks on the value of the serial number.",
			},
			"retry": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, between retries if a secondary server's attempt to contact the primary server for a zone refresh fails.",
			},
			"serial": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The secondary server uses this parameter to determine whether it requires a zone transfer from the primary server.",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.",
			},
		},
	}
}

func dnssoarecGetThePayloadFromthePlan(ctx context.Context, data *DnssoarecResourceModel) dns.Dnssoarec {
	tflog.Debug(ctx, "In dnssoarecGetThePayloadFromthePlan Function")

	// Create API request body from the model
	dnssoarec := dns.Dnssoarec{}
	if !data.Contact.IsNull() && !data.Contact.IsUnknown() {
		dnssoarec.Contact = data.Contact.ValueString()
	}
	if !data.Domain.IsNull() && !data.Domain.IsUnknown() {
		dnssoarec.Domain = data.Domain.ValueString()
	}
	if !data.Ecssubnet.IsNull() && !data.Ecssubnet.IsUnknown() {
		dnssoarec.Ecssubnet = data.Ecssubnet.ValueString()
	}
	if !data.Expire.IsNull() && !data.Expire.IsUnknown() {
		dnssoarec.Expire = utils.IntPtr(int(data.Expire.ValueInt64()))
	}
	if !data.Minimum.IsNull() && !data.Minimum.IsUnknown() {
		dnssoarec.Minimum = utils.IntPtr(int(data.Minimum.ValueInt64()))
	}
	if !data.Nodeid.IsNull() && !data.Nodeid.IsUnknown() {
		dnssoarec.Nodeid = utils.IntPtr(int(data.Nodeid.ValueInt64()))
	}
	if !data.Originserver.IsNull() && !data.Originserver.IsUnknown() {
		dnssoarec.Originserver = data.Originserver.ValueString()
	}
	if !data.Refresh.IsNull() && !data.Refresh.IsUnknown() {
		dnssoarec.Refresh = utils.IntPtr(int(data.Refresh.ValueInt64()))
	}
	if !data.Retry.IsNull() && !data.Retry.IsUnknown() {
		dnssoarec.Retry = utils.IntPtr(int(data.Retry.ValueInt64()))
	}
	if !data.Serial.IsNull() && !data.Serial.IsUnknown() {
		dnssoarec.Serial = utils.IntPtr(int(data.Serial.ValueInt64()))
	}
	if !data.Ttl.IsNull() && !data.Ttl.IsUnknown() {
		dnssoarec.Ttl = utils.IntPtr(int(data.Ttl.ValueInt64()))
	}
	return dnssoarec
}

func dnssoarecSetAttrFromGet(ctx context.Context, data *DnssoarecResourceModel, getResponseData map[string]interface{}) *DnssoarecResourceModel {
	tflog.Debug(ctx, "In dnssoarecSetAttrFromGet Function")

	// Convert API response to model.
	// For Optional+Computed attributes whose value may be omitted from the GET
	// response (e.g. NITRO omits an attribute equal to its default), only reset to
	// null when the model value was unknown. This preserves a user-configured value
	// and prevents "provider produced inconsistent result after apply" errors.
	if val, ok := getResponseData["contact"]; ok && val != nil {
		data.Contact = types.StringValue(val.(string))
	} else if data.Contact.IsUnknown() {
		data.Contact = types.StringNull()
	}
	// domain is the primary key and is always returned by NITRO.
	if val, ok := getResponseData["domain"]; ok && val != nil {
		data.Domain = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["ecssubnet"]; ok && val != nil {
		data.Ecssubnet = types.StringValue(val.(string))
	} else if data.Ecssubnet.IsUnknown() {
		data.Ecssubnet = types.StringNull()
	}
	if val, ok := getResponseData["expire"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Expire = types.Int64Value(intVal)
		}
	} else if data.Expire.IsUnknown() {
		data.Expire = types.Int64Null()
	}
	if val, ok := getResponseData["minimum"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Minimum = types.Int64Value(intVal)
		}
	} else if data.Minimum.IsUnknown() {
		data.Minimum = types.Int64Null()
	}
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		}
	} else if data.Nodeid.IsUnknown() {
		data.Nodeid = types.Int64Null()
	}
	if val, ok := getResponseData["originserver"]; ok && val != nil {
		data.Originserver = types.StringValue(val.(string))
	} else if data.Originserver.IsUnknown() {
		data.Originserver = types.StringNull()
	}
	if val, ok := getResponseData["refresh"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Refresh = types.Int64Value(intVal)
		}
	} else if data.Refresh.IsUnknown() {
		data.Refresh = types.Int64Null()
	}
	if val, ok := getResponseData["retry"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Retry = types.Int64Value(intVal)
		}
	} else if data.Retry.IsUnknown() {
		data.Retry = types.Int64Null()
	}
	if val, ok := getResponseData["serial"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Serial = types.Int64Value(intVal)
		}
	} else if data.Serial.IsUnknown() {
		data.Serial = types.Int64Null()
	}
	if val, ok := getResponseData["ttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ttl = types.Int64Value(intVal)
		}
	} else if data.Ttl.IsUnknown() {
		data.Ttl = types.Int64Null()
	}

	// Set ID for the resource
	// Named resource keyed on the single unique attribute (domain) - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Domain.ValueString()))
	return data
}

// dnssoarecSetAttrFromGetForDatasource populates all model fields from the GET
// response for the datasource. It mirrors dnssoarecSetAttrFromGet but always
// resets omitted attributes to null (a datasource has no prior/planned state to
// preserve) and sets the datasource ID.
func dnssoarecSetAttrFromGetForDatasource(ctx context.Context, data *DnssoarecResourceModel, getResponseData map[string]interface{}) *DnssoarecResourceModel {
	tflog.Debug(ctx, "In dnssoarecSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["contact"]; ok && val != nil {
		data.Contact = types.StringValue(val.(string))
	} else {
		data.Contact = types.StringNull()
	}
	if val, ok := getResponseData["domain"]; ok && val != nil {
		data.Domain = types.StringValue(val.(string))
	} else {
		data.Domain = types.StringNull()
	}
	if val, ok := getResponseData["ecssubnet"]; ok && val != nil {
		data.Ecssubnet = types.StringValue(val.(string))
	} else {
		data.Ecssubnet = types.StringNull()
	}
	if val, ok := getResponseData["expire"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Expire = types.Int64Value(intVal)
		} else {
			data.Expire = types.Int64Null()
		}
	} else {
		data.Expire = types.Int64Null()
	}
	if val, ok := getResponseData["minimum"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Minimum = types.Int64Value(intVal)
		} else {
			data.Minimum = types.Int64Null()
		}
	} else {
		data.Minimum = types.Int64Null()
	}
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		} else {
			data.Nodeid = types.Int64Null()
		}
	} else {
		data.Nodeid = types.Int64Null()
	}
	if val, ok := getResponseData["originserver"]; ok && val != nil {
		data.Originserver = types.StringValue(val.(string))
	} else {
		data.Originserver = types.StringNull()
	}
	if val, ok := getResponseData["refresh"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Refresh = types.Int64Value(intVal)
		} else {
			data.Refresh = types.Int64Null()
		}
	} else {
		data.Refresh = types.Int64Null()
	}
	if val, ok := getResponseData["retry"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Retry = types.Int64Value(intVal)
		} else {
			data.Retry = types.Int64Null()
		}
	} else {
		data.Retry = types.Int64Null()
	}
	if val, ok := getResponseData["serial"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Serial = types.Int64Value(intVal)
		} else {
			data.Serial = types.Int64Null()
		}
	} else {
		data.Serial = types.Int64Null()
	}
	if val, ok := getResponseData["ttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ttl = types.Int64Value(intVal)
		} else {
			data.Ttl = types.Int64Null()
		}
	} else {
		data.Ttl = types.Int64Null()
	}

	// Set ID for the datasource - plain domain value
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Domain.ValueString()))
	return data
}
