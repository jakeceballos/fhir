// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "errors"
    "encoding/json"
    "fmt"
    "bytes"
)

// ExplanationOfBenefitSupportingInfo This resource provides: the claim details; adjudication details from the processing of a Claim; and optionally account balance information, for informing the subscriber of the benefits provided.
type ExplanationOfBenefitSupportingInfo struct {

  // The general class of the information supplied: information; exception; accident, employment; onset, etc.
  Category *CodeableConcept `json:"category"`

  // System and code pertaining to the specific information regarding special conditions relating to the setting, treatment or patient  for which care is sought.
  Code *CodeableConcept `json:"code,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // Provides the reason in the situation where a reason code is required in addition to the content.
  Reason *Coding `json:"reason,omitempty"`

  // A number to uniquely identify supporting information entries.
  Sequence float64 `json:"sequence,omitempty"`

  // The date when or period to which this information refers.
  TimingDate string `json:"timingDate,omitempty"`

  // The date when or period to which this information refers.
  TimingPeriod *Period `json:"timingPeriod,omitempty"`

  // Additional data or information such as resources, documents, images etc. including references to the data or the actual inclusion of the data.
  ValueAttachment *Attachment `json:"valueAttachment,omitempty"`

  // Additional data or information such as resources, documents, images etc. including references to the data or the actual inclusion of the data.
  ValueBoolean bool `json:"valueBoolean,omitempty"`

  // Additional data or information such as resources, documents, images etc. including references to the data or the actual inclusion of the data.
  ValueQuantity *Quantity `json:"valueQuantity,omitempty"`

  // Additional data or information such as resources, documents, images etc. including references to the data or the actual inclusion of the data.
  ValueReference *Reference `json:"valueReference,omitempty"`

  // Additional data or information such as resources, documents, images etc. including references to the data or the actual inclusion of the data.
  ValueString string `json:"valueString,omitempty"`
}

func (strct *ExplanationOfBenefitSupportingInfo) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "Category" field is required
    if strct.Category == nil {
        return nil, errors.New("category is a required field")
    }
    // Marshal the "category" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"category\": ")
	if tmp, err := json.Marshal(strct.Category); err != nil {
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
    // Marshal the "sequence" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"sequence\": ")
	if tmp, err := json.Marshal(strct.Sequence); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "timingDate" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"timingDate\": ")
	if tmp, err := json.Marshal(strct.TimingDate); err != nil {
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
    // Marshal the "valueAttachment" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueAttachment\": ")
	if tmp, err := json.Marshal(strct.ValueAttachment); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueBoolean" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueBoolean\": ")
	if tmp, err := json.Marshal(strct.ValueBoolean); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueQuantity" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueQuantity\": ")
	if tmp, err := json.Marshal(strct.ValueQuantity); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueReference" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueReference\": ")
	if tmp, err := json.Marshal(strct.ValueReference); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueString" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueString\": ")
	if tmp, err := json.Marshal(strct.ValueString); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ExplanationOfBenefitSupportingInfo) UnmarshalJSON(b []byte) error {
    categoryReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "category":
            if err := json.Unmarshal([]byte(v), &strct.Category); err != nil {
                return err
             }
            categoryReceived = true
        case "code":
            if err := json.Unmarshal([]byte(v), &strct.Code); err != nil {
                return err
             }
        case "extension":
            if err := json.Unmarshal([]byte(v), &strct.Extension); err != nil {
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
        case "reason":
            if err := json.Unmarshal([]byte(v), &strct.Reason); err != nil {
                return err
             }
        case "sequence":
            if err := json.Unmarshal([]byte(v), &strct.Sequence); err != nil {
                return err
             }
        case "timingDate":
            if err := json.Unmarshal([]byte(v), &strct.TimingDate); err != nil {
                return err
             }
        case "timingPeriod":
            if err := json.Unmarshal([]byte(v), &strct.TimingPeriod); err != nil {
                return err
             }
        case "valueAttachment":
            if err := json.Unmarshal([]byte(v), &strct.ValueAttachment); err != nil {
                return err
             }
        case "valueBoolean":
            if err := json.Unmarshal([]byte(v), &strct.ValueBoolean); err != nil {
                return err
             }
        case "valueQuantity":
            if err := json.Unmarshal([]byte(v), &strct.ValueQuantity); err != nil {
                return err
             }
        case "valueReference":
            if err := json.Unmarshal([]byte(v), &strct.ValueReference); err != nil {
                return err
             }
        case "valueString":
            if err := json.Unmarshal([]byte(v), &strct.ValueString); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if category (a required property) was received
    if !categoryReceived {
        return errors.New("\"category\" is required but was not present")
    }
    return nil
}
