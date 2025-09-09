# Architecture du Système Multimédia Hoomi

Je documente ici l'architecture que j'ai conçue pour le système multimédia d'Hoomi.

## Vue d'Ensemble

Le système multimédia d'Hoomi est conçu pour être :
- **Sécurisé** : Tous les médias sont chiffrés de bout en bout.
- **Performant** : Optimisé pour des temps de chargement rapides et une navigation fluide.
- **Évolutif** : Capable de gérer un grand nombre de médias et d'utilisateurs.
- **Intuitif** : Interface utilisateur simple et agréable à utiliser.
- **Intégré** : Étroitement lié aux autres composants d'Hoomi (messagerie, événements, jeux).

## Architecture Générale

### Structure en Couches

```
┌─────────────────────────────────────┐
│         UI / Interface              │
├─────────────────────────────────────┤
│      Logique Multimédia           │
├─────────────────────────────────────┤
│      Traitement des Médias          │
├─────────────────────────────────────┤
│      API des Médias                │
├─────────────────────────────────────┤
│      Stockage et Distribution       │
└─────────────────────────────────────┘
```

## Composants Principaux

### 1. Service de Gestion des Médias

#### Responsabilités
- Upload, téléchargement et suppression des médias.
- Génération de métadonnées.
- Gestion des albums et des organisations.
- Contrôle d'accès basé sur l'appartenance au Hoomi.

#### Endpoints API Principaux
- `POST /hoomis/{id}/media` : Upload d'un média.
- `GET /hoomis/{id}/media` : Liste des médias d'un Hoomi.
- `GET /media/{id}` : Récupérer les détails d'un média.
- `DELETE /media/{id}` : Supprimer un média.
- `POST /media/{id}/share` : Partager un média avec d'autres membres.
- `GET /albums` : Liste des albums.
- `POST /albums` : Créer un album.
- `PUT /albums/{id}` : Mettre à jour un album.

#### Modèle de Données

```sql
CREATE TABLE media (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    hoomi_id UUID NOT NULL REFERENCES hoomis(id) ON DELETE CASCADE,
    uploader_id UUID NOT NULL REFERENCES users(id) ON DELETE SET NULL,
    filename VARCHAR(255) NOT NULL,
    filesize BIGINT NOT NULL,
    mimetype VARCHAR(100) NOT NULL,
    storage_key VARCHAR(512) NOT NULL, -- Clé dans le stockage cloud
    thumbnail_key VARCHAR(512), -- Clé de la miniature
    metadata JSONB, -- Métadonnées techniques et de chiffrement
    caption TEXT, -- Légende du média
    uploaded_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE albums (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    hoomi_id UUID NOT NULL REFERENCES hoomis(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    cover_media_id UUID REFERENCES media(id) ON DELETE SET NULL,
    created_by UUID NOT NULL REFERENCES users(id) ON DELETE SET NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE album_media (
    album_id UUID NOT NULL REFERENCES albums(id) ON DELETE CASCADE,
    media_id UUID NOT NULL REFERENCES media(id) ON DELETE CASCADE,
    added_by UUID NOT NULL REFERENCES users(id) ON DELETE SET NULL,
    added_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    
    PRIMARY KEY (album_id, media_id)
);

CREATE INDEX idx_media_hoomi_uploaded ON media(hoomi_id, uploaded_at);
CREATE INDEX idx_albums_hoomi ON albums(hoomi_id);
CREATE INDEX idx_album_media_album ON album_media(album_id);
```

### 2. Service de Traitement des Médias

#### Responsabilités
- Génération de miniatures et d'aperçus.
- Transcodage de vidéos.
- Extraction de métadonnées.
- Compression d'images.

#### Technologies
- **ImageMagick** : Pour le traitement d'images.
- **FFmpeg** : Pour le traitement de vidéos/audio.
- **ExifTool** : Pour l'extraction de métadonnées EXIF.

#### Processus de Traitement

