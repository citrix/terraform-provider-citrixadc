package appqoecustomresp

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AppqoecustomrespDataSource)(nil)

func APpqoecustomrespDataSource() datasource.DataSource {
	return &AppqoecustomrespDataSource{}
}

type AppqoecustomrespDataSource struct {
	client *service.NitroClient
}

func (d *AppqoecustomrespDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appqoecustomresp"
}

func (d *AppqoecustomrespDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AppqoecustomrespDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AppqoecustomrespDataSourceSchema()
}

func (d *AppqoecustomrespDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AppqoecustomrespResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Case 3: Array filter without parent ID

	name_Name := data.Name.ValueString()

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             "appqoecustomresp",
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read appqoecustomresp, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "appqoecustomresp returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {

		if v["name"].(string) == name_Name {
			foundIndex = i
			break
		}

	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("appqoecustomresp with name %s not found", name_Name))
		return
	}

	appqoecustomrespSetAttrFromGet(ctx, &data, dataArr[foundIndex])

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
