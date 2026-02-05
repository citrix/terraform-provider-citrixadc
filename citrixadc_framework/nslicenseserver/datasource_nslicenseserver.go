package nslicenseserver

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*NslicenseserverDataSource)(nil)

func NSlicenseserverDataSource() datasource.DataSource {
	return &NslicenseserverDataSource{}
}

type NslicenseserverDataSource struct {
	client *service.NitroClient
}

func (d *NslicenseserverDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nslicenseserver"
}

func (d *NslicenseserverDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *NslicenseserverDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = NslicenseserverDataSourceSchema()
}

func (d *NslicenseserverDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data NslicenseserverResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Validate that at least one identifier is provided
	hasLicenseserverip := !data.Licenseserverip.IsNull() && !data.Licenseserverip.IsUnknown()
	hasServername := !data.Servername.IsNull() && !data.Servername.IsUnknown()

	if !hasLicenseserverip && !hasServername {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"At least one of 'licenseserverip' or 'servername' must be specified to identify the nslicenseserver resource.",
		)
		return
	}

	// Case 3: Array filter without parent ID

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             "nslicenseserver",
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nslicenseserver, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "nslicenseserver returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Only check licenseserverip if it was provided
		if hasLicenseserverip {
			if v["licenseserverip"].(string) != data.Licenseserverip.ValueString() {
				match = false
			}
		}

		// Only check servername if it was provided
		if hasServername {
			if v["servername"].(string) != data.Servername.ValueString() {
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
		resp.Diagnostics.AddError("Client Error", "nslicenseserver with specified criteria not found")
		return
	}

	nslicenseserverSetAttrFromGet(ctx, &data, dataArr[foundIndex])

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
