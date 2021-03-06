// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "errors"
    "encoding/json"
    "fmt"
)

// Condition A clinical condition, problem, diagnosis, or other event, situation, issue, or clinical concept that has risen to a level of concern.
type Condition struct {

  // The date or estimated date that the condition resolved or went into remission. This is called "abatement" because of the many overloaded connotations associated with "remission" or "resolution" - Conditions are never really resolved, but they can abate.
  AbatementAge *Age `json:"abatementAge,omitempty"`

  // The date or estimated date that the condition resolved or went into remission. This is called "abatement" because of the many overloaded connotations associated with "remission" or "resolution" - Conditions are never really resolved, but they can abate.
  AbatementDateTime string `json:"abatementDateTime,omitempty"`

  // The date or estimated date that the condition resolved or went into remission. This is called "abatement" because of the many overloaded connotations associated with "remission" or "resolution" - Conditions are never really resolved, but they can abate.
  AbatementPeriod *Period `json:"abatementPeriod,omitempty"`

  // The date or estimated date that the condition resolved or went into remission. This is called "abatement" because of the many overloaded connotations associated with "remission" or "resolution" - Conditions are never really resolved, but they can abate.
  AbatementRange *Range `json:"abatementRange,omitempty"`

  // The date or estimated date that the condition resolved or went into remission. This is called "abatement" because of the many overloaded connotations associated with "remission" or "resolution" - Conditions are never really resolved, but they can abate.
  AbatementString string `json:"abatementString,omitempty"`

  // Individual who is making the condition statement.
  Asserter *Reference `json:"asserter,omitempty"`

  // The anatomical location where this condition manifests itself.
  BodySite []*CodeableConcept `json:"bodySite,omitempty"`

  // A category assigned to the condition.
  Category []*CodeableConcept `json:"category,omitempty"`

  // The clinical status of the condition.
  ClinicalStatus *CodeableConcept `json:"clinicalStatus,omitempty"`

  // Identification of the condition, problem or diagnosis.
  Code *CodeableConcept `json:"code,omitempty"`

  // These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
  Contained []interface{} `json:"contained,omitempty"`

  // The Encounter during which this Condition was created or to which the creation of this record is tightly associated.
  Encounter *Reference `json:"encounter,omitempty"`

  // Supporting evidence / manifestations that are the basis of the Condition's verification status, such as evidence that confirmed or refuted the condition.
  Evidence []*ConditionEvidence `json:"evidence,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
  Id string `json:"id,omitempty"`

  // Business identifiers assigned to this condition by the performer or other systems which remain constant as the resource is updated and propagates from server to server.
  Identifier []*Identifier `json:"identifier,omitempty"`

  // A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
  ImplicitRules string `json:"implicitRules,omitempty"`

  // The base language in which the resource is written.
  Language string `json:"language,omitempty"`

  // The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
  Meta *Meta `json:"meta,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // Additional information about the Condition. This is a general notes/comments entry  for description of the Condition, its diagnosis and prognosis.
  Note []*Annotation `json:"note,omitempty"`

  // Estimated or actual date or date-time  the condition began, in the opinion of the clinician.
  OnsetAge *Age `json:"onsetAge,omitempty"`

  // Estimated or actual date or date-time  the condition began, in the opinion of the clinician.
  OnsetDateTime string `json:"onsetDateTime,omitempty"`

  // Estimated or actual date or date-time  the condition began, in the opinion of the clinician.
  OnsetPeriod *Period `json:"onsetPeriod,omitempty"`

  // Estimated or actual date or date-time  the condition began, in the opinion of the clinician.
  OnsetRange *Range `json:"onsetRange,omitempty"`

  // Estimated or actual date or date-time  the condition began, in the opinion of the clinician.
  OnsetString string `json:"onsetString,omitempty"`

  // The recordedDate represents when this particular Condition record was created in the system, which is often a system-generated date.
  RecordedDate string `json:"recordedDate,omitempty"`

  // Individual who recorded the record and takes responsibility for its content.
  Recorder *Reference `json:"recorder,omitempty"`

  // This is a Condition resource
  ResourceType interface{} `json:"resourceType"`

  // A subjective assessment of the severity of the condition as evaluated by the clinician.
  Severity *CodeableConcept `json:"severity,omitempty"`

  // Clinical stage or grade of a condition. May include formal severity assessments.
  Stage []*ConditionStage `json:"stage,omitempty"`

  // Indicates the patient or group who the condition record is associated with.
  Subject *Reference `json:"subject"`

  // A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
  Text *Narrative `json:"text,omitempty"`

  // The verification status to support the clinical status of the condition.
  VerificationStatus *CodeableConcept `json:"verificationStatus,omitempty"`
}

