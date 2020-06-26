// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "errors"
    "encoding/json"
    "fmt"
)

// DeviceDefinitionMaterial The characteristics, operational status and capabilities of a medical-related component of a medical device.
type DeviceDefinitionMaterial struct {

  // Whether the substance is a known or suspected allergen.
  AllergenicIndicator bool `json:"allergenicIndicator,omitempty"`

  // Indicates an alternative material of the device.
  Alternate bool `json:"alternate,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // The substance.
  Substance *CodeableConcept `json:"substance"`
}

func (strct *DeviceDefinitionMaterial) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "allergenicIndicator" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"allergenicIndicator\": ")
	if tmp, err := json.Marshal(strct.AllergenicIndicator); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "alternate" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"alternate\": ")
	if tmp, err := json.Marshal(strct.Alternate); err != nil {
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
    // "Substance" field is required
    if strct.Substance == nil {
        return nil, errors.New("substance is a required field")
    }
    // Marshal the "substance" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"substance\": ")
	if tmp, err := json.Marshal(strct.Substance); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *DeviceDefinitionMaterial) UnmarshalJSON(b []byte) error {
    substanceReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "allergenicIndicator":
            if err := json.Unmarshal([]byte(v), &strct.AllergenicIndicator); err != nil {
                return err
             }
        case "alternate":
            if err := json.Unmarshal([]byte(v), &strct.Alternate); err != nil {
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
        case "substance":
            if err := json.Unmarshal([]byte(v), &strct.Substance); err != nil {
                return err
             }
            substanceReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if substance (a required property) was received
    if !substanceReceived {
        return errors.New("\"substance\" is required but was not present")
    }
    return nil
}
