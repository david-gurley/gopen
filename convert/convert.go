package convert

import (
	"encoding/csv"
	"github.com/PaloAltoNetworks/pango/poli/security"
	"github.com/gocarina/gocsv"
	"github.com/pensando/go-psm/psm_dss"
	"io"
	"io/ioutil"
)

const (
	DEFAULT_POLICY = "default"
)

// ImportBytesCSV => ConvertCSV => Convert
// import files or others sources to byes arrays first. this allows us to take bytes
// from for example a browser through an API.  then parse the bytes into strings.
// finally, convert strings to final data structures.

type ConfigFromCSVFiles struct {
	Entries       string `json:"entries"`
	AddressGroups string `json:"address_groups"`
	ServiceGroups string `json:"service_groups'`
	Zones         string `json:"zones"`
}

type ImportCSVBytes struct {
	Entries       []byte `json:"entries"`
	AddressGroups []byte `json:"address_groups"`
	ServiceGroups []byte `json:"service_groups"`
	Zones         []byte `json:"zones"`
}

func NewImportCSVBytesFromConfig(config *ConfigFromCSVFiles) (*ImportCSVBytes, error) {
	var importCSVBytes ImportCSVBytes
	var err error
	if config.Entries != "" {
		importCSVBytes.Entries, err = ioutil.ReadFile(config.Entries)
		if err != nil {
			return &importCSVBytes, err
		}
	}
	if config.AddressGroups != "" {
		importCSVBytes.AddressGroups, err = ioutil.ReadFile(config.AddressGroups)
		if err != nil {
			return &importCSVBytes, err
		}
	}
	if config.ServiceGroups != "" {
	}
	if config.Zones != "" {
		importCSVBytes.Zones, err = ioutil.ReadFile(config.Zones)
		if err != nil {
			return &importCSVBytes, err
		}
	}
	return &importCSVBytes, nil
}

func (importCSVBytes *ImportCSVBytes) Convert() (*ConvertCSV, error) {
	var convertCSV ConvertCSV
	if importCSVBytes.Entries != nil {
		entries := []*EntryCSV{}
		err := gocsv.UnmarshalBytes(importCSVBytes.Entries, &entries)
		if err != nil {
			// if we get an error, try LazyQuotes
			gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
				return gocsv.LazyCSVReader(in)
			})
			err := gocsv.UnmarshalBytes(importCSVBytes.Entries, &entries)
			if err != nil {
				return &convertCSV, err
			}
			convertCSV.Entries = entries
		} else {
			convertCSV.Entries = entries
		}

	}
	if importCSVBytes.AddressGroups != nil {
		addressGroups := []*AddressGroupCSV{}
		err := gocsv.UnmarshalBytes(importCSVBytes.AddressGroups, &addressGroups)
		if err != nil {
			gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
				return gocsv.LazyCSVReader(in)
			})
			err := gocsv.UnmarshalBytes(importCSVBytes.AddressGroups, &addressGroups)
			if err != nil {
				return &convertCSV, err
			}
			convertCSV.AddressGroups = addressGroups
		} else {
			convertCSV.AddressGroups = addressGroups
		}
	}
	if importCSVBytes.ServiceGroups != nil {
	}
	if importCSVBytes.Zones != nil {
		zones := []*ZoneCSV{}
		err := gocsv.UnmarshalBytes(importCSVBytes.Zones, &zones)
		if err != nil {
			gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
				r := csv.NewReader(in)
				r.LazyQuotes = true
				return r
			})
			err := gocsv.UnmarshalBytes(importCSVBytes.Zones, &zones)
			if err != nil {
				return &convertCSV, err
			}
			convertCSV.Zones = zones
		} else {
			convertCSV.Zones = zones
		}
	}
	return &convertCSV, nil
}

type ConvertCSV struct {
	Entries       []*EntryCSV        `json:"entries"`
	AddressGroups []*AddressGroupCSV `json:"address_groups"`
	ServiceGroups []*ServiceGroupCSV `json:"service_groups"`
	Zones         []*ZoneCSV         `json:"zone"`
}

func ConvertCSVFromFiles(config *ConfigFromCSVFiles) (*ConvertCSV, error) {
	importCSVBytes, err := NewImportCSVBytesFromConfig(config)
	if err != nil {
		return &ConvertCSV{}, err
	}
	return importCSVBytes.Convert()
}

// could end up with multiple policies enforced on different VRFs or VLANs, but for now
// we will generate one policy DEFAULT_POLICY
type Convert struct {
	Entries       []*security.Entry                                 `json:"entries"`
	AddressGroups map[string]*AddressGroup                          `json:"address_groups"`
	ServiceGroups map[string]*ServiceGroup                          `json:"service_groups"`
	Zones         map[string]*Zone                                  `json:"zones"`
	DssPolicy     map[string]*psm_dss.SecurityNetworkSecurityPolicy `json:"dss_policy"`
}

