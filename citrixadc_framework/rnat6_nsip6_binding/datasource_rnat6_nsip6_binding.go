package rnat6_nsip6_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*Rnat6Nsip6BindingDataSource)(nil)

func RNat6Nsip6BindingDataSource() datasource.DataSource {
	return &Rnat6Nsip6BindingDataSource{}
}

type Rnat6Nsip6BindingDataSource struct {
	client *service.NitroClient
}

func (d *Rnat6Nsip6BindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rnat6_nsip6_binding"
}

func (d *Rnat6Nsip6BindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *Rnat6Nsip6BindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = Rnat6Nsip6BindingDataSourceSchema()
}

func (d *Rnat6Nsip6BindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data Rnat6Nsip6BindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	natip6_Name := data.Natip6

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Rnat6_nsip6_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read rnat6_nsip6_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "rnat6_nsip6_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check natip6
		if val, ok := v["natip6"].(string); ok {
			if natip6_Name.IsNull() || val != natip6_Name.ValueString() {
				match = false
				continue
			}
		} else if !natip6_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("rnat6_nsip6_binding with natip6 %s not found", natip6_Name))
		return
	}

	rnat6_nsip6_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
