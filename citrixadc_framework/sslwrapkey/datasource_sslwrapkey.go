package sslwrapkey

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*SslwrapkeyDataSource)(nil)

func SSlwrapkeyDataSource() datasource.DataSource {
	return &SslwrapkeyDataSource{}
}

type SslwrapkeyDataSource struct {
	client *service.NitroClient
}

func (d *SslwrapkeyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslwrapkey"
}

func (d *SslwrapkeyDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *SslwrapkeyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SslwrapkeyDataSourceSchema()
}

func (d *SslwrapkeyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SslwrapkeyResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 3: Array filter without parent ID
	wrapkeyname_Name := data.Wrapkeyname

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Sslwrapkey.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read sslwrapkey, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "sslwrapkey returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check wrapkeyname
		if val, ok := v["wrapkeyname"].(string); ok {
			if wrapkeyname_Name.IsNull() || val != wrapkeyname_Name.ValueString() {
				match = false
				continue
			}
		} else if !wrapkeyname_Name.IsNull() {
			match = false
			continue
		}

		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("sslwrapkey with wrapkeyname %s not found", wrapkeyname_Name))
		return
	}

	sslwrapkeySetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
