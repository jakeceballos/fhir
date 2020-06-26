// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "encoding/json"
    "fmt"
)

// Expression A expression that is evaluated in a specified context and returns a value. The context of use of the expression must specify the context in which the expression is evaluated, and how the result of the expression is used.
type Expression struct {

  // A brief, natural language description of the condition that effectively communicates the intended semantics.
  Description string `json:"description,omitempty"`

  // An expression in the specified language that returns a value.
  Expression string `json:"expression,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // The media type of the language for the expression.
  Language interface{} `json:"language,omitempty"`

  // A short name assigned to the expression to allow for multiple reuse of the expression in the context where it is defined.
  Name string `json:"name,omitempty"`

  // A URI that defines where the expression is found.
  Reference string `json:"reference,omitempty"`
}

func (strct *Expression) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
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
    // Marshal the "expression" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"expression\": ")
	if tmp, err := json.Marshal(strct.Expression); err != nil {
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
    // Marshal the "language" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"language\": ")
	if tmp, err := json.Marshal(strct.Language); err != nil {
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
    // Marshal the "reference" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"reference\": ")
	if tmp, err := json.Marshal(strct.Reference); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Expression) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
                return err
             }
        case "expression":
            if err := json.Unmarshal([]byte(v), &strct.Expression); err != nil {
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
        case "language":
            if err := json.Unmarshal([]byte(v), &strct.Language); err != nil {
                return err
             }
        case "name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
        case "reference":
            if err := json.Unmarshal([]byte(v), &strct.Reference); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}
