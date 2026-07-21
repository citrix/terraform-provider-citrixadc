package lsnappsprofile_port_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*LsnappsprofilePortBindingDataSource)(nil)

func LSnappsprofilePortBindingDataSource() datasource.DataSource {
	return &LsnappsprofilePortBindingDataSource{}
}

type LsnappsprofilePortBindingDataSource struct {
	client *service.NitroClient
}

func (d *LsnappsprofilePortBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnappsprofile_port_binding"
}

func (d *LsnappsprofilePortBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *LsnappsprofilePortBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = LsnappsprofilePortBindingDataSourceSchema()
}

func (d *LsnappsprofilePortBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data LsnappsprofilePortBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	appsprofilename_Name := data.Appsprofilename.ValueString()
	lsnport_Name := data.Lsnport

	var err error

	// NOTE: The direct "lsnappsprofile_port_binding/<appsprofilename>" GET endpoint does not
	// return the bound ports on the ADC (it responds with an empty payload even when a port
	// is bound). The bound ports are only exposed through the aggregate
	// "lsnappsprofile_binding/<appsprofilename>" endpoint, so read from there instead and
	// extract the nested "lsnappsprofile_port_binding" array.
	findParams := service.FindParams{
		ResourceType:             "lsnappsprofile_binding",
		ResourceName:             appsprofilename_Name,
		ResourceMissingErrorCode: 258,
	}
	aggArr, err := d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read lsnappsprofile_port_binding, got error: %s", err))
		return
	}

	// Parent appsprofile is missing
	if len(aggArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "lsnappsprofile_port_binding returned empty array.")
		return
	}

	// Extract the nested lsnappsprofile_port_binding array from the aggregate response
	dataArr := []map[string]interface{}{}
	if raw, ok := aggArr[0]["lsnappsprofile_port_binding"]; ok && raw != nil {
		if arr, ok := raw.([]interface{}); ok {
			for _, item := range arr {
				if m, ok := item.(map[string]interface{}); ok {
					dataArr = append(dataArr, m)
				}
			}
		}
	}

	// No ports are bound to this appsprofile
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "lsnappsprofile_port_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check lsnport
		if val, ok := v["lsnport"].(string); ok {
			if lsnport_Name.IsNull() || val != lsnport_Name.ValueString() {
				match = false
				continue
			}
		} else if !lsnport_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("lsnappsprofile_port_binding with lsnport %s not found", lsnport_Name))
		return
	}

	lsnappsprofile_port_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
