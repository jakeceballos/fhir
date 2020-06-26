// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "errors"
    "encoding/json"
    "fmt"
)

// ObservationComponent Measurements and simple assertions made about a patient, device or other subject.
type ObservationComponent struct {

  // Describes what was observed. Sometimes this is called the observation "code".
  Code *CodeableConcept `json:"code"`

  // Provides a reason why the expected value in the element Observation.component.value[x] is missing.
  DataAbsentReason *CodeableConcept `json:"dataAbsentReason,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // A categorical assessment of an observation value.  For example, high, low, normal.
  Interpretation []*CodeableConcept `json:"interpretation,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // Guidance on how to interpret the value by comparison to a normal or recommended range.
  ReferenceRange []*ObservationReferenceRange `json:"referenceRange,omitempty"`

  // The information determined as a result of making the observation, if the information has a simple value.
  ValueBoolean bool `json:"valueBoolean,omitempty"`

  // The information determined as a result of making the observation, if the information has a simple value.
  ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`

  // The information determined as a result of making the observation, if the information has a simple value.
  ValueDateTime string `json:"valueDateTime,omitempty"`

  // The information determined as a result of making the observation, if the information has a simple value.
  ValueInteger float64 `json:"valueInteger,omitempty"`

  // The information determined as a result of making the observation, if the information has a simple value.
  ValuePeriod *Period `json:"valuePeriod,omitempty"`

  // The information determined as a result of making the observation, if the information has a simple value.
  ValueQuantity *Quantity `json:"valueQuantity,omitempty"`

  // The information determined as a result of making the observation, if the information has a simple value.
  ValueRange *Range `json:"valueRange,omitempty"`

  // The information determined as a result of making the observation, if the information has a simple value.
  ValueRatio *Ratio `json:"valueRatio,omitempty"`

  // The information determined as a result of making the observation, if the information has a simple value.
  ValueSampledData *SampledData `json:"valueSampledData,omitempty"`

  // The information determined as a result of making the observation, if the information has a simple value.
  ValueString string `json:"valueString,omitempty"`

  // The information determined as a result of making the observation, if the information has a simple value.
  ValueTime string `json:"valueTime,omitempty"`
}

func (strct *ObservationComponent) MarshalJSON() ([]byte, error) {
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
    // Marshal the "dataAbsentReason" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"dataAbsentReason\": ")
	if tmp, err := json.Marshal(strct.DataAbsentReason); err != nil {
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
    // Marshal the "interpretation" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"interpretation\": ")
	if tmp, err := json.Marshal(strct.Interpretation); err != nil {
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
    // Marshal the "referenceRange" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"referenceRange\": ")
	if tmp, err := json.Marshal(strct.ReferenceRange); err != nil {
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
    // Marshal the "valuePeriod" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valuePeriod\": ")
	if tmp, err := json.Marshal(strct.ValuePeriod); err != nil {
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
    // Marshal the "valueRatio" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueRatio\": ")
	if tmp, err := json.Marshal(strct.ValueRatio); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueSampledData" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueSampledData\": ")
	if tmp, err := json.Marshal(strct.ValueSampledData); err != nil {
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

func (strct *ObservationComponent) UnmarshalJSON(b []byte) error {
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
        case "dataAbsentReason":
            if err := json.Unmarshal([]byte(v), &strct.DataAbsentReason); err != nil {
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
        case "interpretation":
            if err := json.Unmarshal([]byte(v), &strct.Interpretation); err != nil {
                return err
             }
        case "modifierExtension":
            if err := json.Unmarshal([]byte(v), &strct.ModifierExtension); err != nil {
                return err
             }
        case "referenceRange":
            if err := json.Unmarshal([]byte(v), &strct.ReferenceRange); err != nil {
                return err
             }
        case "valueBoolean":
            if err := json.Unmarshal([]byte(v), &strct.ValueBoolean); err != nil {
                return err
             }
        case "valueCodeableConcept":
            if err := json.Unmarshal([]byte(v), &strct.ValueCodeableConcept); err != nil {
                return err
             }
        case "valueDateTime":
            if err := json.Unmarshal([]byte(v), &strct.ValueDateTime); err != nil {
                return err
             }
        case "valueInteger":
            if err := json.Unmarshal([]byte(v), &strct.ValueInteger); err != nil {
                return err
             }
        case "valuePeriod":
            if err := json.Unmarshal([]byte(v), &strct.ValuePeriod); err != nil {
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
        case "valueRatio":
            if err := json.Unmarshal([]byte(v), &strct.ValueRatio); err != nil {
                return err
             }
        case "valueSampledData":
            if err := json.Unmarshal([]byte(v), &strct.ValueSampledData); err != nil {
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
    // check if code (a required property) was received
    if !codeReceived {
        return errors.New("\"code\" is required but was not present")
    }
    return nil
}
