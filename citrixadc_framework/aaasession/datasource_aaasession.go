package aaasession

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*AaasessionDataSource)(nil)

func AAasessionDataSource() datasource.DataSource {
	return &AaasessionDataSource{}
}

type AaasessionDataSource struct {
	client *service.NitroClient
}

func (d *AaasessionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaasession"
}

func (d *AaasessionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AaasessionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AaasessionDataSourceSchema()
}

// Read backs the datasource with the NITRO get(all) endpoint. All selectors are
// optional filters; the first session matching every supplied filter is
// returned. nodeid is a valid GET filter and is honored here.
func (d *AaasessionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AaasessionResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Aaasession.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read aaasession, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "aaasession returned empty array")
		return
	}

	// Iterate through results to find the first one matching every supplied filter.
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		if !data.Groupname.IsNull() {
			if val, ok := v["groupname"].(string); !ok || val != data.Groupname.ValueString() {
				match = false
			}
		}
		if match && !data.Iip.IsNull() {
			if val, ok := v["iip"].(string); !ok || val != data.Iip.ValueString() {
				match = false
			}
		}
		if match && !data.Netmask.IsNull() {
			if val, ok := v["netmask"].(string); !ok || val != data.Netmask.ValueString() {
				match = false
			}
		}
		if match && !data.Sessionkey.IsNull() {
			if val, ok := v["sessionkey"].(string); !ok || val != data.Sessionkey.ValueString() {
				match = false
			}
		}
		if match && !data.Username.IsNull() {
			if val, ok := v["username"].(string); !ok || val != data.Username.ValueString() {
				match = false
			}
		}

		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", "no aaasession matched the provided filters")
		return
	}

	aaasessionSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])

	// The datasource has no Create; set the synthetic ID here.
	data.Id = types.StringValue("aaasession-query")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
