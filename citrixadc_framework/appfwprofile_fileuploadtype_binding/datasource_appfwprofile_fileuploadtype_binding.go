package appfwprofile_fileuploadtype_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AppfwprofileFileuploadtypeBindingDataSource)(nil)

func APpfwprofileFileuploadtypeBindingDataSource() datasource.DataSource {
	return &AppfwprofileFileuploadtypeBindingDataSource{}
}

type AppfwprofileFileuploadtypeBindingDataSource struct {
	client *service.NitroClient
}

func (d *AppfwprofileFileuploadtypeBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_fileuploadtype_binding"
}

func (d *AppfwprofileFileuploadtypeBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AppfwprofileFileuploadtypeBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AppfwprofileFileuploadtypeBindingDataSourceSchema()
}

func (d *AppfwprofileFileuploadtypeBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AppfwprofileFileuploadtypeBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	asfileuploadtypesurl_Name := data.AsFileuploadtypesUrl
	filetype_Name := ""
	if !data.Filetype.IsNull() && !data.Filetype.IsUnknown() {
		var filetypes []string
		diags := data.Filetype.ElementsAs(ctx, &filetypes, false)
		resp.Diagnostics.Append(diags...)
		filetype_Name = strings.Join(filetypes, ";")
	}
	fileuploadtype_Name := data.Fileuploadtype

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Appfwprofile_fileuploadtype_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_fileuploadtype_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "appfwprofile_fileuploadtype_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check as_fileuploadtypes_url
		if val, ok := v["as_fileuploadtypes_url"].(string); ok {
			if asfileuploadtypesurl_Name.IsNull() || val != asfileuploadtypesurl_Name.ValueString() {
				match = false
				continue
			}
		} else if !asfileuploadtypesurl_Name.IsNull() {
			match = false
			continue
		}

		// Check filetype
		if v["filetype"] != nil {
			if filetypeSlice, ok := v["filetype"].([]interface{}); ok {
				dataFiletype := strings.Join(utils.ToStringList(filetypeSlice), ";")
				if dataFiletype != filetype_Name {
					match = false
				}
			}
		}

		// Check fileuploadtype
		if val, ok := v["fileuploadtype"].(string); ok {
			if fileuploadtype_Name.IsNull() || val != fileuploadtype_Name.ValueString() {
				match = false
				continue
			}
		} else if !fileuploadtype_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("appfwprofile_fileuploadtype_binding with as_fileuploadtypes_url %s not found", asfileuploadtypesurl_Name))
		return
	}

	appfwprofile_fileuploadtype_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
