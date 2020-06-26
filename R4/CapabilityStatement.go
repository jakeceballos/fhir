// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "fmt"
    "errors"
    "bytes"
    "encoding/json"
)

// CapabilityStatement A Capability Statement documents a set of capabilities (behaviors) of a FHIR Server for a particular version of FHIR that may be used as a statement of actual server functionality or a statement of required or desired server implementation.
type CapabilityStatement struct {

  // Contact details to assist a user in finding and communicating with the publisher.
  Contact []*ContactDetail `json:"contact,omitempty"`

  // These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
  Contained []interface{} `json:"contained,omitempty"`

  // A copyright statement relating to the capability statement and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the capability statement.
  Copyright string `json:"copyright,omitempty"`

  // The date  (and optionally time) when the capability statement was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the capability statement changes.
  Date string `json:"date,omitempty"`

  // A free text natural language description of the capability statement from a consumer's perspective. Typically, this is used when the capability statement describes a desired rather than an actual solution, for example as a formal expression of requirements as part of an RFP.
  Description string `json:"description,omitempty"`

  // A document definition.
  Document []*CapabilityStatementDocument `json:"document,omitempty"`

  // A Boolean value to indicate that this capability statement is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
  Experimental bool `json:"experimental,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // The version of the FHIR specification that this CapabilityStatement describes (which SHALL be the same as the FHIR version of the CapabilityStatement itself). There is no default value.
  FhirVersion interface{} `json:"fhirVersion,omitempty"`

  // A list of the formats supported by this implementation using their content types.
  Format []string `json:"format,omitempty"`

  // The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
  Id string `json:"id,omitempty"`

  // Identifies a specific implementation instance that is described by the capability statement - i.e. a particular installation, rather than the capabilities of a software program.
  Implementation *CapabilityStatementImplementation `json:"implementation,omitempty"`

  // A list of implementation guides that the server does (or should) support in their entirety.
  ImplementationGuide []string `json:"implementationGuide,omitempty"`

  // A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
  ImplicitRules string `json:"implicitRules,omitempty"`

  // Reference to a canonical URL of another CapabilityStatement that this software adds to. The capability statement automatically includes everything in the other statement, and it is not duplicated, though the server may repeat the same resources, interactions and operations to add additional details to them.
  Imports []string `json:"imports,omitempty"`

  // Reference to a canonical URL of another CapabilityStatement that this software implements. This capability statement is a published API description that corresponds to a business service. The server may actually implement a subset of the capability statement it claims to implement, so the capability statement must specify the full capability details.
  Instantiates []string `json:"instantiates,omitempty"`

  // A legal or geographic region in which the capability statement is intended to be used.
  Jurisdiction []*CodeableConcept `json:"jurisdiction,omitempty"`

  // The way that this statement is intended to be used, to describe an actual running instance of software, a particular product (kind, not instance of software) or a class of implementation (e.g. a desired purchase).
  Kind interface{} `json:"kind,omitempty"`

  // The base language in which the resource is written.
  Language string `json:"language,omitempty"`

  // A description of the messaging capabilities of the solution.
  Messaging []*CapabilityStatementMessaging `json:"messaging,omitempty"`

  // The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
  Meta *Meta `json:"meta,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // A natural language name identifying the capability statement. This name should be usable as an identifier for the module by machine processing applications such as code generation.
  Name string `json:"name,omitempty"`

  // A list of the patch formats supported by this implementation using their content types.
  PatchFormat []string `json:"patchFormat,omitempty"`

  // The name of the organization or individual that published the capability statement.
  Publisher string `json:"publisher,omitempty"`

  // Explanation of why this capability statement is needed and why it has been designed as it has.
  Purpose string `json:"purpose,omitempty"`

  // This is a CapabilityStatement resource
  ResourceType interface{} `json:"resourceType"`

  // A definition of the restful capabilities of the solution, if any.
  Rest []*CapabilityStatementRest `json:"rest,omitempty"`

  // Software that is covered by this capability statement.  It is used when the capability statement describes the capabilities of a particular software version, independent of an installation.
  Software *CapabilityStatementSoftware `json:"software,omitempty"`

  // The status of this capability statement. Enables tracking the life-cycle of the content.
  Status interface{} `json:"status,omitempty"`

  // A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
  Text *Narrative `json:"text,omitempty"`

  // A short, descriptive, user-friendly title for the capability statement.
  Title string `json:"title,omitempty"`

  // An absolute URI that is used to identify this capability statement when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this capability statement is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the capability statement is stored on different servers.
  Url string `json:"url,omitempty"`

  // The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate capability statement instances.
  UseContext []*UsageContext `json:"useContext,omitempty"`

  // The identifier that is used to identify this version of the capability statement when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the capability statement author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
  Version string `json:"version,omitempty"`
}

