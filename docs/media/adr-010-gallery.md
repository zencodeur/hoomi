# ADR-010: Architecture de la Galerie Multimédia pour Hoomi

## Statut

Proposé

## Contexte

Je dois choisir l'architecture de la galerie multimédia pour Hoomi. Les options principales sont :

### Approche 1 : Galerie Tierce Intégrée
- Utiliser une solution de galerie existante (comme PhotoSwipe, LightGallery) et l'intégrer.
- Personnaliser l'apparence et ajouter les fonctionnalités spécifiques à Hoomi.

### Approche 2 : Galerie Personnalisée from Scratch
- Développer une galerie entièrement personnalisée avec React/Next.js pour le web et Swift/Kotlin pour mobile.
- Contrôle total sur l'expérience utilisateur et les fonctionnalités.

### Approche 3 : Hybridation
- Utiliser des composants tiers pour les parties complexes (visionneuse, slider) mais développer l'interface de gestion (grille, albums) en interne.

Les considérations importantes pour moi sont :
-   **Contrôle** : Capacité à personnaliser l'expérience exactement comme je la veux.
-   **Temps de développement** : Rapidité de mise sur le marché.
-   **Maintenabilité** : Facilité de maintenance et de mise à jour.
-   **Performance** : Fluidité de l'interface, surtout avec de grandes collections.
-   **Intégration** : Facilité d'intégration avec le reste de l'application Hoomi.
-   **Coût** : Coût de développement initial et de maintenance à long terme.

## Décision

Je choisis l'**Approche 3 : Hybridation**.

## Conséquences

-   **Positives** :
    -   **Développement rapide** : Utilisation de composants éprouvés pour les parties complexes.
    -   **Contrôle suffisant** : Développement personnalisé pour l'interface de gestion qui est le cœur de l'expérience Hoomi.
    -   **Maintenabilité** : Moins de code personnalisé à maintenir.
    -   **Performance** : Les composants tiers sont optimisés pour la performance.
    -   **Flexibilité** : Possibilité de remplacer les composants tiers si nécessaire.
-   **Négatives** :
    -   **Dépendances** : Dépendance à des bibliothèques tierces.
    -   **Personnalisation limitée** : Certaines personnalisation avancées pourraient être difficiles.

Je suis conscient de ces compromis. L'approche hybride me semble être le meilleur équilibre entre développement rapide, maintenabilité et contrôle sur l'expérience utilisateur. Elle me permet de me concentrer sur les éléments différenciants d'Hoomi (gestion des albums, recherche, collaboration) tout en bénéficiant de la robustesse des composants tiers pour les tâches standard (affichage d'images, navigation).

Cela correspond à l'approche générale d'Hoomi qui vise à réutiliser ce qui existe lorsque cela fait sens, tout en développant les parties critiques en interne.