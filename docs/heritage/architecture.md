# Architecture du Système d'Héritage Numérique Hoomi

Je documente ici l'architecture que j'ai conçue pour le système d'héritage numérique d'Hoomi.

## Vue d'Ensemble

L'architecture de l'héritage numérique d'Hoomi est conçue pour être :
- **Sécurisée** : Toutes les données d'héritage sont chiffrées de bout en bout.
- **Durable** : Conçue pour durer dans le temps, assurant la pérennité de l'héritage.
- **Modulaire** : Composants distincts pour l'arbre généalogique, les capsules, la transmission.
- **Intégrée** : Étroitement liée aux autres systèmes d'Hoomi (messagerie, médias, événements).
- **Évolutives** : Capable de s'adapter aux besoins futurs et à la croissance.

## Architecture Générale

### Structure en Couches

```
┌─────────────────────────────────────┐
│         UI / Interface              │
├─────────────────────────────────────┤
│     Logique d'Héritage Numérique    │
├─────────────────────────────────────┤
│      Services Spécialisés           │
│  (Arbre, Capsules, Transmission)    │
├─────────────────────────────────────┤
│        API d'Héritage               │
├─────────────────────────────────────┤
│      Stockage et Persistance        │
└─────────────────────────────────────┘
```

## Composants Principaux

### 1. Service de Gestion Généalogique

#### Responsabilités
- Création, modification, suppression de personnes dans l'arbre.
- Gestion des relations familiales (parenté, mariage, adoption).
- Stockage des informations biographiques et multimédias.
- Génération de l'arbre généalogique interactif.

#### Endpoints API Principaux
- `POST /hoomis/{id}/genealogy/persons` : Ajouter une personne.
- `GET /hoomis/{id}/genealogy/persons` : Lister les personnes du Hoomi.
- `GET /hoomis/{id}/genealogy/persons/{personId}` : Récupérer les détails d'une personne.
- `PUT /hoomis/{id}/genealogy/persons/{personId}` : Mettre à jour une personne.
- `DELETE /hoomis/{id}/genealogy/persons/{personId}` : Supprimer une personne.
- `POST /hoomis/{id}/genealogy/relations` : Ajouter une relation (parent-enfant, mariage).
- `GET /hoomis/{id}/genealogy/tree` : Récupérer l'arbre généalogique complet ou partiel.

#### Modèle de Données

```sql
CREATE TABLE genealogy_persons (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    hoomi_id UUID NOT NULL REFERENCES hoomis(id) ON DELETE CASCADE,
    user_id UUID REFERENCES users(id) ON DELETE SET NULL, -- Lien vers un utilisateur Hoomi actif
    
    -- Informations de base
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100),
    birth_date DATE,
    death_date DATE,
    gender VARCHAR(20), -- male, female, other, unknown
    
    -- Biographie et détails
    biography TEXT,
    occupation VARCHAR(255),
    interests TEXT[], -- Liste d'intérêts
    
    -- Métadonnées
    created_by UUID NOT NULL REFERENCES users(id) ON DELETE SET NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    
    -- Chiffrement
    encrypted_data TEXT, -- Données sensibles chiffrées
    encryption_metadata JSONB -- Informations sur le chiffrement
);

CREATE TABLE genealogy_relations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    hoomi_id UUID NOT NULL REFERENCES hoomis(id) ON DELETE CASCADE,
    
    -- Personnes impliquées
    person1_id UUID NOT NULL REFERENCES genealogy_persons(id) ON DELETE CASCADE,
    person2_id UUID NOT NULL REFERENCES genealogy_persons(id) ON DELETE CASCADE,
    
    -- Type de relation
    relation_type VARCHAR(50) NOT NULL, -- parent, spouse, sibling, cousin, etc.
    
    -- Détails de la relation
    start_date DATE, -- Date de mariage, naissance d'enfant, etc.
    end_date DATE, -- Divorce, décès d'un conjoint, etc.
    details TEXT, -- Notes sur la relation
    
    -- Métadonnées
    created_by UUID NOT NULL REFERENCES users(id) ON DELETE SET NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    
    -- Unicité
    UNIQUE(person1_id, person2_id, relation_type)
);

CREATE TABLE genealogy_media (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    person_id UUID NOT NULL REFERENCES genealogy_persons(id) ON DELETE CASCADE,
    media_id UUID NOT NULL REFERENCES media(id) ON DELETE CASCADE,
    relationship VARCHAR(50), -- portrait, wedding, childhood, etc.
    description TEXT,
    is_featured BOOLEAN NOT NULL DEFAULT false, -- Photo principale de la personne
    
    -- Métadonnées
    added_by UUID NOT NULL REFERENCES users(id) ON DELETE SET NULL,
    added_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_genealogy_persons_hoomi ON genealogy_persons(hoomi_id);
CREATE INDEX idx_genealogy_persons_user ON genealogy_persons(user_id);
CREATE INDEX idx_genealogy_relations_person1 ON genealogy_relations(person1_id);
CREATE INDEX idx_genealogy_relations_person2 ON genealogy_relations(person2_id);
CREATE INDEX idx_genealogy_relations_hoomi ON genealogy_relations(hoomi_id);
CREATE INDEX idx_genealogy_media_person ON genealogy_media(person_id);
```

