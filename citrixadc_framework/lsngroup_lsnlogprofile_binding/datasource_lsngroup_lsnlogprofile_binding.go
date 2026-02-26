package lsngroup_lsnlogprofile_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*LsngroupLsnlogprofileBindingDataSource)(nil)

func LSngroupLsnlogprofileBindingDataSource() datasource.DataSource {
	return &LsngroupLsnlogprofileBindingDataSource{}
}

type LsngroupLsnlogprofileBindingDataSource struct {
	client *service.NitroClient
}

func (d *LsngroupLsnlogprofileBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsngroup_lsnlogprofile_binding"
}

func (d *LsngroupLsnlogprofileBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *LsngroupLsnlogprofileBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = LsngroupLsnlogprofileBindingDataSourceSchema()
}

func (d *LsngroupLsnlogprofileBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data LsngroupLsnlogprofileBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	groupname_Name := data.Groupname.ValueString()
	logprofilename_Name := data.Logprofilename

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Lsngroup_lsnlogprofile_binding.Type(),
		ResourceName:             groupname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read lsngroup_lsnlogprofile_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "lsngroup_lsnlogprofile_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check logprofilename
		if val, ok := v["logprofilename"].(string); ok {
			if logprofilename_Name.IsNull() || val != logprofilename_Name.ValueString() {
				match = false
				continue
			}
		} else if !logprofilename_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("lsngroup_lsnlogprofile_binding with logprofilename %s not found", logprofilename_Name))
		return
	}

	lsngroup_lsnlogprofile_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
