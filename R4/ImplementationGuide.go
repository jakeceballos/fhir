// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "encoding/json"
    "fmt"
    "errors"
    "bytes"
)

// ImplementationGuide A set of rules of how a particular interoperability or standards problem is solved - typically through the use of FHIR resources. This resource is used to gather all the parts of an implementation guide into a logical whole and to publish a computable definition of all the parts.
type ImplementationGuide struct {

  // Contact details to assist a user in finding and communicating with the publisher.
  Contact []*ContactDetail `json:"contact,omitempty"`

  // These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
  Contained []interface{} `json:"contained,omitempty"`

  // A copyright statement relating to the implementation guide and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the implementation guide.
  Copyright string `json:"copyright,omitempty"`

  // The date  (and optionally time) when the implementation guide was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the implementation guide changes.
  Date string `json:"date,omitempty"`

  // The information needed by an IG publisher tool to publish the whole implementation guide.
  Definition *ImplementationGuideDefinition `json:"definition,omitempty"`

  // Another implementation guide that this implementation depends on. Typically, an implementation guide uses value sets, profiles etc.defined in other implementation guides.
  DependsOn []*ImplementationGuideDependsOn `json:"dependsOn,omitempty"`

  // A free text natural language description of the implementation guide from a consumer's perspective.
  Description string `json:"description,omitempty"`

  // A Boolean value to indicate that this implementation guide is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
  Experimental bool `json:"experimental,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // The version(s) of the FHIR specification that this ImplementationGuide targets - e.g. describes how to use. The value of this element is the formal version of the specification, without the revision number, e.g. [publication].[major].[minor], which is 4.0.0. for this version.
  FhirVersion []interface{} `json:"fhirVersion,omitempty"`

  // A set of profiles that all resources covered by this implementation guide must conform to.
  Global []*ImplementationGuideGlobal `json:"global,omitempty"`

  // The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
  Id string `json:"id,omitempty"`

  // A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
  ImplicitRules string `json:"implicitRules,omitempty"`

  // A legal or geographic region in which the implementation guide is intended to be used.
  Jurisdiction []*CodeableConcept `json:"jurisdiction,omitempty"`

  // The base language in which the resource is written.
  Language string `json:"language,omitempty"`

  // The license that applies to this Implementation Guide, using an SPDX license code, or 'not-open-source'.
  License interface{} `json:"license,omitempty"`

  // Information about an assembled implementation guide, created by the publication tooling.
  Manifest *ImplementationGuideManifest `json:"manifest,omitempty"`

  // The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
  Meta *Meta `json:"meta,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // A natural language name identifying the implementation guide. This name should be usable as an identifier for the module by machine processing applications such as code generation.
  Name string `json:"name,omitempty"`

  // The NPM package name for this Implementation Guide, used in the NPM package distribution, which is the primary mechanism by which FHIR based tooling manages IG dependencies. This value must be globally unique, and should be assigned with care.
  PackageId string `json:"packageId,omitempty"`

  // The name of the organization or individual that published the implementation guide.
  Publisher string `json:"publisher,omitempty"`

  // This is a ImplementationGuide resource
  ResourceType interface{} `json:"resourceType"`

  // The status of this implementation guide. Enables tracking the life-cycle of the content.
  Status interface{} `json:"status,omitempty"`

  // A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
  Text *Narrative `json:"text,omitempty"`

  // A short, descriptive, user-friendly title for the implementation guide.
  Title string `json:"title,omitempty"`

  // An absolute URI that is used to identify this implementation guide when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this implementation guide is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the implementation guide is stored on different servers.
  Url string `json:"url,omitempty"`

  // The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate implementation guide instances.
  UseContext []*UsageContext `json:"useContext,omitempty"`

  // The identifier that is used to identify this version of the implementation guide when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the implementation guide author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
  Version string `json:"version,omitempty"`
}

func (strct *ImplementationGuide) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
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
    // Marshal the "definition" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"definition\": ")
	if tmp, err := json.Marshal(strct.Definition); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "dependsOn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"dependsOn\": ")
	if tmp, err := json.Marshal(strct.DependsOn); err != nil {
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
    // Marshal the "experimental" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"experimental\": ")
	if tmp, err := json.Marshal(strct.Experimental); err != nil {
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
    // Marshal the "fhirVersion" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"fhirVersion\": ")
	if tmp, err := json.Marshal(strct.FhirVersion); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "global" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"global\": ")
	if tmp, err := json.Marshal(strct.Global); err != nil {
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
    // Marshal the "license" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"license\": ")
	if tmp, err := json.Marshal(strct.License); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "manifest" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"manifest\": ")
	if tmp, err := json.Marshal(strct.Manifest); err != nil {
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
    // Marshal the "packageId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"packageId\": ")
	if tmp, err := json.Marshal(strct.PackageId); err != nil {
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

func (strct *ImplementationGuide) UnmarshalJSON(b []byte) error {
    resourceTypeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
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
        case "definition":
            if err := json.Unmarshal([]byte(v), &strct.Definition); err != nil {
                return err
             }
        case "dependsOn":
            if err := json.Unmarshal([]byte(v), &strct.DependsOn); err != nil {
                return err
             }
        case "description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
                return err
             }
        case "experimental":
            if err := json.Unmarshal([]byte(v), &strct.Experimental); err != nil {
                return err
             }
        case "extension":
            if err := json.Unmarshal([]byte(v), &strct.Extension); err != nil {
                return err
             }
        case "fhirVersion":
            if err := json.Unmarshal([]byte(v), &strct.FhirVersion); err != nil {
                return err
             }
        case "global":
            if err := json.Unmarshal([]byte(v), &strct.Global); err != nil {
                return err
             }
        case "id":
            if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
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
        case "license":
            if err := json.Unmarshal([]byte(v), &strct.License); err != nil {
                return err
             }
        case "manifest":
            if err := json.Unmarshal([]byte(v), &strct.Manifest); err != nil {
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
        case "packageId":
            if err := json.Unmarshal([]byte(v), &strct.PackageId); err != nil {
                return err
             }
        case "publisher":
            if err := json.Unmarshal([]byte(v), &strct.Publisher); err != nil {
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
        case "text":
            if err := json.Unmarshal([]byte(v), &strct.Text); err != nil {
                return err
             }
        case "title":
            if err := json.Unmarshal([]byte(v), &strct.Title); err != nil {
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
    // check if resourceType (a required property) was received
    if !resourceTypeReceived {
        return errors.New("\"resourceType\" is required but was not present")
    }
    return nil
}
