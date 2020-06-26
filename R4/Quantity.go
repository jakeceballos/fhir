// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "fmt"
    "bytes"
    "encoding/json"
)

// Quantity A measured amount (or an amount that can potentially be measured). Note that measured amounts include amounts that are not precisely quantified, including amounts involving arbitrary units and floating currencies.
type Quantity struct {

  // A computer processable form of the unit in some unit representation system.
  Code string `json:"code,omitempty"`

  // How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues; e.g. if the comparator is "<" , then the real value is < stated value.
  Comparator interface{} `json:"comparator,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // The identification of the system that provides the coded form of the unit.
  System string `json:"system,omitempty"`

  // A human-readable form of the unit.
  Unit string `json:"unit,omitempty"`

  // The value of the measured amount. The value includes an implicit precision in the presentation of the value.
  Value float64 `json:"value,omitempty"`
}

func (strct *Quantity) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
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
    // Marshal the "comparator" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"comparator\": ")
	if tmp, err := json.Marshal(strct.Comparator); err != nil {
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
    // Marshal the "unit" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"unit\": ")
	if tmp, err := json.Marshal(strct.Unit); err != nil {
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

func (strct *Quantity) UnmarshalJSON(b []byte) error {
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
        case "comparator":
            if err := json.Unmarshal([]byte(v), &strct.Comparator); err != nil {
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
        case "system":
            if err := json.Unmarshal([]byte(v), &strct.System); err != nil {
                return err
             }
        case "unit":
            if err := json.Unmarshal([]byte(v), &strct.Unit); err != nil {
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
