# Architecture des Applications Mobiles Hoomi

Je documente ici l'architecture que j'ai conçue pour les applications mobiles d'Hoomi, tant pour iOS que pour Android.

## Vue d'Ensemble

Les applications mobiles Hoomi sont conçues pour être :
- **Intuitives** : Interface utilisateur simple et familiale.
- **Sécurisées** : Intégration complète du chiffrement de bout en bout.
- **Performantes** : Optimisées pour les appareils mobiles.
- **Fonctionnelles** : Toutes les fonctionnalités principales accessibles.

## Architecture Générale

### Structure en Couches

```
┌─────────────────────────────────────┐
│           UI / Interface            │
├─────────────────────────────────────┤
│      Logique Applicative Mobile     │
├─────────────────────────────────────┤
│      Services (API, Crypto, etc.)   │
├─────────────────────────────────────┤
│        Plateforme (iOS/Android)     │
└─────────────────────────────────────┘
```

## iOS (Swift)

### Technologies

- **Langage** : Swift 5+
- **Framework UI** : UIKit (pour compatibilité étendue) avec SwiftUI pour de nouveaux composants
- **Gestion des dépendances** : Swift Package Manager
- **Chiffrement** : CryptoKit, CommonCrypto
- **Réseau** : URLSession, WebSocket
- **Stockage local** : CoreData, Keychain

### Structure des Dossiers

```
Hoomi-iOS/
├── HoomiApp/                 # Cible principale de l'application
│   ├── AppDelegate.swift     # Point d'entrée de l'application
│   ├── SceneDelegate.swift   # Gestion des scènes (iOS 13+)
│   ├── Resources/            # Images, fichiers locaux
│   └── Supporting Files/
├── Core/                     # Logique métier partagée
│   ├── Models/               # Modèles de données
│   ├── Services/             # Services réseau, crypto, etc.
│   ├── Managers/             # Gestionnaires d'état
│   └── Extensions/           # Extensions Swift
├── Features/                 # Fonctionnalités principales
│   ├── Authentication/       # Écrans d'authentification
│   ├── Home/                 # Écran principal (liste des Hoomis)
│   ├── Hoomi/                # Écrans d'un Hoomi spécifique
│   ├── Chat/                 # Écrans de messagerie
│   ├── Media/                # Gestion des médias
│   └── Profile/              # Gestion du profil
├── UI/                       # Composants d'interface réutilisables
│   ├── Components/           # Boutons, cellules, vues personnalisées
│   ├── Styles/               # Thèmes, couleurs, polices
│   └── Utils/                # Helpers UI
└── Tests/                    # Tests unitaires et d'intégration
```

## Android (Kotlin)

### Technologies

- **Langage** : Kotlin
- **Framework UI** : Jetpack Compose (moderne) avec compatibilité View System
- **Gestion des dépendances** : Gradle avec Kotlin DSL
- **Chiffrement** : Android Keystore, Tink (ou libsodium)
- **Réseau** : OkHttp, Retrofit, OkHttp-WebSocket
- **Stockage local** : Room, EncryptedSharedPreferences

### Structure des Dossiers

```
hoomi-android/
├── app/                      # Module principal de l'application
│   ├── src/
│   │   ├── main/
│   │   │   ├── java/         # Code source
│   │   │   └── res/          # Ressources
│   │   └── test/             # Tests
│   └── build.gradle.kts      # Configuration Gradle
├── core/                     # Module de base (logique métier partagée)
│   ├── src/
│   │   ├── main/
│   │   │   └── java/
│   │   └── test/
│   └── build.gradle.kts
├── features/                 # Modules de fonctionnalités
│   ├── auth/
│   ├── home/
│   ├── hoomi/
│   ├── chat/
│   ├── media/
│   └── profile/
├── ui/                       # Module UI (composants réutilisables)
│   └── src/
│       └── main/
│           └── java/
└── buildSrc/                 # Scripts de build personnalisés
```

## Intégration Backend

### API Client

Les deux plateformes utilisent le même API REST décrit dans `docs/backend/api.md`, avec :
- **Authentification** : JWT dans les headers.
- **Sérialisation** : JSON.
- **WebSockets** : Pour la messagerie en temps réel.

### Gestion des États

- **iOS** : Utilisation d'un pattern Coordinator pour la navigation, avec des ViewModels pour la logique.
- **Android** : Utilisation de ViewModel (Architecture Components) et Navigation Component.

## Sécurité Mobile

### Stockage des Clés

- **iOS** : Keychain Services pour stocker les clés privées de manière sécurisée.
- **Android** : Android Keystore System pour protéger les clés cryptographiques.

### Biométrie

- **iOS** : LocalAuthentication framework (Touch ID, Face ID).
- **Android** : BiometricPrompt API.

### Sécurité du Réseau

- **Certificate Pinning** : Pour prévenir les attaques MITM.
- **SSL** : Toutes les communications via HTTPS/WSS.

## Performance

### Mise en Cache

- **iOS** : NSCache pour les images, CoreData pour les données persistantes.
- **Android** : LruCache pour les images, Room pour les données persistantes.

### Chargement Différé

- Implémentation de lazy loading pour les listes de messages et de médias.

### Optimisation de la Batterie

- Gestion appropriée des connexions réseau en arrière-plan.
- Utilisation de WorkManager (Android) / BackgroundTasks (iOS) pour les tâches différées.

## Localisation

Support multilingue avec :
- **iOS** : Localizable.strings
- **Android** : res/values/strings.xml avec qualifiers de langue

## Accessibilité

Suivi des lignes directrices de :
- **iOS** : WWDC Accessibility guidelines
- **Android** : Material Design Accessibility

## Tests

### Tests Unitaires

- **iOS** : XCTest
- **Android** : JUnit, Mockito

### Tests UI

- **iOS** : XCTest UI Testing
- **Android** : Espresso

## CI/CD

### Build

- Utilisation de GitHub Actions pour les builds automatisés.
- Génération d'APK/IPA pour les déploiements.

### Distribution

- **iOS** : TestFlight pour la distribution bêta.
- **Android** : Google Play Console pour la distribution bêta et production.

## Dépendances

### Gestionnaires

- **iOS** : Swift Package Manager (SPM) ou CocoaPods.
- **Android** : Gradle.

### Bibliothèques Clés

- **Réseau** : Alamofire (iOS), OkHttp (Android)
- **Chiffrement** : CryptoSwift (iOS), Tink (Android)
- **UI** : SnapKit (iOS), ConstraintLayout (Android)
- **Images** : Kingfisher (iOS), Glide (Android)

## Évolution Future

- **Partage de Code** : Évaluer l'utilisation de Kotlin Multiplatform Mobile (KMM) ou Flutter pour partager de la logique.
- **Widgets** : Implémenter des widgets pour l'accès rapide aux notifications.
- **Watch App** : Développer une application Apple Watch pour les notifications.