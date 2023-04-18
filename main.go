package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const port = ":4000"

func css(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/styles.css")
}

func js(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/main.js")
}

func jsone(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/data.json")
}

func search(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/search.html")
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Notfound(w, r)
		return
	}

	//"https://groupietrackers.herokuapp.com/api/locations"

	renderTemplate(w, "home")
	case1 := strconv.Itoa(rand.Intn(52-34) + 34)
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + case1)
	if err != nil {
		fmt.Println("No response from request")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = err
	}
	defer resp.Body.Close()

	var result Api

	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	loc1, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + case1)
	if err != nil {
		fmt.Println("No response from request")
	}
	bodyLoc, err := ioutil.ReadAll(loc1.Body)
	if err != nil {
		err = err
	}
	defer loc1.Body.Close()

	var location Rel

	if err := json.Unmarshal(bodyLoc, &location); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	case2 := strconv.Itoa(rand.Intn(34-17) + 17)
	resp2, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + case2)
	if err != nil {
		fmt.Println("No response from request")
	}
	body2, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		err = err
	}
	defer resp2.Body.Close()

	var result2 Api

	if err := json.Unmarshal(body2, &result2); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	loc2, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + case2)
	if err != nil {
		fmt.Println("No response from request")
	}
	bodyLoc2, err := ioutil.ReadAll(loc2.Body)
	if err != nil {
		err = err
	}

	var location2 Rel

	defer loc2.Body.Close()
	if err := json.Unmarshal(bodyLoc2, &location2); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	case3 := strconv.Itoa(rand.Intn(17-1) + 1)
	resp3, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + case3)
	if err != nil {
		fmt.Println("No response from request")
	}
	body3, err := ioutil.ReadAll(resp3.Body)
	if err != nil {
		err = err
	}
	defer resp.Body.Close()

	var result3 Api

	if err := json.Unmarshal(body3, &result3); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	loc3, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + case3)
	if err != nil {
		fmt.Println("No response from request")
	}
	bodyLoc3, err := ioutil.ReadAll(loc3.Body)
	if err != nil {
		err = err
	}

	var location3 Rel

	defer loc3.Body.Close()
	if err := json.Unmarshal(bodyLoc3, &location3); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	// https://groupietrackers.herokuapp.com/api/relation
	dat1, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + case1)
	if err != nil {
		fmt.Println("No response from request")
	}
	DatLoc1, err := ioutil.ReadAll(dat1.Body)
	if err != nil {
		err = err
	}

	var final map[string]interface{}
	defer dat1.Body.Close()
	json.Unmarshal([]byte(DatLoc1), &final)

	dat2, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + case2)
	if err != nil {
		fmt.Println("No response from request")
	}
	DatLoc2, err := ioutil.ReadAll(dat2.Body)
	if err != nil {
		err = err
	}

	var final2 map[string]interface{}
	defer dat2.Body.Close()
	json.Unmarshal([]byte(DatLoc2), &final2)

	dat3, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + case3)
	if err != nil {
		fmt.Println("No response from request")
	}
	DatLoc3, err := ioutil.ReadAll(dat3.Body)
	if err != nil {
		err = err
	}

	var final3 map[string]interface{}
	defer dat3.Body.Close()
	json.Unmarshal([]byte(DatLoc3), &final3)

	fmt.Fprintf(w, "<div class=\"page\">")
	fmt.Fprintf(w, "<div class=\"card\">")
	fmt.Fprintf(w, "<div class=\"cardtop\">")
	fmt.Fprint(w, "<img class=\"image\" src="+result.Image+">")
	fmt.Fprintf(w, "</div>")
	fmt.Fprintf(w, "<div class=\"cardbottom\">")
	fmt.Fprint(w, "<h2 class=\"name\">"+result.Name+" • "+strconv.Itoa(result.CreationDate)+"</h2>")
	fmt.Fprint(w, "<p class=\"firstAlbum\">First album released in : "+result.FirstAlbum+"</p>")
	fmt.Fprintf(w, "<h2>Members :</h2>")
	fmt.Fprintf(w, "<ul>")
	for _, rec := range result.Members {
		fmt.Fprint(w, "<li class=\"members\">"+rec+"</li>")
	}
	fmt.Fprintf(w, "</ul>")
	fmt.Fprintf(w, "<h2>Locations :</h2>")
	fmt.Fprintf(w, "<ul>")
	pute := interface_to_string((final["datesLocations"]))
	if strings.Contains(pute, "]") {
		temp := strings.ReplaceAll(pute, "]", "")
		temp1 := strings.ReplaceAll(temp, "[", "")
		temp2 := strings.ReplaceAll(temp1, "map", "")
		temp3 := strings.ReplaceAll(temp2, string(rune(32)), "<br>")
		temp4 := strings.ReplaceAll(temp3, ":", ":<br>")
		temp5 := strings.ReplaceAll(temp4, "_", " ")
		fmt.Fprint(w, temp5)
	}
	fmt.Fprintf(w, "</ul>")
	fmt.Fprintf(w, "</ul>")
	fmt.Fprintf(w, "</div>")
	fmt.Fprintf(w, "</div>")
	fmt.Fprintf(w, "<div class=\"card\">")
	fmt.Fprintf(w, "<div class=\"cardtop\">")
	fmt.Fprint(w, "<img class=\"image\" src="+result2.Image+">")
	fmt.Fprintf(w, "</div>")
	fmt.Fprintf(w, "<div class=\"cardbottom\">")
	fmt.Fprint(w, "<h2 class=\"name\">"+result2.Name+" • "+strconv.Itoa(result2.CreationDate)+"</h2>")
	fmt.Fprint(w, "<p class=\"firstAlbum\">First album released in : "+result2.FirstAlbum+"</p>")
	fmt.Fprintf(w, "<h2>Members :</h2>")
	fmt.Fprintf(w, "<ul>")
	for _, rec := range result2.Members {
		fmt.Fprint(w, "<li class=\"members\">"+rec+"</li>")
	}
	fmt.Fprintf(w, "</ul>")
	fmt.Fprintf(w, "<h2>Locations :</h2>")
	fmt.Fprintf(w, "<ul>")
	pute2 := interface_to_string((final2["datesLocations"]))
	if strings.Contains(pute, "]") {
		temp := strings.ReplaceAll(pute2, "]", "")
		temp1 := strings.ReplaceAll(temp, "[", "")
		temp2 := strings.ReplaceAll(temp1, "map", "")
		temp3 := strings.ReplaceAll(temp2, string(rune(32)), "<br>")
		temp4 := strings.ReplaceAll(temp3, ":", ":<br>")
		temp5 := strings.ReplaceAll(temp4, "_", " ")
		fmt.Fprint(w, temp5)
	}
	fmt.Fprintf(w, "</ul>")
	fmt.Fprintf(w, "</div>")
	fmt.Fprintf(w, "</div>")
	fmt.Fprintf(w, "<div class=\"card\">")
	fmt.Fprintf(w, "<div class=\"cardtop\">")
	fmt.Fprint(w, "<img class=\"image\" src="+result3.Image+">")
	fmt.Fprintf(w, "</div>")
	fmt.Fprintf(w, "<div class=\"cardbottom\">")
	fmt.Fprint(w, "<h2 class=\"name\">"+result3.Name+" • "+strconv.Itoa(result3.CreationDate)+"</h2>")
	fmt.Fprint(w, "<p class=\"firstAlbum\">First album released in : "+result3.FirstAlbum+"</p>")
	fmt.Fprintf(w, "<h2>Members :</h2>")
	fmt.Fprintf(w, "<ul>")
	for _, rec := range result3.Members {
		fmt.Fprint(w, "<li class=\"members\">"+rec+"</li>")
	}
	fmt.Fprintf(w, "</ul>")
	fmt.Fprintf(w, "<h2>Locations :</h2>")
	fmt.Fprintf(w, "<ul>")
	pute3 := interface_to_string((final3["datesLocations"]))
	if strings.Contains(pute, "]") {
		temp := strings.ReplaceAll(pute3, "]", "")
		temp1 := strings.ReplaceAll(temp, "[", "")
		temp2 := strings.ReplaceAll(temp1, "map", "")
		temp3 := strings.ReplaceAll(temp2, string(rune(32)), "<br>")
		temp4 := strings.ReplaceAll(temp3, ":", ":<br>")
		temp5 := strings.ReplaceAll(temp4, "_", " ")
		fmt.Fprint(w, temp5)
	}
	fmt.Fprintf(w, "</ul>")
	fmt.Fprintf(w, "</div>")
	fmt.Fprintf(w, "</div>")
	fmt.Fprintf(w, "</div>")
	fmt.Fprintf(w, `    <div class="footer-basic">
	<footer>
		<div class="social"><a target="_blank" href="https://www.instagram.com/zone01_rouen_normandie"><i class="icon ion-social-instagram"></i></a><a target="_blank" href="https://fr.linkedin.com/company/zone01-rouen-normandie"><i class="icon ion-social-linkedin"></i></a><a target="_blank" href="https://twitter.com/zone01rouen?lang=fr"><i class="icon ion-social-twitter"></i></a><a target="_blank" href="https://www.facebook.com/Zone01Rouen"><i class="icon ion-social-facebook"></i></a></div>
	  
		<p class="copyright">Daryl, Evan, Quentin © 2022</p>
	</footer>
</div>`)
}

