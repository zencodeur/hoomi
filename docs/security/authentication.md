# Authentification dans Hoomi

Je documente ici le système d'authentification que j'ai mis en place pour Hoomi.

## Vue d'Ensemble

Mon système d'authentification repose sur :
1. **JSON Web Tokens (JWT)** : Pour la gestion des sessions côté client.
2. **Hachage de mot de passe** : Pour stocker les mots de passe de manière sécurisée.
3. **OAuth 2.0** : Pour l'authentification tierce (à implémenter plus tard).

## Workflow d'Authentification

### Inscription (Register)

1. L'utilisateur soumet un email, un nom et un mot de passe.
2. Je vérifie que l'email n'est pas déjà utilisé.
3. Je hache le mot de passe avec un sel fort (bcrypt ou argon2).
4. Je crée un nouvel utilisateur dans la base de données.
5. Je génère un JWT pour l'utilisateur nouvellement créé.
6. Je renvoie le token et les informations de l'utilisateur.

### Connexion (Login)

1. L'utilisateur soumet son email et son mot de passe.
2. Je récupère l'utilisateur par son email.
3. Je compare le mot de passe soumis avec le hash stocké.
4. Si la vérification est réussie, je génère un nouveau JWT.
5. Je renvoie le token et les informations de l'utilisateur.

### Validation de Session

1. Pour chaque requête protégée, le client envoie le JWT dans l'en-tête `Authorization: Bearer <token>`.
2. Le serveur vérifie la validité du token (signature, expiration).
3. Si le token est valide, je décode l'ID utilisateur et je charge les informations de l'utilisateur.
4. La requête est traitée avec le contexte de l'utilisateur authentifié.

## JSON Web Tokens (JWT)

### Structure

Je utilise des JWT composés de trois parties :
- **Header** : Algorithme de signature (HS256).
- **Payload** : Données de l'utilisateur (ID, email, nom, etc.) et métadonnées (iat, exp).
- **Signature** : Signature cryptographique pour vérifier l'intégrité.

### Propriétés

- **Durée de vie** : 24 heures.
- **Algorithme** : HS256 (à renforcer avec RS256 dans une version future).
- **Revocation** : Actuellement, je ne gère pas la révocation des tokens. Un utilisateur déconnecté côté client peut toujours utiliser un token valide jusqu'à son expiration.

## Sécurité des Mots de Passe

### Hachage

J'utilise **bcrypt** avec un coût de 12 pour hacher les mots de passe. Cela me fournit :
- Une résistance aux attaques par force brute.
- Un sel automatique et unique pour chaque mot de passe.
- Un temps de hachage constant malgré les variations d'entrée.

### Contraintes

Pour les nouveaux mots de passe, j'impose :
- Longueur minimale de 8 caractères.
- Présence de caractères en majuscule, minuscule, chiffres et caractères spéciaux.
- Interdiction des mots de passe courants (via une liste noire).

## Protection contre les Attaques

### Brute Force

- Je limite le nombre de tentatives de connexion échouées par IP et par utilisateur.
- Après 5 tentatives échouées, je bloque temporairement l'IP pendant 15 minutes.

### Enumeration d'Emails

- Les messages d'erreur pour l'inscription/connexion sont génériques pour ne pas révéler si un email existe déjà.

### Cross-Site Request Forgery (CSRF)

- Les tokens JWT sont envoyés via des en-têtes HTTP, ce qui les immunise naturellement contre les attaques CSRF classiques.

## Authentification à Deux Facteurs (2FA)

Cette fonctionnalité est prévue pour une version future :
- **TOTP** (Time-Based One-Time Password) via une application d'authentification.
- **Backup codes** pour récupérer l'accès en cas de perte du dispositif 2FA.

## Gestion des Sessions

Actuellement, je suis stateless avec JWT. Chaque client gère son propre token. Pour une gestion plus fine des sessions (liste des sessions actives, déconnexion à distance), je pourrais introduire un stockage côté serveur des sessions (Redis) dans une version future.

## Révocation de Tokens

Pour permettre la déconnexion immédiate, je pourrais implémenter :
- Une liste de révocation côté serveur.
- Un raccourcissement de la durée de vie des tokens avec un système de refresh token.