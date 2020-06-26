// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "fmt"
    "bytes"
    "encoding/json"
)

// QuestionnaireResponseAnswer A structured set of questions and their answers. The questions are ordered and grouped into coherent subsets, corresponding to the structure of the grouping of the questionnaire being responded to.
type QuestionnaireResponseAnswer struct {

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // Nested groups and/or questions found within this particular answer.
  Item []*QuestionnaireResponseItem `json:"item,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // The answer (or one of the answers) provided by the respondent to the question.
  ValueAttachment *Attachment `json:"valueAttachment,omitempty"`

  // The answer (or one of the answers) provided by the respondent to the question.
  ValueBoolean bool `json:"valueBoolean,omitempty"`

  // The answer (or one of the answers) provided by the respondent to the question.
  ValueCoding *Coding `json:"valueCoding,omitempty"`

  // The answer (or one of the answers) provided by the respondent to the question.
  ValueDate string `json:"valueDate,omitempty"`

  // The answer (or one of the answers) provided by the respondent to the question.
  ValueDateTime string `json:"valueDateTime,omitempty"`

  // The answer (or one of the answers) provided by the respondent to the question.
  ValueDecimal float64 `json:"valueDecimal,omitempty"`

  // The answer (or one of the answers) provided by the respondent to the question.
  ValueInteger float64 `json:"valueInteger,omitempty"`

  // The answer (or one of the answers) provided by the respondent to the question.
  ValueQuantity *Quantity `json:"valueQuantity,omitempty"`

  // The answer (or one of the answers) provided by the respondent to the question.
  ValueReference *Reference `json:"valueReference,omitempty"`

  // The answer (or one of the answers) provided by the respondent to the question.
  ValueString string `json:"valueString,omitempty"`

  // The answer (or one of the answers) provided by the respondent to the question.
  ValueTime string `json:"valueTime,omitempty"`

  // The answer (or one of the answers) provided by the respondent to the question.
  ValueUri string `json:"valueUri,omitempty"`
}

func (strct *QuestionnaireResponseAnswer) MarshalJSON() ([]byte, error) {
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
    // Marshal the "item" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"item\": ")
	if tmp, err := json.Marshal(strct.Item); err != nil {
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
    // Marshal the "valueAttachment" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueAttachment\": ")
	if tmp, err := json.Marshal(strct.ValueAttachment); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueBoolean" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueBoolean\": ")
	if tmp, err := json.Marshal(strct.ValueBoolean); err != nil {
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
    // Marshal the "valueDateTime" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueDateTime\": ")
	if tmp, err := json.Marshal(strct.ValueDateTime); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueDecimal" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueDecimal\": ")
	if tmp, err := json.Marshal(strct.ValueDecimal); err != nil {
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
    // Marshal the "valueUri" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueUri\": ")
	if tmp, err := json.Marshal(strct.ValueUri); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *QuestionnaireResponseAnswer) UnmarshalJSON(b []byte) error {
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
        case "item":
            if err := json.Unmarshal([]byte(v), &strct.Item); err != nil {
                return err
             }
        case "modifierExtension":
            if err := json.Unmarshal([]byte(v), &strct.ModifierExtension); err != nil {
                return err
             }
        case "valueAttachment":
            if err := json.Unmarshal([]byte(v), &strct.ValueAttachment); err != nil {
                return err
             }
        case "valueBoolean":
            if err := json.Unmarshal([]byte(v), &strct.ValueBoolean); err != nil {
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
        case "valueDateTime":
            if err := json.Unmarshal([]byte(v), &strct.ValueDateTime); err != nil {
                return err
             }
        case "valueDecimal":
            if err := json.Unmarshal([]byte(v), &strct.ValueDecimal); err != nil {
                return err
             }
        case "valueInteger":
            if err := json.Unmarshal([]byte(v), &strct.ValueInteger); err != nil {
                return err
             }
        case "valueQuantity":
            if err := json.Unmarshal([]byte(v), &strct.ValueQuantity); err != nil {
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
        case "valueUri":
            if err := json.Unmarshal([]byte(v), &strct.ValueUri); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}