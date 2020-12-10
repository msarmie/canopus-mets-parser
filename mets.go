package main

import (
  "fmt"
  "flag"
  "log"
  "encoding/xml"
  "encoding/json"
  "io/ioutil"
  "os"
  "strings"
  "strconv"
)

// ********* JSON Structs *********

// New
type ObjectMetsManifest struct { // TODO Data structure name may change
	Title               string           `json:"title"`
	JiraTicketNumber    string           `json:"jira_ticket_number"`
	DepartmentOrLibrary string           `json:"department_or_library"`
	CollectionCall      string           `json:"collection_call"`
	DepositorName       string           `json:"depositor_name"`
	BaggingDate         string           `json:"bagging_date"`
	Description         string           `json:"description"`
	SfErrors            string           `json:"sf_errors"`
	NewTarTechMD        NewTarTechMd     `json:"tar_techMD"`
	ManifestSha256      string           `json:"manifest_sha256"`
	ManifestMd5         string           `json:"manifest_md5"`
	Manifest            manifestMetsJSON `json:"manifest"`
	StorageLocation     string           `json:"storage_location"`
	FileCount           int64            `json:"file_count"`
	SchemaVersion       string           `json:"schema_version"`
}

// NewTarTechMd represents the Tar Tech MD used in Object Metadata
type NewTarTechMd struct {
	Siegfried   string        `json:"siegfried"`
	Scandate    string        `json:"scandate"`
	Signature   string        `json:"signature"`
	Created     string        `json:"created"`
	Identifiers []Identifiers `json:"identifiers"`
	Files       []Files       `json:"files"`
}

// Files represents the Manifest Files Array Structure
type Files struct {
	FileName string        `json:"filename"`
	FileSize int64         `json:"filesize"`
	Modified string        `json:"modified"`
	Errors   string        `json:"errors"`
	Md5      string        `json:"md5"`
  Sha256   string        `json:"sha256"`
	Matches  []Matches     `json:"matches"`
}

// Identifiers represents the Identifiers Structure used in manifest file
type Identifiers struct {
	Name    string `json:"name"`
	Details string `json:"details"`
}

// Matches represents Matches Structure
type Matches struct {
	Ns      string `json:"ns"`
	ID      string `json:"id"`
	Format  string `json:"format"`
	Version string `json:"version"`
	Mime    string `json:"mime"`
	Basis   string `json:"basis"`
	Warning string `json:"warning"`
}

// New
type manifestMetsJSON struct {
	Siegfried   string        `json:"siegfried"`
	Scandate    string        `json:"scandate"`
	Signature   string        `json:"signature"`
	Created     string        `json:"created"`
	Identifiers []Identifiers `json:"identifiers"`
	Files       []FilesMets   `json:"files"`
}

// New
type FilesMets struct {
	FileName      string        `json:"filename"`
	FileSize      int64         `json:"filesize"`
	Modified      string        `json:"modified"`
	Errors        string        `json:"errors"`
	Md5           string        `json:"md5"`
  Sha256        string        `json:"sha256"`
	Matches       []Matches     `json:"matches"`
	DescriptiveMD descriptiveMD `json:"descriptiveMD"`
}

