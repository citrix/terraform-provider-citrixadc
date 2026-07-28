package fis_channel_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*FisChannelBindingDataSource)(nil)

func FIsChannelBindingDataSource() datasource.DataSource {
	return &FisChannelBindingDataSource{}
}

type FisChannelBindingDataSource struct {
	client *service.NitroClient
}

func (d *FisChannelBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fis_channel_binding"
}

func (d *FisChannelBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *FisChannelBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = FisChannelBindingDataSourceSchema()
}

func (d *FisChannelBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data FisChannelBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Array filter with parent ID - the bound channels are only retrievable via the
	// aggregate parent endpoint (fis_binding/<name>).
	name_Name := data.Name.ValueString()
	ifnum_Name := data.Ifnum

	dataArr, err := fis_channel_bindingAggregateRead(d.client, name_Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read fis_channel_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "fis_channel_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right ifnum
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		if val, ok := v["ifnum"].(string); ok {
			if ifnum_Name.IsNull() || val != ifnum_Name.ValueString() {
				match = false
				continue
			}
		} else if !ifnum_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("fis_channel_binding with ifnum %s not found", ifnum_Name))
		return
	}

	fis_channel_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
