package nskeymanagerproxy

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*NskeymanagerproxyDataSource)(nil)

func NSkeymanagerproxyDataSource() datasource.DataSource {
	return &NskeymanagerproxyDataSource{}
}

type NskeymanagerproxyDataSource struct {
	client *service.NitroClient
}

func (d *NskeymanagerproxyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nskeymanagerproxy"
}

func (d *NskeymanagerproxyDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *NskeymanagerproxyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = NskeymanagerproxyDataSourceSchema()
}

func (d *NskeymanagerproxyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data NskeymanagerproxyResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Look up by serverip when set, otherwise servername (both x-unique-attr).
	name := data.Serverip.ValueString()
	if name == "" {
		name = data.Servername.ValueString()
	}

	getResponseData, err := d.client.FindResource(service.Nskeymanagerproxy.Type(), name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nskeymanagerproxy, got error: %s", err))
		return
	}

	nskeymanagerproxySetAttrFromGetForDatasource(ctx, &data, getResponseData)
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
