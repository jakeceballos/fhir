// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "encoding/json"
    "fmt"
)

// EvidenceVariableCharacteristic The EvidenceVariable resource describes a "PICO" element that knowledge (evidence, assertion, recommendation) is about.
type EvidenceVariableCharacteristic struct {

  // Define members of the evidence element using Codes (such as condition, medication, or observation), Expressions ( using an expression language such as FHIRPath or CQL) or DataRequirements (such as Diabetes diagnosis onset in the last year).
  DefinitionCanonical string `json:"definitionCanonical,omitempty"`

  // Define members of the evidence element using Codes (such as condition, medication, or observation), Expressions ( using an expression language such as FHIRPath or CQL) or DataRequirements (such as Diabetes diagnosis onset in the last year).
  DefinitionCodeableConcept *CodeableConcept `json:"definitionCodeableConcept,omitempty"`

  // Define members of the evidence element using Codes (such as condition, medication, or observation), Expressions ( using an expression language such as FHIRPath or CQL) or DataRequirements (such as Diabetes diagnosis onset in the last year).
  DefinitionDataRequirement *DataRequirement `json:"definitionDataRequirement,omitempty"`

  // Define members of the evidence element using Codes (such as condition, medication, or observation), Expressions ( using an expression language such as FHIRPath or CQL) or DataRequirements (such as Diabetes diagnosis onset in the last year).
  DefinitionExpression *Expression `json:"definitionExpression,omitempty"`

  // Define members of the evidence element using Codes (such as condition, medication, or observation), Expressions ( using an expression language such as FHIRPath or CQL) or DataRequirements (such as Diabetes diagnosis onset in the last year).
  DefinitionReference *Reference `json:"definitionReference,omitempty"`

  // Define members of the evidence element using Codes (such as condition, medication, or observation), Expressions ( using an expression language such as FHIRPath or CQL) or DataRequirements (such as Diabetes diagnosis onset in the last year).
  DefinitionTriggerDefinition *TriggerDefinition `json:"definitionTriggerDefinition,omitempty"`

  // A short, natural language description of the characteristic that could be used to communicate the criteria to an end-user.
  Description string `json:"description,omitempty"`

  // When true, members with this characteristic are excluded from the element.
  Exclude bool `json:"exclude,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Indicates how elements are aggregated within the study effective period.
  GroupMeasure interface{} `json:"groupMeasure,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // Indicates what effective period the study covers.
  ParticipantEffectiveDateTime string `json:"participantEffectiveDateTime,omitempty"`

  // Indicates what effective period the study covers.
  ParticipantEffectiveDuration *Duration `json:"participantEffectiveDuration,omitempty"`

  // Indicates what effective period the study covers.
  ParticipantEffectivePeriod *Period `json:"participantEffectivePeriod,omitempty"`

  // Indicates what effective period the study covers.
  ParticipantEffectiveTiming *Timing `json:"participantEffectiveTiming,omitempty"`

  // Indicates duration from the participant's study entry.
  TimeFromStart *Duration `json:"timeFromStart,omitempty"`

  // Use UsageContext to define the members of the population, such as Age Ranges, Genders, Settings.
  UsageContext []*UsageContext `json:"usageContext,omitempty"`
}

func (strct *EvidenceVariableCharacteristic) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
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
    // Marshal the "definitionCodeableConcept" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"definitionCodeableConcept\": ")
	if tmp, err := json.Marshal(strct.DefinitionCodeableConcept); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "definitionDataRequirement" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"definitionDataRequirement\": ")
	if tmp, err := json.Marshal(strct.DefinitionDataRequirement); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "definitionExpression" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"definitionExpression\": ")
	if tmp, err := json.Marshal(strct.DefinitionExpression); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "definitionReference" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"definitionReference\": ")
	if tmp, err := json.Marshal(strct.DefinitionReference); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "definitionTriggerDefinition" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"definitionTriggerDefinition\": ")
	if tmp, err := json.Marshal(strct.DefinitionTriggerDefinition); err != nil {
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
    // Marshal the "exclude" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"exclude\": ")
	if tmp, err := json.Marshal(strct.Exclude); err != nil {
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
    // Marshal the "groupMeasure" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"groupMeasure\": ")
	if tmp, err := json.Marshal(strct.GroupMeasure); err != nil {
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
    // Marshal the "participantEffectiveDateTime" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"participantEffectiveDateTime\": ")
	if tmp, err := json.Marshal(strct.ParticipantEffectiveDateTime); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "participantEffectiveDuration" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"participantEffectiveDuration\": ")
	if tmp, err := json.Marshal(strct.ParticipantEffectiveDuration); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "participantEffectivePeriod" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"participantEffectivePeriod\": ")
	if tmp, err := json.Marshal(strct.ParticipantEffectivePeriod); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "participantEffectiveTiming" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"participantEffectiveTiming\": ")
	if tmp, err := json.Marshal(strct.ParticipantEffectiveTiming); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "timeFromStart" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"timeFromStart\": ")
	if tmp, err := json.Marshal(strct.TimeFromStart); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "usageContext" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"usageContext\": ")
	if tmp, err := json.Marshal(strct.UsageContext); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *EvidenceVariableCharacteristic) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "definitionCanonical":
            if err := json.Unmarshal([]byte(v), &strct.DefinitionCanonical); err != nil {
                return err
             }
        case "definitionCodeableConcept":
            if err := json.Unmarshal([]byte(v), &strct.DefinitionCodeableConcept); err != nil {
                return err
             }
        case "definitionDataRequirement":
            if err := json.Unmarshal([]byte(v), &strct.DefinitionDataRequirement); err != nil {
                return err
             }
        case "definitionExpression":
            if err := json.Unmarshal([]byte(v), &strct.DefinitionExpression); err != nil {
                return err
             }
        case "definitionReference":
            if err := json.Unmarshal([]byte(v), &strct.DefinitionReference); err != nil {
                return err
             }
        case "definitionTriggerDefinition":
            if err := json.Unmarshal([]byte(v), &strct.DefinitionTriggerDefinition); err != nil {
                return err
             }
        case "description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
                return err
             }
        case "exclude":
            if err := json.Unmarshal([]byte(v), &strct.Exclude); err != nil {
                return err
             }
        case "extension":
            if err := json.Unmarshal([]byte(v), &strct.Extension); err != nil {
                return err
             }
        case "groupMeasure":
            if err := json.Unmarshal([]byte(v), &strct.GroupMeasure); err != nil {
                return err
             }
        case "id":
            if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
                return err
             }
        case "modifierExtension":
            if err := json.Unmarshal([]byte(v), &strct.ModifierExtension); err != nil {
                return err
             }
        case "participantEffectiveDateTime":
            if err := json.Unmarshal([]byte(v), &strct.ParticipantEffectiveDateTime); err != nil {
                return err
             }
        case "participantEffectiveDuration":
            if err := json.Unmarshal([]byte(v), &strct.ParticipantEffectiveDuration); err != nil {
                return err
             }
        case "participantEffectivePeriod":
            if err := json.Unmarshal([]byte(v), &strct.ParticipantEffectivePeriod); err != nil {
                return err
             }
        case "participantEffectiveTiming":
            if err := json.Unmarshal([]byte(v), &strct.ParticipantEffectiveTiming); err != nil {
                return err
             }
        case "timeFromStart":
            if err := json.Unmarshal([]byte(v), &strct.TimeFromStart); err != nil {
                return err
             }
        case "usageContext":
            if err := json.Unmarshal([]byte(v), &strct.UsageContext); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}