type descriptiveMD struct {
  XMLName               xml.Name `xml:"dublincore" json:"-"`
  Identifier            string   `xml:"identifier" json:"identifier"`
  Title                 string   `xml:"title" json:"title"`
  Creator               string   `xml:"creator" json:"creator"`
  Date                  string   `xml:"date" json:"date"`
  Type                  string   `xml:"type" json:"type"`
  Format                string   `xml:"format" json:"format"`
  LanguageArr           []string `xml:"language" json:"-"`
  Language              string   `json:"language"`
  Contributor           string   `xml:"contributor" json:"contributor"`
  Provenance            string   `xml:"provenance" json:"provenance"`
  SubjectArr            []string `xml:"subject" json:"-"`
  Subject               string   `json:"subject"`
  Description           string   `xml:"description" json:"description"`
  Publisher             string   `xml:"publisher" json:"publisher"`
  Source                string   `xml:"source" json:"source"`
  Relation              string   `xml:"relation" json:"relation"`
  Coverage              string   `xml:"converge" json:"converge"`
  Rights                string   `xml:"rights" json:"rights"`
  IsPartOf              string   `xml:"isPartOf,omitempty" json:"isPartOf,omitempty"`
  Abstract              string   `xml:"abstract,omitempty" json:"abstract,,omitempty"`
  AccessRights          string   `xml:"accessRights,omitempty" json:"accessRights,omitempty"`
  AccrualMethod         string   `xml:"accrualMethod,omitempty" json:"accrualMethod,omitempty"`
  AccrualPeriodicity    string   `xml:"accrualPeriodicity,omitempty" json:"accrualPeriodicity,omitempty"`
  AccrualPolicy         string   `xml:"accrualPolicy,omitempty" json:"accrualPolicy,omitempty"`
  Alternative           string   `xml:"alternative,omitempty" json:"alternative,omitempty"`
  Audience              string   `xml:"audience,omitempty" json:"audience,omitempty"`
  Available             string   `xml:"available,omitempty" json:"available,omitempty"`
  BibliographicCitation string   `xml:"bibliographicCitation,omitempty" json:"bibliographicCitation,omitempty"`
  ConformsTo            string   `xml:"conformsTo,omitempty" json:"conformsTo,omitempty"`
  Created               string   `xml:"created,omitempty" json:"created,omitempty"`
  DateAccepted          string   `xml:"dateAccepted,omitempty" json:"dateAccepted,omitempty"`
  DateCopyrighted       string   `xml:"dateCopyrighted,omitempty" json:"dateCopyrighted,omitempty"`
  DateSubmitted         string   `xml:"dateSubmitted,omitempty" json:"dateSubmitted,omitempty"`
  EducationLevel        string   `xml:"educationLevel,omitempty" json:"educationLevel,omitempty"`
  Extent                string   `xml:"extent,omitempty" json:"extent,omitempty"`
  HasFormat             string   `xml:"hasFormat,omitempty" json:"hasFormat,omitempty"`
  HasPart               string   `xml:"hasPart,omitempty" json:"hasPart,omitempty"`
  HasVersion            string   `xml:"hasVersion,omitempty" json:"hasVersion,omitempty"`
  InstructionalMethod   string   `xml:"instructionalMethod,omitempty" json:"instructionalMethod,omitempty"`
  IsFormatOf            string   `xml:"isFormatOf,omitempty" json:"isFormatOf,omitempty"`
  IsReferencedBy        string   `xml:"isReferencedBy,omitempty" json:"isReferencedBy,omitempty"`
  IsReplacedBy          string   `xml:"isReplacedBy,omitempty" json:"isReplacedBy,omitempty"`
  IsRequiredBy          string   `xml:"isRequiredBy,omitempty" json:"isRequiredBy,omitempty"`
  Issued                string   `xml:"issued,omitempty" json:"issued,omitempty"`
  IsVersionOf           string   `xml:"isVersionOf,omitempty" json:"isVersionOf,omitempty"`
  License               string   `xml:"license,omitempty" json:"license,omitempty"`
  Mediator              string   `xml:"mediator,omitempty" json:"mediator,omitempty"`
  Modified              string   `xml:"modified,omitempty" json:"modified,omitempty"`
  References            string   `xml:"references,omitempty" json:"references,omitempty"`
  Replaces              string   `xml:"replaces,omitempty" json:"replaces,omitempty"`
  Requires              string   `xml:"requires,omitempty" json:"requires,omitempty"`
  RightsHolder          string   `xml:"rightsHolder,omitempty" json:"rightsHolder,omitempty"`
  Spatial               string   `xml:"spatial,omitempty" json:"spatial,omitempty"`
  TableOfContents       string   `xml:"tableOfContents,omitempty" json:"tableOfContents,omitempty"`
  Temporal              string   `xml:"temporal,omitempty" json:"temporal,omitempty"`
  Valid                 string   `xml:"valid,omitempty" json:"valid,omitempty"`
  Events                []Events `json:"events"`
  Agents                []Agents `json:"agents"`
}

