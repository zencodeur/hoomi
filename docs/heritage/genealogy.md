# Arbre Généalogique Vivant Hoomi

Je documente ici le système d'arbre généalogique vivant d'Hoomi, qui permet aux familles de créer, enrichir et explorer leur histoire familiale de manière interactive et collaborative.

## Vue d'Ensemble

L'**Arbre Généalogique Vivant** d'Hoomi est conçu pour être :
- **Interactif** : Navigation fluide et intuitive dans l'arbre.
- **Collaboratif** : Contribution de tous les membres du Hoomi.
- **Multimédia** : Enrichi avec des photos, vidéos, documents et histoires.
- **Évolutif** : Mise à jour en temps réel avec les événements de la vie.
- **Signifiant** : Allant bien au-delà d'une simple liste de noms et de dates.

## Interface Utilisateur

### Web (Next.js)

#### Vue Arbre

##### Présentation
- **Graphique interactive** : Utilisation d'une bibliothèque de visualisation (D3.js, Sigma.js) pour un arbre dynamique.
- **Nœuds personnalisés** : Chaque personne est représentée par un nœud avec :
  - Photo de profil.
  - Nom.
  - Dates de naissance/décès (si renseignées).
  - Indicateurs visuels (mariage, enfants, nouveaux contenus).
- **Liens dynamiques** : Lignes reliant les personnes avec des couleurs ou styles différents selon le type de relation.

##### Navigation
- **Zoom** : Molette de souris ou geste de pincement pour zoomer/dézoomer.
- **Pan** : Cliquer-glisser ou faire glisser avec deux doigts pour se déplacer dans l'arbre.
- **Focus** : Double-clic sur une personne pour la centrer et la mettre en évidence.
- **Recherche** : Barre de recherche pour trouver rapidement une personne.

##### Interactions
- **Clic sur un nœud** : Ouvre un panneau latéral avec les détails de la personne.
- **Clic sur un lien** : Affiche les détails de la relation (date de mariage, etc.).
- **Menu contextuel** : Clic droit pour ajouter un parent, un enfant, un conjoint.

#### Panneau de Détails

##### Informations de Base
- **Photo de profil** : Grande taille avec options de modification.
- **Nom complet** : Prénom, nom, surnom.
- **Dates clés** : Naissance, mariage, décès avec gestion de l'incertitude (vers 1920, avant 1950).
- **Âge** : Calculé automatiquement ou saisi manuellement.

##### Biographie
- **Éditeur WYSIWYG** : Pour rédiger une biographie riche.
- **Chronologie** : Liste des événements de la vie (naissance, mariage, enfants, décès, accomplissements).
- **Liens vers d'autres personnes** : Mentions des membres de la famille dans la biographie.

##### Médias
- **Galerie** : Grille des photos et vidéos associées à cette personne.
- **Documents** : Liste des documents (lettres, certificats, journaux).
- **Enregistrements** : Fichiers audio/vidéo (interviews, messages).
- **Ajout de médias** : Glisser-déposer ou sélection depuis la galerie Hoomi.

##### Relations
- **Famille immédiate** : Parents, fratrie, conjoints, enfants avec photos et liens rapides.
- **Famille élargie** : Cousins, oncles, tantes, belle-famille.
- **Relations spéciales** : Parrains, marraines, parrainages.

##### Histoires et Souvenirs
- **Histoires racontées** : Par cette personne ou à son sujet.
- **Souvenirs partagés** : Moments mémorables mentionnant cette personne.
- **Création de nouvelles histoires** : Bouton pour ajouter une histoire liée à cette personne.

### Mobile (iOS/Android)

#### Vue Arbre

##### Présentation
- **Vue simplifiée** : Arbre linéaire ascendant/descendant pour une navigation facile.
- **Cartes** : Chaque personne dans une carte empilée verticalement.
- **Indicateurs visuels** : Badges pour les nouveaux contenus, les médias.

##### Navigation
- **Swipe** : Gauche/droite pour naviguer entre les générations.
- **Tap** : Sur une personne pour voir ses détails.
- **Pinch** : Pour zoomer (si vue graphique disponible).

#### Vue Détails

##### Disposition
- **Défilement vertical** : Toutes les informations dans une seule vue défilante.
- **Sections repliables** : Biographie, médias, relations, histoires.

##### Actions
- **Menu flottant** : Pour les actions rapides (modifier, ajouter un parent, partager).
- **Partage natif** : Intégration avec le système de partage du téléphone.

