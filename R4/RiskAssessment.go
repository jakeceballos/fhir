// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "fmt"
    "bytes"
    "errors"
    "encoding/json"
)

// RiskAssessment An assessment of the likely outcome(s) for a patient or other subject as well as the likelihood of each outcome.
type RiskAssessment struct {

  // A reference to the request that is fulfilled by this risk assessment.
  BasedOn *Reference `json:"basedOn,omitempty"`

  // Indicates the source data considered as part of the assessment (for example, FamilyHistory, Observations, Procedures, Conditions, etc.).
  Basis []*Reference `json:"basis,omitempty"`

  // The type of the risk assessment performed.
  Code *CodeableConcept `json:"code,omitempty"`

  // For assessments or prognosis specific to a particular condition, indicates the condition being assessed.
  Condition *Reference `json:"condition,omitempty"`

  // These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
  Contained []interface{} `json:"contained,omitempty"`

  // The encounter where the assessment was performed.
  Encounter *Reference `json:"encounter,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
  Id string `json:"id,omitempty"`

  // Business identifier assigned to the risk assessment.
  Identifier []*Identifier `json:"identifier,omitempty"`

  // A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
  ImplicitRules string `json:"implicitRules,omitempty"`

  // The base language in which the resource is written.
  Language string `json:"language,omitempty"`

  // The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
  Meta *Meta `json:"meta,omitempty"`

  // The algorithm, process or mechanism used to evaluate the risk.
  Method *CodeableConcept `json:"method,omitempty"`

  // A description of the steps that might be taken to reduce the identified risk(s).
  Mitigation string `json:"mitigation,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // Additional comments about the risk assessment.
  Note []*Annotation `json:"note,omitempty"`

  // The date (and possibly time) the risk assessment was performed.
  OccurrenceDateTime string `json:"occurrenceDateTime,omitempty"`

  // The date (and possibly time) the risk assessment was performed.
  OccurrencePeriod *Period `json:"occurrencePeriod,omitempty"`

  // A reference to a resource that this risk assessment is part of, such as a Procedure.
  Parent *Reference `json:"parent,omitempty"`

  // The provider or software application that performed the assessment.
  Performer *Reference `json:"performer,omitempty"`

  // Describes the expected outcome for the subject.
  Prediction []*RiskAssessmentPrediction `json:"prediction,omitempty"`

  // The reason the risk assessment was performed.
  ReasonCode []*CodeableConcept `json:"reasonCode,omitempty"`

  // Resources supporting the reason the risk assessment was performed.
  ReasonReference []*Reference `json:"reasonReference,omitempty"`

  // This is a RiskAssessment resource
  ResourceType interface{} `json:"resourceType"`

  // The status of the RiskAssessment, using the same statuses as an Observation.
  Status string `json:"status,omitempty"`

  // The patient or group the risk assessment applies to.
  Subject *Reference `json:"subject"`

  // A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
  Text *Narrative `json:"text,omitempty"`
}

