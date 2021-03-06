// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "fmt"
    "bytes"
    "errors"
    "encoding/json"
)

// MedicinalProductPharmaceuticalWithdrawalPeriod A pharmaceutical product described in terms of its composition and dose form.
type MedicinalProductPharmaceuticalWithdrawalPeriod struct {

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // Extra information about the withdrawal period.
  SupportingInformation string `json:"supportingInformation,omitempty"`

  // Coded expression for the type of tissue for which the withdrawal period applues, e.g. meat, milk.
  Tissue *CodeableConcept `json:"tissue"`

  // A value for the time.
  Value *Quantity `json:"value"`
}

func (strct *MedicinalProductPharmaceuticalWithdrawalPeriod) MarshalJSON() ([]byte, error) {
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
    // Marshal the "supportingInformation" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"supportingInformation\": ")
	if tmp, err := json.Marshal(strct.SupportingInformation); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Tissue" field is required
    if strct.Tissue == nil {
        return nil, errors.New("tissue is a required field")
    }
    // Marshal the "tissue" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"tissue\": ")
	if tmp, err := json.Marshal(strct.Tissue); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Value" field is required
    if strct.Value == nil {
        return nil, errors.New("value is a required field")
    }
    // Marshal the "value" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"value\": ")
	if tmp, err := json.Marshal(strct.Value); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *MedicinalProductPharmaceuticalWithdrawalPeriod) UnmarshalJSON(b []byte) error {
    tissueReceived := false
    valueReceived := false
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
        case "modifierExtension":
            if err := json.Unmarshal([]byte(v), &strct.ModifierExtension); err != nil {
                return err
             }
        case "supportingInformation":
            if err := json.Unmarshal([]byte(v), &strct.SupportingInformation); err != nil {
                return err
             }
        case "tissue":
            if err := json.Unmarshal([]byte(v), &strct.Tissue); err != nil {
                return err
             }
            tissueReceived = true
        case "value":
            if err := json.Unmarshal([]byte(v), &strct.Value); err != nil {
                return err
             }
            valueReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if tissue (a required property) was received
    if !tissueReceived {
        return errors.New("\"tissue\" is required but was not present")
    }
    // check if value (a required property) was received
    if !valueReceived {
        return errors.New("\"value\" is required but was not present")
    }
    return nil
}
