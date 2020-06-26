// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "encoding/json"
    "fmt"
)

// LocationHoursOfOperation Details and position information for a physical place where services are provided and resources and participants may be stored, found, contained, or accommodated.
type LocationHoursOfOperation struct {

  // The Location is open all day.
  AllDay bool `json:"allDay,omitempty"`

  // Time that the Location closes.
  ClosingTime string `json:"closingTime,omitempty"`

  // Indicates which days of the week are available between the start and end Times.
  DaysOfWeek []string `json:"daysOfWeek,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // Time that the Location opens.
  OpeningTime string `json:"openingTime,omitempty"`
}

func (strct *LocationHoursOfOperation) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "allDay" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"allDay\": ")
	if tmp, err := json.Marshal(strct.AllDay); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "closingTime" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"closingTime\": ")
	if tmp, err := json.Marshal(strct.ClosingTime); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "daysOfWeek" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"daysOfWeek\": ")
	if tmp, err := json.Marshal(strct.DaysOfWeek); err != nil {
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
    // Marshal the "openingTime" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"openingTime\": ")
	if tmp, err := json.Marshal(strct.OpeningTime); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *LocationHoursOfOperation) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "allDay":
            if err := json.Unmarshal([]byte(v), &strct.AllDay); err != nil {
                return err
             }
        case "closingTime":
            if err := json.Unmarshal([]byte(v), &strct.ClosingTime); err != nil {
                return err
             }
        case "daysOfWeek":
            if err := json.Unmarshal([]byte(v), &strct.DaysOfWeek); err != nil {
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
        case "openingTime":
            if err := json.Unmarshal([]byte(v), &strct.OpeningTime); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}
