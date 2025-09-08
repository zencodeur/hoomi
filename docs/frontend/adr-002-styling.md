# ADR-002: Approche de Styling pour le Frontend

## Statut

Proposé

## Contexte

Pour le frontend de Hoomi, je dois choisir une approche de styling cohérente. Les options principales sont :

1.  **CSS Modules** : Fichiers CSS locaux à un composant, avec des noms de classes uniques générés.
2.  **Styled Components** : Styling dans le JS/TS avec tagged template literals, générant des classes uniques.
3.  **Tailwind CSS** : Framework CSS utility-first.
4.  **CSS traditionnel avec PostCSS** : Fichiers CSS globaux avec des outils de post-traitement pour des fonctionnalités avancées.

Les considérations importantes sont :
-   **Maintenabilité** : Comment gérer la complexité du style à grande échelle.
-   **Performance** : Taille du bundle CSS final.
-   **Coût d'apprentissage** : Ma propre familiarité avec chaque option.
-   **Thématisation** : Facilité de changement de thème (clair/sombre).
-   **Réutilisabilité** : Facilité de créer et réutiliser des composants stylisés.

## Décision

Je choisis **CSS Modules** comme approche principale de styling pour le frontend de Hoomi.

## Conséquences

-   **Positives** :
    -   Scope CSS local par défaut, évite les conflits de noms.
    -   Syntaxe CSS familière.
    -   Bon équilibre entre performance et fonctionnalité.
    -   Bien supporté par Next.js.
    -   Permet une transition vers d'autres solutions si nécessaire.
-   **Négatives** :
    -   Moins de dynamisme que Styled Components (pas de props pour changer le style).
    -   Nécessite une discipline pour organiser les fichiers CSS.
    -   Pas de garantie de "zero-runtime" comme avec certaines autres solutions.