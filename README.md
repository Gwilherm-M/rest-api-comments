# Serveur REST

Serveur Rest permettant de stocker et gérer une chaine de commentaires liés à un id d'asset.

***

## Instalation

Language : go 1.19.2

Framework : mux / mongo-driver

```sh
go get -u github.com/gorilla/mux
```

```sh
go get go.mongodb.org/mongo-driver/mongo
```

## Initialisation

Avant de lancer le serveur.

* Avoir accés a une bdd de type MongoDB.
* Créer la variable d'environnement MONGODBPATH et lui assigner l'adresse du serveur MongoDB. Ex.:

```sh
MONGODBPATH = 'mongodb://localhost:27017'
```

## URL Path / Méthode

* `/`
* `/comment/` Permets de créer un commentaire.
  * Méthode: `POST`
    * Body: `assetId`, `text`
* `/comment/{id}` Permets de récupérer, supprimer ou updater un commentaire.
  * Méthode: `GET`, `DELETE`, `PUT`
    * Body pour `PUT`: `text`
* `/assetsComments/` Permets de récupérer la liste de tous les assetId.
  * Méthode: `GET`
* `/assetsComments/{assetId}` Permets de récupérer le fil de commentaire pour un assetId.
  * Méthode: `GET`
  * Query: `limit` *int*, `skip` *int*, `createAt` *string* __asc | des__, `updateAt` *string* __asc | des__

Example:

```url
/assetsComments/{assetId}?limit=5&skip=1&createAt=asc
```

L'attribut `limit` permet de limiter le nombre d'éléments renvoyés.

L'attribut `skip` permet de prendre les éléments suivants.

L'attribut `createAt`, __asc | des__ permet d'ordonner la liste dans l'ordre croissant ou décroissant selon leur date de création.

L'attribut `updateAt`, __asc | des__ permet d'ordonner la liste dans l'ordre croissant ou décroissant selon leur date d'update.

> Note:
>
> Par défaut `/assetsComments/{assetId}` renvoie tous les commentaires liés à un assetId, dans un ordre décroissant selon leur date de création.

***

## Excercice de gestion de projet

### Si un utilisateur ne peut pas se connecter à l'application

1. Je me connecte à l'application avec un autre compte utilisateur.
   * Cela me permet de savoir si le problème est commun à d'autres comptes.

2. Si j'ai pu me connecter. Je me reconnecte avec un compte admin.
   * Ainsi je vais pouvoir identifier le compte de l'utilisateur.
   * Voir sa dernière connexion. Éventuellement l'ip à ce moment-là.
   * Regarder à quand remonte la dernière modification de son mot de passe.

     * L'utilisateur a oublié son password.
     * Son compte a peut-être été piraté.
   * Si rien d'anormal est visible sur le compte, je propose à l'utilisateur de:
     * Vérifier s' il y a des mises à jour à effectuer sur son système ou pour le navigateur.
     * Nettoyer les données de navigation de son navigateur.
     * Eventuellement de tester avec un autre navigateur.
     * Si rien n'y fait, lui proposer un changement de password en suivant les procédures en vigueur.

3. Si je n'ai pas pu me connecter avec un autre compte utilisateur, je teste avec le compte admin.
    * Si rien n'y fait et qu'aucun compte ne me permet de me connecter à l'application.
    * Je regarde les logs côté serveur et j'identifie si le problème est lié au code ou à la BDD.
      * Je regarde l'état de la BDD. Je m'assure que le compte Admin existe. Si jamais la BDD est vide, envisager de remettre la dernière sauvegarde en date. Je rédige un ticket d'alerte pour l'équipe.
    * Si aucun problème notable n'est visible sur la BDD, je regarde la date de la dernière release mise en prod.
      * Dans l'éventualité qu'une mise en production d'une nouvelle release aurait été faite la veille (peut-être buggé) remettre en prod la version précédente, analyser la problématique ou bug dans un environnement de développement. Je rédige un nouveau ticket tout en spécifiant le bug dans le code.  