### 2. Service de Gestion des Capsules Temporelles

#### Responsabilités
- Création, programmation et gestion des capsules temporelles.
- Surveillance des conditions de déclenchement.
- Envoi des capsules aux destinataires au moment opportun.

#### Endpoints API Principaux
- `POST /hoomis/{id}/time-capsules` : Créer une capsule temporelle.
- `GET /hoomis/{id}/time-capsules` : Lister les capsules du Hoomi.
- `GET /time-capsules/{id}` : Récupérer les détails d'une capsule.
- `PUT /time-capsules/{id}` : Mettre à jour une capsule.
- `DELETE /time-capsules/{id}` : Supprimer une capsule.
- `POST /time-capsules/{id}/trigger` : Déclencher manuellement une capsule (pour test).

#### Modèle de Données

```sql
CREATE TABLE time_capsules (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    hoomi_id UUID NOT NULL REFERENCES hoomis(id) ON DELETE CASCADE,
    creator_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    
    -- Contenu
    title VARCHAR(255) NOT NULL,
    description TEXT,
    content JSONB, -- Contenu de la capsule (texte, liens vers médias, etc.)
    
    -- Programmation
    trigger_type VARCHAR(20) NOT NULL, -- 'date', 'event', 'mixed'
    scheduled_at TIMESTAMPTZ, -- Pour les déclenchements temporels
    event_type VARCHAR(50), -- Pour les déclenchements événementiels (birth, death, anniversary)
    event_person_id UUID REFERENCES genealogy_persons(id) ON DELETE SET NULL, -- Personne concernée par l'événement
    
    -- Destinataires
    recipient_ids UUID[], -- Liste des utilisateurs destinataires
    
    -- Statut
    status VARCHAR(20) NOT NULL DEFAULT 'scheduled', -- scheduled, triggered, delivered, cancelled
    triggered_at TIMESTAMPTZ, -- Date de déclenchement réel
    delivered_at TIMESTAMPTZ, -- Date de livraison effective
    
    -- Métadonnées
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    
    -- Chiffrement
    encrypted_content TEXT, -- Contenu chiffré
    encryption_metadata JSONB -- Informations sur le chiffrement
);

CREATE TABLE time_capsule_contents (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    capsule_id UUID NOT NULL REFERENCES time_capsules(id) ON DELETE CASCADE,
    content_type VARCHAR(50) NOT NULL, -- 'text', 'media', 'document'
    content_id UUID, -- Référence à un média ou document dans Hoomi
    content_text TEXT, -- Texte directement inclus (chiffré)
    description TEXT -- Description du contenu
);

CREATE INDEX idx_time_capsules_hoomi_status ON time_capsules(hoomi_id, status);
CREATE INDEX idx_time_capsules_scheduled ON time_capsules(scheduled_at) WHERE status = 'scheduled';
CREATE INDEX idx_time_capsule_contents_capsule ON time_capsule_contents(capsule_id);
```

