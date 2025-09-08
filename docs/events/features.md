# Fonctionnalités du Système d'Événements Hoomi

Je documente ici les fonctionnalités détaillées du système d'événements d'Hoomi.

## 1. Création et Gestion des Événements

### Création Rapide

#### Via le Calendrier
- **Clic sur une date** : Ouvre un formulaire minimal pour créer un événement à cette date.
- **Glisser-déposer** : Création d'un événement en sélectionnant une plage horaire.

#### Via la Messagerie
- **Bouton "Planifier"** : Dans la barre d'outils de conversation.
- **Détection de date** : L'application suggère de créer un événement quand une date est mentionnée.

#### Formulaire de Création

##### Champs Obligatoires
- **Titre** : Nom de l'événement (ex: "Anniversaire de maman").
- **Date et Heure** : Début et fin de l'événement.
  - **Choix intuitif** : Calendrier visuel + horloge.
  - **Événement sur plusieurs jours** : Possibilité de définir une date de fin différente.
  - **Toute la journée** : Case à cocher pour les événements sans heure précise.

##### Champs Optionnels
- **Description** : Détails sur l'événement (activités prévues, tenue, etc.).
- **Lieu** :
  - **Nom** : Nom du lieu (ex: "Chez les Dupont").
  - **Adresse** : Adresse postale complète.
  - **Carte** : Intégration d'une carte interactive (OpenStreetMap/Google Maps).
  - **Coordonnées GPS** : Récupération automatique ou saisie manuelle.
- **Participants** : Sélection des membres du Hoomi à inviter.
- **Rappels** : Configuration des rappels (1 jour, 1 heure, 10 minutes avant).
- **Couleur** : Choix d'une couleur pour catégoriser l'événement.
- **Confidentialité** : Niveau de détail à partager avec les participants.

### Édition d'Événement

#### Modification
- **Accès** : Via le détail de l'événement ou en double-cliquant sur le calendrier.
- **Champs modifiables** : Tous les champs du formulaire de création.
- **Historique** : Conservation des modifications importantes.

#### Suppression
- **Confirmation** : Boîte de dialogue pour confirmer la suppression.
- **Impact** : Notifications envoyées à tous les participants.

### Événements Récurrents

#### Types de Récurrence
- **Quotidien** : Chaque jour.
- **Hebdomadaire** : Chaque semaine, jours de la semaine sélectionnables.
- **Mensuel** : Chaque mois, même date ou même jour de la semaine.
- **Annuel** : Chaque année, même date.
- **Personnalisé** : Fréquence et intervalle configurables.

#### Gestion de la Récurrence
- **Modification d'une instance** : Possibilité de modifier une seule occurrence.
- **Modification de la série** : Mise à jour de toutes les occurrences futures.
- **Exception** : Possibilité d'ajouter/supprimer des dates spécifiques.

## 2. Participants et Invitations

### Invitation de Membres

#### Sélection
- **Liste des membres** : Du Hoomi avec cases à cocher.
- **Suggestions** : Basé sur les interactions récentes dans la messagerie.
- **Groupes** : Possibilité de créer et d'inviter des groupes de membres.

#### Envoi d'Invitation
- **Notification immédiate** : Message dans la conversation du Hoomi.
- **Email (optionnel)** : Envoi d'un email si l'utilisateur l'a activé dans ses préférences.
- **Message personnalisé** : Possibilité d'ajouter un message à l'invitation.

### Réponse aux Invitations

#### Options de Réponse
- **Participe** : Confirme la participation.
- **Ne participe pas** : Indique un empêchement.
- **Peut-être** : Indique une probabilité conditionnelle de participation.

#### Modification de Réponse
- **Flexibilité** : Possibilité de changer de réponse à tout moment.
- **Notification** : Mise à jour automatique pour l'organisateur.

### Liste des Participants

#### Affichage
- **Avatar** : Photo de profil de chaque participant.
- **Statut** : Icône indiquant s'ils participent, ne participent pas ou sont incertains.
- **Rôle** : Organisateur, Participant.
- **Dernière réponse** : Date de la dernière mise à jour du statut.

