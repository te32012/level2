package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.

*/

type ConfigWget struct {
	PathForSafeSite  string
	DepthOfRecursion int
	URI              string
}

func (configwget *ConfigWget) ConfigWget() {

	flag.IntVar(&configwget.DepthOfRecursion, "d", 1, "глубина рекурсивного скачивания файлов")

	flag.Parse()

	agrs := flag.Args()
	fmt.Println(agrs)
	fmt.Println(configwget.DepthOfRecursion)
	if len(agrs) == 2 {
		configwget.PathForSafeSite = agrs[0]
		configwget.URI = agrs[1]
		//isvalide, err := regexp.MatchString(`/^(http|https|ftp):\/\/(([A-Z0-9][A-Z0-9_-]*)(\.[A-Z0-9][A-Z0-9_-]*)+)/i`, configwget.URI)
		//if err != nil {
		//	log.Fatal()
		//}
		//if !isvalide {
		//	log.Fatal("переданный URI адрес не валидный")
		//}
	} else {
		log.Fatal("не было передано досаточное количество аргументов")
	}
}

func (configwget *ConfigWget) Wget(link string, depth int) {
	if link == "" {
		return
	}
	u0, err := url.ParseRequestURI(link)
	if err != nil {
		log.Fatal(err)
	}
	var tmplink string = u0.Scheme + `://` + u0.User.String() + u0.Hostname() + u0.Port() + u0.Path
	fmt.Println("wget: download site from " + tmplink)
	response, err := http.Get(tmplink)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	buffer, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("wget: response " + " status " + response.Status)
	dir := configwget.GetPathDir(u0.Path)
	configwget.CreateDirectory(configwget.PathForSafeSite + "/" + u0.Host + dir)
	filename := configwget.GetPathName(u0.Path)
	site, err := os.Create(configwget.PathForSafeSite + "/" + u0.Host + dir + "/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer site.Close()
	fmt.Println("wget:file " + u0.Host + dir + "/" + filename + " create")
	_, err = site.Write(buffer)
	if err != nil {
		log.Fatal(err)
	}
	dependsLinks := configwget.GetLinksFromFile(buffer)
	var absolutedependslinks []string
	// fmt.Println(dependsLinks)
	for _, allLink := range dependsLinks {
		u, err := url.Parse(allLink)
		if err != nil {
			log.Fatal(err)
		}
		if !u.IsAbs() && u.Host == "" {
			u1, err := u0.Parse(allLink)
			if err != nil {
				log.Fatal(err)
			}
			absolutedependslinks = append(absolutedependslinks, u1.Scheme+`://`+u1.User.String()+u1.Hostname()+u1.Port()+u1.Path)
			continue
		}
		if u.IsAbs() && u.Host == u0.Host {
			absolutedependslinks = append(absolutedependslinks, allLink)
		}
	}
	if depth-1 >= 0 {
		for _, dependlink := range absolutedependslinks {
			configwget.Wget(dependlink, depth-1)
		}
	}
}

func (configwget *ConfigWget) GetPathName(path string) string {
	tmp := strings.Split(path, "/")
	last := tmp[len(tmp)-1]
	tmp = strings.Split(last, "#")
	name := tmp[0]
	if name == "" {
		return "index.html"
	}
	return name
}

func (configwget *ConfigWget) GetPathDir(uri string) string {
	tmp := strings.Split(uri, "/")
	return strings.Join(tmp[:len(tmp)-1], "/")
}

func (configwget *ConfigWget) CreateDirectory(dir string) {
	_, err := os.Stat(dir)
	if err == nil {
		return
	}
	isexist := os.IsExist(err)
	if !isexist {
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (configwget *ConfigWget) GetLinksFromFile(site []byte) []string {
	var links map[string]bool = make(map[string]bool)
	absolutelinkreg := regexp.MustCompile(`(http|https|ftp):\/\/([\w\-_]+(?:(?:\.[\w\-_]+)+))([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?`)
	herflinkreg := regexp.MustCompile(`href="\S+"`)
	scrlinkreg := regexp.MustCompile(`src="\S+"`)
	absolutelink := absolutelinkreg.FindAll(site, -1)
	for _, x := range absolutelink {
		tmp := strings.Split(string(x), "#")[0]
		links[tmp] = true
	}
	herflink := herflinkreg.FindAll(site, -1)
	for _, x := range herflink {
		tmp := strings.Split(string(x[6:len(x)-1]), "#")[0]
		links[tmp] = true
	}
	scrlink := scrlinkreg.FindAll(site, -1)
	for _, x := range scrlink {
		tmp := strings.Split(string(x[5:len(x)-1]), "#")[0]
		links[tmp] = true
	}
	var ans []string
	for keys := range links {
		ans = append(ans, keys)
	}
	return ans
}
func main() {
	var configwget ConfigWget = ConfigWget{}
	configwget.ConfigWget()
	configwget.Wget(configwget.URI, configwget.DepthOfRecursion)
}
