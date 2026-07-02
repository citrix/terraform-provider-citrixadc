package dnsproxyrecords

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DnsproxyrecordsResourceModel describes the resource data model.
type DnsproxyrecordsResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Negrectype types.String `tfsdk:"negrectype"`
	Type       types.String `tfsdk:"type"`
}

func (r *DnsproxyrecordsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dnsproxyrecords resource.",
			},
			"negrectype": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Filter the Negative DNS records i.e NXDOMAIN and NODATA entries to be flushed. e.g flush dns proxyRecords NXDOMAIN will flush only the NXDOMAIN entries from the cache",
			},
			"type": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Filter the DNS records to be flushed.e.g flush dns proxyRecords -type A   will flush only the A records from the cache.",
			},
		},
	}
}

func dnsproxyrecordsGetThePayloadFromthePlan(ctx context.Context, data *DnsproxyrecordsResourceModel) dns.Dnsproxyrecords {
	tflog.Debug(ctx, "In dnsproxyrecordsGetThePayloadFromthePlan Function")

	// Create API request body from the model
	dnsproxyrecords := dns.Dnsproxyrecords{}
	if !data.Negrectype.IsNull() && !data.Negrectype.IsUnknown() {
		dnsproxyrecords.Negrectype = data.Negrectype.ValueString()
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		dnsproxyrecords.Type = data.Type.ValueString()
	}

	return dnsproxyrecords
}
