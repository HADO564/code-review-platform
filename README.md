# Code Review Platfrom

## Introduction
The following project is an attempt to learn a different tech stack for web development. It is currently a work in progress, and a means for me to learn the following languages and technologies

* GraphQL
* Go-lang
* Svelte
* Github Actions
* A monorepo to house them
* Integration of a postgres db (supabase)

## Tech used
* Uses Graphql. Uses the 99designs graphql library to generate schemas and convert them into Go Structs.
* Connected to Supabase. Uses Supabase auth to allow smooth usage of RLS.
* Uses basic Go modules for request handling and error management.
* Sveltekit based frontend used with Supabase Auth for token usage.
* Client-side and server-side rendering for optimal serving of webpages.
* \*Other additions will be mentioned later

## Features
* Collaborate with clients, co-workers, or friends on projects, sharing code, reviewing and commenting on the code.
* Create code snippets for sharing, similar to stack overflow
* UI will be inspired from github where a user will be able to see the snippets they are working on along with the snippets their friends are.
* Will add resend to send notifications via email