func main() {
	json_modifier()
	http.HandleFunc("/", Home)
	http.HandleFunc("/styles.css", css)
	http.HandleFunc("/main.js", js)
	http.HandleFunc("/data.json", jsone)
	http.HandleFunc("/search.html", search)
	fmt.Println("SERVER AWAITS: http://localhost:4000/")
	http.ListenAndServe(port, nil)
}

// PrettyPrint to print struct in a readable way

func renderTemplate(w http.ResponseWriter, s string) {
	t, err := template.ParseFiles("./templates/" + s + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func Notfound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<body style='background-color:black; color:white; text-align:center'>Error 404, this page doesn't exist</body>")
}

type Api struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Rel struct {
	Location []string `json:"Locations"`
	// Id       int      `json: id`
}

// Func pour convertir une interface en string
func interface_to_string(inter interface{}) string {
	return fmt.Sprintf("%v", inter)
}

func json_modifier() {
	file, err := os.ReadFile("./templates/data.json")
	err_nil(err)
	/* file_loc, err := os.ReadFile("locations.json") */
	re := regexp.MustCompile(`("https:\/\/groupietrackers\.herokuapp\.com\/api\/locations\/)(\d*)"`)
	result := re.FindAllSubmatch(file, -1)
	// var replace string
	var End_File string
	End_File = string(file)
	for i := range result {
		Json_Formated := []string{}
		Index := string(result[i][2])
		toReplace := string(result[i][0])
		Locations, err := ScrapeLocations(Atoi(Index))
		for u, Location := range Locations.Location {
			if len(Locations.Location)-1 != u {
				Json_Formated = append(Json_Formated, `"`+Location+`",`)
			} else {
				Json_Formated = append(Json_Formated, `"`+Location+`"`)
			}
		}
		err_nil(err)
		fmt.Println(interface_to_string(Json_Formated))
		// fmt.Println(toReplace)
		End_File = strings.ReplaceAll(End_File, toReplace, interface_to_string(Json_Formated))
		// fmt.Println(Json_Formated)
		fl, err := os.Create("./templates/data.json")
		fl.WriteString(End_File)
		fl.Close()
		err_nil(err)
		// fmt.Println(Index)
	}
}

func err_nil(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ScrapeLocations(id int) (m Rel, err error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + strconv.Itoa(id))
	err_nil(err)

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &m)
	return
}

func Atoi(s string) int {
	inits := s
	var result int
	if len(s) > 0 {
		if s[0] == '+' || s[0] == '-' {
			s = s[1:]
		}
	}
	for _, elem := range []byte(s) {
		elem -= '0'
		if elem > 9 {
			return 0
		}
		result = result*10 + int(elem)
	}
	if inits[0] == '-' {
		result = -result
	}

	return result
}
