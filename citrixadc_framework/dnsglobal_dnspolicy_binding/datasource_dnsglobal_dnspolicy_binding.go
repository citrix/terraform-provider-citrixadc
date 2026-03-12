package dnsglobal_dnspolicy_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*DnsglobalDnspolicyBindingDataSource)(nil)

func DNsglobalDnspolicyBindingDataSource() datasource.DataSource {
	return &DnsglobalDnspolicyBindingDataSource{}
}

type DnsglobalDnspolicyBindingDataSource struct {
	client *service.NitroClient
}

func (d *DnsglobalDnspolicyBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnsglobal_dnspolicy_binding"
}

func (d *DnsglobalDnspolicyBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *DnsglobalDnspolicyBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DnsglobalDnspolicyBindingDataSourceSchema()
}

func (d *DnsglobalDnspolicyBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DnsglobalDnspolicyBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 3: Array filter without parent ID
	policyname_Name := data.Policyname
	type_Name := data.Type

	var dataArr []map[string]interface{}
	var argsMap map[string]string = make(map[string]string)
	var err error
	if !type_Name.IsNull() && type_Name.ValueString() != "" {
		argsMap["type"] = type_Name.ValueString()
	}

	findParams := service.FindParams{
		ResourceType:             service.Dnsglobal_dnspolicy_binding.Type(),
		ArgsMap:                  argsMap,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read dnsglobal_dnspolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "dnsglobal_dnspolicy_binding returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check policyname
		if val, ok := v["policyname"].(string); ok {
			if policyname_Name.IsNull() || val != policyname_Name.ValueString() {
				match = false
				continue
			}
		} else if !policyname_Name.IsNull() {
			match = false
			continue
		}
		if !type_Name.IsNull() && type_Name.ValueString() != "" {
			if v, ok := v["type"]; ok {
				if v.(string) != type_Name.ValueString() {
					match = false
				}
			}
		}

		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("dnsglobal_dnspolicy_binding with policyname %s not found", policyname_Name))
		return
	}

	dnsglobal_dnspolicy_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
