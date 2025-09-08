# Capsules Temporelles Hoomi

Je documente ici le système de capsules temporelles d'Hoomi, qui permet aux membres de la famille de créer des messages, des contenus multimédias ou des instructions programmés pour être révélés à un moment futur prédéterminé ou lors d'un événement spécifique.

## Vue d'Ensemble

Les **Capsules Temporelles** d'Hoomi sont conçues pour être :
- **Personnalisables** : Contenu riche (texte, images, vidéos, documents).
- **Programmables** : Déclenchement basé sur le temps, les événements ou une combinaison.
- **Sécurisées** : Chiffrées de bout en bout et accessibles uniquement aux destinataires désignés.
- **Émotionnelles** : Créées pour transmettre des sentiments, des souvenirs, des conseils ou des surprises.
- **Surprenantes** : Révélées de manière inattendue pour créer des moments magiques.

## Interface Utilisateur

### Web (Next.js)

#### Création d'une Capsule

##### Formulaire de Base
- **Titre** : Champ texte pour nommer la capsule.
- **Description** : Champ texte optionnel pour décrire la capsule.
- **Destinataires** : Sélection des membres du Hoomi qui recevront la capsule.
- **Contenu** : Éditeur riche pour le texte + zone de dépôt pour les médias.

##### Programmation
- **Déclenchement temporel** :
  - **Date et heure précise** : Calendrier + horloge.
  - **Délai** : Dans X jours/semaines/années.
  - **Événement annuel** : Chaque année à une date spécifique.
- **Déclenchement événementiel** :
  - **Naissance** : Quand un nouveau membre rejoint le Hoomi.
  - **Anniversaire** : À l'anniversaire d'un membre spécifique.
  - **Mariage** : Quand un membre se marie (événement créé dans le calendrier).
  - **Décès** : (Futur) Quand un membre est marqué comme décédé.
- **Combinaison** : Possibilité de combiner déclenchement temporel et événementiel.

##### Aperçu et Envoi
- **Aperçu** : Visualisation de la capsule telle qu'elle apparaîtra aux destinataires.
- **Confirmation** : Récapitulatif des paramètres de déclenchement.
- **Bouton d'envoi** : Création définitive de la capsule.

#### Gestion des Capsules

##### Liste des Capsules
- **Grille ou liste** : Affichage des capsules créées par l'utilisateur.
- **Colonnes** :
  - Titre
  - Destinataires
  - Type de déclenchement
  - Date programmée
  - Statut (programmée, déclenchée, livrée)
- **Filtres** : Par statut, par date, par type de déclenchement.

##### Détails d'une Capsule
- **Contenu** : Affichage complet du contenu de la capsule.
- **Métadonnées** :
  - Créateur
  - Date de création
  - Destinataires
  - Conditions de déclenchement
- **Actions** :
  - **Modifier** : (Tant que non déclenchée) Modifier les paramètres.
  - **Tester** : Déclencher manuellement la capsule pour vérifier.
  - **Annuler** : Supprimer la capsule programmée.
  - **Dupliquer** : Créer une nouvelle capsule basée sur celle-ci.

### Mobile (iOS/Android)

#### Création Rapide

##### Assistant de Création
- **Étapes simples** :
  1. **Quoi** : Texte, photo, vidéo.
  2. **À qui** : Sélection des destinataires.
  3. **Quand** : Date ou événement.
- **Modèles** : Suggestions de capsules basées sur le contexte (anniversaire proche, nouveau bébé).

##### Capture Intégrée
- **Caméra/Photothèque** : Capture ou sélection directe de médias.
- **Enregistrement vocal** : Création de messages audio.
- **Dessin** : (Futur) Création de dessins à inclure dans la capsule.

#### Gestion sur Mobile

