package crimeModels

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"time"

	//	"io/ioutil"
	"bufio"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func con() *gorm.DB {
	dsn := "host=192.168.1.94 user=dusty password=Qapla1999 dbname=dusty port=5432 sslmode=disable"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db
}

func init() {
	fmt.Println("vim-go")
	DB = con()
}

type Detroits struct {
	DetroitNodes []Detroit `json:"detroit_nodes"`
}
type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}
type Detroit struct {
	Rownum                string `json:"ROWNUM"`
	Caseid                string `json:"CASEID"`
	Crimeid               string `json:"CRIMEID"`
	Crno                  string `json:"CRNO"`
	Category              string `json:"CATEGORY"`
	Offensedescription    string `json:"OFFENSEDESCRIPTION" gorm:"foreignKey:Name`
	Stateoffensefileclass string `json:"STATEOFFENSEFILECLASS"`
	Incidentdate          string `json:"INCIDENTDATE"`
	Hour                  string `json:"HOUR"`
	Sca                   string `json:"SCA"`
	Precinct              string `json:"PRECINCT"`
	Council               string `json:"COUNCIL"`
	Neighborhood          string `json:"NEIGHBORHOOD"`
	Censustract           string `json:"CENSUSTRACT"`
	Location              string `json:"LOCATION"`
}

type Offense struct {
	ID   int
	Name string `gorm:"references:Offensedescription"`
}

func init() {
	fmt.Println("vim-go")
}

func logn(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}

func humanateBytes(s uint64, base float64, sizes []string) string {
	if s < 10 {
		return fmt.Sprintf("%dB", s)
	}
	e := math.Floor(logn(float64(s), base))
	suffix := sizes[int(e)]
	val := float64(s) / math.Pow(base, math.Floor(e))
	f := "%.0f"
	if val < 10 {
		f = "%.1f"
	}

	return fmt.Sprintf(f+"%s", val, suffix)
}

// FileSize calculates the file size and generate user-friendly string.
func FileSize(s int64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	return humanateBytes(uint64(s), 1024, sizes)
}

func GetDetroit() (Detroit, []*Detroit) {
	//	d := Detroit{}
	//	fmt.Println(d)
	start := time.Now()
	fileName := "models/detroit.json"
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error to read [file=%v]: %v", fileName, err.Error())
	}

	fi, err := f.Stat()
	if err != nil {
		log.Fatalf("Could not obtain stat, handle error: %v", err.Error())
	}

	r := bufio.NewReader(f)
	d := json.NewDecoder(r)

	i := 0

	d.Token()
	data := []*Detroit{}
	for d.More() {
		elm := &Detroit{}
		d.Decode(elm)
		//fmt.Println(d.Token)
		data = append(data, elm)
		//fmt.Printf("%v \n", elm)
		i++
	}
	d.Token()
	elapsed := time.Since(start)

	fmt.Printf("Total of [%v] object created.\n", i)
	fmt.Printf("The [%s] is %s long\n", fileName, FileSize(fi.Size()))
	fmt.Printf("To parse the file took [%v]\n", elapsed)

	//	file, _ := ioutil.ReadFile("models/detroit.json")
	//	file, err := file.Open(path)
	/* handle error */
	//scanner = bufio.NewScanner(file)
	//for scanner.Scan() {
	//    processLine(scanner.Text())
	//}
	//fmt.Println(file)

	//_ = json.Unmarshal([]byte(d), &data)
	//	datar := d.Decode([]Detroit{})
	//fmt.Println(data)
	//fmt.Println(j)
	return Detroit{}, data
}

func CreateDetroit(detroit *Detroit) (err error) {
	if err = DB.Create(detroit).Error; err != nil {
		return err
	}
	return nil
}
