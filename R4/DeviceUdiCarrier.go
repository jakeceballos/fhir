// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "encoding/json"
    "fmt"
)

// DeviceUdiCarrier A type of a manufactured item that is used in the provision of healthcare without being substantially changed through that activity. The device may be a medical or non-medical device.
type DeviceUdiCarrier struct {

  // The full UDI carrier of the Automatic Identification and Data Capture (AIDC) technology representation of the barcode string as printed on the packaging of the device - e.g., a barcode or RFID.   Because of limitations on character sets in XML and the need to round-trip JSON data through XML, AIDC Formats *SHALL* be base64 encoded.
  CarrierAIDC string `json:"carrierAIDC,omitempty"`

  // The full UDI carrier as the human readable form (HRF) representation of the barcode string as printed on the packaging of the device.
  CarrierHRF string `json:"carrierHRF,omitempty"`

  // The device identifier (DI) is a mandatory, fixed portion of a UDI that identifies the labeler and the specific version or model of a device.
  DeviceIdentifier string `json:"deviceIdentifier,omitempty"`

  // A coded entry to indicate how the data was entered.
  EntryType interface{} `json:"entryType,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // Organization that is charged with issuing UDIs for devices.  For example, the US FDA issuers include :
  // 1) GS1: 
  // http://hl7.org/fhir/NamingSystem/gs1-di, 
  // 2) HIBCC:
  // http://hl7.org/fhir/NamingSystem/hibcc-dI, 
  // 3) ICCBBA for blood containers:
  // http://hl7.org/fhir/NamingSystem/iccbba-blood-di, 
  // 4) ICCBA for other devices:
  // http://hl7.org/fhir/NamingSystem/iccbba-other-di.
  Issuer string `json:"issuer,omitempty"`

  // The identity of the authoritative source for UDI generation within a  jurisdiction.  All UDIs are globally unique within a single namespace with the appropriate repository uri as the system.  For example,  UDIs of devices managed in the U.S. by the FDA, the value is  http://hl7.org/fhir/NamingSystem/fda-udi.
  Jurisdiction string `json:"jurisdiction,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
}

func (strct *DeviceUdiCarrier) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "carrierAIDC" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"carrierAIDC\": ")
	if tmp, err := json.Marshal(strct.CarrierAIDC); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "carrierHRF" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"carrierHRF\": ")
	if tmp, err := json.Marshal(strct.CarrierHRF); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "deviceIdentifier" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"deviceIdentifier\": ")
	if tmp, err := json.Marshal(strct.DeviceIdentifier); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "entryType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"entryType\": ")
	if tmp, err := json.Marshal(strct.EntryType); err != nil {
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
    // Marshal the "issuer" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"issuer\": ")
	if tmp, err := json.Marshal(strct.Issuer); err != nil {
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

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *DeviceUdiCarrier) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "carrierAIDC":
            if err := json.Unmarshal([]byte(v), &strct.CarrierAIDC); err != nil {
                return err
             }
        case "carrierHRF":
            if err := json.Unmarshal([]byte(v), &strct.CarrierHRF); err != nil {
                return err
             }
        case "deviceIdentifier":
            if err := json.Unmarshal([]byte(v), &strct.DeviceIdentifier); err != nil {
                return err
             }
        case "entryType":
            if err := json.Unmarshal([]byte(v), &strct.EntryType); err != nil {
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
        case "issuer":
            if err := json.Unmarshal([]byte(v), &strct.Issuer); err != nil {
                return err
             }
        case "jurisdiction":
            if err := json.Unmarshal([]byte(v), &strct.Jurisdiction); err != nil {
                return err
             }
        case "modifierExtension":
            if err := json.Unmarshal([]byte(v), &strct.ModifierExtension); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}
