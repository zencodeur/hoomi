# Intégration du Système d'Événements Hoomi

Je documente ici comment le système d'événements s'intègre avec les autres composants d'Hoomi.

## 1. Intégration avec la Messagerie

### Création d'Événements depuis la Conversation

#### Déclenchement
- **Mot-clé** : Lorsqu'un utilisateur tape des mots comme "rendez-vous", "événement", "planifions" suivi d'une date.
- **Bouton "Planifier"** : Disponible dans la barre d'outils de la conversation.

#### Processus
1. L'utilisateur clique sur "Planifier" ou tape un mot-clé.
2. Une fenêtre modale s'ouvre avec un formulaire pré-rempli :
   - **Titre** : Basé sur le contexte de la conversation.
   - **Date** : Extraite du message ou date courante par défaut.
   - **Participants** : Membres actifs de la conversation.
3. L'utilisateur complète/modifie les détails.
4. L'événement est créé et un message spécial est envoyé dans la conversation :
   ```
   [événement] Anniversaire de maman
   📅 Mercredi 15 mai 2024 à 18h00
   📍 Chez les Dupont
   👥 3 participants : Alice, Bob, Charlie
   ```

### Discussion Dédiée à l'Événement

#### Canal Spécifique
- Chaque événement a un **canal de discussion dédié**.
- Accessible via un onglet "Discussion" dans la vue détaillée de l'événement.

#### Fonctionnalités
- **Messages liés** : Discussions spécifiquement sur l'événement.
- **Partages** : Possibilité de partager des fichiers, liens directement dans ce canal.
- **Décisions** : Votes ou sondages pour les détails de l'événement (lieu, menu, etc.).

#### Notifications
- Les messages dans ce canal génèrent des notifications pour tous les participants à l'événement.

### Liens dans les Messages

#### Auto-détection
- Les liens vers des événements Hoomi sont automatiquement transformés en **cartes interactives**.
- Ces cartes affichent :
  - Titre de l'événement
  - Date et heure
  - Lieu
  - Nombre de participants
  - Boutons "Participer"/"Ne pas participer"

#### Mise à Jour en Temps Réel
- Si l'événement est modifié, la carte dans le message est mise à jour automatiquement.
- Les statuts de participation des utilisateurs sont mis à jour en temps réel.

## 2. Intégration avec le Stockage de Fichiers

### Pièces Jointes aux Événements

#### Types de Fichiers
- **Documents** : Plans, listes de courses, recettes.
- **Images** : Photos du lieu, inspiration pour la décoration.
- **Vidéos** : Souvenirs d'événements passés, tutoriels.

#### Gestion
- **Section "Fichiers"** : Dans la vue détaillée de l'événement.
- **Glisser-déposer** : Pour ajouter facilement des fichiers.
- **Prévisualisation** : Miniatures pour les images, icônes pour les documents.

#### Permissions
- **Accès restreint** : Seuls les membres du Hoomi peuvent accéder aux fichiers.
- **Chiffrement** : Les fichiers liés à l'événement bénéficient du même chiffrement E2EE.

### Invitations Graphiques

#### Création
- **Templates** : Modèles d'invitations personnalisables.
- **Éditeur visuel** : Interface de création d'invitations avec texte, images, formes.

#### Partage
- **Envoi direct** : Possibilité d'envoyer l'invitation par message dans Hoomi.
- **Téléchargement** : Export en PDF ou image pour impression/partage externe.

#### Suivi
- **Ouvertures** : (Futur) Suivi des ouvertures pour les invitations envoyées par email.

### Souvenirs d'Événements Passés

#### Album Automatique
- **Compilation** : Création automatique d'un album avec :
  - Photos partagées dans le canal de discussion de l'événement.
  - Copies d'écran des messages importants.
  - Fichiers liés à l'événement.

#### Organisation
- **Par événement** : Un album par événement dans une section "Souvenirs".
- **Balises** : Possibilité d'ajouter des balises pour retrouver facilement.

#### Partage
- **Au sein du Hoomi** : Album partageable dans la zone "Mémoire".
- **Extérieur** : (Futur) Possibilité de créer un album public pour partager avec des proches externes.

## 3. Intégration avec les Jeux

### Événements de Jeu

#### Création Spéciale
- **Type d'événement** : Nouveau type "Soirée Jeu" avec des options spécifiques :
  - **Choix du jeu** : Sélection parmi les jeux disponibles dans Hoomi.
  - **Nombre de joueurs** : Indication du nombre minimum/maximum.
  - **Matériel** : Liste du matériel nécessaire (jeu physique, jetons, etc.).

#### Préparation
- **Liste de vérification** : Génération automatique d'une liste des choses à préparer.
- **Invitation aux joueurs** : Suggestions basées sur les membres qui jouent souvent aux mêmes jeux.

### Tournois

#### Organisation
- **Système de tournoi** : Intégration directe des jeux dans un système de tournoi.
- **Brackets** : Génération automatique des tableaux d'affrontement.
- **Suivi des scores** : Mise à jour en temps réel des résultats.