1. **Réception du fichier** : Le fichier est reçu via l'API d'upload.
2. **Validation initiale** : Vérification du type MIME et de la taille.
3. **Mise en file d'attente** : Le traitement est mis en file d'attente via **Redis**.
4. **Traitement asynchrone** :
   - **Images** : Génération de miniatures (100x100, 600x600).
   - **Vidéos** : Génération d'une vignette, transcodage en format web-friendly.
   - **Documents** : Génération d'une vignette de la première page (pour PDF).
5. **Stockage des versions** : Les versions traitées sont stockées séparément.
6. **Mise à jour des métadonnées** : Les clés de stockage des versions sont enregistrées.

### 3. Service de Stockage et Distribution

#### Stockage Cloud

##### Fournisseur
- **Google Cloud Storage** : Choisi pour son niveau gratuit généreux et sa fiabilité.

##### Structure
```
hoomi-media/
├── originals/              # Fichiers originaux chiffrés
│   └── {hoomi_id}/
│       └── {year}/{month}/
│           └── {media_id}.enc
├── thumbnails/             # Miniatures chiffrées
│   └── {hoomi_id}/
│       └── {media_id}_thumb.enc
├── previews/              # Aperçus chiffrés
│   └── {hoomi_id}/
│       └── {media_id}_preview.enc
└── temp/                  # Fichiers temporaires pour le traitement
    └── {processing_id}/
```

#### Distribution de Contenu (CDN)

##### Mise en Cache
- Utilisation d'un CDN (comme Cloudflare) pour distribuer les fichiers de manière performante.
- Les fichiers sont servis via des **liens pré-signés à courte durée de vie**.

##### Sécurité
- **Liens pré-signés** : Chaque lien est généré dynamiquement avec une expiration.
- **Vérification d'accès** : Le service vérifie les permissions avant de générer un lien.

### 4. Service de Chiffrement

#### Responsabilités
- Chiffrement des fichiers avant le stockage.
- Déchiffrement des fichiers pour l'affichage.
- Gestion des clés de chiffrement.

#### Processus

1. **Avant l'upload** :
   - Le fichier est chiffré avec une clé symétrique AES-256.
   - La clé symétrique est chiffrée avec la clé publique de chaque membre du Hoomi.
   - Les données chiffrées et les clés chiffrées sont envoyées au serveur.

2. **Pendant le stockage** :
   - Le serveur stocke uniquement le fichier chiffré.
   - Les clés chiffrées sont stockées dans la base de données `media.metadata`.

3. **Au téléchargement** :
   - Le client récupère le fichier chiffré.
   - Le client récupère sa clé chiffrée.
   - Le client déchiffre la clé avec sa clé privée.
   - Le client déchiffre le fichier avec la clé symétrique.

## Interface Utilisateur

### Web (Next.js)

#### Galerie
- **Grille responsive** : Adaptation automatique au format d'écran.
- **Infinite scroll** : Chargement progressif des médias.
- **Sélection multiple** : Pour le partage ou la suppression.
- **Filtres** : Par date, type de média, album.

#### Visionneuse
- **Navigation au clavier** : Flèches gauche/droite pour naviguer.
- **Zoom** : Pour les images.
- **Lecteur vidéo** : Intégré avec contrôles.
- **Partage** : Boutons pour partager dans la messagerie ou créer un événement.

#### Albums
- **Création** : Formulaire pour créer un nouvel album.
- **Gestion** : Ajout/retrait de médias, modification de la couverture.
- **Partage** : Possibilité de partager un album entier.

### Mobile (iOS/Android)

#### Galerie
- **Swipe** : Glisser pour naviguer entre les médias.
- **Appui long** : Pour sélectionner plusieurs médias.
- **Menu contextuel** : Pour les actions rapides (partager, supprimer).

#### Visionneuse
- **Pincement** : Pour zoomer dans les images.
- **Rotation** : Pour les vidéos en mode paysage.
- **Gestes** : Pour fermer la visionneuse.

