package lsnsipalgprofile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*LsnsipalgprofileDataSource)(nil)

func LSnsipalgprofileDataSource() datasource.DataSource {
	return &LsnsipalgprofileDataSource{}
}

type LsnsipalgprofileDataSource struct {
	client *service.NitroClient
}

func (d *LsnsipalgprofileDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnsipalgprofile"
}

func (d *LsnsipalgprofileDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *LsnsipalgprofileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = LsnsipalgprofileDataSourceSchema()
}

func (d *LsnsipalgprofileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data LsnsipalgprofileResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Case 2: Find with single ID attribute
	sipalgprofilename_Name := data.Sipalgprofilename.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = d.client.FindResource(service.Lsnsipalgprofile.Type(), sipalgprofilename_Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read lsnsipalgprofile, got error: %s", err))
		return
	}

	lsnsipalgprofileSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
