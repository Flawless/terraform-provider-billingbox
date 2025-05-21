package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"terraform-provider-billingbox/internal/client"
)

// Ensure provider defined types fully satisfy framework interfaces.

var _ resource.Resource = &RoleResource{}
var _ resource.ResourceWithImportState = &RoleResource{}

func NewRoleResource() resource.Resource {
	return &RoleResource{}
}

// RoleResource defines the resource implementation.
type RoleResource struct {
	client *client.Client
}

type RoleUser struct {
	ID           types.String `json:"id,omitempty" tfsdk:"id"`
	ResourceType types.String `json:"resourceType" tfsdk:"resource_type"`
}

// RoleResourceModel describes the resource data model.
type RoleResourceModel struct {
	ResourceType types.String `json:"resourceType"   tfsdk:"resource_type"`
	ID           types.String `json:"id,omitempty"   tfsdk:"id"`
	User         *RoleUser    `json:"user,omitempty" tfsdk:"user"`
	Name         types.String `json:"name,omitempty" tfsdk:"name"`
	Meta         types.Object `json:"meta,omitempty" tfsdk:"meta"`
}

func (r *RoleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_role"
}

func (r *RoleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Role resource",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Unique identifier for the Role",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				MarkdownDescription: "Resource type of the Role. Always set to 'Role'.",
				Default:             stringdefault.StaticString("Role"),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"user": schema.SingleNestedAttribute{
				Required:            true,
				MarkdownDescription: "User associated with the role",
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "Unique identifier for the User",
					},
					"resource_type": schema.StringAttribute{
						Computed:            true,
						Optional:            true,
						MarkdownDescription: "Resource type of the User. Always set to 'User'.",
						Default:             stringdefault.StaticString("User"),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Name of the Role",
			},
			// shared metadata from schema_metadata.go
			"meta": MetaAttributes(),
		},
	}
}

func (r *RoleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *Client, got: %T", req.ProviderData),
		)
		return
	}
	r.client = client
}

