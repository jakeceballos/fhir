// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "errors"
    "encoding/json"
    "fmt"
    "bytes"
)

// Consent A record of a healthcare consumer’s  choices, which permits or denies identified recipient(s) or recipient role(s) to perform one or more actions within a given policy context, for specific purposes and periods of time.
type Consent struct {

  // A classification of the type of consents found in the statement. This element supports indexing and retrieval of consent statements.
  Category []*CodeableConcept `json:"category"`

  // These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
  Contained []interface{} `json:"contained,omitempty"`

  // When this  Consent was issued / created / indexed.
  DateTime string `json:"dateTime,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
  Id string `json:"id,omitempty"`

  // Unique identifier for this copy of the Consent Statement.
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

  // The organization that manages the consent, and the framework within which it is executed.
  Organization []*Reference `json:"organization,omitempty"`

  // The patient/healthcare consumer to whom this consent applies.
  Patient *Reference `json:"patient,omitempty"`

  // Either the Grantor, which is the entity responsible for granting the rights listed in a Consent Directive or the Grantee, which is the entity responsible for complying with the Consent Directive, including any obligations or limitations on authorizations and enforcement of prohibitions.
  Performer []*Reference `json:"performer,omitempty"`

  // The references to the policies that are included in this consent scope. Policies may be organizational, but are often defined jurisdictionally, or in law.
  Policy []*ConsentPolicy `json:"policy,omitempty"`

  // A reference to the specific base computable regulation or policy.
  PolicyRule *CodeableConcept `json:"policyRule,omitempty"`

  // An exception to the base policy of this consent. An exception can be an addition or removal of access permissions.
  Provision *ConsentProvision `json:"provision,omitempty"`

  // This is a Consent resource
  ResourceType interface{} `json:"resourceType"`

  // A selector of the type of consent being presented: ADR, Privacy, Treatment, Research.  This list is now extensible.
  Scope *CodeableConcept `json:"scope"`

  // The source on which this consent statement is based. The source might be a scanned original paper form, or a reference to a consent that links back to such a source, a reference to a document repository (e.g. XDS) that stores the original consent document.
  SourceAttachment *Attachment `json:"sourceAttachment,omitempty"`

  // The source on which this consent statement is based. The source might be a scanned original paper form, or a reference to a consent that links back to such a source, a reference to a document repository (e.g. XDS) that stores the original consent document.
  SourceReference *Reference `json:"sourceReference,omitempty"`

  // Indicates the current state of this consent.
  Status interface{} `json:"status,omitempty"`

  // A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
  Text *Narrative `json:"text,omitempty"`

  // Whether a treatment instruction (e.g. artificial respiration yes or no) was verified with the patient, his/her family or another authorized person.
  Verification []*ConsentVerification `json:"verification,omitempty"`
}

func (strct *Consent) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "Category" field is required
    // only required object types supported for marshal checking (for now)
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
    // Marshal the "dateTime" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"dateTime\": ")
	if tmp, err := json.Marshal(strct.DateTime); err != nil {
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
    // Marshal the "organization" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"organization\": ")
	if tmp, err := json.Marshal(strct.Organization); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
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
    // Marshal the "policy" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"policy\": ")
	if tmp, err := json.Marshal(strct.Policy); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "policyRule" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"policyRule\": ")
	if tmp, err := json.Marshal(strct.PolicyRule); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "provision" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"provision\": ")
	if tmp, err := json.Marshal(strct.Provision); err != nil {
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
    // "Scope" field is required
    if strct.Scope == nil {
        return nil, errors.New("scope is a required field")
    }
    // Marshal the "scope" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"scope\": ")
	if tmp, err := json.Marshal(strct.Scope); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "sourceAttachment" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"sourceAttachment\": ")
	if tmp, err := json.Marshal(strct.SourceAttachment); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "sourceReference" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"sourceReference\": ")
	if tmp, err := json.Marshal(strct.SourceReference); err != nil {
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
    // Marshal the "verification" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"verification\": ")
	if tmp, err := json.Marshal(strct.Verification); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Consent) UnmarshalJSON(b []byte) error {
    categoryReceived := false
    resourceTypeReceived := false
    scopeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "category":
            if err := json.Unmarshal([]byte(v), &strct.Category); err != nil {
                return err
             }
            categoryReceived = true
        case "contained":
            if err := json.Unmarshal([]byte(v), &strct.Contained); err != nil {
                return err
             }
        case "dateTime":
            if err := json.Unmarshal([]byte(v), &strct.DateTime); err != nil {
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
        case "organization":
            if err := json.Unmarshal([]byte(v), &strct.Organization); err != nil {
                return err
             }
        case "patient":
            if err := json.Unmarshal([]byte(v), &strct.Patient); err != nil {
                return err
             }
        case "performer":
            if err := json.Unmarshal([]byte(v), &strct.Performer); err != nil {
                return err
             }
        case "policy":
            if err := json.Unmarshal([]byte(v), &strct.Policy); err != nil {
                return err
             }
        case "policyRule":
            if err := json.Unmarshal([]byte(v), &strct.PolicyRule); err != nil {
                return err
             }
        case "provision":
            if err := json.Unmarshal([]byte(v), &strct.Provision); err != nil {
                return err
             }
        case "resourceType":
            if err := json.Unmarshal([]byte(v), &strct.ResourceType); err != nil {
                return err
             }
            resourceTypeReceived = true
        case "scope":
            if err := json.Unmarshal([]byte(v), &strct.Scope); err != nil {
                return err
             }
            scopeReceived = true
        case "sourceAttachment":
            if err := json.Unmarshal([]byte(v), &strct.SourceAttachment); err != nil {
                return err
             }
        case "sourceReference":
            if err := json.Unmarshal([]byte(v), &strct.SourceReference); err != nil {
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
        case "verification":
            if err := json.Unmarshal([]byte(v), &strct.Verification); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if category (a required property) was received
    if !categoryReceived {
        return errors.New("\"category\" is required but was not present")
    }
    // check if resourceType (a required property) was received
    if !resourceTypeReceived {
        return errors.New("\"resourceType\" is required but was not present")
    }
    // check if scope (a required property) was received
    if !scopeReceived {
        return errors.New("\"scope\" is required but was not present")
    }
    return nil
}