#### Service de Surveillance

Un **worker dédié** vérifie périodiquement les conditions de déclenchement :
- **Vérification temporelle** : Comparaison de `scheduled_at` avec l'heure actuelle.
- **Vérification événementielle** : Surveillance des événements dans le Hoomi (naissance, décès).

### 3. Service de Transmission Héritage

#### Responsabilités
- Gestion des volontés numériques.
- Désignation des légataires numériques.
- Contrôle d'accès conditionnel aux contenus.
- Planification de l'accès à l'héritage.

#### Endpoints API Principaux
- `POST /hoomis/{id}/heritage/wills` : Créer une volonté numérique.
- `GET /hoomis/{id}/heritage/wills` : Lister les volontés du Hoomi.
- `GET /heritage/wills/{id}` : Récupérer une volonté.
- `PUT /heritage/wills/{id}` : Mettre à jour une volonté.
- `DELETE /heritage/wills/{id}` : Supprimer une volonté.
- `POST /heritage/assets` : Associer un actif à une volonté.
- `POST /heritage/access-rules` : Créer une règle d'accès conditionnel.

#### Modèle de Données

```sql
CREATE TABLE heritage_wills (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    hoomi_id UUID NOT NULL REFERENCES hoomis(id) ON DELETE CASCADE,
    creator_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    
    -- Informations sur la volonté
    title VARCHAR(255) NOT NULL,
    description TEXT,
    is_active BOOLEAN NOT NULL DEFAULT true,
    
    -- Métadonnées
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    activated_at TIMESTAMPTZ -- Date d'activation officielle
);

CREATE TABLE heritage_assets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    will_id UUID NOT NULL REFERENCES heritage_wills(id) ON DELETE CASCADE,
    asset_type VARCHAR(50) NOT NULL, -- 'media', 'document', 'album', 'chat_history', 'custom'
    asset_id UUID, -- Référence à l'actif dans Hoomi
    custom_description TEXT, -- Pour les actifs non liés directement (souvenirs, objets physiques)
    
    -- Instructions
    instruction TEXT -- Que faire de cet actif
);

CREATE TABLE heritage_beneficiaries (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    will_id UUID NOT NULL REFERENCES heritage_wills(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role VARCHAR(50) NOT NULL, -- 'beneficiary', 'executor', 'guardian'
    share_percentage INTEGER, -- Pourcentage de l'héritage (null pour les exécuteurs/gardiens)
    
    -- Conditions
    conditions JSONB, -- Conditions pour recevoir l'héritage (âge, mariage, etc.)
    
    UNIQUE(will_id, user_id)
);

CREATE TABLE heritage_access_rules (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    hoomi_id UUID NOT NULL REFERENCES hoomis(id) ON DELETE CASCADE,
    creator_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    
    -- Actif concerné
    asset_type VARCHAR(50) NOT NULL,
    asset_id UUID, -- Référence à l'actif
    
    -- Conditions d'accès
    condition_type VARCHAR(50) NOT NULL, -- 'age', 'date', 'event', 'password'
    condition_value JSONB, -- Valeur de la condition (ex: {"age": 18}, {"date": "2030-01-01"})
    
    -- Destinataires
    grant_to_user_ids UUID[], -- Utilisateurs qui obtiennent l'accès
    
    -- Métadonnées
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_heritage_wills_hoomi ON heritage_wills(hoomi_id);
CREATE INDEX idx_heritage_assets_will ON heritage_assets(will_id);
CREATE INDEX idx_heritage_beneficiaries_will ON heritage_beneficiaries(will_id);
CREATE INDEX idx_heritage_access_rules_hoomi ON heritage_access_rules(hoomi_id);
```

## Intégration avec les Autres Systèmes

### Messagerie

#### Notifications
- **Création d'actifs** : Notification aux bénéficiaires quand un actif leur est attribué.
- **Capsules déclenchées** : Message spécial dans la conversation du Hoomi.
- **Anniversaires** : Messages automatiques pour les anniversaires des personnes de l'arbre.

