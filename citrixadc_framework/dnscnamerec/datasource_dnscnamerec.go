package dnscnamerec

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*DnscnamerecDataSource)(nil)

func DNscnamerecDataSource() datasource.DataSource {
	return &DnscnamerecDataSource{}
}

type DnscnamerecDataSource struct {
	client *service.NitroClient
}

func (d *DnscnamerecDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnscnamerec"
}

func (d *DnscnamerecDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *DnscnamerecDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DnscnamerecDataSourceSchema()
}

func (d *DnscnamerecDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DnscnamerecResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	aliasname_Name := data.Aliasname.ValueString()

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Dnscnamerec.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read dnscnamerec, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "dnscnamerec returned empty array.")
		return
	}

	// Iterate through results to find the one with the right aliasname
	foundIndex := -1
	for i, v := range dataArr {
		if aliasnameVal, ok := v["aliasname"].(string); ok && aliasnameVal == aliasname_Name {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("dnscnamerec with aliasname %s not found", aliasname_Name))
		return
	}

	dnscnamerecSetAttrFromGet(ctx, &data, dataArr[foundIndex])

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
