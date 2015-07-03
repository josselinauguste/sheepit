# TODO

1. Intégration continue (http://www.synbioz.com/blog/rugged_gem_git)
2. Push to deploy
3. Pipeline de déploiement (https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern)

## Features

- support RCov

## Initialisation

- afficher version (ajax)
- déclarer impacts sur mvx_private, PHP, doc, formation hotline
- ajouter les dates (création, validation)


## Validation

- compléter pipeline
- afficher les infos statiques (ex: changelog)
- persister les infos de validation
- consulter les infos une fois validé


## Déploiement

- Màj le changelog & la version
- tagger le commit de version
- send deployment email
- validation envoi comm externe


## Further

- gérer les PR sur tous les projets (pas seulement mvx)
- automatiser vérif des PRs
- préremplir changelog (git log?)
- deviner PRs (utiliser les merges)
- récupérer statut jenkins
- automatiser les déploiements


## Pipeline

Code review -> Rake test -> Déploiement en préprod -> smoke tests -> tests manuels (changelog) -> tests manuel aléatoires (happy path) -> supports (doc & hotline) -> vérification des métriques -> màj changelog & version & tag -> déploiement prod -> smoke tests -> tests manuels (changelog) -> tests manuel aléatoires (happy path) -> envoi des notifs
