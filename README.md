# r314dash

![r314dash Logo](static/images/favicon.ico)

## Overview

r314dash is a lightweight, web-based dashboard interface designed to manage multiple webui using iframes. This project is particularly useful for environments utilizing reverse proxies like Traefik or Nginx, allowing users to seamlessly switch between various applications hosted on different subpaths or domains. 

## Features

- **Multi-iframe Support**: Display multiple subpaths simultaneously using iframes, enabling quick switching between applications.
- **Configurable Interface**: Customize the dashboard with a configurable `config.yaml` file, allowing for easy management of links and settings.
- **User Information Display**: If configured, the dashboard can show the current active user based on headers provided by the reverse proxy.
- **Performance Stats (WIP)**: Future updates will include quick stats from a Prometheus instance, displaying CPU, RAM, Disk, and Network usage.
- **Custom Assets**: Mount custom front-end assets, such as images, at the path `/static`.
- **Stateless Design**: Built with a stateless approach, making it Kubernetes-friendly and easy to deploy.

## Project Structure

The main components of the r314dash interface include:

1. **Top Navigation Bar**:
   - Displays the site name (configurable).
   - Controls to reload, close, and open the currently selected iframe app in a new tab.
   - User menu to display the current active user (if configured).
   - Future integration for displaying quick stats from Prometheus.

2. **Side Navigation Bar**:
   - Contains links to activate the apps in their respective iframes.
   - Functions as a bookmark manager for quick access to frequently used applications.
   - The list of apps/links is configurable via the `config.yaml` file.

## Development

To build and run the main app you can use ```air``` with the provided ```.air.toml``` (if on windows set ```main.exe``` in cmd).

To watch changes to the templates and reload the css file you can use:
```bash
npm install
npm run watch
```

## Configuration

The configuration for r314dash is managed through the `config.yaml` file. This file contains all possible options for customizing the dashboard. The configuration is reloaded every time the application or container starts.

### Example `config.yaml`

```yaml
# TODO
```

### Acknowledgments

r314dash was inspired by Organizr, aiming to provide a faster, more lightweight, and Kubernetes-friendly alternative.

### TODO
    - doc:
        - screenshot images
        - deployment instructions
        - config options
    - app:
        - actually working light/dark mode switch
        - more efficient stats polling from prometheus
        - better icon management, like dashy
        - unit tests