// New: Premis events
type Events struct {
  Uuid       string `json:"uuid"`
  Type       string `json:"type"`
  DateTime   string `json:"datetime"`
  Outcome    string `json:"outcome"`
  Detail     string `json:"detail"`
  DetailNote string `json:"detail_note"`
}

// New: Premis agents
type Agents struct {
  IdentifierType   string `json:"identifier_type"`
  IdentifierValue  string `json:"identifier_value"`
  Name             string `json:"name"`
  Type             string `json:"type"`
}


// ********* XML Structs *********
type Mets struct {
  XMLName        xml.Name         `xml:"mets"`
  Header struct {
    CreateDate string `xml:"CREATEDATE,attr"`
    ModifyDate string `xml:"LASTMODDATE,attr"`
  } `xml:"metsHdr"`
  // Header         MetsHeader       `xml:"metsHdr"`
  DescriptiveSec []DescriptiveSec `xml:"dmdSec"`
  AdminSec       []AdminSec       `xml:"amdSec"`
  FileSec        FileSec          `xml:"fileSec"`
  StructMap      []StructMap      `xml:"structMap"`
}

// dmdSec
type DescriptiveSec struct {
  XMLName    xml.Name     `xml:"dmdSec"`
  ID         string       `xml:"ID,attr"`
  Dmd        Dmd          `xml:"mdWrap"`
  DigiProvMD []DigiProvMD `xml:"digiprovMD"`
}

type Dmd struct {
  XMLName      xml.Name      `xml:"mdWrap"`
  Mdtype       string        `xml:"MDTYPE,attr"`
  PremisObject PremisObject  `xml:"xmlData>object"`
  DublinCoreMD descriptiveMD `xml:"xmlData>dublincore"`
}

// amdSec
type AdminSec struct {
  XMLName     xml.Name     `xml:"amdSec"`
  ID          string       `xml:"ID,attr"`
  TechnicalMD TechnicalMD  `xml:"techMD"`
  DigiProvMD  []DigiProvMD `xml:"digiprovMD"`
  RightsMD    []RightsMD   `xml:"rightsMD"`
  SourceMD    SourceMD     `xml:"sourceMD"`
}

// amdSec > SourceMD
type SourceMD struct {
  XMLName                   xml.Name `xml:"sourceMD"`
  Payload                   string   `xml:"mdWrap>xmlData>transfer_metadata>Payload-Oxum"`
  BagCount                  string   `xml:"mdWrap>xmlData>transfer_metadata>Bag-Count"`
  ContactName               string   `xml:"mdWrap>xmlData>transfer_metadata>Contact-Name"`
  ContactEmail              string   `xml:"mdWrap>xmlData>transfer_metadata>Contact-Email"`
  BagSize                   string   `xml:"mdWrap>xmlData>transfer_metadata>Bag-Size"`
  BaggingDate               string   `xml:"mdWrap>xmlData>transfer_metadata>Bagging-Date"`
  SourceOrganization        string   `xml:"mdWrap>xmlData>transfer_metadata>Source-Organization"`
  ExternalDescription       string   `xml:"mdWrap>xmlData>transfer_metadata>External-Description"`
  ExternalIdentifier        string   `xml:"mdWrap>xmlData>transfer_metadata>External-Identifier"`
  BagGroupIdentifier        string   `xml:"mdWrap>xmlData>transfer_metadata>Bag-Group-Identifier"`
  InternalSenderIdentifier  string   `xml:"mdWrap>xmlData>transfer_metadata>Internal-Sender-Identifier"`
  InternalSenderDescription string   `xml:"mdWrap>xmlData>transfer_metadata>Internal-Sender-Description"`
}

