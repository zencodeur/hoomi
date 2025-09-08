# Chiffrement de Bout en Bout (E2EE) dans Hoomi

Je documente ici l'implémentation du chiffrement de bout en bout que j'ai mise en place pour Hoomi, qui est un pilier fondamental de la sécurité de l'application.

## Vue d'Ensemble

Le chiffrement de bout en bout (E2EE) garantit que seuls les membres d'un Hoomi peuvent lire le contenu des messages et des fichiers partagés. Même le serveur Hoomi ne peut pas accéder au contenu en clair.

## Principes de Base

### Chiffrement Asymétrique

J'utilise un système de chiffrement asymétrique (clés publiques/privées) pour l'échange sécurisé des clés de chiffrement symétrique.

### Chiffrement Symétrique

Pour le chiffrement réel du contenu (messages, fichiers), j'utilise un chiffrement symétrique avec des clés générées de manière sécurisée.

### Signal Protocol

Je m'inspire du Signal Protocol, largement reconnu pour sa robustesse, bien que mon implémentation soit adaptée aux spécificités de Hoomi.

## Gestion des Clés

### Clés d'Identité

Chaque utilisateur possède une **clé d'identité** (paire de clés RSA 4096 bits) :
- **Clé privée** : Stockée de manière sécurisée sur l'appareil de l'utilisateur.
- **Clé publique** : Partagée avec les autres membres des Hoomis.

### Clés de Session

Pour chaque session de communication (conversation dans un Hoomi), une **clé de session symétrique** (AES-256) est générée.

### Distribution des Clés

1. Lorsqu'un utilisateur rejoint un Hoomi, sa clé publique d'identité est récupérée par les autres membres.
2. Pour envoyer un message, l'expéditeur :
   - Génère une nouvelle clé de session symétrique.
   - Chiffre le message avec cette clé de session.
   - Chiffre la clé de session avec la clé publique de chaque destinataire.
   - Envoie le message chiffré et les clés de session chiffrées.

## Chiffrement des Messages

### Processus d'Envoi

1. **Préparation** :
   - L'utilisateur compose un message en clair.
2. **Génération de la clé de session** :
   - Une clé symétrique AES-256 est générée de manière sécurisée.
3. **Chiffrement du message** :
   - Le message est chiffré avec la clé de session (AES-GCM).
4. **Chiffrement des clés de session** :
   - La clé de session est chiffrée avec la clé publique de chaque membre du Hoomi.
5. **Envoi** :
   - Le message chiffré, les clés de session chiffrées et les métadonnées sont envoyés au serveur.

### Processus de Réception

1. **Récupération** :
   - Le destinataire récupère le message chiffré et sa clé de session chiffrée.
2. **Déchiffrement de la clé de session** :
   - La clé de session est déchiffrée avec la clé privée du destinataire.
3. **Déchiffrement du message** :
   - Le message est déchiffré avec la clé de session.

## Chiffrement des Fichiers

Le même principe s'applique aux fichiers :
1. Le fichier est chiffré avec une clé symétrique.
2. La clé est chiffrée pour chaque membre du Hoomi.
3. Le fichier chiffré et les clés chiffrées sont stockés.

## Stockage des Clés

### Côté Client

- **Clés privées d'identité** : Stockées dans le trousseau sécurisé de l'appareil (Keychain sur iOS, Keystore sur Android, localStorage chiffré sur web).
- **Clés de session** : Générées à la volée, jamais stockées de manière persistante.

### Côté Serveur

- Le serveur ne stocke **jamais** les clés privées.
- Le serveur stocke uniquement les clés publiques d'identité.
- Les clés de session chiffrées sont transmises via le serveur mais celui-ci ne peut pas les déchiffrer.

## Gestion des Nouveaux Membres

Lorsqu'un nouveau membre rejoint un Hoomi :
1. Sa clé publique est récupérée par les membres existants.
2. Les messages et fichiers existants ne sont pas automatiquement accessibles à ce nouveau membre (ils n'ont pas les clés de session correspondantes).
3. Pour les nouveaux contenus, la clé de session est chiffrée avec la clé publique du nouveau membre également.

## Rotation des Clés

### Clés d'Identité

Les clés d'identité ont une longue durée de vie. Elles sont régénérées uniquement si l'utilisateur le demande explicitement ou si le trousseau est compromis.

### Clés de Session

Les clés de session sont **éphémères** :
- Une nouvelle clé est générée pour chaque message ou fichier.
- Elles ne sont jamais réutilisées.

## Sécurité et Menaces

### Protection contre le Serveur

Même si le serveur est compromis, le contenu chiffré reste inaccessible sans les clés privées des utilisateurs.

### Protection contre les Attaques de l'Homme du Milieu

L'authentification des clés publiques via l'application et les vérifications d'identité protègent contre ce type d'attaque.

### Récupération en Cas de Perte d'Appareil

C'est un défi. Si un utilisateur perd son appareil et donc sa clé privée, il perd l'accès aux messages déchiffrés. Une solution de sauvegarde sécurisée des clés (chiffrée avec un mot de passe maître) est prévue.

## Implémentation Technique

### Bibliothèques

- **WebCrypto API** : Pour les opérations cryptographiques dans le navigateur.
- **CryptoKit (iOS)** et **Android Keystore** : Pour les plateformes mobiles.
- **Libsodium** (Go) : Pour les opérations côté serveur.

### Formats

- **Clés** : Format PEM pour le stockage, JWK pour l'échange.
- **Messages chiffrés** : Format personnalisé incluant le contenu chiffré, l'IV et les tags d'authentification.

## Limitations Actuelles

- **Pas de vérification d'identité avancée** : Pour l'instant, je fais confiance au processus d'invitation.
- **Pas de synchronisation des clés entre appareils** : Un utilisateur doit importer ses clés sur chaque appareil.

## Évolution Future

- **Sauvegarde de clés** : Implémenter un système de sauvegarde chiffrée des clés privées.
- **Vérification d'identité** : Ajouter des codes de vérification pour confirmer l'identité des membres.
- **Synchronisation multi-appareils** : Permettre à un utilisateur de synchroniser ses clés entre plusieurs appareils de manière sécurisée.