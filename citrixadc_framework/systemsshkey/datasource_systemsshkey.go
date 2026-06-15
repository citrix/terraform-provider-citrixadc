package systemsshkey

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*SystemsshkeyDataSource)(nil)

func SYstemsshkeyDataSource() datasource.DataSource {
	return &SystemsshkeyDataSource{}
}

type SystemsshkeyDataSource struct {
	client *service.NitroClient
}

func (d *SystemsshkeyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemsshkey"
}

func (d *SystemsshkeyDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *SystemsshkeyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SystemsshkeyDataSourceSchema()
}

func (d *SystemsshkeyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SystemsshkeyResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 3: Array filter without parent ID
	name_Name := data.Name
	sshkeytype_Name := data.Sshkeytype

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Systemsshkey.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read systemsshkey, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "systemsshkey returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check name
		if val, ok := v["name"].(string); ok {
			if name_Name.IsNull() || val != name_Name.ValueString() {
				match = false
				continue
			}
		} else if !name_Name.IsNull() {
			match = false
			continue
		}

		// Check sshkeytype
		if val, ok := v["sshkeytype"].(string); ok {
			if sshkeytype_Name.IsNull() || val != sshkeytype_Name.ValueString() {
				match = false
				continue
			}
		} else if !sshkeytype_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("systemsshkey with name %s not found", name_Name))
		return
	}

	systemsshkeySetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
