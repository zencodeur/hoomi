# Sécurité des Applications Mobiles Hoomi

Je documente ici les considérations et implémentations de sécurité spécifiques aux applications mobiles d'Hoomi.

## Vue d'Ensemble

La sécurité des apps mobiles Hoomi repose sur :
- **Chiffrement de bout en bout** : Intégration complète du système E2EE.
- **Protection des données sensibles** : Stockage sécurisé des clés et informations personnelles.
- **Authentification renforcée** : Mécanismes d'authentification biométrique.
- **Sécurité réseau** : Protection contre les interceptions et attaques MITM.
- **Conformité** : Respect des réglementations (RGPD, etc.).

## Chiffrement de Bout en Bout (E2EE)

### Gestion des Clés

#### Stockage des Clés Privées

- **iOS** : Utilisation du **Keychain Services** avec protection `kSecAttrAccessibleAfterFirstUnlock`.
  - Les clés sont chiffrées par le système d'exploitation.
  - Accès uniquement par l'application Hoomi.
  - Protection contre l'extraction par des outils externes.

- **Android** : Utilisation du **Android Keystore System**.
  - Les clés sont générées et stockées dans un environnement sécurisé.
  - Utilisation de `KeyGenParameterSpec` pour spécifier les conditions d'utilisation.
  - Protection contre l'extraction et la copie.

#### Sauvegarde des Clés

- **Défi** : Les sauvegardes système (iCloud, Google Backup) ne doivent pas inclure les clés privées.
- **Solution** :
  - **iOS** : Utilisation de `kSecAttrAccessibleWhenUnlockedThisDeviceOnly` pour les clés critiques.
  - **Android** : Définition de `setIsStrongBoxBacked(true)` et `setCriticalToDeviceEncryption(false)`.

#### Export/Import des Clés

- Fonctionnalité permettant aux utilisateurs d'exporter leurs clés de manière sécurisée.
- Export chiffré avec un mot de passe maître fourni par l'utilisateur.
- Import avec vérification de l'intégrité et de l'authenticité.

### Chiffrement des Données

#### Messages

- **Processus** :
  1. Génération d'une clé de session AES-256 aléatoire.
  2. Chiffrement du message avec AES-GCM.
  3. Chiffrement de la clé de session avec la clé publique de chaque destinataire.
  4. Envoi du message chiffré et des clés de session chiffrées.

#### Fichiers

- **Processus** :
  1. Chiffrement du fichier avec AES-256 avant l'upload.
  2. La clé de chiffrement est chiffrée pour chaque membre du Hoomi.
  3. Upload du fichier chiffré vers le cloud storage.

## Authentification

### Jetons d'Authentification (JWT)

- **Stockage** :
  - **iOS** : Dans le Keychain, chiffré.
  - **Android** : Dans `EncryptedSharedPreferences`.
- **Renouvellement** :
  - Mise en place d'un système de refresh token.
  - Renouvellement automatique avant expiration.

### Authentification Biométrique

#### Implémentation

- **iOS** : 
  - Utilisation du framework **LocalAuthentication**.
  - Support de Touch ID et Face ID.
  - Fallback sur le code d'accès si biométrie non disponible/échec.

- **Android** :
  - Utilisation de **BiometricPrompt API** (Android 9+).
  - Support d'empreinte digitale, reconnaissance faciale, etc.
  - Fallback sur PIN/mot de passe.

#### Sécurité

- La biométrie est utilisée pour déchiffrer une clé locale, pas pour authentifier directement l'utilisateur.
- En cas d'échec biométrique, l'utilisateur peut utiliser un code PIN de secours.

### Code PIN de Secours

- Un code PIN à 6 chiffres est généré à la première ouverture de l'application.
- Ce PIN est utilisé pour chiffrer une clé maître qui déchiffre les clés de l'application.
- Le PIN est stocké de manière sécurisée :
  - **iOS** : Haché avec PBKDF2 et stocké dans le Keychain.
  - **Android** : Haché avec PBKDF2 et stocké dans `EncryptedSharedPreferences`.

## Sécurité Réseau

### Communication HTTPS/WSS

- Toutes les communications avec le backend se font via HTTPS (API) et WSS (WebSocket).
- Vérification stricte des certificats SSL.

### Certificate Pinning

