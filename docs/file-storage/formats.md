# Formats et Types de Fichiers dans Hoomi

Je documente ici les formats et types de fichiers que je supporte dans Hoomi, ainsi que leur gestion spécifique.

## Vue d'Ensemble

Hoomi supporte différents types de fichiers pour répondre aux besoins variés des familles :
- **Images** : Pour partager des photos de moments précieux.
- **Vidéos** : Pour capturer des événements en mouvement.
- **Documents** : Pour partager des fichiers importants (recettes, documents familiaux, etc.).
- **Audio** : Pour des messages vocaux ou des enregistrements audio.

## Types de Fichiers Supportés

### Images

**Formats Supportés :**
- JPEG/JPG
- PNG
- GIF (animés et statiques)
- WebP
- HEIC (sur iOS, converti en JPG pour la compatibilité)

**Gestion Spécifique :**
- Génération automatique de miniatures.
- Optimisation de la taille pour l'affichage dans les conversations.
- Préservation de l'orientation EXIF.
- Compression avec perte pour réduire la taille tout en maintenant une qualité acceptable.

### Vidéos

**Formats Supportés :**
- MP4 (H.264)
- MOV
- AVI
- WebM

**Gestion Spécifique :**
- Génération d'aperçus (thumbnails).
- Transcodage si nécessaire pour la compatibilité.
- Limitation de la durée (par exemple, max 10 minutes pour l'instant).
- Compression pour réduire la taille de stockage et de transfert.

### Documents

**Formats Supportés :**
- PDF
- DOC/DOCX
- XLS/XLSX
- PPT/PPTX
- TXT
- RTF

**Gestion Spécifique :**
- Génération d'aperçus pour les PDF (première page).
- Pour les documents Office, extraction du texte pour une recherche future (si déchiffrable).
- Limitation de la taille (par exemple, max 50MB).

### Audio

**Formats Supportés :**
- MP3
- WAV
- M4A
- FLAC

**Gestion Spécifique :**
- Conversion vers un format standard (MP3) si nécessaire.
- Limitation de la durée (par exemple, max 5 minutes).
- Compression pour réduire la taille.

## Limitations et Restrictions

### Taille Maximale

- **Par fichier** : 100MB (cette limite pourra être augmentée dans les offres payantes).
- **Par utilisateur/jour** : 1GB (pour prévenir les abus).

### Contenu Interdit

Le système inclut un filtre basique pour détecter :
- Des types de fichiers explicitement interdits (exécutables).
- Du contenu potentiellement inapproprié (système d'IA basique).

### Validation

- Validation du type MIME côté serveur en plus de la vérification côté client.
- Vérification de l'intégrité des fichiers après l'upload.

## Optimisation

### Compression

- **Images** : Utilisation de techniques de compression adaptées.
- **Vidéos** : Transcodage avec des paramètres optimisés.
- **Audio** : Conversion avec un bitrate approprié.

### Mise en Cache

- Les miniatures et aperçus sont mis en cache pour améliorer les performances.

### Déduplication

- (Fonctionnalité future) Déduplication des fichiers identiques pour économiser de l'espace.

## Métadonnées

Pour chaque type de fichier, je stocke des métadonnées spécifiques :

### Images
- Dimensions (largeur x hauteur)
- Résolution DPI
- Profil couleur
- Données EXIF (date, appareil, localisation si disponible)

### Vidéos
- Durée
- Dimensions
- Codec vidéo/audio
- Débit binaire
- Nombre d'images par seconde (FPS)

### Documents
- Nombre de pages (pour PDF)
- Titre, auteur (si disponible dans les métadonnées)
- Date de création/modification

### Audio
- Durée
- Codec
- Débit binaire
- Fréquence d'échantillonnage

## Aperçus et Miniatures

### Génération

- Générés automatiquement lors de l'upload.
- Stockés séparément pour un accès rapide.

### Tailles

- **Miniature** : 100x100 pixels (carré, recadré).
- **Aperçu** : Largeur maximale de 600 pixels, hauteur proportionnelle.

### Format

- JPEG pour les images et aperçus.
- Pour les vidéos, l'aperçu est une image JPEG de la première frame.

## Intégration avec E2EE

Tous les fichiers sont chiffrés avant la génération des aperçus, donc :
- Les aperçus sont générés à partir du fichier chiffré.
- Cela limite les types d'aperçus possibles (pas d'analyse de contenu).
- Pour les PDF, seul le nom du fichier et le nombre de pages peuvent être stockés en clair.

## Traitement Asynchrone

Certaines opérations (comme la génération d'aperçus vidéo) sont effectuées de manière asynchrone :
- Le fichier est d'abord stocké.
- Une tâche de fond traite le fichier pour générer les aperçus.
- L'utilisateur est notifié quand le traitement est terminé.

## Support Multi-plateforme

La gestion des fichiers est conçue pour fonctionner de manière cohérente sur :
- Web (navigateurs modernes)
- iOS (Swift)
- Android (Kotlin)

Les bibliothèques d'encodage/decodage spécifiques à chaque plateforme sont utilisées pour garantir la compatibilité.

## Évolution Future

- **Support de nouveaux formats** : Ajout de formats spécifiques (par exemple, fichiers de projets créatifs).
- **Amélioration de l'IA** : Pour une meilleure détection de contenu et génération d'aperçus.
- **Compression intelligente** : Adaptation dynamique de la qualité en fonction de la connexion.