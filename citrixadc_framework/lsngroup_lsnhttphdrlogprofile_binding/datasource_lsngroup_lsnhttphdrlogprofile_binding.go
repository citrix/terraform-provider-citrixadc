package lsngroup_lsnhttphdrlogprofile_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*LsngroupLsnhttphdrlogprofileBindingDataSource)(nil)

func LSngroupLsnhttphdrlogprofileBindingDataSource() datasource.DataSource {
	return &LsngroupLsnhttphdrlogprofileBindingDataSource{}
}

type LsngroupLsnhttphdrlogprofileBindingDataSource struct {
	client *service.NitroClient
}

func (d *LsngroupLsnhttphdrlogprofileBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsngroup_lsnhttphdrlogprofile_binding"
}

func (d *LsngroupLsnhttphdrlogprofileBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *LsngroupLsnhttphdrlogprofileBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = LsngroupLsnhttphdrlogprofileBindingDataSourceSchema()
}

func (d *LsngroupLsnhttphdrlogprofileBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data LsngroupLsnhttphdrlogprofileBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	groupname_Name := data.Groupname.ValueString()
	httphdrlogprofilename_Name := data.Httphdrlogprofilename

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Lsngroup_lsnhttphdrlogprofile_binding.Type(),
		ResourceName:             groupname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read lsngroup_lsnhttphdrlogprofile_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "lsngroup_lsnhttphdrlogprofile_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check httphdrlogprofilename
		if val, ok := v["httphdrlogprofilename"].(string); ok {
			if httphdrlogprofilename_Name.IsNull() || val != httphdrlogprofilename_Name.ValueString() {
				match = false
				continue
			}
		} else if !httphdrlogprofilename_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("lsngroup_lsnhttphdrlogprofile_binding with httphdrlogprofilename %s not found", httphdrlogprofilename_Name))
		return
	}

	lsngroup_lsnhttphdrlogprofile_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
