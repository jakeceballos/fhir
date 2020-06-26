// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// Contract Legally enforceable, formally recorded unilateral or bilateral directive i.e., a policy or agreement.
type Contract struct {

  // Alternative representation of the title for this Contract definition, derivative, or instance in any legal state., e.g., a domain specific contract number related to legislation.
  Alias []string `json:"alias,omitempty"`

  // Relevant time or time-period when this Contract is applicable.
  Applies *Period `json:"applies,omitempty"`

  // The individual or organization that authored the Contract definition, derivative, or instance in any legal state.
  Author *Reference `json:"author,omitempty"`

  // A formally or informally recognized grouping of people, principals, organizations, or jurisdictions formed for the purpose of achieving some form of collective action such as the promulgation, administration and enforcement of contracts and policies.
  Authority []*Reference `json:"authority,omitempty"`

  // These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
  Contained []interface{} `json:"contained,omitempty"`

  // Precusory content developed with a focus and intent of supporting the formation a Contract instance, which may be associated with and transformable into a Contract.
  ContentDefinition *ContractContentDefinition `json:"contentDefinition,omitempty"`

  // The minimal content derived from the basal information source at a specific stage in its lifecycle.
  ContentDerivative *CodeableConcept `json:"contentDerivative,omitempty"`

  // Recognized governance framework or system operating with a circumscribed scope in accordance with specified principles, policies, processes or procedures for managing rights, actions, or behaviors of parties or principals relative to resources.
  Domain []*Reference `json:"domain,omitempty"`

  // Event resulting in discontinuation or termination of this Contract instance by one or more parties to the contract.
  ExpirationType *CodeableConcept `json:"expirationType,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // The "patient friendly language" versionof the Contract in whole or in parts. "Patient friendly language" means the representation of the Contract and Contract Provisions in a manner that is readily accessible and understandable by a layperson in accordance with best practices for communication styles that ensure that those agreeing to or signing the Contract understand the roles, actions, obligations, responsibilities, and implication of the agreement.
  Friendly []*ContractFriendly `json:"friendly,omitempty"`

  // The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
  Id string `json:"id,omitempty"`

  // Unique identifier for this Contract or a derivative that references a Source Contract.
  Identifier []*Identifier `json:"identifier,omitempty"`

  // A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
  ImplicitRules string `json:"implicitRules,omitempty"`

  // The URL pointing to a FHIR-defined Contract Definition that is adhered to in whole or part by this Contract.
  InstantiatesCanonical *Reference `json:"instantiatesCanonical,omitempty"`

  // The URL pointing to an externally maintained definition that is adhered to in whole or in part by this Contract.
  InstantiatesUri string `json:"instantiatesUri,omitempty"`

  // When this  Contract was issued.
  Issued string `json:"issued,omitempty"`

  // The base language in which the resource is written.
  Language string `json:"language,omitempty"`

  // List of Legal expressions or representations of this Contract.
  Legal []*ContractLegal `json:"legal,omitempty"`

  // Legal states of the formation of a legal instrument, which is a formally executed written document that can be formally attributed to its author, records and formally expresses a legally enforceable act, process, or contractual duty, obligation, or right, and therefore evidences that act, process, or agreement.
  LegalState *CodeableConcept `json:"legalState,omitempty"`

  // Legally binding Contract: This is the signed and legally recognized representation of the Contract, which is considered the "source of truth" and which would be the basis for legal action related to enforcement of this Contract.
  LegallyBindingAttachment *Attachment `json:"legallyBindingAttachment,omitempty"`

  // Legally binding Contract: This is the signed and legally recognized representation of the Contract, which is considered the "source of truth" and which would be the basis for legal action related to enforcement of this Contract.
  LegallyBindingReference *Reference `json:"legallyBindingReference,omitempty"`

  // The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
  Meta *Meta `json:"meta,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // A natural language name identifying this Contract definition, derivative, or instance in any legal state. Provides additional information about its content. This name should be usable as an identifier for the module by machine processing applications such as code generation.
  Name string `json:"name,omitempty"`

  // Links to Provenance records for past versions of this Contract definition, derivative, or instance, which identify key state transitions or updates that are likely to be relevant to a user looking at the current version of the Contract.  The Provence.entity indicates the target that was changed in the update. http://build.fhir.org/provenance-definitions.html#Provenance.entity.
  RelevantHistory []*Reference `json:"relevantHistory,omitempty"`

  // This is a Contract resource
  ResourceType interface{} `json:"resourceType"`

  // List of Computable Policy Rule Language Representations of this Contract.
  Rule []*ContractRule `json:"rule,omitempty"`

  // A selector of legal concerns for this Contract definition, derivative, or instance in any legal state.
  Scope *CodeableConcept `json:"scope,omitempty"`

  // Parties with legal standing in the Contract, including the principal parties, the grantor(s) and grantee(s), which are any person or organization bound by the contract, and any ancillary parties, which facilitate the execution of the contract such as a notary or witness.
  Signer []*ContractSigner `json:"signer,omitempty"`

  // Sites in which the contract is complied with,  exercised, or in force.
  Site []*Reference `json:"site,omitempty"`

  // The status of the resource instance.
  Status string `json:"status,omitempty"`

  // Sub-category for the Contract that distinguishes the kinds of systems that would be interested in the Contract within the context of the Contract's scope.
  SubType []*CodeableConcept `json:"subType,omitempty"`

  // The target entity impacted by or of interest to parties to the agreement.
  Subject []*Reference `json:"subject,omitempty"`

  // An explanatory or alternate user-friendly title for this Contract definition, derivative, or instance in any legal state.t giving additional information about its content.
  Subtitle string `json:"subtitle,omitempty"`

  // Information that may be needed by/relevant to the performer in their execution of this term action.
  SupportingInfo []*Reference `json:"supportingInfo,omitempty"`

  // One or more Contract Provisions, which may be related and conveyed as a group, and may contain nested groups.
  Term []*ContractTerm `json:"term,omitempty"`

  // A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
  Text *Narrative `json:"text,omitempty"`

  // A short, descriptive, user-friendly title for this Contract definition, derivative, or instance in any legal state.t giving additional information about its content.
  Title string `json:"title,omitempty"`

  // Narrows the range of legal concerns to focus on the achievement of specific contractual objectives.
  TopicCodeableConcept *CodeableConcept `json:"topicCodeableConcept,omitempty"`

  // Narrows the range of legal concerns to focus on the achievement of specific contractual objectives.
  TopicReference *Reference `json:"topicReference,omitempty"`

  // A high-level category for the legal instrument, whether constructed as a Contract definition, derivative, or instance in any legal state.  Provides additional information about its content within the context of the Contract's scope to distinguish the kinds of systems that would be interested in the contract.
  Type *CodeableConcept `json:"type,omitempty"`

  // Canonical identifier for this contract, represented as a URI (globally unique).
  Url string `json:"url,omitempty"`

  // An edition identifier used for business purposes to label business significant variants.
  Version string `json:"version,omitempty"`
}

