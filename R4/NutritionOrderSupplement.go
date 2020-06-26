// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "encoding/json"
    "fmt"
)

// NutritionOrderSupplement A request to supply a diet, formula feeding (enteral) or oral nutritional supplement to a patient/resident.
type NutritionOrderSupplement struct {

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // Free text or additional instructions or information pertaining to the oral supplement.
  Instruction string `json:"instruction,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // The product or brand name of the nutritional supplement such as "Acme Protein Shake".
  ProductName string `json:"productName,omitempty"`

  // The amount of the nutritional supplement to be given.
  Quantity *Quantity `json:"quantity,omitempty"`

  // The time period and frequency at which the supplement(s) should be given.  The supplement should be given for the combination of all schedules if more than one schedule is present.
  Schedule []*Timing `json:"schedule,omitempty"`

  // The kind of nutritional supplement product required such as a high protein or pediatric clear liquid supplement.
  Type *CodeableConcept `json:"type,omitempty"`
}

func (strct *NutritionOrderSupplement) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
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
    // Marshal the "instruction" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"instruction\": ")
	if tmp, err := json.Marshal(strct.Instruction); err != nil {
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
    // Marshal the "productName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"productName\": ")
	if tmp, err := json.Marshal(strct.ProductName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "quantity" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"quantity\": ")
	if tmp, err := json.Marshal(strct.Quantity); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "schedule" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"schedule\": ")
	if tmp, err := json.Marshal(strct.Schedule); err != nil {
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

func (strct *NutritionOrderSupplement) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "extension":
            if err := json.Unmarshal([]byte(v), &strct.Extension); err != nil {
                return err
             }
        case "id":
            if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
                return err
             }
        case "instruction":
            if err := json.Unmarshal([]byte(v), &strct.Instruction); err != nil {
                return err
             }
        case "modifierExtension":
            if err := json.Unmarshal([]byte(v), &strct.ModifierExtension); err != nil {
                return err
             }
        case "productName":
            if err := json.Unmarshal([]byte(v), &strct.ProductName); err != nil {
                return err
             }
        case "quantity":
            if err := json.Unmarshal([]byte(v), &strct.Quantity); err != nil {
                return err
             }
        case "schedule":
            if err := json.Unmarshal([]byte(v), &strct.Schedule); err != nil {
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