#### Albums
- **Création rapide** : Depuis la galerie, possibilité de créer un album avec les médias sélectionnés.

## Intégration avec les Autres Systèmes

### Messagerie

#### Partage Instantané
- **Bouton "Photo"** : Dans la barre de message pour prendre ou choisir une photo.
- **Glisser-déposer** : Possibilité de glisser une image depuis la galerie vers la zone de message.
- **Aperçu dans le chat** : Les images partagées apparaissent comme des miniatures cliquables.

#### Médias dans les Messages
- **Cartes multimédias** : Pour les vidéos et documents partagés.
- **Lecture intégrée** : Pour les fichiers audio.

### Événements

#### Médias d'Événement
- **Album automatique** : Création d'un album pour chaque événement.
- **Ajout automatique** : Les photos prises pendant un événement (via l'app) sont ajoutées automatiquement.

#### Invitations Illustrées
- **Création d'invitations** : À partir de templates avec des photos de la famille.
- **Partage** : Via la messagerie ou par lien.

### Jeux

#### Captures d'Écran
- **Sauvegarde automatique** : Captures d'écran des parties gagnantes.
- **Partage** : Possibilité de partager les captures dans la messagerie.

### Profil Utilisateur

#### Avatar et Photos de Profil
- **Gestion** : Possibilité de choisir une photo parmi les médias du Hoomi.
- **Historique** : Suivi des changements d'avatar.

## Sécurité

### Chiffrement de Bout en Bout

#### Implémentation
- **Chiffrement côté client** : Tous les fichiers sont chiffrés avant d'être envoyés.
- **Clés chiffrées** : Chaque membre du Hoomi reçoit une version chiffrée de la clé de chiffrement.
- **Stockage chiffré** : Le serveur ne voit jamais les fichiers en clair.

#### Algorithmes
- **AES-256-GCM** : Pour le chiffrement des fichiers.
- **RSA-4096** : Pour le chiffrement des clés symétriques.

### Contrôle d'Accès

#### Vérification
- **À chaque requête** : Vérification que l'utilisateur est membre du Hoomi propriétaire du média.
- **Liens temporaires** : Les liens de téléchargement sont à courte durée de vie.

### Intégrité

#### Vérification
- **Hash SHA-256** : Calculé avant le chiffrement et stocké pour vérifier l'intégrité.
- **Tags d'authentification** : Utilisés par AES-GCM.

## Performance

### Optimisations

#### Mise en Cache
- **Miniatures** : Mise en cache agressive des miniatures.
- **Métadonnées** : Caching des métadonnées fréquemment demandées.

#### Compression
- **Images** : Compression avec perte contrôlée.
- **Vidéos** : Transcodage vers des codecs web-optimisés.

#### Chargement Progressif
- **Low Quality Image Placeholders (LQIP)** : Chargement d'abord d'une miniature floue, puis de la haute résolution.
- **Lazy Loading** : Chargement des médias uniquement quand ils entrent dans le viewport.

### Évolutivité

#### Traitement Asynchrone
- **Files d'attente** : Utilisation de Redis pour gérer les files d'attente de traitement.
- **Workers** : Plusieurs workers peuvent traiter les médias en parallèle.

#### Partitionnement
- **Par Hoomi** : Organisation du stockage par Hoomi pour faciliter la gestion.
- **Par date** : Organisation par année/mois pour faciliter l'archivage.

## Surveillance

### Métriques

- **Nombre de médias uploadés/téléchargés** par heure.
- **Temps de traitement moyen** des médias.
- **Taux d'erreur** dans le traitement.
- **Utilisation du stockage** par Hoomi.

### Alertes

- **Échecs de traitement** répétés.
- **Seuil d'utilisation de stockage** approché.
- **Erreurs de chiffrement/déchiffrement**.

Cette architecture modulaire et sécurisée permet de gérer efficacement les médias familiaux tout en assurant une intégration fluide avec les autres composants d'Hoomi.