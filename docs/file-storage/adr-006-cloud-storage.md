# ADR-006: Choix de la Solution de Stockage Cloud pour Hoomi

## Statut

Proposé

## Contexte

Je dois choisir une solution de stockage cloud pour les fichiers dans Hoomi. Les options principales sont :

1.  **Amazon S3** : Service de stockage d'objets hautement évolutif.
2.  **Google Cloud Storage** : Service de stockage d'objets de Google Cloud.
3.  **Azure Blob Storage** : Service de stockage d'objets de Microsoft Azure.
4.  **Stockage local** : Gérer mon propre stockage sur les serveurs.

Les considérations importantes pour moi sont :
-   **Coût** : Je dois minimiser les coûts, surtout au début.
-   **Fiabilité** : Le stockage doit être fiable avec une haute disponibilité.
-   **Performance** : Des temps de réponse rapides pour les uploads/downloads.
-   **Évolutivité** : Capacité à grandir avec le nombre d'utilisateurs.
-   **Facilité d'intégration** : Intégration simple avec mon backend Go.
-   **Fonctionnalités** : Fonctionnalités comme les liens pré-signés, les métadonnées.

## Décision

Je choisis **Google Cloud Storage** comme solution de stockage cloud pour Hoomi.

## Conséquences

-   **Positives** :
    -   **Généreux niveau gratuit** : 5 Go de stockage et 1 Go de trafic par mois gratuitement, ce qui est idéal pour commencer.
    -   **Intégration avec Firebase** : Si j'utilise Firebase pour d'autres services, l'intégration est naturelle.
    -   **Haute performance** : Réputé pour ses performances et sa faible latence.
    -   **Bibliothèques Go matures** : Le SDK Google Cloud pour Go est bien documenté.
    -   **Fonctionnalités avancées** : Support des liens pré-signés, des métadonnées, du versioning.
-   **Négatives** :
    -   **Verrouillage fournisseur** : Comme toutes les solutions cloud, il y a un risque de dépendance.
    -   **Coûts à long terme** : Les coûts peuvent augmenter rapidement avec la croissance.

Je suis conscient de ces compromis, mais le niveau gratuit généreux et l'intégration facile avec Go font de Google Cloud Storage le meilleur choix pour démarrer. Je pourrai réévaluer si les besoins ou les coûts changent significativement.