// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.labeler.defs

import (
	comatprototypes "github.com/bluesky-social/indigo/api/atproto"
)

// LabelerDefs_LabelerPolicies is a "labelerPolicies" in the app.bsky.labeler.defs schema.
type LabelerDefs_LabelerPolicies struct {
	// labelValueDefinitions: Label values created by this labeler and scoped exclusively to it. Labels defined here will override global label definitions for this labeler.
	LabelValueDefinitions []*comatprototypes.LabelDefs_LabelValueDefinition `json:"labelValueDefinitions,omitempty" cborgen:"labelValueDefinitions,omitempty"`
	// labelValues: The label values which this labeler publishes. May include global or custom labels.
	LabelValues []*string `json:"labelValues" cborgen:"labelValues"`
}

// LabelerDefs_LabelerView is a "labelerView" in the app.bsky.labeler.defs schema.
//
// RECORDTYPE: LabelerDefs_LabelerView
type LabelerDefs_LabelerView struct {
	LexiconTypeID string                             `json:"$type,const=app.bsky.labeler.defs#labelerView" cborgen:"$type,const=app.bsky.labeler.defs#labelerView"`
	Cid           string                             `json:"cid" cborgen:"cid"`
	Creator       *ActorDefs_ProfileView             `json:"creator" cborgen:"creator"`
	IndexedAt     string                             `json:"indexedAt" cborgen:"indexedAt"`
	Labels        []*comatprototypes.LabelDefs_Label `json:"labels,omitempty" cborgen:"labels,omitempty"`
	LikeCount     *int64                             `json:"likeCount,omitempty" cborgen:"likeCount,omitempty"`
	Uri           string                             `json:"uri" cborgen:"uri"`
	Viewer        *LabelerDefs_LabelerViewerState    `json:"viewer,omitempty" cborgen:"viewer,omitempty"`
}

// LabelerDefs_LabelerViewDetailed is a "labelerViewDetailed" in the app.bsky.labeler.defs schema.
//
// RECORDTYPE: LabelerDefs_LabelerViewDetailed
type LabelerDefs_LabelerViewDetailed struct {
	LexiconTypeID string                             `json:"$type,const=app.bsky.labeler.defs#labelerViewDetailed" cborgen:"$type,const=app.bsky.labeler.defs#labelerViewDetailed"`
	Cid           string                             `json:"cid" cborgen:"cid"`
	Creator       *ActorDefs_ProfileView             `json:"creator" cborgen:"creator"`
	IndexedAt     string                             `json:"indexedAt" cborgen:"indexedAt"`
	Labels        []*comatprototypes.LabelDefs_Label `json:"labels,omitempty" cborgen:"labels,omitempty"`
	LikeCount     *int64                             `json:"likeCount,omitempty" cborgen:"likeCount,omitempty"`
	Policies      *LabelerDefs_LabelerPolicies       `json:"policies" cborgen:"policies"`
	Uri           string                             `json:"uri" cborgen:"uri"`
	Viewer        *LabelerDefs_LabelerViewerState    `json:"viewer,omitempty" cborgen:"viewer,omitempty"`
}

// LabelerDefs_LabelerViewerState is a "labelerViewerState" in the app.bsky.labeler.defs schema.
type LabelerDefs_LabelerViewerState struct {
	Like *string `json:"like,omitempty" cborgen:"like,omitempty"`
}