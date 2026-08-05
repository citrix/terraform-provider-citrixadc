package bridgetable

import (
	"context"
	"fmt"
	"strconv"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*BridgetableDataSource)(nil)

func BRidgetableDataSource() datasource.DataSource {
	return &BridgetableDataSource{}
}

type BridgetableDataSource struct {
	client *service.NitroClient
}

func (d *BridgetableDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bridgetable"
}

func (d *BridgetableDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *BridgetableDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = BridgetableDataSourceSchema()
}

func (d *BridgetableDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data BridgetableResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// bridgetable is a collection; fetch all entries and select the one matching
	// any supplied identity keys (mac/vxlan/vtep). With no keys supplied, the first
	// entry is returned.
	findParams := service.FindParams{
		ResourceType: service.Bridgetable.Type(),
	}
	dataArray, err := d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read bridgetable, got error: %s", err))
		return
	}
	if len(dataArray) == 0 {
		resp.Diagnostics.AddError("Client Error", "No bridgetable entries found")
		return
	}

	foundIndex := -1
	for i, entry := range dataArray {
		match := true
		if !data.Mac.IsNull() && fmt.Sprintf("%v", entry["mac"]) != data.Mac.ValueString() {
			match = false
		}
		if !data.Vxlan.IsNull() && fmt.Sprintf("%v", entry["vxlan"]) != strconv.Itoa(int(data.Vxlan.ValueInt64())) {
			match = false
		}
		if !data.Vtep.IsNull() && fmt.Sprintf("%v", entry["vtep"]) != data.Vtep.ValueString() {
			match = false
		}
		if match {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", "No matching bridgetable entry found")
		return
	}

	bridgetableSetAttrFromGetForDatasource(ctx, &data, dataArray[foundIndex])

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
