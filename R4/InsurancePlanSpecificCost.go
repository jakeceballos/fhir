// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "errors"
    "encoding/json"
    "fmt"
    "bytes"
)

// InsurancePlanSpecificCost Details of a Health Insurance product/plan provided by an organization.
type InsurancePlanSpecificCost struct {

  // List of the specific benefits under this category of benefit.
  Benefit []*InsurancePlanBenefit1 `json:"benefit,omitempty"`

  // General category of benefit (Medical; Dental; Vision; Drug; Mental Health; Substance Abuse; Hospice, Home Health).
  Category *CodeableConcept `json:"category"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
}

func (strct *InsurancePlanSpecificCost) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "benefit" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"benefit\": ")
	if tmp, err := json.Marshal(strct.Benefit); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
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

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *InsurancePlanSpecificCost) UnmarshalJSON(b []byte) error {
    categoryReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "benefit":
            if err := json.Unmarshal([]byte(v), &strct.Benefit); err != nil {
                return err
             }
        case "category":
            if err := json.Unmarshal([]byte(v), &strct.Category); err != nil {
                return err
             }
            categoryReceived = true
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
