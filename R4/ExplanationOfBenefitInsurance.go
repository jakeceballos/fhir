// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "errors"
    "encoding/json"
    "fmt"
)

// ExplanationOfBenefitInsurance This resource provides: the claim details; adjudication details from the processing of a Claim; and optionally account balance information, for informing the subscriber of the benefits provided.
type ExplanationOfBenefitInsurance struct {

  // Reference to the insurance card level information contained in the Coverage resource. The coverage issuing insurer will use these details to locate the patient's actual coverage within the insurer's information system.
  Coverage *Reference `json:"coverage"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // A flag to indicate that this Coverage is to be used for adjudication of this claim when set to true.
  Focal bool `json:"focal,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // Reference numbers previously provided by the insurer to the provider to be quoted on subsequent claims containing services or products related to the prior authorization.
  PreAuthRef []string `json:"preAuthRef,omitempty"`
}

func (strct *ExplanationOfBenefitInsurance) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "Coverage" field is required
    if strct.Coverage == nil {
        return nil, errors.New("coverage is a required field")
    }
    // Marshal the "coverage" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"coverage\": ")
	if tmp, err := json.Marshal(strct.Coverage); err != nil {
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
    // Marshal the "focal" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"focal\": ")
	if tmp, err := json.Marshal(strct.Focal); err != nil {
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
    // Marshal the "preAuthRef" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"preAuthRef\": ")
	if tmp, err := json.Marshal(strct.PreAuthRef); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ExplanationOfBenefitInsurance) UnmarshalJSON(b []byte) error {
    coverageReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "coverage":
            if err := json.Unmarshal([]byte(v), &strct.Coverage); err != nil {
                return err
             }
            coverageReceived = true
        case "extension":
            if err := json.Unmarshal([]byte(v), &strct.Extension); err != nil {
                return err
             }
        case "focal":
            if err := json.Unmarshal([]byte(v), &strct.Focal); err != nil {
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
        case "preAuthRef":
            if err := json.Unmarshal([]byte(v), &strct.PreAuthRef); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if coverage (a required property) was received
    if !coverageReceived {
        return errors.New("\"coverage\" is required but was not present")
    }
    return nil
}