// amdSec > TechnicalMD
type TechnicalMD struct {
  XMLName      xml.Name     `xml:"techMD"`
  ID           string       `xml:"ID,attr"`
  PremisObject PremisObject `xml:"mdWrap>xmlData>object"`
}

// mets > []structmap
type StructMap struct {
  XMLName xml.Name `xml:"structMap"`
  ID      string   `xml:"ID,attr"`
  Label   string   `xml:"LABEL,attr"`
  Type    string   `xml:"TYPE,attr"`
  Parent  Div      `xml:"div"`
}

// structmap > div (File item Div)
type Div struct {
  XMLName xml.Name     `xml:"div"`
  Label    string      `xml:"LABEL,attr"`
  Type     string      `xml:"TYPE,attr"`
  Dmdid    string      `xml:"DMDID,attr"`
  Admid    string      `xml:"ADMID,attr"`
  File     FilePointer `xml:"fptr"`
  Children []Div       `xml:"div"`
}

// structmap > div (item) > fileptr
type FilePointer struct {
  XMLName xml.Name `xml:"fptr"`
  Fileid  string   `xml:"FILEID,attr"`
}

// amdSec > digiprov
type DigiProvMD struct {
  XMLName xml.Name `xml:"digiprovMD"`
  ID      string   `xml:"ID,attr"`
  Mdtype  string   `xml:"MDTYPE,attr"`
  Premis  Premis   `xml:"mdWrap"`
}

// amdSec > digiprov > PremisAgent | PremisEvent
type Premis struct {
  Mdtype      string      `xml:"MDTYPE,attr"`
  PremisAgent PremisAgent `xml:"xmlData>agent"`
  PremisEvent PremisEvent `xml:"xmlData>event"`
}

// amdSec > digiprov > PremisEvent
type PremisEvent struct {
  XMLName              xml.Name `xml:"event"`
  EventIdentifierValue string   `xml:"eventIdentifier>eventIdentifierValue"`
  EventType            string   `xml:"eventType"`
  EventDate            string   `xml:"eventDateTime"`
  // EventDetail          string   `xml:"eventDetail"`  // FRDR METS different structure
  EventDetail          string   `xml:"eventDetailInformation>eventDetail"`
  EventOutcome         string   `xml:"eventOutcomeInformation>eventOutcome"`
  EventOutcomeNote     string   `xml:"eventOutcomeInformation>eventOutcomeDetail>eventOutcomeDetailNote"`
}

// amdSec > digiprov > PremisEvent
type PremisAgent struct {
  XMLName              xml.Name `xml:"agent"`
  AgentIdentifierType  string   `xml:"agentIdentifier>agentIdentifierType"`
  AgentIdentifierValue string   `xml:"agentIdentifier>agentIdentifierValue"`
  AgentName            string   `xml:"agentName"`
  AgentType            string   `xml:"agentType"`
}

// amdSec > rightsmd
type RightsMD struct {
  XMLName xml.Name `xml:"rightsMD"`
  // something here
}

// amdSec > techMd > PremisObject
type PremisObject struct {
  ObjectName         string   `xml:"originalName"`
  Uuid               []string `xml:"objectIdentifier>objectIdentifierValue"`
  Hashtype           string   `xml:"objectCharacteristics>fixity>messageDigestAlgorithm"`
  Hashvalue          string   `xml:"objectCharacteristics>fixity>messageDigest"`
  Bytes              string   `xml:"objectCharacteristics>size"`
  Format             string   `xml:"objectCharacteristics>format>formatDesignation>formatName"`
  Version            string   `xml:"objectCharacteristics>format>formatDesignation>formatVersion"`
  FormatRegistryName string   `xml:"objectCharacteristics>format>formatRegistry>formatRegistryName"`
  FormatRegistryKey  string   `xml:"objectCharacteristics>format>formatRegistry>formatRegistryKey"`
  ModifiedDate       string   `xml:"objectCharacteristics>creatingApplication>dateCreatedByApplication"`
  Fits               Fits     `xml:"objectCharacteristics>objectCharacteristicsExtension>fits"`
}

