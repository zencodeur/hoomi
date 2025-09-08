# Architecture du Système d'Événements Hoomi

Je documente ici l'architecture que j'ai conçue pour le système d'événements d'Hoomi.

## Vue d'Ensemble

Le système d'événements de Hoomi est conçu pour être :
- **Intuitif** : Interface simple pour créer, gérer et participer aux événements.
- **Collaboratif** : Outils pour planifier ensemble et communiquer autour des événements.
- **Intégré** : Intégration étroite avec le système de messagerie et le stockage de fichiers.
- **Connecté** : Synchronisation avec les calendriers externes.
- **Sécurisé** : Respect de la confidentialité des événements familiaux.

## Architecture Générale

### Structure en Couches

```
┌─────────────────────────────────────┐
│         UI / Interface              │
├─────────────────────────────────────┤
│      Logique des Événements        │
├─────────────────────────────────────┤
│         API des Événements         │
├─────────────────────────────────────┤
│      Gestion des Notifications      │
├─────────────────────────────────────┤
│        Stockage des Données        │
└─────────────────────────────────────┘
```

## Composants Principaux

### 1. Service de Gestion des Événements

#### Responsabilités
- Création, modification, suppression des événements.
- Gestion des participants et des invitations.
- Stockage des métadonnées des événements.

#### Endpoints API Principaux
- `POST /hoomis/{id}/events` : Créer un événement.
- `GET /hoomis/{id}/events` : Lister les événements d'un Hoomi.
- `GET /events/{id}` : Récupérer les détails d'un événement.
- `PUT /events/{id}` : Mettre à jour un événement.
- `DELETE /events/{id}` : Supprimer un événement.
- `POST /events/{id}/invite` : Inviter des membres à un événement.
- `POST /events/{id}/responses` : Répondre à une invitation.

#### Modèle de Données

```sql
CREATE TABLE events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    hoomi_id UUID NOT NULL REFERENCES hoomis(id) ON DELETE CASCADE,
    creator_id UUID NOT NULL REFERENCES users(id) ON DELETE SET NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ,
    timezone VARCHAR(50) NOT NULL DEFAULT 'UTC',
    location_name VARCHAR(255),
    location_address TEXT,
    location_latitude DECIMAL(10, 8),
    location_longitude DECIMAL(11, 8),
    is_all_day BOOLEAN NOT NULL DEFAULT false,
    recurrence_rule TEXT, -- iCalendar RRULE format
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE event_participants (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    event_id UUID NOT NULL REFERENCES events(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role VARCHAR(20) NOT NULL DEFAULT 'participant', -- organizer, participant, guest
    status VARCHAR(20) NOT NULL DEFAULT 'invited', -- invited, accepted, declined, maybe
    response_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    
    UNIQUE(event_id, user_id)
);

CREATE INDEX idx_events_hoomi_time ON events(hoomi_id, start_time);
CREATE INDEX idx_event_participants_event ON event_participants(event_id);
CREATE INDEX idx_event_participants_user ON event_participants(user_id);
```

### 2. Service de Calendrier

#### Responsabilités
- Génération des vues calendrier (mois, semaine, jour).
- Calcul des événements récurrents.
- Export/Import de fichiers iCalendar (.ics).

#### Fonctionnalités
- **Vues multiples** : Mois, semaine, jour, agenda.
- **Glisser-déposer** : Pour déplacer/redimensionner les événements.
- **Superposition** : Affichage des événements de plusieurs Hoomis.
- **Filtrage** : Par type d'événement, par participant.

### 3. Service de Notifications

#### Responsabilités
- Envoi de rappels avant les événements.
- Notifications en temps réel des mises à jour.
- Intégration avec les systèmes de notification native (push, email).

#### Types de Notifications
- **Rappels** : 1 jour, 1 heure, 10 minutes avant l'événement.
- **Mises à jour** : Changement de date/heure, ajout de détails.
- **Invitations** : Nouvel événement, modification d'un événement existant.
- **Réponses** : Mise à jour du statut de participation des membres.

#### Configuration
Chaque utilisateur peut configurer ses préférences de notification :
- **Canal** : Push, email, SMS (futur).
- **Anticipation** : Délai des rappels.
- **Types** : Quels types de notifications recevoir.

### 4. Service d'Intégration Calendaire