func (strct *RiskAssessment) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "basedOn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"basedOn\": ")
	if tmp, err := json.Marshal(strct.BasedOn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "basis" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"basis\": ")
	if tmp, err := json.Marshal(strct.Basis); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
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
    // Marshal the "condition" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"condition\": ")
	if tmp, err := json.Marshal(strct.Condition); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "contained" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"contained\": ")
	if tmp, err := json.Marshal(strct.Contained); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "encounter" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"encounter\": ")
	if tmp, err := json.Marshal(strct.Encounter); err != nil {
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
    // Marshal the "identifier" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"identifier\": ")
	if tmp, err := json.Marshal(strct.Identifier); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "implicitRules" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"implicitRules\": ")
	if tmp, err := json.Marshal(strct.ImplicitRules); err != nil {
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
    // Marshal the "meta" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"meta\": ")
	if tmp, err := json.Marshal(strct.Meta); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "method" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"method\": ")
	if tmp, err := json.Marshal(strct.Method); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "mitigation" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"mitigation\": ")
	if tmp, err := json.Marshal(strct.Mitigation); err != nil {
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
    // Marshal the "note" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"note\": ")
	if tmp, err := json.Marshal(strct.Note); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "occurrenceDateTime" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"occurrenceDateTime\": ")
	if tmp, err := json.Marshal(strct.OccurrenceDateTime); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "occurrencePeriod" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"occurrencePeriod\": ")
	if tmp, err := json.Marshal(strct.OccurrencePeriod); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "parent" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"parent\": ")
	if tmp, err := json.Marshal(strct.Parent); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "performer" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"performer\": ")
	if tmp, err := json.Marshal(strct.Performer); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "prediction" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"prediction\": ")
	if tmp, err := json.Marshal(strct.Prediction); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "reasonCode" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"reasonCode\": ")
	if tmp, err := json.Marshal(strct.ReasonCode); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "reasonReference" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"reasonReference\": ")
	if tmp, err := json.Marshal(strct.ReasonReference); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "ResourceType" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "resourceType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"resourceType\": ")
	if tmp, err := json.Marshal(strct.ResourceType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "status" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"status\": ")
	if tmp, err := json.Marshal(strct.Status); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Subject" field is required
    if strct.Subject == nil {
        return nil, errors.New("subject is a required field")
    }
    // Marshal the "subject" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"subject\": ")
	if tmp, err := json.Marshal(strct.Subject); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "text" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"text\": ")
	if tmp, err := json.Marshal(strct.Text); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *RiskAssessment) UnmarshalJSON(b []byte) error {
    resourceTypeReceived := false
    subjectReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "basedOn":
            if err := json.Unmarshal([]byte(v), &strct.BasedOn); err != nil {
                return err
             }
        case "basis":
            if err := json.Unmarshal([]byte(v), &strct.Basis); err != nil {
                return err
             }
        case "code":
            if err := json.Unmarshal([]byte(v), &strct.Code); err != nil {
                return err
             }
        case "condition":
            if err := json.Unmarshal([]byte(v), &strct.Condition); err != nil {
                return err
             }
        case "contained":
            if err := json.Unmarshal([]byte(v), &strct.Contained); err != nil {
                return err
             }
        case "encounter":
            if err := json.Unmarshal([]byte(v), &strct.Encounter); err != nil {
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
        case "identifier":
            if err := json.Unmarshal([]byte(v), &strct.Identifier); err != nil {
                return err
             }
        case "implicitRules":
            if err := json.Unmarshal([]byte(v), &strct.ImplicitRules); err != nil {
                return err
             }
        case "language":
            if err := json.Unmarshal([]byte(v), &strct.Language); err != nil {
                return err
             }
        case "meta":
            if err := json.Unmarshal([]byte(v), &strct.Meta); err != nil {
                return err
             }
        case "method":
            if err := json.Unmarshal([]byte(v), &strct.Method); err != nil {
                return err
             }
        case "mitigation":
            if err := json.Unmarshal([]byte(v), &strct.Mitigation); err != nil {
                return err
             }
        case "modifierExtension":
            if err := json.Unmarshal([]byte(v), &strct.ModifierExtension); err != nil {
                return err
             }
        case "note":
            if err := json.Unmarshal([]byte(v), &strct.Note); err != nil {
                return err
             }
        case "occurrenceDateTime":
            if err := json.Unmarshal([]byte(v), &strct.OccurrenceDateTime); err != nil {
                return err
             }
        case "occurrencePeriod":
            if err := json.Unmarshal([]byte(v), &strct.OccurrencePeriod); err != nil {
                return err
             }
        case "parent":
            if err := json.Unmarshal([]byte(v), &strct.Parent); err != nil {
                return err
             }
        case "performer":
            if err := json.Unmarshal([]byte(v), &strct.Performer); err != nil {
                return err
             }
        case "prediction":
            if err := json.Unmarshal([]byte(v), &strct.Prediction); err != nil {
                return err
             }
        case "reasonCode":
            if err := json.Unmarshal([]byte(v), &strct.ReasonCode); err != nil {
                return err
             }
        case "reasonReference":
            if err := json.Unmarshal([]byte(v), &strct.ReasonReference); err != nil {
                return err
             }
        case "resourceType":
            if err := json.Unmarshal([]byte(v), &strct.ResourceType); err != nil {
                return err
             }
            resourceTypeReceived = true
        case "status":
            if err := json.Unmarshal([]byte(v), &strct.Status); err != nil {
                return err
             }
        case "subject":
            if err := json.Unmarshal([]byte(v), &strct.Subject); err != nil {
                return err
             }
            subjectReceived = true
        case "text":
            if err := json.Unmarshal([]byte(v), &strct.Text); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if resourceType (a required property) was received
    if !resourceTypeReceived {
        return errors.New("\"resourceType\" is required but was not present")
    }
    // check if subject (a required property) was received
    if !subjectReceived {
        return errors.New("\"subject\" is required but was not present")
    }
    return nil
}
