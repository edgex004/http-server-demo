// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    loggingSchema, err := UnmarshalLoggingSchema(bytes)
//    bytes, err = loggingSchema.Marshal()

package logging

import "bytes"
import "errors"
import "encoding/json"

func UnmarshalLoggingSchema(data []byte) (LoggingSchema, error) {
	var r LoggingSchema
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *LoggingSchema) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Container for prototype schema definitions.
type LoggingSchema struct {
	GetLogEntryResponse  *GetLogEntryResponse  `json:"GetLogEntryResponse,omitempty"` 
	HTTPValidationError  *HTTPValidationError  `json:"HTTPValidationError,omitempty"` 
	Location             *LocationClass        `json:"Location,omitempty"`            
	LogEntry             *LogEntry             `json:"LogEntry,omitempty"`            
	PostLogEntryResponse *PostLogEntryResponse `json:"PostLogEntryResponse,omitempty"`
	ValidationError      *ValidationError      `json:"ValidationError,omitempty"`     
}

type GetLogEntryResponse struct {
	Entry   LogEntry `json:"entry"`   
	EntryID int64    `json:"entry_id"`
}

type LogEntry struct {
	DeviceID    string         `json:"device_id"`            
	Entry       string         `json:"entry"`                
	Location    *LocationClass `json:"location,omitempty"`   
	Temperature *float64       `json:"temperature,omitempty"`
	Title       string         `json:"title"`                
}

type LocationClass struct {
	Latitude  float64 `json:"latitude"` 
	Longitude float64 `json:"longitude"`
}

type HTTPValidationError struct {
	Detail []ValidationError `json:"detail,omitempty"`
}

type ValidationError struct {
	LOC  []LOCElement `json:"loc"` 
	Msg  string       `json:"msg"` 
	Type string       `json:"type"`
}

type PostLogEntryResponse struct {
	EntryID int64 `json:"entry_id"`
}

type LOCElement struct {
	Integer *int64
	String  *string
}

func (x *LOCElement) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, &x.Integer, nil, nil, &x.String, false, nil, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *LOCElement) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, nil, nil, x.String, false, nil, false, nil, false, nil, false, nil, false)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
		*pi = nil
	}
	if pf != nil {
		*pf = nil
	}
	if pb != nil {
		*pb = nil
	}
	if ps != nil {
		*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return false, err
	}

	switch v := tok.(type) {
	case json.Number:
		if pi != nil {
			i, err := v.Int64()
			if err == nil {
				*pi = &i
				return false, nil
			}
		}
		if pf != nil {
			f, err := v.Float64()
			if err == nil {
				*pf = &f
				return false, nil
			}
			return false, errors.New("Unparsable number")
		}
		return false, errors.New("Union does not contain number")
	case float64:
		return false, errors.New("Decoder should not return float64")
	case bool:
		if pb != nil {
			*pb = &v
			return false, nil
		}
		return false, errors.New("Union does not contain bool")
	case string:
		if haveEnum {
			return false, json.Unmarshal(data, pe)
		}
		if ps != nil {
			*ps = &v
			return false, nil
		}
		return false, errors.New("Union does not contain string")
	case nil:
		if nullable {
			return false, nil
		}
		return false, errors.New("Union does not contain null")
	case json.Delim:
		if v == '{' {
			if haveObject {
				return true, json.Unmarshal(data, pc)
			}
			if haveMap {
				return false, json.Unmarshal(data, pm)
			}
			return false, errors.New("Union does not contain object")
		}
		if v == '[' {
			if haveArray {
				return false, json.Unmarshal(data, pa)
			}
			return false, errors.New("Union does not contain array")
		}
		return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")

}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
		return json.Marshal(*pi)
	}
	if pf != nil {
		return json.Marshal(*pf)
	}
	if pb != nil {
		return json.Marshal(*pb)
	}
	if ps != nil {
		return json.Marshal(*ps)
	}
	if haveArray {
		return json.Marshal(pa)
	}
	if haveObject {
		return json.Marshal(pc)
	}
	if haveMap {
		return json.Marshal(pm)
	}
	if haveEnum {
		return json.Marshal(pe)
	}
	if nullable {
		return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}
