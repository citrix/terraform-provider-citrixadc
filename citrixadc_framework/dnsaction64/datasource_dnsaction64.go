package dnsaction64

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*Dnsaction64DataSource)(nil)

func DNsaction64DataSource() datasource.DataSource {
	return &Dnsaction64DataSource{}
}

type Dnsaction64DataSource struct {
	client *service.NitroClient
}

func (d *Dnsaction64DataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnsaction64"
}

func (d *Dnsaction64DataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *Dnsaction64DataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = Dnsaction64DataSourceSchema()
}

func (d *Dnsaction64DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data Dnsaction64ResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Case 2: Find with single ID attribute
	actionname_Name := data.Actionname.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = d.client.FindResource(service.Dnsaction64.Type(), actionname_Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read dnsaction64, got error: %s", err))
		return
	}

	dnsaction64SetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