// amdSec > techMd > PremisObject > Fits
type Fits struct {
  XMLName          xml.Name `xml:"fits"`
  ModifiedUnixtime string   `xml:"fits>fileinfo>fslastmodified"`
  Md5              string   `xml:"fileinfo>md5checksum"`
  Filepath         string   `xml:"fileinfo>filepath"`
  Filename         string   `xml:"fileinfo>filename"`
  Identity         Identity `xml:"identification>identity"`
}

// amdSec > techMd > PremisObject > Fits > Identity
type Identity struct {
  XMLName     xml.Name `xml:"identity"`
  Format      string   `xml:"format,attr"`
  Mimetype    string   `xml:"mimetype,attr"`
  Toolname    string   `xml:"toolname,attr"`
  Toolversion string   `xml:"toolversion,attr"`
}

// mets > filesec
type FileSec struct {
  XMLName xml.Name  `xml:"fileSec"`
  FileGrp []FileGrp `xml:"fileGrp"`
}

// mets > filesec > filegrp
type FileGrp struct {
  XMLName  xml.Name `xml:"fileGrp"`
  FileType string   `xml:"USE,attr"`
  Files    []File   `xml:"file"`
}

type File struct {
  XMLName xml.Name `xml:"file"`
  Admid   string   `xml:"ADMID,attr"`
  ID      string   `xml:"ID,attr"`
  FileLocation struct {
    Location string `xml:"href,attr"`
  } `xml:"FLocat"`
}

// Associate object to corresponding mets metadata
type FileMapped struct {
  Admid string
  Dmdid []string
  Name string
}

func main() {
  metsFilePathUserInput := flag.String("mets", "", "Provide a mets filepath")
  outputDirPathUserInput := flag.String("out", "", "Provide an output directory")

  flag.Parse()

  filePath := *metsFilePathUserInput
  dirPath := *outputDirPathUserInput

  if (filePath == "") {
    fmt.Println()
    fmt.Println("PLEASE ENTER A METS FILEPATH")
    fmt.Scanln(&filePath)

    if string(filePath) == "" {
      log.Fatal("ERROR : MUST ENTER A METS FILEPATH")
    }
  }

  if (dirPath == "") {
    fmt.Println()
    fmt.Println("PLEASE ENTER OUTPUT DIRECTORY PATH")
    fmt.Scanln(&dirPath)

    if string(filePath) == "" {
      log.Fatal("ERROR : MUST ENTER OUTPUT DIRECTORY PATH")
    }
  }

  file, err := os.Open(filePath);
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()
  data, err := ioutil.ReadAll(file)
  if err != nil {
      log.Fatal(err)
  }

  val := Mets{}
  // deserialization, transform XML to go object that can be processed
  // From XML string transform to Go struct data structure
  err = xml.Unmarshal(data, &val)
  if err != nil {
      log.Fatal(err)
  }

  buildMetadataMets(val, dirPath)

  fmt.Println("Success!")
}

// Output JSON file with METS metadata in Canopus schema
func buildMetadataMets(mets Mets, target string) (string, string) {
	manifestObject := ObjectMetsManifest{}
  packageName := getParentPackage(mets.StructMap)

  if mets.DescriptiveSec == nil {
    log.Fatal("Descriptive metadata (dmdSec) missing.")
  }

  file_count, files, transferLevelDc := extractMetadataMetsFile(mets)
  manifestObject.Title = transferLevelDc.Title
  if manifestObject.Title == "" {
    manifestObject.Title = packageName
  }
  manifestObject.CollectionCall = transferLevelDc.Identifier
  manifestObject.Description = transferLevelDc.Description
  manifestObject.BaggingDate = mets.Header.CreateDate
  manifestObject.FileCount = file_count

  manifest := manifestMetsJSON{}
  e := getSiegfriedMetadata(mets.AdminSec)
  if e != nil {
    sieg := getSiegfriedVersion(e.Detail)
    manifest.Siegfried = sieg["version"]
    manifest.Scandate = e.DateTime
  }
  manifest.Files = files
  identifier := Identifiers{}
  var identifiers []Identifiers
  identifiers = append(identifiers, identifier)
  manifest.Identifiers = identifiers
  manifestObject.Manifest = manifest
  manifestObject.StorageLocation = packageName
  manifestObject.SchemaVersion = "0.2.0"

	// target += "/" + manifestObject.Title + "_" + "metadata.json"
  target += "/" + packageName + "_" + "metadata.json"

	//Write struct to file
	writeNewStructToFile(target, manifestObject)
	return target, ""
}

