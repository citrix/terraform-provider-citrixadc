package lsngroup_lsnappsprofile_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*LsngroupLsnappsprofileBindingDataSource)(nil)

func LSngroupLsnappsprofileBindingDataSource() datasource.DataSource {
	return &LsngroupLsnappsprofileBindingDataSource{}
}

type LsngroupLsnappsprofileBindingDataSource struct {
	client *service.NitroClient
}

func (d *LsngroupLsnappsprofileBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsngroup_lsnappsprofile_binding"
}

func (d *LsngroupLsnappsprofileBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *LsngroupLsnappsprofileBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = LsngroupLsnappsprofileBindingDataSourceSchema()
}

func (d *LsngroupLsnappsprofileBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data LsngroupLsnappsprofileBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	groupname_Name := data.Groupname.ValueString()
	appsprofilename_Name := data.Appsprofilename

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Lsngroup_lsnappsprofile_binding.Type(),
		ResourceName:             groupname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read lsngroup_lsnappsprofile_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "lsngroup_lsnappsprofile_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check appsprofilename
		if val, ok := v["appsprofilename"].(string); ok {
			if appsprofilename_Name.IsNull() || val != appsprofilename_Name.ValueString() {
				match = false
				continue
			}
		} else if !appsprofilename_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("lsngroup_lsnappsprofile_binding with appsprofilename %s not found", appsprofilename_Name))
		return
	}

	lsngroup_lsnappsprofile_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
