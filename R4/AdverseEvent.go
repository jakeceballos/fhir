// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "errors"
    "encoding/json"
    "fmt"
)

// AdverseEvent Actual or  potential/avoided event causing unintended physical injury resulting from or contributed to by medical care, a research study or other healthcare setting factors that requires additional monitoring, treatment, or hospitalization, or that results in death.
type AdverseEvent struct {

  // Whether the event actually happened, or just had the potential to. Note that this is independent of whether anyone was affected or harmed or how severely.
  Actuality interface{} `json:"actuality,omitempty"`

  // The overall type of event, intended for search and filtering purposes.
  Category []*CodeableConcept `json:"category,omitempty"`

  // These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
  Contained []interface{} `json:"contained,omitempty"`

  // Parties that may or should contribute or have contributed information to the adverse event, which can consist of one or more activities.  Such information includes information leading to the decision to perform the activity and how to perform the activity (e.g. consultant), information that the activity itself seeks to reveal (e.g. informant of clinical history), or information about what activity was performed (e.g. informant witness).
  Contributor []*Reference `json:"contributor,omitempty"`

  // The date (and perhaps time) when the adverse event occurred.
  Date string `json:"date,omitempty"`

  // Estimated or actual date the AdverseEvent began, in the opinion of the reporter.
  Detected string `json:"detected,omitempty"`

  // The Encounter during which AdverseEvent was created or to which the creation of this record is tightly associated.
  Encounter *Reference `json:"encounter,omitempty"`

  // This element defines the specific type of event that occurred or that was prevented from occurring.
  Event *CodeableConcept `json:"event,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
  Id string `json:"id,omitempty"`

  // Business identifiers assigned to this adverse event by the performer or other systems which remain constant as the resource is updated and propagates from server to server.
  Identifier *Identifier `json:"identifier,omitempty"`

  // A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
  ImplicitRules string `json:"implicitRules,omitempty"`

  // The base language in which the resource is written.
  Language string `json:"language,omitempty"`

  // The information about where the adverse event occurred.
  Location *Reference `json:"location,omitempty"`

  // The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
  Meta *Meta `json:"meta,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // Describes the type of outcome from the adverse event.
  Outcome *CodeableConcept `json:"outcome,omitempty"`

  // The date on which the existence of the AdverseEvent was first recorded.
  RecordedDate string `json:"recordedDate,omitempty"`

  // Information on who recorded the adverse event.  May be the patient or a practitioner.
  Recorder *Reference `json:"recorder,omitempty"`

  // AdverseEvent.referenceDocument.
  ReferenceDocument []*Reference `json:"referenceDocument,omitempty"`

  // This is a AdverseEvent resource
  ResourceType interface{} `json:"resourceType"`

  // Includes information about the reaction that occurred as a result of exposure to a substance (for example, a drug or a chemical).
  ResultingCondition []*Reference `json:"resultingCondition,omitempty"`

  // Assessment whether this event was of real importance.
  Seriousness *CodeableConcept `json:"seriousness,omitempty"`

  // Describes the severity of the adverse event, in relation to the subject. Contrast to AdverseEvent.seriousness - a severe rash might not be serious, but a mild heart problem is.
  Severity *CodeableConcept `json:"severity,omitempty"`

  // AdverseEvent.study.
  Study []*Reference `json:"study,omitempty"`

  // This subject or group impacted by the event.
  Subject *Reference `json:"subject"`

  // AdverseEvent.subjectMedicalHistory.
  SubjectMedicalHistory []*Reference `json:"subjectMedicalHistory,omitempty"`

  // Describes the entity that is suspected to have caused the adverse event.
  SuspectEntity []*AdverseEventSuspectEntity `json:"suspectEntity,omitempty"`

  // A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
  Text *Narrative `json:"text,omitempty"`
}