func (strct *CapabilityStatement) MarshalJSON() ([]byte, error) {
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
    // Marshal the "document" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"document\": ")
	if tmp, err := json.Marshal(strct.Document); err != nil {
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
    // Marshal the "format" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"format\": ")
	if tmp, err := json.Marshal(strct.Format); err != nil {
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
    // Marshal the "implementation" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"implementation\": ")
	if tmp, err := json.Marshal(strct.Implementation); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "implementationGuide" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"implementationGuide\": ")
	if tmp, err := json.Marshal(strct.ImplementationGuide); err != nil {
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
    // Marshal the "imports" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"imports\": ")
	if tmp, err := json.Marshal(strct.Imports); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "instantiates" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"instantiates\": ")
	if tmp, err := json.Marshal(strct.Instantiates); err != nil {
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
    // Marshal the "kind" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"kind\": ")
	if tmp, err := json.Marshal(strct.Kind); err != nil {
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
    // Marshal the "messaging" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"messaging\": ")
	if tmp, err := json.Marshal(strct.Messaging); err != nil {
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
    // Marshal the "patchFormat" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"patchFormat\": ")
	if tmp, err := json.Marshal(strct.PatchFormat); err != nil {
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
    // Marshal the "purpose" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"purpose\": ")
	if tmp, err := json.Marshal(strct.Purpose); err != nil {
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
    // Marshal the "rest" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"rest\": ")
	if tmp, err := json.Marshal(strct.Rest); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "software" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"software\": ")
	if tmp, err := json.Marshal(strct.Software); err != nil {
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

func (strct *CapabilityStatement) UnmarshalJSON(b []byte) error {
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
        case "description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
                return err
             }
        case "document":
            if err := json.Unmarshal([]byte(v), &strct.Document); err != nil {
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
        case "format":
            if err := json.Unmarshal([]byte(v), &strct.Format); err != nil {
                return err
             }
        case "id":
            if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
                return err
             }
        case "implementation":
            if err := json.Unmarshal([]byte(v), &strct.Implementation); err != nil {
                return err
             }
        case "implementationGuide":
            if err := json.Unmarshal([]byte(v), &strct.ImplementationGuide); err != nil {
                return err
             }
        case "implicitRules":
            if err := json.Unmarshal([]byte(v), &strct.ImplicitRules); err != nil {
                return err
             }
        case "imports":
            if err := json.Unmarshal([]byte(v), &strct.Imports); err != nil {
                return err
             }
        case "instantiates":
            if err := json.Unmarshal([]byte(v), &strct.Instantiates); err != nil {
                return err
             }
        case "jurisdiction":
            if err := json.Unmarshal([]byte(v), &strct.Jurisdiction); err != nil {
                return err
             }
        case "kind":
            if err := json.Unmarshal([]byte(v), &strct.Kind); err != nil {
                return err
             }
        case "language":
            if err := json.Unmarshal([]byte(v), &strct.Language); err != nil {
                return err
             }
        case "messaging":
            if err := json.Unmarshal([]byte(v), &strct.Messaging); err != nil {
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
        case "patchFormat":
            if err := json.Unmarshal([]byte(v), &strct.PatchFormat); err != nil {
                return err
             }
        case "publisher":
            if err := json.Unmarshal([]byte(v), &strct.Publisher); err != nil {
                return err
             }
        case "purpose":
            if err := json.Unmarshal([]byte(v), &strct.Purpose); err != nil {
                return err
             }
        case "resourceType":
            if err := json.Unmarshal([]byte(v), &strct.ResourceType); err != nil {
                return err
             }
            resourceTypeReceived = true
        case "rest":
            if err := json.Unmarshal([]byte(v), &strct.Rest); err != nil {
                return err
             }
        case "software":
            if err := json.Unmarshal([]byte(v), &strct.Software); err != nil {
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