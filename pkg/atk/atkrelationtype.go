// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"fmt"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "stubs.h"
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_relation_type_get_type()), F: marshalRelationType},
	})
}

// RelationType describes the type of the relation
type RelationType int

const (
	// RelationNull: not used, represens "no relationship" or an error
	// condition.
	RelationNull RelationType = iota
	// RelationControlledBy indicates an object controlled by one or more target
	// objects.
	RelationControlledBy
	// RelationControllerFor indicates an object is an controller for one or
	// more target objects.
	RelationControllerFor
	// RelationLabelFor indicates an object is a label for one or more target
	// objects.
	RelationLabelFor
	// RelationLabelledBy indicates an object is labelled by one or more target
	// objects.
	RelationLabelledBy
	// RelationMemberOf indicates an object is a member of a group of one or
	// more target objects.
	RelationMemberOf
	// RelationNodeChildOf indicates an object is a cell in a treetable which is
	// displayed because a cell in the same column is expanded and identifies
	// that cell.
	RelationNodeChildOf
	// RelationFlowsTo indicates that the object has content that flows
	// logically to another AtkObject in a sequential way, (for instance
	// text-flow).
	RelationFlowsTo
	// RelationFlowsFrom indicates that the object has content that flows
	// logically from another AtkObject in a sequential way, (for instance
	// text-flow).
	RelationFlowsFrom
	// RelationSubwindowOf indicates a subwindow attached to a component but
	// otherwise has no connection in the UI heirarchy to that component.
	RelationSubwindowOf
	// RelationEmbeds indicates that the object visually embeds another object's
	// content, i.e. this object's content flows around another's content.
	RelationEmbeds
	// RelationEmbeddedBy: reciprocal of ATK_RELATION_EMBEDS, indicates that
	// this object's content is visualy embedded in another object.
	RelationEmbeddedBy
	// RelationPopupFor indicates that an object is a popup for another object.
	RelationPopupFor
	// RelationParentWindowOf indicates that an object is a parent window of
	// another object.
	RelationParentWindowOf
	// RelationDescribedBy: reciprocal of ATK_RELATION_DESCRIPTION_FOR.
	// Indicates that one or more target objects provide descriptive information
	// about this object. This relation type is most appropriate for information
	// that is not essential as its presentation may be user-configurable and/or
	// limited to an on-demand mechanism such as an assistive technology
	// command. For brief, essential information such as can be found in a
	// widget's on-screen label, use ATK_RELATION_LABELLED_BY. For an on-screen
	// error message, use ATK_RELATION_ERROR_MESSAGE. For lengthy extended
	// descriptive information contained in an on-screen object, consider using
	// ATK_RELATION_DETAILS as assistive technologies may provide a means for
	// the user to navigate to objects containing detailed descriptions so that
	// their content can be more closely reviewed.
	RelationDescribedBy
	// RelationDescriptionFor: reciprocal of ATK_RELATION_DESCRIBED_BY.
	// Indicates that this object provides descriptive information about the
	// target object(s). See also ATK_RELATION_DETAILS_FOR and
	// ATK_RELATION_ERROR_FOR.
	RelationDescriptionFor
	// RelationNodeParentOf indicates an object is a cell in a treetable and is
	// expanded to display other cells in the same column.
	RelationNodeParentOf
	// RelationDetails: reciprocal of ATK_RELATION_DETAILS_FOR. Indicates that
	// this object has a detailed or extended description, the contents of which
	// can be found in the target object(s). This relation type is most
	// appropriate for information that is sufficiently lengthy as to make
	// navigation to the container of that information desirable. For less
	// verbose information suitable for announcement only, see
	// ATK_RELATION_DESCRIBED_BY. If the detailed information describes an error
	// condition, ATK_RELATION_ERROR_FOR should be used instead. Since:
	// ATK-2.26.
	RelationDetails
	// RelationDetailsFor: reciprocal of ATK_RELATION_DETAILS. Indicates that
	// this object provides a detailed or extended description about the target
	// object(s). See also ATK_RELATION_DESCRIPTION_FOR and
	// ATK_RELATION_ERROR_FOR. Since: ATK-2.26.
	RelationDetailsFor
	// RelationErrorMessage: reciprocal of ATK_RELATION_ERROR_FOR. Indicates
	// that this object has one or more errors, the nature of which is described
	// in the contents of the target object(s). Objects that have this relation
	// type should also contain ATK_STATE_INVALID_ENTRY in their StateSet.
	// Since: ATK-2.26.
	RelationErrorMessage
	// RelationErrorFor: reciprocal of ATK_RELATION_ERROR_MESSAGE. Indicates
	// that this object contains an error message describing an invalid
	// condition in the target object(s). Since: ATK_2.26.
	RelationErrorFor
	// RelationLastDefined: not used, this value indicates the end of the
	// enumeration.
	RelationLastDefined
)

func marshalRelationType(p uintptr) (interface{}, error) {
	return RelationType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for RelationType.
func (r RelationType) String() string {
	switch r {
	case RelationNull:
		return "Null"
	case RelationControlledBy:
		return "ControlledBy"
	case RelationControllerFor:
		return "ControllerFor"
	case RelationLabelFor:
		return "LabelFor"
	case RelationLabelledBy:
		return "LabelledBy"
	case RelationMemberOf:
		return "MemberOf"
	case RelationNodeChildOf:
		return "NodeChildOf"
	case RelationFlowsTo:
		return "FlowsTo"
	case RelationFlowsFrom:
		return "FlowsFrom"
	case RelationSubwindowOf:
		return "SubwindowOf"
	case RelationEmbeds:
		return "Embeds"
	case RelationEmbeddedBy:
		return "EmbeddedBy"
	case RelationPopupFor:
		return "PopupFor"
	case RelationParentWindowOf:
		return "ParentWindowOf"
	case RelationDescribedBy:
		return "DescribedBy"
	case RelationDescriptionFor:
		return "DescriptionFor"
	case RelationNodeParentOf:
		return "NodeParentOf"
	case RelationDetails:
		return "Details"
	case RelationDetailsFor:
		return "DetailsFor"
	case RelationErrorMessage:
		return "ErrorMessage"
	case RelationErrorFor:
		return "ErrorFor"
	case RelationLastDefined:
		return "LastDefined"
	default:
		return fmt.Sprintf("RelationType(%d)", r)
	}
}
