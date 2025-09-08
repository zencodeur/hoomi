# Agents d'Intelligence Artificielle Familiale Hoomi

Je documente ici les différents agents d'intelligence artificielle qui composent le système IA familiale d'Hoomi, chacun spécialisé dans un domaine spécifique pour améliorer la vie familiale.

## Vue d'Ensemble

Les **Agents IA** d'Hoomi sont conçus pour être :
- **Spécialisés** : Chacun excelle dans un domaine particulier.
- **Collaboratifs** : Ils peuvent interagir et se compléter.
- **Adaptables** : Ils apprennent et s'ajustent aux spécificités de chaque famille.
- **Respectueux** : Ils agissent toujours dans le respect de la vie privée et de l'autonomie familiale.
- **Transparents** : Leurs actions et leurs décisions sont compréhensibles.

## Agent Mémoire (Memoria)

### Rôle Principal
Agir comme le **mémorisateur familiale**, stockant et restituant les informations importantes pour chaque membre et pour la famille dans son ensemble.

### Fonctionnalités

#### Mémorisation d'Informations Clés
- **Dates importantes** : Anniversaires, anniversaires de mariage, dates de décès.
- **Préférences personnelles** : Couleurs préférées, plats détestés, hobbies.
- **Allergies et intolérances** : Informations médicales critiques.
- **Faits biographiques** : Lieux de naissance, parcours scolaire, métiers.
- **Souvenirs partagés** : Vacances, événements mémorables, anecdotes.

#### Rappels Intelligents
- **Proactifs** : Rappels discrets avant les événements importants.
- **Contextuels** : Rappels basés sur le lieu, l'heure, la personne présente.
- **Personnalisés** : Adaptés au style de communication de chaque membre.

#### Assistance Contextuelle
- **Dans la messagerie** : Identifier les personnes mentionnées, fournir des informations contextuelles.
- **Dans le calendrier** : Suggérer des dates en se basant sur les disponibilités connues.
- **Dans les listes de tâches** : Rappeler les habitudes (ne pas oublier le lait, arroser les plantes).

### Modèle Cognitif

#### Structure des Connaissances
- **Triplets Sujet-Prédicat-Objet** : Pour représenter les faits ("Alice", "aime", "la tarte aux pommes").
- **Ontologie familiale** : Structure des relations (parent, enfant, frère, sœur, cousin).
- **Chronologie** : Organisation des faits dans le temps.

#### Apprentissage
- **Extraction d'informations** : Analyse des messages pour identifier les faits importants.
- **Confirmation** : Demande de confirmation aux utilisateurs pour les faits incertains.
- **Désambiguïsation** : Gestion des homonymes et des informations contradictoires.

### Capteurs
- **Messages texte** : Analyse des conversations.
- **Événements du calendrier** : Extraction de dates et de faits.
- **Profils utilisateurs** : Informations de base des membres.
- **Arbre généalogique** : Relations familiales et faits généalogiques.

### Actuateurs
- **Notifications discrètes** : Rappels dans la barre de notification.
- **Suggestions dans l'interface** : Dans les champs de texte, calendrier, etc.
- **Messages automatiques** : Envoi de messages pré-programmés (ex: "Joyeux anniversaire !").

### Mémoire Persistante
- **Base de faits** : Stockage chiffré des informations mémorisées.
- **Historique des rappels** : Pour éviter les doublons et analyser l'efficacité.
- **Préférences de rappel** : Personnalisation de la fréquence et du type de rappels.

## Agent Organisation (Ordo)

### Rôle Principal
Agir comme l'**assistant organisationnel** de la famille, aidant à planifier, structurer et automatiser la vie quotidienne.

### Fonctionnalités

#### Planification Assistée
- **Suggestions d'événements** : Basées sur les habitudes et les disponibilités.
- **Optimisation d'emplois du temps** : Aide à trouver des créneaux communs.
- **Gestion de conflits** : Identification et proposition de résolution de conflits d'agenda.