#### Actions
- **Relance** : Possibilité d'envoyer un rappel aux participants qui n'ont pas répondu.
- **Contact** : Lien rapide vers la messagerie pour contacter un participant.

## 3. Calendrier et Vue des Événements

### Vues Calendrier

#### Mois
- **Aperçu global** : Vue d'ensemble des événements sur un mois.
- **Indicateurs** : Points colorés pour les événements de chaque jour.
- **Navigation** : Flèches pour passer au mois précédent/suivant.

#### Semaine
- **Détail hebdomadaire** : Vue détaillée d'une semaine.
- **Grille horaire** : Affichage des heures de 6h à 24h.
- **Superposition** : Possibilité d'afficher les événements de plusieurs Hoomis.

#### Jour
- **Vue détaillée** : Liste chronologique des événements de la journée.
- **Timeline** : Ligne de temps pour situer visuellement les événements.
- **Événements chevauchants** : Gestion des événements qui se chevauchent.

#### Agenda
- **Liste linéaire** : Affichage des événements à venir dans l'ordre chronologique.
- **Filtrage** : Tri par date, par participant, par type.

### Navigation et Interaction

#### Navigation Rapide
- **Aujourd'hui** : Bouton pour revenir à la date du jour.
- **Saut de date** : Saisie directe d'une date.
- **Recherche** : Recherche d'événements par titre, description, participant.

#### Glisser-Déposer
- **Déplacement** : Déplacer un événement vers une autre date/heure par glisser-déposer.
- **Redimensionnement** : Modifier la durée d'un événement en redimensionnant ses bords.

### Personnalisation de l'Affichage

#### Couleurs
- **Catégorisation** : Chaque type d'événement (anniversaire, sortie, tâche ménagère) a une couleur.
- **Personnalisation** : Possibilité de changer les couleurs par défaut.
- **Héritage** : Les couleurs peuvent être héritées du Hoomi parent.

#### Filtres
- **Par participant** : Afficher uniquement les événements auxquels participe un membre spécifique.
- **Par type** : Filtrer par catégories d'événements.
- **Par Hoomi** : Afficher les événements d'un ou plusieurs Hoomis spécifiques.

## 4. Notifications et Rappels

### Notifications en Temps Réel

#### Création d'Événement
- **Pour les invités** : Notification immédiate de l'invitation.
- **Pour l'organisateur** : Confirmation de création.

#### Modification d'Événement
- **Pour les participants** : Notification des changements (horaire, lieu, description).
- **Détail des changements** : Indication précise de ce qui a été modifié.

#### Réponse à l'Invitation
- **Pour l'organisateur** : Notification de la réponse d'un participant.
- **Mise à jour de la liste** : Actualisation automatique de la liste des participants.

### Rappels Configurables

#### Anticipation
- **Par défaut** : 1 jour, 1 heure et 10 minutes avant l'événement.
- **Personnalisable** : Possibilité de définir des rappels à 1 semaine, 3 jours, etc.
- **Multiples** : Plusieurs rappels possibles pour le même événement.

#### Canaux de Rappel
- **Notification push** : Notification dans l'application.
- **Email** : Envoi d'un email rappel.
- **SMS (futur)** : Envoi d'un SMS pour les rappels critiques.

#### Contenu des Rappels
- **Essentiel** : Titre, date, heure, lieu de l'événement.
- **Liens rapides** : Boutons pour participer/ne pas participer directement depuis la notification.
- **Personnalisé** : Possibilité d'ajouter un message personnalisé au rappel.

### Intégration avec les Systèmes Natifs

#### Calendrier Système
- **Synchronisation** : Option pour synchroniser les événements Hoomi avec le calendrier natif.
- **Notifications natives** : Utilisation des rappels du système d'exploitation.

#### Horloge
- **Alarmes** : Possibilité de définir une alarme sonore pour les rappels.

## 5. Intégration et Partage

### Export d'Événements

#### Formats Supportés
- **iCalendar (.ics)** : Format standard pour les échanges calendrier.
- **PDF** : Impression ou partage d'un résumé de l'événement.

#### Contenu Exporté
- **Métadonnées** : Titre, dates, description.
- **Participants** : Liste des participants et leurs réponses.
- **Lieu** : Coordonnées et carte du lieu.

