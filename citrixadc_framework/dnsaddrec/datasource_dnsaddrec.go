package dnsaddrec

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*DnsaddrecDataSource)(nil)

func DNsaddrecDataSource() datasource.DataSource {
	return &DnsaddrecDataSource{}
}

type DnsaddrecDataSource struct {
	client *service.NitroClient
}

func (d *DnsaddrecDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnsaddrec"
}

func (d *DnsaddrecDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *DnsaddrecDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DnsaddrecDataSourceSchema()
}

func (d *DnsaddrecDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DnsaddrecResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	hostname_Name := data.Hostname.ValueString()

	ipaddress_Name := data.Ipaddress.ValueString()

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Dnsaddrec.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read dnsaddrec, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "dnsaddrec returned empty array.")
		return
	}

	// Iterate through results to find the one with the right hostname and ipaddress
	foundIndex := -1
	for i, v := range dataArr {
		hostnameMatch := false
		ipaddressMatch := false

		if hostnameVal, ok := v["hostname"].(string); ok && hostnameVal == hostname_Name {
			hostnameMatch = true
		}

		if ipaddressVal, ok := v["ipaddress"].(string); ok && ipaddressVal == ipaddress_Name {
			ipaddressMatch = true
		}

		if hostnameMatch && ipaddressMatch {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("dnsaddrec with hostname %s and ipaddress %s not found", hostname_Name, ipaddress_Name))
		return
	}

	dnsaddrecSetAttrFromGet(ctx, &data, dataArr[foundIndex])

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