// Return file count, list of files, objects directory (transfer level) metadata
func extractMetadataMetsFile(mets Mets) (int64, []FilesMets, descriptiveMD){
  var total_size int64
  var file_count_all int64
  total_size = 0
  file_count_all = 0
  var files []FilesMets
  var transferLevelDc descriptiveMD

  // get descriptive metadata
  dublincore := getDublinCore(mets)
  structmap := getFileIdDdmdIdStructMap(mets.StructMap)

  for _, id := range structmap["objects"] {
    _, ok := dublincore[id]
    if ok {
      transferLevelDc = dublincore[id]
    }
  }

  // map of files with corresponding admd, dmd,
  filemap := getAmdIdByFileIdFileSec(mets.FileSec, structmap)

  // one adminsec for each file
  for _, a := range mets.AdminSec {
    file := FilesMets{}
    if (a.TechnicalMD.ID != "") {
      t := a.TechnicalMD
      file.Md5 =  t.PremisObject.Fits.Md5
      if t.PremisObject.Hashtype == "sha256" {
        file.Sha256 =  t.PremisObject.Hashvalue
      }
      if (t.PremisObject.Bytes == "") {
        log.Fatal("Empty bytes") // TODO
      }
      byte, err := strconv.Atoi(t.PremisObject.Bytes)
      if err != nil {
        log.Fatal(err)
      }
      total_size += int64(byte)
      file.FileSize = int64(byte)
      file.Modified = t.PremisObject.ModifiedDate // TODO

      // file identification match PRONOM
      match := Matches{}
      match.Format =  t.PremisObject.Format
      match.Version = t.PremisObject.Version
      match.Ns = t.PremisObject.FormatRegistryName
      match.ID = t.PremisObject.FormatRegistryKey
      match.Mime = t.PremisObject.Fits.Identity.Mimetype
      file.Matches = append(file.Matches, match)

      // PREMIS:EVENT and AGENTS
      events, agents := getPremisEvents(a)

      // DublinCore metadata 
      descriptivemd := descriptiveMD{}
      descriptivemd.Events = events
      descriptivemd.Agents = agents

      for _, value := range filemap {
        if value.Admid == a.ID {
          file.FileName = value.Name
          for _, id := range value.Dmdid {
            _, ok := dublincore[id] // [dmdSec_2, dmdSec_3]
            if ok {
              descriptivemd = dublincore[id]
              descriptivemd.Language = strings.Join(descriptivemd.LanguageArr, ",")
              descriptivemd.Subject = strings.Join(descriptivemd.SubjectArr, ",")
            }
          }
        }
      }
      file.DescriptiveMD = descriptivemd

      files = append(files, file)
      file_count_all++
    }
  }
  return file_count_all, files, transferLevelDc
}

// return map of dublincore metadata identified by dmd ID
func getDublinCore(mets Mets) (map[string]descriptiveMD){
  dublincore := make(map[string]descriptiveMD)
  for _, desc := range mets.DescriptiveSec {
    if desc.Dmd.Mdtype == "DC" {
      dublincore[desc.ID] = desc.Dmd.DublinCoreMD
    }
  }
  return dublincore
}

