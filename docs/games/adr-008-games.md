# ADR-008: Architecture et Technologie pour le Système de Jeux Hoomi

## Statut

Proposé

## Contexte

Je dois choisir l'architecture et les technologies pour implémenter le système de jeux familiaux dans Hoomi. Les options principales sont :

### Approche 1 : Jeux Natifs Intégrés
- Développer chaque jeu comme une partie intégrante de l'application principale (iOS/Android/Web).
- Utiliser les frameworks natifs (SwiftUI, Jetpack Compose, React).

### Approche 2 : Jeux Web dans Conteneur
- Développer les jeux comme des applications web légères.
- Les intégrer dans l'application principale via un conteneur (iframe/webview).

### Approche 3 : Moteur de Jeu Spécialisé
- Utiliser un moteur de jeu comme Unity ou Godot.
- Développer les jeux avec ce moteur et les intégrer.

Les considérations importantes pour moi sont :
-   **Performance** : L'expérience de jeu doit être fluide.
-   **Développement** : Facilité et rapidité de développement des jeux.
-   **Intégration** : Comment les jeux s'intègrent avec le reste de l'application.
-   **Mise à jour** : Facilité de mise à jour des jeux sans mise à jour de l'application.
-   **Coût** : Coût de développement et de maintenance.
-   **Équipe** : Mes compétences et celles du marché du travail.

## Décision

Je choisis l'**Approche 2 : Jeux Web dans Conteneur**.

## Conséquences

-   **Positives** :
    -   **Développement rapide** : Utilisation de technologies web standard (HTML/CSS/JS) que je maîtrise bien.
    -   **Mise à jour indépendante** : Les jeux peuvent être mis à jour sans mettre à jour l'application principale.
    -   **Portabilité** : Les mêmes jeux fonctionnent sur web, iOS et Android.
    -   **Coût réduit** : Pas besoin d'apprendre de nouveaux frameworks ou moteurs.
    -   **Écosystème riche** : Accès à de nombreuses bibliothèques et outils web.
    -   **Itération rapide** : Facile de tester et déployer de nouvelles versions.
-   **Négatives** :
    -   **Performance** : Potentiellement moins performant que les jeux natifs pour des graphismes complexes.
    -   **Accès limité** : Accès restreint à certaines APIs natives (caméra, capteurs).
    -   **Isolement** : Nécessité de bien isoler les jeux pour la sécurité.

Je suis conscient de ces compromis. Pour les jeux familiaux simples que je prévois (jeux de société, puzzles, cartes), les performances web seront largement suffisantes. L'avantage de pouvoir déployer rapidement de nouveaux jeux et de les mettre à jour indépendamment de l'application principale est déterminant pour un projet avec des ressources limitées.

Si plus tard je veux développer des jeux plus graphiques, je pourrais envisager une approche hybride ou migrer vers un moteur spécialisé. Mais pour l'instant, l'approche web dans un conteneur répond parfaitement aux besoins de Hoomi.