#### Communication
- **Canal dédié** : Discussion spéciale pour le tournoi avec annonces, résultats.
- **Notifications** : Alertes pour les matchs à venir, résultats disponibles.

## 4. Intégration avec le Profil Utilisateur

### Historique des Événements

#### Vue Personnelle
- **Calendrier personnel** : Agrégation des événements de tous les Hoomis auxquels l'utilisateur appartient.
- **Filtres** : Par Hoomi, par type d'événement, par statut de participation.

#### Statistiques
- **Événements organisés** : Nombre d'événements créés par l'utilisateur.
- **Participations** : Nombre d'événements auxquels il a participé.
- **Disponibilité** : Taux de participation/absence.

### Préférences d'Événements

#### Configuration
- **Types d'événements** : Préférences pour les types d'événements à suivre.
- **Lieux fréquentés** : Historique des lieux pour suggestions rapides.
- **Participants favoris** : Membres avec lesquels l'utilisateur organise souvent des événements.

#### Notifications
- **Personnalisation** : Choix des types d'événements pour lesquels recevoir des notifications.
- **Horaires préférés** : Suggestions d'horaires basées sur les disponibilités passées.

## 5. Intégration avec les Calendriers Externes

### Synchronisation Bidirectionnelle

#### Protocole iCalendar
- **Flux iCalendar** : Génération d'une URL unique pour chaque Hoomi.
- **Mise à jour automatique** : Les événements se synchronisent automatiquement.

#### Champs Synchronisés
- **Titre**
- **Description**
- **Date et heure**
- **Lieu**
- **Participants** (en tant que "attendees" dans iCalendar)

#### Gestion des Conflits
- **Timestamp** : Utilisation des timestamps pour déterminer la version la plus récente.
- **Priorité utilisateur** : Possibilité de définir quelle source a la priorité en cas de conflit.

### Import d'Événements Externes

#### Formats Supportés
- **Fichiers .ics** : Import par glisser-déposer.
- **Liens calendrier** : Ajout d'un calendrier externe à synchroniser.

#### Traitement
- **Création automatique** : Les événements externes sont créés comme des événements Hoomi spéciaux.
- **Marquage** : Indication visuelle que l'événement vient d'un calendrier externe.
- **Mise à jour** : Les modifications dans le calendrier externe sont répercutées.

## 6. Intégration avec les Notifications Système

### Notifications Push

#### Déclencheurs
- **Nouvel événement** : Quand un utilisateur est invité.
- **Modification** : Quand un événement auquel l'utilisateur participe est modifié.
- **Rappel** : Selon les préférences de rappel configurées.
- **Réponse** : Quand un autre participant répond à une invitation.

#### Contenu
- **Titre** : Nom du Hoomi + Type de notification.
- **Corps** : Détails pertinents (nom de l'événement, changement, rappel).
- **Actions rapides** : Boutons "Participer"/"Ne pas participer" directement dans la notification.

### Notifications Email

#### Conditions
- **Activation** : L'utilisateur doit avoir activé les notifications email dans ses préférences.
- **Types** : Configuration des types de notifications à recevoir par email.

#### Template
- **Design responsive** : Adapté aux mobiles et desktops.
- **Informations clés** : Titre, date, lieu, participants, bouton d'action principal.

### Intégration avec l'Horloge Système

#### Alarmes
- **Configuration** : Possibilité de définir une alarme sonore pour les rappels.
- **Son personnalisé** : Choix du son d'alarme parmi plusieurs options.

#### Widget (Mobile)
- **Événements à venir** : Affichage des prochains événements dans un widget écran d'accueil.
- **Mise à jour automatique** : Le widget se met à jour en temps réel.

## 7. Intégration avec l'Authentification et les Autorisations

### Permissions sur les Événements

#### Rôles
- **Organisateur** :
  - Peut modifier/supprimer l'événement.
  - Peut inviter/retirer des participants.
  - Peut modifier les rôles des autres participants.
- **Participant** :
  - Peut voir les détails complets.
  - Peut modifier son statut de participation.
  - Peut participer à la discussion dédiée.
- **Invité** : (Futur)
  - Accès limité via un lien spécial.
  - Possibilité de confirmer sa participation.

#### Vérification
- **À chaque action** : Vérification des permissions avant chaque opération.
- **Messages d'erreur** : Messages clairs si l'utilisateur n'a pas les permissions.

### Sécurité des Données

#### Chiffrement
- **Données sensibles** : Les descriptions détaillées, les lieux précis peuvent être chiffrés.
- **Clés partagées** : Les clés de chiffrement sont partagées avec tous les participants.

#### Audit
- **Journalisation** : Enregistrement des actions importantes (création, modification, suppression).
- **Traçabilité** : Possibilité de voir qui a fait quoi et quand.

Cette intégration étroite entre le système d'événements et les autres composants d'Hoomi permet de créer une expérience unifiée et cohérente, où chaque outil complète les autres pour faciliter la vie familiale numérique.