// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "fmt"
    "bytes"
    "encoding/json"
)

// Identifier An identifier - identifies some entity uniquely and unambiguously. Typically this is used for business identifiers.
type Identifier struct {

  // Organization that issued/manages the identifier.
  Assigner *Reference `json:"assigner,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // Time period during which identifier is/was valid for use.
  Period *Period `json:"period,omitempty"`

  // Establishes the namespace for the value - that is, a URL that describes a set values that are unique.
  System string `json:"system,omitempty"`

  // A coded type for the identifier that can be used to determine which identifier to use for a specific purpose.
  Type *CodeableConcept `json:"type,omitempty"`

  // The purpose of this identifier.
  Use interface{} `json:"use,omitempty"`

  // The portion of the identifier typically relevant to the user and which is unique within the context of the system.
  Value string `json:"value,omitempty"`
}

func (strct *Identifier) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "assigner" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"assigner\": ")
	if tmp, err := json.Marshal(strct.Assigner); err != nil {
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
    // Marshal the "period" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"period\": ")
	if tmp, err := json.Marshal(strct.Period); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "system" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"system\": ")
	if tmp, err := json.Marshal(strct.System); err != nil {
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
    // Marshal the "use" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"use\": ")
	if tmp, err := json.Marshal(strct.Use); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
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

func (strct *Identifier) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "assigner":
            if err := json.Unmarshal([]byte(v), &strct.Assigner); err != nil {
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
        case "period":
            if err := json.Unmarshal([]byte(v), &strct.Period); err != nil {
                return err
             }
        case "system":
            if err := json.Unmarshal([]byte(v), &strct.System); err != nil {
                return err
             }
        case "type":
            if err := json.Unmarshal([]byte(v), &strct.Type); err != nil {
                return err
             }
        case "use":
            if err := json.Unmarshal([]byte(v), &strct.Use); err != nil {
                return err
             }
        case "value":
            if err := json.Unmarshal([]byte(v), &strct.Value); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}