# PHP Dinero REST wrapper

This is a PHP wrapper for Dinero. 

## Installation

1. Require using composer

```
composer require danijwilliams/php-dinero
```

### Requirements

* PHP +5.6

### Getting started

1. [Apply as a developer](https://api.dinero.dk/docs#apply) at Dinero

2. Get your client id and secret

3. Find the organisation id when logged into Dinero (bottom left)

![](https://www.dropbox.com/s/ovwgzkmuu325lco/Screenshot%202017-08-28%2012.53.19.png?raw=1&dl=0)

4. Create an API key inside Dinero

5. Utilize the wrapper as below

``` php
 $dinero = new \danijwilliams\Dinero\Dinero( $clientId, $clientSecret );
 $dinero->auth( $apiKey, $orgId ); // this WILL send a request to the auth API.
 
 $contacts = $dinero->contacts()->perPage(10)->page(2)->get();
 
 // Do something with the contacts.
```

``` php
 $invoices = $dinero->invoices()->all();
```

``` php
 $products = $dinero->products()->deletedOnly()->all();
```

You can also use an old auth token, if you dont want to auth everytime you setup a new instance of Dinero.

``` php
 $dinero = new \danijwilliams\Dinero\Dinero( $clientId, $clientSecret );
 $dinero->setAuth($token, $orgId); // this will NOT send a request to the auth API.
 
 $products = $dinero->products()->deletedOnly()->all();
```

## Usage

### Creating Contacts

``` php
// Create Instance
$dinero = new \danijwilliams\Dinero\Dinero( $clientId, $clientSecret );

// Auth to a Dinero account
$dinero->auth( $apiKey, $orgId );
 
// Create the contact
$contact = $dinero->contacts()->create([ 'IsPerson' => true, 'Name' => 'Test', 'CountryKey' => 'DK' ]);

// if the request succeeded, $contact will be a \LasseRafn\Dinero\Models\Contact object.
```
