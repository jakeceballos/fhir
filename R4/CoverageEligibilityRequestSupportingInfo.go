// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "errors"
    "encoding/json"
    "fmt"
)

// CoverageEligibilityRequestSupportingInfo The CoverageEligibilityRequest provides patient and insurance coverage information to an insurer for them to respond, in the form of an CoverageEligibilityResponse, with information regarding whether the stated coverage is valid and in-force and optionally to provide the insurance details of the policy.
type CoverageEligibilityRequestSupportingInfo struct {

  // The supporting materials are applicable for all detail items, product/servce categories and specific billing codes.
  AppliesToAll bool `json:"appliesToAll,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // Additional data or information such as resources, documents, images etc. including references to the data or the actual inclusion of the data.
  Information *Reference `json:"information"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // A number to uniquely identify supporting information entries.
  Sequence float64 `json:"sequence,omitempty"`
}

func (strct *CoverageEligibilityRequestSupportingInfo) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "appliesToAll" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"appliesToAll\": ")
	if tmp, err := json.Marshal(strct.AppliesToAll); err != nil {
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
    // "Information" field is required
    if strct.Information == nil {
        return nil, errors.New("information is a required field")
    }
    // Marshal the "information" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"information\": ")
	if tmp, err := json.Marshal(strct.Information); err != nil {
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

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *CoverageEligibilityRequestSupportingInfo) UnmarshalJSON(b []byte) error {
    informationReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "appliesToAll":
            if err := json.Unmarshal([]byte(v), &strct.AppliesToAll); err != nil {
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
        case "information":
            if err := json.Unmarshal([]byte(v), &strct.Information); err != nil {
                return err
             }
            informationReceived = true
        case "modifierExtension":
            if err := json.Unmarshal([]byte(v), &strct.ModifierExtension); err != nil {
                return err
             }
        case "sequence":
            if err := json.Unmarshal([]byte(v), &strct.Sequence); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if information (a required property) was received
    if !informationReceived {
        return errors.New("\"information\" is required but was not present")
    }
    return nil
}