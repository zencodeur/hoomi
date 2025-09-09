# Transmission d'Héritage Numérique Hoomi

Je documente ici le système de transmission d'héritage numérique d'Hoomi, qui permet aux membres de la famille de planifier, gérer et transmettre leurs biens numériques et leurs volontés de manière sécurisée et organisée.

## Vue d'Ensemble

Le système de **Transmission d'Héritage Numérique** d'Hoomi est conçu pour être :
- **Sécurisé** : Toutes les informations sont chiffrées de bout en bout.
- **Contrôlé** : Chaque membre décide précisément ce qui est transmis, à qui, et quand.
- **Flexible** : Adapté aux lois et coutumes locales.
- **Intuitif** : Interface simple pour gérer des questions complexes.
- **Intégré** : Lié étroitement aux autres systèmes d'Hoomi (messagerie, médias, arbre généalogique).

## Interface Utilisateur

### Web (Next.js)

#### Tableau de Bord de l'Héritage

##### Vue d'Ensemble
- **Statut** : Indicateur visuel de l'état de l'héritage (complet, en cours, incomplet).
- **Actifs numériques** : Nombre total d'actifs associés à l'héritage.
- **Volontés** : Nombre de volontés créées.
- **Légataires** : Liste des personnes désignées comme légataires.

##### Sections Principales
- **Mes Volontés** : Liste des volontés créées par l'utilisateur.
- **Mes Légataires** : Personnes désignées comme légataires dans les volontés d'autres membres.
- **Actifs Associés** : Contenus numériques liés à l'héritage.
- **Événements Futurs** : Événements (naissance, anniversaire) qui pourraient déclencher des transmissions.

#### Création d'une Volonté Numérique

##### Assistant de Création
- **Étapes guidées** :
  1. **Introduction** : Explication du concept de volonté numérique.
  2. **Identification** : Qui crée la volonté.
  3. **Actifs** : Quels contenus inclure.
  4. **Légataires** : Qui recevra quoi.
  5. **Conditions** : Y a-t-il des conditions (âge, événements).
  6. **Validation** : Résumé et confirmation.

##### Formulaire de Volonté
- **Nom de la volonté** : Champ texte pour identifier la volonté.
- **Description** : Champ texte optionnel pour détails.
- **Statut** : Actif/Inactif avec switch.

##### Sélection des Actifs
- **Navigateur d'actifs** :
  - **Messagerie** : Historique des conversations, messages spécifiques.
  - **Médias** : Photos, vidéos, albums, documents.
  - **Événements** : Calendrier, événements spécifiques.
  - **Arbre Généalogique** : Informations généalogiques, biographies.
  - **Jeux** : Scores, parties sauvegardées.
- **Glisser-déposer** : Pour ajouter des actifs à la volonté.
- **Sélection multiple** : Cases à cocher pour ajouter plusieurs actifs.

##### Désignation des Légataires
- **Liste des membres du Hoomi** : Avec photos et noms.
- **Rôles** :
  - **Bénéficiaire** : Reçoit les actifs.
  - **Exécuteur** : Gère l'exécution de la volonté.
  - **Gardien** : Veille à ce que les conditions soient respectées.
- **Parts** : Pourcentage ou quantité spécifique pour chaque bénéficiaire.
- **Conditions** : Âge, événement, mot de passe.

#### Gestion des Volontés

##### Liste des Volontés
- **Grille** :
  - Nom de la volonté
  - Date de création
  - Statut
  - Nombre de légataires
  - Nombre d'actifs
- **Filtres** : Par statut, par date, par nombre d'actifs.

##### Détails d'une Volonté
- **Actifs inclus** : Liste avec miniatures et descriptions.
- **Légataires** : Liste avec rôles et parts.
- **Conditions** : Détail des conditions d'exécution.
- **Historique** : Journal des modifications, activations, exécutions.
- **Actions** :
  - **Modifier** : Changer les actifs, les légataires, les conditions.
  - **Activer/Désactiver** : Changer le statut.
  - **Dupliquer** : Créer une nouvelle volonté basée sur celle-ci.
  - **Supprimer** : Effacer définitivement la volonté.

### Mobile (iOS/Android)

#### Accès Rapide

##### Raccourci Principal
- **Onglet "Héritage"** : Dans la barre de navigation principale.
- **Carte de statut** : Résumé visuel de l'état de l'héritage.

