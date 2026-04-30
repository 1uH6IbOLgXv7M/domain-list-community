package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/v2fly/domain-list-community/internal/dlc"
	router "github.com/v2fly/v2ray-core/v5/app/router/routercommon"
	"google.golang.org/protobuf/proto"
)

var (
	inputData   = flag.String("inputdata", "dlc.dat", "Name of the geosite dat file")
	outputDir   = flag.String("outputdir", "./output", "Directory to place all generated files")
	exportLists = flag.String("exportlists", "", "Lists to be exported, separated by ',' (empty for _all_)")
)

type DomainRule struct {
	Type  string
	Value string
	Attrs []string
}

type DomainList struct {
	Name  string
	Rules []DomainRule
}

func (d *DomainRule) domain2String() string {
	var dstr strings.Builder
	dstr.Grow(len(d.Type) + len(d.Value) + 10)
	dstr.WriteString(d.Type)
	dstr.WriteByte(':')
	dstr.WriteString(d.Value)
	for i, attr := range d.Attrs {
		if i == 0 {
			dstr.WriteByte(':')
		} else {
			dstr.WriteByte(',')
		}
		dstr.WriteByte('@')
		dstr.WriteString(attr)
	}
	return dstr.String()
}

func loadGeosite(path string) ([]DomainList, map[string]*DomainList, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read geosite file: %w", err)
	}
	vgeositeList := new(router.GeoSiteList)
	if err := proto.Unmarshal(data, vgeositeList); err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	domainLists := make([]DomainList, len(vgeositeList.Entry))
	domainListByName := make(map[string]*DomainList, len(vgeositeList.Entry))
	for i, vsite := range vgeositeList.Entry {
		rules := make([]DomainRule, 0, len(vsite.Domain))
		for _, vdomain := range vsite.Domain {
			rule := DomainRule{Value: vdomain.Value}
			switch vdomain.Type {
			case router.Domain_RootDomain:
				rule.Type = dlc.RuleTypeDomain
			case router.Domain_Regex:
				rule.Type = dlc.RuleTypeRegexp
			case router.Domain_Plain:
				rule.Type = dlc.RuleTypeKeyword
			case router.Domain_Full:
				rule.Type = dlc.RuleTypeFullDomain
			default:
				return nil, nil, fmt.Errorf("invalid rule type: %+v", vdomain.Type)
			}
			for _, vattr := range vdomain.Attribute {
				rule.Attrs = append(rule.Attrs, vattr.Key)
			}
			rules = append(rules, rule)
		}
		domainLists[i] = DomainList{
			Name:  strings.ToUpper(vsite.CountryCode),
			Rules: rules,
		}
		domainListByName[domainLists[i].Name] = &domainLists[i]
	}
	return domainLists, domainListByName, nil
}

func exportSite(name string, domainListByName map[string]*DomainList) error {
	domainList, ok := domainListByName[strings.ToUpper(name)]
	if !ok {
		return fmt.Errorf("list %q does not exist", name)
	}
	if len(domainList.Rules) == 0 {
		return fmt.Errorf("list %q is empty", name)
	}
	file, err := os.Create(filepath.Join(*outputDir, name+".yml"))
	if err != nil {
		return err
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	fmt.Fprintf(w, "%s:\n", name)
	for _, domain := range domainList.Rules {
		fmt.Fprintf(w, "  - %q\n", domain.domain2String())
	}
	return w