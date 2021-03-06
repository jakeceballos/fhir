// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "encoding/json"
    "fmt"
)

// ParameterDefinition The parameters to the module. This collection specifies both the input and output parameters. Input parameters are provided by the caller as part of the $evaluate operation. Output parameters are included in the GuidanceResponse.
type ParameterDefinition struct {

  // A brief discussion of what the parameter is for and how it is used by the module.
  Documentation string `json:"documentation,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // The maximum number of times this element is permitted to appear in the request or response.
  Max string `json:"max,omitempty"`

  // The minimum number of times this parameter SHALL appear in the request or response.
  Min float64 `json:"min,omitempty"`

  // The name of the parameter used to allow access to the value of the parameter in evaluation contexts.
  Name string `json:"name,omitempty"`

  // If specified, this indicates a profile that the input data must conform to, or that the output data will conform to.
  Profile string `json:"profile,omitempty"`

  // The type of the parameter.
  Type string `json:"type,omitempty"`

  // Whether the parameter is input or output for the module.
  Use string `json:"use,omitempty"`
}

func (strct *ParameterDefinition) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "documentation" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"documentation\": ")
	if tmp, err := json.Marshal(strct.Documentation); err != nil {
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
    // Marshal the "max" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"max\": ")
	if tmp, err := json.Marshal(strct.Max); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "min" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"min\": ")
	if tmp, err := json.Marshal(strct.Min); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "name" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "profile" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"profile\": ")
	if tmp, err := json.Marshal(strct.Profile); err != nil {
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

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ParameterDefinition) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "documentation":
            if err := json.Unmarshal([]byte(v), &strct.Documentation); err != nil {
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
        case "max":
            if err := json.Unmarshal([]byte(v), &strct.Max); err != nil {
                return err
             }
        case "min":
            if err := json.Unmarshal([]byte(v), &strct.Min); err != nil {
                return err
             }
        case "name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
        case "profile":
            if err := json.Unmarshal([]byte(v), &strct.Profile); err != nil {
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
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}
