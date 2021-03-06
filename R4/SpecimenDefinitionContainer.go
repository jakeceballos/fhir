// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "encoding/json"
    "fmt"
)

// SpecimenDefinitionContainer A kind of specimen with associated set of requirements.
type SpecimenDefinitionContainer struct {

  // Substance introduced in the kind of container to preserve, maintain or enhance the specimen. Examples: Formalin, Citrate, EDTA.
  Additive []*SpecimenDefinitionAdditive `json:"additive,omitempty"`

  // Color of container cap.
  Cap *CodeableConcept `json:"cap,omitempty"`

  // The capacity (volume or other measure) of this kind of container.
  Capacity *Quantity `json:"capacity,omitempty"`

  // The textual description of the kind of container.
  Description string `json:"description,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // The type of material of the container.
  Material *CodeableConcept `json:"material,omitempty"`

  // The minimum volume to be conditioned in the container.
  MinimumVolumeQuantity *Quantity `json:"minimumVolumeQuantity,omitempty"`

  // The minimum volume to be conditioned in the container.
  MinimumVolumeString string `json:"minimumVolumeString,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // Special processing that should be applied to the container for this kind of specimen.
  Preparation string `json:"preparation,omitempty"`

  // The type of container used to contain this kind of specimen.
  Type *CodeableConcept `json:"type,omitempty"`
}

func (strct *SpecimenDefinitionContainer) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "additive" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"additive\": ")
	if tmp, err := json.Marshal(strct.Additive); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "cap" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"cap\": ")
	if tmp, err := json.Marshal(strct.Cap); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "capacity" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"capacity\": ")
	if tmp, err := json.Marshal(strct.Capacity); err != nil {
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
    // Marshal the "material" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"material\": ")
	if tmp, err := json.Marshal(strct.Material); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "minimumVolumeQuantity" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"minimumVolumeQuantity\": ")
	if tmp, err := json.Marshal(strct.MinimumVolumeQuantity); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "minimumVolumeString" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"minimumVolumeString\": ")
	if tmp, err := json.Marshal(strct.MinimumVolumeString); err != nil {
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
    // Marshal the "preparation" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"preparation\": ")
	if tmp, err := json.Marshal(strct.Preparation); err != nil {
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

func (strct *SpecimenDefinitionContainer) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "additive":
            if err := json.Unmarshal([]byte(v), &strct.Additive); err != nil {
                return err
             }
        case "cap":
            if err := json.Unmarshal([]byte(v), &strct.Cap); err != nil {
                return err
             }
        case "capacity":
            if err := json.Unmarshal([]byte(v), &strct.Capacity); err != nil {
                return err
             }
        case "description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
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
        case "material":
            if err := json.Unmarshal([]byte(v), &strct.Material); err != nil {
                return err
             }
        case "minimumVolumeQuantity":
            if err := json.Unmarshal([]byte(v), &strct.MinimumVolumeQuantity); err != nil {
                return err
             }
        case "minimumVolumeString":
            if err := json.Unmarshal([]byte(v), &strct.MinimumVolumeString); err != nil {
                return err
             }
        case "modifierExtension":
            if err := json.Unmarshal([]byte(v), &strct.ModifierExtension); err != nil {
                return err
             }
        case "preparation":
            if err := json.Unmarshal([]byte(v), &strct.Preparation); err != nil {
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
