package lsnhttphdrlogprofile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*LsnhttphdrlogprofileDataSource)(nil)

func LSnhttphdrlogprofileDataSource() datasource.DataSource {
	return &LsnhttphdrlogprofileDataSource{}
}

type LsnhttphdrlogprofileDataSource struct {
	client *service.NitroClient
}

func (d *LsnhttphdrlogprofileDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnhttphdrlogprofile"
}

func (d *LsnhttphdrlogprofileDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *LsnhttphdrlogprofileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = LsnhttphdrlogprofileDataSourceSchema()
}

func (d *LsnhttphdrlogprofileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data LsnhttphdrlogprofileResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Case 2: Find with single ID attribute
	httphdrlogprofilename_Name := data.Httphdrlogprofilename.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = d.client.FindResource(service.Lsnhttphdrlogprofile.Type(), httphdrlogprofilename_Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read lsnhttphdrlogprofile, got error: %s", err))
		return
	}

	lsnhttphdrlogprofileSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
