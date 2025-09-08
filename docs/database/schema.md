# Schéma de la Base de Données Hoomi

Je documente ici le schéma de la base de données PostgreSQL que j'utilise pour le backend de Hoomi.

## Entités Principales

### Users

Cette table représente un utilisateur de mon application.

**Attributs :**
- `id` (UUID) : Identifiant unique de l'utilisateur. **(Clé primaire)**
- `email` (string, unique) : Adresse e-mail de l'utilisateur.
- `password_hash` (string) : Mot de passe haché.
- `name` (string) : Nom complet de l'utilisateur.
- `created_at` (timestamp) : Date de création du compte.
- `updated_at` (timestamp) : Date de dernière mise à jour du compte.

### Hoomis

Cette table représente un groupe familial.

**Attributs :**
- `id` (UUID) : Identifiant unique du Hoomi. **(Clé primaire)**
- `name` (string) : Nom du Hoomi.
- `created_at` (timestamp) : Date de création du Hoomi.
- `updated_at` (timestamp) : Date de dernière mise à jour du Hoomi.

### Memberships

Cette table représente l'appartenance d'un utilisateur à un Hoomi.

**Attributs :**
- `id` (UUID) : Identifiant unique de l'adhésion. **(Clé primaire)**
- `user_id` (UUID) : Référence à l'utilisateur. **(Clé étrangère vers `users.id`)**
- `hoomi_id` (UUID) : Référence au Hoomi. **(Clé étrangère vers `hoomis.id`)**
- `role` (string) : Rôle de l'utilisateur dans le Hoomi (admin, member, ...).
- `joined_at` (timestamp) : Date d'adhésion au Hoomi.
- `created_at` (timestamp) : Date de création de l'adhésion.
- `updated_at` (timestamp) : Date de dernière mise à jour de l'adhésion.

**Contraintes :**
- Unicité de `(user_id, hoomi_id)` : Un utilisateur ne peut pas rejoindre le même Hoomi plusieurs fois.

### Messages

Cette table représente un message textuel échangé dans un Hoomi.

**Attributs :**
- `id` (UUID) : Identifiant unique du message. **(Clé primaire)**
- `hoomi_id` (UUID) : Référence au Hoomi dans lequel le message est posté. **(Clé étrangère vers `hoomis.id`)**
- `sender_id` (UUID) : Référence à l'utilisateur qui a envoyé le message. **(Clé étrangère vers `users.id`)**
- `content` (text) : Contenu du message (chiffré côté client avant d'être stocké).
- `created_at` (timestamp) : Date d'envoi du message.
- `updated_at` (timestamp) : Date de dernière mise à jour du message.

### Media

Cette table représente un fichier média (photo, vidéo, document) partagé dans un Hoomi.

**Attributs :**
- `id` (UUID) : Identifiant unique du média. **(Clé primaire)**
- `hoomi_id` (UUID) : Référence au Hoomi dans lequel le média est partagé. **(Clé étrangère vers `hoomis.id`)**
- `uploader_id` (UUID) : Référence à l'utilisateur qui a uploadé le média. **(Clé étrangère vers `users.id`)**
- `file_name` (string) : Nom original du fichier.
- `file_path` (string) : Chemin vers le fichier dans le stockage.
- `file_size` (integer) : Taille du fichier en octets.
- `mime_type` (string) : Type MIME du fichier.
- `metadata` (jsonb) : Métadonnées supplémentaires (dimensions pour une image, durée pour une vidéo, etc.). Ce champ peut aussi contenir des informations de chiffrement.
- `created_at` (timestamp) : Date d'upload du média.
- `updated_at` (timestamp) : Date de dernière mise à jour du média.

## Relations

- Un **User** peut appartenir à plusieurs **Hoomis** (via **Memberships**).
- Un **Hoomi** peut avoir plusieurs **Users** (via **Memberships**).
- Un **Hoomi** peut contenir plusieurs **Messages**.
- Un **User** peut envoyer plusieurs **Messages**.
- Un **Hoomi** peut contenir plusieurs **Media**.
- Un **User** peut uploader plusieurs **Media**.

## Diagramme (Textuel)

```
Users ||--o{ Memberships }o--|| Hoomis
Hoomis ||--o{ Messages
Users >--o{ Messages
Hoomis ||--o{ Media
Users >--o{ Media
```

## Considérations de Sécurité

Voici les mesures de sécurité que j'ai prévues :
- Les mots de passe ne sont jamais stockés en clair.
- Le contenu des messages est chiffré côté client avant d'être stocké.
- Les informations sensibles dans `Media.metadata` peuvent être chiffrées.
- Les accès à la base de données sont restreints et authentifiés.

## Indexes

(À définir plus précisément selon les requêtes les plus fréquentes)
- Index sur `users.email` (pour l'authentification).
- Index sur `memberships.user_id` et `memberships.hoomi_id` (pour les requêtes d'appartenance).
- Index sur `messages.hoomi_id` et `messages.created_at` (pour lister les messages d'un Hoomi).
- Index sur `media.hoomi_id` et `media.created_at` (pour lister les médias d'un Hoomi).