#### Contenus d'Héritage
- **Messages historiques** : Possibilité d'inclure des conversations historiques dans l'héritage.
- **Discussions dédiées** : Création de fils de discussion spécifiques pour planifier l'héritage.

### Médias

#### Médias Généalogiques
- **Photos de famille** : Intégration automatique des photos de profil et de couverture dans les profils généalogiques.
- **Albums historiques** : Création automatique d'albums par décennie ou par personne.
- **Vidéos d'histoire** : Intégration de vidéos interviews ou souvenirs dans les profils.

#### Médias dans les Capsules
- **Contenus multimédias** : Possibilité d'inclure des photos, vidéos, documents dans les capsules temporelles.

### Événements

#### Événements Généalogiques
- **Dates importantes** : Anniversaires, dates de mariage automatiquement ajoutés à l'arbre.
- **Événements de vie** : Naissances, mariages, décès mis à jour automatiquement dans l'arbre.

#### Événements d'Héritage
- **Célébrations** : Création d'événements pour célébrer les anniversaires des ancêtres.
- **Commémorations** : Événements pour commémorer les décès.

### Profils Utilisateurs

#### Liaison
- **Personne généalogique** : Chaque utilisateur peut être lié à une personne dans l'arbre généalogique.
- **Profil enrichi** : Informations du profil utilisateur intégrées dans le profil généalogique.

#### Héritage Numérique Personnel
- **Contenus personnels** : Chaque utilisateur peut désigner des contenus personnels comme faisant partie de son héritage.

## Sécurité

### Chiffrement de Bout en Bout

#### Données Sensibles
- **Biographies** : Textes détaillés sur les personnes de l'arbre.
- **Contenus des capsules** : Messages et médias dans les capsules temporelles.
- **Volontés numériques** : Détails des volontés et des répartitions.

#### Implémentation
- **AES-256-GCM** : Pour le chiffrement des contenus.
- **RSA-4096** : Pour le chiffrement des clés symétriques.
- **Clés partagées** : Chaque membre du Hoomi reçoit une version chiffrée de la clé nécessaire.

### Contrôle d'Accès

#### Permissions
- **Contributeurs à l'arbre** : Droits pour ajouter/modifier des personnes et relations.
- **Légataires** : Droits spécifiques pour accéder à certains contenus selon les volontés.
- **Exécuteurs** : Droits pour gérer l'exécution des volontés numériques.

#### Vérification
- **À chaque requête** : Vérification des permissions avant d'accéder aux données d'héritage.
- **Journalisation** : Enregistrement de toutes les actions sur les données d'héritage.

## Performance et Évolutivité

### Optimisations

#### Caching
- **Arbre généalogique** : Mise en cache des structures d'arbre fréquemment demandées.
- **Métadonnées** : Caching des métadonnées des capsules et des volontés.

#### Pagination
- **Listes longues** : Pagination des listes de personnes, de capsules, de volontés.

#### Chargement Différé
- **Détails** : Chargement des détails d'une personne, d'une capsule ou d'une volonté uniquement quand nécessaire.

### Évolutivité

#### Partitionnement
- **Par Hoomi** : Organisation des données par Hoomi pour faciliter la gestion.
- **Par date** : Pour les capsules temporelles programmées à long terme.

#### Workers Asynchrones
- **Déclenchement** : Workers dédiés pour vérifier les conditions de déclenchement des capsules.
- **Livraison** : Workers pour gérer l'envoi des capsules aux destinataires.

## Surveillance

### Métriques

- **Nombre de personnes dans l'arbre** par Hoomi.
- **Nombre de capsules créées/déclenchées**.
- **Nombre de volontés actives**.
- **Taux d'utilisation des fonctionnalités d'héritage**.

### Alertes

- **Erreurs de déclenchement** de capsules.
- **Problèmes d'accès** aux données d'héritage.
- **Seuils d'utilisation** approchés (nombre de personnes, de capsules).

Cette architecture modulaire et sécurisée permet de gérer efficacement l'héritage numérique familial tout en assurant une intégration fluide avec les autres composants d'Hoomi.