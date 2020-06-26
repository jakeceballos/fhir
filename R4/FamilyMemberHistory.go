// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "errors"
    "encoding/json"
    "fmt"
)

// FamilyMemberHistory Significant health conditions for a person related to the patient relevant in the context of care for the patient.
type FamilyMemberHistory struct {

  // The age of the relative at the time the family member history is recorded.
  AgeAge *Age `json:"ageAge,omitempty"`

  // The age of the relative at the time the family member history is recorded.
  AgeRange *Range `json:"ageRange,omitempty"`

  // The age of the relative at the time the family member history is recorded.
  AgeString string `json:"ageString,omitempty"`

  // The actual or approximate date of birth of the relative.
  BornDate string `json:"bornDate,omitempty"`

  // The actual or approximate date of birth of the relative.
  BornPeriod *Period `json:"bornPeriod,omitempty"`

  // The actual or approximate date of birth of the relative.
  BornString string `json:"bornString,omitempty"`

  // The significant Conditions (or condition) that the family member had. This is a repeating section to allow a system to represent more than one condition per resource, though there is nothing stopping multiple resources - one per condition.
  Condition []*FamilyMemberHistoryCondition `json:"condition,omitempty"`

  // These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
  Contained []interface{} `json:"contained,omitempty"`

  // Describes why the family member's history is not available.
  DataAbsentReason *CodeableConcept `json:"dataAbsentReason,omitempty"`

  // The date (and possibly time) when the family member history was recorded or last updated.
  Date string `json:"date,omitempty"`

  // Deceased flag or the actual or approximate age of the relative at the time of death for the family member history record.
  DeceasedAge *Age `json:"deceasedAge,omitempty"`

  // Deceased flag or the actual or approximate age of the relative at the time of death for the family member history record.
  DeceasedBoolean bool `json:"deceasedBoolean,omitempty"`

  // Deceased flag or the actual or approximate age of the relative at the time of death for the family member history record.
  DeceasedDate string `json:"deceasedDate,omitempty"`

  // Deceased flag or the actual or approximate age of the relative at the time of death for the family member history record.
  DeceasedRange *Range `json:"deceasedRange,omitempty"`

  // Deceased flag or the actual or approximate age of the relative at the time of death for the family member history record.
  DeceasedString string `json:"deceasedString,omitempty"`

  // If true, indicates that the age value specified is an estimated value.
  EstimatedAge bool `json:"estimatedAge,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
  Id string `json:"id,omitempty"`

  // Business identifiers assigned to this family member history by the performer or other systems which remain constant as the resource is updated and propagates from server to server.
  Identifier []*Identifier `json:"identifier,omitempty"`

  // A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
  ImplicitRules string `json:"implicitRules,omitempty"`

  // The URL pointing to a FHIR-defined protocol, guideline, orderset or other definition that is adhered to in whole or in part by this FamilyMemberHistory.
  InstantiatesCanonical []string `json:"instantiatesCanonical,omitempty"`

  // The URL pointing to an externally maintained protocol, guideline, orderset or other definition that is adhered to in whole or in part by this FamilyMemberHistory.
  InstantiatesUri []string `json:"instantiatesUri,omitempty"`

  // The base language in which the resource is written.
  Language string `json:"language,omitempty"`

  // The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
  Meta *Meta `json:"meta,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // This will either be a name or a description; e.g. "Aunt Susan", "my cousin with the red hair".
  Name string `json:"name,omitempty"`

  // This property allows a non condition-specific note to the made about the related person. Ideally, the note would be in the condition property, but this is not always possible.
  Note []*Annotation `json:"note,omitempty"`

  // The person who this history concerns.
  Patient *Reference `json:"patient"`

  // Describes why the family member history occurred in coded or textual form.
  ReasonCode []*CodeableConcept `json:"reasonCode,omitempty"`

  // Indicates a Condition, Observation, AllergyIntolerance, or QuestionnaireResponse that justifies this family member history event.
  ReasonReference []*Reference `json:"reasonReference,omitempty"`

  // The type of relationship this person has to the patient (father, mother, brother etc.).
  Relationship *CodeableConcept `json:"relationship"`

  // This is a FamilyMemberHistory resource
  ResourceType interface{} `json:"resourceType"`

  // The birth sex of the family member.
  Sex *CodeableConcept `json:"sex,omitempty"`

  // A code specifying the status of the record of the family history of a specific family member.
  Status interface{} `json:"status,omitempty"`

  // A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
  Text *Narrative `json:"text,omitempty"`
}

