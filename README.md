# Network_sh_recon

Ce script en Go est un outil de reconnaissance réseau permettant d'interroger plusieurs API pour collecter des informations sur des IPs ou des domaines. Il utilise les services suivants :

- **ipinfo.io** : Pour obtenir des informations sur une adresse IP.
- **Shodan** : Pour rechercher des informations sur des appareils connectés à Internet.
- **GreyNoise** : Pour analyser des IPs et leurs activités.
- **Censys** : Pour obtenir des informations détaillées sur des hôtes.
- **crt.sh** : Pour rechercher des sous-domaines et des certificats associés à un domaine.

## Prérequis

- Go version 1.16 ou supérieure.
- Les clés API pour les services ci-dessus sont nécessaires pour effectuer les requêtes. Vous devrez vous inscrire à chaque service et remplacer les valeurs des clés API dans le script.

## Installation

### 1. Cloner le repository

```bash
git clone https://github.com/votre-utilisateur/Network_sh_recon.git
cd Network_sh_recon

2. Installer Go (si ce n'est pas déjà fait)

Assurez-vous d'avoir Go installé sur votre machine. Vous pouvez le télécharger et l'installer depuis le site officiel de Go.
3. Ajouter vos clés API

Avant d'exécuter le script, remplacez les valeurs de vos clés API dans le script network_tool.go à l'endroit suivant :

const (
    ipinfoAPIKey    = "VOTRE_CLE_IPINFO"
    shodanAPIKey    = "VOTRE_CLE_SHODAN"
    greynoiseAPIKey = "VOTRE_CLE_GREYNOISE"
    censysAPIKey    = "VOTRE_CLE_CENSYS"
    censysSecret    = "VOTRE_SECRET_CENSYS"
)

4. Exécuter le script

Une fois que vous avez configuré vos clés API, exécutez le script avec la commande suivante :

go run Network_sh_recon.go

Utilisation

Une fois le script lancé, un menu interactif s'affichera dans le terminal, offrant plusieurs options :

    1. Rechercher avec ipinfo : Obtenez des informations détaillées sur une adresse IP.
    2. Rechercher avec Shodan : Recherchez des informations sur un appareil connecté via Shodan.
    3. Rechercher avec GreyNoise : Analysez une IP pour en savoir plus sur ses activités avec GreyNoise.
    4. Rechercher avec Censys : Obtenez des informations détaillées sur un hôte via Censys.
    5. Rechercher avec crt.sh : Recherchez les sous-domaines associés à un domaine en utilisant crt.sh.
    6. Quitter : Quittez l'outil.

![network_tool2](https://github.com/user-attachments/assets/740d5b82-abfb-4802-81a4-ec81fc5a7454)


Pour chaque option, il vous sera demandé d'entrer un domaine ou une adresse IP à analyser. Vous pourrez ensuite choisir d'enregistrer les résultats dans un fichier texte.
Contributions

Les contributions sont les bienvenues ! Si vous avez des suggestions d'améliorations ou si vous avez corrigé un bug, n'hésitez pas à soumettre une pull request.
Auteurs

    Prehelios : Développeur principal
