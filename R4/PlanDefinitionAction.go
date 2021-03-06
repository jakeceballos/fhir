// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "encoding/json"
    "fmt"
    "bytes"
)

// PlanDefinitionAction This resource allows for the definition of various types of plans as a sharable, consumable, and executable artifact. The resource is general enough to support the description of a broad range of clinical artifacts such as clinical decision support rules, order sets and protocols.
type PlanDefinitionAction struct {

  // Sub actions that are contained within the action. The behavior of this action determines the functionality of the sub-actions. For example, a selection behavior of at-most-one indicates that of the sub-actions, at most one may be chosen as part of realizing the action definition.
  Action []*PlanDefinitionAction `json:"action,omitempty"`

  // Defines whether the action can be selected multiple times.
  CardinalityBehavior interface{} `json:"cardinalityBehavior,omitempty"`

  // A code that provides meaning for the action or action group. For example, a section may have a LOINC code for the section of a documentation template.
  Code []*CodeableConcept `json:"code,omitempty"`

  // An expression that describes applicability criteria or start/stop conditions for the action.
  Condition []*PlanDefinitionCondition `json:"condition,omitempty"`

  // A reference to an ActivityDefinition that describes the action to be taken in detail, or a PlanDefinition that describes a series of actions to be taken.
  DefinitionCanonical string `json:"definitionCanonical,omitempty"`

  // A reference to an ActivityDefinition that describes the action to be taken in detail, or a PlanDefinition that describes a series of actions to be taken.
  DefinitionUri string `json:"definitionUri,omitempty"`

  // A brief description of the action used to provide a summary to display to the user.
  Description string `json:"description,omitempty"`

  // Didactic or other informational resources associated with the action that can be provided to the CDS recipient. Information resources can include inline text commentary and links to web resources.
  Documentation []*RelatedArtifact `json:"documentation,omitempty"`

  // Customizations that should be applied to the statically defined resource. For example, if the dosage of a medication must be computed based on the patient's weight, a customization would be used to specify an expression that calculated the weight, and the path on the resource that would contain the result.
  DynamicValue []*PlanDefinitionDynamicValue `json:"dynamicValue,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Identifies goals that this action supports. The reference must be to a goal element defined within this plan definition.
  GoalId []string `json:"goalId,omitempty"`

  // Defines the grouping behavior for the action and its children.
  GroupingBehavior interface{} `json:"groupingBehavior,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // Defines input data requirements for the action.
  Input []*DataRequirement `json:"input,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // Defines the outputs of the action, if any.
  Output []*DataRequirement `json:"output,omitempty"`

  // Indicates who should participate in performing the action described.
  Participant []*PlanDefinitionParticipant `json:"participant,omitempty"`

  // Defines whether the action should usually be preselected.
  PrecheckBehavior interface{} `json:"precheckBehavior,omitempty"`

  // A user-visible prefix for the action.
  Prefix string `json:"prefix,omitempty"`

  // Indicates how quickly the action should be addressed with respect to other actions.
  Priority string `json:"priority,omitempty"`

  // A description of why this action is necessary or appropriate.
  Reason []*CodeableConcept `json:"reason,omitempty"`

  // A relationship to another action such as "before" or "30-60 minutes after start of".
  RelatedAction []*PlanDefinitionRelatedAction `json:"relatedAction,omitempty"`

  // Defines the required behavior for the action.
  RequiredBehavior interface{} `json:"requiredBehavior,omitempty"`

  // Defines the selection behavior for the action and its children.
  SelectionBehavior interface{} `json:"selectionBehavior,omitempty"`

  // A code or group definition that describes the intended subject of the action and its children, if any.
  SubjectCodeableConcept *CodeableConcept `json:"subjectCodeableConcept,omitempty"`

  // A code or group definition that describes the intended subject of the action and its children, if any.
  SubjectReference *Reference `json:"subjectReference,omitempty"`

  // A text equivalent of the action to be performed. This provides a human-interpretable description of the action when the definition is consumed by a system that might not be capable of interpreting it dynamically.
  TextEquivalent string `json:"textEquivalent,omitempty"`

  // An optional value describing when the action should be performed.
  TimingAge *Age `json:"timingAge,omitempty"`

  // An optional value describing when the action should be performed.
  TimingDateTime string `json:"timingDateTime,omitempty"`

  // An optional value describing when the action should be performed.
  TimingDuration *Duration `json:"timingDuration,omitempty"`

  // An optional value describing when the action should be performed.
  TimingPeriod *Period `json:"timingPeriod,omitempty"`

  // An optional value describing when the action should be performed.
  TimingRange *Range `json:"timingRange,omitempty"`

  // An optional value describing when the action should be performed.
  TimingTiming *Timing `json:"timingTiming,omitempty"`

  // The title of the action displayed to a user.
  Title string `json:"title,omitempty"`

  // A reference to a StructureMap resource that defines a transform that can be executed to produce the intent resource using the ActivityDefinition instance as the input.
  Transform string `json:"transform,omitempty"`

  // A description of when the action should be triggered.
  Trigger []*TriggerDefinition `json:"trigger,omitempty"`

  // The type of action to perform (create, update, remove).
  Type *CodeableConcept `json:"type,omitempty"`
}

func (strct *PlanDefinitionAction) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "action" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"action\": ")
	if tmp, err := json.Marshal(strct.Action); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "cardinalityBehavior" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"cardinalityBehavior\": ")
	if tmp, err := json.Marshal(strct.CardinalityBehavior); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "code" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"code\": ")
	if tmp, err := json.Marshal(strct.Code); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "condition" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"condition\": ")
	if tmp, err := json.Marshal(strct.Condition); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "definitionCanonical" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"definitionCanonical\": ")
	if tmp, err := json.Marshal(strct.DefinitionCanonical); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "definitionUri" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"definitionUri\": ")
	if tmp, err := json.Marshal(strct.DefinitionUri); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "description" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"description\": ")
	if tmp, err := json.Marshal(strct.Description); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "documentation" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"documentation\": ")
	if tmp, err := json.Marshal(strct.Documentation); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "dynamicValue" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"dynamicValue\": ")
	if tmp, err := json.Marshal(strct.DynamicValue); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "extension" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"extension\": ")
	if tmp, err := json.Marshal(strct.Extension); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "goalId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"goalId\": ")
	if tmp, err := json.Marshal(strct.GoalId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "groupingBehavior" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"groupingBehavior\": ")
	if tmp, err := json.Marshal(strct.GroupingBehavior); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "id" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"id\": ")
	if tmp, err := json.Marshal(strct.Id); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "input" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"input\": ")
	if tmp, err := json.Marshal(strct.Input); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "modifierExtension" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"modifierExtension\": ")
	if tmp, err := json.Marshal(strct.ModifierExtension); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "output" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"output\": ")
	if tmp, err := json.Marshal(strct.Output); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "participant" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"participant\": ")
	if tmp, err := json.Marshal(strct.Participant); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "precheckBehavior" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"precheckBehavior\": ")
	if tmp, err := json.Marshal(strct.PrecheckBehavior); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "prefix" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"prefix\": ")
	if tmp, err := json.Marshal(strct.Prefix); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "priority" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"priority\": ")
	if tmp, err := json.Marshal(strct.Priority); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "reason" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"reason\": ")
	if tmp, err := json.Marshal(strct.Reason); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "relatedAction" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"relatedAction\": ")
	if tmp, err := json.Marshal(strct.RelatedAction); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "requiredBehavior" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"requiredBehavior\": ")
	if tmp, err := json.Marshal(strct.RequiredBehavior); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "selectionBehavior" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"selectionBehavior\": ")
	if tmp, err := json.Marshal(strct.SelectionBehavior); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "subjectCodeableConcept" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"subjectCodeableConcept\": ")
	if tmp, err := json.Marshal(strct.SubjectCodeableConcept); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "subjectReference" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"subjectReference\": ")
	if tmp, err := json.Marshal(strct.SubjectReference); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "textEquivalent" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"textEquivalent\": ")
	if tmp, err := json.Marshal(strct.TextEquivalent); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "timingAge" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"timingAge\": ")
	if tmp, err := json.Marshal(strct.TimingAge); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "timingDateTime" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"timingDateTime\": ")
	if tmp, err := json.Marshal(strct.TimingDateTime); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "timingDuration" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"timingDuration\": ")
	if tmp, err := json.Marshal(strct.TimingDuration); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "timingPeriod" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"timingPeriod\": ")
	if tmp, err := json.Marshal(strct.TimingPeriod); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "timingRange" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"timingRange\": ")
	if tmp, err := json.Marshal(strct.TimingRange); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "timingTiming" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"timingTiming\": ")
	if tmp, err := json.Marshal(strct.TimingTiming); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "title" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"title\": ")
	if tmp, err := json.Marshal(strct.Title); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "transform" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"transform\": ")
	if tmp, err := json.Marshal(strct.Transform); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "trigger" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"trigger\": ")
	if tmp, err := json.Marshal(strct.Trigger); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "type" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"type\": ")
	if tmp, err := json.Marshal(strct.Type); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *PlanDefinitionAction) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "action":
            if err := json.Unmarshal([]byte(v), &strct.Action); err != nil {
                return err
             }
        case "cardinalityBehavior":
            if err := json.Unmarshal([]byte(v), &strct.CardinalityBehavior); err != nil {
                return err
             }
        case "code":
            if err := json.Unmarshal([]byte(v), &strct.Code); err != nil {
                return err
             }
        case "condition":
            if err := json.Unmarshal([]byte(v), &strct.Condition); err != nil {
                return err
             }
        case "definitionCanonical":
            if err := json.Unmarshal([]byte(v), &strct.DefinitionCanonical); err != nil {
                return err
             }
        case "definitionUri":
            if err := json.Unmarshal([]byte(v), &strct.DefinitionUri); err != nil {
                return err
             }
        case "description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
                return err
             }
        case "documentation":
            if err := json.Unmarshal([]byte(v), &strct.Documentation); err != nil {
                return err
             }
        case "dynamicValue":
            if err := json.Unmarshal([]byte(v), &strct.DynamicValue); err != nil {
                return err
             }
        case "extension":
            if err := json.Unmarshal([]byte(v), &strct.Extension); err != nil {
                return err
             }
        case "goalId":
            if err := json.Unmarshal([]byte(v), &strct.GoalId); err != nil {
                return err
             }
        case "groupingBehavior":
            if err := json.Unmarshal([]byte(v), &strct.GroupingBehavior); err != nil {
                return err
             }
        case "id":
            if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
                return err
             }
        case "input":
            if err := json.Unmarshal([]byte(v), &strct.Input); err != nil {
                return err
             }
        case "modifierExtension":
            if err := json.Unmarshal([]byte(v), &strct.ModifierExtension); err != nil {
                return err
             }
        case "output":
            if err := json.Unmarshal([]byte(v), &strct.Output); err != nil {
                return err
             }
        case "participant":
            if err := json.Unmarshal([]byte(v), &strct.Participant); err != nil {
                return err
             }
        case "precheckBehavior":
            if err := json.Unmarshal([]byte(v), &strct.PrecheckBehavior); err != nil {
                return err
             }
        case "prefix":
            if err := json.Unmarshal([]byte(v), &strct.Prefix); err != nil {
                return err
             }
        case "priority":
            if err := json.Unmarshal([]byte(v), &strct.Priority); err != nil {
                return err
             }
        case "reason":
            if err := json.Unmarshal([]byte(v), &strct.Reason); err != nil {
                return err
             }
        case "relatedAction":
            if err := json.Unmarshal([]byte(v), &strct.RelatedAction); err != nil {
                return err
             }
        case "requiredBehavior":
            if err := json.Unmarshal([]byte(v), &strct.RequiredBehavior); err != nil {
                return err
             }
        case "selectionBehavior":
            if err := json.Unmarshal([]byte(v), &strct.SelectionBehavior); err != nil {
                return err
             }
        case "subjectCodeableConcept":
            if err := json.Unmarshal([]byte(v), &strct.SubjectCodeableConcept); err != nil {
                return err
             }
        case "subjectReference":
            if err := json.Unmarshal([]byte(v), &strct.SubjectReference); err != nil {
                return err
             }
        case "textEquivalent":
            if err := json.Unmarshal([]byte(v), &strct.TextEquivalent); err != nil {
                return err
             }
        case "timingAge":
            if err := json.Unmarshal([]byte(v), &strct.TimingAge); err != nil {
                return err
             }
        case "timingDateTime":
            if err := json.Unmarshal([]byte(v), &strct.TimingDateTime); err != nil {
                return err
             }
        case "timingDuration":
            if err := json.Unmarshal([]byte(v), &strct.TimingDuration); err != nil {
                return err
             }
        case "timingPeriod":
            if err := json.Unmarshal([]byte(v), &strct.TimingPeriod); err != nil {
                return err
             }
        case "timingRange":
            if err := json.Unmarshal([]byte(v), &strct.TimingRange); err != nil {
                return err
             }
        case "timingTiming":
            if err := json.Unmarshal([]byte(v), &strct.TimingTiming); err != nil {
                return err
             }
        case "title":
            if err := json.Unmarshal([]byte(v), &strct.Title); err != nil {
                return err
             }
        case "transform":
            if err := json.Unmarshal([]byte(v), &strct.Transform); err != nil {
                return err
             }
        case "trigger":
            if err := json.Unmarshal([]byte(v), &strct.Trigger); err != nil {
                return err
             }
        case "type":
            if err := json.Unmarshal([]byte(v), &strct.Type); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}
