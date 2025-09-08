# Composants des Applications Mobiles Hoomi

Je documente ici les principaux composants que je vais implémenter dans les applications mobiles Hoomi, pour iOS et Android.

## Composants de Haut Niveau

### Écran de Démarrage (Splash Screen)

- **Description** : Premier écran affiché lors du lancement de l'application.
- **Fonction** : Afficher le logo pendant l'initialisation de l'application.
- **Durée** : Maximum 2 secondes.
- **Accessibilité** : Aucune interaction nécessaire.

### Écrans d'Authentification

#### Écran de Bienvenue

- **Description** : Premier écran pour les utilisateurs non connectés.
- **Contenu** : 
  - Logo de l'application.
  - Message de bienvenue.
  - Boutons "Se connecter" et "S'inscrire".
- **Navigation** : Vers l'écran de connexion ou d'inscription.

#### Écran de Connexion

- **Description** : Formulaire pour que les utilisateurs existants se connectent.
- **Champs** :
  - Email (champ de saisie)
  - Mot de passe (champ de saisie masqué)
- **Actions** :
  - Bouton "Se connecter".
  - Lien "Mot de passe oublié".
  - Lien "Créer un compte".
- **Validation** : Vérification du format de l'email et de la présence du mot de passe.
- **Sécurité** : Masquage du mot de passe, protection contre le brute force.

#### Écran d'Inscription

- **Description** : Formulaire pour que les nouveaux utilisateurs créent un compte.
- **Champs** :
  - Nom (champ de saisie)
  - Email (champ de saisie)
  - Mot de passe (champ de saisie masqué)
  - Confirmation du mot de passe (champ de saisie masqué)
- **Actions** :
  - Bouton "S'inscrire".
  - Lien "Déjà un compte? Se connecter".
- **Validation** : 
  - Format de l'email.
  - Force du mot de passe.
  - Correspondance des mots de passe.
- **Sécurité** : Politique de mot de passe stricte.

### Écran Principal (Home)

- **Description** : Écran principal après connexion, affichant la liste des Hoomis.
- **Contenu** :
  - Liste des Hoomis de l'utilisateur.
  - Bouton "Créer un Hoomi".
  - Profil utilisateur (menu déroulant ou page séparée).
- **Navigation** : 
  - Toucher un Hoomi pour y accéder.
  - Bouton "+" pour créer un nouveau Hoomi.
- **État** : Affichage d'un indicateur de chargement si la liste est en cours de récupération.

### Écran d'un Hoomi

- **Description** : Écran principal d'un Hoomi spécifique, avec les différentes zones.
- **Navigation par onglets** :
  - **Le Coeur** (Mur Émotionnel)
  - **La Mémoire** (Albums)
  - **Les Rendez-vous** (Événements)
  - **Le Jeu** (Loisirs)
  - **Le Cadre** (Profil/Paramètres)
- **Barre d'outils** :
  - Nom du Hoomi.
  - Indicateur du nombre de membres en ligne.
  - Menu des options du Hoomi.

## Composants de la Zone "Le Coeur" (Mur Émotionnel)

### Liste des Messages

- **Description** : Affichage des messages du Hoomi, triés du plus récent au plus ancien.
- **Éléments** :
  - Bulle de message avec avatar de l'expéditeur.
  - Contenu du message.
  - Horodatage.
  - Indicateur de statut (lu/envoyé).
- **Interactions** :
  - Scroll infini pour charger plus d'historique.
  - Swipe pour des actions rapides (répondre, copier, etc.).