func (strct *FamilyMemberHistory) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "ageAge" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ageAge\": ")
	if tmp, err := json.Marshal(strct.AgeAge); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ageRange" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ageRange\": ")
	if tmp, err := json.Marshal(strct.AgeRange); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "ageString" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"ageString\": ")
	if tmp, err := json.Marshal(strct.AgeString); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "bornDate" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"bornDate\": ")
	if tmp, err := json.Marshal(strct.BornDate); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "bornPeriod" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"bornPeriod\": ")
	if tmp, err := json.Marshal(strct.BornPeriod); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "bornString" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"bornString\": ")
	if tmp, err := json.Marshal(strct.BornString); err != nil {
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
    // Marshal the "deceasedAge" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"deceasedAge\": ")
	if tmp, err := json.Marshal(strct.DeceasedAge); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "deceasedBoolean" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"deceasedBoolean\": ")
	if tmp, err := json.Marshal(strct.DeceasedBoolean); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "deceasedDate" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"deceasedDate\": ")
	if tmp, err := json.Marshal(strct.DeceasedDate); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "deceasedRange" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"deceasedRange\": ")
	if tmp, err := json.Marshal(strct.DeceasedRange); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "deceasedString" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"deceasedString\": ")
	if tmp, err := json.Marshal(strct.DeceasedString); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "estimatedAge" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"estimatedAge\": ")
	if tmp, err := json.Marshal(strct.EstimatedAge); err != nil {
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
    // Marshal the "instantiatesCanonical" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"instantiatesCanonical\": ")
	if tmp, err := json.Marshal(strct.InstantiatesCanonical); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "instantiatesUri" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"instantiatesUri\": ")
	if tmp, err := json.Marshal(strct.InstantiatesUri); err != nil {
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
    // Marshal the "name" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
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
    // "Patient" field is required
    if strct.Patient == nil {
        return nil, errors.New("patient is a required field")
    }
    // Marshal the "patient" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"patient\": ")
	if tmp, err := json.Marshal(strct.Patient); err != nil {
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
    // "Relationship" field is required
    if strct.Relationship == nil {
        return nil, errors.New("relationship is a required field")
    }
    // Marshal the "relationship" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"relationship\": ")
	if tmp, err := json.Marshal(strct.Relationship); err != nil {
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
    // Marshal the "sex" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"sex\": ")
	if tmp, err := json.Marshal(strct.Sex); err != nil {
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

func (strct *FamilyMemberHistory) UnmarshalJSON(b []byte) error {
    patientReceived := false
    relationshipReceived := false
    resourceTypeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "ageAge":
            if err := json.Unmarshal([]byte(v), &strct.AgeAge); err != nil {
                return err
             }
        case "ageRange":
            if err := json.Unmarshal([]byte(v), &strct.AgeRange); err != nil {
                return err
             }
        case "ageString":
            if err := json.Unmarshal([]byte(v), &strct.AgeString); err != nil {
                return err
             }
        case "bornDate":
            if err := json.Unmarshal([]byte(v), &strct.BornDate); err != nil {
                return err
             }
        case "bornPeriod":
            if err := json.Unmarshal([]byte(v), &strct.BornPeriod); err != nil {
                return err
             }
        case "bornString":
            if err := json.Unmarshal([]byte(v), &strct.BornString); err != nil {
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
        case "dataAbsentReason":
            if err := json.Unmarshal([]byte(v), &strct.DataAbsentReason); err != nil {
                return err
             }
        case "date":
            if err := json.Unmarshal([]byte(v), &strct.Date); err != nil {
                return err
             }
        case "deceasedAge":
            if err := json.Unmarshal([]byte(v), &strct.DeceasedAge); err != nil {
                return err
             }
        case "deceasedBoolean":
            if err := json.Unmarshal([]byte(v), &strct.DeceasedBoolean); err != nil {
                return err
             }
        case "deceasedDate":
            if err := json.Unmarshal([]byte(v), &strct.DeceasedDate); err != nil {
                return err
             }
        case "deceasedRange":
            if err := json.Unmarshal([]byte(v), &strct.DeceasedRange); err != nil {
                return err
             }
        case "deceasedString":
            if err := json.Unmarshal([]byte(v), &strct.DeceasedString); err != nil {
                return err
             }
        case "estimatedAge":
            if err := json.Unmarshal([]byte(v), &strct.EstimatedAge); err != nil {
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
        case "instantiatesCanonical":
            if err := json.Unmarshal([]byte(v), &strct.InstantiatesCanonical); err != nil {
                return err
             }
        case "instantiatesUri":
            if err := json.Unmarshal([]byte(v), &strct.InstantiatesUri); err != nil {
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
        case "name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
        case "note":
            if err := json.Unmarshal([]byte(v), &strct.Note); err != nil {
                return err
             }
        case "patient":
            if err := json.Unmarshal([]byte(v), &strct.Patient); err != nil {
                return err
             }
            patientReceived = true
        case "reasonCode":
            if err := json.Unmarshal([]byte(v), &strct.ReasonCode); err != nil {
                return err
             }
        case "reasonReference":
            if err := json.Unmarshal([]byte(v), &strct.ReasonReference); err != nil {
                return err
             }
        case "relationship":
            if err := json.Unmarshal([]byte(v), &strct.Relationship); err != nil {
                return err
             }
            relationshipReceived = true
        case "resourceType":
            if err := json.Unmarshal([]byte(v), &strct.ResourceType); err != nil {
                return err
             }
            resourceTypeReceived = true
        case "sex":
            if err := json.Unmarshal([]byte(v), &strct.Sex); err != nil {
                return err
             }
        case "status":
            if err := json.Unmarshal([]byte(v), &strct.Status); err != nil {
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
    // check if patient (a required property) was received
    if !patientReceived {
        return errors.New("\"patient\" is required but was not present")
    }
    // check if relationship (a required property) was received
    if !relationshipReceived {
        return errors.New("\"relationship\" is required but was not present")
    }
    // check if resourceType (a required property) was received
    if !resourceTypeReceived {
        return errors.New("\"resourceType\" is required but was not present")
    }
    return nil
}