func (strct *AdverseEvent) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "actuality" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"actuality\": ")
	if tmp, err := json.Marshal(strct.Actuality); err != nil {
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
    // Marshal the "contributor" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"contributor\": ")
	if tmp, err := json.Marshal(strct.Contributor); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "date" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"date\": ")
	if tmp, err := json.Marshal(strct.Date); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "detected" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"detected\": ")
	if tmp, err := json.Marshal(strct.Detected); err != nil {
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
    // Marshal the "event" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"event\": ")
	if tmp, err := json.Marshal(strct.Event); err != nil {
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
    // Marshal the "location" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"location\": ")
	if tmp, err := json.Marshal(strct.Location); err != nil {
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
    // Marshal the "outcome" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"outcome\": ")
	if tmp, err := json.Marshal(strct.Outcome); err != nil {
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
    // Marshal the "referenceDocument" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"referenceDocument\": ")
	if tmp, err := json.Marshal(strct.ReferenceDocument); err != nil {
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
    // Marshal the "resultingCondition" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"resultingCondition\": ")
	if tmp, err := json.Marshal(strct.ResultingCondition); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "seriousness" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"seriousness\": ")
	if tmp, err := json.Marshal(strct.Seriousness); err != nil {
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
    // Marshal the "study" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"study\": ")
	if tmp, err := json.Marshal(strct.Study); err != nil {
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
    // Marshal the "subjectMedicalHistory" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"subjectMedicalHistory\": ")
	if tmp, err := json.Marshal(strct.SubjectMedicalHistory); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "suspectEntity" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"suspectEntity\": ")
	if tmp, err := json.Marshal(strct.SuspectEntity); err != nil {
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

func (strct *AdverseEvent) UnmarshalJSON(b []byte) error {
    resourceTypeReceived := false
    subjectReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "actuality":
            if err := json.Unmarshal([]byte(v), &strct.Actuality); err != nil {
                return err
             }
        case "category":
            if err := json.Unmarshal([]byte(v), &strct.Category); err != nil {
                return err
             }
        case "contained":
            if err := json.Unmarshal([]byte(v), &strct.Contained); err != nil {
                return err
             }
        case "contributor":
            if err := json.Unmarshal([]byte(v), &strct.Contributor); err != nil {
                return err
             }
        case "date":
            if err := json.Unmarshal([]byte(v), &strct.Date); err != nil {
                return err
             }
        case "detected":
            if err := json.Unmarshal([]byte(v), &strct.Detected); err != nil {
                return err
             }
        case "encounter":
            if err := json.Unmarshal([]byte(v), &strct.Encounter); err != nil {
                return err
             }
        case "event":
            if err := json.Unmarshal([]byte(v), &strct.Event); err != nil {
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
        case "location":
            if err := json.Unmarshal([]byte(v), &strct.Location); err != nil {
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
        case "outcome":
            if err := json.Unmarshal([]byte(v), &strct.Outcome); err != nil {
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
        case "referenceDocument":
            if err := json.Unmarshal([]byte(v), &strct.ReferenceDocument); err != nil {
                return err
             }
        case "resourceType":
            if err := json.Unmarshal([]byte(v), &strct.ResourceType); err != nil {
                return err
             }
            resourceTypeReceived = true
        case "resultingCondition":
            if err := json.Unmarshal([]byte(v), &strct.ResultingCondition); err != nil {
                return err
             }
        case "seriousness":
            if err := json.Unmarshal([]byte(v), &strct.Seriousness); err != nil {
                return err
             }
        case "severity":
            if err := json.Unmarshal([]byte(v), &strct.Severity); err != nil {
                return err
             }
        case "study":
            if err := json.Unmarshal([]byte(v), &strct.Study); err != nil {
                return err
             }
        case "subject":
            if err := json.Unmarshal([]byte(v), &strct.Subject); err != nil {
                return err
             }
            subjectReceived = true
        case "subjectMedicalHistory":
            if err := json.Unmarshal([]byte(v), &strct.SubjectMedicalHistory); err != nil {
                return err
             }
        case "suspectEntity":
            if err := json.Unmarshal([]byte(v), &strct.SuspectEntity); err != nil {
                return err
             }
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