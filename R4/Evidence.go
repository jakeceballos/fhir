// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "errors"
    "encoding/json"
    "fmt"
)

// Evidence The Evidence resource describes the conditional state (population and any exposures being compared within the population) and outcome (if specified) that the knowledge (evidence, assertion, recommendation) is about.
type Evidence struct {

  // The date on which the resource content was approved by the publisher. Approval happens once when the content is officially approved for usage.
  ApprovalDate string `json:"approvalDate,omitempty"`

  // An individiual or organization primarily involved in the creation and maintenance of the content.
  Author []*ContactDetail `json:"author,omitempty"`

  // Contact details to assist a user in finding and communicating with the publisher.
  Contact []*ContactDetail `json:"contact,omitempty"`

  // These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
  Contained []interface{} `json:"contained,omitempty"`

  // A copyright statement relating to the evidence and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the evidence.
  Copyright string `json:"copyright,omitempty"`

  // The date  (and optionally time) when the evidence was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the evidence changes.
  Date string `json:"date,omitempty"`

  // A free text natural language description of the evidence from a consumer's perspective.
  Description string `json:"description,omitempty"`

  // An individual or organization primarily responsible for internal coherence of the content.
  Editor []*ContactDetail `json:"editor,omitempty"`

  // The period during which the evidence content was or is planned to be in active use.
  EffectivePeriod *Period `json:"effectivePeriod,omitempty"`

  // An individual or organization responsible for officially endorsing the content for use in some setting.
  Endorser []*ContactDetail `json:"endorser,omitempty"`

  // A reference to a EvidenceVariable resource that defines the population for the research.
  ExposureBackground *Reference `json:"exposureBackground"`

  // A reference to a EvidenceVariable resource that defines the exposure for the research.
  ExposureVariant []*Reference `json:"exposureVariant,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
  Id string `json:"id,omitempty"`

  // A formal identifier that is used to identify this evidence when it is represented in other formats, or referenced in a specification, model, design or an instance.
  Identifier []*Identifier `json:"identifier,omitempty"`

  // A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
  ImplicitRules string `json:"implicitRules,omitempty"`

  // A legal or geographic region in which the evidence is intended to be used.
  Jurisdiction []*CodeableConcept `json:"jurisdiction,omitempty"`

  // The base language in which the resource is written.
  Language string `json:"language,omitempty"`

  // The date on which the resource content was last reviewed. Review happens periodically after approval but does not change the original approval date.
  LastReviewDate string `json:"lastReviewDate,omitempty"`

  // The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
  Meta *Meta `json:"meta,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // A natural language name identifying the evidence. This name should be usable as an identifier for the module by machine processing applications such as code generation.
  Name string `json:"name,omitempty"`

  // A human-readable string to clarify or explain concepts about the resource.
  Note []*Annotation `json:"note,omitempty"`

  // A reference to a EvidenceVariable resomece that defines the outcome for the research.
  Outcome []*Reference `json:"outcome,omitempty"`

  // The name of the organization or individual that published the evidence.
  Publisher string `json:"publisher,omitempty"`

  // Related artifacts such as additional documentation, justification, or bibliographic references.
  RelatedArtifact []*RelatedArtifact `json:"relatedArtifact,omitempty"`

  // This is a Evidence resource
  ResourceType interface{} `json:"resourceType"`

  // An individual or organization primarily responsible for review of some aspect of the content.
  Reviewer []*ContactDetail `json:"reviewer,omitempty"`

  // The short title provides an alternate title for use in informal descriptive contexts where the full, formal title is not necessary.
  ShortTitle string `json:"shortTitle,omitempty"`

  // The status of this evidence. Enables tracking the life-cycle of the content.
  Status interface{} `json:"status,omitempty"`

  // An explanatory or alternate title for the Evidence giving additional information about its content.
  Subtitle string `json:"subtitle,omitempty"`

  // A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
  Text *Narrative `json:"text,omitempty"`

  // A short, descriptive, user-friendly title for the evidence.
  Title string `json:"title,omitempty"`

  // Descriptive topics related to the content of the Evidence. Topics provide a high-level categorization grouping types of Evidences that can be useful for filtering and searching.
  Topic []*CodeableConcept `json:"topic,omitempty"`

  // An absolute URI that is used to identify this evidence when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this evidence is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the evidence is stored on different servers.
  Url string `json:"url,omitempty"`

  // The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate evidence instances.
  UseContext []*UsageContext `json:"useContext,omitempty"`

  // The identifier that is used to identify this version of the evidence when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the evidence author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence. To provide a version consistent with the Decision Support Service specification, use the format Major.Minor.Revision (e.g. 1.0.0). For more information on versioning knowledge assets, refer to the Decision Support Service specification. Note that a version is required for non-experimental active artifacts.
  Version string `json:"version,omitempty"`
}

