// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// SubstanceSourceMaterial Source material shall capture information on the taxonomic and anatomical origins as well as the fraction of a material that can result in or can be modified to form a substance. This set of data elements shall be used to define polymer substances isolated from biological matrices. Taxonomic and anatomical origins shall be described using a controlled vocabulary as required. This information is captured for naturally derived polymers ( . starch) and structurally diverse substances. For Organisms belonging to the Kingdom Plantae the Substance level defines the fresh material of a single species or infraspecies, the Herbal Drug and the Herbal preparation. For Herbal preparations, the fraction information will be captured at the Substance information level and additional information for herbal extracts will be captured at the Specified Substance Group 1 information level. See for further explanation the Substance Class: Structurally Diverse and the herbal annex.
type SubstanceSourceMaterial struct {

  // These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
  Contained []interface{} `json:"contained,omitempty"`

  // The country where the plant material is harvested or the countries where the plasma is sourced from as laid down in accordance with the Plasma Master File. For “Plasma-derived substances” the attribute country of origin provides information about the countries used for the manufacturing of the Cryopoor plama or Crioprecipitate.
  CountryOfOrigin []*CodeableConcept `json:"countryOfOrigin,omitempty"`

  // Stage of life for animals, plants, insects and microorganisms. This information shall be provided only when the substance is significantly different in these stages (e.g. foetal bovine serum).
  DevelopmentStage *CodeableConcept `json:"developmentStage,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Many complex materials are fractions of parts of plants, animals, or minerals. Fraction elements are often necessary to define both Substances and Specified Group 1 Substances. For substances derived from Plants, fraction information will be captured at the Substance information level ( . Oils, Juices and Exudates). Additional information for Extracts, such as extraction solvent composition, will be captured at the Specified Substance Group 1 information level. For plasma-derived products fraction information will be captured at the Substance and the Specified Substance Group 1 levels.
  FractionDescription []*SubstanceSourceMaterialFractionDescription `json:"fractionDescription,omitempty"`

  // The place/region where the plant is harvested or the places/regions where the animal source material has its habitat.
  GeographicalLocation []string `json:"geographicalLocation,omitempty"`

  // The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
  Id string `json:"id,omitempty"`

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

  // This subclause describes the organism which the substance is derived from. For vaccines, the parent organism shall be specified based on these subclause elements. As an example, full taxonomy will be described for the Substance Name: ., Leaf.
  Organism *SubstanceSourceMaterialOrganism `json:"organism,omitempty"`

  // The unique identifier associated with the source material parent organism shall be specified.
  OrganismId *Identifier `json:"organismId,omitempty"`

  // The organism accepted Scientific name shall be provided based on the organism taxonomy.
  OrganismName string `json:"organismName,omitempty"`

  // The parent of the herbal drug Ginkgo biloba, Leaf is the substance ID of the substance (fresh) of Ginkgo biloba L. or Ginkgo biloba L. (Whole plant).
  ParentSubstanceId []*Identifier `json:"parentSubstanceId,omitempty"`

  // The parent substance of the Herbal Drug, or Herbal preparation.
  ParentSubstanceName []string `json:"parentSubstanceName,omitempty"`

  // To do.
  PartDescription []*SubstanceSourceMaterialPartDescription `json:"partDescription,omitempty"`

  // This is a SubstanceSourceMaterial resource
  ResourceType interface{} `json:"resourceType"`

  // General high level classification of the source material specific to the origin of the material.
  SourceMaterialClass *CodeableConcept `json:"sourceMaterialClass,omitempty"`

  // The state of the source material when extracted.
  SourceMaterialState *CodeableConcept `json:"sourceMaterialState,omitempty"`

  // The type of the source material shall be specified based on a controlled vocabulary. For vaccines, this subclause refers to the class of infectious agent.
  SourceMaterialType *CodeableConcept `json:"sourceMaterialType,omitempty"`

  // A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
  Text *Narrative `json:"text,omitempty"`
}

func (strct *SubstanceSourceMaterial) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
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
    // Marshal the "countryOfOrigin" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"countryOfOrigin\": ")
	if tmp, err := json.Marshal(strct.CountryOfOrigin); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "developmentStage" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"developmentStage\": ")
	if tmp, err := json.Marshal(strct.DevelopmentStage); err != nil {
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
    // Marshal the "fractionDescription" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"fractionDescription\": ")
	if tmp, err := json.Marshal(strct.FractionDescription); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "geographicalLocation" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"geographicalLocation\": ")
	if tmp, err := json.Marshal(strct.GeographicalLocation); err != nil {
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
    // Marshal the "organism" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"organism\": ")
	if tmp, err := json.Marshal(strct.Organism); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "organismId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"organismId\": ")
	if tmp, err := json.Marshal(strct.OrganismId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "organismName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"organismName\": ")
	if tmp, err := json.Marshal(strct.OrganismName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "parentSubstanceId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"parentSubstanceId\": ")
	if tmp, err := json.Marshal(strct.ParentSubstanceId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "parentSubstanceName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"parentSubstanceName\": ")
	if tmp, err := json.Marshal(strct.ParentSubstanceName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "partDescription" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"partDescription\": ")
	if tmp, err := json.Marshal(strct.PartDescription); err != nil {
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
    // Marshal the "sourceMaterialClass" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"sourceMaterialClass\": ")
	if tmp, err := json.Marshal(strct.SourceMaterialClass); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "sourceMaterialState" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"sourceMaterialState\": ")
	if tmp, err := json.Marshal(strct.SourceMaterialState); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "sourceMaterialType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"sourceMaterialType\": ")
	if tmp, err := json.Marshal(strct.SourceMaterialType); err != nil {
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

func (strct *SubstanceSourceMaterial) UnmarshalJSON(b []byte) error {
    resourceTypeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "contained":
            if err := json.Unmarshal([]byte(v), &strct.Contained); err != nil {
                return err
             }
        case "countryOfOrigin":
            if err := json.Unmarshal([]byte(v), &strct.CountryOfOrigin); err != nil {
                return err
             }
        case "developmentStage":
            if err := json.Unmarshal([]byte(v), &strct.DevelopmentStage); err != nil {
                return err
             }
        case "extension":
            if err := json.Unmarshal([]byte(v), &strct.Extension); err != nil {
                return err
             }
        case "fractionDescription":
            if err := json.Unmarshal([]byte(v), &strct.FractionDescription); err != nil {
                return err
             }
        case "geographicalLocation":
            if err := json.Unmarshal([]byte(v), &strct.GeographicalLocation); err != nil {
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
        case "organism":
            if err := json.Unmarshal([]byte(v), &strct.Organism); err != nil {
                return err
             }
        case "organismId":
            if err := json.Unmarshal([]byte(v), &strct.OrganismId); err != nil {
                return err
             }
        case "organismName":
            if err := json.Unmarshal([]byte(v), &strct.OrganismName); err != nil {
                return err
             }
        case "parentSubstanceId":
            if err := json.Unmarshal([]byte(v), &strct.ParentSubstanceId); err != nil {
                return err
             }
        case "parentSubstanceName":
            if err := json.Unmarshal([]byte(v), &strct.ParentSubstanceName); err != nil {
                return err
             }
        case "partDescription":
            if err := json.Unmarshal([]byte(v), &strct.PartDescription); err != nil {
                return err
             }
        case "resourceType":
            if err := json.Unmarshal([]byte(v), &strct.ResourceType); err != nil {
                return err
             }
            resourceTypeReceived = true
        case "sourceMaterialClass":
            if err := json.Unmarshal([]byte(v), &strct.SourceMaterialClass); err != nil {
                return err
             }
        case "sourceMaterialState":
            if err := json.Unmarshal([]byte(v), &strct.SourceMaterialState); err != nil {
                return err
             }
        case "sourceMaterialType":
            if err := json.Unmarshal([]byte(v), &strct.SourceMaterialType); err != nil {
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
