package ntpserver

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*NtpserverDataSource)(nil)

func NTpserverDataSource() datasource.DataSource {
	return &NtpserverDataSource{}
}

type NtpserverDataSource struct {
	client *service.NitroClient
}

func (d *NtpserverDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ntpserver"
}

func (d *NtpserverDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *NtpserverDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = NtpserverDataSourceSchema()
}

func (d *NtpserverDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data NtpserverResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Validate that at least one identifier is provided
	hasServerip := !data.Serverip.IsNull() && !data.Serverip.IsUnknown()
	hasServername := !data.Servername.IsNull() && !data.Servername.IsUnknown()

	if !hasServerip && !hasServername {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"At least one of 'serverip' or 'servername' must be specified to identify the ntpserver resource.",
		)
		return
	}

	// Case 3: Array filter without parent ID

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             "ntpserver",
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read ntpserver, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "ntpserver returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Only check serverip if it was provided
		if hasServerip {
			serveripVal, okServerip := v["serverip"].(string)
			servernameVal, okServername := v["servername"].(string)

			// Check if the serverip matches either the serverip or servername field in the response
			// NetScaler sometimes stores IP addresses in the servername field
			matchFound := false
			if okServerip && serveripVal == data.Serverip.ValueString() {
				matchFound = true
			}
			if !matchFound && okServername && servernameVal == data.Serverip.ValueString() {
				matchFound = true
			}
			if !matchFound {
				match = false
			}
		}

		// Only check servername if it was provided
		if hasServername {
			servernameVal, okServername := v["servername"].(string)
			serveripVal, okServerip := v["serverip"].(string)

			// Check if the servername matches either the servername or serverip field in the response
			matchFound := false
			if okServername && servernameVal == data.Servername.ValueString() {
				matchFound = true
			}
			if !matchFound && okServerip && serveripVal == data.Servername.ValueString() {
				matchFound = true
			}
			if !matchFound {
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
		resp.Diagnostics.AddError("Client Error", "ntpserver with specified criteria not found")
		return
	}

	ntpserverSetAttrFromGet(ctx, &data, dataArr[foundIndex])

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
