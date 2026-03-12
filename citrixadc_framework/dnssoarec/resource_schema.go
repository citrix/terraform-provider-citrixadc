package dnssoarec

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
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
				Required:    true,
				Description: "Email address of the contact to whom domain issues can be addressed. In the email address, replace the @ sign with a period (.). For example, enter domainadmin.example.com instead of domainadmin@example.com.",
			},
			"domain": schema.StringAttribute{
				Required:    true,
				Description: "Domain name for which to add the SOA record.",
			},
			"ecssubnet": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Subnet for which the cached SOA record need to be removed.",
			},
			"expire": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(3600),
				Description: "Time, in seconds, after which the zone data on a secondary name server can no longer be considered authoritative because all refresh and retry attempts made during the period have failed. After the expiry period, the secondary server stops serving the zone. Typically one week. Not used by the primary server.",
			},
			"minimum": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(5),
				Description: "Default time to live (TTL) for all records in the zone. Can be overridden for individual records.",
			},
			"nodeid": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Unique number that identifies the cluster node.",
			},
			"originserver": schema.StringAttribute{
				Required:    true,
				Description: "Domain name of the name server that responds authoritatively for the domain.",
			},
			"refresh": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(3600),
				Description: "Time, in seconds, for which a secondary server must wait between successive checks on the value of the serial number.",
			},
			"retry": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(3),
				Description: "Time, in seconds, between retries if a secondary server's attempt to contact the primary server for a zone refresh fails.",
			},
			"serial": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(100),
				Description: "The secondary server uses this parameter to determine whether it requires a zone transfer from the primary server.",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(3600),
				Description: "Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.",
			},
		},
	}
}

func dnssoarecGetThePayloadFromtheConfig(ctx context.Context, data *DnssoarecResourceModel) dns.Dnssoarec {
	tflog.Debug(ctx, "In dnssoarecGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	dnssoarec := dns.Dnssoarec{}
	if !data.Contact.IsNull() {
		dnssoarec.Contact = data.Contact.ValueString()
	}
	if !data.Domain.IsNull() {
		dnssoarec.Domain = data.Domain.ValueString()
	}
	if !data.Ecssubnet.IsNull() {
		dnssoarec.Ecssubnet = data.Ecssubnet.ValueString()
	}
	if !data.Expire.IsNull() {
		dnssoarec.Expire = utils.IntPtr(int(data.Expire.ValueInt64()))
	}
	if !data.Minimum.IsNull() {
		dnssoarec.Minimum = utils.IntPtr(int(data.Minimum.ValueInt64()))
	}
	if !data.Nodeid.IsNull() {
		dnssoarec.Nodeid = utils.IntPtr(int(data.Nodeid.ValueInt64()))
	}
	if !data.Originserver.IsNull() {
		dnssoarec.Originserver = data.Originserver.ValueString()
	}
	if !data.Refresh.IsNull() {
		dnssoarec.Refresh = utils.IntPtr(int(data.Refresh.ValueInt64()))
	}
	if !data.Retry.IsNull() {
		dnssoarec.Retry = utils.IntPtr(int(data.Retry.ValueInt64()))
	}
	if !data.Serial.IsNull() {
		dnssoarec.Serial = utils.IntPtr(int(data.Serial.ValueInt64()))
	}
	if !data.Ttl.IsNull() {
		dnssoarec.Ttl = utils.IntPtr(int(data.Ttl.ValueInt64()))
	}
	return dnssoarec
}

func dnssoarecSetAttrFromGet(ctx context.Context, data *DnssoarecResourceModel, getResponseData map[string]interface{}) *DnssoarecResourceModel {
	tflog.Debug(ctx, "In dnssoarecSetAttrFromGet Function")

	// Convert API response to model
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
		}
	} else {
		data.Expire = types.Int64Null()
	}
	if val, ok := getResponseData["minimum"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Minimum = types.Int64Value(intVal)
		}
	} else {
		data.Minimum = types.Int64Null()
	}
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
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
		}
	} else {
		data.Refresh = types.Int64Null()
	}
	if val, ok := getResponseData["retry"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Retry = types.Int64Value(intVal)
		}
	} else {
		data.Retry = types.Int64Null()
	}
	if val, ok := getResponseData["serial"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Serial = types.Int64Value(intVal)
		}
	} else {
		data.Serial = types.Int64Null()
	}
	if val, ok := getResponseData["ttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ttl = types.Int64Value(intVal)
		}
	} else {
		data.Ttl = types.Int64Null()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated
	data.Id = types.StringValue(fmt.Sprintf("%s", data.Domain.ValueString()))
	return data
}
