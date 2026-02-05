package dnsnsrec

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*DnsnsrecDataSource)(nil)

func DNsnsrecDataSource() datasource.DataSource {
	return &DnsnsrecDataSource{}
}

type DnsnsrecDataSource struct {
	client *service.NitroClient
}

func (d *DnsnsrecDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnsnsrec"
}

func (d *DnsnsrecDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *DnsnsrecDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DnsnsrecDataSourceSchema()
}

func (d *DnsnsrecDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DnsnsrecResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read DNS NS record by type
	domain_Name := data.Domain.ValueString()

	var dataArr []map[string]interface{}
	var err error

	// Query using type parameter
	findParams := service.FindParams{
		ResourceType:             service.Dnsnsrec.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read dnsnsrec, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "dnsnsrec returned empty array.")
		return
	}

	// Filter by domain from the results
	foundIndex := -1
	for i, v := range dataArr {
		if domainVal, ok := v["domain"]; ok && domainVal != nil {
			if domainVal.(string) == domain_Name {
				foundIndex = i
				break
			}
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("dnsnsrec with domain %s not found in results", domain_Name))
		return
	}

	dnsnsrecSetAttrFromGet(ctx, &data, dataArr[foundIndex])

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
