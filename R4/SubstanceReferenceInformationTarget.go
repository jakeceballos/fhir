// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "encoding/json"
    "fmt"
)

// SubstanceReferenceInformationTarget Todo.
type SubstanceReferenceInformationTarget struct {

  // Todo.
  AmountQuantity *Quantity `json:"amountQuantity,omitempty"`

  // Todo.
  AmountRange *Range `json:"amountRange,omitempty"`

  // Todo.
  AmountString string `json:"amountString,omitempty"`

  // Todo.
  AmountType *CodeableConcept `json:"amountType,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // Todo.
  Interaction *CodeableConcept `json:"interaction,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // Todo.
  Organism *CodeableConcept `json:"organism,omitempty"`

  // Todo.
  OrganismType *CodeableConcept `json:"organismType,omitempty"`

  // Todo.
  Source []*Reference `json:"source,omitempty"`

  // Todo.
  Target *Identifier `json:"target,omitempty"`

  // Todo.
  Type *CodeableConcept `json:"type,omitempty"`
}

func (strct *SubstanceReferenceInformationTarget) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "amountQuantity" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"amountQuantity\": ")
	if tmp, err := json.Marshal(strct.AmountQuantity); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "amountRange" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"amountRange\": ")
	if tmp, err := json.Marshal(strct.AmountRange); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "amountString" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"amountString\": ")
	if tmp, err := json.Marshal(strct.AmountString); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "amountType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"amountType\": ")
	if tmp, err := json.Marshal(strct.AmountType); err != nil {
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
    // Marshal the "interaction" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"interaction\": ")
	if tmp, err := json.Marshal(strct.Interaction); err != nil {
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
    // Marshal the "organism" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"organism\": ")
	if tmp, err := json.Marshal(strct.Organism); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "organismType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"organismType\": ")
	if tmp, err := json.Marshal(strct.OrganismType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "source" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"source\": ")
	if tmp, err := json.Marshal(strct.Source); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "target" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"target\": ")
	if tmp, err := json.Marshal(strct.Target); err != nil {
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

func (strct *SubstanceReferenceInformationTarget) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "amountQuantity":
            if err := json.Unmarshal([]byte(v), &strct.AmountQuantity); err != nil {
                return err
             }
        case "amountRange":
            if err := json.Unmarshal([]byte(v), &strct.AmountRange); err != nil {
                return err
             }
        case "amountString":
            if err := json.Unmarshal([]byte(v), &strct.AmountString); err != nil {
                return err
             }
        case "amountType":
            if err := json.Unmarshal([]byte(v), &strct.AmountType); err != nil {
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
        case "interaction":
            if err := json.Unmarshal([]byte(v), &strct.Interaction); err != nil {
                return err
             }
        case "modifierExtension":
            if err := json.Unmarshal([]byte(v), &strct.ModifierExtension); err != nil {
                return err
             }
        case "organism":
            if err := json.Unmarshal([]byte(v), &strct.Organism); err != nil {
                return err
             }
        case "organismType":
            if err := json.Unmarshal([]byte(v), &strct.OrganismType); err != nil {
                return err
             }
        case "source":
            if err := json.Unmarshal([]byte(v), &strct.Source); err != nil {
                return err
             }
        case "target":
            if err := json.Unmarshal([]byte(v), &strct.Target); err != nil {
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
