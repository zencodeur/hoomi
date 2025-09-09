# Traitement des Médias dans Hoomi

Je documente ici comment le système de traitement des médias fonctionne dans Hoomi.

## Vue d'Ensemble

Le traitement des médias dans Hoomi est conçu pour :
- **Optimiser** les fichiers pour un affichage rapide et une utilisation efficace de la bande passante.
- **Extraire** les métadonnées utiles pour l'organisation et la recherche.
- **Générer** des versions adaptées aux différents usages (miniatures, aperçus, formats web).
- **Sécuriser** le processus pour protéger la vie privée des utilisateurs.

## Processus de Traitement

### 1. Réception du Média

#### Upload Initial
1. **Validation côté client** :
   - Vérification du type MIME.
   - Vérification de la taille maximale autorisée.
   - Pré-chiffrement du fichier (selon les paramètres de l'utilisateur).
2. **Envoi au serveur** :
   - Le fichier chiffré est envoyé via une requête `POST multipart/form-data`.
   - Un token de progression peut être utilisé pour suivre l'upload.

#### Validation Côté Serveur
1. **Vérification de l'authentification** :
   - Le token JWT est validé.
   - L'appartenance de l'utilisateur au Hoomi cible est vérifiée.
2. **Validation du fichier** :
   - Double vérification du type MIME (ne pas faire confiance uniquement au client).
   - Vérification de la taille.
3. **Stockage temporaire** :
   - Le fichier est stocké temporairement dans un bucket dédié avant traitement.

### 2. Mise en File d'Attente

#### Système de Files d'Attente
- **Redis** est utilisé pour gérer les files d'attente de traitement.
- Chaque type de média peut avoir sa propre file d'attente (images, vidéos, documents).

#### Priorisation
- **Médias de la messagerie** : Priorité plus élevée pour une expérience de conversation fluide.
- **Médias des événements** : Priorité normale.
- **Batch uploads** : Priorité plus basse pour ne pas bloquer l'interface.

### 3. Traitement Asynchrone

#### Workers de Traitement
- Des **workers dédiés** consomment les files d'attente.
- Ces workers peuvent être **scalables** selon la charge.
- Chaque worker est spécialisé dans un type de média.

#### Environnement de Traitement
- **Conteneurs Docker** : Pour isoler les processus de traitement.
- **Ressources limitées** : Pour éviter qu'un fichier volumineux ne bloque tout le système.
- **Timeout** : Chaque traitement a un timeout maximal.

## Traitement par Type de Média

### Images

#### Formats Supportés
- **JPEG/JPG**
- **PNG**
- **GIF** (animés et statiques)
- **WebP**
- **HEIC** (converti en JPG)

#### Processus de Traitement

1. **Extraction de Métadonnées** :
   - Utilisation d'**ExifTool** pour extraire :
     - Dimensions
     - Résolution
     - Orientation
     - Données EXIF (date, appareil photo)
     - Profil couleur

2. **Génération de Miniatures** :
   - **Carrée (100x100)** : Pour les grilles de galerie.
   - **Adaptée (max 600px)** : Pour l'affichage dans les conversations.
   - **Grande (max 1920px)** : Version optimisée pour l'affichage plein écran.

3. **Optimisation** :
   - **Compression avec perte** : Pour réduire la taille sans perte de qualité perceptible.
   - **Conversion de couleur** : sRGB pour la compatibilité web.
   - **Suppression des métadonnées sensibles** : Suppression optionnelle des données EXIF de localisation.

4. **Formats de Sortie** :
   - **Miniatures** : WebP (avec fallback JPEG).
   - **Aperçus** : WebP ou JPEG selon le navigateur.
   - **Versions optimisées** : WebP ou JPEG.

#### Algorithmes d'Optimisation

- **Redimensionnement intelligent** :
  - **Lanczos** : Pour une qualité maximale.
  - **Conservation de l'aspect ratio** : Avec padding ou recadrage intelligent.
- **Qualité adaptative** :
  - Ajustement dynamique de la qualité en fonction de la complexité de l'image.

### Vidéos

#### Formats Supportés
- **MP4 (H.264)**
- **MOV**
- **AVI**
- **WebM**

#### Processus de Traitement

1. **Analyse Initiale** :
   - Utilisation de **ffprobe** pour extraire :
     - Durée
     - Dimensions
     - Codec vidéo/audio
     - Débit binaire
     - Nombre d'images par seconde

2. **Génération de Vignettes** :
   - **Capture d'image** : À 25% de la durée de la vidéo.
   - **Plusieurs tailles** : Carrée (100x100), Adaptée (max 600px).

3. **Transcodage** :
   - **Format cible** : MP4 avec codec H.264 pour une compatibilité maximale.
   - **Paramètres** :
     - Résolution maximale : 1080p (Full HD) pour l'affichage web.
     - Débit adaptatif selon la complexité.
     - Codec audio : AAC.
   - **Qualité** : Préréglages équilibrant taille de fichier et qualité.

4. **Streaming Adaptatif** :
   - **(Futur)** : Génération de plusieurs qualités pour le streaming HLS.

#### Optimisations Vidéo

- **Compression efficace** :
  - Utilisation de **CRF (Constant Rate Factor)** pour maintenir une qualité visuelle constante.
- **Accélération matérielle** :
  - Utilisation de l'accélération GPU si disponible (NVENC pour NVIDIA).

### Documents

#### Formats Supportés
- **PDF**
- **DOC/DOCX**
- **XLS/XLSX**
- **PPT/PPTX**
- **TXT**
- **RTF**

#### Processus de Traitement

1. **Analyse du Document** :
   - **Nombre de pages** (pour PDF).
   - **Méta-informations** : Titre, auteur, date de création.

2. **Génération de Vignettes** :
   - **Première page** : Pour les PDF.
   - **Icône représentative** : Pour les documents Office.

3. **Conversion (Optionnelle)** :
   - **(Futur)** : Conversion en PDF pour uniformiser l'affichage.
   - **OCR** : (Futur) Reconnaissance optique de caractères pour les PDF scannés.

4. **Extraction de Texte** :
   - **(Futur)** : Pour la recherche dans les documents (si déchiffrable).

#### Sécurité pour les Documents

- **Aperçus sécurisés** :
  - Les documents ne sont jamais déchiffrés côté serveur.
  - Les aperçus sont générés à partir de la version chiffrée, limitant les informations extraites.

## Sécurité du Traitement

### Environnement Isolé

#### Conteneurs
- Chaque traitement s'exécute dans un **conteneur Docker** éphémère.
- **Permissions minimales** : Le conteneur n'a accès qu'aux fichiers nécessaires.
- **Destruction** : Le conteneur est détruit après le traitement.

#### Fichiers Temporaires
- **Nettoyage automatique** : Les fichiers temporaires sont supprimés après traitement.
- **Chemin aléatoire** : Utilisation de chemins aléatoires et non prédictibles.

### Chiffrement

#### Données en Transit
- **HTTPS** : Toutes les communications avec les workers de traitement sont chiffrées.
- **Volumes chiffrés** : Les volumes de stockage utilisés par les workers sont chiffrés.

#### Données au Repos
- **Fichiers originaux** : Toujours chiffrés avant d'être écrits sur le disque.
- **Versions traitées** : Également chiffrées.

## Gestion des Erreurs

### Types d'Erreurs

#### Erreurs de Traitement
- **Fichier corrompu** : Impossible à décoder.
- **Format non supporté** : Même si validé initialement, le fichier est incorrect.
- **Timeout** : Le traitement prend trop de temps.
- **Ressources insuffisantes** : Le worker manque de mémoire/CPU.

#### Erreurs Réseau
- **Perte de connexion** : Pendant le transfert vers/depuis le stockage.
- **Erreur de stockage** : Impossible d'écrire/lire dans le bucket cloud.

### Stratégies de Gestion

#### Retry Mechanism
- **Politique exponentielle** : Tentatives avec délais croissants (1s, 2s, 4s, 8s).
- **Maximum de tentatives** : 3 tentatives par défaut.

#### Dead Letter Queue
- **Échecs persistants** : Après 3 échecs, le job est déplacé dans une file d'attente d'erreurs.
- **Alerte** : Notification aux administrateurs pour intervention manuelle.

#### Compensation
- **Nettoyage** : Si une version traitée ne peut être sauvegardée, les versions partielles sont nettoyées.
- **Retour à l'état précédent** : Si possible, restauration de l'état avant le traitement.

## Surveillance et Métriques

### Suivi des Traitements

#### Logs
- **Début/Fin** : Timestamp du début et de la fin de chaque traitement.
- **Étapes** : Journalisation des étapes importantes (extraction, redimensionnement, encodage).
- **Erreurs** : Détails complets en cas d'erreur.

#### Métriques
- **Temps de traitement** : Par type de média.
- **Taux de succès** : Pour chaque type de traitement.
- **Utilisation des ressources** : CPU, mémoire, disque par worker.

### Dashboards

#### Pour les Administrateurs
- **Files d'attente** : Nombre de jobs en attente, en traitement, en erreur.
- **Performance** : Temps moyen de traitement par type.
- **Erreurs** : Types d'erreurs les plus fréquents.

#### Pour les Utilisateurs
- **Progression** : Pour les uploads en cours (si applicable).
- **Statut** : Indication si le traitement est en cours ou terminé.

## Évolutivité

### Scaling Horizontal

#### Workers
- **Ajout dynamique** : Possibilité d'ajouter/supprimer des workers selon la charge.
- **Auto-scaling** : (Futur) Basé sur la longueur des files d'attente.

#### Load Balancing
- **Distribution équitable** : Des jobs entre les workers disponibles.

### Partitionnement

#### Par Type de Média
- **Files d'attente dédiées** : Pour gérer les priorités et les spécialisations.
- **Workers spécialisés** : Pour optimiser les performances par type.

Cette architecture de traitement robuste et sécurisée permet de gérer efficacement tous les types de médias tout en maintenant une excellente expérience utilisateur.