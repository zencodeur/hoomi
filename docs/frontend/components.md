# Composants Frontend Hoomi

Ce document décrit les principaux composants de l'interface utilisateur de l'application Hoomi Web.

## Composants de Haut Niveau

### Layout Principal

- **`MainLayout`** : Le layout principal de l'application, englobant toutes les pages une fois l'utilisateur connecté.
  - **Props** : `children` (ReactNode)
  - **Slots** :
    - `header` : En-tête de l'application.
    - `sidebar` : Barre latérale pour la navigation entre les Hoomis.
    - `main` : Contenu principal de la page.
    - `footer` : Pied de page (optionnel).

### En-tête (Header)

- **`AppHeader`** : La barre de navigation en haut de l'application.
  - **Props** : `user` (objet utilisateur), `onLogout` (fonction)
  - **Slots** :
    - `logo` : Logo de l'application.
    - `userMenu` : Menu utilisateur (profil, paramètres, déconnexion).
    - `notifications` : Zone pour les notifications (à implémenter plus tard).

### Barre Latérale (Sidebar)

- **`HoomiSidebar`** : La barre latérale affichant la liste des Hoomis de l'utilisateur.
  - **Props** : `hoomis` (array), `activeHoomiId` (string), `onSelectHoomi` (fonction), `onCreateHoomi` (fonction)
  - **Slots** :
    - `hoomiList` : Liste des Hoomis.
    - `createButton` : Bouton pour créer un nouveau Hoomi.

### Liste des Hoomis

- **`HoomiList`** : Une liste de cartes représentant les Hoomis.
  - **Props** : `hoomis` (array), `activeHoomiId` (string), `onSelectHoomi` (fonction)
  - **Slots** :
    - `hoomiCard` : Carte individuelle pour chaque Hoomi.

### Carte Hoomi

- **`HoomiCard`** : Une carte représentant un Hoomi individuel dans la liste.
  - **Props** : `hoomi` (objet), `isActive` (bool), `onClick` (fonction)
  - **Slots** :
    - `name` : Nom du Hoomi.
    - `members` : Indicateur du nombre de membres.
    - `avatar` : Avatar/icone du Hoomi.

## Pages Principales

### Page d'Accueil (Landing)

- **`LandingPage`** : La page d'accueil pour les utilisateurs non connectés.
  - **Slots** :
    - `hero` : Section hero avec message principal.
    - `features` : Liste des fonctionnalités clés.
    - `cta` : Appel à l'action (liens vers login/register).

### Page de Connexion

- **`LoginPage`** : Le formulaire de connexion.
  - **Props** : `onLogin` (fonction)
  - **Slots** :
    - `form` : Formulaire de connexion.
    - `links` : Liens vers la page d'inscription et la récupération de mot de passe.

### Page d'Inscription

- **`RegisterPage`** : Le formulaire d'inscription.
  - **Props** : `onRegister` (fonction)
  - **Slots** :
    - `form` : Formulaire d'inscription.
    - `links` : Liens vers la page de connexion.

### Page du Hoomi

- **`HoomiPage`** : La page principale d'un Hoomi, une fois sélectionné.
  - **Props** : `hoomi` (objet)
  - **Slots** :
    - `header` : En-tête spécifique au Hoomi (nom, membres).
    - `tabs` : Onglets pour naviguer entre les différentes zones (Mur, Mémoire, Rendez-vous, Jeu, Cadre).
    - `content` : Contenu principal basé sur l'onglet actif.

## Composants de Zone d'Affinité

### Zone "Le Coeur" (Mur Émotionnel)

- **`HeartZone`** : La zone principale de communication.
  - **Slots** :
    - `messageList` : Liste des messages.
    - `messageInput` : Champ de saisie pour envoyer un message.
    - `emotions` : Sélecteur d'émotion pour le message (à implémenter plus tard).

### Zone "La Mémoire" (Albums)

- **`MemoryZone`** : La zone de gestion des photos et documents.
  - **Slots** :
    - `albumList` : Liste des albums.
    - `mediaViewer` : Visualiseur de médias sélectionnés.
    - `uploadArea` : Zone pour uploader de nouveaux médias.

### Zone "Les Rendez-vous" (Événements)

- **`AppointmentsZone`** : La zone de gestion des événements.
  - **Slots** :
    - `calendar` : Calendrier des événements.
    - `eventList` : Liste des événements à venir.
    - `eventForm` : Formulaire pour créer/modifier un événement.

### Zone "Le Jeu" (Loisirs)

- **`GameZone`** : La zone d'accès aux jeux familiaux.
  - **Slots** :
    - `gameList` : Liste des jeux disponibles.
    - `gameLauncher` : Interface pour lancer un jeu.

### Zone "Le Cadre" (Profil/Paramètres)

- **`FrameZone`** : La zone de gestion du profil utilisateur et des paramètres du Hoomi.
  - **Slots** :
    - `userProfile` : Section pour le profil utilisateur.
    - `hoomiSettings` : Section pour les paramètres du Hoomi.
    - `members` : Section pour gérer les membres du Hoomi.

## Composants UI de Base

### Bouton

- **`Button`** : Un bouton générique.
  - **Props** : `variant` (primary, secondary, danger), `size` (small, medium, large), `disabled` (bool), `onClick` (fonction), `children` (ReactNode)

### Champ de Saisie

- **`Input`** : Un champ de saisie générique.
  - **Props** : `type` (text, email, password), `label`, `placeholder`, `value`, `onChange` (fonction), `error` (string)

### Carte

- **`Card`** : Un conteneur de contenu générique.
  - **Props** : `title`, `children` (ReactNode), `actions` (ReactNode)

### Modal

- **`Modal`** : Une boîte de dialogue modale.
  - **Props** : `isOpen` (bool), `onClose` (fonction), `title`, `children` (ReactNode), `actions` (ReactNode)

## Composants Spécifiques à Hoomi

### Message

- **`MessageBubble`** : Une bulle de message individuelle.
  - **Props** : `message` (objet), `isOwn` (bool), `senderName` (string)

### Silhouette de Personne

- **`PersonSilhouette`** : Une représentation stylisée d'une personne.
  - **Props** : `size` (small, medium, large), `color` (string)

### Bulle d'Émotion

- **`EmotionBubble`** : Une petite bulle représentant une émotion.
  - **Props** : `emotion` (string), `onClick` (fonction)

Cette liste n'est pas exhaustive et sera complétée au fur et à mesure du développement.