func (strct *Condition) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "abatementAge" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"abatementAge\": ")
	if tmp, err := json.Marshal(strct.AbatementAge); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "abatementDateTime" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"abatementDateTime\": ")
	if tmp, err := json.Marshal(strct.AbatementDateTime); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "abatementPeriod" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"abatementPeriod\": ")
	if tmp, err := json.Marshal(strct.AbatementPeriod); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "abatementRange" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"abatementRange\": ")
	if tmp, err := json.Marshal(strct.AbatementRange); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "abatementString" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"abatementString\": ")
	if tmp, err := json.Marshal(strct.AbatementString); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "asserter" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"asserter\": ")
	if tmp, err := json.Marshal(strct.Asserter); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "bodySite" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"bodySite\": ")
	if tmp, err := json.Marshal(strct.BodySite); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "category" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"category\": ")
	if tmp, err := json.Marshal(strct.Category); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "clinicalStatus" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"clinicalStatus\": ")
	if tmp, err := json.Marshal(strct.ClinicalStatus); err != nil {
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
    // Marshal the "evidence" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"evidence\": ")
	if tmp, err := json.Marshal(strct.Evidence); err != nil {
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
    // Marshal the "onsetAge" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"onsetAge\": ")
	if tmp, err := json.Marshal(strct.OnsetAge); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "onsetDateTime" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"onsetDateTime\": ")
	if tmp, err := json.Marshal(strct.OnsetDateTime); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "onsetPeriod" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"onsetPeriod\": ")
	if tmp, err := json.Marshal(strct.OnsetPeriod); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "onsetRange" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"onsetRange\": ")
	if tmp, err := json.Marshal(strct.OnsetRange); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "onsetString" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"onsetString\": ")
	if tmp, err := json.Marshal(strct.OnsetString); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "recordedDate" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"recordedDate\": ")
	if tmp, err := json.Marshal(strct.RecordedDate); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "recorder" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"recorder\": ")
	if tmp, err := json.Marshal(strct.Recorder); err != nil {
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
    // Marshal the "severity" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"severity\": ")
	if tmp, err := json.Marshal(strct.Severity); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "stage" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"stage\": ")
	if tmp, err := json.Marshal(strct.Stage); err != nil {
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
    // Marshal the "verificationStatus" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"verificationStatus\": ")
	if tmp, err := json.Marshal(strct.VerificationStatus); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Condition) UnmarshalJSON(b []byte) error {
    resourceTypeReceived := false
    subjectReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "abatementAge":
            if err := json.Unmarshal([]byte(v), &strct.AbatementAge); err != nil {
                return err
             }
        case "abatementDateTime":
            if err := json.Unmarshal([]byte(v), &strct.AbatementDateTime); err != nil {
                return err
             }
        case "abatementPeriod":
            if err := json.Unmarshal([]byte(v), &strct.AbatementPeriod); err != nil {
                return err
             }
        case "abatementRange":
            if err := json.Unmarshal([]byte(v), &strct.AbatementRange); err != nil {
                return err
             }
        case "abatementString":
            if err := json.Unmarshal([]byte(v), &strct.AbatementString); err != nil {
                return err
             }
        case "asserter":
            if err := json.Unmarshal([]byte(v), &strct.Asserter); err != nil {
                return err
             }
        case "bodySite":
            if err := json.Unmarshal([]byte(v), &strct.BodySite); err != nil {
                return err
             }
        case "category":
            if err := json.Unmarshal([]byte(v), &strct.Category); err != nil {
                return err
             }
        case "clinicalStatus":
            if err := json.Unmarshal([]byte(v), &strct.ClinicalStatus); err != nil {
                return err
             }
        case "code":
            if err := json.Unmarshal([]byte(v), &strct.Code); err != nil {
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
        case "evidence":
            if err := json.Unmarshal([]byte(v), &strct.Evidence); err != nil {
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
        case "modifierExtension":
            if err := json.Unmarshal([]byte(v), &strct.ModifierExtension); err != nil {
                return err
             }
        case "note":
            if err := json.Unmarshal([]byte(v), &strct.Note); err != nil {
                return err
             }
        case "onsetAge":
            if err := json.Unmarshal([]byte(v), &strct.OnsetAge); err != nil {
                return err
             }
        case "onsetDateTime":
            if err := json.Unmarshal([]byte(v), &strct.OnsetDateTime); err != nil {
                return err
             }
        case "onsetPeriod":
            if err := json.Unmarshal([]byte(v), &strct.OnsetPeriod); err != nil {
                return err
             }
        case "onsetRange":
            if err := json.Unmarshal([]byte(v), &strct.OnsetRange); err != nil {
                return err
             }
        case "onsetString":
            if err := json.Unmarshal([]byte(v), &strct.OnsetString); err != nil {
                return err
             }
        case "recordedDate":
            if err := json.Unmarshal([]byte(v), &strct.RecordedDate); err != nil {
                return err
             }
        case "recorder":
            if err := json.Unmarshal([]byte(v), &strct.Recorder); err != nil {
                return err
             }
        case "resourceType":
            if err := json.Unmarshal([]byte(v), &strct.ResourceType); err != nil {
                return err
             }
            resourceTypeReceived = true
        case "severity":
            if err := json.Unmarshal([]byte(v), &strct.Severity); err != nil {
                return err
             }
        case "stage":
            if err := json.Unmarshal([]byte(v), &strct.Stage); err != nil {
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
        case "verificationStatus":
            if err := json.Unmarshal([]byte(v), &strct.VerificationStatus); err != nil {
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