- Mise en place de **Certificate Pinning** pour prévenir les attaques MITM.
- **iOS** : Configuration dans `Info.plist` ou via URLSession.
- **Android** : Configuration dans le fichier de ressources réseau ou via OkHttpClient.

### Protection contre les Attaques

#### Brute Force

- Limitation du nombre de tentatives d'authentification.
- Blocage temporaire après plusieurs échecs.

#### Session Management

- Invalidations des sessions à la déconnexion.
- Gestion sécurisée des sessions en arrière-plan.

## Protection des Données

### Données Sensibles

- **Clés cryptographiques** : Stockées dans le trousseau sécurisé.
- **Tokens d'authentification** : Stockés de manière chiffrée.
- **Données personnelles** : Chiffrées localement si stockées.

### Cache et Données Temporaires

- Nettoyage automatique du cache sensible.
- Chiffrement des données temporaires si nécessaire.
- Éviction du cache en arrière-plan.

### Screenshots et Enregistrement

- **Prévention** : Désactivation des screenshots dans les écrans sensibles (clés, messages déchiffrés).
- **iOS** : Utilisation de `UIScreen.screens.first?.windows.forEach { $0.secure = true }`.
- **Android** : Utilisation de `FLAG_SECURE` sur les fenêtres sensibles.

## Gestion des Permissions

### Permissions Minimales

- Demande de permissions uniquement quand nécessaire.
- Explication claire de pourquoi chaque permission est requise.

### Permissions Critiques

- **Stockage** : Pour accéder aux photos/vidéos à partager.
- **Appareil photo** : Pour prendre des photos directement.
- **Microphone** : Pour les messages vocaux (futur).
- **Contacts** : Pour faciliter l'invitation (avec consentement explicite).

## Détection d'Environnement Non-Sécurisé

### Root/Jailbreak Detection

- **Android** : Vérification de l'environnement rooté.
- **iOS** : Vérification de l'environnement jailbreaké.
- **Action** : Avertissement à l'utilisateur et limitation des fonctionnalités sensibles.

### Emulateur/Simulateur Detection

- Détection de l'exécution dans un environnement émulé.
- Limitation des fonctionnalités de sécurité dans ces environnements.

## Intégrité de l'Application

### Code Obfuscation

- **Android** : Utilisation de R8/ProGuard pour obscurcir le code.
- **iOS** : Le langage Swift offre une certaine protection par défaut.

### Runtime Integrity Checks

- Vérifications périodiques de l'intégrité de l'application.
- Détection de tentatives de modification du code.

## Mise à Jour de Sécurité

### Mises à Jour Obligatoires

- Possibilité de forcer les utilisateurs à mettre à jour pour des correctifs de sécurité critiques.
- Vérification de la version de l'application côté serveur.

### Patching Dynamique

- Interdiction des mécanismes de mise à jour dynamique de code (comme les frameworks de hotfix).

## Tests de Sécurité

### Tests Automatisés

- Intégration de tests de sécurité dans le pipeline CI/CD.
- Utilisation d'outils comme MobSF (Mobile Security Framework).

### Tests Manuels

- Audits de sécurité réguliers.
- Tests d'intrusion sur les applications.

## Conformité et Réglementation

### RGPD

- Mise en œuvre des droits des utilisateurs (accès, rectification, effacement).
- Minimisation des données collectées.
- Transparence sur l'utilisation des données.

### COPPA (Protection de la vie privée des enfants)

- Conformité pour les utilisateurs mineurs.
- Mécanismes de consentement parental.

## Surveillance et Journalisation

### Logs de Sécurité

- Journalisation des événements de sécurité importants.
- **Attention** : Ne jamais logger de données sensibles (clés, mots de passe, messages).

### Détection d'Anomalies

- Surveillance des comportements suspects (tentatives de connexion inhabituelles, etc.).

## Plan de Réponse aux Incidents

### Procédures

- Étapes à suivre en cas de découverte d'une vulnérabilité.
- Communication avec les utilisateurs si nécessaire.

### Contact

- Moyen pour les chercheurs en sécurité de signaler des vulnérabilités.

Cette approche de sécurité vise à protéger les utilisateurs de Hoomi contre les menaces spécifiques aux applications mobiles, tout en maintenant une expérience utilisateur fluide.