#### Automatisation de Tâches
- **Listes de courses intelligentes** : Génération basée sur les habitudes alimentaires et les événements à venir.
- **Routines quotidiennes** : Création et gestion de routines (lever, coucher, devoirs).
- **Tâches récurrentes** : Gestion des tâches qui reviennent régulièrement.

#### Gestion de Projets Familiaux
- **Découpage** : Aide à diviser les grands projets en tâches gérables.
- **Assignation** : Suggestion d'assignation basée sur les compétences et disponibilités.
- **Suivi** : Suivi de l'avancement et rappels des échéances.

### Modèle Cognitif

#### Structure des Connaissances
- **Modèles d'habitudes** : Représentation des routines quotidiennes de chaque membre.
- **Préférences organisationnelles** : Matin/soir, week-end/jour de semaine.
- **Compétences** : Identification des compétences de chaque membre (cuisine, bricolage, devoirs).

#### Apprentissage
- **Analyse de comportement** : Observation des habitudes pour prédire les besoins.
- **Feedback** : Apprentissage à partir des retours des utilisateurs sur les suggestions.
- **Adaptation** : Ajustement des modèles en fonction des changements de routine.

### Capteurs
- **Emplois du temps** : Données du calendrier.
- **Messages** : Identification des demandes implicites d'organisation.
- **Capteurs domestiques** : (Futur) Données des objets connectés (frigo, machine à laver).
- **Historique d'achats** : (Futur) Pour anticiper les besoins (papier toilette, produits ménagers).

### Actuateurs
- **Création d'événements** : Ajout automatique d'événements au calendrier.
- **Listes de tâches** : Création et mise à jour de listes.
- **Notifications** : Rappels pour les tâches et les événements.
- **Suggestions dans les conversations** : Proposer des dates pour un événement.

### Mémoire Persistante
- **Modèles d'habitudes** : Stockage des routines familiales.
- **Historique de projets** : Pour améliorer les suggestions futures.
- **Préférences d'organisation** : Personnalisation des méthodes d'organisation.

## Agent Connexion (Vinculo)

### Rôle Principal
Agir comme le **facilitateur de liens familiaux**, encourageant les interactions, les moments partagés et le renforcement des relations entre les membres.

### Fonctionnalités

#### Suggestion de Moments de Connexion
- **Activités à faire ensemble** : Basées sur les intérêts communs et les disponibilités.
- **Jeux adaptés** : Recommandation de jeux en fonction de l'âge, du nombre de joueurs et de l'humeur.
- **Initiatives spontanées** : Suggestions pour casser la routine et créer des surprises.

#### Analyse des Dynamiques Relationnelles
- **Identification des liens forts/faibles** : Suivi des interactions pour identifier les membres moins en contact.
- **Médiation discrète** : (Futur) Aide à désamorcer les tensions en suggérant des activités apaisantes.
- **Renforcement des liens** : Proposer des activités pour rapprocher des membres spécifiques.

#### Gestion de l'Ambiance Familiale
- **Détection d'humeur** : (Subtile) Analyse des messages pour percevoir l'ambiance générale.
- **Suggestions d'ambiance** : Proposer de la musique, des vidéos, des jeux selon l'humeur perçue.
- **Moments de partage** : Créer des occasions de se rassembler (repas, soirées cinéma à la maison).

### Modèle Cognitif

