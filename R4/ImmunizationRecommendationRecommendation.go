// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "encoding/json"
    "fmt"
    "bytes"
    "errors"
)

// ImmunizationRecommendationRecommendation A patient's point-in-time set of recommendations (i.e. forecasting) according to a published schedule with optional supporting justification.
type ImmunizationRecommendationRecommendation struct {

  // Vaccine(s) which should not be used to fulfill the recommendation.
  ContraindicatedVaccineCode []*CodeableConcept `json:"contraindicatedVaccineCode,omitempty"`

  // Vaccine date recommendations.  For example, earliest date to administer, latest date to administer, etc.
  DateCriterion []*ImmunizationRecommendationDateCriterion `json:"dateCriterion,omitempty"`

  // Contains the description about the protocol under which the vaccine was administered.
  Description string `json:"description,omitempty"`

  // Nominal position of the recommended dose in a series (e.g. dose 2 is the next recommended dose).
  DoseNumberPositiveInt float64 `json:"doseNumberPositiveInt,omitempty"`

  // Nominal position of the recommended dose in a series (e.g. dose 2 is the next recommended dose).
  DoseNumberString string `json:"doseNumberString,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // The reason for the assigned forecast status.
  ForecastReason []*CodeableConcept `json:"forecastReason,omitempty"`

  // Indicates the patient status with respect to the path to immunity for the target disease.
  ForecastStatus *CodeableConcept `json:"forecastStatus"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // One possible path to achieve presumed immunity against a disease - within the context of an authority.
  Series string `json:"series,omitempty"`

  // The recommended number of doses to achieve immunity.
  SeriesDosesPositiveInt float64 `json:"seriesDosesPositiveInt,omitempty"`

  // The recommended number of doses to achieve immunity.
  SeriesDosesString string `json:"seriesDosesString,omitempty"`

  // Immunization event history and/or evaluation that supports the status and recommendation.
  SupportingImmunization []*Reference `json:"supportingImmunization,omitempty"`

  // Patient Information that supports the status and recommendation.  This includes patient observations, adverse reactions and allergy/intolerance information.
  SupportingPatientInformation []*Reference `json:"supportingPatientInformation,omitempty"`

  // The targeted disease for the recommendation.
  TargetDisease *CodeableConcept `json:"targetDisease,omitempty"`

  // Vaccine(s) or vaccine group that pertain to the recommendation.
  VaccineCode []*CodeableConcept `json:"vaccineCode,omitempty"`
}

func (strct *ImmunizationRecommendationRecommendation) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "contraindicatedVaccineCode" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"contraindicatedVaccineCode\": ")
	if tmp, err := json.Marshal(strct.ContraindicatedVaccineCode); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "dateCriterion" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"dateCriterion\": ")
	if tmp, err := json.Marshal(strct.DateCriterion); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
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
    // Marshal the "doseNumberPositiveInt" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"doseNumberPositiveInt\": ")
	if tmp, err := json.Marshal(strct.DoseNumberPositiveInt); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "doseNumberString" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"doseNumberString\": ")
	if tmp, err := json.Marshal(strct.DoseNumberString); err != nil {
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
    // Marshal the "forecastReason" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"forecastReason\": ")
	if tmp, err := json.Marshal(strct.ForecastReason); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "ForecastStatus" field is required
    if strct.ForecastStatus == nil {
        return nil, errors.New("forecastStatus is a required field")
    }
    // Marshal the "forecastStatus" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"forecastStatus\": ")
	if tmp, err := json.Marshal(strct.ForecastStatus); err != nil {
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
    // Marshal the "series" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"series\": ")
	if tmp, err := json.Marshal(strct.Series); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "seriesDosesPositiveInt" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"seriesDosesPositiveInt\": ")
	if tmp, err := json.Marshal(strct.SeriesDosesPositiveInt); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "seriesDosesString" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"seriesDosesString\": ")
	if tmp, err := json.Marshal(strct.SeriesDosesString); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "supportingImmunization" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"supportingImmunization\": ")
	if tmp, err := json.Marshal(strct.SupportingImmunization); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "supportingPatientInformation" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"supportingPatientInformation\": ")
	if tmp, err := json.Marshal(strct.SupportingPatientInformation); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "targetDisease" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"targetDisease\": ")
	if tmp, err := json.Marshal(strct.TargetDisease); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "vaccineCode" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"vaccineCode\": ")
	if tmp, err := json.Marshal(strct.VaccineCode); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ImmunizationRecommendationRecommendation) UnmarshalJSON(b []byte) error {
    forecastStatusReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "contraindicatedVaccineCode":
            if err := json.Unmarshal([]byte(v), &strct.ContraindicatedVaccineCode); err != nil {
                return err
             }
        case "dateCriterion":
            if err := json.Unmarshal([]byte(v), &strct.DateCriterion); err != nil {
                return err
             }
        case "description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
                return err
             }
        case "doseNumberPositiveInt":
            if err := json.Unmarshal([]byte(v), &strct.DoseNumberPositiveInt); err != nil {
                return err
             }
        case "doseNumberString":
            if err := json.Unmarshal([]byte(v), &strct.DoseNumberString); err != nil {
                return err
             }
        case "extension":
            if err := json.Unmarshal([]byte(v), &strct.Extension); err != nil {
                return err
             }
        case "forecastReason":
            if err := json.Unmarshal([]byte(v), &strct.ForecastReason); err != nil {
                return err
             }
        case "forecastStatus":
            if err := json.Unmarshal([]byte(v), &strct.ForecastStatus); err != nil {
                return err
             }
            forecastStatusReceived = true
        case "id":
            if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
                return err
             }
        case "modifierExtension":
            if err := json.Unmarshal([]byte(v), &strct.ModifierExtension); err != nil {
                return err
             }
        case "series":
            if err := json.Unmarshal([]byte(v), &strct.Series); err != nil {
                return err
             }
        case "seriesDosesPositiveInt":
            if err := json.Unmarshal([]byte(v), &strct.SeriesDosesPositiveInt); err != nil {
                return err
             }
        case "seriesDosesString":
            if err := json.Unmarshal([]byte(v), &strct.SeriesDosesString); err != nil {
                return err
             }
        case "supportingImmunization":
            if err := json.Unmarshal([]byte(v), &strct.SupportingImmunization); err != nil {
                return err
             }
        case "supportingPatientInformation":
            if err := json.Unmarshal([]byte(v), &strct.SupportingPatientInformation); err != nil {
                return err
             }
        case "targetDisease":
            if err := json.Unmarshal([]byte(v), &strct.TargetDisease); err != nil {
                return err
             }
        case "vaccineCode":
            if err := json.Unmarshal([]byte(v), &strct.VaccineCode); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if forecastStatus (a required property) was received
    if !forecastStatusReceived {
        return errors.New("\"forecastStatus\" is required but was not present")
    }
    return nil
}