func (strct *Contract) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "alias" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"alias\": ")
	if tmp, err := json.Marshal(strct.Alias); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "applies" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"applies\": ")
	if tmp, err := json.Marshal(strct.Applies); err != nil {
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
    // Marshal the "authority" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"authority\": ")
	if tmp, err := json.Marshal(strct.Authority); err != nil {
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
    // Marshal the "contentDefinition" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"contentDefinition\": ")
	if tmp, err := json.Marshal(strct.ContentDefinition); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "contentDerivative" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"contentDerivative\": ")
	if tmp, err := json.Marshal(strct.ContentDerivative); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "domain" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"domain\": ")
	if tmp, err := json.Marshal(strct.Domain); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "expirationType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"expirationType\": ")
	if tmp, err := json.Marshal(strct.ExpirationType); err != nil {
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
    // Marshal the "friendly" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"friendly\": ")
	if tmp, err := json.Marshal(strct.Friendly); err != nil {
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
    // Marshal the "issued" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"issued\": ")
	if tmp, err := json.Marshal(strct.Issued); err != nil {
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
    // Marshal the "legal" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"legal\": ")
	if tmp, err := json.Marshal(strct.Legal); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "legalState" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"legalState\": ")
	if tmp, err := json.Marshal(strct.LegalState); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "legallyBindingAttachment" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"legallyBindingAttachment\": ")
	if tmp, err := json.Marshal(strct.LegallyBindingAttachment); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "legallyBindingReference" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"legallyBindingReference\": ")
	if tmp, err := json.Marshal(strct.LegallyBindingReference); err != nil {
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
    // Marshal the "relevantHistory" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"relevantHistory\": ")
	if tmp, err := json.Marshal(strct.RelevantHistory); err != nil {
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
    // Marshal the "rule" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"rule\": ")
	if tmp, err := json.Marshal(strct.Rule); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
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
    // Marshal the "signer" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"signer\": ")
	if tmp, err := json.Marshal(strct.Signer); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "site" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"site\": ")
	if tmp, err := json.Marshal(strct.Site); err != nil {
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
    // Marshal the "subType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"subType\": ")
	if tmp, err := json.Marshal(strct.SubType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
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
    // Marshal the "supportingInfo" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"supportingInfo\": ")
	if tmp, err := json.Marshal(strct.SupportingInfo); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "term" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"term\": ")
	if tmp, err := json.Marshal(strct.Term); err != nil {
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
    // Marshal the "topicCodeableConcept" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"topicCodeableConcept\": ")
	if tmp, err := json.Marshal(strct.TopicCodeableConcept); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "topicReference" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"topicReference\": ")
	if tmp, err := json.Marshal(strct.TopicReference); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "type" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"type\": ")
	if tmp, err := json.Marshal(strct.Type); err != nil {
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

func (strct *Contract) UnmarshalJSON(b []byte) error {
    resourceTypeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "alias":
            if err := json.Unmarshal([]byte(v), &strct.Alias); err != nil {
                return err
             }
        case "applies":
            if err := json.Unmarshal([]byte(v), &strct.Applies); err != nil {
                return err
             }
        case "author":
            if err := json.Unmarshal([]byte(v), &strct.Author); err != nil {
                return err
             }
        case "authority":
            if err := json.Unmarshal([]byte(v), &strct.Authority); err != nil {
                return err
             }
        case "contained":
            if err := json.Unmarshal([]byte(v), &strct.Contained); err != nil {
                return err
             }
        case "contentDefinition":
            if err := json.Unmarshal([]byte(v), &strct.ContentDefinition); err != nil {
                return err
             }
        case "contentDerivative":
            if err := json.Unmarshal([]byte(v), &strct.ContentDerivative); err != nil {
                return err
             }
        case "domain":
            if err := json.Unmarshal([]byte(v), &strct.Domain); err != nil {
                return err
             }
        case "expirationType":
            if err := json.Unmarshal([]byte(v), &strct.ExpirationType); err != nil {
                return err
             }
        case "extension":
            if err := json.Unmarshal([]byte(v), &strct.Extension); err != nil {
                return err
             }
        case "friendly":
            if err := json.Unmarshal([]byte(v), &strct.Friendly); err != nil {
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
        case "issued":
            if err := json.Unmarshal([]byte(v), &strct.Issued); err != nil {
                return err
             }
        case "language":
            if err := json.Unmarshal([]byte(v), &strct.Language); err != nil {
                return err
             }
        case "legal":
            if err := json.Unmarshal([]byte(v), &strct.Legal); err != nil {
                return err
             }
        case "legalState":
            if err := json.Unmarshal([]byte(v), &strct.LegalState); err != nil {
                return err
             }
        case "legallyBindingAttachment":
            if err := json.Unmarshal([]byte(v), &strct.LegallyBindingAttachment); err != nil {
                return err
             }
        case "legallyBindingReference":
            if err := json.Unmarshal([]byte(v), &strct.LegallyBindingReference); err != nil {
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
        case "relevantHistory":
            if err := json.Unmarshal([]byte(v), &strct.RelevantHistory); err != nil {
                return err
             }
        case "resourceType":
            if err := json.Unmarshal([]byte(v), &strct.ResourceType); err != nil {
                return err
             }
            resourceTypeReceived = true
        case "rule":
            if err := json.Unmarshal([]byte(v), &strct.Rule); err != nil {
                return err
             }
        case "scope":
            if err := json.Unmarshal([]byte(v), &strct.Scope); err != nil {
                return err
             }
        case "signer":
            if err := json.Unmarshal([]byte(v), &strct.Signer); err != nil {
                return err
             }
        case "site":
            if err := json.Unmarshal([]byte(v), &strct.Site); err != nil {
                return err
             }
        case "status":
            if err := json.Unmarshal([]byte(v), &strct.Status); err != nil {
                return err
             }
        case "subType":
            if err := json.Unmarshal([]byte(v), &strct.SubType); err != nil {
                return err
             }
        case "subject":
            if err := json.Unmarshal([]byte(v), &strct.Subject); err != nil {
                return err
             }
        case "subtitle":
            if err := json.Unmarshal([]byte(v), &strct.Subtitle); err != nil {
                return err
             }
        case "supportingInfo":
            if err := json.Unmarshal([]byte(v), &strct.SupportingInfo); err != nil {
                return err
             }
        case "term":
            if err := json.Unmarshal([]byte(v), &strct.Term); err != nil {
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
        case "topicCodeableConcept":
            if err := json.Unmarshal([]byte(v), &strct.TopicCodeableConcept); err != nil {
                return err
             }
        case "topicReference":
            if err := json.Unmarshal([]byte(v), &strct.TopicReference); err != nil {
                return err
             }
        case "type":
            if err := json.Unmarshal([]byte(v), &strct.Type); err != nil {
                return err
             }
        case "url":
            if err := json.Unmarshal([]byte(v), &strct.Url); err != nil {
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
