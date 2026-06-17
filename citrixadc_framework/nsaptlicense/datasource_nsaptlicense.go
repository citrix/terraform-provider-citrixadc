package nsaptlicense

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*NsaptlicenseDataSource)(nil)

func NSaptlicenseDataSource() datasource.DataSource {
	return &NsaptlicenseDataSource{}
}

type NsaptlicenseDataSource struct {
	client *service.NitroClient
}

func (d *NsaptlicenseDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsaptlicense"
}

func (d *NsaptlicenseDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *NsaptlicenseDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = NsaptlicenseDataSourceSchema()
}

func (d *NsaptlicenseDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data NsaptlicenseResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// serialno is the GET-only filter key.
	serialno_Name := data.Serialno

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Nsaptlicense.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nsaptlicense, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "nsaptlicense returned empty array")
		return
	}

	// Iterate through results to find the one with the right serialno
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["serialno"].(string); ok {
			if !serialno_Name.IsNull() && val == serialno_Name.ValueString() {
				foundIndex = i
				break
			}
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("nsaptlicense with serialno %s not found", serialno_Name))
		return
	}

	nsaptlicenseSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