- **États** :
  - Chargement.
  - Liste vide (message d'incitation).
  - Erreur de chargement.

### Saisie de Message

- **Description** : Zone de composition pour envoyer un nouveau message.
- **Éléments** :
  - Champ de texte multiligne.
  - Bouton d'envoi.
  - Bouton pour attacher un média.
  - Bouton pour sélectionner une émotion.
- **Fonctionnalités** :
  - Expansion automatique du champ de texte.
  - Envoi avec Entrée (option configurable).
  - Indicateur de saisie ("en train d'écrire").

### Message avec Média

- **Description** : Affichage d'un message contenant un média (image, vidéo, document).
- **Éléments** :
  - Aperçu du média.
  - Légende (optionnelle).
  - Actions (agrandir, télécharger, partager).
- **Types** :
  - Image : Affichage dans un viewer avec zoom.
  - Vidéo : Lecteur intégré avec contrôles.
  - Document : Icône avec nom du fichier et taille.

## Composants de la Zone "La Mémoire" (Albums)

### Liste des Albums

- **Description** : Grille d'albums de médias.
- **Éléments** :
  - Vignettes représentatives de chaque album.
  - Nom de l'album.
  - Nombre de médias dans l'album.
- **Actions** :
  - Toucher pour ouvrir l'album.
  - Swipe pour des options (renommer, supprimer).

### Album de Médias

- **Description** : Affichage des médias d'un album spécifique.
- **Présentation** :
  - Grille de vignettes.
  - Vue en liste (alternative).
- **Actions** :
  - Toucher pour agrandir un média.
  - Sélection multiple pour partager/supprimer.
  - Bouton pour ajouter des médias.

### Visualiseur de Médias

- **Description** : Affichage plein écran d'un média.
- **Fonctionnalités** :
  - Navigation entre les médias (swipe gauche/droite).
  - Zoom pour les images.
  - Contrôles de lecture pour les vidéos.
  - Informations sur le média (date, auteur).
  - Actions (télécharger, partager, supprimer).

## Composants de la Zone "Les Rendez-vous" (Événements)

### Calendrier des Événements

- **Description** : Vue calendrier des événements du Hoomi.
- **Modes** :
  - Vue mois.
  - Vue semaine.
  - Vue jour.
- **Éléments** :
  - Événements affichés sur les dates correspondantes.
  - Indicateur pour les jours avec des événements.
- **Actions** :
  - Toucher une date pour voir les événements.
  - Toucher un événement pour le voir en détail.

### Liste des Événements

- **Description** : Liste des événements à venir, triés par date.
- **Éléments** :
  - Titre de l'événement.
  - Date et heure.
  - Lieu (optionnel).
  - Nombre de participants.
- **Actions** :
  - Toucher pour voir les détails.
  - Swipe pour s'inscrire/se désinscrire.

### Détail d'un Événement

- **Description** : Page détaillant un événement spécifique.
- **Contenu** :
  - Titre.
  - Description.
  - Date et heure.
  - Lieu avec carte (si applicable).
  - Liste des participants.
  - Options d'inscription.
- **Actions** :
  - S'inscrire/se désinscrire.
  - Ajouter au calendrier du système.
  - Partager l'événement.

## Composants de la Zone "Le Jeu" (Loisirs)

### Liste des Jeux

- **Description** : Grille ou liste des jeux disponibles dans le Hoomi.
- **Éléments** :
  - Icône du jeu.
  - Nom du jeu.
  - Nombre de joueurs en ligne.
  - Bouton "Jouer".
- **Filtres** :
  - Par catégorie.
  - Par nombre de joueurs.

### Écran de Jeu

- **Description** : Interface d'un jeu spécifique.
- **Contenu** :
  - Zone de jeu principale.
  - Informations sur les joueurs.
  - Chat intégré (optionnel).
  - Contrôles du jeu.
- **Actions** :
  - Interactions spécifiques au jeu.
  - Quitter le jeu.

## Composants de la Zone "Le Cadre" (Profil/Paramètres)

### Profil Utilisateur

- **Description** : Informations et paramètres du compte utilisateur.
- **Contenu** :
  - Avatar et nom.
  - Email.
  - Date d'inscription.
  - Statistiques (nombre de messages, de médias partagés).
- **Actions** :
  - Modifier le profil.
  - Changer le mot de passe.
  - Gérer les appareils.

### Paramètres du Hoomi

- **Description** : Paramètres spécifiques au Hoomi actuel.
- **Sections** :
  - Informations du Hoomi (nom, date de création).
  - Liste des membres et leurs rôles.
  - Options de notification.
  - Paramètres de confidentialité.
- **Actions** :
  - Modifier le nom du Hoomi.
  - Inviter des membres.
  - Changer les rôles.
  - Quitter le Hoomi.
  - Supprimer le Hoomi (pour admin).

## Composants UI de Base

### Barre de Navigation

- **Description** : Barre supérieure de navigation dans l'application.
- **Éléments** :
  - Titre de la page.
  - Boutons d'action (options, recherche, etc.).
  - Indicateur de navigation (retour, menu).

### Menu de Navigation

- **Description** : Menu latéral ou bottom bar pour naviguer entre les sections principales.
- **Éléments** :
  - Liens vers les zones principales (Cœur, Mémoire, etc.).
  - Lien vers l'écran principal (liste des Hoomis).
  - Lien vers les paramètres.

### Boutons

- **Types** :
  - Primaire (action principale).
  - Secondaire (action alternative).
  - Destructif (suppression).
  - Flottant (action rapide).
- **États** :
  - Normal.
  - Désactivé.
  - En chargement.

### Cartes

- **Description** : Conteneur pour afficher des informations groupées.
- **Éléments** :
  - Titre.
  - Contenu principal.
  - Actions secondaires.
- **Variants** :
  - Carte simple.
  - Carte avec image.
  - Carte interactive.

### Listes et Cellules

- **Description** : Composants pour afficher des listes de données.
- **Types** :
  - Liste simple.
  - Liste avec sections.
  - Liste avec pull-to-refresh.
- **Cellules** :
  - Différents styles selon le contenu.
  - Swipe actions.
  - Sélection multiple.

### Dialogues et Alertes

- **Description** : Fenêtres modales pour les confirmations et les informations importantes.
- **Types** :
  - Alerte simple (information).
  - Confirmation (oui/non).
  - Saisie (formulaire court).
  - Action Sheet (liste d'actions).

## Composants Spécifiques à Hoomi

### Bulle de Message avec Émotion

- **Description** : Variante de la bulle de message qui inclut une icône d'émotion.
- **Éléments** :
  - Icône d'émotion.
  - Contenu du message.
  - Avatar de l'expéditeur.

### Silhouette de Personne

- **Description** : Représentation stylisée d'un membre du Hoomi.
- **Utilisation** :
  - Dans la liste des membres.
  - Comme avatar par défaut.
- **Variants** :
  - Différentes couleurs selon le rôle.
  - Indicateur d'en ligne/hors ligne.

### Indicateur de Saisie

- **Description** : Animation montrant qu'un membre est en train de taper un message.
- **Affichage** : 
  - Dans la barre d'état du Hoomi.
  - Dans la conversation en temps réel.

Cette liste couvre les principaux composants que je prévois d'implémenter. Elle sera complétée et affinée au fur et à mesure du développement.