// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "errors"
    "encoding/json"
    "fmt"
)

// MedicinalProductIngredientReferenceStrength An ingredient of a manufactured item or pharmaceutical product.
type MedicinalProductIngredientReferenceStrength struct {

  // The country or countries for which the strength range applies.
  Country []*CodeableConcept `json:"country,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // For when strength is measured at a particular point or distance.
  MeasurementPoint string `json:"measurementPoint,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // Strength expressed in terms of a reference substance.
  Strength *Ratio `json:"strength"`

  // Strength expressed in terms of a reference substance.
  StrengthLowLimit *Ratio `json:"strengthLowLimit,omitempty"`

  // Relevant reference substance.
  Substance *CodeableConcept `json:"substance,omitempty"`
}

func (strct *MedicinalProductIngredientReferenceStrength) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "country" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"country\": ")
	if tmp, err := json.Marshal(strct.Country); err != nil {
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
    // Marshal the "measurementPoint" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"measurementPoint\": ")
	if tmp, err := json.Marshal(strct.MeasurementPoint); err != nil {
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
    // "Strength" field is required
    if strct.Strength == nil {
        return nil, errors.New("strength is a required field")
    }
    // Marshal the "strength" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"strength\": ")
	if tmp, err := json.Marshal(strct.Strength); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "strengthLowLimit" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"strengthLowLimit\": ")
	if tmp, err := json.Marshal(strct.StrengthLowLimit); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "substance" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"substance\": ")
	if tmp, err := json.Marshal(strct.Substance); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *MedicinalProductIngredientReferenceStrength) UnmarshalJSON(b []byte) error {
    strengthReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "country":
            if err := json.Unmarshal([]byte(v), &strct.Country); err != nil {
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
        case "measurementPoint":
            if err := json.Unmarshal([]byte(v), &strct.MeasurementPoint); err != nil {
                return err
             }
        case "modifierExtension":
            if err := json.Unmarshal([]byte(v), &strct.ModifierExtension); err != nil {
                return err
             }
        case "strength":
            if err := json.Unmarshal([]byte(v), &strct.Strength); err != nil {
                return err
             }
            strengthReceived = true
        case "strengthLowLimit":
            if err := json.Unmarshal([]byte(v), &strct.StrengthLowLimit); err != nil {
                return err
             }
        case "substance":
            if err := json.Unmarshal([]byte(v), &strct.Substance); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if strength (a required property) was received
    if !strengthReceived {
        return errors.New("\"strength\" is required but was not present")
    }
    return nil
}