// get PREMIS:EVENTS and PREMIS:AGENT for an object identified by AdminSec
func getPremisEvents(a AdminSec) ([]Events, []Agents){
  var events []Events
  var agents []Agents
  for _, digiprov := range a.DigiProvMD {
    if digiprov.Premis.Mdtype == "PREMIS:EVENT" {
      event := Events{}
      event.Uuid = digiprov.Premis.PremisEvent.EventIdentifierValue
      event.Type = digiprov.Premis.PremisEvent.EventType
      event.DateTime = digiprov.Premis.PremisEvent.EventDate
      event.Detail = digiprov.Premis.PremisEvent.EventDetail
      event.Outcome = digiprov.Premis.PremisEvent.EventOutcome
      event.DetailNote = digiprov.Premis.PremisEvent.EventOutcomeNote
      events = append(events, event)
    }
    if digiprov.Premis.Mdtype == "PREMIS:AGENT" {
      agent := Agents{}
      agent.IdentifierType = digiprov.Premis.PremisAgent.AgentIdentifierType
      agent.IdentifierValue = digiprov.Premis.PremisAgent.AgentIdentifierValue
      agent.Name = digiprov.Premis.PremisAgent.AgentName
      agent.Type = digiprov.Premis.PremisAgent.AgentType
      agents = append(agents, agent)
    }
  }
  return events, agents
}

// get file ID of object from structmap
func getFileIdDdmdIdStructMap(structmap []StructMap) (map[string][]string) {
  sm := make(map[string][]string)
  for _, s1 := range structmap {
    if s1.Label == "Archivematica default" {
      unpackDiv(s1.Parent, sm)
    }
  }
  return sm
}

// recursively iterate over Directory objects to get file ID of non-Directory objects
func unpackDiv(div Div, sm map[string][]string) (map[string][]string){
  dmdIds := strings.Split(div.Dmdid, " ")
  if div.Type == "Item" {
    sm[div.File.Fileid] = dmdIds // DMDID="dmdSec_3 dmdSec_4"
  } else if (div.Type == "Directory") {
    if (div.Label == "objects") {
      sm[div.Label] = dmdIds
    }
    for _, c := range div.Children {
      unpackDiv(c,sm)
    }
  }
  return sm
}

// get administrative ID of file in File Section
func getAmdIdByFileIdFileSec(filesec FileSec, structmap map[string][]string) (map[string]FileMapped){
  filemap := make(map[string]FileMapped)
  for _, grp := range filesec.FileGrp {
    for _, file := range grp.Files {
      _, ok := structmap[file.ID]
      if ok {
        filemapped := FileMapped{}
        filemapped.Admid = file.Admid
        filemapped.Dmdid = structmap[file.ID]
        filemapped.Name = file.FileLocation.Location
        filemap[file.ID] = filemapped
        }
      }
    }
    return filemap
  }

// return Siefried information from the first metadata it finds
func getSiegfriedMetadata(adminsec []AdminSec) (*Events){
  for _, a := range adminsec {
    events, _ := getPremisEvents(a)
    for _, value := range events {
      if strings.Contains(value.Detail, "Siegfried") {
         return &value
      }
    }
  }
  return nil
}

// parse Siegfried metadata into map
func getSiegfriedVersion(sieg string) (map[string]string) {
  sieg1 := strings.Split(sieg, "; ") // [program="Siegfried", version="1.8.0"]
  siegProgram := strings.Split(sieg1[0], "=")
  siegVersion := strings.Split(sieg1[1], "=")
  siegMap := map[string]string {
    "program":strings.Trim(siegProgram[1], "\""),
    "version":strings.Trim(siegVersion[1], "\""),
  }
  return siegMap
}

// get parent package name
func getParentPackage(structMap []StructMap) string {
  packageName := ""
  for _, sm := range structMap {
    if sm.Label == "Archivematica default" {
      if sm.Parent.Type == "Directory" {
        packageName = sm.Parent.Label
      }
    }
  }
  return packageName
}

//taken from upload.go
func writeNewStructToFile(file string, m ObjectMetsManifest) {
	output, err := json.MarshalIndent(&m, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(file, output, 0750)
	if err != nil {
		log.Fatal(err)
	}
}