#### Responsabilités
- Synchronisation avec les calendriers externes (Google Calendar, Apple Calendar, Outlook).
- Génération de liens d'abonnement iCalendar.
- Gestion des permissions d'accès aux calendriers externes.

#### Fonctionnalités
- **Synchronisation bidirectionnelle** : Mise à jour des deux côtés.
- **Filtres** : Synchroniser uniquement certains types d'événements.
- **Sécurité** : Authentification OAuth2 avec portée limitée.

## Intégration avec les Autres Systèmes

### Messagerie

- **Création rapide** : Bouton pour créer un événement à partir d'une conversation.
- **Discussion dédiée** : Canal de discussion spécifique à chaque événement.
- **Partages** : Liens vers l'événement dans les messages.

### Stockage de Fichiers

- **Pièces jointes** : Documents, images, plans liés à l'événement.
- **Invitations** : Création et envoi d'invitations graphiques.
- **Souvenirs** : Album photo automatique pour les événements passés.

### Jeux

- **Événements de jeu** : Organisation de soirées jeux spécifiques.
- **Tournois** : Création de tournois liés à des événements.

## Interface Utilisateur

### Web (Next.js)

#### Composants Principaux
- **CalendarView** : Vue principale du calendrier avec navigation.
- **EventModal** : Formulaire pour créer/éditer un événement.
- **EventDetails** : Page de détail d'un événement.
- **EventParticipants** : Gestion des participants et des réponses.

#### Navigation
- **Sidebar** : Accès rapide aux Hoomis et aux vues calendrier.
- **Toolbar** : Outils de vue, filtres, création rapide.
- **Responsive** : Adaptation aux écrans mobiles.

### Mobile (iOS/Android)

#### Composants Natifs
- **CalendarScreen** : Vue calendrier optimisée pour mobile.
- **EventFormScreen** : Formulaire tactile pour créer/éditer.
- **EventDetailScreen** : Détail de l'événement avec carte, participants, discussion.
- **Notifications** : Intégration avec les notifications système.

#### Gestes
- **Swipe** : Pour accepter/refuser rapidement une invitation.
- **Long Press** : Pour créer un événement à une date spécifique.
- **Pinch** : Pour zoomer dans les vues calendrier.

## Sécurité

### Contrôle d'Accès

- **Appartenance** : Seuls les membres du Hoomi peuvent voir les événements.
- **Rôles** : 
  - **Organisateur** : Peut modifier/supprimer l'événement, gérer les participants.
  - **Participant** : Peut voir les détails, modifier sa réponse.
  - **Invité** : Accès limité via un lien spécial (futur).

### Chiffrement

- **Données sensibles** : Chiffrement des lieux précis, descriptions détaillées.
- **Communication** : Toutes les communications via HTTPS/WSS.

### Confidentialité

- **Partage restreint** : Impossible de partager un événement en dehors du Hoomi.
- **Export contrôlé** : Les exports iCalendar ne contiennent que les informations nécessaires.

## Performance

### Optimisations

- **Pagination** : Chargement des événements par périodes.
- **Mise en cache** : Caching des événements fréquemment consultés.
- **Indexation** : Index de base de données optimisés pour les requêtes temporelles.

### Évolutivité

- **Partitionnement** : Possibilité de partitionner les événements par Hoomi ou date.
- **Chargement différé** : Chargement des détails uniquement quand nécessaire.

## Tests

### Tests Unitaires

- Validation des règles métiers (dates, récurrence).
- Test des calculs de rappels.
- Validation des permissions.

### Tests d'Intégration

- Création d'événements avec différents scénarios.
- Synchronisation calendrier.
- Notifications et rappels.

### Tests de Charge

- Simulation de plusieurs milliers d'événements.
- Test des performances de l'interface avec beaucoup de données.

## Surveillance

### Métriques

- **Nombre d'événements créés** par jour.
- **Taux de réponse** aux invitations.
- **Utilisation des rappels**.
- **Erreurs de synchronisation** calendrier.

### Alertes

- **Échecs de rappels**.
- **Problèmes de synchronisation**.
- **Pic d'utilisation** anormal.

Cette architecture modulaire permet de gérer efficacement les événements familiaux tout en assurant une intégration fluide avec les autres composants d'Hoomi.