package dnscnamerec

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

// DnscnamerecResourceModel describes the resource data model.
type DnscnamerecResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Aliasname     types.String `tfsdk:"aliasname"`
	Canonicalname types.String `tfsdk:"canonicalname"`
	Ecssubnet     types.String `tfsdk:"ecssubnet"`
	Nodeid        types.Int64  `tfsdk:"nodeid"`
	Ttl           types.Int64  `tfsdk:"ttl"`
}

func (r *DnscnamerecResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dnscnamerec resource.",
			},
			"aliasname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Alias for the canonical domain name.",
			},
			"canonicalname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Canonical domain name.",
			},
			"ecssubnet": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Subnet for which the cached CNAME record need to be removed.",
			},
			"nodeid": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Unique number that identifies the cluster node.",
			},
			"ttl": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Default:     int64default.StaticInt64(3600),
				Description: "Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.",
			},
		},
	}
}

func dnscnamerecGetThePayloadFromtheConfig(ctx context.Context, data *DnscnamerecResourceModel) dns.Dnscnamerec {
	tflog.Debug(ctx, "In dnscnamerecGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	dnscnamerec := dns.Dnscnamerec{}
	if !data.Aliasname.IsNull() {
		dnscnamerec.Aliasname = data.Aliasname.ValueString()
	}
	if !data.Canonicalname.IsNull() {
		dnscnamerec.Canonicalname = data.Canonicalname.ValueString()
	}
	if !data.Ecssubnet.IsNull() {
		dnscnamerec.Ecssubnet = data.Ecssubnet.ValueString()
	}
	if !data.Nodeid.IsNull() {
		dnscnamerec.Nodeid = utils.IntPtr(int(data.Nodeid.ValueInt64()))
	}
	if !data.Ttl.IsNull() {
		dnscnamerec.Ttl = utils.IntPtr(int(data.Ttl.ValueInt64()))
	}

	return dnscnamerec
}

func dnscnamerecSetAttrFromGet(ctx context.Context, data *DnscnamerecResourceModel, getResponseData map[string]interface{}) *DnscnamerecResourceModel {
	tflog.Debug(ctx, "In dnscnamerecSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["aliasname"]; ok && val != nil {
		data.Aliasname = types.StringValue(val.(string))
	} else {
		data.Aliasname = types.StringNull()
	}
	if val, ok := getResponseData["canonicalname"]; ok && val != nil {
		data.Canonicalname = types.StringValue(val.(string))
	} else {
		data.Canonicalname = types.StringNull()
	}
	if val, ok := getResponseData["ecssubnet"]; ok && val != nil {
		data.Ecssubnet = types.StringValue(val.(string))
	} else {
		data.Ecssubnet = types.StringNull()
	}
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		}
	} else {
		data.Nodeid = types.Int64Null()
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
	data.Id = types.StringValue(fmt.Sprintf("%s", data.Aliasname.ValueString()))

	return data
}