#### Structure des Connaissances
- **Graphe relationnel** : Représentation des liens entre les membres (force, fréquence d'interaction).
- **Profils d'intérêts** : Cartographie des goûts de chaque membre.
- **Modèles de personnalité** : (Basiques) Pour adapter les suggestions (introverti/extraverti, calme/énergique).

#### Apprentissage
- **Analyse des interactions** : Fréquence, ton, sujets de discussion.
- **Feedback sur les suggestions** : Apprentissage à partir des réactions aux propositions.
- **Évolution des intérêts** : Suivi des changements de goûts et d'intérêts.

### Capteurs
- **Historique de messagerie** : Analyse des interactions textuelles.
- **Participation aux événements** : Qui participe à quoi et avec qui.
- **Partage de médias** : Qui partage quoi avec qui.
- **(Futur) Données biométriques** : Pouls, activité (via montres connectées) pour percevoir le stress ou l'énergie.

### Actuateurs
- **Suggestions dans la messagerie** : Proposer des activités ou des sujets de discussion.
- **Création d'événements** : Proposer et aider à créer des événements de connexion.
- **Notifications discrètes** : Pour suggérer un moment de qualité.
- **Interface des jeux** : Intégration avec le système de jeux pour recommander des parties.

### Mémoire Persistante
- **Historique des interactions** : Pour analyser les dynamiques relationnelles.
- **Préférences de connexion** : Types d'activités préférées pour se rassembler.
- **Historique des suggestions** : Pour améliorer les recommandations futures.

## Agent Histoire (Historia)

### Rôle Principal
Agir comme le **gardien de la mémoire familiale**, aidant à découvrir, structurer, préserver et célébrer l'histoire de la famille.

### Fonctionnalités

#### Découverte d'Histoires
- **"Souvenirs du jour"** : Partage d'événements qui se sont produits un an plus tôt.
- **Histoires croisées** : Relier des événements de différents membres pour raconter une histoire familiale.
- **Indices généalogiques** : Identifier des lacunes dans l'arbre généalogique et suggérer des pistes.

#### Structuration des Souvenirs
- **Chronologie personnalisée** : Création de frises chronologiques pour chaque membre ou pour la famille.
- **Thèmes** : Organisation des souvenirs par thèmes (vacances, mariages, naissances).
- **Génération de récits** : Transformation de faits en histoires narratives.

#### Préservation du Patrimoine
- **Assistance à l'enrichissement** : Suggérer d'ajouter des détails aux profils généalogiques.
- **Création de documents** : Aide à la création de documents familiaux (arbres, livres de souvenirs).
- **Interviews** : (Futur) Suggestion de poser des questions à des membres aînés pour préserver leurs témoignages.

### Modèle Cognitif

#### Structure des Connaissances
- **Chronologie familiale** : Organisation des événements dans le temps.
- **Thésaurus sémantique** : Mots-clés pour classifier les souvenirs (vacances, fête, réussite).
- **Modèles narratifs** : Structures de récits pour guider la génération d'histoires.

#### Apprentissage
- **Analyse de contenu** : Extraction de faits et d'émotions à partir des messages, médias et documents.
- **Identification de motifs** : Reconnaissance de cycles ou de traditions familiales.
- **Feedback narratif** : Apprentissage à partir des réactions aux histoires générées.

### Capteurs
- **Arbre généalogique** : Source principale d'informations historiques.
- **Médias** : Photos, vidéos, documents contenant des souvenirs.
- **Messages** : Histoires racontées dans les conversations.
- **Événements** : Dates et détails des événements marquants.

### Actuateurs
- **Fil d'actualité "Souvenirs"** : Partage régulier de souvenirs dans une section dédiée.
- **Suggestions d'enrichissement** : Dans les profils et l'arbre généalogique.
- **Génération de documents** : Création de pdf, d'albums ou de présentations.
- **Notifications discrètes** : Pour alerter de dates importantes dans l'histoire familiale.

### Mémoire Persistante
- **Base de connaissances historiques** : Ensemble structuré des faits et événements familiaux.
- **Corpus narratif** : Collection des histoires générées et validées.
- **Préférences de découverte** : Types d'histoires et de souvenirs préférés par la famille.

## Agent Sécurité (Tutela)

### Rôle Principal
Agir comme le **veilleur bienveillant**, surveillant discrètement le bien-être familial et prêt à intervenir de manière appropriée en cas de besoin.

### Fonctionnalités

#### Surveillance Bienveillante
- **Détection d'isolement** : Identifier les membres qui se retirent ou ne communiquent plus.
- **Signaux préoccupants** : (Très subtil) Analyse de mots-clés ou de changements de comportement.
- **Événements de vie critiques** : Suivre les événements marquants (maladie, deuil, déménagement).

#### Intervention Discrète
- **Alertes aux tuteurs** : Notifications discrètes et personnalisables pour les parents/gardiens.
- **Suggestions de soutien** : Proposer des actions ou des ressources en cas de besoin perçu.
- **Ressources** : Fournir des liens vers des ressources d'aide (psychologue, associations).

#### Prévention
- **Rappels de bien-être** : Suggestions discrètes pour prendre soin de soi.
- **Création de filets de sécurité** : (Futur) Permettre aux membres de définir des contacts à prévenir en cas d'urgence.

### Modèle Cognitif

#### Structure des Connaissances
- **Modèles de bien-être** : Profils de comportement normaux pour chaque membre.
- **Seuils d'alerte** : Paramètres définis pour déclencher des alertes.
- **Ressources d'aide** : Base de connaissances sur les ressources d'assistance disponibles.

#### Apprentissage
- **Modélisation comportementale** : Création de profils normaux d'activité et d'interaction.
- **Détection d'anomalies** : Identification des écarts significatifs par rapport aux comportements habituels.
- **Apprentissage des faux positifs** : Réduction des alertes intempestives grâce au feedback.

### Capteurs
- **Activité sur l'application** : Fréquence d'ouverture, messages envoyés/reçus.
- **Contenu des messages** : (Chiffré, analyse locale) Pour détecter des mots-clés préoccupants.
- **Événements du calendrier** : Changements brutaux d'emploi du temps.
- **(Futur) Capteurs de santé** : Données de montres connectées (sommeil, activité).

### Actuateurs
- **Notifications discrètes** : Aux membres désignés (parents, grands-parents).
- **Suggestions dans la messagerie** : Messages subtils pour engager la conversation.
- **Suggestions d'activités** : Proposer des activités relaxantes ou sociales.
- **Fourniture de ressources** : Liens vers des sites ou numéros d'assistance.

### Mémoire Persistante
- **Profils de bien-être** : Modèles de comportement normal pour chaque membre.
- **Historique des alertes** : Pour améliorer la précision et éviter les répétitions.
- **Préférences de surveillance** : Niveau de surveillance souhaité par chaque membre.

## Collaboration entre Agents

### Bus d'Événements

#### Communication
- **Événements IA** : Les agents publient des événements (ex: "nouveau souvenir découvert", "membre isolé détecté").
- **Souscriptions** : Les agents s'abonnent aux événements qui les intéressent.
- **Réactions** : Les agents réagissent aux événements pour coordonner leurs actions.

#### Exemples
- **Memoria** découvre l'anniversaire d'un membre → Publie un événement.
- **Ordo** reçoit l'événement → Crée un rappel pour acheter un cadeau.
- **Vinculo** reçoit l'événement → Suggère une activité spéciale pour l'occasion.

### Partage de Connaissances

#### Pool de Connaissances
- **Base de faits commune** : Les agents partagent certaines connaissances (liens familiaux, dates).
- **API d'accès** : Chaque agent peut requérir des informations auprès des autres agents.

#### Exemples
- **Historia** enrichit le profil d'un membre → Met à jour la base de connaissances.
- **Memoria** et **Ordo** peuvent accéder à cette information mise à jour.

### Coordination des Actions

#### Scénarios Complexes
- **Actions orchestrées** : Pour des situations nécessitant l'intervention de plusieurs agents.
- **Planification** : Création de plans d'action impliquant plusieurs agents.

#### Exemple
- **Scenario** : Préparation d'un anniversaire surprise.
  1. **Memoria** identifie la date.
  2. **Ordo** planifie les tâches.
  3. **Vinculo** suggère des activités secrètes.
  4. **Tutela** veille à ce que le secret soit bien gardé (pas de fuite dans les messages).

Cette architecture d'agents spécialisés et collaboratifs permet de créer une IA familiale riche, adaptative et centrée sur le bien-être et la cohésion familiale.