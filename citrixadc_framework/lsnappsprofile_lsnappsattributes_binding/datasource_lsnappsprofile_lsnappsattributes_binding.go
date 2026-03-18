package lsnappsprofile_lsnappsattributes_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*LsnappsprofileLsnappsattributesBindingDataSource)(nil)

func LSnappsprofileLsnappsattributesBindingDataSource() datasource.DataSource {
	return &LsnappsprofileLsnappsattributesBindingDataSource{}
}

type LsnappsprofileLsnappsattributesBindingDataSource struct {
	client *service.NitroClient
}

func (d *LsnappsprofileLsnappsattributesBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnappsprofile_lsnappsattributes_binding"
}

func (d *LsnappsprofileLsnappsattributesBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *LsnappsprofileLsnappsattributesBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = LsnappsprofileLsnappsattributesBindingDataSourceSchema()
}

func (d *LsnappsprofileLsnappsattributesBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data LsnappsprofileLsnappsattributesBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	appsprofilename_Name := data.Appsprofilename.ValueString()
	appsattributesname_Name := data.Appsattributesname

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Lsnappsprofile_lsnappsattributes_binding.Type(),
		ResourceName:             appsprofilename_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read lsnappsprofile_lsnappsattributes_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "lsnappsprofile_lsnappsattributes_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check appsattributesname
		if val, ok := v["appsattributesname"].(string); ok {
			if appsattributesname_Name.IsNull() || val != appsattributesname_Name.ValueString() {
				match = false
				continue
			}
		} else if !appsattributesname_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("lsnappsprofile_lsnappsattributes_binding with appsattributesname %s not found", appsattributesname_Name))
		return
	}

	lsnappsprofile_lsnappsattributes_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
