package wasmfile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*WasmfileDataSource)(nil)

func NewWasmfileDataSource() datasource.DataSource {
	return &WasmfileDataSource{}
}

type WasmfileDataSource struct {
	client *service.NitroClient
}

func (d *WasmfileDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wasmfile"
}

func (d *WasmfileDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *WasmfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = WasmfileDataSourceSchema()
}

func (d *WasmfileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data WasmfileResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	name := data.Name.ValueString()

	findParams := service.FindParams{
		ResourceType:             wasmfileResourceType,
		FilterMap:                map[string]string{"name": name},
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read wasmfile, got error: %s", err))
		return
	}
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("wasmfile %s not found", name))
		return
	}

	// Match the exact name defensively, fall back to the first entry.
	matched := dataArr[0]
	for _, entry := range dataArr {
		if v, ok := entry["name"]; ok && v != nil && v.(string) == name {
			matched = entry
			break
		}
	}
	wasmfileSetAttrFromGet(ctx, &data, matched)

	// Datasource has no Create step that would seed the ID; set it explicitly.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