func (strct *Evidence) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "approvalDate" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"approvalDate\": ")
	if tmp, err := json.Marshal(strct.ApprovalDate); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "author" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"author\": ")
	if tmp, err := json.Marshal(strct.Author); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "contact" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"contact\": ")
	if tmp, err := json.Marshal(strct.Contact); err != nil {
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
    // Marshal the "copyright" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"copyright\": ")
	if tmp, err := json.Marshal(strct.Copyright); err != nil {
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
    // Marshal the "editor" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"editor\": ")
	if tmp, err := json.Marshal(strct.Editor); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "effectivePeriod" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"effectivePeriod\": ")
	if tmp, err := json.Marshal(strct.EffectivePeriod); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "endorser" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"endorser\": ")
	if tmp, err := json.Marshal(strct.Endorser); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "ExposureBackground" field is required
    if strct.ExposureBackground == nil {
        return nil, errors.New("exposureBackground is a required field")
    }
    // Marshal the "exposureBackground" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"exposureBackground\": ")
	if tmp, err := json.Marshal(strct.ExposureBackground); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "exposureVariant" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"exposureVariant\": ")
	if tmp, err := json.Marshal(strct.ExposureVariant); err != nil {
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
    // Marshal the "jurisdiction" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"jurisdiction\": ")
	if tmp, err := json.Marshal(strct.Jurisdiction); err != nil {
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
    // Marshal the "lastReviewDate" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"lastReviewDate\": ")
	if tmp, err := json.Marshal(strct.LastReviewDate); err != nil {
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
    // Marshal the "publisher" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"publisher\": ")
	if tmp, err := json.Marshal(strct.Publisher); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "relatedArtifact" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"relatedArtifact\": ")
	if tmp, err := json.Marshal(strct.RelatedArtifact); err != nil {
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
    // Marshal the "reviewer" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"reviewer\": ")
	if tmp, err := json.Marshal(strct.Reviewer); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "shortTitle" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"shortTitle\": ")
	if tmp, err := json.Marshal(strct.ShortTitle); err != nil {
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
    // Marshal the "subtitle" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"subtitle\": ")
	if tmp, err := json.Marshal(strct.Subtitle); err != nil {
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
    // Marshal the "title" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"title\": ")
	if tmp, err := json.Marshal(strct.Title); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "topic" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"topic\": ")
	if tmp, err := json.Marshal(strct.Topic); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "url" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"url\": ")
	if tmp, err := json.Marshal(strct.Url); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "useContext" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"useContext\": ")
	if tmp, err := json.Marshal(strct.UseContext); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "version" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"version\": ")
	if tmp, err := json.Marshal(strct.Version); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Evidence) UnmarshalJSON(b []byte) error {
    exposureBackgroundReceived := false
    resourceTypeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "approvalDate":
            if err := json.Unmarshal([]byte(v), &strct.ApprovalDate); err != nil {
                return err
             }
        case "author":
            if err := json.Unmarshal([]byte(v), &strct.Author); err != nil {
                return err
             }
        case "contact":
            if err := json.Unmarshal([]byte(v), &strct.Contact); err != nil {
                return err
             }
        case "contained":
            if err := json.Unmarshal([]byte(v), &strct.Contained); err != nil {
                return err
             }
        case "copyright":
            if err := json.Unmarshal([]byte(v), &strct.Copyright); err != nil {
                return err
             }
        case "date":
            if err := json.Unmarshal([]byte(v), &strct.Date); err != nil {
                return err
             }
        case "description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
                return err
             }
        case "editor":
            if err := json.Unmarshal([]byte(v), &strct.Editor); err != nil {
                return err
             }
        case "effectivePeriod":
            if err := json.Unmarshal([]byte(v), &strct.EffectivePeriod); err != nil {
                return err
             }
        case "endorser":
            if err := json.Unmarshal([]byte(v), &strct.Endorser); err != nil {
                return err
             }
        case "exposureBackground":
            if err := json.Unmarshal([]byte(v), &strct.ExposureBackground); err != nil {
                return err
             }
            exposureBackgroundReceived = true
        case "exposureVariant":
            if err := json.Unmarshal([]byte(v), &strct.ExposureVariant); err != nil {
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
        case "jurisdiction":
            if err := json.Unmarshal([]byte(v), &strct.Jurisdiction); err != nil {
                return err
             }
        case "language":
            if err := json.Unmarshal([]byte(v), &strct.Language); err != nil {
                return err
             }
        case "lastReviewDate":
            if err := json.Unmarshal([]byte(v), &strct.LastReviewDate); err != nil {
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
        case "outcome":
            if err := json.Unmarshal([]byte(v), &strct.Outcome); err != nil {
                return err
             }
        case "publisher":
            if err := json.Unmarshal([]byte(v), &strct.Publisher); err != nil {
                return err
             }
        case "relatedArtifact":
            if err := json.Unmarshal([]byte(v), &strct.RelatedArtifact); err != nil {
                return err
             }
        case "resourceType":
            if err := json.Unmarshal([]byte(v), &strct.ResourceType); err != nil {
                return err
             }
            resourceTypeReceived = true
        case "reviewer":
            if err := json.Unmarshal([]byte(v), &strct.Reviewer); err != nil {
                return err
             }
        case "shortTitle":
            if err := json.Unmarshal([]byte(v), &strct.ShortTitle); err != nil {
                return err
             }
        case "status":
            if err := json.Unmarshal([]byte(v), &strct.Status); err != nil {
                return err
             }
        case "subtitle":
            if err := json.Unmarshal([]byte(v), &strct.Subtitle); err != nil {
                return err
             }
        case "text":
            if err := json.Unmarshal([]byte(v), &strct.Text); err != nil {
                return err
             }
        case "title":
            if err := json.Unmarshal([]byte(v), &strct.Title); err != nil {
                return err
             }
        case "topic":
            if err := json.Unmarshal([]byte(v), &strct.Topic); err != nil {
                return err
             }
        case "url":
            if err := json.Unmarshal([]byte(v), &strct.Url); err != nil {
                return err
             }
        case "useContext":
            if err := json.Unmarshal([]byte(v), &strct.UseContext); err != nil {
                return err
             }
        case "version":
            if err := json.Unmarshal([]byte(v), &strct.Version); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if exposureBackground (a required property) was received
    if !exposureBackgroundReceived {
        return errors.New("\"exposureBackground\" is required but was not present")
    }
    // check if resourceType (a required property) was received
    if !resourceTypeReceived {
        return errors.New("\"resourceType\" is required but was not present")
    }
    return nil
}
