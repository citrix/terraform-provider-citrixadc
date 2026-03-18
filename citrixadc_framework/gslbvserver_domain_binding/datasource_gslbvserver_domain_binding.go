package gslbvserver_domain_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*GslbvserverDomainBindingDataSource)(nil)

func GSlbvserverDomainBindingDataSource() datasource.DataSource {
	return &GslbvserverDomainBindingDataSource{}
}

type GslbvserverDomainBindingDataSource struct {
	client *service.NitroClient
}

func (d *GslbvserverDomainBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gslbvserver_domain_binding"
}

func (d *GslbvserverDomainBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *GslbvserverDomainBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = GslbvserverDomainBindingDataSourceSchema()
}

func (d *GslbvserverDomainBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data GslbvserverDomainBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	backupipflag_Name := data.Backupipflag
	cookiedomainflag_Name := data.CookieDomainflag
	domainname_Name := data.Domainname

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Gslbvserver_domain_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read gslbvserver_domain_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "gslbvserver_domain_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check backupipflag
		if val, ok := v["backupipflag"].(bool); ok {
			if backupipflag_Name.IsNull() || val != backupipflag_Name.ValueBool() {
				match = false
				continue
			}
		} else if !backupipflag_Name.IsNull() {
			match = false
			continue
		}

		// Check cookie_domainflag
		if val, ok := v["cookie_domainflag"].(bool); ok {
			if cookiedomainflag_Name.IsNull() || val != cookiedomainflag_Name.ValueBool() {
				match = false
				continue
			}
		} else if !cookiedomainflag_Name.IsNull() {
			match = false
			continue
		}

		// Check domainname
		if val, ok := v["domainname"].(string); ok {
			if domainname_Name.IsNull() || val != domainname_Name.ValueString() {
				match = false
				continue
			}
		} else if !domainname_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("gslbvserver_domain_binding with backupipflag %s not found", backupipflag_Name))
		return
	}

	gslbvserver_domain_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
