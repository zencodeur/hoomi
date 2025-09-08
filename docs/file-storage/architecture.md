# Architecture du Stockage de Fichiers Hoomi

Je documente ici l'architecture du système de stockage de fichiers que j'ai conçue pour Hoomi.

## Vue d'Ensemble

Le système de stockage de fichiers de Hoomi est conçu pour être :
- **Sécurisé** : Avec un chiffrement de bout en bout intégré.
- **Évolutif** : Pour gérer un grand nombre de fichiers et d'utilisateurs.
- **Fiable** : Pour garantir la disponibilité des fichiers.
- **Performant** : Pour des temps de téléchargement et d'upload rapides.

## Composants Principaux

### Service de Médias (Go)

Responsable de :
- Gérer les requêtes d'upload et de téléchargement de fichiers.
- Coordonner le chiffrement E2EE des fichiers.
- Interagir avec le stockage cloud.
- Mettre à jour la base de données avec les métadonnées.

### Stockage Cloud

J'utilise un service de stockage cloud (comme AWS S3, Google Cloud Storage) pour stocker les fichiers de manière fiable et évolutive.

### Base de Données (PostgreSQL)

Stockage des métadonnées des fichiers dans la table `media`, y compris :
- Les chemins vers les fichiers dans le stockage cloud.
- Les informations de chiffrement.
- Les métadonnées techniques (taille, type MIME, dimensions).

## Flux de Données

### Upload d'un Fichier

1. **Client** : L'utilisateur sélectionne un fichier à uploader.
2. **Client** : Le fichier est chiffré avec une clé symétrique AES-256.
3. **Client** : La clé de chiffrement est chiffrée avec la clé publique de chaque membre du Hoomi.
4. **Client** : Le fichier chiffré est envoyé via une requête `POST multipart/form-data` à l'API (`POST /hoomis/{id}/media`).
5. **API** : Le service de médias reçoit le fichier.
6. **API** : Le fichier est validé (taille, type MIME).
7. **API** : Le fichier chiffré est uploadé vers le stockage cloud.
8. **API** : Les métadonnées (chemin, taille, type MIME, clés chiffrées) sont stockées dans PostgreSQL.
9. **API** : Une notification est envoyée aux membres connectés via WebSocket.

### Téléchargement d'un Fichier

1. **Client** : L'utilisateur clique sur un fichier pour le télécharger.
2. **Client** : Une requête `GET` est envoyée à l'API (`GET /media/{id}`).
3. **API** : Le service de médias récupère les métadonnées depuis PostgreSQL.
4. **API** : Le fichier chiffré est récupéré depuis le stockage cloud.
5. **API** : Le fichier chiffré est envoyé dans la réponse HTTP.
6. **Client** : Reçoit le fichier chiffré.
7. **Client** : Déchiffre le fichier avec la clé de session E2EE.
8. **Client** : Affiche ou sauvegarde le fichier déchiffré.

## Sécurité

### Chiffrement de Bout en Bout

- **Chiffrement côté client** : Le fichier est chiffré avant d'être envoyé.
- **Stockage chiffré** : Le serveur ne voit jamais le fichier en clair.
- **Clés chiffrées** : Chaque membre du Hoomi reçoit une version chiffrée de la clé de chiffrement.

### Accès Contrôlé

- L'accès aux fichiers est contrôlé par le service d'authentification/autorisation.
- Seuls les membres d'un Hoomi peuvent accéder à ses fichiers.

## Optimisation

### Mise en Cache

- Les miniatures et les aperçus peuvent être mis en cache pour améliorer les performances.

### Compression

- Pour certains types de fichiers (images, textes), une compression peut être appliquée avant le chiffrement.

### Déduplication

- À l'avenir, je pourrais implémenter une déduplication pour éviter de stocker plusieurs fois le même fichier.

## Gestion des Versions

Actuellement, chaque upload crée une nouvelle entrée. Une gestion de versions plus sophistiquée pourrait être implémentée plus tard.

## Sauvegarde et Récupération

### Sauvegardes Cloud

- Le stockage cloud offre généralement des mécanismes de sauvegarde intégrés.

### Récupération après Sinistre

- Possibilité de restaurer les fichiers à partir des sauvegardes du fournisseur cloud.

## Scalabilité

### Horizontal Scaling

- Le stockage cloud s'adapte automatiquement à la demande.
- Plusieurs instances du service de médias peuvent être déployées.

### Limites

- Des quotas peuvent s'appliquer au niveau du fournisseur cloud.
- La bande passante peut être un facteur limitant.

## Surveillance

Je prévois d'implémenter des métriques pour surveiller :
- Le nombre de fichiers uploadés/téléchargés par seconde.
- L'utilisation du stockage.
- Les taux d'erreur.
- La latence des opérations.