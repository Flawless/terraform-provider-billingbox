package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Meta represents the metadata of a resource
type Meta struct {
	VersionID   types.String `json:"versionId,omitempty" tfsdk:"version_id"`
	CreatedAt   types.String `json:"createdAt,omitempty" tfsdk:"created_at"`
	LastUpdated types.String `json:"lastUpdated,omitempty" tfsdk:"last_updated"`
}

// MetaAttributes returns the schema attributes for metadata
func MetaAttributes() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Computed:            true,
		MarkdownDescription: "Server-side resource metadata",
		Attributes: map[string]schema.Attribute{
			"version_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Version identifier for the resource",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"created_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Time when the resource was created",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"last_updated": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Time when the resource was last updated",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}
