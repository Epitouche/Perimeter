# Nuxt POC
Nuxt is a fullstack Vue.js framework.

## Table Of Contents

- [Nuxt POC](#nuxt-poc)
  - [Table Of Contents](#table-of-contents)
  - [Other Documents](#other-documents)
  - [Description](#description)
    - [Prerequisites](#prerequisites)
      - [Nodejs \& npm](#nodejs--npm)
    - [Run and View](#run-and-view)
    - [Functionalities Tested](#functionalities-tested)
      - [Modules:](#modules)
      - [Components:](#components)
      - [Utils:](#utils)
      - [NuxtUI Components:](#nuxtui-components)
  - [Further information](#further-information)

## Other Documents

[Main Documentation](../../../README.md)
[Comparative Study](../README.md)

## Description

A simple Nuxt project to test it's capabilities and efficiency.
The project is a quiz that fetches data from an api to display it's results.

### Prerequisites

Command lines for installation (linux)

#### Nodejs & npm
**Fedora:**
- sudo dnf install nodejs -y
**Ubuntu:**
- sudo apt install nodejs
- sudo apt install npm

Nodejs must be v18.0.0 or newer. You can check your version with the command line: `node -v`

### Run and View

Run command line: `npm run dev` for a dev server.
Then navigate to `http://localhost:3000` to see the project web page(s).

### Functionalities Tested

#### Modules:
- Nuxt UI
- TailwindCSS
- Pinia

#### Components:
- NuxtPage
- NuxtLayout
- NuxtLink

#### Utils:
- definePageMeta
- $fetch

#### NuxtUI Components:
- UButton
- UPopover

## Further information

- [Nuxt Website](https://nuxt.com/)
- [Tailwind for Nuxt](https://tailwindcss.nuxtjs.org/getting-started/installation)