##### Notifications
- **Création** : Notification quand une capsule est programmée.
- **Déclenchement** : Notification quand une capsule est déclenchée.
- **Livraison** : Notification quand une capsule est livrée (si l'utilisateur est destinataire).

##### Interface de Gestion
- **Onglet dédié** : Dans le menu principal de l'application.
- **Liste simplifiée** : Cartes pour chaque capsule avec informations essentielles.
- **Actions rapides** : Swipe pour annuler/tester/dupliquer.

## Types de Capsules

### 1. Capsules Personnelles

#### Messages Intimes
- **Lettres d'amour** : Pour un conjoint à lire à un anniversaire.
- **Conseils de vie** : Pour ses enfants à lire à un certain âge.
- **Souvenirs** : Partage de souvenirs d'enfance avec les petits-enfants.

#### Vidéos de Vœux
- **Anniversaires** : Vidéos de vœux programmées pour s'envoyer automatiquement.
- **Fêtes** : Messages pour Noël, Pâques, aid, etc.
- **Moments spéciaux** : Vidéos pour les diplômes, les mariages.

### 2. Capsules Collectives

#### Messages du Clan
- **Bienvenue** : Messages de bienvenue pour un nouveau-né.
- **Souvenirs partagés** : Vidéos rassemblant les souvenirs de vacances.
- **Instructions** : Pour un événement familial futur (organisation de retrouvailles).

#### Albums Surprise
- **Souvenirs d'événements** : Albums photos de vacances ou fêtes de famille à révéler plus tard.
- **Hommages** : Albums créés après un décès pour honorer la mémoire.

### 3. Capsules Événementielles

#### Pour les Décès
- **Messages d'adieu** : Lettres à lire en cas de décès.
- **Instructions post-mortem** : Que faire des comptes numériques, des objets personnels.
- **Souvenirs** : Albums ou vidéos de souvenirs à partager avec la famille.

#### Pour les Naissances
- **Lettres aux nouveau-nés** : Messages des grands-parents, oncles, tantes.
- **Prédictions affectueuses** : Ce qu'on souhaite à l'enfant pour son avenir.
- **Souvenirs d'enfance** : Histoires de la propre enfance des parents/grands-parents.

#### Pour les Mariages
- **Conseils conjugaux** : Lettres de conseil d'anciens mariés.
- **Souvenirs de rencontre** : Comment les grands-parents se sont rencontrés.
- **Vœux pour le couple** : Messages de soutien pour le couple.

## Fonctionnalités Avancées

### Contenu Riche

#### Éditeur de Texte
- **Mise en forme** : Gras, italique, souligné, titres.
- **Listes** : À puces, numérotées.
- **Liens** : Vers d'autres personnes de l'arbre, vers des médias.
- **Émoticônes** : Pour ajouter de l'émotion.

#### Médias dans les Capsules
- **Images** : Jusqu'à 10 images par capsule.
- **Vidéos** : Vidéos courtes (jusqu'à 2 minutes) pour les messages personnels.
- **Documents** : PDF, lettres scannées, documents importants.
- **Audio** : Messages vocaux, chansons, enregistrements.

### Programmation Sophistiquée

#### Déclenchements Multiples
- **ET logique** : La capsule se déclenche uniquement si toutes les conditions sont remplies.
  - Ex : Le 18ème anniversaire *ET* si la personne est membre du Hoomi.
- **OU logique** : La capsule se déclenche si au moins une condition est remplie.
  - Ex : Au mariage *OU* à l'achat d'une maison.

#### Conditions Complexes
- **Âge** : La capsule se déclenche quand un destinataire atteint un certain âge.
- **Statut** : La capsule se déclenche si un destinataire a un certain statut (marié, parent).
- **Localisation** : (Futur) La capsule se déclenche si un destinataire est à un endroit spécifique.

### Interactivité

#### Capsules Réactives
- **Réponses** : (Futur) Possibilité pour les destinataires de répondre à une capsule, créant un dialogue temporel.
- **Quizz** : (Futur) Capsules avec des questions/réponses pour tester les connaissances familiales.
- **Défis** : (Futur) Capsules proposant des défis à relever (généalogie, cuisine familiale).

## Intégration avec les Autres Systèmes Hoomi

### Messagerie

#### Livraison
- **Message spécial** : Les capsules sont livrées sous forme de messages spéciaux dans la conversation du Hoomi.
- **Notification push** : Notification système en plus du message dans la conversation.

#### Création Assistée
- **Depuis une conversation** : Bouton pour créer une capsule à partir d'un message important.
- **Médias joints** : Possibilité d'inclure automatiquement les médias de la conversation.

### Médias

#### Médias dans les Capsules
- **Sélection depuis la galerie** : Choix parmi les médias existants dans Hoomi.
- **Albums spéciaux** : Création d'albums qui deviennent des capsules à une date donnée.

#### Capsules vers Médias
- **Archivage** : Les capsules déclenchées peuvent être automatiquement archivées dans un album spécial.
- **Partage** : Possibilité de partager le contenu d'une capsule déclenchée dans la galerie.

### Événements

#### Déclenchement Événementiel
- **Liaison directe** : Lier une capsule à un événement du calendrier Hoomi.
- **Création depuis un événement** : Bouton pour créer une capsule liée lors de la création d'un événement.

#### Capsules dans les Événements
- **Programme** : Inclure des capsules dans l'organisation d'un événement futur.
- **Souvenirs** : Créer automatiquement des capsules à partir des médias d'un événement passé.

### Arbre Généalogique

#### Personnes dans les Capsules
- **Destinataires spécifiques** : Lier une capsule à une personne spécifique de l'arbre (ex: message du grand-père à son petit-fils).
- **Événements généalogiques** : Déclenchement basé sur les dates de l'arbre (anniversaire d'un ancêtre).

#### Capsules Généalogiques
- **Histoires familiales** : Capsules contenant des histoires liées à des branches de l'arbre.
- **Souvenirs d'ancêtres** : Messages ou médias d'ancêtres à transmettre aux générations futures.

## Sécurité et Confidentialité

### Chiffrement

#### Données au Repos
- **Chiffrement AES-256** : Toutes les capsules sont chiffrées avant d'être stockées.
- **Clés chiffrées** : Les clés de chiffrement sont chiffrées avec la clé publique de chaque destinataire.

#### Données en Transit
- **HTTPS/WSS** : Toutes les communications sont chiffrées.
- **Chiffrement bout-en-bout** : Le serveur ne voit jamais le contenu en clair.

### Contrôle d'Accès

#### Destinataires
- **Vérification stricte** : Seuls les destinataires désignés peuvent accéder au contenu.
- **Ajout de destinataires** : (Futur) Possibilité d'ajouter des destinataires tant que la capsule n'est pas déclenchée.

#### Statut de la Capsule
- **Non modifiable après déclenchement** : Une fois déclenchée, une capsule ne peut plus être modifiée.
- **Journalisation** : Enregistrement de toutes les actions sur les capsules.

## Performance et Fiabilité

### Surveillance des Déclenchements

#### Service de Surveillance
- **Vérification régulière** : Workers qui vérifient périodiquement les conditions de déclenchement.
- **File d'attente** : Gestion des capsules prêtes à être déclenchées.

#### Tolérance aux Pannes
- **Redondance** : Plusieurs workers pour éviter les points de défaillance unique.
- **Journalisation** : Enregistrement de toutes les vérifications et déclenchements.

### Livraison Garantie

#### Mécanismes de Rejeu
- **Retry** : Tentatives multiples en cas d'échec de livraison.
- **Dead Letter Queue** : Capsules dont la livraison a échoué de manière permanente.

#### Notifications d'Échec
- **Alertes administrateurs** : En cas de problème persistant avec les déclenchements.
- **Logs détaillés** : Pour diagnostiquer les problèmes.

## Accessibilité

### Interfaces Adaptées

#### Navigation au Clavier
- **Raccourcis** : Tab, Entrée, Flèches pour naviguer dans les formulaires de création.
- **Validation** : Entrée pour valider/soumettre.

#### Lecteurs d'Écran
- **Balises ARIA** : Pour décrire les éléments interactifs.
- **Labels** : Textes alternatifs pour les boutons et les champs.

### Options Visuelles

#### Personnalisation
- **Taille du texte** : Ajustement pour les utilisateurs malvoyants.
- **Contraste élevé** : Palettes de couleurs avec contraste élevé.
- **Animations réduites** : Option pour réduire les animations.

## Internationalisation

### Langues Supportées
- **Français** : Langue par défaut.
- **Anglais** : Traduction complète.
- **(Futur)** Espagnol : Traduction automatique avec vérification humaine.

### Formats Locaux
- **Dates** : Adaptation au format local (DD/MM/YYYY vs MM/DD/YYYY).
- **Fuseaux horaires** : Gestion automatique du fuseau horaire de l'utilisateur.
- **Premier jour de la semaine** : Dimanche ou lundi selon la région.

Le système de capsules temporelles d'Hoomi offre une manière unique et émotionnelle de relier les générations, de préserver les souvenirs et de créer des surprises inattendues qui renforcent les liens familiaux à travers le temps.