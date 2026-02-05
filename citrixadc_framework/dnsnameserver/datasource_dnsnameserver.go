package dnsnameserver

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*DnsnameserverDataSource)(nil)

func DNsnameserverDataSource() datasource.DataSource {
	return &DnsnameserverDataSource{}
}

type DnsnameserverDataSource struct {
	client *service.NitroClient
}

func (d *DnsnameserverDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnsnameserver"
}

func (d *DnsnameserverDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *DnsnameserverDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DnsnameserverDataSourceSchema()
}

func (d *DnsnameserverDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DnsnameserverResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Validate that at least one identifier is provided
	hasDnsvservername := !data.Dnsvservername.IsNull() && !data.Dnsvservername.IsUnknown()
	hasIp := !data.Ip.IsNull() && !data.Ip.IsUnknown()

	if !hasDnsvservername && !hasIp {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"At least one of 'dnsvservername' or 'ip' must be specified to identify the dnsnameserver resource.",
		)
		return
	}

	// Case 3: Array filter without parent ID

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             "dnsnameserver",
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read dnsnameserver, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "dnsnameserver returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Only check dnsvservername if it was provided
		if hasDnsvservername {
			if v["dnsvservername"].(string) != data.Dnsvservername.ValueString() {
				match = false
			}
		}

		// Only check ip if it was provided
		if hasIp {
			if v["ip"].(string) != data.Ip.ValueString() {
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
		resp.Diagnostics.AddError("Client Error", "dnsnameserver with specified criteria not found")
		return
	}

	dnsnameserverSetAttrFromGet(ctx, &data, dataArr[foundIndex])

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