##### Création Mobile
- **Assistant simplifié** :
  - **Étapes réduites** : Moins d'étapes que sur web.
  - **Sélection rapide** : Modèles prédéfinis (testament complet, cadeau d'anniversaire, etc.).
- **Capture intégrée** : Possibilité de prendre une photo ou une vidéo directement dans le flux de création.

#### Gestion sur Mobile

##### Notifications
- **Rappels** : Notifications pour mettre à jour l'héritage.
- **Événements déclencheurs** : Alertes quand un événement pourrait affecter l'héritage.
- **Demandes de legs** : Notification quand on est désigné comme légataire.

##### Interface de Gestion
- **Liste verticale** : Pour les petits écrans.
- **Cartes interactives** : Pour chaque volonté avec actions rapides (swipe).
- **Recherche** : Pour trouver rapidement une volonté par nom.

## Types de Volontés

### 1. Testament Numérique Complet

#### Description
- **Portée** : Ensemble complet des biens numériques d'un membre.
- **Actifs** :
  - Historique de messagerie complet.
  - Galerie multimédia entière.
  - Documents personnels.
  - Informations généalogiques.
- **Légataires** :
  - **Principal** : Généralement un conjoint ou un enfant.
  - **Secondaire** : Autres enfants, fratrie.
  - **Exécuteur** : Personne de confiance pour gérer l'exécution.
- **Conditions** :
  - **Événement** : Décès confirmé.
  - **Délai** : Attente de X jours pour confirmation.

### 2. Legs Spécifiques

#### Description
- **Portée** : Actifs spécifiques à des bénéficiaires spécifiques.
- **Exemples** :
  - **Album de photos d'enfance** : Pour un parent.
  - **Recettes de famille** : Pour une fille passionnée de cuisine.
  - **Collection de musique** : Pour un fils musicien.
- **Actifs** :
  - Sélection spécifique de médias, documents.
- **Légataires** :
  - Personne directement liée à l'actif.
- **Conditions** :
  - **Âge** : Par exemple, à 18 ans.
  - **Événement** : Par exemple, mariage.

### 3. Volontés Événementielles

#### Description
- **Portée** : Actions déclenchées par des événements spécifiques.
- **Exemples** :
  - **Message d'anniversaire** : Envoyé automatiquement chaque année.
  - **Lettre de mariage** : Pour les futurs mariés.
  - **Souvenirs de vacances** : Partagés après un certain événement.
- **Actifs** :
  - Messages, vidéos, albums.
- **Légataires** :
  - Membres impliqués dans l'événement.
- **Conditions** :
  - **Date** : Anniversaire, date spécifique.
  - **Événement** : Création d'un événement dans le calendrier.

### 4. Volontés de Secours

#### Description
- **Portée** : Informations critiques à transmettre en cas d'urgence.
- **Actifs** :
  - **Mots de passe** : (Futur) Stockage sécurisé de mots de passe critiques.
  - **Informations médicales** : Allergies, groupes sanguins.
  - **Contacts d'urgence** : Liste de personnes à contacter.
- **Légataires** :
  - **Exécuteurs** : Personnes habilitées à accéder en urgence.
  - **Services** : Pompiers, hôpitaux (via intégration future).
- **Conditions** :
  - **Inactivité** : Si l'utilisateur ne se connecte pas pendant X jours.
  - **Déclenchement manuel** : Par un exécuteur en cas d'urgence.

## Fonctionnalités Avancées

### Gestion des Conditions

#### Types de Conditions
- **Temporelles** :
  - **Âge** : Légataire doit avoir un certain âge.
  - **Date** : À une date précise ou après une date.
- **Événementielles** :
  - **Naissance** : Dans le Hoomi.
  - **Mariage** : D'un membre spécifique.
  - **Décès** : D'un membre (confirmé).
- **Combinées** :
  - **ET** : Deux conditions doivent être remplies.
  - **OU** : Au moins une condition doit être remplie.

#### Interface de Configuration
- **Constructeur visuel** : Pour créer des conditions complexes sans code.
- **Prévisualisation** : Affichage de qui serait concerné par la condition.

### Contrôle d'Accès Progressif

#### Niveaux d'Accès
- **Consultation** : Possibilité de voir que l'actif existe.
- **Accès limité** : Aperçu partiel (miniature, extrait).
- **Accès complet** : Accès au contenu original.
- **Modification** : Possibilité de modifier l'actif (réservé aux créateurs).

#### Attribution par Volonté
- **Personnalisable** : Chaque volonté peut définir des niveaux d'accès différents pour différents légataires.
- **Héritage** : Possibilité de définir des niveaux d'accès par défaut pour les nouveaux actifs.

### Audit et Traçabilité

#### Journal des Activités
- **Création** : Qui a créé quoi et quand.
- **Modification** : Qui a modifié quoi et quand.
- **Accès** : Qui a accédé à quoi et quand.
- **Exécution** : Qui a exécuté quoi et quand.

#### Rapports
- **Rapport annuel** : Résumé des activités liées à l'héritage.
- **Rapport de sécurité** : Alertes et événements de sécurité.
- **Rapport de santé** : (Futur) Indicateurs sur la complétude de l'héritage.

## Intégration avec les Autres Systèmes Hoomi

### Messagerie

#### Historique dans l'Héritage
- **Archivage** : Possibilité d'inclure l'historique de conversation dans une volonté.
- **Messages spécifiques** : Sélection de messages particuliers à inclure.

#### Notifications d'Héritage
- **Alertes** : Notifications quand on est nommé légataire.
- **Mises à jour** : Notifications des changements dans les volontés qui nous concernent.

### Médias

#### Actifs Multimédias
- **Inclusion directe** : Médias, albums, documents inclus dans les volontés.
- **Liens** : Possibilité de lier des actifs sans les inclure physiquement.

#### Albums d'Héritage
- **Création automatique** : Albums spéciaux pour chaque volonté.
- **Partage** : Possibilité de partager l'album avec les légataires.

### Événements

#### Déclenchement Événementiel
- **Liaison** : Volontés liées à des événements du calendrier.
- **Création assistée** : Assistant pour créer une volonté lors de la création d'un événement.

#### Événements d'Héritage
- **Célébrations** : Événements pour célébrer les anniversaires d'ancêtres.
- **Commemorations** : Événements pour commémorer les décès.

### Arbre Généalogique

#### Personnes dans les Volontés
- **Légataires spécifiques** : Liaison directe avec des personnes de l'arbre.
- **Actifs généalogiques** : Inclusion d'informations généalogiques dans les volontés.

#### Héritage Généalogique
- **Transmission de l'arbre** : Volontés spécifiques pour transmettre l'arbre généalogique.
- **Souvenirs d'ancêtres** : Messages ou médias d'ancêtres à transmettre.

## Sécurité et Confidentialité

### Chiffrement de Bout en Bout

#### Données Sensibles
- **Contenus des volontés** : Textes, listes d'actifs, conditions.
- **Actifs inclus** : Médias, documents, messages spécifiques.
- **Métadonnées** : Informations sur les légataires, les conditions.

#### Implémentation
- **AES-256-GCM** : Pour le chiffrement des contenus.
- **RSA-4096** : Pour le chiffrement des clés symétriques.
- **Clés partagées** : Chaque légataire reçoit une version chiffrée de la clé nécessaire.

### Contrôle d'Accès Strict

#### Vérification
- **À chaque requête** : Vérification des permissions avant d'accéder aux données d'héritage.
- **Journalisation** : Enregistrement de toutes les actions sur les données d'héritage.

#### Permissions
- **Créateur** : Droits complets sur la volonté.
- **Légataire** : Droits selon les conditions définies.
- **Exécuteur** : Droits pour gérer l'exécution.
- **Gardien** : Droits pour vérifier les conditions.

### Sécurité des Contenus Critiques

#### Mots de Passe
- **(Futur)** **Stockage sécurisé** : Coffre-fort numérique pour mots de passe critiques.
- **(Futur)** **Partage conditionnel** : Possibilité de partager des mots de passe selon des conditions.

#### Informations Médicales
- **(Futur)** **Champ dédié** : Informations médicales critiques.
- **(Futur)** **Accès d'urgence** : Partage rapide avec services médicaux en cas d'urgence.

## Performance et Fiabilité

### Surveillance des Conditions

#### Service de Surveillance
- **Vérification régulière** : Workers qui vérifient périodiquement les conditions.
- **File d'attente** : Gestion des volontés prêtes à être exécutées.

#### Tolérance aux Pannes
- **Redondance** : Plusieurs workers pour éviter les points de défaillance unique.
- **Journalisation** : Enregistrement de toutes les vérifications et exécutions.

### Exécution Garantie

#### Mécanismes de Rejeu
- **Retry** : Tentatives multiples en cas d'échec d'exécution.
- **Dead Letter Queue** : Volontés dont l'exécution a échoué de manière permanente.

#### Notifications d'Échec
- **Alertes administrateurs** : En cas de problème persistant avec les exécutions.
- **Logs détaillés** : Pour diagnostiquer les problèmes.

## Accessibilité

### Interfaces Adaptées

#### Navigation au Clavier
- **Raccourcis** : Tab, Entrée, Flèches pour naviguer dans les formulaires.
- **Validation** : Entrée pour valider/soumettre.

#### Lecteurs d'Écran
- **Balises ARIA** : Pour décrire les éléments interactifs.
- **Labels** : Textes alternatifs pour les boutons et les champs.

### Options Visuelles

#### Personnalisation
- **Taille du texte** : Ajustement pour les utilisateurs malvoyants.
- **Contraste élevé** : Palettes de couleurs avec contraste élevé.
- **Animations réduites** : Option pour réduire les animations.

## Internationalisation

### Langues Supportées
- **Français** : Langue par défaut.
- **Anglais** : Traduction complète.
- **(Futur)** Espagnol : Traduction automatique avec vérification humaine.

### Formats Locaux
- **Dates** : Adaptation au format local (DD/MM/YYYY vs MM/DD/YYYY).
- **Monnaies** : (Futur) Pour les biens matériels.
- **Lois** : (Futur) Adaptation aux lois locales sur l'héritage numérique.

Le système de transmission d'héritage numérique d'Hoomi offre une solution complète, sécurisée et émotionnellement riche pour permettre aux familles de préserver et transmettre leur héritage numérique de manière organisée et contrôlée.