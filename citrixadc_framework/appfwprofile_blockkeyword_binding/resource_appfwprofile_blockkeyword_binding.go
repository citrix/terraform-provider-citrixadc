package appfwprofile_blockkeyword_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppfwprofileBlockkeywordBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileBlockkeywordBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileBlockkeywordBindingResource)(nil)

func NewAppfwprofileBlockkeywordBindingResource() resource.Resource {
	return &AppfwprofileBlockkeywordBindingResource{}
}

// AppfwprofileBlockkeywordBindingResource defines the resource implementation.
type AppfwprofileBlockkeywordBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileBlockkeywordBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileBlockkeywordBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_blockkeyword_binding"
}

func (r *AppfwprofileBlockkeywordBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileBlockkeywordBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileBlockkeywordBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_blockkeyword_binding resource")
	appfwprofile_blockkeyword_binding := appfwprofile_blockkeyword_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Appfwprofile_blockkeyword_binding.Type(), &appfwprofile_blockkeyword_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_blockkeyword_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwprofile_blockkeyword_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("as_blockkeyword_formurl:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsBlockkeywordFormurl.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("blockkeyword:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Blockkeyword.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("fieldname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Fieldname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readAppfwprofileBlockkeywordBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileBlockkeywordBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileBlockkeywordBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_blockkeyword_binding resource")

	r.readAppfwprofileBlockkeywordBindingFromApi(ctx, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	// Resource was deleted out-of-band - remove from state for self-healing
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileBlockkeywordBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwprofileBlockkeywordBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Pattern 5: Update is a no-op for this binding. NITRO exposes no update
	// endpoint for appfwprofile_blockkeyword_binding (only add/delete/get), and
	// every schema attribute is RequiresReplace, so Terraform never invokes Update
	// with an actual changed value. Just re-read and persist state.
	tflog.Debug(ctx, "Update is a no-op for appfwprofile_blockkeyword_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readAppfwprofileBlockkeywordBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileBlockkeywordBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileBlockkeywordBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_blockkeyword_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name_value, ok := idMap["name"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'name' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["as_blockkeyword_formurl"]; ok && val != "" {
		// as_blockkeyword_formurl is a form action URL containing reserved characters (':' and '/').
		// nitro-go does not encode ?args= values, so encode it to avoid a 400 from NITRO.
		argsMap["as_blockkeyword_formurl"] = utils.UrlEncode(val)
	}
	if val, ok := idMap["blockkeyword"]; ok && val != "" {
		argsMap["blockkeyword"] = val
	}
	if val, ok := idMap["fieldname"]; ok && val != "" {
		argsMap["fieldname"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Appfwprofile_blockkeyword_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwprofile_blockkeyword_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwprofile_blockkeyword_binding binding")
}

// Helper function to read appfwprofile_blockkeyword_binding data from API
func (r *AppfwprofileBlockkeywordBindingResource) readAppfwprofileBlockkeywordBindingFromApi(ctx context.Context, data *AppfwprofileBlockkeywordBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	name_Name, ok := idMap["name"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'name' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Appfwprofile_blockkeyword_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_blockkeyword_binding, got error: %s", err))
		return
	}

	// Resource is missing - signal deletion for self-healing
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check as_blockkeyword_formurl
		if idVal, ok := idMap["as_blockkeyword_formurl"]; ok {
			if val, ok := v["as_blockkeyword_formurl"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["as_blockkeyword_formurl"].(string); ok {
			match = false
			continue
		}

		// Check blockkeyword
		if idVal, ok := idMap["blockkeyword"]; ok {
			if val, ok := v["blockkeyword"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["blockkeyword"].(string); ok {
			match = false
			continue
		}

		// Check fieldname
		if idVal, ok := idMap["fieldname"]; ok {
			if val, ok := v["fieldname"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["fieldname"].(string); ok {
			match = false
			continue
		}
		if match {
			foundIndex = i
			break
		}
	}

	//  Resource is missing - signal deletion for self-healing
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	appfwprofile_blockkeyword_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