### Partage d'Événements

#### Lien Partagé
- **URL unique** : Génération d'un lien pour partager un événement.
- **Accès restreint** : Le lien ne fonctionne que pour les membres du Hoomi.

#### Intégration dans la Messagerie
- **Carte d'événement** : Génération automatique d'une carte cliquable dans les messages.
- **Mise à jour dynamique** : La carte se met à jour si l'événement change.

### Synchronisation avec les Calendriers Externes

#### Fournisseurs Supportés
- **Google Calendar** : Synchronisation via OAuth2.
- **Apple Calendar** : Synchronisation via CalDAV.
- **Microsoft Outlook** : Synchronisation via Exchange ActiveSync ou CalDAV.

#### Configuration
- **Authentification** : OAuth2 avec consentement explicite de l'utilisateur.
- **Sélection des calendriers** : Choix des calendriers externes à synchroniser.
- **Direction** : Synchronisation unidirectionnelle (Hoomi → Externe) ou bidirectionnelle.

#### Gestion des Conflits
- **Priorité** : Hoomi a la priorité par défaut en cas de conflit.
- **Résolution manuelle** : Interface pour résoudre les conflits de synchronisation.

## 6. Événements Spéciaux

### Anniversaires

#### Détection Automatique
- **Profil utilisateur** : Récupération de la date de naissance du profil.
- **Création automatique** : Génération d'un événement annuel pour chaque anniversaire.

#### Personnalisation
- **Template** : Modèle d'anniversaire personnalisable par Hoomi.
- **Rappels** : Rappels spéciaux (1 semaine avant, le matin de l'anniversaire).

#### Célébration
- **Bannière spéciale** : Affichage d'une bannière sur l'interface le jour de l'anniversaire.
- **Message automatique** : Envoi d'un message de félicitations dans le fil de discussion.

### Tâches Récurrentes

#### Gestion des Corvées
- **Rotation** : Attribution automatique des tâches aux membres du Hoomi.
- **Historique** : Suivi de qui a effectué quoi et quand.

#### Rappels de Tâches
- **Anticipation** : Rappels 1 jour à l'avance.
- **Suivi** : Notification quand une tâche est marquée comme terminée.

### Voyages et Événements sur Plusieurs Jours

#### Planificateur de Voyage
- **Jours multiples** : Création de plans détaillés pour chaque jour du voyage.
- **Logements** : Gestion des réservations d'hôtel, Airbnb, etc.
- **Transports** : Ajout de billets de train, avion, etc.

#### Itinéraire
- **Carte interactive** : Affichage du trajet sur une carte.
- **Points d'intérêt** : Ajout de lieux à visiter.
- **Partage** : Possibilité de partager l'itinéraire avec des non-membres.

## 7. Accessibilité et Internationalisation

### Accessibilité

#### Navigation au Clavier
- **Raccourcis** : Ctrl+N pour nouveau, flèches pour naviguer.
- **Focus visible** : Indication claire de l'élément ayant le focus.

#### Lecteurs d'Écran
- **Balises ARIA** : Utilisation correcte des attributs d'accessibilité.
- **Alternatives textuelles** : Pour les icônes et images.

#### Options Visuelles
- **Mode contraste élevé** : Palette de couleurs pour les utilisateurs malvoyants.
- **Taille de texte** : Ajustement de la taille des textes.
- **Animations** : Possibilité de réduire les animations.

### Internationalisation

#### Langues Supportées
- **Français** : Langue par défaut.
- **Anglais** : Traduction complète.
- **Espagnol** : (Futur) Traduction automatique avec vérification humaine.

#### Formats Locaux
- **Dates et heures** : Adaptation au format local (DD/MM/YYYY vs MM/DD/YYYY).
- **Fuseaux horaires** : Gestion automatique du fuseau horaire de l'utilisateur.
- **Premier jour de la semaine** : Dimanche ou lundi selon la région.

Cette liste de fonctionnalités couvre les aspects essentiels du système d'événements d'Hoomi, en mettant l'accent sur la simplicité d'utilisation, la collaboration et l'intégration avec les autres parties de l'application.