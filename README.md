# README

## Overview

This project is a React application built with NextJS, Yarn for package management, Tailwind for styles and DaisyUI for components, and a Go API.

## Thought process

The tools were chosen with simplicity and speed of development in mind. The API is a simple Go application. The React application uses NextJS for SSR and fetches data from the API to format and display on the feed screen, which was inspired by the Instagram feed view.

## Getting Started

1. Clone the repository:

```bash
git clone https://github.com/Maverick5000/moneylion-coding-challenge.git
```

2. Navigate to the project directory:

```bash
cd moneylion-coding-challenge/
```

## Running the API and UI Development environments

To start the development server, run:

```bash
make up
```

Or

```bash
docker-compose up
```

This will launch the application in development mode. Open http://localhost:3000 in your browser to view it.

## Running Tests

### Unit Tests on the API

To run unit tests on the API, use the following command:

```bash
make test
```
