// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "encoding/json"
    "fmt"
)

// MolecularSequenceStructureVariant Raw data describing a biological sequence.
type MolecularSequenceStructureVariant struct {

  // Used to indicate if the outer and inner start-end values have the same meaning.
  Exact bool `json:"exact,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // Structural variant inner.
  Inner *MolecularSequenceInner `json:"inner,omitempty"`

  // Length of the variant chromosome.
  Length float64 `json:"length,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // Structural variant outer.
  Outer *MolecularSequenceOuter `json:"outer,omitempty"`

  // Information about chromosome structure variation DNA change type.
  VariantType *CodeableConcept `json:"variantType,omitempty"`
}

func (strct *MolecularSequenceStructureVariant) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "exact" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"exact\": ")
	if tmp, err := json.Marshal(strct.Exact); err != nil {
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
    // Marshal the "inner" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"inner\": ")
	if tmp, err := json.Marshal(strct.Inner); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "length" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"length\": ")
	if tmp, err := json.Marshal(strct.Length); err != nil {
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
    // Marshal the "outer" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"outer\": ")
	if tmp, err := json.Marshal(strct.Outer); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "variantType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"variantType\": ")
	if tmp, err := json.Marshal(strct.VariantType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *MolecularSequenceStructureVariant) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "exact":
            if err := json.Unmarshal([]byte(v), &strct.Exact); err != nil {
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
        case "inner":
            if err := json.Unmarshal([]byte(v), &strct.Inner); err != nil {
                return err
             }
        case "length":
            if err := json.Unmarshal([]byte(v), &strct.Length); err != nil {
                return err
             }
        case "modifierExtension":
            if err := json.Unmarshal([]byte(v), &strct.ModifierExtension); err != nil {
                return err
             }
        case "outer":
            if err := json.Unmarshal([]byte(v), &strct.Outer); err != nil {
                return err
             }
        case "variantType":
            if err := json.Unmarshal([]byte(v), &strct.VariantType); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}
