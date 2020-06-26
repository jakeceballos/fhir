// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "encoding/json"
    "fmt"
    "errors"
    "bytes"
)

// OrganizationAffiliation Defines an affiliation/assotiation/relationship between 2 distinct oganizations, that is not a part-of relationship/sub-division relationship.
type OrganizationAffiliation struct {

  // Whether this organization affiliation record is in active use.
  Active bool `json:"active,omitempty"`

  // Definition of the role the participatingOrganization plays in the association.
  Code []*CodeableConcept `json:"code,omitempty"`

  // These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
  Contained []interface{} `json:"contained,omitempty"`

  // Technical endpoints providing access to services operated for this role.
  Endpoint []*Reference `json:"endpoint,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Healthcare services provided through the role.
  HealthcareService []*Reference `json:"healthcareService,omitempty"`

  // The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
  Id string `json:"id,omitempty"`

  // Business identifiers that are specific to this role.
  Identifier []*Identifier `json:"identifier,omitempty"`

  // A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
  ImplicitRules string `json:"implicitRules,omitempty"`

  // The base language in which the resource is written.
  Language string `json:"language,omitempty"`

  // The location(s) at which the role occurs.
  Location []*Reference `json:"location,omitempty"`

  // The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
  Meta *Meta `json:"meta,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // Health insurance provider network in which the participatingOrganization provides the role's services (if defined) at the indicated locations (if defined).
  Network []*Reference `json:"network,omitempty"`

  // Organization where the role is available (primary organization/has members).
  Organization *Reference `json:"organization,omitempty"`

  // The Participating Organization provides/performs the role(s) defined by the code to the Primary Organization (e.g. providing services or is a member of).
  ParticipatingOrganization *Reference `json:"participatingOrganization,omitempty"`

  // The period during which the participatingOrganization is affiliated with the primary organization.
  Period *Period `json:"period,omitempty"`

  // This is a OrganizationAffiliation resource
  ResourceType interface{} `json:"resourceType"`

  // Specific specialty of the participatingOrganization in the context of the role.
  Specialty []*CodeableConcept `json:"specialty,omitempty"`

  // Contact details at the participatingOrganization relevant to this Affiliation.
  Telecom []*ContactPoint `json:"telecom,omitempty"`

  // A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
  Text *Narrative `json:"text,omitempty"`
}

func (strct *OrganizationAffiliation) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "active" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"active\": ")
	if tmp, err := json.Marshal(strct.Active); err != nil {
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
    // Marshal the "endpoint" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"endpoint\": ")
	if tmp, err := json.Marshal(strct.Endpoint); err != nil {
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
    // Marshal the "healthcareService" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"healthcareService\": ")
	if tmp, err := json.Marshal(strct.HealthcareService); err != nil {
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
    // Marshal the "network" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"network\": ")
	if tmp, err := json.Marshal(strct.Network); err != nil {
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
    // Marshal the "participatingOrganization" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"participatingOrganization\": ")
	if tmp, err := json.Marshal(strct.ParticipatingOrganization); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "period" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"period\": ")
	if tmp, err := json.Marshal(strct.Period); err != nil {
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
    // Marshal the "specialty" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"specialty\": ")
	if tmp, err := json.Marshal(strct.Specialty); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "telecom" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"telecom\": ")
	if tmp, err := json.Marshal(strct.Telecom); err != nil {
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

func (strct *OrganizationAffiliation) UnmarshalJSON(b []byte) error {
    resourceTypeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "active":
            if err := json.Unmarshal([]byte(v), &strct.Active); err != nil {
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
        case "endpoint":
            if err := json.Unmarshal([]byte(v), &strct.Endpoint); err != nil {
                return err
             }
        case "extension":
            if err := json.Unmarshal([]byte(v), &strct.Extension); err != nil {
                return err
             }
        case "healthcareService":
            if err := json.Unmarshal([]byte(v), &strct.HealthcareService); err != nil {
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
        case "network":
            if err := json.Unmarshal([]byte(v), &strct.Network); err != nil {
                return err
             }
        case "organization":
            if err := json.Unmarshal([]byte(v), &strct.Organization); err != nil {
                return err
             }
        case "participatingOrganization":
            if err := json.Unmarshal([]byte(v), &strct.ParticipatingOrganization); err != nil {
                return err
             }
        case "period":
            if err := json.Unmarshal([]byte(v), &strct.Period); err != nil {
                return err
             }
        case "resourceType":
            if err := json.Unmarshal([]byte(v), &strct.ResourceType); err != nil {
                return err
             }
            resourceTypeReceived = true
        case "specialty":
            if err := json.Unmarshal([]byte(v), &strct.Specialty); err != nil {
                return err
             }
        case "telecom":
            if err := json.Unmarshal([]byte(v), &strct.Telecom); err != nil {
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
    return nil
}
