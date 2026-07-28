package lsngroup_lsnsipalgprofile_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*LsngroupLsnsipalgprofileBindingDataSource)(nil)

func LSngroupLsnsipalgprofileBindingDataSource() datasource.DataSource {
	return &LsngroupLsnsipalgprofileBindingDataSource{}
}

type LsngroupLsnsipalgprofileBindingDataSource struct {
	client *service.NitroClient
}

func (d *LsngroupLsnsipalgprofileBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsngroup_lsnsipalgprofile_binding"
}

func (d *LsngroupLsnsipalgprofileBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *LsngroupLsnsipalgprofileBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = LsngroupLsnsipalgprofileBindingDataSourceSchema()
}

func (d *LsngroupLsnsipalgprofileBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data LsngroupLsnsipalgprofileBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	groupname_Name := data.Groupname.ValueString()
	sipalgprofilename_Name := data.Sipalgprofilename

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Lsngroup_lsnsipalgprofile_binding.Type(),
		ResourceName:             groupname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read lsngroup_lsnsipalgprofile_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "lsngroup_lsnsipalgprofile_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check sipalgprofilename
		if val, ok := v["sipalgprofilename"].(string); ok {
			if sipalgprofilename_Name.IsNull() || val != sipalgprofilename_Name.ValueString() {
				match = false
				continue
			}
		} else if !sipalgprofilename_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("lsngroup_lsnsipalgprofile_binding with sipalgprofilename %s not found", sipalgprofilename_Name))
		return
	}

	lsngroup_lsnsipalgprofile_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