func (r *RoleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Creating role resource")

	var data *RoleResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Convert to API type
	role := &client.Role{
		Resource: client.Resource{
			ResourceType: "Role",
			ID:           data.ID.ValueString(),
		},
		Name: data.Name.ValueString(),
	}

	if data.User != nil && !data.User.ID.IsNull() && !data.User.ID.IsUnknown() {
		role.User = &client.RoleUser{
			ResourceType: "User",
			ID:           data.User.ID.ValueString(),
		}
	}

	// Create the role
	result, err := r.client.CreateResource("Role", role)
	if err != nil {
		resp.Diagnostics.AddError("Error creating role", err.Error())
		return
	}

	// Update the model with the response data
	if id, ok := result["id"].(string); ok {
		data.ID = types.StringValue(id)
	}
	data.ResourceType = types.StringValue("Role")
	if name, ok := result["name"].(string); ok {
		data.Name = types.StringValue(name)
	}

	// Initialize user if not already initialized
	if data.User == nil {
		data.User = &RoleUser{}
	}

	// Set the user field from the response
	if user, ok := result["user"].(map[string]interface{}); ok {
		if userID, ok := user["id"].(string); ok {
			data.User.ID = types.StringValue(userID)
		}
		if resourceType, ok := user["resourceType"].(string); ok {
			data.User.ResourceType = types.StringValue(resourceType)
		}
	} else {
		// If user is not in response but was in request, preserve the request values
		if data.User.ID.IsNull() {
			data.User.ID = types.StringValue(role.User.ID)
		}
		if data.User.ResourceType.IsNull() {
			data.User.ResourceType = types.StringValue("User")
		}
	}

	if meta, ok := result["meta"].(map[string]interface{}); ok {
		metaValues := map[string]attr.Value{}
		if versionID, ok := meta["versionId"].(string); ok {
			metaValues["version_id"] = types.StringValue(versionID)
		}
		if createdAt, ok := meta["createdAt"].(string); ok {
			metaValues["created_at"] = types.StringValue(createdAt)
		}
		if lastUpdated, ok := meta["lastUpdated"].(string); ok {
			metaValues["last_updated"] = types.StringValue(lastUpdated)
		}
		metaTypes := map[string]attr.Type{
			"version_id":   types.StringType,
			"created_at":   types.StringType,
			"last_updated": types.StringType,
		}
		metaObj, diags := types.ObjectValue(metaTypes, metaValues)
		if !diags.HasError() {
			data.Meta = metaObj
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RoleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Reading role resource")

	var data RoleResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	role, err := r.client.GetResource("Role", data.ID.ValueString())
	if err != nil {
		// If the resource doesn't exist, mark it for recreation
		if client.IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error reading role", err.Error())
		return
	}

	// Update the model with the response data
	if id, ok := role["id"].(string); ok {
		data.ID = types.StringValue(id)
	}
	data.ResourceType = types.StringValue("Role")
	if name, ok := role["name"].(string); ok {
		data.Name = types.StringValue(name)
	}

	// Initialize user if not already initialized
	if data.User == nil {
		data.User = &RoleUser{}
	}

	// Set the user field from the response
	if user, ok := role["user"].(map[string]interface{}); ok {
		if userID, ok := user["id"].(string); ok {
			data.User.ID = types.StringValue(userID)
		}
		if resourceType, ok := user["resourceType"].(string); ok {
			data.User.ResourceType = types.StringValue(resourceType)
		}
	} else {
		// If user is not in response but exists in state, preserve the state values
		if !data.User.ID.IsNull() {
			// Ensure resource_type is set to "User" when we have a user ID
			data.User.ResourceType = types.StringValue("User")
		} else {
			// If we don't have a user ID in state, we should error as user is required
			resp.Diagnostics.AddError(
				"Missing Required Field",
				"User ID is required but not found in state or API response",
			)
			return
		}
	}

	if meta, ok := role["meta"].(map[string]interface{}); ok {
		metaValues := map[string]attr.Value{}
		if versionID, ok := meta["versionId"].(string); ok {
			metaValues["version_id"] = types.StringValue(versionID)
		}
		if createdAt, ok := meta["createdAt"].(string); ok {
			metaValues["created_at"] = types.StringValue(createdAt)
		}
		if lastUpdated, ok := meta["lastUpdated"].(string); ok {
			metaValues["last_updated"] = types.StringValue(lastUpdated)
		}
		metaTypes := map[string]attr.Type{
			"version_id":   types.StringType,
			"created_at":   types.StringType,
			"last_updated": types.StringType,
		}
		metaObj, diags := types.ObjectValue(metaTypes, metaValues)
		if !diags.HasError() {
			data.Meta = metaObj
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RoleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Updating role resource")

	var data RoleResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Convert to API type
	role := &client.Role{
		User: &client.RoleUser{
			ResourceType: "User",
			ID:           data.User.ID.ValueString(),
		},
		Name: data.Name.ValueString(),
	}

	// Update the role
	result, err := r.client.UpdateResource("Role", data.ID.ValueString(), role)
	if err != nil {
		resp.Diagnostics.AddError("Error updating role", err.Error())
		return
	}

	// Update the model with the response data
	if id, ok := result["id"].(string); ok {
		data.ID = types.StringValue(id)
	}
	if meta, ok := result["meta"].(map[string]interface{}); ok {
		metaValues := map[string]attr.Value{}
		if versionID, ok := meta["versionId"].(string); ok {
			metaValues["version_id"] = types.StringValue(versionID)
		}
		if createdAt, ok := meta["createdAt"].(string); ok {
			metaValues["created_at"] = types.StringValue(createdAt)
		}
		if lastUpdated, ok := meta["lastUpdated"].(string); ok {
			metaValues["last_updated"] = types.StringValue(lastUpdated)
		}
		metaTypes := map[string]attr.Type{
			"version_id":   types.StringType,
			"created_at":   types.StringType,
			"last_updated": types.StringType,
		}
		metaObj, diags := types.ObjectValue(metaTypes, metaValues)
		if !diags.HasError() {
			data.Meta = metaObj
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RoleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Deleting role resource")

	var data RoleResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteResource("Role", data.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error deleting role", err.Error())
		return
	}
}

func (r *RoleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
