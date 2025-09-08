# ADR-007: Technologies pour les Applications Mobiles Hoomi

## Statut

Proposé

## Contexte

Je dois choisir les technologies pour développer les applications mobiles d'Hoomi. Les options principales sont :

### Approche Native Séparée
1.  **iOS** : Swift avec UIKit/SwiftUI
2.  **Android** : Kotlin avec Jetpack Compose/View System

### Approche Multiplateforme
1.  **Flutter** : SDK de Google utilisant Dart
2.  **React Native** : Framework de Facebook utilisant JavaScript/TypeScript
3.  **Xamarin** : Framework de Microsoft utilisant C#
4.  **Kotlin Multiplatform Mobile (KMM)** : Partage de code Kotlin entre iOS et Android

Les considérations importantes pour moi sont :
-   **Performance** : L'expérience doit être fluide et native.
-   **Accès aux APIs natives** : Besoin d'accéder aux fonctionnalités spécifiques de chaque plateforme (biométrie, chiffrement, etc.).
-   **Équipe et expertise** : Mon expérience et celle du marché du travail.
-   **Temps de développement** : Rapidité de mise sur le marché.
-   **Maintenabilité** : Facilité de maintenance à long terme.
-   **Coût** : Coût de développement et de maintenance.

## Décision

Je choisis une **approche native séparée** :
- **iOS** : Swift avec UIKit (et SwiftUI pour de nouveaux composants)
- **Android** : Kotlin avec Jetpack Compose (et View System pour la compatibilité)

## Conséquences

-   **Positives** :
    -   **Performance optimale** : Applications parfaitement optimisées pour chaque plateforme.
    -   **Meilleure intégration** : Accès complet aux APIs et guidelines de design natives.
    -   **Expérience utilisateur native** : Respect des conventions de chaque plateforme.
    -   **Écosystème mature** : Outils et bibliothèques bien établis.
    -   **Maîtrise totale** : Contrôle complet sur chaque aspect de l'application.
-   **Négatives** :
    -   **Double développement** : Deux bases de code à maintenir.
    -   **Ressources** : Nécessite deux ensembles de compétences ou un développeur full-stack mobile.
    -   **Synchronisation** : Risque de décalage fonctionnel entre les deux versions.

Je suis conscient de ces compromis. En tant que développeur unique avec un budget limité, je pourrais envisager Flutter ou React Native à l'avenir pour réduire la duplication. Cependant, pour une application qui met l'accent sur la sécurité et la performance comme Hoomi, l'approche native me semble la meilleure pour démarrer. Elle me permet de tirer parti de toutes les fonctionnalités spécifiques de chaque plateforme, notamment en matière de chiffrement et de gestion sécurisée des clés.

J'envisage de partager une partie de la logique métier (modèles, services non-UI) via une approche comme Kotlin Multiplatform Mobile dans une version future si le besoin de réduction de duplication se fait sentir.