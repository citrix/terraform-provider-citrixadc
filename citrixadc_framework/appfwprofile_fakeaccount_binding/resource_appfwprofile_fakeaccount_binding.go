package appfwprofile_fakeaccount_binding

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
var _ resource.Resource = &AppfwprofileFakeaccountBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileFakeaccountBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileFakeaccountBindingResource)(nil)
var _ resource.ResourceWithValidateConfig = (*AppfwprofileFakeaccountBindingResource)(nil)

func NewAppfwprofileFakeaccountBindingResource() resource.Resource {
	return &AppfwprofileFakeaccountBindingResource{}
}

// AppfwprofileFakeaccountBindingResource defines the resource implementation.
type AppfwprofileFakeaccountBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileFakeaccountBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileFakeaccountBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_fakeaccount_binding"
}

func (r *AppfwprofileFakeaccountBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

// ValidateConfig enforces the mutual exclusion between formexpression and
// formurl_fad. Per the CLI synopsis (man bind appfw profile, fakeAccount
// branch): (-fakeAccount <expression> -tag <expression> [-isFieldNameRegex (
// REGEX | NOTREGEX )] [-formExpression <expression>] [-formURL <expression>]).
// Both formExpression and formURL are bracketed-optional, but the appliance
// rejects setting both at once with NITRO errorcode 390 ("Only one of FormURL
// or FormExpression is allowed"). The synopsis does not mark either as
// mandatory, so neither is required here - only the at-most-one constraint is
// enforced.
func (r *AppfwprofileFakeaccountBindingResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data AppfwprofileFakeaccountBindingResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	formExprSet := !data.Formexpression.IsNull() && !data.Formexpression.IsUnknown()
	formUrlSet := !data.FormurlFad.IsNull() && !data.FormurlFad.IsUnknown()

	if formExprSet && formUrlSet {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"Only one of \"formexpression\" or \"formurl_fad\" may be set for an "+
				"appfwprofile_fakeaccount_binding. They are mutually exclusive "+
				"(NITRO errorcode 390: \"Only one of FormURL or FormExpression is allowed\").",
		)
	}
}

func (r *AppfwprofileFakeaccountBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileFakeaccountBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_fakeaccount_binding resource")
	appfwprofile_fakeaccount_binding := appfwprofile_fakeaccount_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Appfwprofile_fakeaccount_binding.Type(), &appfwprofile_fakeaccount_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_fakeaccount_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwprofile_fakeaccount_binding resource")

	// Set ID for the resource before reading state.
	// formexpression and formurl_fad are an at-most-one (mutually exclusive)
	// pair; only the populated arm is composed into the ID so the absent arm is
	// never emitted as an empty "key:" segment.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("fakeaccount:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Fakeaccount.ValueString()))))
	if !data.Formexpression.IsNull() && !data.Formexpression.IsUnknown() && data.Formexpression.ValueString() != "" {
		idParts = append(idParts, fmt.Sprintf("formexpression:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Formexpression.ValueString()))))
	}
	if !data.FormurlFad.IsNull() && !data.FormurlFad.IsUnknown() && data.FormurlFad.ValueString() != "" {
		idParts = append(idParts, fmt.Sprintf("formurl_fad:%s", utils.UrlEncode(fmt.Sprintf("%v", data.FormurlFad.ValueString()))))
	}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("tag:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Tag.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readAppfwprofileFakeaccountBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileFakeaccountBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileFakeaccountBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_fakeaccount_binding resource")

	r.readAppfwprofileFakeaccountBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileFakeaccountBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwprofileFakeaccountBindingResourceModel

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
	// endpoint for appfwprofile_fakeaccount_binding (only add/delete/get), and
	// every schema attribute is RequiresReplace, so Terraform never invokes Update
	// with an actual changed value. Just re-read and persist state.
	tflog.Debug(ctx, "Update is a no-op for appfwprofile_fakeaccount_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readAppfwprofileFakeaccountBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileFakeaccountBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileFakeaccountBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_fakeaccount_binding resource")
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
	if val, ok := idMap["fakeaccount"]; ok && val != "" {
		argsMap["fakeaccount"] = val
	}
	if val, ok := idMap["formexpression"]; ok && val != "" {
		// formexpression is a regular expression that may contain reserved characters.
		// nitro-go does not encode ?args= values, so encode it to avoid a 400 from NITRO.
		argsMap["formexpression"] = utils.UrlEncode(val)
	}
	if val, ok := idMap["formurl_fad"]; ok && val != "" {
		// formurl_fad is a URL containing reserved characters (':' and '/').
		// nitro-go does not encode ?args= values, so encode it to avoid a 400 from NITRO.
		argsMap["formurl_fad"] = utils.UrlEncode(val)
	}
	if val, ok := idMap["tag"]; ok && val != "" {
		// tag is a tag expression that may contain reserved characters.
		// nitro-go does not encode ?args= values, so encode it to avoid a 400 from NITRO.
		argsMap["tag"] = utils.UrlEncode(val)
	}

	err = r.client.DeleteResourceWithArgsMap(service.Appfwprofile_fakeaccount_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwprofile_fakeaccount_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwprofile_fakeaccount_binding binding")
}

// Helper function to read appfwprofile_fakeaccount_binding data from API
func (r *AppfwprofileFakeaccountBindingResource) readAppfwprofileFakeaccountBindingFromApi(ctx context.Context, data *AppfwprofileFakeaccountBindingResourceModel, diags *diag.Diagnostics) {

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
		ResourceType:             service.Appfwprofile_fakeaccount_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_fakeaccount_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "appfwprofile_fakeaccount_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check fakeaccount
		if idVal, ok := idMap["fakeaccount"]; ok {
			if val, ok := v["fakeaccount"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["fakeaccount"].(string); ok {
			match = false
			continue
		}

		// Check formexpression (at-most-one arm; may be absent from the ID).
		if idVal, ok := idMap["formexpression"]; ok {
			if val, ok := v["formexpression"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if val, ok := v["formexpression"].(string); ok && val != "" {
			// formexpression not part of this binding's ID, but the record has a
			// non-empty value - not our record.
			match = false
			continue
		}

		// Check formurl_fad (at-most-one arm; may be absent from the ID).
		if idVal, ok := idMap["formurl_fad"]; ok {
			if val, ok := v["formurl_fad"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if val, ok := v["formurl_fad"].(string); ok && val != "" {
			match = false
			continue
		}

		// Check tag
		if idVal, ok := idMap["tag"]; ok {
			if val, ok := v["tag"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["tag"].(string); ok {
			match = false
			continue
		}
		if match {
			foundIndex = i
			break
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("appfwprofile_fakeaccount_binding not found with the provided ID attributes"))
		return
	}

	appfwprofile_fakeaccount_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
