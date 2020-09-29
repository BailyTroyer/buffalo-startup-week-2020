## Example Applications

## Overview

The following two folders hold our web application and backend API. Don't worry about the framework or the language of the services becuase that's not the focus here. We just want to run a process that's exposed on a port somewhere, or bundle static assets to be hosted on a CDN (Content Delivery Network).

## API - `/api`

In `/api` you'll find a basic HTTP request Multiplexer written in Go, which is a fancy term for a request router and dispatcher for matching incoming requests to their respective handler; i.e. an endpoint that matches to a handler, or, function. 

Example:

`/api/v1/endpoint` -> `func endpoint(...) { ... }`

## Web App - `/bflobox`

In `/bflobox` you'll find a basic webapp written in React with a few screens, a router and a simple redux store. From there you'd add the business logic for managing users, orders, etc. 

Some companies separate the marketing site from the webapp, however that's a matter of personal preference, and for this case we'll just combine the two.

**Note** The file structure and other explanations for the repo can be found in the corresponding services' directories `README.md`. As always if you have questions or want to just chat about anything feel free to hit me up on `@TODO add email,slack,Linkedin,Discord,Steam`