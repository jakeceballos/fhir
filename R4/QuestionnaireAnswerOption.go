// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "encoding/json"
    "fmt"
)

// QuestionnaireAnswerOption A structured set of questions intended to guide the collection of answers from end-users. Questionnaires provide detailed control over order, presentation, phraseology and grouping to allow coherent, consistent data collection.
type QuestionnaireAnswerOption struct {

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // Indicates whether the answer value is selected when the list of possible answers is initially shown.
  InitialSelected bool `json:"initialSelected,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // A potential answer that's allowed as the answer to this question.
  ValueCoding *Coding `json:"valueCoding,omitempty"`

  // A potential answer that's allowed as the answer to this question.
  ValueDate string `json:"valueDate,omitempty"`

  // A potential answer that's allowed as the answer to this question.
  ValueInteger float64 `json:"valueInteger,omitempty"`

  // A potential answer that's allowed as the answer to this question.
  ValueReference *Reference `json:"valueReference,omitempty"`

  // A potential answer that's allowed as the answer to this question.
  ValueString string `json:"valueString,omitempty"`

  // A potential answer that's allowed as the answer to this question.
  ValueTime string `json:"valueTime,omitempty"`
}

func (strct *QuestionnaireAnswerOption) MarshalJSON() ([]byte, error) {
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
    // Marshal the "initialSelected" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"initialSelected\": ")
	if tmp, err := json.Marshal(strct.InitialSelected); err != nil {
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
    // Marshal the "valueCoding" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueCoding\": ")
	if tmp, err := json.Marshal(strct.ValueCoding); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueDate" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueDate\": ")
	if tmp, err := json.Marshal(strct.ValueDate); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueInteger" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueInteger\": ")
	if tmp, err := json.Marshal(strct.ValueInteger); err != nil {
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
    // Marshal the "valueString" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueString\": ")
	if tmp, err := json.Marshal(strct.ValueString); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueTime" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueTime\": ")
	if tmp, err := json.Marshal(strct.ValueTime); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *QuestionnaireAnswerOption) UnmarshalJSON(b []byte) error {
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
        case "initialSelected":
            if err := json.Unmarshal([]byte(v), &strct.InitialSelected); err != nil {
                return err
             }
        case "modifierExtension":
            if err := json.Unmarshal([]byte(v), &strct.ModifierExtension); err != nil {
                return err
             }
        case "valueCoding":
            if err := json.Unmarshal([]byte(v), &strct.ValueCoding); err != nil {
                return err
             }
        case "valueDate":
            if err := json.Unmarshal([]byte(v), &strct.ValueDate); err != nil {
                return err
             }
        case "valueInteger":
            if err := json.Unmarshal([]byte(v), &strct.ValueInteger); err != nil {
                return err
             }
        case "valueReference":
            if err := json.Unmarshal([]byte(v), &strct.ValueReference); err != nil {
                return err
             }
        case "valueString":
            if err := json.Unmarshal([]byte(v), &strct.ValueString); err != nil {
                return err
             }
        case "valueTime":
            if err := json.Unmarshal([]byte(v), &strct.ValueTime); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}
