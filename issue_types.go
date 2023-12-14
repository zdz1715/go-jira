package jira

// IssueType represents a type of a Jira issue.
// Typical types are "Request", "Bug", "Story", ...
type IssueType struct {
	Self           string `json:"self,omitempty"`
	ID             string `json:"id,omitempty"`
	EntityId       string `json:"entityId,omitempty"`
	Description    string `json:"description,omitempty"`
	IconUrl        string `json:"iconUrl,omitempty"`
	Name           string `json:"name,omitempty"`
	Subtask        bool   `json:"subtask,omitempty"`
	AvatarId       int    `json:"avatarId,omitempty"`
	HierarchyLevel int    `json:"hierarchyLevel,omitempty"`
}
