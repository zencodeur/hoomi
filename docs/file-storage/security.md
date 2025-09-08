# Sécurité du Stockage de Fichiers Hoomi

Je documente ici comment je sécurise le stockage de fichiers dans Hoomi, qui est un aspect critique de l'application.

## Vue d'Ensemble

La sécurité du stockage de fichiers repose sur :
- **Chiffrement de bout en bout** : Les fichiers sont chiffrés côté client avant d'être envoyés au serveur.
- **Contrôle d'accès** : Seuls les membres d'un Hoomi peuvent accéder à ses fichiers.
- **Stockage sécurisé** : Le stockage cloud utilise des protocoles de sécurité robustes.

## Chiffrement de Bout en Bout

### Processus d'Upload

1. **Génération de Clé** :
   - Une clé symétrique AES-256 unique est générée pour chaque fichier.
2. **Chiffrement du Fichier** :
   - Le fichier est chiffré avec cette clé symétrique (AES-GCM).
3. **Chiffrement de la Clé** :
   - La clé symétrique est chiffrée avec la clé publique de chaque membre du Hoomi.
4. **Envoi** :
   - Le fichier chiffré et les clés chiffrées sont envoyés au serveur.

### Processus de Téléchargement

1. **Récupération** :
   - Le client récupère le fichier chiffré et sa clé chiffrée.
2. **Déchiffrement de la Clé** :
   - La clé symétrique est déchiffrée avec la clé privée du destinataire.
3. **Déchiffrement du Fichier** :
   - Le fichier est déchiffré avec la clé symétrique.

## Stockage des Clés

### Côté Client

- **Clés privées** : Stockées dans le trousseau sécurisé de l'appareil.
- **Clés chiffrées** : Reçues du serveur et déchiffrées localement.

### Côté Serveur

- **Clés privées** : Jamais stockées.
- **Clés chiffrées** : Stockées dans la base de données `media.metadata` avec les autres métadonnées.
- **Fichiers** : Stockés chiffrés, donc inutilisables sans les clés.

## Contrôle d'Accès

### Authentification

- Toutes les requêtes d'upload/download nécessitent un token JWT valide.

### Autorisation

- Le service vérifie que l'utilisateur est membre du Hoomi auquel le fichier appartient.

### Liens Pré-signés

- Pour le téléchargement direct depuis le stockage cloud, des liens pré-signés à courte durée de vie sont générés.

## Intégrité

### Hachage

- Un hachage (SHA-256) du fichier original est calculé avant le chiffrement.
- Ce hachage peut être utilisé pour vérifier l'intégrité après déchiffrement.

### Tags d'Authentification

- L'algorithme AES-GCM fournit des tags d'authentification qui garantissent l'intégrité du fichier chiffré.

## Sécurité du Stockage Cloud

### Chiffrement côté Serveur

- Le fournisseur cloud chiffre les données au repos par défaut.
- Notre chiffrement E2EE s'applique en plus de ce chiffrement.

### Accès API

- Les accès au bucket cloud sont restreints aux services backend uniquement.
- Les utilisateurs n'ont jamais d'accès direct au stockage cloud.

## Gestion des Nouveaux Membres

Lorsqu'un nouveau membre rejoint un Hoomi :
1. Il ne peut pas accéder aux fichiers existants (il n'a pas les clés).
2. Pour les nouveaux fichiers, la clé est chiffrée avec sa clé publique également.

## Rotation des Clés

### Clés de Fichiers

- Chaque fichier a sa propre clé symétrique, qui n'est jamais réutilisée.

### Clés d'Identité

- Les clés d'identité (RSA) ont une longue durée de vie mais peuvent être régénérées.

## Sauvegarde des Clés

C'est un défi majeur. Si un utilisateur perd son appareil :
- Il perd ses clés privées.
- Il perd l'accès aux fichiers qu'il ne peut pas déchiffrer.
- Une solution de sauvegarde chiffrée (avec mot de passe maître) est prévue.

## Protection contre les Attaques

### Attaque par Force Brute

- Les clés AES-256 sont résistantes aux attaques par force brute.
- Les clés RSA de 4096 bits offrent une sécurité élevée.

### Attaque de l'Homme du Milieu

- L'authentification des clés publiques via l'application protège contre ce type d'attaque.

### Exposition du Stockage

- Même si le stockage cloud est compromis, les fichiers restent chiffrés.

## Limitations Actuelles

- **Pas de vérification d'identité avancée** : Pour l'instant, je fais confiance au processus d'invitation.
- **Pas de synchronisation des clés entre appareils** : Un utilisateur doit importer ses clés sur chaque appareil.

## Évolution Future

- **Sauvegarde de clés** : Implémenter un système de sauvegarde chiffrée des clés privées.
- **Vérification d'identité** : Ajouter des codes de vérification pour confirmer l'identité des membres.
- **Synchronisation multi-appareils** : Permettre à un utilisateur de synchroniser ses clés entre plusieurs appareils de manière sécurisée.