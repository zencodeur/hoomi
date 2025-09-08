# Stockage des Messages dans Hoomi

Je documente ici comment je gère le stockage des messages dans Hoomi, en tenant compte de l'aspect critique de la sécurité et de la performance.

## Vue d'Ensemble

Le stockage des messages dans Hoomi est conçu pour :
- **Garantir la sécurité** : Les messages sont chiffrés de bout en bout et stockés chiffrés.
- **Assurer la performance** : Récupération rapide de l'historique, même pour les conversations actives.
- **Permettre la scalabilité** : Adapter le stockage à la croissance du nombre d'utilisateurs et de messages.
- **Maintenir la fiabilité** : Ne jamais perdre de messages.

## Base de Données Principale (PostgreSQL)

### Table `messages`

```sql
CREATE TABLE messages (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    hoomi_id UUID NOT NULL REFERENCES hoomis(id) ON DELETE CASCADE,
    sender_id UUID NOT NULL REFERENCES users(id) ON DELETE SET NULL,
    content TEXT NOT NULL, -- Contenu chiffré
    nonce TEXT, -- Nonce utilisé pour le chiffrement
    session_id TEXT, -- ID de la session de chiffrement
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

### Colonnes Détail

- **`id`** : Identifiant unique du message.
- **`hoomi_id`** : Référence au Hoomi dans lequel le message a été posté.
- **`sender_id`** : Référence à l'utilisateur qui a envoyé le message. `ON DELETE SET NULL` pour préserver le message même si l'utilisateur est supprimé.
- **`content`** : Le contenu du message, chiffré de bout en bout. Stocké en texte (base64).
- **`nonce`** : Le nonce utilisé lors du chiffrement AES-GCM, nécessaire pour le déchiffrement.
- **`session_id`** : L'identifiant de la session de chiffrement utilisée.
- **`created_at`** : Horodatage de création du message.
- **`updated_at`** : Horodatage de la dernière mise à jour.

### Index

```sql
-- Pour la récupération des messages d'un Hoomi, triés par date
CREATE INDEX idx_messages_hoomi_created ON messages(hoomi_id, created_at);

-- Pour la récupération des messages d'un utilisateur
CREATE INDEX idx_messages_sender ON messages(sender_id);

-- Pour la pagination avec un curseur (before/after)
CREATE INDEX idx_messages_created ON messages(created_at);
```

## Sécurité du Stockage

### Chiffrement

- **Côté client** : Le message est chiffré avec AES-GCM avant d'être envoyé au serveur.
- **Stockage** : Le serveur reçoit et stocke uniquement le contenu chiffré.
- **Clés** : Les clés de chiffrement ne sont jamais connues du serveur.

### Séparation des Données

- Chaque Hoomi a son propre espace de messages.
- Les messages ne peuvent être accédés que par les membres du Hoomi concerné.

## Performance et Optimisation

### Pagination

Pour éviter de surcharger le client ou le serveur, l'historique est paginé :
- La requête API accepte des paramètres `limit` et `before`.
- Les index sur `(hoomi_id, created_at)` permettent une récupération efficace.

### Chargement Différé

- Pour les Hoomis avec de longs historiques, seuls les messages récents sont chargés initialement.
- Les messages plus anciens sont chargés à la demande (infinite scroll).

### Archivage

À l'avenir, je pourrais implémenter un système d'archivage :
- Déplacer les messages très anciens vers un stockage moins cher (objet storage).
- Garder les métadonnées dans la base principale pour la recherche.

## Sauvegarde et Récupération

### Sauvegardes PostgreSQL

- Sauvegardes régulières de la base de données.
- Réplication pour la haute disponibilité.

### Récupération après Sinistre

- Possibilité de restaurer l'historique complet des messages à partir des sauvegardes.
- Les messages restent chiffrés dans les sauvegardes.

## Gestion de la Croissance

### Partitionnement

Pour les très grandes instances :
- Partitionner la table `messages` par `hoomi_id` ou par date.
- Cela permet de distribuer la charge sur plusieurs serveurs de base de données.

### Purge (Optionnelle)

Une fonctionnalité de purge pourrait être offerte aux Hoomis :
- Permettre aux admins de supprimer automatiquement les messages après une certaine durée.
- Cela nécessiterait des politiques de rétention configurables.

## Intégrité des Données

### Contraintes

- Clés étrangères pour maintenir la cohérence référentielle.
- `NOT NULL` sur les champs critiques.

### Audits

- J'envisage de mettre en place un système de log pour les opérations de suppression de messages, à des fins de sécurité.

## Limites Actuelles

- Pas de compression du contenu chiffré (le chiffrer le compresse déjà fortement).
- Pas de recherche full-text dans le contenu (impossible sans déchiffrer).
- Pas de mécanisme de "suppression pour tous" qui efface réellement le contenu du serveur (seul le marquage côté client existe).

## Évolution Future

- **Compression** : Évaluer si une compression avant chiffrement apporte des bénéfices.
- **Recherche** : Explorer des solutions de recherche sur données chiffrées (cryptographie fonctionnelle).
- **Migration** : Prévoir un mécanisme pour migrer les messages si le schéma de chiffrement change.