// making this a method so we have access to groupings and other metadata.
// return error if the entry is not convertable to penando security group rule.
func (convert *Convert) EntryToDssRule(entry *security.Entry) (*psm_dss.SecuritySGRule, error) {
	var rule psm_dss.SecuritySGRule
	if IsEntryDisabled(entry) {
		return &rule, ErrConvertEntryDisabled
	}
	if IsEntryUser(entry) {
		return &rule, ErrConvertEntryUser
	}
	if IsEntryDevice(entry) {
		return &rule, ErrConvertEntryDevice
	}
	if IsEntryAppID(entry) {
		return &rule, ErrConvertSourceEntryAppID
	}
	// action
	if entry.Action != "" {
		rule.SetAction(entry.Action)
	}
	// source address
	sourceAddressGroups := EntrySourceAddressGroupNames(entry)
	if len(sourceAddressGroups) > 0 {
		found := false
		for _, sourceAddressGroup := range sourceAddressGroups {
			if _, ok := convert.AddressGroups[sourceAddressGroup]; ok {
				found = true
				break
			}
		}
		if !found {
			return &rule, ErrConvertAddressSrcMissingMapping
		}
	}
	// destination address
	destinationAddressGroups := EntryDestinationAddressGroupNames(entry)
	if len(destinationAddressGroups) > 0 {
		found := false
		for _, destinationAddressGroup := range destinationAddressGroups {
			if _, ok := convert.AddressGroups[destinationAddressGroup]; ok {
				found = true
				break
			}
		}
		if !found {
			return &rule, ErrConvertAddressDstMissingMapping
		}
	}
	// service
	if len(entry.Services) > 0 {
		for _, service := range entry.Services {
			dssProtoPort, match := EntryServiceToDssProtoPort(service)
			if match {
				if rule.HasProtoPorts() {
					rule.SetProtoPorts(append(*rule.ProtoPorts, dssProtoPort))
				} else {
					rule.SetProtoPorts([]psm_dss.SecurityProtoPort{dssProtoPort})
				}
			} else {
			}
		}
	}
	// we can skip looking up zone mapppings
	if IsEntryZoneAnyAny(entry) {
	}
	return &rule, nil
}

func (convert *Convert) NumZoneAnyAny() int {
	num := 0
	for _, entry := range convert.Entries {
		if IsEntryZoneAnyAny(entry) {
			num++
		}
	}
	return num
}
func (convert *Convert) NumZoneNotAnyAny() int {
	num := 0
	for _, entry := range convert.Entries {
		if IsEntryZoneAnyAny(entry) {
		} else {
			num++
		}
	}
	return num
}

func (convert *Convert) NumAppID() int {
	num := 0
	for _, entry := range convert.Entries {
		if IsEntryAppID(entry) {
			num++
		}
	}
	return num
}

func (convert *Convert) NumNotAppID() int {
	num := 0
	for _, entry := range convert.Entries {
		if IsEntryAppID(entry) {
		} else {
			num++
		}
	}
	return num
}

func (convert *Convert) NumDisabled() int {
	num := 0
	for _, entry := range convert.Entries {
		if IsEntryDisabled(entry) {
			num++
		}
	}
	return num
}

func NewConvert() *Convert {
	convert := Convert{
		Entries:       make([]*security.Entry, 0),
		AddressGroups: make(map[string]*AddressGroup),
		ServiceGroups: make(map[string]*ServiceGroup),
		Zones:         make(map[string]*Zone),
		DssPolicy:     make(map[string]*psm_dss.SecurityNetworkSecurityPolicy),
	}
	convert.DssPolicy[DEFAULT_POLICY] = psm_dss.NewSecurityNetworkSecurityPolicy()
	convert.DssPolicy[DEFAULT_POLICY].Spec = psm_dss.NewSecurityNetworkSecurityPolicySpec()
	return &convert
}
func NewFromCSVFiles(config *ConfigFromCSVFiles) (*Convert, error) {
	// read files into byte arrays for processing
	conv := NewConvert()
	convertCSV, err := ConvertCSVFromFiles(config)
	if err != nil {
		return conv, err
	}
	for _, entry := range convertCSV.Entries {
		// this parses fields - e.g. split strings to arrays
		e, err := entry.Convert()
		if err != nil {
			continue
		}
		conv.Entries = append(conv.Entries, e)
	}
	for _, addressGroup := range convertCSV.AddressGroups {
		a, err := addressGroup.Convert()
		if err != nil {
			continue
		}
		conv.AddressGroups[addressGroup.Name] = a
	}
	for _, zone := range convertCSV.Zones {
		z, err := zone.Convert()
		if err != nil {
			continue
		}
		conv.Zones[zone.Name] = z
	}
	for _, entry := range conv.Entries {
		for _, zone := range entry.SourceZones {
			if _, ok := conv.Zones[zone]; ok {
				conv.Zones[zone].EntryCount++
			} else {
				conv.Zones[zone] = &Zone{
					Name:       zone,
					EntryCount: 1,
				}
			}
		}
		for _, zone := range entry.DestinationZones {
			if _, ok := conv.Zones[zone]; ok {
				conv.Zones[zone].EntryCount++
			} else {
				conv.Zones[zone] = &Zone{
					Name:       zone,
					EntryCount: 1,
				}
			}
		}
		for _, addressGroup := range entry.SourceAddresses {
			if _, ok := conv.AddressGroups[addressGroup]; ok {
				conv.AddressGroups[addressGroup].EntryCount++
			}
		}
		for _, addressGroup := range entry.DestinationAddresses {
			if _, ok := conv.AddressGroups[addressGroup]; ok {
				conv.AddressGroups[addressGroup].EntryCount++
			}
		}
	}
	return conv, nil
}

func (convert *Convert) ToDss() []error {
	errors := make([]error, 0)
	for _, entry := range convert.Entries {
		rule, err := convert.EntryToDssRule(entry)
		if err != nil {
			errors = append(errors, err)
		} else if rule != nil {
			if convert.DssPolicy[DEFAULT_POLICY].Spec.HasRules() {
				convert.DssPolicy[DEFAULT_POLICY].Spec.SetRules(append(*convert.DssPolicy[DEFAULT_POLICY].Spec.Rules, *rule))
			} else {
				convert.DssPolicy[DEFAULT_POLICY].Spec.SetRules([]psm_dss.SecuritySGRule{*rule})
			}
		}
	}
	return errors
}
