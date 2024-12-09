package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

const (
    ipinfoAPIKey    = "VOTRE_CLE_IPINFO"
    shodanAPIKey    = "VOTRE_CLE_SHODAN"
    greynoiseAPIKey = "VOTRE_CLE_GREYNOISE"
    censysAPIKey    = "VOTRE_CLE_CENSYS"
    censysSecret    = "VOTRE_SECRET_CENSYS"
)


// Définition de la structure de la réponse de crt.sh
type CrtShResponse struct {
	NameValue string `json:"name_value"`
}

// Fonction pour sauvegarder les résultats dans un fichier texte
func saveToFile(filename, content string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Erreur lors de la création du fichier : %v\n", err)
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		log.Fatalf("Erreur lors de l'écriture dans le fichier : %v\n", err)
	}
	fmt.Printf("Résultats enregistrés dans le fichier : %s\n", filename)
}

// Fonction pour gérer l'enregistrement des résultats
func handleSave(result string) {
	fmt.Println("Résultats :")
	fmt.Println(result)
	fmt.Print("Voulez-vous enregistrer les résultats dans un fichier texte ? (y/n) : ")
	var saveChoice string
	fmt.Scanln(&saveChoice)
	if strings.ToLower(saveChoice) == "y" {
		fmt.Print("Entrez le nom du fichier (avec extension, ex : resultat.txt) : ")
		var filename string
		fmt.Scanln(&filename)
		saveToFile(filename, result)
	}
}

// Fonction pour interroger crt.sh
func queryCrtSh(domain string) string {
	url := fmt.Sprintf("https://crt.sh/?q=%s&output=json", domain)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Erreur lors de la requête à crt.sh : %v\n", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Erreur lors de la lecture de la réponse crt.sh : %v\n", err)
	}

	var crtShResponse []CrtShResponse
	err = json.Unmarshal(body, &crtShResponse)
	if err != nil {
		log.Fatalf("Erreur lors de l'analyse JSON de crt.sh : %v\n", err)
	}

	// Créer un résultat sous forme de chaîne de caractères
	var result string
	for _, entry := range crtShResponse {
		result += entry.NameValue + "\n"
	}

	return result // Retourner une seule chaîne de caractères
}

// Fonction pour interroger ipinfo
func queryIpinfo(target string) string {
	// Si l'entrée est un domaine, résolvez-le en IP
	ips, err := net.LookupHost(target)
	if err != nil {
		log.Fatalf("Erreur lors de la résolution DNS : %v\n", err)
	}

	// Prenez la première IP résolue
	ip := ips[0]
	fmt.Printf("Résolution DNS : %s → %s\n", target, ip)

	// Appelez l'API ipinfo avec l'IP
	url := fmt.Sprintf("https://ipinfo.io/%s?token=%s", ip, ipinfoAPIKey)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Erreur lors de la requête à ipinfo : %v\n", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

// Fonction pour interroger Shodan
func queryShodan(target string) string {
	url := fmt.Sprintf("https://api.shodan.io/shodan/host/%s?key=%s", target, shodanAPIKey)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Erreur lors de la requête à Shodan : %v\n", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

// Fonction pour interroger GreyNoise
func queryGreyNoise(target string) string {
	url := fmt.Sprintf("https://api.greynoise.io/v3/community/%s", target)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Erreur lors de la création de la requête GreyNoise : %v\n", err)
	}
	req.Header.Set("key", greynoiseAPIKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Erreur lors de la requête à GreyNoise : %v\n", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

// Fonction pour interroger Censys
func queryCensys(target string) string {
	url := fmt.Sprintf("https://search.censys.io/api/v1/hosts/%s", target)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Erreur lors de la création de la requête Censys : %v\n", err)
	}
	req.SetBasicAuth(censysAPIKey, censysSecret)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Erreur lors de la requête à Censys : %v\n", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Censys a retourné un code de statut : %d\n", resp.StatusCode)
		return "Erreur : Impossible d'obtenir les résultats de Censys"
	}

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

// Fonction principale
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Menu avec décoration
		fmt.Println("==================================================")
		fmt.Println("🌐 Bienvenue dans l'outil de reconnaissance réseau 🌐")
		fmt.Println("==================================================")
		fmt.Println("Options disponibles :")
		fmt.Println("1. Rechercher avec ipinfo")
		fmt.Println("2. Rechercher avec Shodan")
		fmt.Println("3. Rechercher avec GreyNoise")
		fmt.Println("4. Rechercher avec Censys")
		fmt.Println("5. Rechercher avec crt.sh")
		fmt.Println("6. Quitter")
		fmt.Println("==================================================")
		fmt.Print("\nChoisissez une option : ")

		scanner.Scan()
		choice := scanner.Text()

		if choice == "6" {
			fmt.Println("Merci d'avoir utilisé l'outil. Au revoir !")
			break
		}

		fmt.Print("Entrez l'IP ou le domaine à analyser : ")
		scanner.Scan()
		target := scanner.Text()

		// Créez un espace pour les résultats
		var result string

		// Option pour seulement crt.sh
		switch choice {
		case "1":
			result = queryIpinfo(target)
		case "2":
			result = queryShodan(target)
		case "3":
			result = queryGreyNoise(target)
		case "4":
			result = queryCensys(target)
		case "5":
			// Recherche uniquement via crt.sh
			result = queryCrtSh(target) // Cela retourne maintenant une chaîne
		default:
			fmt.Println("Option invalide. Veuillez réessayer.")
			continue
		}

		// Affichage des résultats et gestion de l'enregistrement
		handleSave(result)
	}
}