## Fonctionnalités Collaboratives

### Contribution

#### Ajout de Personnes
- **Formulaire complet** : Champs pour toutes les informations (nom, dates, biographie).
- **Import depuis les contacts** : Possibilité d'importer des personnes depuis le carnet d'adresses.
- **Création à partir d'un mariage** : Ajout automatique des conjoints lors de l'ajout d'un mariage.

#### Édition
- **Permissions** : Configuration de qui peut modifier quelles informations.
- **Historique des modifications** : Suivi des changements avec auteur et date.
- **Restauration** : Possibilité de restaurer une version antérieure.

#### Validation
- **Système de suggestions** : Les membres peuvent suggérer des modifications qui doivent être approuvées.
- **Vote** : (Futur) Pour les informations incertaines, vote des membres pour trouver un consensus.

### Enrichissement

#### Médias
- **Ajout depuis la galerie Hoomi** : Lien rapide vers les médias existants.
- **Upload direct** : Possibilité d'uploader de nouveaux médias directement dans le profil.
- **Légendage** : Ajout de légendes aux médias avec identification des personnes présentes.

#### Histoires
- **Création de récit** : Éditeur pour raconter une histoire liée à une personne.
- **Liaison à des événements** : Association de l'histoire à un événement de la chronologie.
- **Médias dans les histoires** : Insertion de photos, vidéos, documents dans les récits.

#### Documents
- **Scan de documents** : (Futur) Utilisation de la caméra pour scanner des documents anciens.
- **OCR** : (Futur) Reconnaissance optique de caractères pour rendre les documents recherchables.
- **Organisation** : Classement des documents par type, date, personne.

## Intégration avec les Autres Systèmes Hoomi

### Messagerie

#### Mentions
- **@mention** : Possibilité de mentionner des personnes de l'arbre dans les messages.
- **Lien automatique** : Les noms de personnes dans les messages sont automatiquement liés à leur profil.

#### Partage d'histoires
- **Publication d'histoires** : Partage d'une histoire de l'arbre comme message dans la conversation du Hoomi.
- **Discussion** : Création de fils de discussion dédiés à une personne ou un événement.

### Médias

#### Galerie Généalogique
- **Albums automatiques** : Création d'albums par personne, par couple, par génération.
- **Reconnaissance faciale** : (Futur) Identification automatique des personnes sur les photos.

#### Médias dans les Profils
- **Glisser-déposer** : Depuis la galerie vers un profil pour l'associer.
- **Taggage** : Identification des personnes présentes sur une photo.

### Événements

#### Événements de Vie
- **Ajout automatique** : Les événements (naissance, mariage, décès) créés dans le calendrier sont automatiquement ajoutés à l'arbre.
- **Chronologie synchronisée** : La chronologie d'une personne dans l'arbre est synchronisée avec les événements du calendrier.

#### Célébrations
- **Anniversaires** : Notification et événement automatique pour les anniversaires des personnes de l'arbre.
- **Commemorations** : Événements pour les dates importantes (anniversaire de mariage, décès).

### Héritage Numérique

#### Lien avec les Volontés
- **Personnes dans les volontés** : Liaison directe entre les personnes de l'arbre et les bénéficiaires des volontés numériques.
- **Actifs généalogiques** : Possibilité d'associer des éléments de l'arbre (biographies, photos) à des volontés.

## Recherche et Exploration

### Recherche Textuelle

#### Champs Recherchés
- **Noms** : Prénom, nom, surnom.
- **Biographies** : Contenu des biographies et des histoires.
- **Métadonnées** : Dates, lieux, professions.
- **Tags** : Mots-clés associés aux personnes.

#### Résultats
- **Pertinence** : Tri des résultats par pertinence.
- **Aperçu** : Extrait du texte contenant les termes recherchés.
- **Navigation rapide** : Clic sur un résultat pour aller directement à la personne.

### Filtres Avancés

#### Par Génération
- **Ancêtres** : Afficher uniquement les ancêtres d'une personne.
- **Descendants** : Afficher uniquement les descendants d'une personne.
- **Cousins** : Afficher les cousins d'un certain degré.

#### Par Époque
- **Plage de dates** : Afficher les personnes nées/vivantes pendant une période donnée.
- **Siècle** : Filtrer par siècle (XIXe, XXe, XXIe).

