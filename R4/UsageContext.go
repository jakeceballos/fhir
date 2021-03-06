// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "errors"
    "encoding/json"
    "fmt"
)

// UsageContext Specifies clinical/business/etc. metadata that can be used to retrieve, index and/or categorize an artifact. This metadata can either be specific to the applicable population (e.g., age category, DRG) or the specific context of care (e.g., venue, care setting, provider of care).
type UsageContext struct {

  // A code that identifies the type of context being specified by this usage context.
  Code *Coding `json:"code"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // A value that defines the context specified in this context of use. The interpretation of the value is defined by the code.
  ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`

  // A value that defines the context specified in this context of use. The interpretation of the value is defined by the code.
  ValueQuantity *Quantity `json:"valueQuantity,omitempty"`

  // A value that defines the context specified in this context of use. The interpretation of the value is defined by the code.
  ValueRange *Range `json:"valueRange,omitempty"`

  // A value that defines the context specified in this context of use. The interpretation of the value is defined by the code.
  ValueReference *Reference `json:"valueReference,omitempty"`
}

func (strct *UsageContext) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "Code" field is required
    if strct.Code == nil {
        return nil, errors.New("code is a required field")
    }
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
    // Marshal the "valueCodeableConcept" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueCodeableConcept\": ")
	if tmp, err := json.Marshal(strct.ValueCodeableConcept); err != nil {
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
    // Marshal the "valueRange" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueRange\": ")
	if tmp, err := json.Marshal(strct.ValueRange); err != nil {
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

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *UsageContext) UnmarshalJSON(b []byte) error {
    codeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "code":
            if err := json.Unmarshal([]byte(v), &strct.Code); err != nil {
                return err
             }
            codeReceived = true
        case "extension":
            if err := json.Unmarshal([]byte(v), &strct.Extension); err != nil {
                return err
             }
        case "id":
            if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
                return err
             }
        case "valueCodeableConcept":
            if err := json.Unmarshal([]byte(v), &strct.ValueCodeableConcept); err != nil {
                return err
             }
        case "valueQuantity":
            if err := json.Unmarshal([]byte(v), &strct.ValueQuantity); err != nil {
                return err
             }
        case "valueRange":
            if err := json.Unmarshal([]byte(v), &strct.ValueRange); err != nil {
                return err
             }
        case "valueReference":
            if err := json.Unmarshal([]byte(v), &strct.ValueReference); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if code (a required property) was received
    if !codeReceived {
        return errors.New("\"code\" is required but was not present")
    }
    return nil
}
