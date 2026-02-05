package dnsmxrec

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
				Required:    true,
				Description: "Domain name for which to add the MX record.",
			},
			"ecssubnet": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Subnet for which the cached MX record need to be removed.",
			},
			"mx": schema.StringAttribute{
				Required:    true,
				Description: "Host name of the mail exchange server.",
			},
			"nodeid": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Unique number that identifies the cluster node.",
			},
			"pref": schema.Int64Attribute{
				Required:    true,
				Description: "Priority number to assign to the mail exchange server. A domain name can have multiple mail servers, with a priority number assigned to each server. The lower the priority number, the higher the mail server's priority. When other mail servers have to deliver mail to the specified domain, they begin with the mail server with the lowest priority number, and use other configured mail servers, in priority order, as backups.",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(3600),
				Description: "Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.",
			},
		},
	}
}

func dnsmxrecGetThePayloadFromtheConfig(ctx context.Context, data *DnsmxrecResourceModel) dns.Dnsmxrec {
	tflog.Debug(ctx, "In dnsmxrecGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	dnsmxrec := dns.Dnsmxrec{}
	if !data.Domain.IsNull() {
		dnsmxrec.Domain = data.Domain.ValueString()
	}
	if !data.Ecssubnet.IsNull() {
		dnsmxrec.Ecssubnet = data.Ecssubnet.ValueString()
	}
	if !data.Mx.IsNull() {
		dnsmxrec.Mx = data.Mx.ValueString()
	}
	if !data.Nodeid.IsNull() {
		dnsmxrec.Nodeid = utils.IntPtr(int(data.Nodeid.ValueInt64()))
	}
	if !data.Pref.IsNull() {
		dnsmxrec.Pref = utils.IntPtr(int(data.Pref.ValueInt64()))
	}
	if !data.Ttl.IsNull() {
		dnsmxrec.Ttl = utils.IntPtr(int(data.Ttl.ValueInt64()))
	}

	return dnsmxrec
}

func dnsmxrecSetAttrFromGet(ctx context.Context, data *DnsmxrecResourceModel, getResponseData map[string]interface{}) *DnsmxrecResourceModel {
	tflog.Debug(ctx, "In dnsmxrecSetAttrFromGet Function")

	// Convert API response to model
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
	if val, ok := getResponseData["mx"]; ok && val != nil {
		data.Mx = types.StringValue(val.(string))
	} else {
		data.Mx = types.StringNull()
	}
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		}
	} else {
		data.Nodeid = types.Int64Null()
	}
	if val, ok := getResponseData["pref"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Pref = types.Int64Value(intVal)
		}
	} else {
		data.Pref = types.Int64Null()
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