#### Par Localisation
- **Carte interactive** : Afficher les personnes sur une carte selon leur lieu de naissance/résidence.
- **Région/Pays** : Filtrer par lieu géographique.

#### Par Événement
- **Guerres** : Identifier les personnes nées/vivantes pendant une période de guerre.
- **Migrations** : Suivre les déplacements familiaux à travers les générations.

### Arbres Thématiques

#### Arbres Prédéfinis
- **Arbre Paternel/Maternel** : Focalisation sur une lignée spécifique.
- **Arbre des Femmes/Hommes** : Mettre en lumière les femmes ou les hommes de la famille.
- **Arbre des Mariages** : Mettre en avant les unions et les alliances familiales.

#### Arbres Personnalisés
- **Création** : Possibilité de créer des arbres personnalisés basés sur des critères spécifiques.
- **Sauvegarde** : Enregistrement des arbres personnalisés pour consultation ultérieure.

## Sécurité et Confidentialité

### Contrôle d'Accès

#### Permissions
- **Visibilité** : Qui peut voir les informations sensibles (dates de décès, causes, etc.).
- **Modification** : Qui peut modifier les informations.
- **Suppression** : Qui peut supprimer une personne ou une information.

#### Héritage des Permissions
- **Par Hoomi** : Les permissions de base sont héritées du Hoomi parent.
- **Personnalisables** : Possibilité d'affiner les permissions pour l'arbre généalogique.

### Chiffrement

#### Données Sensibles
- **Biographies détaillées** : Textes longs et intimes chiffrés.
- **Documents privés** : Lettres, journaux intimes chiffrés.
- **Informations médicales** : (Si ajoutées) Historique médical chiffré.

#### Gestion des Clés
- **Clés partagées** : Chaque membre du Hoomi reçoit une version chiffrée des clés nécessaires.
- **Renouvellement** : Possibilité de renouveler les clés de chiffrement.

## Performance et Optimisation

### Rendu de l'Arbre

#### Virtualisation
- **Nœuds virtuels** : Seuls les nœuds visibles sont rendus, pour des arbres très larges.
- **Niveaux de détail** : Affichage simplifié pour les parties éloignées du centre de l'attention.

#### Mise en Cache
- **Arbres fréquents** : Mise en cache des configurations d'arbres les plus demandées.
- **Métadonnées** : Caching des informations de base pour un affichage rapide.

### Chargement Progressif

#### Données
- **Détails au besoin** : Chargement des informations détaillées seulement quand l'utilisateur les consulte.
- **Miniatures** : Chargement d'abord de miniatures basses résolutions.

#### Médias
- **Lazy Loading** : Chargement des médias uniquement lorsqu'ils entrent dans le viewport.
- **Formats adaptatifs** : Servir des formats optimisés selon la bande passante.

## Accessibilité

### Navigation au Clavier

#### Raccourcis
- **Tab** : Naviguer entre les éléments interactifs.
- **Flèches** : Naviguer dans l'arbre (haut/bas pour génération, gauche/droite pour fratrie).
- **Entrée** : Ouvrir les détails d'une personne.
- **Échap** : Fermer les panneaux ouverts.

### Lecteurs d'Écran

#### Balises ARIA
- **Rôles** : Utilisation des rôles ARIA pour décrire la structure de l'arbre.
- **Labels** : Textes alternatifs détaillés pour les nœuds et les relations.
- **Notifications** : Annonces des changements dynamiques (ajout de personne, etc.).

### Options Visuelles

#### Personnalisation
- **Taille du texte** : Ajustement de la taille des textes dans l'interface.
- **Contraste élevé** : Palettes de couleurs pour les utilisateurs malvoyants.
- **Animations** : Possibilité de réduire ou désactiver les animations.

## Internationalisation

### Langues Supportées
- **Français** : Langue par défaut.
- **Anglais** : Traduction complète.
- **(Futur)** Espagnol : Traduction automatique avec vérification humaine.

### Formats Locaux
- **Dates** : Adaptation au format local (DD/MM/YYYY vs MM/DD/YYYY).
- **Nombres** : Format des milliers et des décimaires selon la locale.
- **Généalogie** : Termes spécifiques à la culture (oncle paternel/maternel).

Cette interface riche et collaborative de l'Arbre Généalogique Vivant permet aux familles Hoomi de créer une véritable chronique de leur histoire, tout en renforçant les liens entre les générations.